package active

import (
	"github.com/co-native-ab/pimctl/cmd/group/active/list"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "active",
	Short: "Manage Azure PIM active groups",
	Long:  "Manage Azure PIM active groups",
}

func init() {
	Cmd.AddCommand(list.Cmd)
}
