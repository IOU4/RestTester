package core

import (
	"io"
	"log"
	"net/http"
)

func Proceed(test RestTestRequest) {
	ss := GetRestTest(test)
	endpoints := []string{ss.Url.String()}
	response, err := http.Get(endpoints[0])
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	log.Println("data:", string(data))
}
