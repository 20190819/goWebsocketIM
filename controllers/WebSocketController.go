package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/yangliang4488/webIM/models"
)

type WebSocketController struct {
	baseControler
}

func (_this *WebSocketController) Get() {
	uname := _this.GetString("uname")

	_this.redirect302(uname)

	_this.TplName = "websocket.html"
	_this.Data["IsWebSocket"] = true
	_this.Data["UserName"] = uname
}

func (_this *WebSocketController) Join() {
	uname := _this.GetString("uname")

	_this.redirect302(uname)

	ws, err := websocket.Upgrade(_this.Ctx.ResponseWriter, _this.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); ok {
			http.Error(_this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		}
	} else {
		beego.Error("Cannot setup WebSocket connection:", err)
	}

	Join(uname, ws)
	defer Leave(uname)
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		chanPublish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}

func broadcastWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event", err)
		return
	}

	// 订阅者循环
	for sub := subscribers.Front(); sub != nil; sub.Next() {
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			// 发送消息失败
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				chanUnsubscribe <- sub.Value.(Subscriber).Name
			}
		}

	}
}

// 私有 func
func (_this *WebSocketController) redirect302(uname string) {
	if len(uname) == 0 {
		_this.Redirect("/", 302)
		return
	}
}
