package server

import (
	"encoding/json"
	"net/http"
)

func router() {
	http.Handle("/", new(helloHandler))
	http.Handle("/test", new(testHandler))
}

type helloHandler struct{}

func (h1 *helloHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	body, _ := json.Marshal("hello from rest tester")
	rw.Write(body)
}

type testHandler struct{}

func (h2 *testHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	body, _ := json.Marshal("your a post")
	rw.Write(body)
}
