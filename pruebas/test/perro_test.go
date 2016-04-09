package perro_test

import (
	"platzi/pruebas"
	"testing"
)

func TestBien(t *testing.T) {
	var p = pruebas.Perro{}
	if ladra := p.Ladra(); ladra != "guau" {
		t.Error("No ladró")
	}
}

func TestFalla(t *testing.T) {
	var p = pruebas.Perro{}
	if ladra := p.Ladra(); ladra != "miau" {
		t.Error("No ladró")
	}
}
