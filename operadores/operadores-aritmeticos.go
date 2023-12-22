package operadores

import "fmt"

func ArithmeticOperators() {
	// Operadores Aritméticos (), *, /, %, +, -
	var a = (2 + 3) * 2
	fmt.Println(a)
	// Operadores de asignación: =, +=, -=, *=, /=, %=
	var b int = 5
	b += 2
	fmt.Println(b)
	// declaración post-incremento y post-decremento: ++, --
	var c int = 6
	c++
	fmt.Println(c)
	// (No son una expresión sino una declaración)
}
