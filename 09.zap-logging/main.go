package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	url := "https://www.google.com"
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	sugar.Errorw("failed to process message",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Errorf("Failed to process message: %s", url)
	// sugar.Fatal("Fatal message.  Will call os.Exit")
	// sugar.Info("Info after fatal")
}
