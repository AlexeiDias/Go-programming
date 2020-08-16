//	Using a COMPOSITE LITERAL
//	Create a slice of type int
//	assign 10 values
//	Range over the slice an print the values out.
//	Using format printing
	//	print out the type of the array 
	
	func main() {
		x := []int{4, 5, 7, 8, 44, 50, 65 ,67, 55, 88}
		// fmt.Println(len(x))
		// fmt.Println(x[0])
		// fmt.Println(x[1])
		// fmt.Println(x[2])
		// fmt.Println(x[3])
		// fmt.Println(x[4])
	
		for i, v := range x {
			fmt.Println(i, v)
		}
	
	}
	