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
	Short: "Approve or Deny a pending Azure role assignment request",
	Long:  "Approve or Deny a pending Azure Role assignment request",
	RunE:  runCreateAzureRoleApproval,
}

func init() {
	Cmd.Flags().String("azure-role-name", "", "The name of the Azure role to grant or deny the assignment to")
	Cmd.MarkFlagRequired("azure-role-name")
	Cmd.Flags().String("justification", "", "The justification for the approval or denial")
	Cmd.MarkFlagRequired("justification")
	Cmd.Flags().String("user-principal-name", "", "The user principal name of the user to approve or deny the request for")
	Cmd.MarkFlagRequired("user-principal-name")

	reviewResult := azurerm.ApproveReviewResult
	Cmd.Flags().Var(&reviewResult, "review-result", reviewResult.HelpText())
	Cmd.RegisterFlagCompletionFunc("review-result", reviewResult.CobraCompletion)
}

type createAzureRoleApprovalOptions struct {
	azureRoleName     string
	justification     string
	userPrincipalName string
	reviewResult      azurerm.ReviewResult
}

func newCreateAzureRoleApprovalOptions(flags *pflag.FlagSet) (*createAzureRoleApprovalOptions, error) {
	azureRoleName, err := flags.GetString("azure-role-name")
	if err != nil {
		return nil, fmt.Errorf("failed to get azure-role-name: %w", err)
	}

	justification, err := flags.GetString("justification")
	if err != nil {
		return nil, fmt.Errorf("failed to get justification: %w", err)
	}

	userPrincipalName, err := flags.GetString("user-principal-name")
	if err != nil {
		return nil, fmt.Errorf("failed to get user-principal-name: %w", err)
	}

	reviewResult := azurerm.UnknownReviewResult
	reviewResultString := flags.Lookup("review-result").Value.String()
	err = reviewResult.Set(reviewResultString)
	if err != nil {
		return nil, fmt.Errorf("failed to set review-result: %w", err)
	}

	return &createAzureRoleApprovalOptions{
		azureRoleName:     azureRoleName,
		justification:     justification,
		userPrincipalName: userPrincipalName,
		reviewResult:      reviewResult,
	}, nil
}

func runCreateAzureRoleApproval(cmd *cobra.Command, _ []string) error {
	opts, err := newCreateAzureRoleApprovalOptions(cmd.Flags())
	if err != nil {
		return fmt.Errorf("failed to get options: %w", err)
	}

	return createAzureRoleApproval(cmd.Context(), opts)
}

func createAzureRoleApproval(ctx context.Context, opts *createAzureRoleApprovalOptions) error {
	azurermClient, err := cmdhelper.NewAzurermClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	azureRoleApprovalRequsts, err := azurermClient.PIMAzureRoleApprovalRequests(ctx)
	if err != nil {
		return fmt.Errorf("failed to get azure role approval requests: %w", err)
	}

	matchingRequests := azurerm.AzureRoleAssignmentRequests{}
	for _, azureRoleApprovalRequst := range azureRoleApprovalRequsts {
		if azureRoleApprovalRequst.Role() == opts.azureRoleName && azureRoleApprovalRequst.PrincipalEmail() == opts.userPrincipalName {
			matchingRequests = append(matchingRequests, azureRoleApprovalRequst)
		}
	}

	if len(matchingRequests) != 1 {
		return fmt.Errorf("found %d matching requests, expected 1", len(matchingRequests))
	}

	azureRoleApprovalRequst := matchingRequests[0]
	err = azurermClient.PIMAzureRoleAssignmentApprovalByApprovalID(ctx, *azureRoleApprovalRequst.Properties.ApprovalID, opts.justification, opts.reviewResult)
	if err != nil {
		return fmt.Errorf("failed to approve or deny azure role assignment request: %w", err)
	}

	fmt.Printf("Successfully set azure role assignment request to: %s\n", opts.reviewResult.String())

	return nil
}
