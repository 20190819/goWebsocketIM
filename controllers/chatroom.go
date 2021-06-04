package controllers

import (
	"container/list"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/yangliang4488/webIM/models"
)

var (
	chanSubscribe   = make(chan Subscriber, 10)
	chanUnsubscribe = make(chan string, 10)
	chanPublish     = make(chan models.Event, 10)
	subscribers     = list.New()
)

type Subscription struct {
	Archive []models.Event
	New     <-chan models.Event
}

type Subscriber struct {
	Name string
	Conn *websocket.Conn
}

func newEvent(ep models.EventType, username, msg string) models.Event {
	return models.Event{ep, username, int(time.Now().Unix()), msg}
}

func Join(username string, ws *websocket.Conn) {
	chanSubscribe <- Subscriber{username, ws}
}

func Leave(username string) {
	chanUnsubscribe <- username
}

func chatroom() {
	for {
		select {
		// 订阅
		case sub := <-chanSubscribe:
			if !userExist(subscribers, sub.Name) {
				subscribers.PushBack(sub)
				chanPublish <- newEvent(models.EVENT_JOIN, sub.Name, "")
				beego.Info("new user:", sub.Name, ";websocket conn:", sub.Conn != nil)
			} else {
				beego.Info("old user:", sub.Name, ";websocket conn:", sub.Conn != nil)
			}
		case unsub := <-chanUnsubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					subscribers.Remove(sub)
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub)
					}
					chanPublish <- newEvent(models.EVENT_LEAVE, unsub, "")
					break
				}
			}
		}
	}
}

// 判断用户是否存在
func userExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}
