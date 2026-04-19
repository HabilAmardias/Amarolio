package logger

import (
	"amary/src/customerror"

	"go.uber.org/zap"
)

func NewLogger(isProd bool) (*zap.SugaredLogger, error) {
	logger, err := zap.NewDevelopment()
	if isProd {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return nil, customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	return logger.Sugar(), nil
}
