package main

// Não tem operador tenário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	} else {
		return "reprovado"
	}
}

func main() {
	obterResultado(6.2)
}
