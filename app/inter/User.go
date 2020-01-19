package inter

import "first_web/app/model"

type User interface {
	//创建一个用户
	CreateUser(username,password string)(model.User,error)
	//修改用户的密码
	UpdateUserPasswordById(id int,password string)(model.User,error)
	//删除用户
	DeleteById(id int) error
	//根据id查找用户
	FindUserById(id int) (model.User,error)
	//根据名称查找用户
	FindUserByName(name string) (model.User,error)
	//查找全部的用户
	FindAll()([]model.User,error)

}
