package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"

	"github.com/TiagoSMarques/Snippetbox/internal/models"
)

type config struct {
	addr      string
	staticDir string
	dsn       string
}

// Define an application struct to hold the application-wide dependencies for the web application
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {

	var cfg config
	// Define a new command-line flag with the name 'addr', a default value of ":4000"

	// addr := flag.String("addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.dsn, "dsn", "web:123@/snippetbox?parseTime=true", "MySQL data source name")

	//if we want to have a config struct we can use flag.StringVar to store the flag in an existing variable

	flag.Parse()

	// Use log.New() to create a logger for writing information messages This takes
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// For example, we could redirect the stdout and stderr streams to on-disk files when starting the application like so:
	// $go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log

	//The sql.Open() function initializes a new sql.DB object, which is essentially a pool of database connections go manages connections in this pool as needed, automatically opening and closing them
	//parseTime=true is a DSN parameter that instructs our driver to converte SQL TIME to go time.time
	db, err := openDB(cfg.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	//make sure the connection pool closes before the main func exits
	defer db.Close()

	// Initialize a new instance of our application struct, containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(cfg.staticDir),
	}

	//Use the http.ListenAndServe() to start a new Webserver. We pass in 1 parameters:
	// the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If the http.ListenAndServe returns an error
	//we use the log.Fatal() to log the error message and exit
	infoLog.Print("Starting server on", cfg.addr)
	// err := http.ListenAndServe(cfg.addr, mux)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
