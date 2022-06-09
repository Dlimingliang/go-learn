package main

import (
	"go.uber.org/zap"
	"time"
)

func InitZapConfig() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./myproject.log",
		"stderr",
		"stdout",
	}
}

func main() {

	logger, _ := zap.NewProduction()
	//logger := zap.NewExample()
	//logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	url := "htts://github.com"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
