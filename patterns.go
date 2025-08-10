package main

import (
	"fmt"
)

func DrawPattern(){
	var n int = 5
	rows := 2 * n

	for i := 1; i < rows; i++ {
		if i <= n {
			for j := 1; j <= i; j++ {
				if (i == j) || ((i - j) % 2) == 0 {
					fmt.Print(" * ")
				} else{
					fmt.Print(" ")
				}
			}
		} else{
				for k := 1; k <= rows - i; k++ {
					if ((i - k) % 2) == 0 {
						fmt.Print(" * ")
					} else{
						fmt.Print(" ")
					}
				}
		}
		fmt.Println()
	}
}