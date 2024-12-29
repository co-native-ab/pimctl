package approval

import (
	"github.com/co-native-ab/pimctl/cmd/role/entra/approval/create"
	"github.com/co-native-ab/pimctl/cmd/role/entra/approval/list"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "approval",
	Short: "Manage Entra PIM role approvals",
	Long:  "Manage Entra PIM role approvals",
}

func init() {
	Cmd.AddCommand(create.Cmd)
	Cmd.AddCommand(list.Cmd)
}
