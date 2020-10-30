package main

import (
	"fmt"
	"os"

	chnksrv "github.com/distributed-fs/internal/chunkserver"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./run <port>")
		os.Exit(1)
	}
	chunkserver := chnksrv.NewChunkserver()
	port := fmt.Sprintf(":%s", os.Args[1])
	if err := chunkserver.Start(port); err != nil {
		fmt.Printf("Chunkserver failed to start: %s", err.Error())
		os.Exit(1)
	}
}
