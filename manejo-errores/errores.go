package manejo_errores

import (
	"errors"
	"fmt"
	"strconv"
)

// De esta forma podemos crear errores

// ErrNotFound Error personalizado para elemento no encontrado
var ErrNotFound = errors.New("not found")

var food = map[int]string{
	1: "🍕",
	2: "🍔",
}

func ErrorsFundamentals() {
	// Siempre debemos procurar manejar los errores
	num, err := strconv.Atoi("EDteam")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(num)
}

func Search(key string) (string, error) {
	num, err := strconv.Atoi(key)

	// %w: Verbo para envolver el error dentro de ese error
	if err != nil {
		return "", fmt.Errorf("strconv.Atoi(): %w", err)
	}

	emoji, err := findFood(num)

	if err != nil {
		return "", fmt.Errorf("findFood(): %w", err)
	}

	return emoji, nil
}

func findFood(id int) (string, error) {
	value, ok := food[id]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}
