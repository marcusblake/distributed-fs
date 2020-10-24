package main

import (
	"fmt"
	"os"

	cmn "github.com/distributed-fs/pkg/common"
	"github.com/distributed-fs/pkg/dfs"
)

var (
	chunkserver = ""
)

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	client := dfs.NewClient("")
	_, err := client.IssueFileIORequest(cmn.Open, "hello.txt", nil, 0, 0, chunkserver)
	if err != nil {
		fail(err)
	}

	data := []byte{'h', 'e', 'l', 'l', 'o'}
	_, err = client.IssueFileIORequest(cmn.Append, "hello.txt", data, 0, 0, chunkserver)
	if err != nil {
		fail(err)
	}

	response, err := client.IssueFileIORequest(cmn.Read, "hello.txt", nil, 10, 0, chunkserver)
	if err != nil {
		fail(err)
	}

	_, err = client.IssueFileIORequest(cmn.Close, "hello.txt", nil, 0, 0, chunkserver)
	if err != nil {
		fail(err)
	}

	fmt.Println(response)
}
