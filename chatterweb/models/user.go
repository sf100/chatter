package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id         string `orm:"column(id);pk"`
	Name       string
	NickName   string
	NamePy     string
	Avatar     string
	Status     int
	Password   string
	Sex        int
	Level      int
	Location   string
	Mobile     string
	Email      string
	Occupation string
	URL        string
	Created    time.Time
	Updated    time.Time
	QunId      string
}

/**根据用户名获取用户*
*  @name 用户登陆名
 */
func GetUserByName(name string) *User {
	o := orm.NewOrm()
	var users []User
	mun, err := o.Raw("select * from user where name = ?", name).QueryRows(&users)
	if err != nil {
		beego.Error(err)
		return nil
	}
	if mun > 0 {
		return &users[0]
	}
	return nil

}

//验证用户是否登陆
var IsUserLogin = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/")
	}
}
