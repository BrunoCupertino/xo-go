package client

import (
	"testing"
	"time"

	"github.com/BrunoCupertino/xo-go/internal/encoding"
	"github.com/BrunoCupertino/xo-go/internal/server"
	"github.com/BrunoCupertino/xo-go/internal/state"
	"github.com/stretchr/testify/assert"
)

func TestGameClientManager(t *testing.T) {
	var (
		serverOpts = server.NewConnectorAcceptorOpts("8090")
		serverAcc  = server.NewTCPConnectionAcceptor(serverOpts)
		encoder    = encoding.NewStringStatementEncoder()
		server     = server.NewGameManager(serverAcc, encoder)
		opts       = NewTCPGameClientManagerOpts(":8090")
		renderer   = NewNoOpRenderer()
		client1    = NewTCPGameClientManager(opts, encoder, renderer)
		client2    = NewTCPGameClientManager(opts, encoder, renderer)
	)

	go func() {
		go server.Start()
		go client1.Start()
		go client2.Start()
	}()

	time.Sleep(time.Millisecond * 5)

	square := state.Square0
	client1.Send(square)
	time.Sleep(time.Millisecond * 5)
	assert.Equal(t, client1.state.board[square], client1.state.MyTeam())
	assert.Equal(t, client2.state.board[square], client1.state.MyTeam())

	square = state.Square1
	client2.Send(square)
	time.Sleep(time.Millisecond * 5)
	assert.Equal(t, client1.state.board[square], client2.state.MyTeam())
	assert.Equal(t, client2.state.board[square], client2.state.MyTeam())
}
