package tutorial

import (
	"fmt"
	"os"
)

func CliArgs(){

	// os.Args is a slice of strings
	// os.Args[0] = Name of the command itself
	// os.Args[1:len(os.Args)] = list of command line arguments
	// os.Args[1:] is a shorthand for the above statement
	
	//Note: in s[m:n] if m is omitted it defaults to 0
	// if n is omitted it defaults to len(s)

	fmt.Println("\nCommand Line Arguments: ")
	arg := os.Args[1:]
	fmt.Println(arg)


	fmt.Println("\nEcho Clones")
	var s, sep string
	for i := 1 ; i<len(os.Args) ; i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("ECHO#1: %s", s)

}