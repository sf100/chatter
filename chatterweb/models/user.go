package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sf100/go-uuid/uuid"
	"time"
)

type User struct {
	Id         string `orm:"column(id);pk"`
	Name       string
	NickName   string
	Avatar     string
	Status     int
	Password   string
	Sex        int
	Level      int
	Location   string
	Mobile     string
	Email      string
	Occupation string
	Url        string
	Created    time.Time
	Updated    time.Time
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

/*注册用户*/
func Register(userName, password string) bool {

	if ValidUserName(userName) {
		o := orm.NewOrm()
		user := &User{
			Id:       uuid.New(),
			Name:     userName,
			Password: password,
			Created:  time.Now().Local(),
			Updated:  time.Now().Local(),
		}

		id, err := o.Insert(user)
		fmt.Println("1-->>", id)
		if err != nil {
			beego.Error(err)
			return false
		}
		fmt.Println(id)
	}

	return true
}

/*校验用户名是否可用*/
func ValidUserName(userName string) bool {
	o := orm.NewOrm()
	var users []User
	mun, err := o.Raw("select * from user where name = ? ", userName).QueryRows(&users)
	if err != nil {
		beego.Error(err)
		return false
	}
	if mun > 0 {
		return false
	}
	return true
}

//验证用户是否登陆
var IsUserLogin = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/")
	}
}
