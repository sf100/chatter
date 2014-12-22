package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/sf100/go-uuid/uuid"
	"time"
)

type Qun struct {
	Id          string `orm:"column(id);pk"`
	Name        string
	Avatar      string
	TypeId      string
	CreatorId   string
	Liveness    string
	Description string
	MaxMember   string
	Created     time.Time
	Updated     time.Time
}

/*获取用户所属的群*/
func GetUserQuns(userId string) []Qun {
	o := orm.NewOrm()
	var quns []Qun
	_, err := o.Raw("select t2.* from qun_user t1 left join qun t2 on t1.qun_id = t2.id where t1.user_id = ? order by sort", userId).QueryRows(&quns)
	if err != nil {
		beego.Error(err)
		return nil
	}
	return quns
}

/*获取用户关注的群的数量*/
func GetUserQunNums(userId string) int {
	o := orm.NewOrm()
	num := 0
	err := o.Raw("select count(*) from qun_user t1 left join qun t2 on t1.qun_id = t2.id where t1.user_id = ? ", userId).QueryRow(&num)
	if err != nil {
		beego.Error(err)
	}
	return num
}
