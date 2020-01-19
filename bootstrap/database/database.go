package database

import (
	"first_web/bootstrap"
	"first_web/bootstrap/ini"
	"first_web/bootstrap/log"
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

var Db *gorm.DB
func init()  {
	bootstrap.Func.AddProviders(func() {
		New()
	})
}
func New()  {
	getSelect := ini.Cfg.GetSelect("Mysql")
	host := fmt.Sprintf("%s:%s",
			getSelect.GetKey("host"),
			getSelect.GetKey("port"))
	username := getSelect.GetKey("username")
	password := getSelect.GetKey("password")
	database :=getSelect.GetKey("database")
	charset := getSelect.GetKey("charset")
	client := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,password,host,database,charset)
	db, err := gorm.Open("mysql",client)
	db.DB().SetMaxIdleConns(8)
	db.DB().SetMaxOpenConns(100)
	if err != nil {
		log.Logger.Panic("数据库连接失败",err)
	}
	Db = db
}