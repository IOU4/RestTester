package main

import (
	"log"
	"os"

	"ou.emad/cli"
	"ou.emad/server"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Println("staring server ...")
		server.Run()
	} else {
		log.Println("running rest tester for cli args ...")
		cli.Run(args)
	}
}
