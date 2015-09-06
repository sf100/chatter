package app

import (
	"encoding/json"
	"fmt"
	myrpc "github.com/sf100/chatter/rpc"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func UserPush(w http.ResponseWriter, r *http.Request) {
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
	/*回调函数*/
	tmp := r.FormValue("callbackparam")
	callback = &tmp

	// Token 校验
	token := r.FormValue("baseRequest[token]")
	fmt.Println("token--->", token)
	user := CheckUserByToken(token)

	if nil == user {
		baseRes.Ret = AuthErr
		baseRes.ErrMsg = "auth failure"
		return
	}

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
		msg["toUserName"] = r.FormValue("msg[toUserName]")
		msg["msgType"] = r.FormValue("msg[msgType]")
		msg["fromUserName"] = r.FormValue("msg[fromUserName]")
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
		log.Fatal("msg convert byte error:[%s] ", err)
		return
	}
	fmt.Println(keys)
	for _, v := range keys {
		push(v, msgBytes, expire)
	}

}

// 按 key 推送.
func push(key string, msgBytes []byte, expire int) int {

	node := myrpc.GetComet(key)
	if node == nil || node.Rpc == nil {

		log.Fatal("Get comet node failed [key=%s]", key)
		return NotFoundServer
	}

	client := node.Rpc.Get()
	if client == nil {

		log.Fatal("Get comet node RPC client failed [key=%s]", key)
		return NotFoundServer
	}

	pushArgs := &myrpc.CometPushPrivateArgs{Msg: json.RawMessage(msgBytes), Expire: uint(expire), Key: key}

	ret := OK
	if err := client.Call(myrpc.CometServicePushPrivate, pushArgs, &ret); err != nil {
		log.Fatal("client.Call(\"%s\", \"%v\", &ret) error(%v)", myrpc.CometServicePushPrivate, string(msgBytes), err)
		return InternalErr
	}

	log.Printf("Pushed a message to [key=%s] \n", key)

	return ret
}
