package show

import (
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "show",
	Short: "Show pimctl account cache",
	Long:  "Show pimctl account cache",
	RunE:  runShowAccount,
}

func init() {}

func runShowAccount(cmd *cobra.Command, _ []string) error {
	cachedRecord, err := cmdhelper.GetCachedCredentialAuthenticationRecord()
	if err != nil {
		return fmt.Errorf("failed to clear cache: %w", err)
	}

	if cachedRecord == (azidentity.AuthenticationRecord{}) {
		fmt.Println("No cached account")
		return nil
	}

	output, err := json.MarshalIndent(cachedRecord, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal cached account: %w", err)
	}

	fmt.Println(string(output))

	return nil
}
