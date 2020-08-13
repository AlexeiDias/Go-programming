// create a typed and untyped constant. Print their values


package main

import (
	"fmt"
)

const (
	a     = 44
	b int = 43
)

func main() {

	fmt.Println(a,b)
	fmt.Println(b)

}