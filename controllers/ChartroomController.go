package controllers

import (
	"container/list"
	"fmt"
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
	waitinglist     = list.New()
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

func JoinRoom(username string, ws *websocket.Conn) {
	fmt.Println("chatroom join")
	chanSubscribe <- Subscriber{Name: username, Conn: ws}
}

func LeaveRoom(username string) {
	chanUnsubscribe <- username
}

func chatroom() {
	for {
		select {
		case sub := <-chanSubscribe:
			// 订阅
			if !userExist(subscribers, sub.Name) {
				subscribers.PushBack(sub)
				chanPublish <- newEvent(models.EVENT_JOIN, sub.Name, "")
				beego.Info("new user:", sub.Name, ";websocket conn:", sub.Conn != nil)
			} else {
				beego.Info("old user:", sub.Name, ";websocket conn:", sub.Conn != nil)
			}
		case event := <-chanPublish:
			// 广播消息
			for ch := waitinglist.Back(); ch != nil; ch.Prev() {
				ch.Value.(chan bool) <- true
				waitinglist.Remove(ch)
			}
			broadcastWebSocket(event)
			models.NewArchive(event)
			if event.Type == models.EVENT_MESSAGE {
				beego.Info("message from : ", event.User, "message: ", event.Content)
			}
		case unsub := <-chanUnsubscribe:
			// 取消订阅
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

// goroutine 异步
func init() {
	go chatroom()
}

// 判断用户是否存在
func userExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}
