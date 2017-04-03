package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	common "github.com/amine7536/echo-server/common"
	docopt "github.com/docopt/docopt-go"
	uuid "github.com/satori/go.uuid"
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

	server, err := net.Listen("tcp", service)
	checkerror(err, "ERROR")
	common.Log("INFO", fmt.Sprintf("Server started listening on %s", service))

	for {
		conn, err := server.Accept()
		checkerror(err, "ERROR")
		connUUID := uuid.NewV4()
		go handleConnection(conn, connUUID)
	}
}

func handleConnection(conn net.Conn, connUUID uuid.UUID) {
	bufr := bufio.NewReader(conn)
	buf := make([]byte, 1024)

	common.Log("INFO", fmt.Sprintf("Accepted new connection from %s with connUUID: %s", conn.RemoteAddr(), connUUID))

	for {
		readBytes, err := bufr.Read(buf)
		if err != nil {
			common.Log("WARNING", fmt.Sprintf("connUUID=%s, client=%s, err=%s", connUUID, err, conn.RemoteAddr()))
			conn.Close()
			return
		}

		bufStr := string(buf[:readBytes])
		bufHex := hex.Dump(buf[:readBytes])
		echoStr := fmt.Sprintf("echoserver >> %s\n", bufStr)

		common.Log("INFO", fmt.Sprintf("connUUID: %s >> echoserver\n%s", connUUID, bufHex))
		conn.Write([]byte(echoStr))
		common.Log("INFO", fmt.Sprintf("echoserver >> connUUID: %s\n%s", connUUID, bufHex))

	}
}

func checkerror(err error, level string) {
	if err != nil {
		common.Log(level, fmt.Sprintf("err=%s", err))
	}
}
