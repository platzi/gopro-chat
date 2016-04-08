package controllers

import (
	"github.com/revel/revel"
	"golang.org/x/net/websocket"
	"platzi/chat/app/lib/chat"
)

type WebSocket struct {
	*revel.Controller
}

func (c WebSocket) Room(user string) revel.Result {
	return c.Render(user)
}

func (c WebSocket) RoomSocket(user string, ws *websocket.Conn) revel.Result {
	// Se une al chat
	subscription := chat.Subscribe()
	defer subscription.Cancel()

	chat.Join(user)
	defer chat.Leave(user)

	// Se limpia el archivo
	for _, event := range subscription.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			// Se ha desconectado
			return nil
		}
	}

	// Seleccionamos entre mensajes de websocket y suscripciones.
	// Se requieren eventos de websocket dentro del canal.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			if err := websocket.Message.Receive(ws, &msg); err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Ahora se escuchan nuevos eventos desde el chat.
	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				// Se ha desconectado
				return nil
			}
		case msg, ok := <-newMessages:
			// Si el canal se ha cerrado, lo desconectamos
			if ok == false {
				return nil
			}

			// Se manda el mensaje en caso contrario.
			chat.Say(user, msg)
		}
	}
	return nil
}
