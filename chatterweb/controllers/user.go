package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/sf100/chatter/chatterweb/models"
)

/*user控制器*/
type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.TplNames = "login.html"
}

func (this *UserController) Post() {

	//返回结果
	ret := Result{
		Success: true,
	}
	uname := this.Input().Get("userName")
	pwd := this.Input().Get("password")
	user := models.GetUserByName(uname)

	if user != nil && user.Password == pwd {
		ret.Data = user
		this.SetSession("user", user)

	} else {
		ret.Success = false
		ret.Msg = "用户名或密码错误"
	}

	this.Data["json"] = ret
	this.ServeJson()
	this.StopRun()
}

func (this *UserController) Logout() {
	//清除缓存
	this.DelSession("user")
	this.Redirect("/", 302)
}
