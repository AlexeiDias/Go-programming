//	Create a for loop using this syntax
//	for { }
//	Have it print out the years you have been alive

package main

import (
	"fmt"
)

func main() {
	i := 1976
	for {
		if i > 2017 {
			break
		}
		fmt.Println(i)
		i++

	}
}
