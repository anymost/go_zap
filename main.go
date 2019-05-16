package main
import "go.uber.org/zap"

func sugarLog() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info("hello world")
	sugar.Warn("warning")
	sugar.Debug("hello world")
}

func standardLog() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(logger)
	}
	defer logger.Sync()
	logger.Info("hello world", zap.Int("age", 20), zap.String("name", "jack"))
}

func main() {
	// sugarLog()
	standardLog()
}
