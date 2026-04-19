package customerror

import "fmt"

const (
	Unauthenticate       = 40101
	CommonErr            = 50001
	ItemNotFound         = 40401
	DatabaseExecutionErr = 50002
	InvalidAction        = 40001
)

func NewError(userErr string, sysErr error, errCode int) *CustomError {
	return &CustomError{
		UserErr: userErr,
		SysErr:  sysErr,
		ErrCode: errCode,
	}
}

type CustomError struct {
	UserErr string
	SysErr  error
	ErrCode int
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("user: %s, system: %s", ce.UserErr, ce.SysErr.Error())
}

func (ce *CustomError) GetErrStatusCode() int {
	return ce.ErrCode / 100
}
