package services

import (
	"fmt"
	"github.com/iamMarkchu/shop/lib/orm"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(userName string, password string, rePassword string) interface{} {
	// 判断 用户名是否存在
	// 判断 密码和重复密码是否相等
	ormEngine := orm.EngineInstance()
	sql := "SELECT * FROM `users`"
	var result interface{}
	var err error
	if result, err = ormEngine.QueryString(sql); err != nil {
		fmt.Println("err", err.Error())
	}
	return result
}
