package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type AppController struct {
	baseControler
}

func (_this *AppController) Get() {
	_this.TplName = "welcome.html"
	fmt.Println("welcome")
}

type joinResult struct {
	Code int8
	Msg  string
	Data struct {
		Uname string
		Tech  string
	}
}

func (_this *AppController) Join() {
	_this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")

	uname := _this.GetString("uname")
	tech := _this.GetString("tech")

	beego.Info(uname, tech)

	_this.Data["json"] = joinResult{
		Code: 0,
		Msg:  "success",
		Data: struct {
			Uname string
			Tech  string
		}{uname, tech},
	}
	_this.ServeJSON()
}
