//	Using iota, create 4 constant for the last 4 years
//	Print the constan value

//	Create a variable type spring using a raw string literal.
//	Print it

package main

import (
	"fmt"
)

const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
	d = 2017 + iota
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
