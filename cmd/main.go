package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/akamensky/argparse"
	"github.com/distributed-fs/internal/chunkserver"
	"github.com/distributed-fs/internal/master"
)

const (
	help = `dfs is a tool to help you deploy and manage resources for your distributed file system

Usage:

	dfs <command> [arguments]

You can use any of the following commands:

	deploy: Deploy service based on .yaml configuration file
`
)

var (
	arguments []string
)

func run() error {
	parser := argparse.NewParser("run", "runs a server")

	servertype := parser.String("t", "type", &argparse.Options{Required: true, Help: "Type of server to run"})
	port := parser.Int("p", "port", &argparse.Options{Required: true, Help: "Port to run server on"})

	if err := parser.Parse(arguments); err != nil {
		return errors.New(parser.Usage(err))
	}

	address := fmt.Sprintf("localhost:%d", *port)
	switch *servertype {
	case "master":
		master := master.NewMaster()
		if err := master.Start(address); err != nil {
			return err
		}
	case "chunkserver":
		chunkserver := chunkserver.NewChunkserver()
		if err := chunkserver.Start(address); err != nil {
			return err
		}
	default:
		return errors.New("Unknown server type")
	}

	for {
		time.Sleep(10 * time.Second)
	}

	return nil
}

func deploy() error {
	parser := argparse.NewParser("deploy", " deploys resources for you")

	filename := parser.String("f", "file", &argparse.Options{Required: true, Help: "YAML file which contains deployment configuration"})

	if err := parser.Parse(arguments); err != nil {
		return errors.New(parser.Usage(err))
	}

	fmt.Println(*filename)

	return nil
}

func main() {

	command := os.Args[1]

	var cmd func() error

	switch command {
	case "deploy":
		cmd = deploy
		break
	case "run":
		cmd = run
	default:
		fmt.Println(help)
		os.Exit(1)
	}
	arguments = os.Args[1:]
	if err := cmd(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
