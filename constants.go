package main
import (
	"fmt"
)

const PI float32 = 3.14 

func contanstTest(){
	var radius float32

	fmt.Print("Enter the radius of a circle: ")
	fmt.Scan(&radius)

	circumference := 2 * PI * radius

	fmt.Println("The circumference of a circle with radius",radius,"units is",circumference)
	// fmt.Println()


}