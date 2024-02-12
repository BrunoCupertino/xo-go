package main

import (
	"fmt"
	"time"

	"github.com/BrunoCupertino/xo-go/internal/client"
)

func main() {
	fmt.Println("connectig to server...")

	var (
		opts      = client.NewTCPConnectorOpts(":8088")
		connector = client.NewTCPConnector(opts)
	)

	err := connector.Connect()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			time.Sleep(time.Second * 1)
			// connector.Send([]byte("hi"))
		}
	}()

	connector.Start()

	fmt.Println("shutdown")
}
