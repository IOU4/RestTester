package core

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func RunTest(request RestTestRequest) (*RestTestResult, error) {
	log.Println("running test for: {}", request)
	testObj, err := request.GetRestTest()
	if err != nil {
		return nil, err
	}
	took := time.Now()
	response, err := http.Get(testObj.Url.String())
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("after %vms, got reponse status: %v and response content: %v\n",
		time.Now().Sub(took).Milliseconds(),
		response.Status, response.Header.Get("Content-Type"))
	bodyMatch := string(body) == testObj.Body
	statusMatch := getStatusNumber(response.Status) == testObj.Status
	result := &RestTestResult{BodyMatch: bodyMatch, StatusMatch: statusMatch}
	return result, nil
}

func getStatusNumber(rawStatus string) string {
	return strings.Split(rawStatus, " ")[0]
}
