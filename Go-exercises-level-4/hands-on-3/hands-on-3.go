//	Using the code from the previous example, use Slicing to create the following new slices
//	[42 43 44 45 46]
//	[47 48 49 50 51]
//	[44 45 46 47 48]
//	[43 44 45 46 47]
//	print them out

//	Using a COMPOSITE LITERAL
//	Create a slice of type int
//	assign 10 values
//	Range over the slice an print the values out.
//	Using format printing
	//	print out the type of the array 
	
	package main

import (
	"fmt"
)

//	Using a COMPOSITE LITERAL
//	Create a slice of type int
//	assign 10 values
//	Range over the slice an print the values out.
//	Using format printing
//	print out the type of the array

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x[:5])
	z := append(x[5:])
	k := append(x[2:7])
	j := append(x[1:6])
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(k)
	fmt.Println(j)

	// or
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}
