package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/integrii/flaggy"
)

var (
	dir    = "."
	port   = ":8080"
	portRe = regexp.MustCompile(`^:?\d+$`)
)

func main() {
	flaggy.SetDescription("A simple command line tool to just host this")

	flaggy.AddPositionalValue(&dir, "Directory", 1, false, `"This" (The directory to host)`)
	flaggy.String(&port, "p", "port", "Specify a port")
	flaggy.Parse()

	if !portRe.MatchString(port) {
		log.Fatal("Invalid port")
	}
	if port[0] != ':' {
		port = ":" + port
	}

	fs := http.FileServer(http.Dir(dir))
	log.Printf("Hosting %v on http://localhost%v", dir, port)
	log.Fatal(http.ListenAndServe(port, fs))
}
