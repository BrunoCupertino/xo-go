package server

import (
	"fmt"
	"io"
	"net"
)

type ConnectionAcceptor interface {
	io.Closer
	ListenAndAccept() (net.Conn, error)
}

type TCPConnectionAcceptor struct {
	opts     *ConnectionAcceptorOpts
	listener net.Listener
}

type ConnectionAcceptorOpts struct {
	port string
}

func NewConnectorAcceptorOpts(port string) *ConnectionAcceptorOpts {
	return &ConnectionAcceptorOpts{
		port: port,
	}
}

func NewTCPConnectionAcceptor(opts *ConnectionAcceptorOpts) *TCPConnectionAcceptor {
	ln, err := net.Listen("tcp", ":"+opts.port)
	if err != nil {
		panic(err)
	}

	fmt.Printf("listening on port %s \n", opts.port)

	return &TCPConnectionAcceptor{
		listener: ln,
		opts:     opts,
	}
}

func (c *TCPConnectionAcceptor) ListenAndAccept() (net.Conn, error) {
	conn, err := c.listener.Accept()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *TCPConnectionAcceptor) Close() error {
	return c.listener.Close()
}
