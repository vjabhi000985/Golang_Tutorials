package main

import (
	"fmt"
)

func UserInput() {
	var name string 
	var currentSalary float32


	// Scan()
	/*
		The Scan() function takes input from the user. However, it can only take input values up to a space.
	*/
	fmt.Println("Please Enter your name: ")
	// fmt.Scan(&name) // Bittu Kumar
	// fmt.Println("My name is",name) // Bittu
	
	// Multiple Input using Scan()
	fmt.Scan(&name,&currentSalary)

	fmt.Println("My name is",name) // Bittu
	fmt.Println("My current salary is",currentSalary) // 12000.00

	// Scanln()

	//Scanf()
}
