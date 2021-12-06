package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
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

func main() {
	slice := []string{"hola", "que", "hace"}

	//for i, valor := range slice {
	//para tener el indice y el "valor" el paso del slice

	//for _, valor := range slice {
	//para solo tener en "valor" el paso del slice

	//for i := range slice {
	//para solo tener el indice del for
	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo("Ama")
}
