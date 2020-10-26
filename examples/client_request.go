package main

import (
	"fmt"
	"os"

	"github.com/distributed-fs/pkg/dfs"
)

var (
	chunkserver = ":8080"
)

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	client := dfs.NewClient("")
	err := client.Open("hello.txt", chunkserver)
	if err != nil {
		fail(err)
	}

	data := []byte{'h', 'e', 'l', 'l', 'o'}
	err = client.Append("hello.txt", data, chunkserver)
	if err != nil {
		fail(err)
	}

	response, err := client.Read("hello.txt", 5, 0, chunkserver)
	if err != nil {
		fail(err)
	}

	err = client.Close("hello.txt", chunkserver)
	if err != nil {
		fail(err)
	}

	fmt.Println(response)
}
