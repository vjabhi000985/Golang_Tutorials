package main

import (
	"fmt"
)

func GreatestNumber(){
	var x,y int

	fmt.Print("Enter first number: ")
	fmt.Scan(&x)

	fmt.Print("\nEnter second number: ")
	fmt.Scan(&y)

	if x > y {
		fmt.Printf("\n%d is greater", x)
	} else if x < y{
		fmt.Printf("\n%d is greater\n", y)
	} else{
		fmt.Printf("\n%d and %d are equal.", x, y)
	}
}