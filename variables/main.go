package variables

import "fmt"

func ShowVariables() {
	apple, banana, orange := "🍎", "🍌", "🍊"
	apple, lemon := "manzana", "🍋" // Se puede reasignar el valor de la variable, pero DEBE ser del mismo tipo, en este caso: string.

	fmt.Println(apple, banana, orange, lemon)
}
