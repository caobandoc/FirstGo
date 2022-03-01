package main

import (
	"fmt"
)

// Cuando esta al lado derecho de chan(<-) significa de salida
// Cuando esta al lado izquierdo de chan(<-) significa de entrada
func say(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c)
}
