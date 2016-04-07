package controllers

import (
	"github.com/revel/revel"
)

type Refresh struct {
	*revel.Controller
}

func (c *Refresh) Index() revel.Result {
	// TODO: Agregar lógica para usuario que se une
	return c.Render()
}
func (c *Refresh) Room() revel.Result {
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
