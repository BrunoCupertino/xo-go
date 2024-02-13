package main

import (
	"fmt"

	"github.com/BrunoCupertino/xo-go/internal/encoding"
	"github.com/BrunoCupertino/xo-go/internal/server"
)

func main() {
	acceptorOpts := server.NewConnectorAcceptorOpts("8088")
	acceptor := server.NewTCPConnectionAcceptor(acceptorOpts)
	encoding := encoding.NewStringStatementEncoder()

	r := server.NewGameManager(acceptor, encoding)
	if r != nil {
		fmt.Println("ready to start")
		r.Start()
	}

	fmt.Println("shutdown")
}
