package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	port := flag.String("p", "80", "server port")
	directory := flag.String("d", ".", "root directory")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on http://localhost:%s/\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}