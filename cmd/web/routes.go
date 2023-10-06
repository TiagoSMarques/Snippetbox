package main

import "net/http"

func (app *application) routes(staticDir string) *http.ServeMux {
	// Use the http.NewServeMux() function to initialize a new sermux (router), then
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.

	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	fileServer := http.FileServer(http.Dir(staticDir))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// register the home funtion as the handler for the "/" URL pattern
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
