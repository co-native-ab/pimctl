package cmd

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/co-native-ab/pimctl/cmd/account"
	"github.com/co-native-ab/pimctl/cmd/group"
	"github.com/co-native-ab/pimctl/cmd/login"
	"github.com/co-native-ab/pimctl/cmd/role"
	"github.com/co-native-ab/pimctl/internal/build"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/credentials"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:               "pimctl",
	Short:             "CLI to manage Azure PIM roles and assignments",
	Long:              "A CLI tool to manage Azure Privileged Identity Management (PIM) roles and assignments",
	Version:           fmt.Sprintf("%s (%s)", build.Version, build.Date),
	DisableAutoGenTag: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		command := getCommand(cmd.CommandPath())
		switch {
		case slices.Contains([]string{"login", "account clear", "account show", "account token"}, command):
			return nil
		case strings.HasPrefix(command, "group"):
			if !isLoggedInMSGraph(cmd.Context()) {
				return fmt.Errorf("not logged in, run 'login --scope MicrosoftGraph' to continue")
			}
		case strings.Contains(command, "role entra"):
			if !isLoggedInMSGraph(cmd.Context()) {
				return fmt.Errorf("not logged in, run 'login --scope MicrosoftGraph' to continue")
			}
		case strings.Contains(command, "role azure"):
			if !isLoggedInARM(cmd.Context()) {
				return fmt.Errorf("not logged in, run 'login --scope Azure' to continue")
			}
		default:
			return fmt.Errorf("unknown command: %s", command)
		}

		return nil
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.AutomaticEnv()
	envPrefix := strings.ToUpper(RootCmd.Use)
	viper.SetEnvPrefix(envPrefix)
	RootCmd.PersistentFlags().StringP("profile", "p", "default", fmt.Sprintf("The cache profile name for pimctl. Enables running multiple sessions with different users at the same time. [env: %s_PROFILE]", envPrefix))
	viper.BindPFlag("PROFILE", RootCmd.PersistentFlags().Lookup("profile"))
	RootCmd.AddCommand(account.Cmd)
	RootCmd.AddCommand(group.Cmd)
	RootCmd.AddCommand(role.Cmd)
	RootCmd.AddCommand(login.Cmd)
}

func getCommand(input string) string {
	parts := strings.Split(input, " ")
	if len(parts) == 1 {
		return ""
	}
	return strings.Join(parts[1:], " ")
}

func isLoggedInMSGraph(ctx context.Context) bool {
	cred, err := cmdhelper.NewCachedCredential()
	if err != nil {
		return false
	}

	scope := credentials.MicrosoftGraphPimctlScope
	_, err = cred.GetToken(ctx, policy.TokenRequestOptions{
		EnableCAE: true,
		Scopes:    scope.Scopes(),
	})

	return err == nil
}

func isLoggedInARM(ctx context.Context) bool {
	cred, err := cmdhelper.NewCachedCredential()
	if err != nil {
		return false
	}

	scope := credentials.AzurePimctlScope
	_, err = cred.GetToken(ctx, policy.TokenRequestOptions{
		EnableCAE: true,
		Scopes:    scope.Scopes(),
	})

	return err == nil
}
