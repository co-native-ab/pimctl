package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/co-native-ab/pimctl/cmd/account"
	"github.com/co-native-ab/pimctl/cmd/group"
	"github.com/co-native-ab/pimctl/cmd/login"
	"github.com/co-native-ab/pimctl/cmd/role"
	"github.com/co-native-ab/pimctl/internal/build"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"

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
		switch command {
		case "login", "account clear", "account show":
			return nil
		}

		if !isLoggedIn(cmd.Context()) {
			return fmt.Errorf("not logged in, run 'login' to continue")
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

func isLoggedIn(ctx context.Context) bool {
	cred, scopes, err := cmdhelper.NewCachedCredential()
	if err != nil {
		return false
	}

	_, err = cred.GetToken(ctx, policy.TokenRequestOptions{
		EnableCAE: true,
		Scopes:    scopes,
	})
	return err == nil
}
