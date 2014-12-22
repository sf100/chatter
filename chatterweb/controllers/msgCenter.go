package controllers

import (
	"github.com/astaxie/beego"
)

type MsgController struct {
	beego.Controller
}

func (this *MsgController) Get() {
	this.TplNames = "msgCenter/msgCenter.html"
}

func (this *MsgController) Post() {
	this.TplNames = "msgCenter/msgCenter.html"
}
