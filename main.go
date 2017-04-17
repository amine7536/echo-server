package main

import (
	"fmt"
	"io/ioutil"
	"os"

	common "github.com/amine7536/echo-server/common"
	"github.com/amine7536/echo-server/tcpserver"
	docopt "github.com/docopt/docopt-go"
)

const (
	version     = "0.1.0"
	programName = "Echo Server"
)

func main() {
	common.InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	usage := fmt.Sprintf(`%s

Usage:
  echo-server [--host=HOST] [--port=PORT]
  echo-server -h | --help
  echo-server --version

Options:
  --host=HOST   destination host [default: localhost]
  --port=PORT   destination port [default: 3000]
  -h --help         show this help message and exit
  --version        show version and exit`, programName)

	versionStr := fmt.Sprintf("%s %s", programName, version)

	arguments, _ := docopt.Parse(usage, nil, true, versionStr, false)

	host := arguments["--host"].(string)
	port := arguments["--port"].(string)
	service := host + ":" + port

	tcpserver.Start(service, handleConnection)
}

func handleConnection(c3r tcpserver.C3Request) {
	// c3r.Read()
	c3r.Write()
}

func checkerror(err error, level string) {
	if err != nil {
		common.Log(level, fmt.Sprintf("err=%s", err))
	}
}
