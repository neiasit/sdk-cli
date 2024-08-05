package vendor_proto

import (
	"errors"
	"github.com/neiasit/sdk-cli/internal/vendor_proto/usecase"
	"os"

	"github.com/labstack/gommon/color"
	"github.com/spf13/cobra"
)

const (
	stdLibsFlag          = "std"
	vendorProtobufFolder = "vendor.protobuf"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "vendor-proto",
		Aliases: []string{"vend"},
		Short:   "Vendor proto files from github",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := os.Mkdir(vendorProtobufFolder, os.ModePerm)
			if err != nil {
				if !errors.Is(err, os.ErrExist) {
					return err
				}
				println(color.Yellow("vendor.protobuf directory already exists"))
			}
			val, err := cmd.Flags().GetBool(stdLibsFlag)
			if err != nil {
				return err
			}
			if val {
				// Вендоринг стандартных библиотек
				err = usecase.VendorStandardLibraries()
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().Bool(stdLibsFlag, false, "use standard proto files")
	return cmd
}
