package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler funtion which writes a byte slice containing
// the string as the response body
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it
	// doesn't, use the http.NotFound() function to send a 404 response to the clien
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function.
	//If it can't be converted to an integer, or the value is less than 1, we return a 404 page not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("Display a specific snippet..."))
	fmt.Fprintf(w, "Display a specific snippet with id %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	// Use r.Method to check whether the request is using POST or not.
	//If its not send a method not allowed
	if r.Method != "POST" {
		// response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {

	// Use the http.NewServeMux() function to initialize a new sermux (router), then
	// register the home funtion as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	//Use the http.ListenAndServe() to start a new Webserver. We pass in 2 parameters:
	// the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If the http.ListenAndServe returns an error
	//we use the log.Fatal() to log the error message and exit
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
