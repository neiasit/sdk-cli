package main

import (
	"log"
	"sdk-cli/internal/generate_proto"
	"sdk-cli/internal/initialize"
	"sdk-cli/internal/root"
)

func main() {
	root.Cmd.AddCommand(
		initialize.NewCmd(),
		generate_proto.Cmd,
	)
	err := root.Cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
