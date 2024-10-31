package cli

import (
	"log"
	"strconv"

	"ou.emad/core"
)

func Run(url string, status int, body, file string) {
	if file != "" {
		log.Println("running tests from file", file, "...")
		err := core.RunFileTest(file)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range core.Results {
			v.Print()
		}
		return
	}
	log.Println("running test for cli args...")
	result, err := core.RunTest(&core.Request{Url: url, Status: strconv.Itoa(status), Body: []byte(body)})
	if err != nil {
		log.Fatal(err)
	}
	result.Print()
}
