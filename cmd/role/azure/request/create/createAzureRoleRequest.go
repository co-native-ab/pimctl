package create

import (
	"context"
	"fmt"

	"github.com/co-native-ab/pimctl/internal/azurerm"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Azure role assignment request",
	Long:  "Create a new Azure role assignment request",
	RunE:  runCreateAzureRoleRequest,
}

func init() {
	Cmd.Flags().String("scope", "", "The scope to request assignment for")
	Cmd.MarkFlagRequired("scope")
	Cmd.Flags().String("azure-role-name", "", "The name of the Azure role to request assignment for (can't be used with --azure-role-definition-id and can't have two roles with the same name)")
	Cmd.Flags().String("azure-role-definition-id", "", "The id of the Azure role to request assignment for (can't be used with --azure-role-name)")
	Cmd.MarkFlagsOneRequired("azure-role-name", "azure-role-definition-id")
	Cmd.MarkFlagsMutuallyExclusive("azure-role-name", "azure-role-definition-id")
	Cmd.Flags().String("justification", "", "The justification for the request")
	Cmd.MarkFlagRequired("justification")
	Cmd.Flags().Int("duration", 0, "The duration of the request in hours (default: configured maximum duration)")
}

type createAzureRoleRequestOptions struct {
	azureRoleName         string
	azureRoleDefinitionID string
	azureRoleScope        string
	justification         string
	duration              int
}

func newCreateAzureRoleRequestOptions(flags *pflag.FlagSet) (*createAzureRoleRequestOptions, error) {
	azureRoleName, err := flags.GetString("azure-role-name")
	if err != nil {
		return nil, fmt.Errorf("failed to get azure-role-name: %w", err)
	}

	azureRoleDefinitionID, err := flags.GetString("azure-role-definition-id")
	if err != nil {
		return nil, fmt.Errorf("failed to get azure-role-definition-id: %w", err)
	}

	azureRoleScope, err := flags.GetString("scope")
	if err != nil {
		return nil, fmt.Errorf("failed to get scope: %w", err)
	}

	justification, err := flags.GetString("justification")
	if err != nil {
		return nil, fmt.Errorf("failed to get justification: %w", err)
	}

	duration, err := flags.GetInt("duration")
	if err != nil {
		return nil, fmt.Errorf("failed to get duration: %w", err)
	}

	return &createAzureRoleRequestOptions{
		azureRoleName:         azureRoleName,
		azureRoleDefinitionID: azureRoleDefinitionID,
		azureRoleScope:        azureRoleScope,
		justification:         justification,
		duration:              duration,
	}, nil
}

func runCreateAzureRoleRequest(cmd *cobra.Command, _ []string) error {
	opts, err := newCreateAzureRoleRequestOptions(cmd.Flags())
	if err != nil {
		return fmt.Errorf("failed to get options: %w", err)
	}

	return createAzureRoleRequest(cmd.Context(), opts)
}

func createAzureRoleRequest(ctx context.Context, opts *createAzureRoleRequestOptions) error {
	azurermClient, err := cmdhelper.NewAzurermClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	matchingAzureRoleEligibleAssignment, err := getMatchingAzureRoleEligibleAssignment(ctx, azurermClient, opts.azureRoleDefinitionID, opts.azureRoleName, opts.azureRoleScope)
	if err != nil {
		return fmt.Errorf("failed to get azure role id: %w", err)
	}

	status, err := cmdhelper.PIMAzureRoleAssignmentScheduleRequest(ctx, azurermClient, matchingAzureRoleEligibleAssignment, opts.duration, opts.justification)
	if err != nil {
		return fmt.Errorf("failed to create azure role assignment requests: %w", err)
	}

	fmt.Printf("Request status: %s\n", status)

	return nil
}

func getMatchingAzureRoleEligibleAssignment(ctx context.Context, azurermClient *azurerm.Client, flagAzureRoleDefinitionID string, flagAzureRoleName string, scope string) (azurerm.AzureRoleEligibleAssignment, error) {
	if flagAzureRoleName == "" {
		return azurerm.AzureRoleEligibleAssignment{}, fmt.Errorf("--azure-role-name not set")
	}

	azureRoleEligibleAssignments, err := azurermClient.PIMAzureRoleEligibleAssignments(ctx)
	if err != nil {
		return azurerm.AzureRoleEligibleAssignment{}, fmt.Errorf("failed to get azure role eligible assignments: %w", err)
	}

	matchingAzureRoleEligibleAssignments := azurerm.AzureRoleEligibleAssignments{}
	for _, azureRoleAssignment := range azureRoleEligibleAssignments {
		if flagAzureRoleDefinitionID != "" && azureRoleAssignment.RoleDefinitionID() == flagAzureRoleDefinitionID {
			if azureRoleAssignment.Scope() == scope {
				matchingAzureRoleEligibleAssignments = append(matchingAzureRoleEligibleAssignments, azureRoleAssignment)
			}
		}
		if flagAzureRoleName != "" && azureRoleAssignment.Role() == flagAzureRoleName {
			if azureRoleAssignment.Scope() == scope {
				matchingAzureRoleEligibleAssignments = append(matchingAzureRoleEligibleAssignments, azureRoleAssignment)
			}
		}
	}

	if len(matchingAzureRoleEligibleAssignments) != 1 {
		return azurerm.AzureRoleEligibleAssignment{}, fmt.Errorf("expected exactly one azure role with name %s, got %d", flagAzureRoleName, len(matchingAzureRoleEligibleAssignments))
	}

	return matchingAzureRoleEligibleAssignments[0], nil
}
