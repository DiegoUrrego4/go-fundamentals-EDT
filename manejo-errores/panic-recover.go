package manejo_errores

import "fmt"

// La función panic nos permite entrar en pánico = detener ejecución
// La función recover nos permite recuperarnos de un panic

func Division(dividen, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Me recuperé del pánico")
		}
	}()

	validateZero(divisor)
	fmt.Println(dividen / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("🤕 no puedes dividir por cero")
	}
}
