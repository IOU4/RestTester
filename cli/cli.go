package cli

import (
	"strconv"

	"ou.emad/core"
)

func Run(url string, status int, body, file string) error {
	result, err := core.RunTest(core.Request{Url: url, Status: strconv.Itoa(status), Body: body})
	if err != nil {
		return err
	}
	result.Print()
	return nil
}
