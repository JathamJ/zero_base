package httpo

import (
	"bytes"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpc"
	"io"
	"net/http"
)

type Client struct {
	Service httpc.Service
}

func NewService(name string, opts ...httpc.Option) httpc.Service {
	return &Client{
		Service: httpc.NewService(name, opts...),
	}
}

func NewServiceWithClient(name string, cli *http.Client, opts ...httpc.Option) httpc.Service {
	return &Client{
		Service: httpc.NewServiceWithClient(name, cli, opts...),
	}
}

func (c *Client) Do(ctx context.Context, method, url string, data any) (*http.Response, error) {
	return c.Service.Do(ctx, method, url, data)
}

func (c *Client) DoRequest(r *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	defer func() {
		// 读取请求体
		var reqBodyBytes, respBodyBytes []byte
		reqBodyBytes, _ = io.ReadAll(r.Body)
		respBodyBytes, _ = io.ReadAll(resp.Body)
		// 重置请求体
		r.Body = io.NopCloser(bytes.NewBuffer(respBodyBytes))
		logx.WithContext(r.Context()).Debugf("http do request, [%s]%s, body: %s, headers: %v, statusCode: %d, respHeader: %v, respBody: %s", r.Method, r.RequestURI, string(reqBodyBytes), r.Header, resp.StatusCode, resp.Header, string(respBodyBytes))
	}()
	resp, err = c.Service.DoRequest(r)
	return resp, err
}
