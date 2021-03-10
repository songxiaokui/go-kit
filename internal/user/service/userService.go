package service

import (
	"fmt"
	"sync"
)

/*
@Time    : 2021/3/9 15:25
@Author  : austsxk
@Email   : austsxk@163.com
@File    : userService.go
@Software: GoLand
*/
type User struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type UserServer interface {
	SearchUser(id int) string
	AddUser(name string, id int) bool
}

// 验证结构体是否实现了该接口
var _ UserServer = (*UserImpl)(nil)

type UserImpl struct {
	m sync.Mutex
	d map[int]string
}

func (u *UserImpl) SearchUser(id int) string {
	if v, ok := u.d[id]; ok {
		return v
	}
	return fmt.Sprintf("Id: %d User Not Found", id)
}

func (u *UserImpl) AddUser(name string, id int) bool {
	u.m.Lock()
	defer u.m.Unlock()
	u.d[id] = name
	return true
}

// 最好继承自dao层的持久化存储操作
func NewUserImpl() *UserImpl {
	return &UserImpl{d: make(map[int]string)}
}
