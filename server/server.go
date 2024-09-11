package server

import (
	"log"
	"net/http"

	"ou.emad/server/handlers"
)

func Run() {
	http.Handle("/", new(handlers.HelloHandler))
	http.Handle("/test", new(handlers.TestHandler))

	port := "8080"
	log.Printf("running on port %s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
