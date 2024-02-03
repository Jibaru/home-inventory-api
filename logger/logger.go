package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

type AppLogger struct {
	l *logrus.Logger
}

var appLogger *AppLogger

func init() {
	appLogger = NewAppLogger()
}

func NewAppLogger() *AppLogger {
	file, err := os.OpenFile("logs/app.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.999999999Z07:00",
	})
	logger.SetOutput(io.MultiWriter(file, os.Stdout))
	logger.SetLevel(logrus.TraceLevel)

	return &AppLogger{logger}
}

func LogRequest(
	method string,
	path string,
) {
	appLogger.l.WithFields(logrus.Fields{
		"context": "http",
		"method":  method,
		"path":    path,
	}).Info("request")
}

func LogDBTransactionTrace(
	msg string,
	rows int64,
	elapsed time.Duration,
	data interface{},
) {
	appLogger.l.WithFields(logrus.Fields{
		"context": "db",
		"rows":    rows,
		"elapsed": elapsed,
		"data":    data,
	}).Trace(msg)
}

func LogDBTransactionError(
	msg string,
	data interface{},
) {
	appLogger.l.WithFields(logrus.Fields{
		"context": "db",
		"data":    data,
	}).Error(msg)
}

func LogDBTransactionWarn(
	msg string,
	data interface{},
) {
	appLogger.l.WithFields(logrus.Fields{
		"context": "db",
		"data":    data,
	}).Warn(msg)
}

func LogDBTransactionInfo(
	msg string,
	data interface{},
) {
	appLogger.l.WithFields(logrus.Fields{
		"context": "db",
		"data":    data,
	}).Info(msg)
}

func LogInfo(
	msg string,
) {
	appLogger.l.WithFields(logrus.Fields{
		"context": "app",
	}).Info(msg)
}

func LogError(
	err error,
) {
	appLogger.l.WithFields(logrus.Fields{
		"context": "app",
	}).Error(err)
}
