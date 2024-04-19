package core

import "net/url"

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
	statusMatch bool
	bodyMatch   bool
}
