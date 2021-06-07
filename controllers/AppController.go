package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

type AppController struct {
	baseControler
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

	// 匿名结构体
	var data struct {
		Uname string
		Tech  string
	}
	// 接收 json
	jsonStr := _this.Ctx.Input.RequestBody
	// 解析
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		beego.Error("json 解析错误")
	}
	_this.Data["json"] = joinResult{
		Code: 0,
		Msg:  "success",
		Data: struct {
			Uname string
			Tech  string
		}{data.Uname, data.Tech},
	}
	_this.ServeJSON()
}
