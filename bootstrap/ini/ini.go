package ini

import (
	"first_web/bootstrap"
	"fmt"
	"github.com/go-ini/ini"
	"log"
)

type iniCfg struct {
	file    *ini.File
	section *ini.Section
}

const fileName = "D:/go/src/first_web/config.ini";
//全局变量
var Cfg *iniCfg

func init()  {
	bootstrap.Func.AddProviders(func() {
		New(fileName)
	})
}


func New(fileName string)  {
	cfg, err := ini.InsensitiveLoad(fileName)
	if err != nil {
		log.Fatal(fmt.Sprintf("读取配置文件失败,错误原因为%v", err))
	}
	Cfg = &iniCfg{
		file:    cfg,
		section: nil,
	}
}

func (f *iniCfg) GetSelect(groupName string) *iniCfg {
	section, err := f.file.GetSection(groupName)
	if err != nil {
		log.Fatal(fmt.Sprintf("读取配置文件的section失败,错误为%v", err))
	}
	f.section = section
	return f
}
func (f *iniCfg) GetKey(key string) string {
	getKey, err := f.section.GetKey(key)
	if err != nil {
		return ""
	}
	return getKey.Value()
}

func (f *iniCfg ) GetKeyDefault(key string,defaultValue string) string  {
	getKey, err := f.section.GetKey(key)
	if err != nil {
		return defaultValue
	}
	return getKey.Value()
}


