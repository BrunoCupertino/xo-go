package enconding

import (
	"testing"

	"github.com/BrunoCupertino/xo-go/internal/state"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	var (
		stringEncoder = NewStringStatementEncoder()
		statement     = state.NewStatement(state.TeamSelected, state.OTeam, state.Square1)
	)

	encodedStatement, _ := stringEncoder.Encode(statement)
	decodedStatement, _ := stringEncoder.Decode(encodedStatement)

	assert.Equal(t, statement.State, decodedStatement.State)
	assert.Equal(t, statement.Team, decodedStatement.Team)
	assert.Equal(t, statement.Square, decodedStatement.Square)
}
