package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	var (
		p1 = NewHumanPlayer(XTeam)
		g  = NewGame(p1)
	)

	assert.Equal(t, GameCreated, g.state)
	assert.Equal(t, p1, g.player1)
}

func TestJoin(t *testing.T) {
	var (
		p1 = NewHumanPlayer(XTeam)
		p2 = NewHumanPlayer(XTeam)
		g  = NewGame(p1)
	)

	err := g.Join(p2)

	assert.ErrorIs(t, err, ErrTeamAlreadySelected)
	assert.Equal(t, GameCreated, g.state)
	assert.Nil(t, g.player2)

	p2 = NewHumanPlayer(OTeam)

	err = g.Join(p2)

	assert.NoError(t, err)
	assert.Equal(t, GameStarted, g.state)
	assert.Equal(t, p2, g.player2)

	err = g.Join(p2)
	assert.ErrorIs(t, err, ErrCantJoinStartedGame)
	assert.Equal(t, GameStarted, g.state)
}

func TestPlay(t *testing.T) {
	var (
		p1 = NewHumanPlayer(XTeam)
		p2 = NewHumanPlayer(OTeam)
		g  = NewGame(p1)
	)

	err := g.Play(p1, Statement0)
	assert.ErrorIs(t, err, ErrWaitingPlayerJoin)

	g.Join(p2)

	err = g.Play(p1, Statement0)
	assert.NoError(t, err)

	err = g.Play(p1, Statement0)
	assert.ErrorIs(t, err, ErrCanPlayTwice)

	err = g.Play(p2, Statement0)
	assert.ErrorIs(t, err, ErrInvalidStatement)

	err = g.Play(p2, Statement1)
	assert.NoError(t, err)
}
