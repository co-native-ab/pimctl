package azure

import (
	"github.com/co-native-ab/pimctl/cmd/role/azure/eligible"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "azure",
	Short: "Manage Azure PIM roles",
	Long:  "Manage Azure PIM roles",
}

func init() {
	Cmd.AddCommand(eligible.Cmd)
}