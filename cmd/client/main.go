package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BrunoCupertino/xo-go/cmd/client/renderer"
	"github.com/BrunoCupertino/xo-go/internal/client"
	"github.com/BrunoCupertino/xo-go/internal/encoding"
	"github.com/BrunoCupertino/xo-go/internal/state"
)

func main() {
	fmt.Println("connectig to server...")

	var (
		opts          = client.NewTCPGameClientManagerOpts(":8088")
		encoder       = encoding.NewStringStatementEncoder()
		renderer      = renderer.NewConsoleRenderer()
		clientManager = client.NewTCPGameClientManager(opts, encoder, renderer)
	)

	go func() {
		reader := bufio.NewReader(os.Stdin)

		for {
			input, _ := reader.ReadString('\n')
			input = strings.ReplaceAll(input, "\n", "")

			num, err := strconv.Atoi(input)
			if err != nil {
				panic(err)
			}

			clientManager.Send(state.Square(num))
		}
	}()

	clientManager.Start()

	fmt.Println("shutdown")
}
