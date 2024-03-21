package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Dup() {
	fmt.Println("\nFinding Duplicate Lines")
	Dup1()
}

// Dup1 prints the text of each line that appears more than 
// once in the standard input, preceded by its count.
//Note: to send EOF signal to scanner press Ctrl+Z, Enter (in windows)
func Dup1() {
	counts := make(map[string]int)  //make() creates a new empty map
	input := bufio.NewScanner(os.Stdin) //Scanner breaks input into lines

	// Scan() method returns true or false if there is another line
	for input.Scan() {
		//Text() provides the value of the current line
		counts[input.Text()]++
	}
	//NOTE: ignoring potentenial errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	fmt.Println("Counts Map: ", counts)
}