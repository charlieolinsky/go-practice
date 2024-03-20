package tutorial

import (
	"fmt"
	"os"
)

func CliArgs(){
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