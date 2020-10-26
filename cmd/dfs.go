package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/akamensky/argparse"
	"github.com/distributed-fs/internal/chunkserver"
	"github.com/distributed-fs/internal/master"
	"github.com/go-yaml/yaml"
)

const (
	help = `dfs is a tool to help you deploy and manage resources for your distributed file system

Usage:

	dfs <command> [arguments]

You can use any of the following commands:

	deploy: Deploy service based on .yaml configuration file
	run: 	Run a server locally
`
)

var (
	arguments []string
)

// DeployYaml is the struct for the yaml
type DeployYaml struct {
	DeployServer string `yaml:"deployserver"`
	Master       struct {
		Port uint16 `yaml:"port"`
	}
	Chunkserver struct {
		Port uint16 `yaml:"port"`
	}
}

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

	deployment, err := ioutil.ReadFile(*filename)
	if err != nil {
		return err
	}

	deployYaml := DeployYaml{}

	if err := yaml.Unmarshal(deployment, &deployYaml); err != nil {
		return err
	}

	requestURL := "http://%s/startchunkserver"
	deployserver := fmt.Sprintf(requestURL, deployYaml.DeployServer)
	port := strconv.Itoa(int(deployYaml.Chunkserver.Port))
	body := url.Values{"Port": {port}}
	resp, err := http.PostForm(deployserver, body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func main() {
	var cmd func() error

	if len(os.Args) < 2 {
		fmt.Println(help)
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "deploy":
		cmd = deploy
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
