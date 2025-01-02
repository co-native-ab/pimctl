package login

import (
	"context"
	"fmt"

	"github.com/co-native-ab/pimctl/internal/azurerm"
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
	Cmd.Flags().String("tenant-id", "", "Azure AD Tenant ID for the current user")
	pimctlScope := credentials.MicrosoftGraphPimctlScope
	Cmd.Flags().Var(&pimctlScope, "scope", pimctlScope.HelpText())
	Cmd.RegisterFlagCompletionFunc("scope", pimctlScope.CobraCompletion)
	credentialMethod := credentials.InteractiveBrowserCredentialMethod
	Cmd.Flags().Var(&credentialMethod, "credential-method", credentialMethod.HelpText())
	Cmd.RegisterFlagCompletionFunc("credential-method", credentialMethod.CobraCompletion)
}

type loginOptions struct {
	clientID         string
	tenantID         string
	pimctlScope      credentials.PimctlScope
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

	pimctlScopeString := flags.Lookup("scope").Value.String()
	pimctlScope := credentials.UnknownPimctlScope
	err = pimctlScope.Set(pimctlScopeString)
	if err != nil {
		return nil, fmt.Errorf("failed to set scope: %w", err)
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
		pimctlScope:      pimctlScope,
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

	tenantID, err := discoverTenantID(ctx, opts.tenantID)
	if err != nil {
		return fmt.Errorf("failed to get tenant ID: %w", err)
	}

	cred, err := cmdhelper.NewUncachedCredential(opts.credentialMethod, tenantID, clientID, opts.pimctlScope.Scopes())
	if err != nil {
		return fmt.Errorf("failed to create cached credential: %w", err)
	}

	_, err = cred.Authenticate(ctx, &policy.TokenRequestOptions{
		Scopes:    opts.pimctlScope.Scopes(),
		EnableCAE: true,
		TenantID:  tenantID,
	})
	if err != nil {
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	switch opts.pimctlScope {
	case credentials.MicrosoftGraphPimctlScope:
		graphClient, err := graph.NewClient(cred, opts.pimctlScope.Scopes())
		if err != nil {
			return fmt.Errorf("failed to create pim client: %w", err)
		}

		me, err := graphClient.Me(ctx)
		if err != nil {
			return fmt.Errorf("failed to get me: %w", err)
		}

		fmt.Printf("Successfully logged in to Microsoft Graph as: %s\n", me.UserPrincipalName)
	case credentials.AzurePimctlScope:
		azurermClient, err := azurerm.NewClient(cred)
		if err != nil {
			return fmt.Errorf("failed to create azurerm client: %w", err)
		}

		upn, err := azurermClient.Me(ctx, opts.pimctlScope.Scopes(), tenantID)
		if err != nil {
			return fmt.Errorf("failed to get me: %w", err)
		}

		fmt.Printf("Successfully logged in to Azure Resource Manager as: %s\n", upn)
	default:
		return fmt.Errorf("unsupported scope: %s", opts.pimctlScope.String())
	}

	return nil
}
