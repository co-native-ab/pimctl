package clearAccount

import (
	"fmt"

	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear pimctl account cache",
	Long:  "Clear pimctl account cache",
	RunE:  runClearAccount,
}

func init() {}

func runClearAccount(cmd *cobra.Command, _ []string) error {
	err := cmdhelper.ClearCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to clear cache: %w", err)
	}
	return nil
}
