// Hands-on exercise #1
// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p := person{
		first:      "Alexei",
		last:       "Dias",
		favFlavors: []string{"filhos", "surf"},
	}
	fmt.Println(p)
}
