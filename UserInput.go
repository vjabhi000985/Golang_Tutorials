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
	fmt.Scan(&name) // Bittu Kumar
	fmt.Println("My name is",name) // Bittu
	
	// Multiple Input using Scan()
	fmt.Scan(&name,&currentSalary)

	fmt.Println("My name is",name) // Bittu
	fmt.Println("My current salary is",currentSalary) // 12000.00
	
	// Scanln()
	/*
		We use the Scanln() function to get input values up to the new line. When it encounters a new line, it stops taking input values. 
		*/
		fmt.Scanln(&name,&currentSalary)
		fmt.Println("My name is",name) // Bittu
		fmt.Print("Cuurent Salary: ",currentSalary) // 0

		
		
		//Scanf()
		/*
			The fmt.Scanf() function is similar to Scanln() function. The only difference is that Scanf() takes inputs using format specifiers. 
		*/
		fmt.Scanf("%s", &name)
		fmt.Scanf("%d", &currentSalary)
		
		fmt.Println("Name  : ", name)
		fmt.Println("Salary: ", currentSalary)
}
