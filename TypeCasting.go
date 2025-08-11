package main
import (
	"fmt"
)

func TypeCastingExample(){

	var floatValue float32 = 5.45

  // type conversion from float to int
  var intValue int = int(floatValue)
 
  fmt.Printf("Float Value is %g\n", floatValue)
  fmt.Printf("Integer Value is %d\n", intValue)

	var x int 
	x = 4

	// type conversion from int to float
	var y float32 = float32(x)

	fmt.Println(y)
	fmt.Printf("Value : %f | Type : %T",y,y)
}