package server

import (
	"testing"

	"github.com/BrunoCupertino/xo-go/internal/state"
	"github.com/stretchr/testify/assert"
)

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

	r.Play(*state.NewBoardStatement(p1.GetTeam(), state.Square0))
	r.Play(*state.NewBoardStatement(p2.GetTeam(), state.Square3))

	r.Play(*state.NewBoardStatement(p1.GetTeam(), state.Square1))
	r.Play(*state.NewBoardStatement(p2.GetTeam(), state.Square4))

	statement, gameStatus, err := r.Play(
		*state.NewBoardStatement(p1.GetTeam(), state.Square2))

	assert.Equal(t, state.GameOvered, statement.State)
	assert.Equal(t, state.OTeam, statement.Team)
	assert.Equal(t, state.Square2, statement.Square)
	assert.Equal(t, state.GameOver, gameStatus)
	assert.NoError(t, err)
}
