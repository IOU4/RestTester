package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
)

var Results []*Result
var cursor int

type Request struct {
	Url    string          `json:"url"`
	Status string          `json:"status"`
	Body   json.RawMessage `json:"body"`
}

type Test struct {
	Url    *url.URL
	Status string
	Body   interface{}
}

type Result struct {
	Request *Request
	Took    int64
	Status  bool `json:"statusMatch"`
	Body    bool `json:"bodyMatch"`
}

func getUrl(rawUrl string) (*url.URL, error) {
	urlPattern := regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+(:[0-9]+)?|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w\-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)`)
	if !urlPattern.MatchString(rawUrl) {
		return nil, errors.New("invalid url: " + rawUrl)
	}
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	return parsedUrl, nil
}

func getStatus(rawStatus string) string {
	statusPattern := regexp.MustCompile(`[12345]{\d}{2}`)
	if statusPattern.MatchString(rawStatus) {
		errors.New("invalid status")
	}
	return rawStatus
}

func (test *Request) GetRestTest() (*Test, error) {
	url, err := getUrl(test.Url)
	if err != nil {
		return nil, err
	}
	return &Test{Url: url, Status: getStatus(test.Status), Body: test.Body}, nil
}

func (result *Result) Print() {
	fmt.Println("[+]", result.Request.Url, ":", result.Took, "ms")
	fmt.Println("   matched_status:", result.Status)
	fmt.Println("   matched_body:", result.Body)
	fmt.Println()
}

func (result *Result) HasNext() bool {
	return cursor > len(Results)-1 || cursor < 0
}

func (result *Result) Next() *Result {
	res := Results[cursor]
	cursor++
	return res
}
