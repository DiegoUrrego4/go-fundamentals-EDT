package arrays

import "fmt"

func SlicesFundamentals() {
	// Slice: Son apuntadores a un Array, no poseen datos
	things := [7]string{"🚓", "🚙", "🚕", "🏎️", "🚒", "🚨", "🎈"}

	// Slices:
	// El indice final es excluyente y el inicial incluyente
	cars := things[:5] // 🚓 🚙 🚕 🏎️ 🚒]
	red := things[3:]  // 🏎️ 🚒 🚨 🎈
	all := things[:]
	red[1] = "🚑" // Esto modifica todos los demás slices

	fmt.Println("things", things)
	fmt.Println("cars", cars)
	fmt.Println("red", red)
	fmt.Println("all", all)

	fmt.Println("Red[0]", red[0])
}

func SlicesDeeper() {
	// len(): # de elementos en el slice
	// cap(): # de elementos del array origen, a partir del índice donde se creó

	//animals := [5]string{"🦍", "🦮", "🐕", "🦜", "🐘"}
	//pets := animals[1:3] // "🦮", "🐕"

	// Asignación de valores en slices
	// NOTA: Este slice sigue haciendo referencia al array original, por lo que los elementos también cambiarán allí
	//pets = append(pets, "🐈", "🐩", "🐸")

	// Cuando creamos el array pets referenciamos el array animals Array[4]{"🦮", "🐕", "🦜", "🐘"}
	// Al exceder la capacidad Go crea un nuevo array y agrega los elementos new Array[8]{"🦮", "🐕", "🐈", "🐩", "🐸"}

	//fmt.Println("capacidad pets:", cap(pets))
	//fmt.Println("animals:", animals)
	//fmt.Println("pets:", pets)
	//fmt.Println("tamaño pets:", len(pets))

	//pets := []string{"🦮", "🐕"}
	// make() debemos indicarle el tipo, tamaño y capacidad
	//pets := make([]string, 0, 3)
	//pets = append(pets, "🐈", "🐩", "🐒", "🐍")

	pets := []string{}

	fmt.Println("pets:", pets)
	fmt.Println("tamaño pets:", len(pets))
	fmt.Println("capacidad pets:", cap(pets))
	fmt.Println("Valor cero:", pets == nil)

}
