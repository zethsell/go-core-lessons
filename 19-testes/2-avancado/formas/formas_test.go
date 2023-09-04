package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		arearRecebida := ret.area()

		if areaEsperada != arearRecebida {
			t.Fatalf("a area recebida %f é diferente da esperada %f", arearRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * math.Pow(10, 2))
		arearRecebida := circ.area()

		if areaEsperada != arearRecebida {
			t.Fatalf("a area recebida %f é diferente da esperada %f", arearRecebida, areaEsperada)
		}
	})
}
