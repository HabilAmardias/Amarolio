package users

type Logger interface {
	Errorln(args ...interface{})
	Infoln(args ...interface{})
}
