// Package main is the entry point into our application.
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":8000", "HTTP Network Address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		errorLog: errLog,
		infoLog:  infoLog,
	}

	srv := http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: app.errorLog,
	}

	app.infoLog.Printf("Starting server on port %s", *addr)
	err := srv.ListenAndServe()
	errLog.Fatal(err)

}
