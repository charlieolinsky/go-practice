package ch4

import "fmt"

func Arrays() {
	//Array of 3 Integers
	var a [3]int
	
	fmt.Println(a[0]) //print first element
	fmt.Printf("%v\n\n", a[len(a)-1]) //print last element

	//Elements in a new array default to the zero value for their type. 
	var a2 [5]int
	for i, v := range a2 {
		fmt.Printf("%d\t%d\n\n", i, v)
	} 

	//We can use an array literal to initialize an array with a list of values
	//var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	//using an ellipsis "..." in place of an array length implies it to be the length of the given iniitalizers
	var q2 = [...]int{1, 4, 6}
	fmt.Println(len(q2))

	//The size of an array is part of its type. [3]int != [4]int
	//the following is an illegal statement
	// z := [3]int{1, 2, 3}
	// z = [4]int{1, 2, 3, 4}

	//If an arrays element type is comparable, then the array type is comparable too
	//Order matters
	c := [...]int{3, 2, 1}
	d := [3]int{1, 2, 3}
	// e := [2]int{3, 2}  Cannot compare different array types  
	fmt.Println(c == d, c != d)


	//Passing large arrays to functions is inefficient and arrays are generally inflexible due to their fixed size. 
	//We could pass a pointer to the array instead, but generally, Go programmers use slices. 
}
