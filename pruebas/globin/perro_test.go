package globin_test

import (
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"platzi/pruebas"
	"testing"
)

func Inicia(t *testing.T) (g *goblin.G) {
	g = goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	return g
}

var ladraOk = `guau`
var ladraFail = `miau`

func TestBien(t *testing.T) {
	g := Inicia(t)
	g.Describe("Probando al perro", func() {
		var p = pruebas.Perro{}
		var l = p.Ladra()
		g.It("Si ladra", func() {
			Ω(l).Should(Equal(ladraOk))
		})
		g.It("No maulla", func() {
			Ω(l).ShouldNot(Equal(ladraFail))
		})
	})
}
