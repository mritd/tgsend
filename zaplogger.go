package main

import (
	"log"
	"sync"

	"github.com/mritd/zaplogger"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

var logOnce sync.Once

func init() {
	logOnce.Do(func() {
		zc, err := zaplogger.NewConfig(zaplogger.ZapConfig{
			Encoder:      zaplogger.EncoderConsole,
			Level:        zaplogger.LevelInfo,
			StackLevel:   zaplogger.LevelError,
			TimeEncoding: zaplogger.TimeEncoderDefault,
		})
		if err != nil {
			log.Fatal(err)
		}
		logger = zaplogger.NewLogger(zc).Sugar()
	})
}
