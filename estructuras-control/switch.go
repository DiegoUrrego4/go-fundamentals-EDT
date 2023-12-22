package estructuras_control

import "fmt"

func SwitchControlFlow() {
	character := "🦹🏻"
	canSearch := true

	// Agrupación de casos
	switch {
	case !canSearch:
		fmt.Println("NO está permitida la búsqueda")
	case character == "🦸🏻" || character == "🧞‍":
		fmt.Println("Es un súper héroe")
	case character == "🦹🏻" || character == "🧟‍":
		fmt.Println("Es un súper villano")
	default:
		fmt.Println("Es un personaje normal")

	}
}
