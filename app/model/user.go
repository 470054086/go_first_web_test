package model

import (
	"first_web/config"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Status  config.StatusType `json:"status"`
}

func (u * User) TableName() string  {
	return  "user"
}

