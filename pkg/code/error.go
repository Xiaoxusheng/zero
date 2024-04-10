package code

type Code interface {
	Error() string
	GetCode() int32
}

type ErrorCode struct {
	Msg  string
	Code int32
}

func New(code int32, msg string) *ErrorCode {
	return &ErrorCode{
		Msg:  msg,
		Code: code,
	}
}

func (e *ErrorCode) Error() string {
	return e.Msg
}

func (e *ErrorCode) GetCode() int32 { return e.Code }
