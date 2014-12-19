package app

import (
	log "code.google.com/p/log4go"
	"github.com/sf100/chatter/db"
	"strings"
	"time"
)

type User struct {
	Id         string
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

func CheckUserByToken(token string) *User {

	conn := rs.getConn("token")
	if conn == nil {
		return nil
	}
	defer conn.Close()

	if err := conn.Send("EXISTS", token); err != nil {
		log.Error(err)
		return nil
	}

	if err := conn.Flush(); err != nil {
		log.Error(err)
		return nil
	}

	reply, err := conn.Receive()
	if err != nil {
		log.Error(err)
		return nil
	}
	if 0 == reply.(int64) { // 令牌不存在
		return nil
	}

	idx := strings.Index(token, "_")
	if -1 == idx {
		return nil
	}
	user_id := token[:idx]
	// 从数据库加载用户
	ret := getUserByUid(user_id)
	if nil == ret {
		return nil
	}

	// 刷新令牌
	confExpire := int64(Conf.TokenExpire)
	if err := conn.Send("EXPIRE", token, confExpire); err != nil {
		log.Error(err)
	}
	if err := conn.Flush(); err != nil {
		log.Error(err)
	}
	_, err = conn.Receive()
	if err != nil {
		log.Error(err)
	}

	return ret
}

//更具user_id获取用户
func getUserByUid(id string) *User {
	row := db.MySQL.QueryRow("select * from user where id =?", id)
	user := &User{}
	if err := row.Scan(user.Id, user.Name, user.NickName, user.NamePy, user.Avatar, user.Status, user.Password, user.Sex,
		user.Level, user.Location, user.Mobile, user.Email, user.Occupation, user.URL); err != nil {

		log.Error(err)
		return nil
	}
	return user
}
