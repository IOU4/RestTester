package core

import (
	"net/url"
	"regexp"
)

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

func GetRestTest(test RestTestRequest) RestTest {
	return RestTest{Url: getUrl(test.Url), Status: getStatus(test.Status), Body: test.Body}
}
