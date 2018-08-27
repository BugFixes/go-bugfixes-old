package formatter

import (
  "fmt"
  "io"
  "log"
  "runtime"
  "text/template"
  "time"
)

var stackLocation = 17

func asctime() string {
  return time.Now().Format("2006-01-02 01:02:03")
}

func created() int64 {
  return time.Now().Unix()
}

func filename() string {
  _, filename, _, ok := runtime.Caller(stackLocation)
  if !ok {
    return "<unknown>"
  }
  return filename
}

func lineno() int {
  _, _, lineno, ok := runtime.Caller(stackLocation)
  if !ok {
    return -1
  }
  return lineno
}

func linenoAndLineno() string {
  _, filename, lineno, ok := runtime.Caller(stackLocation)
  if !ok {
    return "<unknown>: -1"
  }

  return fmt.Sprintf("%s: %d", filename, lineno)
}

var functions = template.FuncMap{
  "asctime": asctime,
  "created": created,
  "filename": filename,
  "lineno": lineno,
  "fileline": linenoAndLineno,
}

type LogMessage struct {
  Message string
  Level string
  LevelNum int
}

type LogFormatter struct {
  Template *template.Template
}

func New(pattern string) (*LogFormatter, error) {
  template, err := template.New("logTemplate").Funcs(functions).Parse(pattern)
  if err != nil {
    return nil, err
  }
  return &LogFormatter{
    template,
  }, nil
}

func (logFormatter *LogFormatter) Format(writer *io.Writer, message *LogMessage) {
  err := logFormatter.Template.Execute(*writer, message)
  if err != nil {
    log.Printf("bugfixes package failed to emit %s to %v", err, writer)
  }
}

const DefaultFormatterPattern = "{{ asctime }}; {{ fileline }}; {{.Level}}; {{.Message}}"

var DefaultFormatter, _ = New(DefaultFormatterPattern)
