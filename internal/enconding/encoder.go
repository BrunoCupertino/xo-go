package enconding

import (
	"fmt"
	"strconv"

	"github.com/BrunoCupertino/xo-go/internal/state"
)

type StatementEncoder interface {
	Encode(s *state.Statement) ([]byte, error)
	Decode(data []byte) (*state.Statement, error)
}

type StringStatementEncoder struct{}

func NewStringStatementEncoder() *StringStatementEncoder {
	return &StringStatementEncoder{}
}

func (e *StringStatementEncoder) Encode(s *state.Statement) ([]byte, error) {
	var content = fmt.Sprintf("%s%s%d", s.State, s.Team, s.Square)

	return []byte(content), nil
}

func (e *StringStatementEncoder) Decode(data []byte) (*state.Statement, error) {
	var content = string(data)

	stateChange := state.StateChange(content[0])
	team := state.Team(content[1])

	s, err := strconv.Atoi(string(content[2]))
	if err != nil {
		return nil, err
	}

	square := state.Square(s)

	return state.NewStatement(stateChange, team, square), nil
}
