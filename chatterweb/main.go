package main

import (
	"github.com/astaxie/beego"
	_ "github.com/sf100/chatter/chatterweb/routers"
	_ "github.com/sf100/chatter/chatterweb/setting"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	beego.Run()
	beego.SetStaticPath("/static", "static")
}
