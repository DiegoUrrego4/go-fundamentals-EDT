package tipos_datos

import "fmt"

func ShowOtherTypes() {
	// bool string, numeric
	// El uint dependerá de nuestra arquitectura x32 o x64
	//var a float32 = 123.24

	// NOTA: no se pueden hacer operaciones entre tipos de datos diferentes
	//var a uint8 = 255
	//var b uint16 = 2550
	// ¿Cómo se puede solucionar? --> Casting, uint16(a)
	// _ --> Identificador blank _ , NO se puede usar el operador de asignación corta
	//_ = uint16(a) + b // El tipo de dato de la variable "a" nunca cambia

	// VALOR CERO:
	//var a string // Este por defecto es igual a un string vacío""
	//var a uint8 // Este por defecto es igual a un 0
	var a bool // Este por defecto es igual a falso

	//fmt.Printf("Tipo: %T, Valor: %v\n", c, c)
	fmt.Printf("Tipo: %T, Valor: %q\n", a, a)

}
