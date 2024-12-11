package httpo

import (
	"context"
	"github.com/JathamJ/zero_base/errx"
)

func DefaultOkHandler(ctx context.Context, data any) any {
	return &Response{
		Code: errx.Success,
		Msg:  "success",
		Data: data,
	}
}
