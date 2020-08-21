//Hands-on exercise #1
//Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
//first name
//last name
//favorite ice cream flavors
//Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	favFlav   []string
}

func main() {
	p1 := person{
		firstName: "Zen",
		lastName:  "Dias",
		favFlav: []string{
			"chocolate",
			"strawberry",
			"blueberry",
		},
	}

	p2 := person{
		firstName: "Joy",
		lastName:  "Dias",
		favFlav: []string{
			"mango",
			"banana",
			"melancia",
		},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.favFlav {
		fmt.Println(i, v)

	}
	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for i, v := range p2.favFlav {
		fmt.Println(i, v)

	}

}
