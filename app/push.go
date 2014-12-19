package app

import (
	glog "code.google.com/p/log4go"
	"encoding/json"
	"fmt"
	myrpc "github.com/sf100/chatter/rpc"
	"net/http"
	"strconv"
	"strings"
)

func UserPush(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---------------qing qiu-------------------")
	baseRes := baseResponse{OK, ""}
	res := map[string]interface{}{"baseResponse": &baseRes}
	var callback *string
	defer func() {
		// 返回结果格式化
		resJsonStr := ""
		if resJson, err := json.Marshal(res); err != nil {
			baseRes.ErrMsg = err.Error()
			baseRes.Ret = InternalErr
		} else {
			resJsonStr = string(resJson)
		}
		fmt.Fprintln(w, *callback, "(", resJsonStr, ")")
	}()

	var err error
	var msg = make(map[string]interface{})

	//获取请求数据
	r.ParseForm()

	// Token 校验
	token := r.FormValue("baseRequest[token]")
	user := CheckUserByToken(token)
	if nil == user {
		baseRes.Ret = AuthErr
		baseRes.ErrMsg = "auth failure"
		return
	}

	tmp := r.FormValue("callbackparam")
	callback = &tmp
	if err != nil {
		baseRes.Ret = ParamErr
		baseRes.ErrMsg = "msgType not is int"
		return
	}

	fromUserName := r.FormValue("msg[fromUserName]")
	toUserName := r.FormValue("msg[toUserName]")
	toUserID := toUserName[:strings.Index(toUserName, "@")]

	keys := []string{}
	if strings.HasSuffix(toUserName, USER_SUFFIX) { // 如果是推人

		msg["fromDisplayName"] = user.NickName
		msg["content"] = r.FormValue("msg[content]")
		keys = append(keys, toUserName)

	} else if strings.HasSuffix(toUserName, QUN_SUFFIX) {

		qun := GetQunById(toUserID)
		if nil == qun {
			baseRes.Ret = InternalErr
			return
		}
		msg["content"] = fromUserName + "|" + user.Name + "|" + user.NickName + "&&" + r.FormValue("msg[content]")
		msg["fromDisplayName"] = qun.Name
		msg["fromUserName"] = toUserName

		userIds := GetQunMemberIDs(toUserID)
		keys = append(keys, userIds...)

	}

	// 消息过期时间（单位：秒）
	exp := r.FormValue("msg[expire]")
	expire := 600
	if len(exp) > 0 {
		expire, err = strconv.Atoi(exp)
		if err != nil {
			baseRes.Ret = ParamErr
			return
		}
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		glog.Error("msg convert byte error:[%s] ", err)
		return
	}
	for _, v := range keys {
		push(v, msgBytes, expire)
	}

}

// 按 key 推送.
func push(key string, msgBytes []byte, expire int) int {

	node := myrpc.GetComet(key)
	if node == nil || node.Rpc == nil {

		glog.Error("Get comet node failed [key=%s]", key)
		return NotFoundServer
	}

	client := node.Rpc.Get()
	if client == nil {

		glog.Error("Get comet node RPC client failed [key=%s]", key)
		return NotFoundServer
	}

	pushArgs := &myrpc.CometPushPrivateArgs{Msg: json.RawMessage(msgBytes), Expire: uint(expire), Key: key}

	ret := OK
	if err := client.Call(myrpc.CometServicePushPrivate, pushArgs, &ret); err != nil {
		glog.Error("client.Call(\"%s\", \"%v\", &ret) error(%v)", myrpc.CometServicePushPrivate, string(msgBytes), err)
		return InternalErr
	}

	glog.Info("Pushed a message to [key=%s]", key)

	return ret
}