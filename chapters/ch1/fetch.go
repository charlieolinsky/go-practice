package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Fetch() {
	for _, url := range os.Args[1:] {

		//Ensure there is a proper prefix (Exercise 1.8) 
		if !strings.HasPrefix(url, "http://") {
			fmt.Println("Old URL: " + url)
			url = "http://" + url
			fmt.Println("New URL: "+url+"\n")
		}

		resp, err := http.Get(url) //Makes an HTTP Request. Returns a struct resp, or an error.
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
 		
		//resp.Body contains the server repsonse as a readable stream
		//io.ReadAll() reads the entire response and stores the result.
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		//Write to os.Stdout, print HTTP response status (Exercise 1.9)
		fmt.Printf("%s\n%s", b, resp.Status)
	}
}
