package router

import (
	"first_web/app/controller"
	"first_web/bootstrap"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Router struct {
	Router *httprouter.Router
}
//添加路由
func (r *Router) addRoute()  {
	var Hello= &controller.HelloController{}
	r.Router.POST("/create",Hello.Create)
	r.Router.POST("/updatename",Hello.UpdateName)
	r.Router.POST("/deleteuser",Hello.DeleteUser)
	r.Router.GET("/getall",Hello.FindAll)
	r.Router.GET("/get/:name",Hello.FindUserByName)

	var send = &controller.SendController{}
	r.Router.GET("/send",send.Send)
}

func init()  {
	bootstrap.Func.AddProviders(func() {
		router := Router{Router:httprouter.New()}
		router.addRoute()
		log.Fatal(http.ListenAndServe(":8080",router.Router))
	})
}