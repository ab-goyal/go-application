package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var message string

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	fmt.Fprintf(w, "\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request, name string) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	if name == "" {
		// If the name parameter is not present in the query string, show a generic greeting
		fmt.Fprint(w, "Hello, guest!\n")
	} else {
		// If the name parameter is present in the query string, show a personalized greeting
		fmt.Fprintf(w, "Hello, %s!\n", name)
	}
	fmt.Fprintf(w, message)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/mytime" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	t := time.Now()

	//t.Format("01-02-2006 15:04:05 Monday")
	fmt.Fprintf(w, "The current time is: %s\n", t.Format("01-02-2006 15:04:05 Monday"))
	fmt.Fprintf(w, message)
}
func main() {

	flag.StringVar(&message, "message", "Hello this is the default message", "message to be printed on the / and /hello endpoints")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/mytime", timeHandler)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Read the name parameter from the query string
		name := r.URL.Query().Get("name")

		// Call the helloHandler with the name parameter
		helloHandler(w, r, name)
	})

	var port = "3000"
	var host = "localhost"
	flag.StringVar(&port, "port", port, "Port number")
	flag.Parse()
	//fmt.Println("You seem to prefer", port)

	fmt.Printf("Starting server at port%v\n", port)
	if err := http.ListenAndServe(host+":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
