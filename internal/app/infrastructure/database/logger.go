package database

import (
	"context"
	"github.com/jibaru/home-inventory-api/m/logger"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
}

func NewGormLogger() *GormLogger {
	return &GormLogger{}
}

func (l *GormLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *GormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	logger.LogDBTransactionInfo(msg, data)
}

func (l *GormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	logger.LogDBTransactionWarn(msg, data)
}

func (l *GormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	logger.LogDBTransactionError(msg, data)
}

func (l *GormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	logger.LogDBTransactionTrace("sql", rows, elapsed, sql)
}
