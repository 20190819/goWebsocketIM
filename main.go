package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	_ "github.com/yangliang4488/webIM/routers"
)

const (
	APP_VER = "1.0.0"
)

func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)
	// 注册模板方法
	beego.AddFuncMap("i18n", i18n.Tr)

	
	beego.Run()
}
