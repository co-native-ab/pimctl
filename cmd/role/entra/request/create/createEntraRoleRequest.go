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
	Short: "Create a new Entra role assignment request",
	Long:  "Create a new Entra role assignment request",
	RunE:  runCreateEntraRoleRequest,
}

func init() {
	Cmd.Flags().String("entra-role-name", "", "The name of the Entra role to request assignment for (can't be used with --entra-role-id and can't have two roles with the same name)")
	Cmd.Flags().String("entra-role-id", "", "The id of the Entra role to request assignment for (can't be used with --entra-role-name)")
	Cmd.MarkFlagsOneRequired("entra-role-name", "entra-role-id")
	Cmd.MarkFlagsMutuallyExclusive("entra-role-name", "entra-role-id")
	Cmd.Flags().String("justification", "", "The justification for the request")
	Cmd.MarkFlagRequired("justification")
	Cmd.Flags().Int("duration", 0, "The duration of the request in hours (default: configured maximum duration)")
}

type createEntraRoleRequestOptions struct {
	entraRoleName string
	entraRoleID   string
	justification string
	duration      int
}

func newCreateEntraRoleRequestOptions(flags *pflag.FlagSet) (*createEntraRoleRequestOptions, error) {
	entraRoleName, err := flags.GetString("entra-role-name")
	if err != nil {
		return nil, fmt.Errorf("failed to get entra-role-name: %w", err)
	}

	entraRoleID, err := flags.GetString("entra-role-id")
	if err != nil {
		return nil, fmt.Errorf("failed to get entra-role-id: %w", err)
	}

	justification, err := flags.GetString("justification")
	if err != nil {
		return nil, fmt.Errorf("failed to get justification: %w", err)
	}

	duration, err := flags.GetInt("duration")
	if err != nil {
		return nil, fmt.Errorf("failed to get duration: %w", err)
	}

	return &createEntraRoleRequestOptions{
		entraRoleName: entraRoleName,
		entraRoleID:   entraRoleID,
		justification: justification,
		duration:      duration,
	}, nil
}

func runCreateEntraRoleRequest(cmd *cobra.Command, _ []string) error {
	opts, err := newCreateEntraRoleRequestOptions(cmd.Flags())
	if err != nil {
		return fmt.Errorf("failed to get options: %w", err)
	}

	return createEntraRoleRequest(cmd.Context(), opts)
}

func createEntraRoleRequest(ctx context.Context, opts *createEntraRoleRequestOptions) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	entraRoleID, err := getEntraRoleID(ctx, graphClient, opts.entraRoleID, opts.entraRoleName)
	if err != nil {
		return fmt.Errorf("failed to get entra role id: %w", err)
	}

	status, err := cmdhelper.PIMEntraRoleAssignmentScheduleRequest(ctx, graphClient, opts.duration, entraRoleID, opts.justification)
	if err != nil {
		return fmt.Errorf("failed to create entra role assignment requests: %w", err)
	}

	fmt.Printf("Request status: %s\n", status)

	return nil
}

func getEntraRoleID(ctx context.Context, graphClient *graph.Client, flagEntraRoleID string, flagEntraRoleName string) (string, error) {
	if flagEntraRoleID != "" {
		return flagEntraRoleID, nil
	}

	if flagEntraRoleName == "" {
		return "", fmt.Errorf("--entra-role-name not set")
	}

	entraRoleEligibleAssignments, err := graphClient.PIMEntraRoleEligibleAssignments(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get entra role eligible assignments: %w", err)
	}

	entraRoleIDs := []string{}
	for _, entraRoleAssignment := range entraRoleEligibleAssignments {
		if entraRoleAssignment.RoleDefinition.DisplayName == flagEntraRoleName {
			entraRoleIDs = append(entraRoleIDs, entraRoleAssignment.RoleDefinitionID)
		}
	}

	if len(entraRoleIDs) != 1 {
		return "", fmt.Errorf("expected exactly one entra role with name %s, got %d", flagEntraRoleName, len(entraRoleIDs))
	}

	return entraRoleIDs[0], nil
}
