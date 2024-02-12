package enconding

import (
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
	return nil, nil
}

func (e *StringStatementEncoder) Decode(data []byte) (*state.Statement, error) {
	return nil, nil
}
