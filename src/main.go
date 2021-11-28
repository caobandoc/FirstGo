package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnvalue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a + 1, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "accion")

	value := returnvalue(3)
	fmt.Println("Value:", value)

	value1, _ := doubleReturn(2)
	fmt.Println("value1", value1)
}
