package controller

import (
	"first_web/bootstrap/send"
	"first_web/bootstrap/send/mobile"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type SendController struct {

}

func (s * SendController) Send(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	message := &send.SendMessage{
		"13708080808",
		"我就是个韩寒",
	}
	mobile.G_Send.Send(message)
}