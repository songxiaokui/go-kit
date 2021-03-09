package service

/*
@Time    : 2021/3/9 15:25
@Author  : austsxk
@Email   : austsxk@163.com
@File    : userService.go
@Software: GoLand
*/

type UserServer interface {
	SearchUser(id int) string
}

// 验证结构体是否实现了该接口
var _ UserServer = (*UserImpl)(nil)

type UserImpl struct {
}

func (u *UserImpl) SearchUser(id int) string {
	if id == 1 {
		return "austsxk"
	}
	return "not found"
}

// 最好继承自dao层的持久化存储操作
func NewUserImpl() *UserImpl {
	return &UserImpl{}
}
