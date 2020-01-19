package log

import (
	"first_web/bootstrap"
	"first_web/bootstrap/ini"
	"first_web/tools"
	"github.com/Sirupsen/logrus"
	"io"
	"log"
	"os"
	"strconv"
)

var Logger *logrus.Entry

func init() {
	bootstrap.Func.AddProviders(func() {
		New()
	})
}
func New() {
	logger := logrus.New()
	format := ini.Cfg.GetSelect("Log").GetKey("textFormat")
	if format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:  "2006-01-02 15:04:05",
		})
	}else {
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})
	}

	debug := ini.Cfg.GetSelect("App").GetKey("debug")
	parseBool, err := strconv.ParseBool(debug)
	if err != nil {
		log.Fatal("转化错误啦", err)
	}
	if parseBool {
		logger.SetOutput(os.Stdout)
	} else {
		logger.SetOutput(write())

	}
	//设置固定变量
	logInit := logger.WithFields(logrus.Fields{
		"version": "1.2.3",
		"flag":    true,
	})
	Logger = logInit
}

func write() io.Writer {
	Select := ini.Cfg.GetSelect("log")
	key := Select.GetKey("driver")

	if key == "file" {
		//创建文件的处理
		filePath := Select.GetKey("filePath")
		file := &tools.ToolFile{FileName: filePath}
		Osfile, err := file.CreateFile()
		if err != nil {
			logrus.Fatal(err)
		}
		return Osfile

	} else if key == "redis" {

	}
	return nil
}
