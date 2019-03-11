package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port int
)

func main() {
	flag.IntVar(&port, "port", 8000, "port where the mock server should be listening on")
	flag.Parse()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "OK")
	})

	log.Printf("Starting mock server at port %d...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
