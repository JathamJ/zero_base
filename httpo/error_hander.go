package httpo

import (
	"context"
	"errors"
	"github.com/JathamJ/zero_base/errx"
	"net/http"
)

func DefaultErrorHandler(ctx context.Context, err error) (int, any) {
	resp := &Response{}
	var e *CodeMsg
	if errors.As(err, &e) { //自定义错误类型（默认）
		resp.Code = e.GetErrCode()
		resp.Msg = e.GetErrMsg()
	} else {
		resp.Code = errx.InvalidParam //无错误码按参数错误处理
		resp.Msg = err.Error()
	}

	return http.StatusOK, resp
}
