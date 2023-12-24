package manejo_errores

import (
	"fmt"
	"os"
)

// La función defer como su nombre lo dice se usa para diferir algo
// Diferir = aplazar. ¿Aplazar qué? La ejecución de una función

func DeferBasic() {
	// Ejemplo #1:
	//defer fmt.Println(3) // Con esto aplazamos la ejecución
	//fmt.Println(1)
	//fmt.Println(2)
	// NOTA: La función se aplaza hasta que la función principal (en este caso DeferBasic) retorne
	// El resultado será: 1,2,3 de igual forma, pues el defer se ejecuta una vez la función retorna

	// Ejemplo #2:
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)

	// Cuando hay multiples funciones diferidas se agregan a una pila y estas se ejecutaran de última a primera
	// En este caso: 3,2,1

	// Ejemplo #3:
	//num := 4
	//defer fmt.Println("imprime", num)
	//
	//num = 10
	//defer fmt.Println(num)

	// Pero … ¿Para qué me sirve?
	// Limpiar recursos, cerrar archivos, cerrar conexiones de red o cerrar controladores de bases de datos

	// Ejemplo #4:
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close() // Esto siempre se ejecutará sin necesidad de repetir código

	_, err = file.Write([]byte("Hello gophers"))
	if err != nil {
		//file.Close()
		fmt.Println(err)
		return
	}
	//file.Close()
}
