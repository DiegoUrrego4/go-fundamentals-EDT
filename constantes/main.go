package constantes

import "fmt"

// Notas importantes:
// Los valores constantes se deben asignar al momento de declarar la variable
// No se puede usar el operador de asignación de variable corta
// Al dejar las variables fuera de la función quedan tipo global en el paquete
const (
	os     = "Mac OS"
	domain = "ed.team"
)

/*
Con el identificador iota podemos autoincrementar las variables, sin necesidad de escribir una por una
A tener en cuenta:
iota inicia en 0, por lo que si queremos que comience en un número específico debemos sumar
las constantes NO necesitan ser usadas en su totalidad
*/
const (
	January = iota + 1 // En este caso iniciaría en 1
	February
	March
	April
	May
	June
)

func ShowConstants() {
	fmt.Println(May)
}
