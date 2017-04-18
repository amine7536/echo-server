package tcpserver

import (
	"bufio"
	"errors"
	"fmt"
	"net"

	"github.com/amine7536/echo-server/common"
)

type handlerFunc func(EchoRequest)

// EchoRequest incoming EchoRequest
type EchoRequest struct {
	Proto      string
	ProtoMajor int
	ProtMinor  int
	Body       []byte
	RemoteAddr net.Addr
	Conn       net.Conn
}

func (req *EchoRequest) Write() error {
	defer req.Conn.Close()

	if req.Conn == nil {
		return errors.New("Error no connection")
	}

	req.Conn.Write(req.Body)

	return nil
}

func (req *EchoRequest) Read() error {
	bufr := bufio.NewReader(req.Conn)
	buf := make([]byte, 1024)

	readBytes, err := bufr.Read(buf)
	if err != nil {
		req.Conn.Close()
		return err
	}
	req.Body = buf[:readBytes]

	return nil
}

// Start the server
func Start(service string, fn handlerFunc) error {
	server, err := net.Listen("tcp", service)
	if err != nil {
		return err
	}
	common.Log("INFO", fmt.Sprintf("Server started listening on %s", service))

mainloop:
	for {
		conn, err := server.Accept()
		if err != nil {
			break mainloop
		}

		EchoRequest := EchoRequest{
			Proto:      "Echo",
			ProtoMajor: 1,
			ProtMinor:  2,
			Conn:       conn,
			RemoteAddr: conn.RemoteAddr(),
		}

		go fn(EchoRequest)
	}
	return nil
}
