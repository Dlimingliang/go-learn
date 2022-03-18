package main

import (
	"go.uber.org/zap"
	"time"
)

func InitZapConfig() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./myproject.log",
		"stderr",
		"stdout",
	}
	return cfg.Build()
}

func main() {

	logger, _ := InitZapConfig()
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

	logger.Info("Failed to fetch URL", zap.String("url", url), zap.Int("attempt", 3))
}
