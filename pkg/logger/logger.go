package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	mainLog "log"
)

var log *zap.Logger

func init() {
	cfg := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{"stdout"},
	}

	var err error
	if log, err = cfg.Build(); err != nil {
		mainLog.Fatalln(err)
	}
}

// Info log
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	err := log.Sync()
	if err != nil {
		return
	}
}

// Info log
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))

	log.Info(msg, tags...)
	se := log.Sync()
	if se != nil {
		return
	}
}
