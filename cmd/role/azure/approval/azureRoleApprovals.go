package approval

import (
	"github.com/co-native-ab/pimctl/cmd/role/azure/approval/create"
	"github.com/co-native-ab/pimctl/cmd/role/azure/approval/list"
	"github.com/co-native-ab/pimctl/cmd/role/azure/approval/menu"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "approval",
	Short: "Manage Azure PIM role approvals",
	Long:  "Manage Azure PIM role approvals",
}

func init() {
	Cmd.AddCommand(create.Cmd)
	Cmd.AddCommand(list.Cmd)
	Cmd.AddCommand(menu.Cmd)
}
