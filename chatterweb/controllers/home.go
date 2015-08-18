package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sf100/chatter/chatterweb/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["token"] = this.GetSession("token").(string)
	this.Data["User"] = this.GetSession("user").(*models.User)
	this.TplNames = "home.html"
}
