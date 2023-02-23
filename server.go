package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

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

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Println("Enter the name:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Fprintf(w, "Hello %s", text)
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
	fmt.Fprintf(w, "The current time is: %s", t.Format("01-02-2006 15:04:05 Monday"))
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/mytime", timeHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
