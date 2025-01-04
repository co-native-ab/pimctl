package list

import (
	"context"
	"fmt"

	"github.com/co-native-ab/pimctl/internal/azurerm"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/tableprinter"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List eligible Azure roles you have access to",
	Long:  "List all the eligible Azure roles you have access to in PIM",
	RunE:  runListEligibleAzureRoles,
}

func init() {}

func runListEligibleAzureRoles(cmd *cobra.Command, _ []string) error {
	return listEligibleAzureRoles(cmd.Context())
}

func listEligibleAzureRoles(ctx context.Context) error {
	azurermClient, err := cmdhelper.NewAzurermClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	azureRoleEligibleAssignments, err := azurermClient.PIMAzureRoleEligibleAssignments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get eligible entra roles: %w", err)
	}

	io := iostreams.System()
	err = printEligibleAzureRoleList(io, azureRoleEligibleAssignments)
	if err != nil {
		return fmt.Errorf("failed to print eligible azure roles: %w", err)
	}

	return nil
}

func printEligibleAzureRoleList(io *iostreams.IOStreams, azureRoleEligibleAssignments azurerm.AzureRoleEligibleAssignments) error {
	headers := []string{
		"ROLE",
		"RESOURCE",
		"RESOURCE TYPE",
		"MEMBERSHIP",
		"CONDITION",
		"END TIME",
		"SCOPE",
	}

	table := tableprinter.New(io, tableprinter.WithHeader(headers...))
	for _, azureRoleEligibleAssignment := range azureRoleEligibleAssignments {
		table.AddField(azureRoleEligibleAssignment.Role())
		table.AddField(azureRoleEligibleAssignment.Resource())
		table.AddField(azureRoleEligibleAssignment.ResourceType())
		table.AddField(azureRoleEligibleAssignment.Membership())
		table.AddField(azureRoleEligibleAssignment.Condition())
		table.AddField(azureRoleEligibleAssignment.EndTime())
		table.AddField(azureRoleEligibleAssignment.Scope())
		table.EndRow()
	}

	err := table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}
