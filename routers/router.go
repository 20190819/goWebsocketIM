package routers

import (
	"github.com/astaxie/beego"
	"github.com/yangliang4488/webIM/controllers"
	"github.com/yangliang4488/webIM/routers/middlewares"
)

func init() {

	// 允许跨域
	middlewares.EnableCors()
	beego.Router("/login", &controllers.AppController{}, "post:Login")	// 登录
	beego.Router("/logout", &controllers.AppController{}, "post:Logout")	// 退出

	// websocket
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")	// 进入聊天室
}
