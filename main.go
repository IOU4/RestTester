package main

import (
	"flag"
	"log"

	"ou.emad/cli"
	"ou.emad/server"
)

func main() {
	sFlag := flag.String("s", "", "run local server with specified port.")
	urlFlag := flag.String("url", "", "run quick test for specified url.")
	statusFlag := flag.Int("status", 0, "expected status for quick test.")
	bodyFlag := flag.String("body", "", "expected body for quick test.")
	fileFlag := flag.String("file", "", "json file containing test scenarios")
	flag.Parse()
	if port := *sFlag; port != "" {
		log.Println("staring server ...")
		server.Run(port)
	} else {
		log.Println("running rest tester for cli args ...")
		cli.Run(*urlFlag, *statusFlag, *bodyFlag, *fileFlag)
	}
}
