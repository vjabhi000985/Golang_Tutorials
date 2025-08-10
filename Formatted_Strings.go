package main

import (
	"fmt"
)

func FormattedStringsTutorial(){
	fmt.Println("Formatted Strings Tutorial: ")

	// fmt.Printf() function formats the strings and sends them to the screen. 

	/*
		In Go, every data type has a unique format specifier.
		|-- integer	%d
		|-- float	%g or %f or specify the specific no. of digits : %0.10f
		|-- string	%s
    |-- bool	%t
	*/
	name := "Abhi"
	age := 25
	currentSalary := 15000
	isEmployed := true
	temperature := 67.999999202


	fmt.Printf("Hello, I am %s.\n", name)
	fmt.Printf("I am %d years old.\n", age)
	fmt.Printf("My current Salary is %d.\n", currentSalary)
	fmt.Printf("Is %s work in GTEC Computer Education? %t\n",name,isEmployed)
	fmt.Printf("Today's Temperature: %0.1f",temperature)

	/*
		Apart from that, we also, have two different specific format specifier.
		|-- %v - Value
		|-- %T - Data type
	*/
	fmt.Printf("\nValue: %v | Type: %T", isEmployed, isEmployed)
	// Output: Value: true | Type: bool
}