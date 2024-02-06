package main

import (
	"fmt"

	"github.com/BrunoCupertino/xo-go/internal/client"
)

type ConsoleRenderer struct{}

func (r *ConsoleRenderer) Render(s client.ClientState) {
	fmt.Println("rendering...")
	fmt.Println("rendered")
}
