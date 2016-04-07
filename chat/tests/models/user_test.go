package models_test

import (
	"log"
	"platzi/chat/app/lib"
	"platzi/chat/app/models"
	"testing"
)

func TestSave(t *testing.T) {
	var user = models.User{
		Username: "Romi",
		Email:    "romi@iver.mx",
		Password: "",
	}
	var access = lib.NewAccess("users")
	var iface interface{}
	var err error
	if iface, err = access.Save(user); err != nil {
		t.Error("TestSave.Fail ", err)
		return
	}
	var id = iface.(int64)
	log.Printf("Id %v", id)
	if id == 0 {
		t.Error("Se esperaba un id mayor a 0")
		return
	}
	t.Log("Todo OK OK OK OK")
}
