package log

import (
	"go.uber.org/zap"
)

var sugarLog *zap.SugaredLogger

func Init() {
	logger, _ := zap.NewProduction()

	sugarLog = logger.Sugar()

}
