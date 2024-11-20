package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/co-native-ab/pimctl/cmd/account"
	"github.com/co-native-ab/pimctl/cmd/group"
	"github.com/co-native-ab/pimctl/cmd/login"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "pimctl",
	Short:   "CLI to manage Azure PIM roles and assignments",
	Long:    "A CLI tool to manage Azure Privileged Identity Management (PIM) roles and assignments",
	Version: getVersion(),
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
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func getBuildInfoKey(buildInfo *debug.BuildInfo, key string) string {
	for _, setting := range buildInfo.Settings {
		if setting.Key == key {
			return setting.Value
		}
	}

	return ""
}

func toPtrIfNotEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func toNilIfFalse(b bool) *bool {
	if !b {
		return nil
	}
	return &b
}

func getVersion() string {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return `{"error": "failed to read build info"}`
	}

	vcsRevision := getBuildInfoKey(buildInfo, "vcs.revision")
	lastCommit := ""
	lastCommitDateTime, err := time.Parse(time.RFC3339, getBuildInfoKey(buildInfo, "vcs.time"))
	if err == nil {
		lastCommit = lastCommitDateTime.Format(time.RFC3339)
	}
	dirtyBuild := getBuildInfoKey(buildInfo, "vcs.dirty") == "true"
	version := buildInfo.Main.Version

	output := struct {
		Version    *string `json:"version,omitempty"`
		Revision   *string `json:"revision,omitempty"`
		LastCommit *string `json:"lastCommit,omitempty"`
		Dirty      *bool   `json:"dirty,omitempty"`
	}{
		Version:    toPtrIfNotEmpty(version),
		Revision:   toPtrIfNotEmpty(vcsRevision),
		LastCommit: toPtrIfNotEmpty(lastCommit),
		Dirty:      toNilIfFalse(dirtyBuild),
	}

	outputBytes, err := json.Marshal(output)
	if err != nil {
		return fmt.Sprintf(`{"error": "failed to marshal version info: %s"}`, err)
	}

	if string(outputBytes) == "{}" {
		return `{"error": "failed to get version info"}`
	}

	return string(outputBytes)
}

func init() {
	viper.AutomaticEnv()
	envPrefix := strings.ToUpper(rootCmd.Use)
	viper.SetEnvPrefix(envPrefix)
	rootCmd.PersistentFlags().StringP("profile", "p", "default", fmt.Sprintf("The cache profile name for pimctl. Enables running multiple sessions with different users at the same time. [env: %s_PROFILE]", envPrefix))
	viper.BindPFlag("PROFILE", rootCmd.PersistentFlags().Lookup("profile"))
	rootCmd.AddCommand(account.Cmd)
	rootCmd.AddCommand(group.Cmd)
	rootCmd.AddCommand(login.Cmd)
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
