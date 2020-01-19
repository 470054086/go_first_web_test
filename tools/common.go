package tools

import (
	"encoding/json"
)
type ResponseStruct struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func Response (data interface{},err error) []byte  {

	var rs ResponseStruct
	if err != nil {
		rs = ResponseStruct{
			Code:    500,
			Message: "服务器错误",
			Data:    make(map[string]interface{}),
		}
	}else{
		rs = ResponseStruct{
			Code:    200,
			Message: "请求成功",
			Data:    data,
		}
	}

	marshal, _ := json.Marshal(rs)
	return marshal
}