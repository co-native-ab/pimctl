package eligible

import (
	"github.com/co-native-ab/pimctl/cmd/group/eligible/list"
	"github.com/co-native-ab/pimctl/cmd/group/eligible/menu"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "eligible",
	Short: "Manage Azure PIM eligible groups",
	Long:  "Manage Azure PIM eligible groups",
}

func init() {
	Cmd.AddCommand(list.Cmd)
	Cmd.AddCommand(menu.Cmd)
}
