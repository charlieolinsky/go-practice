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
	fmt.Println("Q2: ", Q2)
	fmt.Println("Summer: ", summer)

	// the len() function returns the number of elements in a given slice
	fmt.Println("Summer Len: ", len(summer))
	// the cap() function returns the distance between the start of the slice and the end of the underlying array.
	fmt.Println("Summer Cap: ", cap(summer))

	//Slicing beyond cap(s) causes a panic
	//fmt.Println(summer[:20])

	//but slicing beyond len(s) extends the slice.
	//Therefore, the result may be longer than the original
	endlessSummer := summer[:7] //extend a slice within capacity
	fmt.Println("Endless Summer: ", endlessSummer)

	//Passing a slice to a function permits the function to modify the underlying array elements
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println("Reverse: ", a) // [5, 4, 3, 2, 1]

	//Rotate the array left by two positions
	q := []int{0, 1, 2, 3, 4, 5}

	fmt.Println("Rotate (OG): ", q)
	reverse(q[:2])
	fmt.Println("Rotate (Step 1): ", q)
	reverse(q[2:])
	fmt.Println("Rotate (Step 2): ", q)
	reverse(q) //do this first to rotate right
	fmt.Println("Rotate (FINAL): ", q)

	/* Slices are not comparable using == */
	// The only legal slice comparison is against nil
}

//Cool for loop!
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
