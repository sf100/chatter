package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/sf100/chatter/chatterweb/models"
)

type HuaLaoController struct {
	beego.Controller
}

func (this *HuaLaoController) Get() {
	this.TplNames = "hualao/hualao.html"
}

func (this *HuaLaoController) Post() {

	this.TplNames = "hualao/hualao.html"
	user := this.GetSession("user").(*models.User)
	users := models.GetUserFriend(user.Id)
	fmt.Println("-------------------")
	fmt.Println(users)
	fmt.Println("-------------------")
	//默认打开用户好友列表
	this.Data["friendLen"] = len(users)
	num := models.GetUserQunNums(user.Id)
	fmt.Println("---------------------", num)
	this.Data["qunsLen"] = models.GetUserQunNums(user.Id)
	this.Data["Users"] = users
}

/*获取好友信息列表*/
func (this *HuaLaoController) GetUserFriend() {

	user := this.GetSession("user").(*models.User)
	users := models.GetUserFriend(user.Id)
	this.TplNames = "hualao/userList.html"
	this.Data["Users"] = users
}

func (this *HuaLaoController) GetUserQuns() {
	user := this.GetSession("user").(*models.User)
	quns := models.GetUserQuns(user.Id)
	this.TplNames = "hualao/qunList.html"
	this.Data["quns"] = quns

}
