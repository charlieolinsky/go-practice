package ch4

import "fmt"

func Slices() {

	/*
			Slices are variable length sequences whose elements are all of the same type.
		 	A slice is written in the form []T (it looks like an array type without a size).

		 	Arrays and slices are intimately connected. A slice is a lightweight data structure that gives
		 	access to a subsequence (or perhaps all) of the elements of an array, which is known as the
		 	slice's underlying array.

		     A slice has 3 components: a pointer, a length, and a capacity.

			 Multiple slices can share the same underlying array and may refer to overlapping parts of that array
	*/

	//Underlying Array Example
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	//The slice operator s[i:j], where 0 <= i <= j <= cap(s) creates a new slice that refers to elements i through j-1 of the sequence s.
	//The resulting slice has j-i elements
	s := months[3:9]
	fmt.Println(s)

	//If i is omitted, it is 0. If j is omitted it is len(s).
	//s := months[1:13] //refers to whole range of valid months
	//s := months[1:]   // as does this
	//s := months[:]    // refers to the whole array

	/* Lets define an overlapping slice. Both contain the month June */
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

}
