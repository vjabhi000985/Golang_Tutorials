/*
// In Golang, every program is a part of package.
// We define it using `package` keyword.
package main

// Let us import the `fmt` package that is used for formatting and basic printing.
import ("fmt")

// Go ignores white spaces so we can use white spaces to increase readability.

// Function main is the start of code and any code enclosed within the curly braces can be executed
func main(){

		// `Println` is a method inside the `fmt` package and it is used to output/print text.
		fmt.Println("Hello world")
	}
*/
package main

import (
	"fmt"
)

func variables(){
	// Declaring Variables Syntax : `var variable_name type = value
	var name string = "Abhi" // Type is string

	var age = 25 // Type is infered

	x := "Abhishek" // Type is infered

	// using `:=` : variable_name := value. It can't be left blank. 
	// fmt.Println("Hello world")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println("You name is",x)
}

// We can use `;` semilion to create compact statements.
// Means that we can write all out code in one line.
// package main; import ("fmt"); func main(){ fmt.Println("Hello world");}

/*
	`//` - single line comment
*/

/* This is a multi-line comment. */
