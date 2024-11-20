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
	Short: "Create a new group assignment request",
	Long:  "Create a new group assignment request",
	RunE:  runCreateGroupRequest,
}

func init() {
	Cmd.Flags().String("group-name", "", "The name of the group to request assignment for (can't be used with --group-id and can't have two groups with the same name)")
	Cmd.Flags().String("group-id", "", "The id of the group to request assignment for (can't be used with --group-name)")
	Cmd.MarkFlagsOneRequired("group-name", "group-id")
	Cmd.MarkFlagsMutuallyExclusive("group-name", "group-id")
	Cmd.Flags().String("justification", "", "The justification for the request")
	Cmd.MarkFlagRequired("justification")
	Cmd.Flags().Int("duration", 0, "The duration of the request in hours (default: configured maximum duration)")
}

type createGroupRequestOptions struct {
	groupName     string
	groupID       string
	justification string
	duration      int
}

func newCreateGroupRequestOptions(flags *pflag.FlagSet) (*createGroupRequestOptions, error) {
	groupName, err := flags.GetString("group-name")
	if err != nil {
		return nil, fmt.Errorf("failed to get group-name: %w", err)
	}

	groupID, err := flags.GetString("group-id")
	if err != nil {
		return nil, fmt.Errorf("failed to get group-id: %w", err)
	}

	justification, err := flags.GetString("justification")
	if err != nil {
		return nil, fmt.Errorf("failed to get justification: %w", err)
	}

	duration, err := flags.GetInt("duration")
	if err != nil {
		return nil, fmt.Errorf("failed to get duration: %w", err)
	}

	return &createGroupRequestOptions{
		groupName:     groupName,
		groupID:       groupID,
		justification: justification,
		duration:      duration,
	}, nil
}

func runCreateGroupRequest(cmd *cobra.Command, _ []string) error {
	opts, err := newCreateGroupRequestOptions(cmd.Flags())
	if err != nil {
		return fmt.Errorf("failed to get options: %w", err)
	}

	return createGroupRequest(cmd.Context(), opts)
}

func createGroupRequest(ctx context.Context, opts *createGroupRequestOptions) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	groupID, err := getGroupID(ctx, graphClient, opts.groupID, opts.groupName)
	if err != nil {
		return fmt.Errorf("failed to get group id: %w", err)
	}

	status, err := cmdhelper.PIMGroupAssignmentScheduleRequest(ctx, graphClient, opts.duration, groupID, opts.justification)
	if err != nil {
		return fmt.Errorf("failed to get group assignment requests: %w", err)
	}

	fmt.Printf("Request status: %s\n", status)

	return nil
}

func getGroupID(ctx context.Context, graphClient *graph.Client, flagGroupID string, flagGroupName string) (string, error) {
	if flagGroupID != "" {
		return flagGroupID, nil
	}

	if flagGroupName == "" {
		return "", fmt.Errorf("--group-name not set")
	}

	groupEligibleAssignments, err := graphClient.PIMGroupEligibleAssignments(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get group eligible assignments: %w", err)
	}

	groupIDs := []string{}
	for _, groupAssignment := range groupEligibleAssignments {
		if groupAssignment.Group.DisplayName == flagGroupName {
			groupIDs = append(groupIDs, groupAssignment.GroupID)
		}
	}

	if len(groupIDs) != 1 {
		return "", fmt.Errorf("expected exactly one group with name %s, got %d", flagGroupName, len(groupIDs))
	}

	return groupIDs[0], nil
}
