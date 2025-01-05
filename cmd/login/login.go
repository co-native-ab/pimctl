package login

import (
	"context"
	"fmt"

	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/credentials"
	"github.com/co-native-ab/pimctl/internal/graph"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var Cmd = &cobra.Command{
	Use:   "login",
	Short: "Login to pimctl",
	Long:  "Login to pimctl. This command will authenticate you to Azure PIM",
	RunE:  runLogin,
}

func init() {
	Cmd.Flags().String("client-id", "", "Entra Application Client ID for pimctl")
	Cmd.Flags().String("tenant-id", credentials.DefaultTenantID, "Entra Tenant ID to use for login")
	credentialMethod := credentials.InteractiveBrowserCredentialMethod
	Cmd.Flags().Var(&credentialMethod, "credential-method", credentialMethod.HelpText())
	Cmd.RegisterFlagCompletionFunc("credential-method", credentialMethod.CobraCompletion)
}

type loginOptions struct {
	clientID         string
	tenantID         string
	credentialMethod credentials.CredentialMethod
}

func newLoginOptions(flags *pflag.FlagSet) (*loginOptions, error) {
	clientID, err := flags.GetString("client-id")
	if err != nil {
		return nil, fmt.Errorf("failed to get client-id: %w", err)
	}

	tenantID, err := flags.GetString("tenant-id")
	if err != nil {
		return nil, fmt.Errorf("failed to get tenant-id: %w", err)
	}

	credentialMethodString := flags.Lookup("credential-method").Value.String()
	credentialMethod := credentials.UnknownCredentialMethod
	err = credentialMethod.Set(credentialMethodString)
	if err != nil {
		return nil, fmt.Errorf("failed to set credential-method: %w", err)
	}

	return &loginOptions{
		clientID:         clientID,
		tenantID:         tenantID,
		credentialMethod: credentialMethod,
	}, nil
}

func runLogin(cmd *cobra.Command, _ []string) error {
	opts, err := newLoginOptions(cmd.Flags())
	if err != nil {
		return fmt.Errorf("failed to get login options: %w", err)
	}

	return login(cmd.Context(), opts)
}

func login(ctx context.Context, opts *loginOptions) error {
	clientID, err := discoverClientID(ctx, opts.clientID)
	if err != nil {
		return fmt.Errorf("failed to get client ID: %w", err)
	}

	cred, err := cmdhelper.NewUncachedCredential(opts.credentialMethod, opts.tenantID, clientID, []string{credentials.MicrosoftGraphScope})
	if err != nil {
		return fmt.Errorf("failed to create cached credential: %w", err)
	}

	_, err = cred.Authenticate(ctx, &policy.TokenRequestOptions{
		Scopes:    []string{credentials.MicrosoftGraphScope},
		EnableCAE: true,
		TenantID:  opts.tenantID,
	})
	if err != nil {
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	graphClient, err := graph.NewClient(cred, []string{credentials.MicrosoftGraphScope})
	if err != nil {
		return fmt.Errorf("failed to create pim client: %w", err)
	}

	me, err := graphClient.Me(ctx)
	if err != nil {
		return fmt.Errorf("failed to get me: %w", err)
	}

	fmt.Printf("Successfully logged in to as: %s\n", me.UserPrincipalName)

	return nil
}
