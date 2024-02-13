package server

import (
	"net"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		conn, err = acceptor.ListenAndAccept()

		wg.Done()
	}()

	_, err2 := net.Dial("tcp", "localhost:8088")
	if err != nil {
		panic(err2)
	}

	wg.Wait()

	assert.NotNil(t, conn)
	assert.NoError(t, err)
}
