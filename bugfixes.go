package bugfixes

import (
  "fmt"
  "go-bugfixes/formatter"
  "io"
  "log"
  "os"
  "sync"
)

const (
  LOG = 0
  INFO = 1
  WARN = 2
  ERROR = 3
  FATAL = 4
)

var logLevels = map[string]int {
  "LOG": LOG,
  "INFO": INFO,
  "WARN": WARN,
  "ERROR": ERROR,
  "FATAL": FATAL,
}

var initalized uint32
var mutex sync.Mutex

type LoggerInterface interface {
  AddHandler(handler *io.Writer)
  Log(message string, a ...interface{})
  Info(message string, a ...interface{})
  Warn(message string, a ...interface{})
  Error(message string, a ...interface{})
  Fatal(message string, a ...interface{})
  SetFormatter(template string, a ...interface{})
}

type rootLogger struct {
  handlers []*LogHandler
  formatter string
}

var instance *rootLogger
var once sync.Once

func get() *rootLogger {
  once.Do(func() {
    instance = &rootLogger{}
  })

  return instance
}

type LogHandler struct {
  Writer *io.Writer
  Level int
  Formatter *formatter.LogFormatter
}

func (h *LogHandler) SetFormatter(pattern string) (*LogHandler, error) {
  handlerFormatter, err := formatter.New(pattern)
  if err != nil {
    log.Printf("bugfixes package failed to create and set new formatter from pattern: %s", pattern)
    return h, err
  }
  h.Formatter = handlerFormatter

  return h, nil
}

func (h *LogHandler) emit(message *formatter.LogMessage) {
  logFomatter := *h.Formatter
  logFomatter.Format(h.Writer, message)
}

func NewHandler(level interface{}, w io.Writer) (*LogHandler, error) {
  switch level.(type) {
  case int:
    if level.(int) < LOG || level.(int) > FATAL {
      log.Print("Level is not supported, use LOG/INFO/WARN/ERROR/FATAL.")
    }
    return &LogHandler{Writer: &w, Level: level.(int), Formatter: formatter.DefaultFormatter}, nil
  case string:
    if _, ok := logLevels[level.(string)]; !ok {
      log.Print("Level is not supported, use LOG/INFO/WARN/ERROR/FATAL.")
    }
    return &LogHandler{Writer: &w, Level: logLevels[level.(string)], Formatter: formatter.DefaultFormatter}, nil
  default:
    return &LogHandler{Writer: &w, Level: LOG, Formatter: formatter.DefaultFormatter}, nil
  }
}

func (l *rootLogger) AddHandler(h *LogHandler) {
  l.handlers = append(l.handlers, h)
}

var defaultHandler, err = NewHandler(LOG, os.Stdout)

func formatLevel(levelno int, level, message string, args ...interface{}) *formatter.LogMessage {
  return &formatter.LogMessage{
    Message: fmt.Sprintf(message + "\n", args...),
    Level: level,
    LevelNum: levelno,
  }
}

func dispatchMessage(handlers []*LogHandler, levelNum int, level string, message string, args ...interface{}) {
  if len(handlers) == 0 {
    handlers = []*LogHandler{defaultHandler}
  }
  for _, handler := range handlers {
    if handler.Level <= levelNum {
      handler.emit(formatLevel(levelNum, level, message, args...))
    }
  }
}

func (l *rootLogger) Log(message string, a ...interface{}) {
  dispatchMessage(l.handlers, LOG, "LOG", message, a...)
}

func (l *rootLogger) Info(message string, a ...interface{}) {
  dispatchMessage(l.handlers, INFO, "INFO", message, a...)
}

func (l *rootLogger) Warn(message string, a ...interface{}) {
  dispatchMessage(l.handlers, WARN, "WARN", message, a...)
}

func (l *rootLogger) Error(message string, a ...interface{}) {
  dispatchMessage(l.handlers, ERROR, "ERROR", message, a...)
}

func (l *rootLogger) Fatal(message string, a ...interface{}) {
  dispatchMessage(l.handlers, FATAL, "FATAL", message, a...)
}

func AddHander(h *LogHandler) {
  get().AddHandler(h)
}

func Log(message string, a ...interface{}) {
  get().Log(message, a...)
}

func Info(message string, a ...interface{}) {
  get().Info(message, a...)
}

func Warn(message string, a ...interface{}) {
  get().Warn(message, a...)
}

func Error(message string, a ...interface{}) {
  get().Error(message, a...)
}

func Fatal(message string, a ...interface{}) {
  get().Fatal(message, a...)
}
