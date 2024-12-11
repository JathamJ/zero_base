package httpo

import "fmt"

// CodeMsg 基础错误
type CodeMsg struct {
	Code int
	Msg  string
}

func (e *CodeMsg) GetErrCode() int {
	return e.Code
}

func (e *CodeMsg) GetErrMsg() string {
	return e.Msg
}

func (e *CodeMsg) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

func NewCodeMsg(code int, message string) *CodeMsg {
	return &CodeMsg{
		Code: code,
		Msg:  message,
	}
}
