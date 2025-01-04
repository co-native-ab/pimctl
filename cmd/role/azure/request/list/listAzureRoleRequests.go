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
	Short: "List all Azure role PIM requests current user can see",
	Long:  "List all Azure role PIM requests current user can see",
	RunE:  runListAzureRoleRequests,
}

func init() {}

func runListAzureRoleRequests(cmd *cobra.Command, _ []string) error {
	return listAzureRoleRequests(cmd.Context())
}

func listAzureRoleRequests(ctx context.Context) error {
	azurermClient, err := cmdhelper.NewAzurermClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	azureRoleAssignmentRequests, err := azurermClient.PIMAzureRoleAssignmentRequests(ctx)
	if err != nil {
		return fmt.Errorf("failed to get azure role assignment requests: %w", err)
	}

	io := iostreams.System()
	err = printAzureRoleRequestList(io, azureRoleAssignmentRequests)
	if err != nil {
		return fmt.Errorf("failed to print azure role assignment requests: %w", err)
	}

	return nil
}

func printAzureRoleRequestList(io *iostreams.IOStreams, azureRoleAssignmentRequests azurerm.AzureRoleAssignmentRequests) error {
	headers := []string{
		"ROLE",
		"RESOURCE",
		"REQUEST TYPE",
		"REASON",
		"CONDITION",
		"REQUEST TIME",
		"START TIME",
		"REQUEST STATUS",
	}

	table := tableprinter.New(io, tableprinter.WithHeader(headers...))
	for _, azureRoleAssignmentRequest := range azureRoleAssignmentRequests {
		table.AddField(azureRoleAssignmentRequest.Role())
		table.AddField(azureRoleAssignmentRequest.Resource())
		table.AddField(azureRoleAssignmentRequest.RequestType())
		table.AddField(azureRoleAssignmentRequest.Reason())
		table.AddField(azureRoleAssignmentRequest.Condition())
		table.AddField(azureRoleAssignmentRequest.RequestTime())
		table.AddField(azureRoleAssignmentRequest.StartTime())
		table.AddField(azureRoleAssignmentRequest.RequestStatus())
		table.EndRow()
	}

	err := table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}
