package punteros

import "fmt"

func Pointers() {
	// Puntero: Variables que almacenan la dirección en memoria de un valor
	var color string = "🔴"
	// Para declarar un puntero (dirección en memoria):
	var pointerColor *string = &color
	// Para modificar el valor de un puntero:
	*pointerColor = "🟦"

	// Para ver dirección en memoria usamos el & antes de la variable
	// Para dereferenciarlo lo hacemos usando *

	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v\n", color, color, &color)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", pointerColor, pointerColor, *pointerColor)

}
