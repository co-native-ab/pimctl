package role

import (
	"github.com/co-native-ab/pimctl/cmd/role/azure"
	"github.com/co-native-ab/pimctl/cmd/role/entra"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "role",
	Short: "Manage PIM roles",
	Long:  "Manage PIM roles",
}

func init() {
	Cmd.AddCommand(azure.Cmd)
	Cmd.AddCommand(entra.Cmd)
}
