package app

import (
	log "code.google.com/p/log4go"
	"encoding/json"
	"github.com/sf100/chatter/db"
	"io/ioutil"
	"net/http"
	"strings"
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

func getUserByName(userName string) *User {
	row := db.MySQL.QueryRow("select id ,name,nick_name,signature,avatar,status,password,sex,level,location,mobile,email,occupation,url from user where name =?", userName)
	user := &User{}
	if err := row.Scan(&user.Id, &user.Name, &user.NickName, &user.Signature, &user.Avatar, &user.Status, &user.Password, &user.Sex,
		&user.Level, &user.Location, &user.Mobile, &user.Email, &user.Occupation, &user.Url); err != nil {
		log.Info(err)
		return nil
	}
	return user
}

//根据user_id获取用户
func getUserByUid(uid string) *User {
	row := db.MySQL.QueryRow("select id ,name,nick_name,signature,avatar,status,password,sex,level,location,mobile,email,occupation,url from user where id =?", uid)
	user := &User{}
	if err := row.Scan(&user.Id, &user.Name, &user.NickName, &user.Signature, &user.Avatar, &user.Status, &user.Password, &user.Sex,
		&user.Level, &user.Location, &user.Mobile, &user.Email, &user.Occupation, &user.Url); err != nil {
		log.Info(err)
		return nil
	}
	return user
}

// 客户端设备登录.
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	baseRes := baseResponse{OK, ""}
	body := ""
	res := map[string]interface{}{"baseResponse": &baseRes}
	defer RetPWriteJSON(w, r, res, &body, time.Now())

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res["ret"] = ParamErr
		log.Error("ioutil.ReadAll() failed (%s)", err.Error())
		return
	}
	body = string(bodyBytes)

	var args map[string]interface{}

	if err := json.Unmarshal(bodyBytes, &args); err != nil {
		baseRes.ErrMsg = err.Error()
		baseRes.Ret = ParamErr
		return
	}

	//TODO:备用
	/*
		baseReq := args["baseRequest"].(map[string]interface{})
		uid := baseReq["uid"].(string)
		deviceId := baseReq["deviceID"].(string)
		deviceType := baseReq["deviceType"].(string)
	*/
	userName := args["userName"].(string)
	password := args["password"].(string)
	// TODO: 登录验证逻辑
	user := getUserByName(userName)
	if nil == user || user.Password != password {
		baseRes.ErrMsg = "auth failed"
		baseRes.Ret = AuthErr
		return
	}
	token, err := genToken(user)
	if err != nil {
		baseRes.ErrMsg = err.Error()
		baseRes.Ret = InternalErr
		return
	}
	res["token"] = token
	res["user"] = user
}
