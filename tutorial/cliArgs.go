package tutorial

import (
	"fmt"
	"os"
	"strings"
)

func CliArgs(){

	// os.Args is a slice of strings
	// os.Args[0] = Name of the command itself
	// os.Args[1:len(os.Args)] = list of command line arguments
	// os.Args[1:] is a shorthand for the above statement
	
	//Note: in s[m:n] if m is omitted it defaults to 0
	// if n is omitted it defaults to len(s)

	fmt.Println("\nCommand Line Arguments: ")
	argZ := os.Args[0]
	fmt.Println("Args[0]   " + argZ)
	arg := os.Args[1:]
	fmt.Println("Args[1:] ", arg)


	fmt.Println("\nEcho Clones")
	var s, sep string
	for i := 1 ; i<len(os.Args) ; i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("ECHO#1: %s", s)

	s2, sep2 := "", ""
	for _,arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Printf("\nECHO#2: %s", s2)

	fmt.Printf("\nECHO#3: ")
	fmt.Println(strings.Join(os.Args[1:], " "))	
	//strings.Join() concats the elements of first arg to create a single string. The second param is a seperator placed 
	//between those elements.
}