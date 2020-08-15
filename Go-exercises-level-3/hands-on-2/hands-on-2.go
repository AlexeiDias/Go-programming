// Print every rune code point of the uppercase alphabet three times. 

package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		//fmt.Printf("the outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%v\t%#U\n", i, i)
		}
	}
}
