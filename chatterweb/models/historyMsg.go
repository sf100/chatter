package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type HistoryMsg struct {
	Id         int `orm:"column(id);pk"`
	SourceId   string
	TargetId   string
	Content    string
	MsgType    int
	SourceType int
	Status     int
	Created    time.Time
}

/*保存消息记录，返回插入记录后的ID*/
func SavaHistoryMsg(historyMsg *HistoryMsg) (int64, error) {

	o := orm.NewOrm()
	historyMsg.Created = time.Now().Local()
	historyMsg.Id = 0
	id, err := o.Insert(historyMsg)
	if err != nil {
		beego.Error(err)
		if err := o.Rollback(); err != nil {
			beego.Error(err)
		}
		return 0, err
	}

	return id, nil
}

/**获取
@userId 用户id
*/
func GetHistoryMsgCount(userId string) int {
	o := orm.NewOrm()
	count := 0
	err := o.Raw("select count(*) from histtory_msg where target_id=? ", userId).QueryRow(&count)
	if err != nil {
		beego.Error(err)
	}
	return count
}

/**获取用户历史聊天消息
@userId 用户id
@page 分页对象
*/
func GetHistoryMsg(userId string, page *Page) ([]HistoryMsg, error) {
	o := orm.NewOrm()
	var historyMsgs []HistoryMsg
	_, err := o.Raw("select id ,source_id,target_id,msg_type,created from histtory_msg where source_id = ? limit ? ?", userId, page.Begin, page.PageSize).QueryRows(&historyMsgs)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return historyMsgs, nil
}

/*获取聊天记录的用户*/
func GetHistoryUser(userId string, page *Page) ([]HistoryMsg, error) {

	o := orm.NewOrm()
	o.Raw(`select t3.source_id,count(*) , t4.nick_name , t3.status
           from history_msg t3 , user t4 
           where  t3.target_id = ? and t3.status = 0 and t3.source_id = t4.id
           GROUP BY t3.source_id , t3.target_id  limit ? , ?
         `, userId, page.Begin, page.PageSize)

	return nil, nil
}
