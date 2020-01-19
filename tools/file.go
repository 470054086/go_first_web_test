package tools

import "os"

type ToolFile struct {
	FileName string
}
var File = new(ToolFile)

func (t *ToolFile) CreateFile() (*os.File,error)  {
	//判断文件是否存在
	exist, err := t.IsExist()
	if err != nil {
		return nil,err
	}
	if !exist {
		_, err = os.Create(t.FileName)
		if err != nil {
			return  nil,err
		}
	}
	openFile, err := os.OpenFile(t.FileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return  nil,err
	}

	return openFile,nil
}
/**
判断或者目录
 */
func (t *ToolFile) IsExist() (bool,error){
	_, err := os.Stat(t.FileName)
	if err == nil {
		return true,nil
	}
	if os.IsExist(err) {
		return false,err
	}
	return false,err

}