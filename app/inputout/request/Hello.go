package request

//Hello struct 请求参数

type Create struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UpdateName struct {
	Id  int `json:"id"`
	Password string `json:"password"`
}

type DeleteUser struct {
	Id int `json:"id"`
}