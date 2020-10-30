package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/distributed-fs/pkg/logger"
)

var (
	gopath = ""
)

const (
	startMasterserver = "%s/src/github.com/distributed-fs/cmd/master/run"
	startChunkserver  = "%s/src/github.com/distributed-fs/cmd/chunkserver/run"
)

func StartServer(port string, program string) error {
	if port == "" {
		return fmt.Errorf("Must provide port")
	}

	cmd := exec.Command(program, port)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Failed to run chunkserver: %s", err.Error())
	}

	return cmd.Process.Release()
}

func StartMaster(res http.ResponseWriter, req *http.Request) {
	program := fmt.Sprintf(startMasterserver, gopath)
	logger.Infof("Received request to start master server from %s", req.RemoteAddr)
	port := req.FormValue("Port")
	if err := StartServer(port, program); err != nil {
		logger.Warning(err.Error())
		fmt.Fprintf(res, "Failed to run master: %s", err.Error())
	}
}

func StartChunkserver(res http.ResponseWriter, req *http.Request) {
	program := fmt.Sprintf(startChunkserver, gopath)
	logger.Infof("Received request to start chunkserver from %s", req.RemoteAddr)
	port := req.FormValue("Port")
	if err := StartServer(port, program); err != nil {
		logger.Warning(err.Error())
		fmt.Fprintf(res, "Failed to run master: %s", err.Error())
	}
}

func StartApiServer(port string) {
	gopath = os.Getenv("GOPATH")
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
