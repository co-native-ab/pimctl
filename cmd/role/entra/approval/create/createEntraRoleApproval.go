package create

import (
	"context"
	"fmt"

	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/graph"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Approve or Deny a pending Entra role assignment request",
	Long:  "Approve or Deny a pending Entra Role assignment request",
	RunE:  runCreateEntraRoleApproval,
}

func init() {
	Cmd.Flags().String("entra-role-name", "", "The name of the Entra role to grant or deny the assignment to")
	Cmd.MarkFlagRequired("entra-role-name")
	Cmd.Flags().String("justification", "", "The justification for the approval or denial")
	Cmd.MarkFlagRequired("justification")
	Cmd.Flags().String("user-principal-name", "", "The user principal name of the user to approve or deny the request for")
	Cmd.MarkFlagRequired("user-principal-name")

	reviewResult := graph.ApproveReviewResult
	Cmd.Flags().Var(&reviewResult, "review-result", reviewResult.HelpText())
	Cmd.RegisterFlagCompletionFunc("review-result", reviewResult.CobraCompletion)
}

type createEntraRoleApprovalOptions struct {
	entraRoleName     string
	justification     string
	userPrincipalName string
	reviewResult      graph.ReviewResult
}

func newCreateEntraRoleApprovalOptions(flags *pflag.FlagSet) (*createEntraRoleApprovalOptions, error) {
	entraRoleName, err := flags.GetString("entra-role-name")
	if err != nil {
		return nil, fmt.Errorf("failed to get entra-role-name: %w", err)
	}

	justification, err := flags.GetString("justification")
	if err != nil {
		return nil, fmt.Errorf("failed to get justification: %w", err)
	}

	userPrincipalName, err := flags.GetString("user-principal-name")
	if err != nil {
		return nil, fmt.Errorf("failed to get user-principal-name: %w", err)
	}

	reviewResult := graph.UnknownReviewResult
	reviewResultString := flags.Lookup("review-result").Value.String()
	err = reviewResult.Set(reviewResultString)
	if err != nil {
		return nil, fmt.Errorf("failed to set review-result: %w", err)
	}

	return &createEntraRoleApprovalOptions{
		entraRoleName:     entraRoleName,
		justification:     justification,
		userPrincipalName: userPrincipalName,
		reviewResult:      reviewResult,
	}, nil
}

func runCreateEntraRoleApproval(cmd *cobra.Command, _ []string) error {
	opts, err := newCreateEntraRoleApprovalOptions(cmd.Flags())
	if err != nil {
		return fmt.Errorf("failed to get options: %w", err)
	}

	return createEntraRoleApproval(cmd.Context(), opts)
}

func createEntraRoleApproval(ctx context.Context, opts *createEntraRoleApprovalOptions) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	entraRoleApprovalRequsts, err := graphClient.PIMEntraRoleApprovalRequests(ctx)
	if err != nil {
		return fmt.Errorf("failed to get entra role approval requests: %w", err)
	}

	matchingRequests := graph.EntraRoleAssignmentRequests{}
	for _, request := range entraRoleApprovalRequsts {
		if request.RoleDefinition.DisplayName == opts.entraRoleName && request.Principal.UserPrincipalName == opts.userPrincipalName {
			matchingRequests = append(matchingRequests, request)
		}
	}

	if len(matchingRequests) != 1 {
		return fmt.Errorf("found %d matching requests, expected 1", len(matchingRequests))
	}

	approvalRequest := matchingRequests[0]
	err = graphClient.PIMEntraRoleAssignmentApprovalByApprovalID(ctx, approvalRequest.ApprovalID, opts.justification, opts.reviewResult)
	if err != nil {
		return fmt.Errorf("failed to approve or deny entra role assignment request: %w", err)
	}

	fmt.Printf("Successfully set entra role assignment request to: %s\n", opts.reviewResult.String())

	return nil
}
