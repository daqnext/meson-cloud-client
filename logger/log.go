package logger

import (
    "go.uber.org/zap"
)

var L *zap.SugaredLogger

func RegisterLogger(level string) {
    var zLogger *zap.Logger
    switch level {
    case "dev":
        zLogger, _ = zap.NewDevelopment()
    default:
        zLogger, _ = zap.NewProduction()
    }
    L = zLogger.Sugar()
}

// TODO: Register performance tuned logger
