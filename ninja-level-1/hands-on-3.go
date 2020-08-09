// Using the code from the previous exercise,
// At the package level scope, assign the following values to the three variables
// for x assign 51
// for y assign “Ace Ventura”
// for z assign false
// for k assign true
// in func main
// use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
// print out the value stored by variable “s”
		
package main

import (
	"fmt"
)

var x int = 51
var y string = "Ace Ventura"
var z bool = false
var k bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\t%v", x, y, z, k)
		// "S"	in Sprint prints into a string
		// %v	the value in a default format
		// \t	adds a tab 
		// \n	when is present adds a new line  
	fmt.Println(s)
	
	
}
