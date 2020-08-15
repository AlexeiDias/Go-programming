//	Create a program that uses a switch statement with the switch expression specified as a variable TYPE string with the IDENTIFIER "favSport"

package main

import (
	"fmt"
)

func main() {
	favSport := "golf"
	switch favSport {
	case "futebol":
		fmt.Println("i like it")
	case "golf":
		fmt.Println("nao gosto")

	}
}
