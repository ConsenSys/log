package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	Json FormatterType = "json"
	Text FormatterType = "text"
)

const (
	Info  = LogLevel(logrus.InfoLevel)
	Debug = LogLevel(logrus.DebugLevel)
	Warn  = LogLevel(logrus.WarnLevel)
	Error = LogLevel(logrus.ErrorLevel)
	Fatal = LogLevel(logrus.FatalLevel)
	Panic = LogLevel(logrus.PanicLevel)
)

var (
	defaultFormatter = &logrus.TextFormatter{
		ForceColors:      true,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	}
	defaultLevel   = logrus.DebugLevel
	defaultOptions = &Options{
		Level:     defaultLevel,
		Formatter: defaultFormatter,
	}
)

type FormatterType string

type LogLevel logrus.Level

type Options struct {
	Level     logrus.Level
	Formatter logrus.Formatter
}

type Option func(*Options)


func New(loggerName string, setters ...Option) *logrus.Entry {

	args := defaultOptions

	for _, setter := range setters {
		setter(args)
	}

	logger := &logrus.Logger{
		Formatter: args.Formatter,
		Level:     args.Level,
		Out:       os.Stdout,
	}

	return logrus.NewEntry(logger).WithFields(logrus.Fields{
		"logger": loggerName,
	})
}



func Level(level LogLevel) Option {
	return func(options *Options) {
		options.Level = logrus.Level(level)
	}
}

func Formatter(formatterType FormatterType) Option {
	return func(options *Options) {
		f := options.Formatter
		switch formatterType {
		case Json:
			f = &logrus.JSONFormatter{}
		case Text:
			f = &logrus.TextFormatter{
				ForceColors:      true,
				FullTimestamp:    true,
				QuoteEmptyFields: true,
			}
		}
		options.Formatter = f
	}
}
