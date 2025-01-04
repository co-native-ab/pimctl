package list

import (
	"context"
	"fmt"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/co-native-ab/pimctl/internal/azurerm"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/tableprinter"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List Azure role PIM requests waiting for approval from current user",
	Long:  "List Azure role PIM request waiting for approval from current user",
	RunE:  runListAzureRoleApprovals,
}

func init() {}

func runListAzureRoleApprovals(cmd *cobra.Command, _ []string) error {
	return listAzureRoleApprovals(cmd.Context())
}

func listAzureRoleApprovals(ctx context.Context) error {
	azurermClient, err := cmdhelper.NewAzurermClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	azureRoleApprovalRequests, err := azurermClient.PIMAzureRoleApprovalRequests(ctx)
	if err != nil {
		return fmt.Errorf("failed to get azure role approval requests: %w", err)
	}

	io := iostreams.System()
	err = printAzureRoleApprovalRequestList(io, azureRoleApprovalRequests)
	if err != nil {
		return fmt.Errorf("failed to print approval requests: %w", err)
	}

	return nil
}

func printAzureRoleApprovalRequestList(io *iostreams.IOStreams, azureRoleApprovalRequests azurerm.AzureRoleAssignmentRequests) error {
	headers := []string{
		"ROLE",
		"RESOURCE",
		"RESOURCE TYPE",
		"REQUESTOR",
		"REQUEST TIME",
		"REASON",
		"CONDITION",
		"REQUEST TYPE",
		"START TIME",
		"END TIME",
		"SCOPE",
	}

	table := tableprinter.New(io, tableprinter.WithHeader(headers...))
	for _, azureRoleAssignmentRequest := range azureRoleApprovalRequests {
		table.AddField(azureRoleAssignmentRequest.Role())
		table.AddField(azureRoleAssignmentRequest.Resource())
		table.AddField(azureRoleAssignmentRequest.ResourceType())
		table.AddField(azureRoleAssignmentRequest.Requestor())
		table.AddField(azureRoleAssignmentRequest.RequestTime())
		table.AddField(azureRoleAssignmentRequest.Reason())
		table.AddField(azureRoleAssignmentRequest.Condition())
		table.AddField(azureRoleAssignmentRequest.RequestType())
		table.AddField(azureRoleAssignmentRequest.StartTime())
		table.AddField(azureRoleAssignmentRequest.EndTime())
		table.AddField(azureRoleAssignmentRequest.Scope())
		table.EndRow()
	}

	err := table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}
