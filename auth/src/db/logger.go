package db

type Logger interface {
	Infoln(args ...interface{})
	Errorln(args ...interface{})
}
