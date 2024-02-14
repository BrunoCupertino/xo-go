package client

import (
	"testing"

	"github.com/BrunoCupertino/xo-go/internal/state"
	"github.com/stretchr/testify/assert"
)

func TestChange(t *testing.T) {
	var (
		gs       = NewClientGameState(state.OTeam)
		s        = state.NewStatement(state.TeamSelected, state.OTeam, state.Square0)
		noWinner = state.Team("")
	)

	assert.Equal(t, s.Team, gs.MyTeam())

	gs.Change(s)
	assert.Equal(t, s.State, gs.CurrentChange())
	assert.NotEqual(t, s.Team, gs.board[s.Square])

	s = state.NewBoardStatement(state.OTeam, state.Square0)

	gs.Change(s)
	assert.Equal(t, s.State, gs.CurrentChange())
	assert.Equal(t, s.Team, gs.board[s.Square])
	assert.Equal(t, byte(1), gs.round)
	assert.False(t, gs.IsMyTurn())
	assert.Equal(t, noWinner, gs.Winner())

	s = state.NewBoardStatement(state.XTeam, state.Square1)

	gs.Change(s)
	assert.Equal(t, s.State, gs.CurrentChange())
	assert.Equal(t, s.Team, gs.board[s.Square])
	assert.Equal(t, byte(2), gs.round)
	assert.True(t, gs.IsMyTurn())
	assert.Equal(t, noWinner, gs.Winner())

	s = state.NewStatement(state.GameOvered, state.OTeam, state.Square2)

	gs.Change(s)
	assert.Equal(t, s.State, gs.CurrentChange())
	assert.Equal(t, s.Team, gs.board[s.Square])
	assert.Equal(t, byte(3), gs.round)
	assert.False(t, gs.IsMyTurn())
	assert.Equal(t, s.Team, gs.Winner())
}
