package tcpserver

import (
	"bufio"
	"errors"
	"fmt"
	"net"

	"github.com/amine7536/echo-server/common"
)

type handlerFunc func(C3Request)

// C3Request incoming C3Request
type C3Request struct {
	Proto      string
	ProtoMajor int
	ProtMinor  int
	GetBody    func() ([]byte, error)
	RemoteAddr net.Addr
	Conn       net.Conn
}

func (c3r *C3Request) Write() error {
	if c3r.Conn == nil {
		return errors.New("Error no connection")
	}

	c3r.Conn.Write([]byte("Hello"))

	return nil
}

func (c3r *C3Request) Read() error {
	bufr := bufio.NewReader(c3r.Conn)
	return nil
}

// A Handler responds to an C3 request.
//
// ServeC3 should write reply data to the ResponseWriter
// and then return.
type Handler interface {
	ServeC3(C3ResponseWriter, *C3Request)
}

// TCPResponseWriter interface
type C3ResponseWriter interface {
	Write([]byte) error
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

		c3request := C3Request{
			Proto:      "C3",
			ProtoMajor: 1,
			ProtMinor:  2,
			Conn:       conn,
			RemoteAddr: conn.RemoteAddr(),
		}

		go fn(c3request)
	}
	return nil
}
