package formatter

import "text/template"

type LogMessage struct {
	Message string
	Level string
	LevelNum int
}

type LogFormatter struct {
	Template *template.Template
}
