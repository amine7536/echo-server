package server

import (
	"bufio"
	"errors"
	"io"
	"net"

	"github.com/Sirupsen/logrus"
)

type handlerFunc func(EchoRequest)

// EchoRequest incoming EchoRequest
type EchoRequest struct {
	Proto      string
	ProtoMajor int
	ProtMinor  int
	Body       []byte
	RemoteAddr net.Addr
	Conn       io.ReadWriteCloser
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

// EchoServer server struct
type EchoServer struct {
	conn   net.Listener
	logger *logrus.Entry
}

// New initialize a new EchoServer
func New(conn net.Listener, logger *logrus.Entry) *EchoServer {
	s := &EchoServer{
		conn:   conn,
		logger: logger,
	}

	return s
}

// ListenAndEcho listen for connection a echo
func (s *EchoServer) ListenAndEcho(fn handlerFunc) error {
	s.logger.Infof("Server started listening on %s", s.conn.Addr())

	// mainloop:
	for {
		conn, err := s.conn.Accept()
		if err != nil {
			return err
			// break mainloop
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

}
