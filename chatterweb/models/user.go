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
	Signature  string
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

type UserInfo struct {
	Id          string
	Name        string
	NickName    string
	Signature   string
	Avatar      string
	Status      int
	Password    string
	Sex         int
	Level       int
	Location    string
	Mobile      string
	Email       string
	Occupation  string
	Url         string
	Remark_name string
}

/**根据用户名获取用户*
*  @name 用户登陆名
 */
func GetUserByName(name string) *User {
	o := orm.NewOrm()
	var users []User
	mun, err := o.Raw("select * from user where name = ? ", name).QueryRows(&users)
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
		o.QueryTable("user")
		id, err := o.Insert(user)
		if err != nil {
			beego.Error(err)
			if err := o.Rollback(); err != nil {
				beego.Error(err)
			}
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

/*获取用户关注好友*/
func GetUserFriend(userId string) []UserInfo {
	if len(userId) == 0 {
		return nil
	}
	users := []UserInfo{}
	o := orm.NewOrm()
	_, err := o.Raw("SELECT t2.* , t1.remark_name from user_user t1 left JOIN user t2 on t1.to_user_id = t2.id  where from_user_id =? order by sort", userId).QueryRows(&users)
	if err != nil {
		beego.Error(err)
		return nil
	}
	return users
}

//验证用户是否登陆
var IsUserLogin = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/")
	}
}
