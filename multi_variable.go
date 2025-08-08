package main
import (
	"fmt"
)

func multiVariable(){
	// Declaring Multiple Variable Declaration
	var a, b, c, d int = 10, 20, 30, 40

	// If `type` keword is not specified, you can declare different types of variables on the same line.
	var a,b = 6, "Hello"
	c, d := 7, "world"

	
	/* 
	Declaration in a Block
	Multiple variables can be declared can also be grouped together into a block for greater readability. 
	*/

	var (
		a int = 2
		b int = 3
		c string = "Abhishek"
		d bool = false
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}