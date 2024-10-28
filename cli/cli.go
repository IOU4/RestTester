package cli

import (
	"log"
	"strconv"

	"ou.emad/core"
)

func Run(url string, status int, body, file string) {
	if file != "" {
		results, err := core.RunFileTest(file)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range results {
			v.Print()
		}
		return
	}
	result, err := core.RunTest(core.Request{Url: url, Status: strconv.Itoa(status), Body: body})
	if err != nil {
		log.Fatal(err)
	}
	result.Print()
}
