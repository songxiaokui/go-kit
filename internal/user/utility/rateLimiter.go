package utility

import (
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"sync"
)

/*
@Time    : 2021/3/24 19:51
@Author  : austsxk
@Email   : austsxk@163.com
@File    : rateLimiter.go
@Software: GoLand
*/

// create a ips rateLimiter

// ipRateLimiter
type IpRateLimiter struct {
	mux *sync.RWMutex
	ips map[string]*rate.Limiter
	b   int
	r   rate.Limit
}

// construct function
func NewIpRateLimiter(b int, r rate.Limit) *IpRateLimiter {
	return &IpRateLimiter{
		mux: &sync.RWMutex{},
		ips: make(map[string]*rate.Limiter),
		b:   b,
		r:   r,
	}
}

// add ip, then return a rate.Limiter to use
func (i *IpRateLimiter) AddIp(ip string) *rate.Limiter {
	i.mux.Lock()
	defer i.mux.Unlock()
	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	return limiter
}

// depend ip get ipLimiter
func (i *IpRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mux.RLock()
	if limiter, ok := i.ips[ip]; ok {
		i.mux.RUnlock()
		return limiter
	} else {
		i.mux.RUnlock()
		return i.AddIp(ip)
	}
}

// func use server test ip limiter
func HttpRateLimiterServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello, world"))
	})
	// crate a limiter
	ipLimiter := NewIpRateLimiter(5, 1)
	fmt.Println("server is running...")
	err := http.ListenAndServe("127.0.0.1:8080", RateLimiterWare(ipLimiter, mux))
	if err != nil {
		log.Fatal(err)
	}
}

// 使用闭包实现一个装饰器,需要一个限流器
func RateLimiterWare(ipLimiter *IpRateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// get ip rateLimier
		limiter := ipLimiter.GetLimiter(request.RemoteAddr)
		if limiter.Allow() {
			next.ServeHTTP(writer, request)
			return
		} else {
			http.Error(writer, http.StatusText(http.StatusTooManyRequests),
				http.StatusTooManyRequests)
			return
		}
	})
}
