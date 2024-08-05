package initialize

import (
	"fmt"
	"github.com/labstack/gommon/color"
	"github.com/neiasit/sdk-cli/internal/initialize/ui"
	"github.com/spf13/cobra"
	"os"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "initialize",
		Aliases: []string{"init", "i"},
		Short:   "Initialize a new project",
		Run: func(cmd *cobra.Command, args []string) {
			isThis, err := cmd.Flags().GetBool("this")
			if err != nil {
				println(color.Red(fmt.Sprintf("Error: %s", err), color.B))
				os.Exit(1)
			}
			if err := ui.DisplayMenu(isThis); err != nil {
				println(color.Red(fmt.Sprintf("Error: %s", err), color.B))
				os.Exit(1)
			}
		},
	}
	cmd.Flags().Bool("this", false, "Create a new project in current directory")
	return cmd
}
