package entra

import (
	"github.com/co-native-ab/pimctl/cmd/role/entra/active"
	"github.com/co-native-ab/pimctl/cmd/role/entra/approval"
	"github.com/co-native-ab/pimctl/cmd/role/entra/eligible"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "entra",
	Short: "Manage Entra PIM roles",
	Long:  "Manage Entra PIM roles",
}

func init() {
	Cmd.AddCommand(active.Cmd)
	Cmd.AddCommand(approval.Cmd)
	Cmd.AddCommand(eligible.Cmd)
}
