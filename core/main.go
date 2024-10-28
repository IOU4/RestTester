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
	log.Printf("test done after %vms.\n", time.Now().Sub(took).Milliseconds())
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

func RunFileTest(path string) ([]*Result, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var requests []Request
	var data []byte
	buf := make([]byte, 100)
	for {
		n, err := file.Read(buf)
		data = append(data, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return []*Result{}, err
		}
	}
	err = json.Unmarshal(data, &requests)
	if err != nil {
		return []*Result{}, err
	}
	var results []*Result
	for _, request := range requests {
		result, err := RunTest(request)
		if err != nil {
			return []*Result{}, err
		}
		results = append(results, result)
	}
	return results, nil
}
