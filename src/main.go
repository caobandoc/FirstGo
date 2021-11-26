package main

import "fmt"

func main() {
	//Variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Carlos"
	edad := 21
	//Cuando sabes los tipos de datos
	fmt.Printf("%s tiene mas de %d años\n", nombre, edad)
	//Cuando no sabes los tipos de datos
	fmt.Printf("%v tiene mas de %v años\n", nombre, edad)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d años", nombre, edad)
	fmt.Println(message)

	// Tipo datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("edad: %T\n", edad)

}
