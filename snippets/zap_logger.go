package snippets

import (
	"go.uber.org/zap"
	"time"
)

func playWithZap() {
	logger, _ := zap.NewProduction()
	defer func() {
		err := logger.Sync()
		if err != nil {
			return
		}
	}()

	url := "localhost"

	// flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	fields := make([]zap.Field, 0)

	logger = logger.With(fields...)
	logger.Info("hello Mouni madam this is without fields")

	fields = append(fields, zap.String("k1", "v1"))
	fields = append(fields, zap.String("k2", "v2"))

	logger = logger.With(fields...)
	logger.Info("hello Mouni madam this is with fields :)")
}
