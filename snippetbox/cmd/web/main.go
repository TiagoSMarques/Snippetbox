package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

// Define an application struct to hold the application-wide dependencies for the web application
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	var cfg config
	// Define a new command-line flag with the name 'addr', a default value of ":4000"

	// addr := flag.String("addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	//if we want to have a config struct we can use flag.StringVar to store the flag in an existing variable

	flag.Parse()

	// Use log.New() to create a logger for writing information messages This takes
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// For example, we could redirect the stdout and stderr streams to on-disk files when starting the application like so:
	// $go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log

	// Initialize a new instance of our application struct, containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(cfg.staticDir),
	}

	//Use the http.ListenAndServe() to start a new Webserver. We pass in 2 parameters:
	// the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If the http.ListenAndServe returns an error
	//we use the log.Fatal() to log the error message and exit
	infoLog.Print("Starting server on", cfg.addr)
	// err := http.ListenAndServe(cfg.addr, mux)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
