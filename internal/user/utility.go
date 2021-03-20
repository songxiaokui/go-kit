package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
	"net/http"
)

/*
@Time    : 2021/3/20 16:52
@Author  : austsxk
@Email   : austsxk@163.com
@File    : utility.go
@Software: GoLand
*/

// 限流方法,使用闭包，中间键方式装饰endpoint
func RateLimiterEndpoint(rate *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !rate.Allow() {
				return nil, NewError(419, "太快")
			}
			return next(ctx, request)
		}
	}
}

// 自定义函数处理，参考默认的处理方法
/*
func DefaultErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			contentType, body = "application/json; charset=utf-8", jsonBody
		}
	}
	w.Header().Set("Content-Type", contentType)
	if headerer, ok := err.(Headerer); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	w.Write(body)
}
*/

func MyErrorHandlerEncoder(_ context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	w.Header().Set("Content-Type", contentType)
	// 断言err为自定义error类型
	var code int = 500
	if e, ok := err.(*AustErr); ok {
		code = e.Code
	}
	w.WriteHeader(code)
	w.Write(body)
}
