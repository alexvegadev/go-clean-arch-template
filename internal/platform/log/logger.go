package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &MercuryFormatter{DisableColors: true},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}

func Info(msg ...interface{}) {
	log.Info(msg)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args)
}

func Infoln(msg ...interface{}) {
	log.Infoln(msg)
}

func Warn(args ...interface{}) {
	log.Warn(args)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args)
}

func Warnln(msg ...interface{}) {
	log.Warnln(msg)
}

func Error(args ...interface{}) {
	log.Error(args)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args)
}

func Errorln(msg ...interface{}) {
	log.Errorln(msg)
}

func GetOut() io.Writer {
	return log.Out
}
