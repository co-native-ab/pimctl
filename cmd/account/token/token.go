package token

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/credentials"
	"github.com/lestrrat-go/jwx/jwt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "token",
	Short: "Get pimctl token",
	Long:  "Get pimctl token",
	RunE:  runToken,
}

func init() {
	Cmd.Flags().String("output", "json", "output format. can be 'json' or 'raw'")
	pimctlScope := credentials.MicrosoftGraphPimctlScope
	Cmd.Flags().Var(&pimctlScope, "scope", pimctlScope.HelpText())
	Cmd.RegisterFlagCompletionFunc("scope", pimctlScope.CobraCompletion)
}

func runToken(cmd *cobra.Command, _ []string) error {
	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		return fmt.Errorf("failed to get output: %w", err)
	}

	pimctlScopeString := cmd.Flags().Lookup("scope").Value.String()
	pimctlScope := credentials.UnknownPimctlScope
	err = pimctlScope.Set(pimctlScopeString)
	if err != nil {
		return fmt.Errorf("failed to set scope: %w", err)
	}

	return getToken(cmd.Context(), outputFormat, pimctlScope)
}

func getToken(ctx context.Context, outputFormat string, pimctlScope credentials.PimctlScope) error {
	cred, err := cmdhelper.NewCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	cachedAuthenticationRecord, err := cmdhelper.GetCachedCredentialAuthenticationRecord()
	if err != nil {
		return fmt.Errorf("failed to get cached authentication record: %w", err)
	}

	token, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes:    pimctlScope.Scopes(),
		EnableCAE: true,
		TenantID:  cachedAuthenticationRecord.TenantID,
	})

	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	if outputFormat == "raw" {
		fmt.Println(token.Token)
		return nil
	}

	parsedToken, err := jwt.ParseString(token.Token)
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	parsedJsonToken, err := json.MarshalIndent(parsedToken, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal token: %w", err)
	}

	fmt.Println(string(parsedJsonToken))

	return nil
}
