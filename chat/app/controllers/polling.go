package controllers

import (
	"github.com/revel/revel"
	"platzi/chat/app/lib/chat"
	"platzi/chat/app/routes"
)

type Polling struct {
	*revel.Controller
}

func (c *Polling) Room(user string) revel.Result {
	revel.INFO.Printf("Polling.Room() %v", user)
	chat.Join(user)
	return c.Render(user)
}

func (c *Polling) Say(user, message string) revel.Result {
	revel.INFO.Printf("Usuario: %v message: %v", user, message)
	chat.Say(user, message)
	return c.Redirect(routes.Polling.Room(user))
}

func (c *Polling) WaitMessages(lastReceived int) revel.Result {
	subscription := chat.Subscribe()
	defer subscription.Cancel()
	var events []chat.Event
	for _, event := range subscription.Archive {
		if event.Timestamp > lastReceived {
			events = append(events, event)
		}
	}

	if len(events) > 0 {
		return c.RenderJson(events)
	}

	event := <-subscription.New
	return c.RenderJson([]chat.Event{event})
}

func (c *Polling) Leave(user string) revel.Result {
	chat.Leave(user)
	return c.Redirect(App.Index)
}
