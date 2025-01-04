package request

import (
	"github.com/co-native-ab/pimctl/cmd/role/azure/request/list"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "request",
	Short: "Manage Azure PIM role assignment requests",
	Long:  "Manage Azure PIM role assignment requests",
}

func init() {
	Cmd.AddCommand(list.Cmd)
}
