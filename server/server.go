package server

import (
	"log"
	"net/http"
)

func Serve() {
	router()
	log.Fatal(http.ListenAndServe(":8088", nil))
}
