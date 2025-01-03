package eligible

import (
	"github.com/co-native-ab/pimctl/cmd/role/azure/eligible/list"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "eligible",
	Short: "Manage Azure PIM eligible roles",
	Long:  "Manage Azure PIM eligible roles",
}

func init() {
	Cmd.AddCommand(list.Cmd)
}
