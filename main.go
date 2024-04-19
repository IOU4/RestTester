package main

import (
	"ou.emad/cli"
	"ou.emad/core"
)

func main() {
	core.Proceed(cli.GetRestTestRequest())
}
