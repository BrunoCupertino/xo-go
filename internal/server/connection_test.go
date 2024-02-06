package server

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPAcceptor(t *testing.T) {
	var (
		conn net.Conn
		err  error

		opts     = NewConnectorAcceptorOpts("8088")
		acceptor = NewTCPConnectionAcceptor(opts)
	)

	chanRes := make(chan struct{})

	go func() {
		conn, err = acceptor.Listen()

		chanRes <- struct{}{}
	}()

	_, err2 := net.Dial("tcp", "localhost:8088")
	if err != nil {
		panic(err2)
	}

	<-chanRes

	assert.NotNil(t, conn)
	assert.NoError(t, err)
}
