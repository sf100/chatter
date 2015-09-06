package controllers

import (
	"github.com/astaxie/beego"
)

type MsgController struct {
	beego.Controller
}

func (this *MsgController) Get() {
	this.Data["token"] = this.GetSession("token").(string)
	this.TplNames = "msgCenter/msgCenter.html"
}

func (this *MsgController) Post() {
	this.TplNames = "msgCenter/msgCenter.html"
}

/*获取用户聊天记录*/
func (this *MsgController) GetUserHistoryMsg() {

}
