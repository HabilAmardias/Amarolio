package logger

import "go.uber.org/zap"

func CreateNewLogger(isProd bool) (*zap.SugaredLogger, error) {
	if isProd {
		logger, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}
		return logger.Sugar(), nil
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
