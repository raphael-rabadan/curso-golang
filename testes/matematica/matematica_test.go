package matematica

import "testing"

const erroPadrao = "Valor esperado espero %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	const valorEsperado float64 = 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
