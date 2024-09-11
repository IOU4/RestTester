package handlers

import (
	"encoding/json"
	"net/http"
)

type HelloHandler struct{}

func (h1 *HelloHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	body, _ := json.Marshal("hello from rest tester")
	rw.Write(body)
}
