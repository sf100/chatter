package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/sf100/chatter/chatterweb/models"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*user控制器*/
type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.TplNames = "login.html"
}

func (this *UserController) Post() {

	//返回结果
	ret := Result{
		Success: true,
	}
	uname := this.Input().Get("userName")
	pwd := this.Input().Get("password")
	user := models.GetUserByName(uname)

	if user != nil && user.Password == pwd {
		//登陆消息系统
		token := loginMsg(uname, pwd)
		fmt.Println(token)
		if len(token) > 0 {
			ret.Data = user
			this.SetSession("token", token)
			this.SetSession("user", user)
		} else {
			ret.Success = false
			ret.Msg = "服务器繁忙请稍后再试"
		}

	} else {
		ret.Success = false
		ret.Msg = "用户名或密码错误"
	}

	this.Data["json"] = ret
	this.ServeJson()
	this.StopRun()
}

func (this *UserController) Register() {
	this.TplNames = "register.html"
}

func (this *UserController) DoRegister() {

	fmt.Println("------------------注册-----------------------")
	//返回结果
	ret := Result{
		Success: true,
	}
	uname := this.Input().Get("userName")
	pwd := this.Input().Get("password")
	if len(uname) == 0 || len(pwd) == 0 {
		ret.Success = false
		ret.Msg = "用户名和密码不能为空"
		return
	} else if len(pwd) < 6 {
		ret.Success = false
		ret.Msg = "密码必须6位以上"

	}

	if !models.Register(uname, pwd) {
		ret.Success = false
		ret.Msg = "服务器繁忙，请稍后在试！"
	}

	/*注册成功，添加会话session*/
	if ret.Success {
		this.SetSession("user", "222")
	}
	this.Data["json"] = ret
	this.ServeJson()
	this.StopRun()

}

/*登录消息系统*/
func loginMsg(userName, pwd string) string {

	bodyData := []byte(`
		{
		    "baseRequest" : {
		        "uid": "",
		        "deviceID": "",
		        "deviceType": "web",
		        "token": ""
		    },
			"userName": "` + userName + `",
		    "password": "` + pwd + `"
		}

	`)
	body := bytes.NewReader(bodyData)
	url := beego.AppConfig.String("msgLoginAddr") //获取消息服务地址
	res, err := http.Post(url, "text/plain;charset=UTF-8", body)

	if err != nil {
		beego.Error(err)
		return ""
	}
	defer res.Body.Close()

	if res.StatusCode == 200 { //请求成功

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			beego.Error("读取response信息异常：[s%]", err)
			return ""
		}
		var args map[string]interface{}
		if err = json.Unmarshal(resBody, &args); err != nil {
			beego.Error("读取response信息转换为异常：[%s]", err)
			return ""
		}
		baseResponse := args["baseResponse"].(map[string]interface{})
		ret := baseResponse["ret"].(float64)
		if ret != 0 {
			beego.Error("服务器处理异常，返回吗：[%d]", int(ret))
			return ""
		}
		return args["token"].(string)
	} else {
		beego.Error("请求失败,错误码:[%d]", res.StatusCode)
		return ""
	}
	return ""
}

func (this *UserController) Logout() {
	//清除缓存
	this.DelSession("user")
	this.Redirect("/", 302)
}
