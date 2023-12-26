package fn_no_genericas_any

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type MyInt int
type MyIntV2 int

func NoGenerics() {
	var num1 MyInt = 2
	var num2 MyInt = 6

	var num3 MyIntV2 = 3
	var num4 MyIntV2 = 4

	printList("EDteam", "gopher", "Go desde cero")
	printList(1, 2, 8)
	fmt.Println("Suma Int:", sum(2, 4, 67))
	fmt.Println("Suma Int MyInt:", sum(num1, num2))
	fmt.Println("Suma Int:", sum(num3, num4))
	fmt.Println("Suma Float:", sum(2, 4.6, 67.1))

	fmt.Println("INCLUDES V1", Includes([]string{"a", "b", "c"}, "c")) // true
	fmt.Println("INCLUDES V2", Includes([]string{"a", "b", "c"}, "d")) // false
	fmt.Println("INCLUDES V2", Includes([]int{1, 12, 24}, 24))         // true

	fmt.Println(filter([]int{2, 12, 23, 98, 21, 79}, func(num int) bool {
		return num < 20
	}))

}

// Las funciones genéricas permiten trabajar con diferentes tipos de datos
// any: cualquier tipo de dato
func printList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}

type Number interface {
	~int | ~float64 | ~float32 | ~uint
}

// Parámetros de tipo:
// Si solo queremos sumar números NO deberíamos usar any
// constrain tipo arbitrario: sum[T int]
// constrain unión de elementos. Elementos separados por pipe
// Elementos de aproximación: ~int: tipo int Y SUS DERIVADOS
func sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// comparable, permite comparar valores

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func filter[T constraints.Ordered](nums []T, callback func(T) bool) []T {
	result := make([]T, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}
