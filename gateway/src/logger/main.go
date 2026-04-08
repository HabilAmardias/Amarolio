package logger

import (
	"amarolio-gateway/src/customerrors"

	"go.uber.org/zap"
)

type Logger interface {
	Infoln(args ...interface{})
	Errorln(args ...interface{})
	Panicln(args ...interface{})
}

func NewLogger(isProd bool) (*zap.SugaredLogger, error) {
	logger, err := zap.NewDevelopment()
	if isProd {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return logger.Sugar(), nil
}
