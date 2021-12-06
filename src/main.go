package main

import (
	pk "curso_golang/src/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ford"
	myCar.Year = 2021
	fmt.Println(myCar)

	pk.PrintMessage("Hola mundo")

	pk.printMessage("Hola mundo")

}
