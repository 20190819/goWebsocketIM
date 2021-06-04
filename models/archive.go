package models

import "container/list"

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
	SIZE = 20
)

type Event struct {
	Type      EventType
	User      string
	Timestamp int
	Content   string
}

var archive = list.New()

func NewArchive(event Event) {
	if archive.Len() > SIZE {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}

func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}
