package bugfixes

import (
	"github.com/bugfixes/go-bugfixes/formatter"
	"io"
)

type LogHandler struct {
	Writer *io.Writer
	Level int
	Formatter *formatter.LogFormatter
}

type rootLogger struct {
	handlers []*LogHandler
	formatter string
}