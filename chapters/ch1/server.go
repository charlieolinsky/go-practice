package ch1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func Server() {
	//Server1()
	//Server2()
	//Server3()
	GifServer()
}

// A minimal Go server
func Server1() {
	http.HandleFunc("/", handler1) //each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// echos the path component of the request URL, r
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Server2 is a minimal "echo" and counter server
var mu sync.Mutex
var count int

func Server2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//Echos the path component of the requested URL
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// echos the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// Server 3 echos the HTTP Request Information
func Server3() {
	http.HandleFunc("/", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	
	//Shorthand error handling
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func GifServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Gif(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}