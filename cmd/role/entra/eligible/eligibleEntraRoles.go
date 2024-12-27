package eligible

import (
	"github.com/co-native-ab/pimctl/cmd/role/entra/eligible/list"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "eligible",
	Short: "Manage Entra PIM eligible roles",
	Long:  "Manage Entra PIM eligible roles",
}

func init() {
	Cmd.AddCommand(list.Cmd)
}
