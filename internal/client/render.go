package client

import "fmt"

type Renderer interface {
	Render(s ClientState)
}

type ConsoleRenderer struct{}

func (r *ConsoleRenderer) Render(s ClientState) {
	fmt.Println("rendering...")
	fmt.Println("rendered")
}
