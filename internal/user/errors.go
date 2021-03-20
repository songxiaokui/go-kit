package user

/*
@Time    : 2021/3/20 17:16
@Author  : austsxk
@Email   : austsxk@163.com
@File    : errors.go
@Software: GoLand
*/

type AustErr struct {
	Msg  string
	Code int
}

func (a *AustErr) Error() string {
	return a.Msg
}

func NewError(code int, msg string) error {
	return &AustErr{msg, code}
}
