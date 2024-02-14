package client

type Renderer interface {
	Render(s *ClientGameState)
}

type NoOpRenderer struct{}

func NewNoOpRenderer() *NoOpRenderer {
	return &NoOpRenderer{}
}

func (r *NoOpRenderer) Render(s *ClientGameState) {}
