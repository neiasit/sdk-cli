package generate_proto

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "generate-proto",
	Aliases: []string{"gen"},
	Short:   "Generate code from proto files",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
