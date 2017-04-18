package main

import (
	"log"

	"github.com/amine7536/echo-server/cmd"
)

const (
	// Version : app version
	Version = "0.3.0"
	// ProgramName : app name
	ProgramName = "echo-server"
)

func main() {

	if err := cmd.NewRootCmd(Version, ProgramName).Execute(); err != nil {
		log.Fatal(err)
	}
}
