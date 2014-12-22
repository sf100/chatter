package controllers

import (
	"fmt"
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

func (this *UserController) Register() {
	this.TplNames = "register.html"
}

func (this *UserController) DoRegister() {

	fmt.Println("------------------注册-----------------------")
	//返回结果
	ret := Result{
		Success: true,
	}
	uname := this.Input().Get("userName")
	pwd := this.Input().Get("password")
	fmt.Println(uname, "---", pwd)
	if len(uname) == 0 || len(pwd) == 0 {
		ret.Success = false
		ret.Msg = "用户名和密码不能为空"
		return
	} else if len(pwd) < 6 {
		ret.Success = false
		ret.Msg = "密码必须6位以上"

	}

	if !models.Register(uname, pwd) {
		ret.Success = false
		ret.Msg = "服务器繁忙，请稍后在试！"
	}

	/*注册成功，添加会话session*/
	if ret.Success {
		this.SetSession("user", "222")
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
