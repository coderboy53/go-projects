package main

import (
	"fmt" //fmt package for terminal and web page printing
	"log" //for logging out errors
	"net/http" // for all http server operations
	"os" //for fetching the command line args
	"strconv"
)


// basically prints out hello when user goes to the /hello route
//if route is not /hello bu function has been invoked return 404
//if route request method is not GET return 404
//otherwise print hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")

}

// fetches data from form.html and displays it at /form route
//if there is some error parsing form data, print the error to web page
//otherwise print post request was successful
//and display details entered in form
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful!\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
}

//get the port to start server on from command line
// create the fileServer handler to redirect requests to / route to open ./static directory containing index.html
// create http handler functions for each routes
// print the start server message to terminal
// listen on the provided port and in case of error log out the error
func main() {
	port, _ := strconv.Atoi(os.Args[1])
	filesServer := http.FileServer(http.Dir("./static")) //tells golang to check static directory for HTTP web pages
	http.Handle("/",filesServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello",helloHandler)
	
	fmt.Printf("Starting server at port %v", port)
	addr := ":"+strconv.Itoa(port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}


}