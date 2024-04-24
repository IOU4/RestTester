package core

import (
	"net/url"
	"regexp"
)

type RestTestRequest struct {
	Url    string
	Status string
	Body   string
}

type RestTest struct {
	Url    *url.URL
	Status string
	Body   string
}

type RestTestResult struct {
	StatusMatch bool `json:"satausMatch"`
	BodyMatch   bool `json:"bodyMatch"`
}

func getUrl(rawUrl string) *url.URL {
	urlPattern := regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+(:[0-9]+)?|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w\-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)`)
	if !urlPattern.MatchString(rawUrl) {
		panic("invalid url: " + rawUrl)
	}

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		panic("couldn't parse url")
	}
	return parsedUrl
}

func getStatus(rawStatus string) string {
	statusPattern := regexp.MustCompile(`[12345]{\d}{2}`)
	if statusPattern.MatchString(rawStatus) {
		panic("invalid status")
	}
	return rawStatus
}

func (test *RestTestRequest) GetRestTest() RestTest {
	return RestTest{Url: getUrl(test.Url), Status: getStatus(test.Status), Body: test.Body}
}
