package bugfixes

import (
  "go-bugfixes/formatter"
  "io"
  "log"
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
