package core

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

func RunTest(request *Request) *Result {
	test, err := request.GetRestTest()
	if err != nil {
		return &Result{err: err}
	}
	now := time.Now()
	response, err := http.Get(test.Url.String())
	if err != nil {
		return &Result{err: err}
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &Result{err: err}
	}
	took := time.Now().Sub(now).Milliseconds()
	bodyMatch, err := testBody(body, request.Body)
	if err != nil {
		return &Result{err: err}
	}
	statusMatch := getStatusNumber(response.Status) == test.Status
	result := &Result{Body: bodyMatch, Status: statusMatch, request: request, Took: took}
	return result
}

func getStatusNumber(rawStatus string) string {
	return strings.Split(rawStatus, " ")[0]
}

func testBody(respBody, expectedBody []byte) (bool, error) {
	var a, b interface{}
	if err := json.Unmarshal(respBody, &a); err != nil {
		return false, errors.New("couldn't parse response body as json")
	}
	if err := json.Unmarshal(expectedBody, &b); err != nil {
		return false, err
	}
	return IsSameJSON(a, b), nil
}

func RunMultipleTests(requests []*Request, ch chan *Result) {
	defer close(ch)
	for _, request := range requests {
		result := RunTest(request)
		if result.err != nil {
			result.request = request
		}
		ch <- result
	}
}
