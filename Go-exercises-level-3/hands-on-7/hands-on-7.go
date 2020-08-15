//	Building on the previous hands-on exercise, create a program that uses "else if" and "else"

package main

import (
	"fmt"
)

func main() {
	x := "James Bond"
	if x == "ding dong" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BondBondBond")
	} else {
		fmt.Println("neither")
	}
}
