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

func (c App) Platzi(user, action string) revel.Result {
	c.Validation.Required(user)
	c.Validation.Required(action)

	if c.Validation.HasErrors() {
		c.Flash.Error("El usuario no ha sido ingresado")
		return c.Redirect(App.Index)
	}

	switch action {
	case "refresh":
		return c.Redirect("/refresh?user=%s", user)
	case "polling":
		return c.Redirect("/polling/room?user=%s", user)
	}

	return c.Render()
}
