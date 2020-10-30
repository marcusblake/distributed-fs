package main

import (
	"fmt"
	"os"
)

const (
	help = `filesysctl is a tool to help you deploy and manage resources for your distributed file system

Usage:

	filesysctl <command> [arguments]

You can use any of the following commands:

	deploy: Deploy service based on .yaml configuration file
	run: 	Run a server locally
`
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(help)
		os.Exit(1)
	}

	command := os.Args[1]
	cmd, ok := commands[command]
	if !ok {
		fmt.Printf("Command %s not found", command)
		fmt.Println(help)
		os.Exit(1)
	}

	arguments := os.Args[1:]
	if err := cmd(arguments); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
