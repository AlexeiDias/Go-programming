//Hands-on exercise #2
//Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

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
		lastName:  "Days",
		favFlav: []string{
			"mango",
			"banana",
			"melancia",
		},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(m)

	for _, v := range m {

		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favFlav {
			fmt.Println(i, val)
		}
		fmt.Println("------------")
	}

}
