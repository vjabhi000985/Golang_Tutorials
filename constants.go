package main
import (
	"fmt"
)

/*
	`const` keyword declares the variables "constant". It is unchangeable and read-only.
	# Syntax: `const CONSTNAME type = value` // Typed constant
	# Syntax: `const CONSTNAME = value` // Untyped constant
	# General Naming Rules:
	1. Constants should be usually written in Uppercase letters.
	2. It can be declared both inside or outside of a function. 
*/
const PI float32 = 3.14 

// Multiple constant declaration using Blocks
const (
	A int = 1
	B string = "Abhi"
	C float32 = 32.4
)

func contanstTest(){
	var radius float32

	fmt.Print("Enter the radius of a circle: ")
	fmt.Scan(&radius)

	circumference := 2 * PI * radius

	fmt.Println("The circumference of a circle with radius",radius,"units is",circumference)
	// fmt.Println()


}