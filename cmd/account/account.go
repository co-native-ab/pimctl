package account

import (
	clearAccount "github.com/co-native-ab/pimctl/cmd/account/clear"
	"github.com/co-native-ab/pimctl/cmd/account/show"
	"github.com/co-native-ab/pimctl/cmd/account/token"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "account",
	Short: "Manage pimctl account",
	Long:  "Manage pimctl account",
}

func init() {
	Cmd.AddCommand(clearAccount.Cmd)
	Cmd.AddCommand(show.Cmd)
	Cmd.AddCommand(token.Cmd)
}
