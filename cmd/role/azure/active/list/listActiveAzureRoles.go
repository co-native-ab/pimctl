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
	Short: "List active Azure roles you have access to",
	Long:  "List all the active Azure roles you have access to in PIM",
	RunE:  runListActiveAzureRoles,
}

func init() {}

func runListActiveAzureRoles(cmd *cobra.Command, _ []string) error {
	return listActiveAzureRoles(cmd.Context())
}

func listActiveAzureRoles(ctx context.Context) error {
	azurermClient, err := cmdhelper.NewAzurermClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	azureRoleActiveAssignments, err := azurermClient.PIMAzureRoleActiveAssignments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get active azure roles: %w", err)
	}

	io := iostreams.System()
	err = printActiveAzureRolesList(io, azureRoleActiveAssignments)
	if err != nil {
		return fmt.Errorf("failed to print active azure roles: %w", err)
	}

	return nil
}

func printActiveAzureRolesList(io *iostreams.IOStreams, azureRoleActiveAssignments azurerm.AzureRoleActiveAssignments) error {
	headers := []string{
		"ROLE",
		"RESOURCE",
		"RESOURCE TYPE",
		"MEMBERSHIP",
		"CONDITION",
		"STATE",
		"END TIME",
		"SCOPE",
	}

	table := tableprinter.New(io, tableprinter.WithHeader(headers...))
	for _, azureRoleActiveAssignment := range azureRoleActiveAssignments {
		table.AddField(azureRoleActiveAssignment.Role())
		table.AddField(azureRoleActiveAssignment.Resource())
		table.AddField(azureRoleActiveAssignment.ResourceType())
		table.AddField(azureRoleActiveAssignment.Membership())
		table.AddField(azureRoleActiveAssignment.Condition())
		table.AddField(azureRoleActiveAssignment.State())
		table.AddField(azureRoleActiveAssignment.EndTime())
		table.AddField(azureRoleActiveAssignment.Scope())
		table.EndRow()
	}

	err := table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}
