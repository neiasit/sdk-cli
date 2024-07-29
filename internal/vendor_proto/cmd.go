package generate_grpc

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "vendor-proto",
	Aliases: []string{"vend"},
	Short:   "Vendor proto files from github",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
