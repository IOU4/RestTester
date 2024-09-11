package cli

import (
	"fmt"
	"ou.emad/core"
)

func Run(args []string) error {
	if len(args) != 4 {
		printUsage()
		return fmt.Errorf("invalid argument number: %d\n", len(args))
	}
	core.RunTest(GetRestTestRequest(args))
	return nil
}

func GetRestTestRequest(args []string) core.RestTestRequest {
	return core.RestTestRequest{Url: args[1], Status: args[2], Body: args[3]}
}

func printUsage() {
	fmt.Println(`usage: rest-tester <url> <expectedStatus> <expectedBody>`)
}
