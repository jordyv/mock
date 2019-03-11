package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

var (
	port int
)

func main() {
	flag.IntVar(&port, "port", 8000, "port where the mock server should be listening on")
	flag.Parse()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.
			WithFields(log.Fields{"IP": request.RemoteAddr, "userAgent": request.UserAgent()}).
			Infof("Got %s request for '%s'", request.Method, request.RequestURI)
		if _, err := fmt.Fprintln(writer, "OK"); err != nil {
			log.WithError(err).Error("could not write client response")
		}
	})

	log.Infof("Starting mock server at port %d...", port)
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
