package gormx

import (
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
	LogZap bool
}

func newWriter(logZap bool, w logger.Writer) *writer {
	return &writer{
		Writer: w,
		LogZap: logZap,
	}
}
