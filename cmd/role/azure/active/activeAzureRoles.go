package active

import (
	"github.com/co-native-ab/pimctl/cmd/role/azure/active/list"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "active",
	Short: "Manage active Azure PIM roles",
	Long:  "Manage active Azure PIM roles",
}

func init() {
	Cmd.AddCommand(list.Cmd)
}
