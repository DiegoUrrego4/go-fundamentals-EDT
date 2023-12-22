package arrays

import "fmt"

func Arrays() {
	// NOTA: Los arrays son de tamaño fijo, NO se pueden modificar
	// Declaración del array:
	//var flags [3]string
	//
	// Asignación de valores
	//flags[0] = "🇲🇽"
	//flags[1] = "🇺🇸"
	//flags[2] = "🇨🇴"

	// Los … se utilizan para que GO infiera el tamaño, dependiendo de los valores que le pasemos
	// Sin embargo, sigue siendo de tamaño fijo
	flags := [...]string{"🇲🇽", "🇺🇸", "🇨🇴"}

	fmt.Println(flags)

}
