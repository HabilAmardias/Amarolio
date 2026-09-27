package users

type Logger interface {
	Errorln(args ...interface{})
}
