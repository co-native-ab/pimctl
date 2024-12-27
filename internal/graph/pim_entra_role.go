package graph

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/microsoftgraph/msgraph-sdk-go/rolemanagement"
)

type RoleDefinition struct {
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
	IsBuiltIn   bool   `json:"isBuiltIn"`
	IsEnabled   bool   `json:"isEnabled"`
}

type EntraRoleEligibleAssignment struct {
	ID               string         `json:"id"`
	DirectoryScopeID string         `json:"directoryScopeId"`
	MemberType       string         `json:"memberType"`
	PrincipalID      string         `json:"principalId"`
	RoleDefinitionID string         `json:"roleDefinitionId"`
	Status           string         `json:"status"`
	ScheduleInfo     ScheduleInfo   `json:"scheduleInfo"`
	Principal        User           `json:"principal"`
	RoleDefinition   RoleDefinition `json:"roleDefinition"`
}

func (g EntraRoleEligibleAssignment) Scope() string {
	if g.DirectoryScopeID == "/" {
		return "Directory"
	}
	return g.DirectoryScopeID
}

type EntraRoleEligibleAssignments []EntraRoleEligibleAssignment

func (c *Client) PIMEntraRoleEligibleAssignments(ctx context.Context) (EntraRoleEligibleAssignments, error) {
	roleEligibilitySchedulesResponse, err := c.client.RoleManagement().Directory().RoleEligibilitySchedules().FilterByCurrentUserWithOn(to.Ptr("principal")).Get(ctx, &rolemanagement.DirectoryRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration{
		QueryParameters: &rolemanagement.DirectoryRoleEligibilitySchedulesFilterByCurrentUserWithOnRequestBuilderGetQueryParameters{
			Expand: []string{"roleDefinition", "principal"},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get eligibility schedules: %w", err)
	}

	graphValues := roleEligibilitySchedulesResponse.GetValue()
	if len(graphValues) == 0 {
		return nil, nil
	}

	var entraRoleEligibleAssignments EntraRoleEligibleAssignments
	for _, graphValue := range graphValues {
		value := EntraRoleEligibleAssignment{}
		err := unmarshalGraphValue(graphValue, &value)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal graph value: %w", err)
		}

		entraRoleEligibleAssignments = append(entraRoleEligibleAssignments, value)
	}

	return entraRoleEligibleAssignments, nil
}

// type EntraRoleActiveAssignment struct {
// 	ID                   string    `json:"id"`
// 	AcessID              string    `json:"accessId"`
// 	AssignmentScheduleId string    `json:"assignmentScheduleId"`
// 	AssignmentType       string    `json:"assignmentType"`
// 	EndDateTime          time.Time `json:"endDateTime"`
// 	GroupID              string    `json:"groupId"`
// 	PricipalID           string    `json:"principalId"`
// 	MemberType           string    `json:"memberType"`
// 	StartDateTime        time.Time `json:"startDateTime"`
// 	Group                Group     `json:"group"`
// 	Principal            User      `json:"principal"`
// }

// func (g EntraRoleActiveAssignment) EndTime() string {
// 	return g.EndDateTime.Local().Format(time.RFC3339)
// }

// type EntraRoleActiveAssignments []EntraRoleActiveAssignment

// func (c *Client) PIMEntraRoleActiveAssignments(ctx context.Context) (EntraRoleActiveAssignments, error) {
// 	assignmentScheduleInstancesResponse, err := c.client.IdentityGovernance().PrivilegedAccess().Group().AssignmentScheduleInstances().FilterByCurrentUserWithOn(to.Ptr("principal")).Get(ctx, &identitygovernance.PrivilegedAccessGroupAssignmentScheduleInstancesFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration{
// 		QueryParameters: &identitygovernance.PrivilegedAccessGroupAssignmentScheduleInstancesFilterByCurrentUserWithOnRequestBuilderGetQueryParameters{
// 			Expand: []string{"group", "principal"},
// 		},
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get assignment schedule instances: %w", err)
// 	}

// 	graphValues := assignmentScheduleInstancesResponse.GetValue()
// 	if len(graphValues) == 0 {
// 		return nil, nil
// 	}

// 	var groupActiveAssignments EntraRoleActiveAssignments
// 	for _, graphValue := range graphValues {
// 		value := EntraRoleActiveAssignment{}
// 		err := unmarshalGraphValue(graphValue, &value)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to unmarshal graph value: %w", err)
// 		}

// 		groupActiveAssignments = append(groupActiveAssignments, value)
// 	}

// 	return groupActiveAssignments, nil
// }

// type EntraRoleAssignmentRequest struct {
// 	ID                string    `json:"id"`
// 	AccessID          string    `json:"accessId"`
// 	Action            string    `json:"action"`
// 	ApprovalID        string    `json:"approvalId"`
// 	CompletedDateTime time.Time `json:"completedDateTime"`
// 	CreatedBy         struct {
// 		User struct {
// 			ID string `json:"id"`
// 		} `json:"user"`
// 	} `json:"createdBy"`
// 	CreatedDateTime  time.Time    `json:"createdDateTime"`
// 	Group            Group        `json:"group"`
// 	GroupID          string       `json:"groupId"`
// 	IsValidationOnly bool         `json:"isValidationOnly"`
// 	Justification    string       `json:"justification"`
// 	Principal        User         `json:"principal"`
// 	PrincipalID      string       `json:"principalId"`
// 	ScheduleInfo     ScheduleInfo `json:"scheduleInfo"`
// 	Status           string       `json:"status"`
// 	TargetSchduleID  string       `json:"targetScheduleId"`
// 	TicketInfo       struct {
// 		TicketNumber string `json:"ticketNumber"`
// 		TicketSystem string `json:"ticketSystem"`
// 	} `json:"ticketInfo"`
// }

// func (g EntraRoleAssignmentRequest) RequestTime() string {
// 	return g.CreatedDateTime.Local().Format(time.RFC3339)
// }

// type EntraRoleAssignmentRequests []EntraRoleAssignmentRequest

// func (c *Client) PIMEntraRoleAssignmentRequests(ctx context.Context) (EntraRoleAssignmentRequests, error) {
// 	return c.pimEntraRoleAssignmentRequests(ctx, "principal", "status eq 'PendingApproval'")
// }

// func (c *Client) pimEntraRoleAssignmentRequests(ctx context.Context, filterOn string, queryFilter string) (EntraRoleAssignmentRequests, error) {
// 	assignmentScheduleRequestsResponse, err := c.client.IdentityGovernance().PrivilegedAccess().Group().AssignmentScheduleRequests().FilterByCurrentUserWithOn(to.Ptr(filterOn)).Get(ctx, &identitygovernance.PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration{
// 		QueryParameters: &identitygovernance.PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetQueryParameters{
// 			Expand: []string{"group", "principal"},
// 			Filter: to.Ptr(queryFilter),
// 		},
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get assignment schedule requests: %w", err)
// 	}

// 	graphValues := assignmentScheduleRequestsResponse.GetValue()
// 	if len(graphValues) == 0 {
// 		return nil, nil
// 	}

// 	var groupAssignmentRequests EntraRoleAssignmentRequests
// 	for _, graphValue := range graphValues {
// 		value := EntraRoleAssignmentRequest{}
// 		err := unmarshalGraphValue(graphValue, &value)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to unmarshal graph value: %w", err)
// 		}

// 		groupAssignmentRequests = append(groupAssignmentRequests, value)
// 	}

// 	return groupAssignmentRequests, nil
// }

// func (c *Client) PIMEntraRoleApprovalRequests(ctx context.Context) (EntraRoleAssignmentRequests, error) {
// 	return c.pimEntraRoleAssignmentRequests(ctx, "approver", "status eq 'PendingApproval'")
// }

// func (c *Client) PIMEntraRoleAssignmentScheduleRequest(ctx context.Context, principalID string, groupID string, justification string, startDateTime time.Time, durationHours string) (string, error) {
// 	requestBody := graphmodels.NewPrivilegedAccessGroupAssignmentScheduleRequest()
// 	accessId := graphmodels.MEMBER_PRIVILEGEDACCESSGROUPRELATIONSHIPS
// 	requestBody.SetAccessId(&accessId)
// 	requestBody.SetPrincipalId(to.Ptr(principalID))
// 	requestBody.SetGroupId(to.Ptr(groupID))
// 	action := graphmodels.SELFACTIVATE_SCHEDULEREQUESTACTIONS
// 	requestBody.SetAction(&action)
// 	scheduleInfo := graphmodels.NewRequestSchedule()
// 	scheduleInfo.SetStartDateTime(to.Ptr(startDateTime))
// 	expiration := graphmodels.NewExpirationPattern()
// 	expirationType := graphmodels.AFTERDURATION_EXPIRATIONPATTERNTYPE
// 	expiration.SetTypeEscaped(&expirationType)
// 	duration, err := serialization.ParseISODuration(durationHours)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to parse ISO duration %q: %w", durationHours, err)
// 	}
// 	expiration.SetDuration(duration)
// 	scheduleInfo.SetExpiration(expiration)
// 	requestBody.SetScheduleInfo(scheduleInfo)
// 	requestBody.SetJustification(to.Ptr(justification))

// 	assignmentScheduleRequestsResponse, err := c.client.IdentityGovernance().PrivilegedAccess().Group().AssignmentScheduleRequests().Post(ctx, requestBody, &identitygovernance.PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderPostRequestConfiguration{})
// 	if err != nil {
// 		return "", fmt.Errorf("failed to post assignment schedule request: %w", err)
// 	}

// 	status := assignmentScheduleRequestsResponse.GetStatus()
// 	if status == nil {
// 		return "", fmt.Errorf("response status is nil")
// 	}

// 	return *status, nil
// }

// func (c *Client) PIMEntraRoleGetMaximumExpirationByGroupID(ctx context.Context, groupID string) (string, error) {
// 	roleManagementPolicyAssignmentsResponse, err := c.client.Policies().RoleManagementPolicyAssignments().Get(context.Background(), &policies.RoleManagementPolicyAssignmentsRequestBuilderGetRequestConfiguration{
// 		QueryParameters: &policies.RoleManagementPolicyAssignmentsRequestBuilderGetQueryParameters{
// 			Filter: to.Ptr(fmt.Sprintf("scopeId eq '%s' and scopeType eq 'Group' and roleDefinitionId eq 'member'", groupID)),
// 			Expand: []string{"policy($expand=rules)"},
// 		},
// 	})

// 	if err != nil {
// 		return "", fmt.Errorf("failed to get role management policy assignments: %w", err)
// 	}

// 	values := roleManagementPolicyAssignmentsResponse.GetValue()
// 	if len(values) != 1 {
// 		return "", fmt.Errorf("expected 1 role management policy assignment, got %d", len(values))
// 	}

// 	policy := values[0].GetPolicy()
// 	if policy == nil {
// 		return "", fmt.Errorf("policy is nil")
// 	}

// 	rules := policy.GetRules()
// 	if rules == nil {
// 		return "", fmt.Errorf("rules is nil")
// 	}

// 	if len(rules) == 0 {
// 		return "", fmt.Errorf("no rules found")
// 	}

// 	for _, rule := range rules {
// 		ruleID := rule.GetId()
// 		if ruleID == nil || *ruleID != "Expiration_EndUser_Assignment" {
// 			continue
// 		}

// 		value := UnifiedRoleManagementPolicyExpirationRule{}
// 		err := unmarshalGraphValue(rule, &value)
// 		if err != nil {
// 			return "", fmt.Errorf("failed to unmarshal graph value: %w", err)
// 		}

// 		isoDuration, err := serialization.ParseISODuration(value.MaximumDuration)
// 		if err != nil {
// 			return "", fmt.Errorf("failed to parse ISO duration %q: %w", value.MaximumDuration, err)
// 		}

// 		return isoDuration.String(), nil
// 	}

// 	return "", fmt.Errorf("no expiration rule found")
// }

// func (c *Client) PIMEntraRoleAssignmentApprovalByApprovalID(ctx context.Context, approvalID string, justification string, reviewResult ReviewResult) error {
// 	assignmentApprovalResponse, err := c.client.IdentityGovernance().PrivilegedAccess().Group().AssignmentApprovals().ByApprovalId(approvalID).Get(ctx, &identitygovernance.PrivilegedAccessGroupAssignmentApprovalsApprovalItemRequestBuilderGetRequestConfiguration{})
// 	if err != nil {
// 		return fmt.Errorf("failed to get assignment approvals: %w", err)
// 	}

// 	value := AssignmentApproval{}
// 	err = unmarshalGraphValue(assignmentApprovalResponse, &value)
// 	if err != nil {
// 		return fmt.Errorf("failed to unmarshal graph value: %w", err)
// 	}

// 	if len(value.Stages) != 1 {
// 		return fmt.Errorf("expected 1 approval stage, got %d", len(value.Stages))
// 	}

// 	stage := value.Stages[0]
// 	if stage.ReviewResult != "NotReviewed" {
// 		return fmt.Errorf("expected review result to be NotReviewed, got %q", stage.ReviewResult)
// 	}

// 	if stage.Status != "InProgress" {
// 		return fmt.Errorf("expected status to be InProgress, got %q", stage.Status)
// 	}

// 	if !stage.AssignedToMe {
// 		return fmt.Errorf("expected assignedToMe to be true, got %v", stage.AssignedToMe)
// 	}

// 	if stage.ID == "" {
// 		return fmt.Errorf("expected ID to be non-empty")
// 	}

// 	approvalStageID := stage.ID

// 	requestBody := graphmodels.NewApprovalStage()
// 	requestBody.SetReviewResult(to.Ptr(reviewResult.String()))
// 	requestBody.SetJustification(&justification)

// 	_, err = c.client.IdentityGovernance().PrivilegedAccess().Group().AssignmentApprovals().ByApprovalId(approvalID).Stages().ByApprovalStageId(approvalStageID).Patch(ctx, requestBody, &identitygovernance.PrivilegedAccessGroupAssignmentApprovalsItemStagesApprovalStageItemRequestBuilderPatchRequestConfiguration{})
// 	if err != nil {
// 		return fmt.Errorf("failed to patch assignment approval stage: %w", err)
// 	}

// 	return nil
// }
