package main

import (
	"fmt"
)

func Operators(){
	var x,y int
	
	/*
		1. Arithmetic Operators: Used for basic math operations.

			Example a=10, b=3
			+	Addition: 	a + b	13
			-	Subtraction: 	a - b	7
			*	Multiplication: 	a * b	30
			/	Division:	a / b	3 (integer division)
			%	Modulus: (remainder)	a % b	1
	
	*/
	x = 1
	y = 2
	
	fmt.Println(x+y)
	// fmt.Print(x-y)
	// fmt.Print(x*y)
	// fmt.Print(x/y)
	// fmt.Print(x%y)

	/*
		2. Assignment Operators: Used to assign values to variables.
		=, +=, -=, *=, /=, %=
	*/
	
	y += 2
	fmt.Println(y) // 4

	y *= 1
	fmt.Println(y) // 4

	y -= 4
	fmt.Println(y) // 0

	y /= 2
	fmt.Println(y) // 0
	
	y %= 3
	fmt.Println(y) // 0
	
	/*
		3. Comparison Operators: Used to compare two values.
		==, !=, <, >, <=, >=
	*/

	fmt.Println(x != y) // true
	fmt.Println(x == y) // false
	fmt.Println(x < y) // false
	fmt.Println(x > y) // true
	fmt.Println(x <= y) // false
	fmt.Println(x >= y) // true

	/*
		4. Logical Operators: Used with boolean values.
		&& - Logical AND
		|| - Logical OR
		! - Logical NOT
	*/

	isRain := false
	isCold := true
	fmt.Println(isRain && isCold) // false
	fmt.Println(isRain || isCold) // true
	fmt.Println(!isCold) //false

	
}