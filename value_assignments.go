package main 
import (
	"fmt"
)

func main(){
	// Variable declaration without initial value
	// Their value wiil be set to the default value of its type.  
	var a string
	var b int 
	var c float32
	var d bool

	// Value Assignment after Declaration
	var x int
	var y int  
	x = 10
	y = 20

	z := x + y 

	fmt.Println(a) // Default value of a string : "" - blank
	fmt.Println(b) // Default value of a int : 0
	fmt.Println(c) // Default value of a float : 0
	fmt.Println(d) // Default value of a bool : 
	
	fmt.Println("Sum of",x,"+",y,"=",z)
}