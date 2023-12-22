package arrays

import "fmt"

func Maps() {
	// Estructuras de clave - valor
	music := make(map[string]string)
	music["guitar"] = "🎸"
	music["violin"] = "🎻"

	fmt.Println(music)

	tech := map[string]string{
		"computer": "💻",
		"mouse":    "🖱️",
	}

	fmt.Println(tech)

	// eliminar elementos:
	delete(tech, "computer")
	fmt.Println(tech)

	// Lectura de valores:
	content, ok := music["guitar"]
	fmt.Println(content, ok)

}
