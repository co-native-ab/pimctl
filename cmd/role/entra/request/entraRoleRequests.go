package request

import (
	"github.com/co-native-ab/pimctl/cmd/role/entra/request/list"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "request",
	Short: "Manage Entra PIM role assignment requests",
	Long:  "Manage Entra PIM role assignment requests",
}

func init() {
	Cmd.AddCommand(list.Cmd)
}
