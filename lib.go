package lib

import "go.uber.org/zap"

func Live() int {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "url",
		"attempt", 3,
		"backoff", "time",
	)
	return 12
}

func Dead() int {

	return 34
}
