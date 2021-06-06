package controllers

import (
	"fmt"
)

type AppController struct {
	baseControler
}

func (_this *AppController) Get() {
	_this.TplName = "welcome.html"
	fmt.Println("welcome")
}

func (_this *AppController) Join() {
	uname := _this.GetString("uname")
	tech := _this.GetString("tech")

	if len(uname) == 0 {
		_this.Redirect("/", 302)
	}

	if tech == "websocket" {
		_this.Redirect("/ws?uname="+uname, 302)
	} else {
		_this.Redirect("/", 302)
	}
}
