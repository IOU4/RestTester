package cli

import (
	"fmt"
	"os"

	"ou.emad/core"
)

func GetRestTestRequest() core.RestTestRequest {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("invalid argument number:", len(args))
		printUsage()
		panic(1)
	}
	return core.RestTestRequest{Url: args[1], Status: args[2], Body: args[3]}
}

func printUsage() {
	fmt.Println(` usage: rest-tester <url> <expectedStatus> <expectedBody> `)
}
