package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new sermux (router), then
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// register the home funtion as the handler for the "/" URL pattern
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
