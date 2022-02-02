package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"hola","que","hace"}

	// _ se usa para escapar el indice
	for _, valor := range slice {
		fmt.Println(valor)
	}

	var input string
	fmt.Scan(&input)
	isPalindromo(input)
}

func isPalindromo(text string){
	var textReverse string
	text = strings.ToLower(text)
	
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}