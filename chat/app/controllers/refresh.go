package controllers

import (
	"github.com/revel/revel"
	"platzi/chat/app/lib/chat"
	"platzi/chat/app/routes"
)

type Refresh struct {
	*revel.Controller
}

func (c *Refresh) Index(user string) revel.Result {
	chat.Join(user)
	return c.Redirect(routes.Refresh.Room(user))
}

func (c *Refresh) Room(user string) revel.Result {
	// TODO: Lógica para identificar el usuario en el chat
	return c.Render()
}

func (c *Refresh) Say() revel.Result {
	// TODO: Lógica  para escribir el texto del usuario
	return c.Render()
}
func (c *Refresh) Leave() revel.Result {
	// TODO: Lógica para dejar el salón
	return c.Render()
}
