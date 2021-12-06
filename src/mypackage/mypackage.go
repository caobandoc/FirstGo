package mypackage

import "fmt"

// CarPublic Car con acceso publico
// La inicial indica si va a ser publico o privado
// Mayuscula para publico
type CarPublic struct {
	Brand string
	Year  int
}

//minuscula para privado
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}

func printMessage(text string) {
	fmt.Println(text)
}
