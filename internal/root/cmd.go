package root

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:     "neiasit",
	Short:   "Neiasit platform sdk cli",
	Long:    "Neiasit platform sdk cli\nI hate makefile if i can make cli ><",
	Version: "1.0.0",
}
