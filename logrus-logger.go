package log

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	*logrus.Entry
}

func NewLogrusLogger(entry *logrus.Entry){
	return LogrusLogger{entry}
}

func (l LogrusLogger) Log(keyvals ...interface{}) error {
	if len(keyvals)%2 == 0 {
		fields := logrus.Fields{}
		for i := 0; i < len(keyvals); i += 2 {
			fields[fmt.Sprint(keyvals[i])] = keyvals[i+1]
		}
		l.WithFields(fields).Info()
	} else {
		l.Info(keyvals)
	}
	return nil
}
