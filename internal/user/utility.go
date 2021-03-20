package user

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
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
				return nil, errors.New("429 请求太快了")
			}
			return next(ctx, request)
		}
	}
}
