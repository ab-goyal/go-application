package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var message string
var kvFile string
var kvMap = make(map[string]string)

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
	printKVMap(w)
}

func printKVMap(w http.ResponseWriter) {
	if len(kvMap) == 0 {
		return
	}
	fmt.Fprintln(w, "Key-Value Pairs:")
	for k, v := range kvMap {
		fmt.Fprintf(w, "%s=%s\n", k, v)
	}
}
func main() {

	flag.StringVar(&message, "message", "Hello this is the default message", "message to be printed on the / and /hello endpoints")
	flag.StringVar(&kvFile, "kvfile", "", "Path to file containing key-value pairs")

	// Load key-value pairs from file
	if kvFile != "" {
		file, err := os.Open(kvFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open key-value file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			pair := strings.Split(line, "=")
			if len(pair) != 2 {
				fmt.Fprintf(os.Stderr, "Invalid key-value pair: %v\n", line)
				os.Exit(1)
			}
			kvMap[pair[0]] = pair[1]
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read key-value file: %v\n", err)
			os.Exit(1)
		}
	}

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/mytime", timeHandler)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Read the name parameter from the query string
		name := r.URL.Query().Get("name")

		// Call the helloHandler with the name parameter
		helloHandler(w, r, name)
		printKVMap(w)
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
