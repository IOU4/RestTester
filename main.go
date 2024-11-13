package main

import (
	"flag"

	"ou.emad/cli"
	"ou.emad/server"
)

func main() {
	serverFlag := flag.String("server", "", "run local server with specified port.")
	urlFlag := flag.String("url", "", "run quick test for specified url.")
	statusFlag := flag.Int("status", 0, "expected status for quick test.")
	bodyFlag := flag.String("body", "", "expected body for quick test.")
	fileFlag := flag.String("file", "", "json file containing test scenarios")
	flag.Parse()
	switch {
	case *serverFlag != "":
		server.Run(*serverFlag)
		break
	case *fileFlag != "":
		cli.RunFromFile(*fileFlag)
		break
	case *urlFlag != "" && *statusFlag != 0:
		cli.Run(*urlFlag, *statusFlag, *bodyFlag)
		break
	default:
		flag.Usage()
	}
}
