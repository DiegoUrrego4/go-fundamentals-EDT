package operadores

import "fmt"

func ComparisonOperators() {
	// Operadores Comparación: >, <, ==, !=, >=, <=
	fmt.Println(6 >= 4)
	// Operadores lógicos &&, ||
	var age uint8 = 70
	fmt.Println("Es adulto?", age >= 18 && age <= 60)
	fmt.Println("Es niño o anciano?", age < 18 || age > 60)
	// Operador lógico Unario: !
	fmt.Println(!(4 == 4))

}
