package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestPlayersOnConnected(t *testing.T) {
// 	var (
// 		opts     = NewConnectorAcceptorOpts("8089")
// 		acceptor = NewTCPConnectionAcceptor(opts)
// 		room     = NewRoom(acceptor)
// 		err      error
// 	)

// 	assert.Nil(t, room.game)

// 	_, err = net.Dial("tcp", "localhost:8089")
// 	time.Sleep(time.Millisecond * 5)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, room.game)

// 	_, err = net.Dial("tcp", "localhost:8089")
// 	assert.NoError(t, err)
// }

func TestCreatePlayerAndJoin(t *testing.T) {
	r := NewRoom()

	p, err := r.CreatePlayerAndJoin()

	assert.NoError(t, err)
	assert.Equal(t, p, r.GetPlayerByTeam(p.GetTeam()))
}

func TestPlay(t *testing.T) {
	r := NewRoom()

	p1, err := r.CreatePlayerAndJoin()
	assert.NoError(t, err)

	p2, err := r.CreatePlayerAndJoin()
	assert.NoError(t, err)

	assert.NotEqual(t, p1.GetTeam(), p2.GetTeam())
	assert.Equal(t, p2, r.GetPlayerByTeam(p2.GetTeam()))
}
