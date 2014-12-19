package routers

import (
	"github.com/astaxie/beego"
	"github.com/sf100/chatter/chatterweb/controllers"
)

func init() {
	beego.Router("/", &controllers.UserController{})
	beego.Router("/login", &controllers.UserController{})
	beego.Router("/register", &controllers.UserController{}, "get:Register")
	beego.Router("/doRegister", &controllers.UserController{}, "post:DoRegister")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/index", &controllers.IndexController{})

}
