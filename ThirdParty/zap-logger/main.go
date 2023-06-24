package main

import "go.uber.org/zap"

//func main() {
//	logger := zap.NewExample()
//	defer logger.Sync()
//
//	url := "http://example.org/api"
//	logger.Info("failed to fetch URL",
//		zap.String("url", url),
//		zap.Int("attempt", 3),
//		zap.Duration("backoff", time.Second),
//	)
//
//	sugar := logger.Sugar()
//	sugar.Infow("failed to fetch URL",
//		"url", url,
//		"attempt", 3,
//		"backoff", time.Second,
//	)
//	sugar.Infof("Failed to fetch URL: %s", url)
//}

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("tracked some metrics",
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	)

	logger2 := logger.With(
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	)
	logger2.Info("tracked some metrics")
}
