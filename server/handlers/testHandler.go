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
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()
	var testRequest core.Request
	if err := json.Unmarshal(reqBody, &testRequest); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	result := core.RunTest(&testRequest)
	if result.Error() != nil {
		log.Println(result.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	body, _ := json.Marshal(result)
	rw.Write(body)
}
