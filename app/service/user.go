package service

import (
	"errors"
	"first_web/app/model"
	"first_web/bootstrap/database"
	"first_web/config"
)

type UserService struct {
}

/**
创建一个用户
*/
func (u *UserService) CreateUser(username, password string) ([]model.User, error) {
	//生成user对象
	var user = &model.User{
		UserName: username,
		Password: password,
		Status:   0,
	}
	save := database.Db.Save(user)
	if save.Error != nil {
		return []model.User{}, errors.New(save.Error.Error())
	}
	return  u.FindUserByName(username)
}

/**
修改用户
*/
func (u *UserService) UpdateUserPasswordById(id int, password string) (model.User, error) {
	user, err := u.FindUserById(id)
	if err!= nil {
		return user,err
	}
	if err := database.Db.Model(&user).Update("password", password).Error; err != nil {
		return model.User{}, nil
	}
	user.Password = password
	return user, nil
}

/**
删除用户
*/
func (u *UserService) DeleteById(id int) error {
	user, err := u.FindUserById(id)
	if err != nil {
		return err
	}
	if err := database.Db.Model(&user).Update("status", config.Status).Error; err != nil {
		return err
	}
	return nil
}

/**
  通过id查询用户
*/
func (u *UserService) FindUserById(id int) (model.User, error) {
	var user model.User
	if err := database.Db.Where("id=?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

/**
  通过名称查询用户
*/
func (u *UserService) FindUserByName(name string) ([]model.User, error) {
	var user []model.User
	if err := database.Db.Where("username=?", name).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

/**
查询出全部的数据
*/
func (u *UserService) FindAll() ([]model.User, error) {
	var user []model.User
	if err := database.Db.Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil

}
