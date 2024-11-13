package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"ou.emad/core"
)

func Run(url string, status int, body string) {
	log.Println("running test for cli args...")
	result := core.RunTest(&core.Request{Url: url, Status: strconv.Itoa(status), Body: []byte(body)})
	if result.Error() != nil {
		log.Fatal(result.Error())
	}
	result.Print()
}
func RunFromFile(path string) error {
	log.Println("running tests from file", path, "...")
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	var requests []*core.Request
	var data []byte
	buf := make([]byte, 100)
	for {
		n, err := file.Read(buf)
		data = append(data, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	err = json.Unmarshal(data, &requests)
	if err != nil {
		log.Println("unmarshal requests error")
		return err
	}
	ch := make(chan *core.Result)
	go core.RunMultipleTests(requests, ch)
	for r := range ch {
		r.PrintUrl()
		if r.Error() != nil {
			fmt.Println()
			fmt.Println("error:", r.Error())
		} else {
			r.Print()
		}
	}
	return err
}
