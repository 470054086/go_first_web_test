package controller

import (
	"encoding/json"
	"first_web/app/inputout/request"
	"first_web/app/service"
	tools "first_web/tools"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HelloController struct {
	userService service.UserService
}

//接受参数
func (h *HelloController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var rs = &request.Create{}
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(rs)
	//检测 user 和password
	username := rs.UserName
	password := rs.Password
	user, err := h.userService.CreateUser(username, password)
	response := tools.Response(user, err)
	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(response))
}

/**
修改用户名
*/
func (h *HelloController) UpdateName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var rs = &request.UpdateName{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(rs)
	id := rs.Id
	password := rs.Password
	user, err := h.userService.UpdateUserPasswordById(id, password)
	response := tools.Response(user, err)
	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(response))
}

/**
 删除用户
 */
func (h *HelloController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)   {
	var rs = &request.DeleteUser{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(rs)
	id := rs.Id
	err := h.userService.DeleteById(id)
	var response interface{}
	resp := tools.Response(response, err)
	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(resp))
}

func (h *HelloController) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	users, err := h.userService.FindAll()
	response := tools.Response(users, err)
	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w,string(response))
}
func (h *HelloController) FindUserByName(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	name := p.ByName("name")
	name = name[0:]
	user , err := h.userService.FindUserByName(name)
	response := tools.Response(user, err)
	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w,string(response))
}
