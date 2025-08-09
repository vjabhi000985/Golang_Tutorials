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

}