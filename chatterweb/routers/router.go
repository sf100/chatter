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
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/home/msgCenter", &controllers.MsgController{})
	beego.Router("/hualao", &controllers.HuaLaoController{})
	beego.Router("/hualao/friend", &controllers.HuaLaoController{}, "post:GetUserFriend")
	beego.Router("/hualao/quns", &controllers.HuaLaoController{}, "post:GetUserQuns")
}
