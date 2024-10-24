package core

import (
	"testing"
)

func TestRunTest(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	status := "200"
	body := `{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}`
	resutl, err := RunTest(Request{Url: url, Status: status, Body: body})
	if err != nil {
		t.Fatal(err)
	}
	if !resutl.BodyMatch || !resutl.StatusMatch {
		t.FailNow()
	}
}

func TestBadUrlRunTest(t *testing.T) {
	url := "something_bad"
	status := "201"
	body := "alkdsj"
	_, err := RunTest(Request{Url: url, Status: status, Body: body})
	if err == nil {
		t.Fatal("should fail on invalid url")
	}
}
