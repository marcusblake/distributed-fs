package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/akamensky/argparse"
	"github.com/distributed-fs/internal/chunkserver"
	"github.com/distributed-fs/internal/master"
	"github.com/go-yaml/yaml"
)

// Command defines the type for a command
type Command func([]string) error

var (
	commands = map[string]Command{
		"run":   run,
		"deply": deploy,
	}
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

func run(arguments []string) error {
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

	return nil
}

func deploy(arguments []string) error {
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
