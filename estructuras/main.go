package estructuras

import "fmt"

type Person struct {
	Name        string
	Age         uint8
	HasChildren bool
}

func Structures() {
	// Para crear instancias
	alexys := Person{Name: "Alexys", Age: 42, HasChildren: false}
	//beto := Person{"Beto", 33, true}
	//alejo := Person{Age: 32}

	// Desde el puntero podemos hacer modificaciones a donde estemos apuntando
	// Si estamos trabajando con punteros no es necesario la desreferenciación
	alvaro := &alexys
	alvaro.Age = 60

	// +v es para que también imprima el nombre en la estructura
	//fmt.Printf("%+v\n", beto.HasChildren)
	//fmt.Printf("%+v\n", alejo.Age)
	//fmt.Printf("%+v\n", alexys.Name)
	fmt.Printf("%+v\n", alexys)
	fmt.Printf("%+v\n", alvaro)

}
