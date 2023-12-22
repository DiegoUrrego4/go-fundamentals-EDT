package estructuras_control

import "fmt"

func IfElseControlFlow() {

	// NOTA: Si solo vamos a usar la variable dentro del scope debemos declararla adentro del if
	if character := "🦸🏻"; character == "🦸🏻" {
		fmt.Println("Es un super héroe")
	} else if character == "🦹🏻‍" {
		fmt.Println("Es un super villano")
	} else {
		fmt.Println("Es un personaje normal")
	}

}
