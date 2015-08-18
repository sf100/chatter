package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/sf100/chatter/chatterweb/models"
	"strconv"
	"strings"
)

type MsgController struct {
	beego.Controller
}

func (this *MsgController) Get() {
	this.Data["token"] = this.GetSession("token").(string)
	this.TplNames = "msgCenter/msgCenter.html"
}

func (this *MsgController) Post() {
	this.Data["token"] = this.GetSession("token").(string)
	user := this.GetSession("user").(*models.User)
	this.Data["User"] = user
	page := &models.Page{Begin: 0, PageSize: 10}
	this.Data["historyMsgs"], _ = models.GetHistoryMsg(user.Id, page)
	this.TplNames = "msgCenter/msgCenter.html"
}

/*记录用户聊天记录*/
func (this *MsgController) SavaHistory() {

	fmt.Println("-----------------------------------")
	ret := Result{Success: true}

	historyMsg := models.HistoryMsg{}

	defer func() {
		this.Data["json"] = ret
		this.ServeJson()
		this.StopRun()
	}()

	if err := this.ParseForm(&historyMsg); err != nil {
		beego.Error(err)
		ret.Success = false
		return
	}
	/*解析类型*/
	if strings.HasSuffix(historyMsg.SourceId, USER_SUFFIX) {
		historyMsg.SourceType = SOURCE_TYPE_USER
	} else if strings.HasSuffix(historyMsg.SourceId, QUN_SUFFIX) {
		historyMsg.SourceType = SOURCE_TYPE_QUN
	} else if strings.HasSuffix(historyMsg.SourceId, APP_SUFFIX) {
		historyMsg.SourceType = SOURCE_TYPE_APP
	}
	historyMsg.SourceId = historyMsg.SourceId[:strings.Index(historyMsg.SourceId, "@")]
	_, err := models.SavaHistoryMsg(&historyMsg)
	if err != nil {
		beego.Error(err)
		ret.Success = false
		return
	}

}

/*获取用户聊天记录*/
func (this *MsgController) GetUserHistoryMsg() {

	ret := Result{Success: true}

	defer func() {
		this.Data["json"] = ret
		this.ServeJson()
		this.StopRun()
	}()

	userId := this.Input().Get("userId")
	beginStr := this.Input().Get("begin")
	begin, err := strconv.Atoi(beginStr)
	if err != nil {
		beego.Error()
		ret.Success = false
		return
	}

	page := &models.Page{Begin: begin, PageSize: 10}
	historyMsgs, _ := models.GetHistoryMsg(userId, page)
	ret.Data = historyMsgs
}
