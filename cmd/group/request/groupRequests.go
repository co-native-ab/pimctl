package request

import (
	"github.com/co-native-ab/pimctl/cmd/group/request/create"
	"github.com/co-native-ab/pimctl/cmd/group/request/list"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "request",
	Short: "Manage Azure PIM group assignment requests",
	Long:  "Manage Azure PIM group assignment requests",
}

func init() {
	Cmd.AddCommand(list.Cmd)
	Cmd.AddCommand(create.Cmd)
}
