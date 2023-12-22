package funciones

import (
	"fmt"
	"strings"
)

func Greet(firstName, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// Funciones con parámetros de referencia:

// ToUpperCase función para devolver texto en mayúsculas
// text Parámetro por referencia = puntero
func ToUpperCase(text *string) {
	*text = strings.ToUpper(*text)
}

// Funciones con retorno:

func Suma(a, b int) int {
	return a + b
}

// Convert Convierte textos a minúsculas y mayúsculas
// Devuelve el valor lower: minúsculas, upper: mayúsculas
func Convert(text string) (lower string, upper string) {
	lower = strings.ToLower(text)
	upper = strings.ToUpper(text)

	return // Esto solo es recomendable hacerlo cuando las funciones son pequeñas
}

// Funciones con funciones:

func Filter(nums []int, callback func(int) bool) []int {
	result := make([]int, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}

// Funciones que devuelven funciones:

func Sum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// Funciones variáticas

func SumWithManyParameters(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}
