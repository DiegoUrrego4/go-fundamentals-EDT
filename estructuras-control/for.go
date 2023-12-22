package estructuras_control

import "fmt"

func ForLoop() {
	// NOTA: En GO solo existe el for. Casos como Do - While, While, …, etc no existen
	//i := 1
	//for {
	//	if i == 6 {
	//		break
	//	}
	//	fmt.Println(i)
	//	i++
	//}

	// ! NOTA: El range devuelve dos valores: índice y contenidoElemento
	// Slices:
	// * Opción1: Podemos pasar el slice o array directamente:
	//for i, v := range []string{"🍕", "🍔", "🍎", "🌭"} {
	//	fmt.Println("indice:", i, "valor:", v)
	//}

	// Slices #2:
	// Opción 2: Podemos declararlo fuera del for range
	//numbers := []uint8{2, 4, 6, 8}
	//
	//for i := range numbers {
	//	numbers[i] *= 2
	//}
	//
	//fmt.Println(numbers)

	// Mapas:
	//food := map[string]string{
	//	"pizza":      "🍕",
	//	"hamburguer": "🍔",
	//	"apple":      "🍎",
	//	"hotdog":     "🌭",
	//}
	//
	//for key, value := range food {
	//	fmt.Println("key:", key, "value:", value)
	//}

	// Strings
	// NOTA: En este caso debemos castear la letra, pues si no lo hacemos obtendremos su valor unicode
	for i, letter := range "EDteam" {
		fmt.Println("indice:", i, "valor:", string(letter))
	}

}
