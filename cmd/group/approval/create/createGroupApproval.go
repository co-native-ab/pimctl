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
	Short: "Approve or Deny a pending group assignment request",
	Long:  "Approve or Deny a pending group assignment request",
	RunE:  runCreateGroupApproval,
}

func init() {
	Cmd.Flags().String("group-name", "", "The name of the group to grant or deny the assignment to")
	Cmd.MarkFlagRequired("group-name")
	Cmd.Flags().String("justification", "", "The justification for the approval or denial")
	Cmd.MarkFlagRequired("justification")
	Cmd.Flags().String("user-principal-name", "", "The user principal name of the user to approve or deny the request for")
	Cmd.MarkFlagRequired("user-principal-name")

	reviewResult := graph.ApproveReviewResult
	Cmd.Flags().Var(&reviewResult, "review-result", reviewResult.HelpText())
	Cmd.RegisterFlagCompletionFunc("review-result", reviewResult.CobraCompletion)

}

type createGroupApprovalOptions struct {
	groupName         string
	justification     string
	userPrincipalName string
	reviewResult      graph.ReviewResult
}

func newCreateGroupApprovalOptions(flags *pflag.FlagSet) (*createGroupApprovalOptions, error) {
	groupName, err := flags.GetString("group-name")
	if err != nil {
		return nil, fmt.Errorf("failed to get group-name: %w", err)
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

	return &createGroupApprovalOptions{
		groupName:         groupName,
		justification:     justification,
		userPrincipalName: userPrincipalName,
		reviewResult:      reviewResult,
	}, nil
}

func runCreateGroupApproval(cmd *cobra.Command, _ []string) error {
	opts, err := newCreateGroupApprovalOptions(cmd.Flags())
	if err != nil {
		return fmt.Errorf("failed to get options: %w", err)
	}

	return createGroupApproval(cmd.Context(), opts)
}

func createGroupApproval(ctx context.Context, opts *createGroupApprovalOptions) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	approvalRequsts, err := graphClient.PIMGroupApprovalRequests(ctx)
	if err != nil {
		return fmt.Errorf("failed to get group assignment requests: %w", err)
	}

	matchingRequests := graph.GroupAssignmentRequests{}
	for _, request := range approvalRequsts {
		if request.Group.DisplayName == opts.groupName && request.Principal.UserPrincipalName == opts.userPrincipalName {
			matchingRequests = append(matchingRequests, request)
		}
	}

	if len(matchingRequests) != 1 {
		return fmt.Errorf("found %d matching requests, expected 1", len(matchingRequests))
	}

	approvalRequest := matchingRequests[0]

	err = graphClient.PIMGroupAssignmentApprovalByApprovalID(ctx, approvalRequest.ApprovalID, opts.justification, opts.reviewResult)
	if err != nil {
		return fmt.Errorf("failed to approve or deny group assignment request: %w", err)
	}

	fmt.Printf("Successfully set group assignment request to: %s\n", opts.reviewResult.String())

	return nil
}
