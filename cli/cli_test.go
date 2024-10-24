package cli

import (
	"testing"
)

func TestCLIRun(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	status := 200
	body := `{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}`
	err := Run(url, status, body, "")
	if err != nil {
		t.Fatal("got errors from cli run: ", err)
	}
}
