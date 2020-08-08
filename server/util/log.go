package util

import "fyne.io/fyne/widget"

type Log struct {
	entry *widget.Label
}

type LogType int

const (
	LOG_INFO LogType = 1 + iota
	LOG_ERROR
)

var MainLog *Log

func NewLog(entry *widget.Label) *Log {
	return &Log{entry: entry}
}

func (l *Log) AddLog(text string, logType LogType) {
	var prefix string
	switch logType {
	case LOG_INFO:
		prefix = "[INFO] "
	case LOG_ERROR:
		prefix = "[ERROR] "
	}
	l.entry.SetText(l.entry.Text + "\n" + prefix + text)
}
