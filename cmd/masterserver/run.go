package main

import (
	"fmt"
	"os"

	mstr "github.com/distributed-fs/internal/master"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./run <port>")
		os.Exit(1)
	}
	masterserver := mstr.NewMaster()
	port := fmt.Sprintf(":%s", os.Args[1])
	if err := masterserver.Start(port); err != nil {
		fmt.Printf("Chunkserver failed to start: %s", err.Error())
		os.Exit(1)
	}
}
