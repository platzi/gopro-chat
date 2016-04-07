package controllers

import (
	"github.com/revel/revel"
	"platzi/chat/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var user = models.User{
		Username: "Romi",
		Email:    "romi@iver.mx",
		Password: "hola",
	}
	revel.INFO.Printf("User %v", user.Username)
	return c.Render()
}

func (c App) Platzi() revel.Result {
	revel.INFO.Printf("Dentro de Platzi")
	return c.Render()
}
