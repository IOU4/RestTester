package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"ou.emad/core"
)

type TestHandler struct{}

func (h2 *TestHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("eror reading body")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()
	var testRequest core.RestTestRequest
	if err := json.Unmarshal(reqBody, &testRequest); err != nil {
		log.Println("coding decode body")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := core.RunTest(testRequest)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(result)
	rw.Write(body)
}
