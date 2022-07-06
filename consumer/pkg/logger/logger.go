package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"sync"
)

type Logger struct {
	*logrus.Entry
}

var (
	instance Logger
	once     sync.Once
)

func GetLogger(level string) Logger {
	once.Do(func() {
		logrusLevel, err := logrus.ParseLevel(level)
		if err != nil {
			log.Fatal(err)
		}
		l := logrus.New()
		l.Level = logrusLevel
		l.SetReportCaller(true)
		l.Formatter = &logrus.TextFormatter{
			DisableColors:    false,
			DisableTimestamp: false,
			FullTimestamp:    true,
		}
		l.SetOutput(os.Stdout)

		instance = Logger{logrus.NewEntry(l)}
	})
	return instance
}
