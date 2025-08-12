package main

import (
	"fmt"
)

func GreatestNumber(){
	var x,y,z int

	x = 19
	y = 17
	z = 15

	if x > y && x > z{
		fmt.Println(x,"is the greatest number.")
	}else if y > z{
		fmt.Println(y,"is the greatest number.")
	}else{
		fmt.Println(z,"is the greatest number.")
	}
	// Greatest among two number
	// fmt.Print("Enter first number: ")
	// fmt.Scan(&x)

	// fmt.Print("\nEnter second number: ")
	// fmt.Scan(&y)

	// if x > y {
	// 	fmt.Printf("\n%d is greater", x)
	// } else if x < y{
	// 	fmt.Printf("\n%d is greater\n", y)
	// } else{
	// 	fmt.Printf("\n%d and %d are equal.", x, y)
	// }
}