// Package constants содержит все используемые в проекте константы
package constants

import (
	"go.uber.org/zap/zapcore"
)

// Логгер.
const (
	LogFile  string = "log.log"
	LogLevel        = zapcore.InfoLevel
)
