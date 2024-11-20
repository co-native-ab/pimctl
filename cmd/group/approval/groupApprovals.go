package approval

import (
	"github.com/co-native-ab/pimctl/cmd/group/approval/create"
	"github.com/co-native-ab/pimctl/cmd/group/approval/list"
	"github.com/co-native-ab/pimctl/cmd/group/approval/menu"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "approval",
	Short: "Manage Azure PIM group approvals",
	Long:  "Manage Azure PIM group approvals",
}

func init() {
	Cmd.AddCommand(list.Cmd)
	Cmd.AddCommand(create.Cmd)
	Cmd.AddCommand(menu.Cmd)
}
