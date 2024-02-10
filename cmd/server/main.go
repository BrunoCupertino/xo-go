package main

import (
	"fmt"

	"github.com/BrunoCupertino/xo-go/internal/server"
)

func main() {
	acceptorOpts := server.NewConnectorAcceptorOpts("8088")
	acceptor := server.NewTCPConnectionAcceptor(acceptorOpts)

	r := server.NewRoom(acceptor)
	if r != nil {
		fmt.Println("room created")
		r.Start()
	}

	fmt.Println("shutdown")
}
