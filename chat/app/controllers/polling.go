package controllers

import "github.com/revel/revel"

type Polling struct {
	*revel.Controller
}

func (c *Polling) Room(user string) revel.Result {
	// TODO: Agregar lógica del poll del chat
	return c.Render()
}

func (c *Polling) Say(user, message string) revel.Result {
	// TODO: Agregar mensaje para poner el en chat
	return nil
}

func (c *Polling) WaitMessages(received int) revel.Result {
	// TODO: publicar los mensajes en json
	return c.RenderJson("")
}

func (c *Polling) Leave(user string) revel.Result {
	// TODO: Lógica para salir
	return c.Redirect(App.Index)
}
