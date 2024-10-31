package core

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func RunTest(request *Request) (*Result, error) {
	test, err := request.GetRestTest()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	response, err := http.Get(test.Url.String())
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	took := time.Now().Sub(now).Milliseconds()
	bodyMatch, err := testBody(body, request.Body)
	if err != nil {
		return nil, err
	}
	statusMatch := getStatusNumber(response.Status) == test.Status
	result := &Result{Body: bodyMatch, Status: statusMatch, Request: request, Took: took}
	return result, nil
}

func getStatusNumber(rawStatus string) string {
	return strings.Split(rawStatus, " ")[0]
}

func testBody(respBody, expectedBody []byte) (bool, error) {
	var a, b interface{}
	if err := json.Unmarshal(respBody, &a); err != nil {
		return false, err
	}
	if err := json.Unmarshal(expectedBody, &b); err != nil {
		return false, err
	}
	return IsSameJSON(a, b), nil
}

func RunFileTest(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	var requests []*Request
	var data []byte
	buf := make([]byte, 100)
	for {
		n, err := file.Read(buf)
		data = append(data, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	err = json.Unmarshal(data, &requests)
	if err != nil {
		log.Println("unmarshal requests error")
		return err
	}
	for _, request := range requests {
		result, err := RunTest(request)
		if err != nil {
			return err
		}
		Results = append(Results, result)
	}
	return nil
}
