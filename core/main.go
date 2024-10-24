package core

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func RunTest(request Request) (*Result, error) {
	log.Println("running test for url: ", request.Url)
	test, err := request.GetRestTest()
	if err != nil {
		return nil, err
	}
	took := time.Now()
	response, err := http.Get(test.Url.String())
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
	bodyMatch := testBody(body, test.Body)
	statusMatch := getStatusNumber(response.Status) == test.Status
	result := &Result{BodyMatch: bodyMatch, StatusMatch: statusMatch}
	return result, nil
}

func getStatusNumber(rawStatus string) string {
	return strings.Split(rawStatus, " ")[0]
}

func testBody(respBody []byte, expectedBody string) bool {
	return strings.TrimSpace(string(respBody)) == strings.TrimSpace(expectedBody)
}
