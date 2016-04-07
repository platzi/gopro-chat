package controllers

import (
	"github.com/revel/revel"
	//	 "platzi/chat/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var user = "Ivan"
	return c.Render(user)
}

func (c App) Platzi(user, messageType string) revel.Result {
	c.Validation.Required(user)
	c.Validation.Required(messageType)

	if c.Validation.HasErrors() {
		c.Flash.Error("Falta algún parámetro")
		return c.Redirect(App.Index)
	}

	switch messageType {
	case "refresh":
		return c.Redirect("/refresh?user=%s", user)
	case "polling":
		return c.Redirect("/polling/room?user=%s", user)
	}

	return c.Render()
}
