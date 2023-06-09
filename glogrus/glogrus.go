package glogrus

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func Setup(mode string) {
	Logger = logrus.New()
	if _, err := os.Stat("./log"); os.IsNotExist(err) {
		_ = os.MkdirAll("./log", os.ModePerm)
	}
	errorLogW, _ := os.OpenFile("./log/error.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o777)
	infoLogW, _ := os.OpenFile("./log/info.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o777)

	formatter := new(ErrorFormatter)
	Logger.SetFormatter(formatter)
	Logger.Out = os.Stdout

	var writerMap lfshook.WriterMap
	switch mode {
	case "production":
		Logger.Level = logrus.ErrorLevel
		writerMap = lfshook.WriterMap{
			logrus.PanicLevel: errorLogW,
			logrus.FatalLevel: errorLogW,
			logrus.ErrorLevel: errorLogW,
		}
	case "staging":
		Logger.Level = logrus.DebugLevel
		writerMap = lfshook.WriterMap{
			logrus.PanicLevel: infoLogW,
			logrus.FatalLevel: infoLogW,
			logrus.ErrorLevel: infoLogW,
			logrus.WarnLevel:  infoLogW,
			logrus.InfoLevel:  infoLogW,
		}
	case "debug":
		Logger.Level = logrus.DebugLevel
		writerMap = lfshook.WriterMap{
			logrus.PanicLevel: infoLogW,
			logrus.FatalLevel: infoLogW,
			logrus.ErrorLevel: infoLogW,
			logrus.WarnLevel:  infoLogW,
			logrus.InfoLevel:  infoLogW,
		}
	}

	lfHook := lfshook.NewHook(writerMap, formatter)
	Logger.Hooks.Add(lfHook)
}

type ErrorFormatter struct{}

func (f *ErrorFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	msg := fmt.Sprintf("[%s/%s]:%s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		entry.Level.String(),
		entry.Message)
	return []byte(msg), nil
}

func Debug(msg ...interface{}) {
	Logger.Debug(msg...)
}

func Info(msg ...interface{}) {
	Logger.Info(msg...)
}

func Error(msg string) {
	msg += string(debug.Stack())
	Logger.Error(msg)
}

func Fatal(msg string) {
	msg += string(debug.Stack())
	Logger.Error(msg)
}

func Panic(msg string) {
	msg += string(debug.Stack())
	Logger.Error(msg)
}

func Warn(msg string) {
	msg += string(debug.Stack())
	Logger.Warn(msg)
}
