package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func main() {

	//格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//encoder := zapcore.NewJSONEncoder(encoderConfig)
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	//写入到哪
	file, _ := os.Create("./test.log")
	writeSyncer := zapcore.AddSync(file)

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) //指定级别
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()
	sugar := logger.Sugar()

	url := "htts://github.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
