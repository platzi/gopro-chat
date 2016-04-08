package chat

import (
	"time"
)

type Event struct {
	Type      string // "join", "leave", or "message"
	User      string
	Timestamp int    // Unix timestamp (secs)
	Text      string // What the user said (if Type == "message")
}

func newEvent(typ, user, msg string) Event {
	return Event{typ, user, int(time.Now().Unix()), msg}
}

func Join(user string) {
	publish <- newEvent("join", user, "")
}

func Say(user, message string) {
	publish <- newEvent("message", user, message)
}

func Leave(user string) {
	publish <- newEvent("leave", user, "")
}

// Limpia un canal de cualquier mensaje
func drain(ch <-chan Event) {
	for {
		select {
		case _, ok := <-ch:
			if ok == false {
				return
			}
		default:
			return
		}
	}
}
