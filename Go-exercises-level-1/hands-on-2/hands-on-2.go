// Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).

// identifier “x” type int
// identifier “y” type string
// identifier “z” type bool
// identifier "k" type bool = true

// in func main
// print out the values for each identifier.

// The compiler assigned values to the variables. What are these values called?
		// the values that the compiler assigns to the variables are called Type.
		// there are diferent types of Types or vairables these are some of them
			// string
				//A string type represents the set of string values. Its value is a sequence of bytes. Strings are immutable types that is once created, it is not possible to change the contents of a string. The predeclared string type is string.
			// numeric
				// They are again arithmetic types and they represents a) integer types or b) floating point values throughout the program.
			// boolean
				// They are boolean types and consists of the two predefined constants: (a) true (b) false
			// devivad types
				//They include (a) Pointer types, (b) Array types, (c) Structure types, (d) Union types and (e) Function types f) Slice types g) Interface types h) Map types i) Channel Types.


package main

import (
	"fmt"
)

var x int
var y string
var z bool
var k bool = true 

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(k)
	
}
