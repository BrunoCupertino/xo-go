package server

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPlayersOnConnected(t *testing.T) {
	var (
		opts     = NewConnectorAcceptorOpts("8089")
		acceptor = NewTCPConnectionAcceptor(opts)
		room     = NewRoom(acceptor)
		err      error
	)

	assert.Nil(t, room.game)

	_, err = net.Dial("tcp", "localhost:8089")

	time.Sleep(time.Microsecond * 50)

	assert.NoError(t, err)
	assert.NotNil(t, room.game)

	_, err = net.Dial("tcp", "localhost:8089")

	time.Sleep(time.Microsecond * 50)

	assert.NoError(t, err)
}
