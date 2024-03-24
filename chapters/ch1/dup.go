package ch1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Dup() {
	fmt.Println("\nFinding Duplicate Lines")
	//Dup1()
	Dup2()
	//Dup3()
}

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
//Note: to send EOF signal to scanner press Ctrl+Z, Enter (in windows)
func Dup1() {
	counts := make(map[string]int)      //make() creates a new empty map
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

// This version can read from the standard input or from a seqeunce of named files
// Note, there is a file called text.txt in /test to use
func Dup2() {
	counts := make(map[string]int)
	//Takes file name from Args
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file) // os.Open() returns both an open file (*os.File), the second is an error

			//if something went wrong display the error
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERR_DUP2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close() //closes the file and releases any resources
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//NOTE: ignoring potential errors from input.Err()
}

//Here we read the entire input into memory in one shot, split it into lines, and the process.
func Dup3() {
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := os.ReadFile(fileName)
		if err != nil{
			fmt.Fprintf(os.Stderr, "ERRDUP3: %v\n", err) 
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}