package main

import (
	"github.com/neiasit/sdk-cli/internal/generate_proto"
	"github.com/neiasit/sdk-cli/internal/initialize"
	"github.com/neiasit/sdk-cli/internal/root"
	"github.com/neiasit/sdk-cli/internal/vendor_proto"
	"log"
)

func main() {
	root.Cmd.AddCommand(
		initialize.NewCmd(),
		generate_proto.Cmd,
		vendor_proto.NewCmd(),
	)
	err := root.Cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
