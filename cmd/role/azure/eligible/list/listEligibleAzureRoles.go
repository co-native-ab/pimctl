package list

import (
	"context"
	"fmt"

	"github.com/co-native-ab/pimctl/internal/cmdhelper"

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

	err = azurermClient.PIMAzureRoleEligibleAssignments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get eligible entra roles: %w", err)
	}

	return nil

	// io := iostreams.System()
	// err = printEligibleGroupsList(io, azureRoleEligibleAssignments)
	// if err != nil {
	// 	return fmt.Errorf("failed to print eligible groups: %w", err)
	// }

	// return nil
}

// func printEligibleGroupsList(io *iostreams.IOStreams, azureRoleEligibleAssignments graph.AzureRoleEligibleAssignments) error {
// 	headers := []string{
// 		"ROLE",
// 		"SCOPE",
// 		"MEMBERSHIP",
// 		"END TIME",
// 	}

// 	table := tableprinter.New(io, tableprinter.WithHeader(headers...))
// 	for _, azureRoleEligibleAssignment := range azureRoleEligibleAssignments {
// 		table.AddField(azureRoleEligibleAssignment.RoleDefinition.DisplayName)
// 		table.AddField(azureRoleEligibleAssignment.Scope())
// 		table.AddField(azureRoleEligibleAssignment.MemberType)
// 		table.AddField(azureRoleEligibleAssignment.ScheduleInfo.EndTime())
// 		table.EndRow()
// 	}

// 	err := table.Render()
// 	if err != nil {
// 		return fmt.Errorf("failed to render table: %w", err)
// 	}

// 	return nil
// }
