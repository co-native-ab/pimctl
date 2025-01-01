package active

import (
	"github.com/co-native-ab/pimctl/cmd/role/entra/active/list"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "active",
	Short: "Manage active Entra PIM roles",
	Long:  "Manage active Entra PIM roles",
}

func init() {
	Cmd.AddCommand(list.Cmd)
}
