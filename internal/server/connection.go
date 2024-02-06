package server

import (
	"net"
)

type ConnectionAcceptor interface {
	Listen() (net.Conn, error)
}

type TCPConnectorAcceptor struct {
	opts     *ConnectorAcceptorOpts
	listener net.Listener
}

type ConnectorAcceptorOpts struct {
	port string
}

func NewConnectorAcceptorOpts(port string) *ConnectorAcceptorOpts {
	return &ConnectorAcceptorOpts{
		port: port,
	}
}

func NewTCPConnectionAcceptor(opts *ConnectorAcceptorOpts) *TCPConnectorAcceptor {
	ln, err := net.Listen("tcp", ":"+opts.port)
	if err != nil {
		panic(err)
	}

	return &TCPConnectorAcceptor{
		listener: ln,
		opts:     opts,
	}
}

func (c *TCPConnectorAcceptor) Listen() (net.Conn, error) {
	conn, err := c.listener.Accept()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
