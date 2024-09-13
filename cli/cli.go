package cli

import (
	"fmt"
	"log"
	"reflect"

	"ou.emad/core"
)

func Run(args []string) error {
	if len(args) != 4 {
		printUsage()
		return fmt.Errorf("invalid argument number: %d\n", len(args))
	}
	result, err := core.RunTest(GetRestTestRequest(args))
	if err != nil {
		return err
	}
	printResult(*result)
	return nil
}

func GetRestTestRequest(args []string) core.RestTestRequest {
	return core.RestTestRequest{Url: args[1], Status: args[2], Body: args[3]}
}

func printUsage() {
	fmt.Println(`usage: rest-tester <url> <expectedStatus> <expectedBody>`)
}

func printResult(result core.RestTestResult) {
	log.Println("test result:")
	v := reflect.ValueOf(result)
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fmt.Printf("%v: %v\n", fieldName, v.Field(i))
	}
}
