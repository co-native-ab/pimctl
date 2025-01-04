package approval

import (
	"github.com/co-native-ab/pimctl/cmd/role/azure/approval/list"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "approval",
	Short: "Manage Azure PIM role approvals",
	Long:  "Manage Azure PIM role approvals",
}

func init() {
	Cmd.AddCommand(list.Cmd)
}
