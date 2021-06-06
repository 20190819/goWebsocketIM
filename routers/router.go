package routers

import (
	"github.com/astaxie/beego"
	"github.com/yangliang4488/webIM/controllers"
)

func init() {
	beego.Router("/", &controllers.AppController{})
	beego.Router("/join", &controllers.AppController{}, "post:Join")

	// websocket
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
}
