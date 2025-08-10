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

	/*
		5. Bitwise Operators: Work on binary representations.
		& (AND) : Sets a bit to 1 if both corresponding bits are 1.
		| (OR)  : Sets a bit to 1 if at least one corresponding bit is 1.
		^ (XOR) : Sets a bit to 1 if bits are different (1 and 0 or 0 and 1).
		&^ (AND NOT) : Clears (sets to 0) the bits in the first operand where the second operand has bits set to 1.
		>> (RIGHT SHITF) : Shifts bits to the right, discarding bits on the right (divides by 2ⁿ).
		<< (LEFT SHIFT)  : Shifts bits to the left, adding zeros on the right (multiplies by 2ⁿ).
	*/

	fmt.Println(x & y) // 0
	fmt.Println(x | y) // 1
	fmt.Println(x ^ 2) // 3
	fmt.Println(x &^ 2) // 1
	fmt.Println(x << 2) // 4
	fmt.Println(7 >> 2) // 1

	/*
		6. Increment/Decrement operators : They are generally used with loops and their value cannot be assigned to other variables.
		Eg. x++, y--
	*/


}