package main

import (
	"fmt"
	"strings"
)

func stringManipulation(){
	fmt.Println("String Manipulation")

	var name string

	name = "abhi"

	name = strings.ToTitle(name)

	fmt.Println("My name is", name)

	// How to use printf
	// %v - value and %T - type of the variable
	fmt.Printf("Hello %v",name)

}