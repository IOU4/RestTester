package server

import (
	"log"
	"net/http"

	"ou.emad/server/handlers"
)

func Run(port string) {
	http.Handle("/", new(handlers.HelloHandler))
	http.Handle("/test", new(handlers.TestHandler))

	log.Printf("running on port %s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
