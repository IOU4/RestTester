package core

import (
	"errors"
	"net/url"
	"regexp"
)

type RestTestRequest struct {
	Url    string `json:"url"`
	Status string `json:"status"`
	Body   string `json:"body"`
}

type RestTest struct {
	Url    *url.URL
	Status string
	Body   string
}

type RestTestResult struct {
	StatusMatch bool `json:"statusMatch"`
	BodyMatch   bool `json:"bodyMatch"`
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

func (test *RestTestRequest) GetRestTest() (*RestTest, error) {
	url, err := getUrl(test.Url)
	if err != nil {
		return nil, err
	}
	return &RestTest{Url: url, Status: getStatus(test.Status), Body: test.Body}, nil
}

type RestTestOutput interface {
	Output()
}
