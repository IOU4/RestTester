package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Proceed(request RestTestRequest) {
	testObj := request.GetRestTest()
	endpoints := []string{testObj.Url.String()}
	response, err := http.Get(endpoints[0])
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	bodyMatch := string(body) == testObj.Body
	statusMatch := getStatusNumber(response.Status) == testObj.Status
	result := RestTestResult{BodyMatch: bodyMatch, StatusMatch: statusMatch}
	printResult(result)
}

func getStatusNumber(rawStatus string) string {
	return strings.Split(rawStatus, " ")[0]
}

func printResult(result RestTestResult) {
	json, err := json.Marshal(result)
	if err != nil {
		panic("error encoding json")
	}
	fmt.Println(string(json))
}
