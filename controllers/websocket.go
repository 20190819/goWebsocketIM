package controllers

import (
	"encoding/json"
	"net/http"
	"webIM/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	baseControler
}

func (this *WebSocketController) Get() {
	username := this.GetString("username")

	this.redirect302(username)

	this.TplName = "websocket.html"
	this.Data["IsWebSocket"] = true
	this.Data["Username"] = username
}

func (this *WebSocketController) Join() {
	username := this.GetString("username")

	this.redirect302(username)

	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); ok {
			http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		}
	} else {
		beego.Error("Cannot setup WebSocket connection:", err)
	}

	// Join(username, ws)
	// defer Leave(username)
	// for {
	// 	_, p, err := ws.ReadMessage()
	// 	if err != nil {
	// 		return
	// 	}
	// 	publish <- newEvent(models.EVENT_MESSAGE, username, string(p))
	// }
}

func (this *WebSocketController) redirect302(username string) {
	if len(username) == 0 {
		this.Redirect("/", 302)
		return
	}
}

func broadcastWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
	}

	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		ws := sub.Value.(Subscriber).Conn
	}

}
