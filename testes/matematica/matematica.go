package matematica

import (
	"fmt"
	"strconv"
)

// Média é responsável por calcular a média aritimetica
func Media(numeros ...float64) (mediaArredonda float64) {
	total := 0.0
	for _, n := range numeros {
		total += n
	}
	media := total / float64(len(numeros))
	mediaArredonda, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return
}
