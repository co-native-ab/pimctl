package azure

import (
	"github.com/co-native-ab/pimctl/cmd/role/azure/active"
	"github.com/co-native-ab/pimctl/cmd/role/azure/approval"
	"github.com/co-native-ab/pimctl/cmd/role/azure/eligible"
	"github.com/co-native-ab/pimctl/cmd/role/azure/request"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "azure",
	Short: "Manage Azure PIM roles",
	Long:  "Manage Azure PIM roles",
}

func init() {
	Cmd.AddCommand(active.Cmd)
	Cmd.AddCommand(approval.Cmd)
	Cmd.AddCommand(eligible.Cmd)
	Cmd.AddCommand(request.Cmd)
}
