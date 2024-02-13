package server

import (
	// "encoding"
	"net"
	"testing"

	"github.com/BrunoCupertino/xo-go/internal/encoding"
	"github.com/BrunoCupertino/xo-go/internal/state"
	"github.com/stretchr/testify/assert"
)

func TestGameManager(t *testing.T) {
	var (
		acceptorOpts = NewConnectorAcceptorOpts("8089")
		acceptor     = NewTCPConnectionAcceptor(acceptorOpts)
		encoder      = encoding.NewStringStatementEncoder()
		msg          = make([]byte, 10)
		m            = NewGameManager(acceptor, encoder)
	)

	go func() {
		m.Start()
	}()

	p1Conn, err := net.Dial("tcp", ":8089")
	assert.NoError(t, err)

	n, err := p1Conn.Read(msg)
	content := msg[:n]
	statement, _ := encoder.Encode(state.NewStatement(state.TeamSelected, state.OTeam, 0))
	assert.Equal(t, statement, content)

	p2Conn, err := net.Dial("tcp", ":8089")
	assert.NoError(t, err)

	n, err = p2Conn.Read(msg)
	content = msg[:n]
	statement, _ = encoder.Encode(state.NewStatement(state.TeamSelected, state.XTeam, 0))
	assert.Equal(t, statement, content)

	playContent, _ := encoder.Encode(state.NewBoardStatement(state.OTeam, state.Square0))
	p1Conn.Write(playContent)

	n, err = p1Conn.Read(msg)
	content = msg[:n]
	statement, _ = encoder.Encode(state.NewBoardStatement(state.OTeam, state.Square0))
	assert.Equal(t, statement, content)

	n, err = p2Conn.Read(msg)
	content = msg[:n]
	statement, _ = encoder.Encode(state.NewBoardStatement(state.OTeam, state.Square0))
	assert.Equal(t, statement, content)
}
