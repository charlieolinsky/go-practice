package main

import (
	"log"
	"os"

	"github.com/charlieolinsky/go-practice/chapters/ch1"
)

func main() {
	/* Chapter 1 */

	//1.1: Hello, World
	//ch1.HelloWorld()

	//1.2: Command Line Arguments
	//ch1.CliArgs()

	//1.3: Finding Duplicate Lines
	//ch1.Dup()

	//1.4: Animated Gifs
	file, err := os.Create("gif-out.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = ch1.Gif(file)
	if err != nil {
		log.Fatal(err)
	}
}
