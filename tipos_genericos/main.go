package tipos_genericos

import "fmt"

// Estructura genérica

type Product[T uint | string] struct {
	ID          T
	Description string
	Price       float64
}

func GenericTypesFundamentals() {
	product1 := Product[uint]{ID: 1, Description: "Zapatos", Price: 30}
	product2 := Product[string]{ID: "2fc6e67a-28ca-4887-8315-4f2ef96533ea", Description: "Camisa", Price: 40}
	fmt.Println(product1)
	fmt.Println(product2)
}
