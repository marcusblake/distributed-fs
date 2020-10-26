package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/distributed-fs/internal"
)

var (
	gopath = ""
)

const (
	startChunkserver = "%s/src/github.com/distributed-fs/cmd/startchunkserver/run"
)

func StartMaster(res http.ResponseWriter, req *http.Request) {
}

func StartChunkserver(res http.ResponseWriter, req *http.Request) {
	program := fmt.Sprintf(startChunkserver, gopath)
	fmt.Println("received request")
	port := req.FormValue("Port")
	if port == "" {
		fmt.Fprint(res, "Must provide port")
		fmt.Println("failed")
		return
	}
	cmd := exec.Command(program, port)
	if err := cmd.Start(); err != nil {
		fmt.Fprintf(res, "Failed to run chunkserver: %s", err.Error())
		fmt.Println("Failed to run chunkserver: %s", err.Error())
	}
	cmd.Process.Release()
	internal.Success("success")
}

func StartApiServer(port string) {
	gopath = os.Getenv("GOPATH")
	fmt.Println(gopath)
	http.HandleFunc("/startmaster", StartMaster)
	http.HandleFunc("/startchunkserver", StartChunkserver)

	portStr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(portStr, nil)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./deployserver <port>")
		os.Exit(1)
	}

	port := os.Args[1]
	StartApiServer(port)
}
