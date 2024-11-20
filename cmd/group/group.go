package group

import (
	"github.com/co-native-ab/pimctl/cmd/group/active"
	"github.com/co-native-ab/pimctl/cmd/group/approval"
	"github.com/co-native-ab/pimctl/cmd/group/eligible"
	"github.com/co-native-ab/pimctl/cmd/group/request"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "group",
	Short: "Manage Azure PIM groups",
	Long:  "Manage Azure PIM groups",
}

func init() {
	Cmd.AddCommand(active.Cmd)
	Cmd.AddCommand(approval.Cmd)
	Cmd.AddCommand(eligible.Cmd)
	Cmd.AddCommand(request.Cmd)
}
