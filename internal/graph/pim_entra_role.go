package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	graphmodelsbeta "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	rolemanagementbeta "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement"
	graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/policies"
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

type EntraRoleActiveAssignment struct {
	ID                       string         `json:"id"`
	AssignmentType           string         `json:"assignmentType"`
	DirectoryScopeID         string         `json:"directoryScopeId"`
	MemberType               string         `json:"memberType"`
	PrincipalID              string         `json:"principalId"`
	RoleDefinitionID         string         `json:"roleDefinitionId"`
	RoleAssignmentOriginID   string         `json:"roleAssignmentOriginId"`
	RoleAssignmentScheduleID string         `json:"roleAssignmentScheduleId"`
	Principal                User           `json:"principal"`
	RoleDefinition           RoleDefinition `json:"roleDefinition"`
	StartDateTime            time.Time      `json:"startDateTime"`
	EndDateTime              time.Time      `json:"endDateTime"`
}

func (g EntraRoleActiveAssignment) EndTime() string {
	return g.EndDateTime.Local().Format(time.RFC3339)
}

func (g EntraRoleActiveAssignment) Scope() string {
	if g.DirectoryScopeID == "/" {
		return "Directory"
	}
	return g.DirectoryScopeID
}

type EntraRoleActiveAssignments []EntraRoleActiveAssignment

func (c *Client) PIMEntraRoleActiveAssignments(ctx context.Context) (EntraRoleActiveAssignments, error) {
	roleAssignmentScheduleInstancesResponse, err := c.client.RoleManagement().Directory().RoleAssignmentScheduleInstances().FilterByCurrentUserWithOn(to.Ptr("principal")).Get(ctx, &rolemanagement.DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration{
		QueryParameters: &rolemanagement.DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnRequestBuilderGetQueryParameters{
			Expand: []string{"roleDefinition", "principal"},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get role assignment schedule instances: %w", err)
	}

	graphValues := roleAssignmentScheduleInstancesResponse.GetValue()
	if len(graphValues) == 0 {
		return nil, nil
	}

	var entraRoleActiveAssignments EntraRoleActiveAssignments
	for _, graphValue := range graphValues {
		value := EntraRoleActiveAssignment{}
		err := unmarshalGraphValue(graphValue, &value)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal graph value: %w", err)
		}

		entraRoleActiveAssignments = append(entraRoleActiveAssignments, value)
	}

	return entraRoleActiveAssignments, nil
}

type EntraRoleAssignmentRequest struct {
	ID                string    `json:"id"`
	Action            string    `json:"action"`
	ApprovalID        string    `json:"approvalId"`
	CompletedDateTime time.Time `json:"completedDateTime"`
	CreatedBy         struct {
		User struct {
			ID string `json:"id"`
		} `json:"user"`
	} `json:"createdBy"`
	CreatedDateTime  time.Time      `json:"createdDateTime"`
	DirectoryScopeID string         `json:"directoryScopeId"`
	IsValidationOnly bool           `json:"isValidationOnly"`
	Justification    string         `json:"justification"`
	Principal        User           `json:"principal"`
	PrincipalID      string         `json:"principalId"`
	RoleDefinition   RoleDefinition `json:"roleDefinition"`
	RoleDefinitionID string         `json:"roleDefinitionId"`
	ScheduleInfo     ScheduleInfo   `json:"scheduleInfo"`
	Status           string         `json:"status"`
	TargetScheduleID string         `json:"targetScheduleId"`
	TicketInfo       struct {
		TicketNumber string `json:"ticketNumber"`
		TicketSystem string `json:"ticketSystem"`
	} `json:"ticketInfo"`
}

func (g EntraRoleAssignmentRequest) RequestTime() string {
	return g.CreatedDateTime.Local().Format(time.RFC3339)
}

func (g EntraRoleAssignmentRequest) Scope() string {
	if g.DirectoryScopeID == "/" {
		return "Directory"
	}
	return g.DirectoryScopeID
}

type EntraRoleAssignmentRequests []EntraRoleAssignmentRequest

func (c *Client) PIMEntraRoleAssignmentRequests(ctx context.Context) (EntraRoleAssignmentRequests, error) {
	return c.pimEntraRoleAssignmentRequests(ctx, "principal", "status eq 'PendingApproval'")
}

func (c *Client) pimEntraRoleAssignmentRequests(ctx context.Context, filterOn string, queryFilter string) (EntraRoleAssignmentRequests, error) {
	roleAssignmentScheduleRequestsResponse, err := c.client.RoleManagement().Directory().RoleAssignmentScheduleRequests().FilterByCurrentUserWithOn(to.Ptr(filterOn)).Get(ctx, &rolemanagement.DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration{
		QueryParameters: &rolemanagement.DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnRequestBuilderGetQueryParameters{
			Expand: []string{"roleDefinition", "principal"},
			Filter: to.Ptr(queryFilter),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get assignment schedule requests: %w", err)
	}

	graphValues := roleAssignmentScheduleRequestsResponse.GetValue()
	if len(graphValues) == 0 {
		return nil, nil
	}

	var entraRoleAssignmentRequests EntraRoleAssignmentRequests
	for _, graphValue := range graphValues {
		value := EntraRoleAssignmentRequest{}
		err := unmarshalGraphValue(graphValue, &value)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal graph value: %w", err)
		}

		entraRoleAssignmentRequests = append(entraRoleAssignmentRequests, value)
	}

	return entraRoleAssignmentRequests, nil
}

func (c *Client) PIMEntraRoleApprovalRequests(ctx context.Context) (EntraRoleAssignmentRequests, error) {
	return c.pimEntraRoleAssignmentRequests(ctx, "approver", "status eq 'PendingApproval'")
}

func (c *Client) PIMEntraRoleAssignmentScheduleRequest(ctx context.Context, principalID string, entraRoleID string, justification string, startDateTime time.Time, durationHours string, entraRoleScopeID string) (string, error) {
	// requestBody := graphmodels.NewPrivilegedAccessGroupAssignmentScheduleRequest()
	requestBody := graphmodels.NewUnifiedRoleAssignmentScheduleRequest()
	action := graphmodels.SELFACTIVATE_UNIFIEDROLESCHEDULEREQUESTACTIONS
	requestBody.SetAction(&action)
	requestBody.SetPrincipalId(&principalID)
	requestBody.SetRoleDefinitionId(&entraRoleID)
	requestBody.SetDirectoryScopeId(&entraRoleScopeID)
	requestBody.SetJustification(&justification)
	scheduleInfo := graphmodels.NewRequestSchedule()
	scheduleInfo.SetStartDateTime(to.Ptr(startDateTime))
	expiration := graphmodels.NewExpirationPattern()
	expirationType := graphmodels.AFTERDURATION_EXPIRATIONPATTERNTYPE
	expiration.SetTypeEscaped(&expirationType)
	duration, err := serialization.ParseISODuration(durationHours)
	if err != nil {
		return "", fmt.Errorf("failed to parse ISO duration %q: %w", durationHours, err)
	}
	expiration.SetDuration(duration)
	scheduleInfo.SetExpiration(expiration)
	requestBody.SetScheduleInfo(scheduleInfo)
	requestBody.SetJustification(to.Ptr(justification))

	roleAssignmentScheduleRequestsResponse, err := c.client.RoleManagement().Directory().RoleAssignmentScheduleRequests().Post(ctx, requestBody, &rolemanagement.DirectoryRoleAssignmentScheduleRequestsRequestBuilderPostRequestConfiguration{})
	if err != nil {
		return "", fmt.Errorf("failed to post role assignment schedule request: %w", err)
	}

	status := roleAssignmentScheduleRequestsResponse.GetStatus()
	if status == nil {
		return "", fmt.Errorf("response status is nil")
	}

	return *status, nil
}

func (c *Client) PIMEntraRoleGetMaximumExpirationByGroupID(ctx context.Context, entraRoleID string, entraRoleScopeID string) (string, error) {
	roleManagementPolicyAssignmentsResponse, err := c.client.Policies().RoleManagementPolicyAssignments().Get(context.Background(), &policies.RoleManagementPolicyAssignmentsRequestBuilderGetRequestConfiguration{
		QueryParameters: &policies.RoleManagementPolicyAssignmentsRequestBuilderGetQueryParameters{
			Filter: to.Ptr(fmt.Sprintf("scopeId eq '%s' and scopeType eq 'Directory' and roleDefinitionId eq '%s'", entraRoleScopeID, entraRoleID)),
			Expand: []string{"policy($expand=rules)"},
		},
	})

	if err != nil {
		return "", fmt.Errorf("failed to get role management policy assignments: %w", err)
	}

	values := roleManagementPolicyAssignmentsResponse.GetValue()
	if len(values) != 1 {
		return "", fmt.Errorf("expected 1 role management policy assignment, got %d", len(values))
	}

	policy := values[0].GetPolicy()
	if policy == nil {
		return "", fmt.Errorf("policy is nil")
	}

	rules := policy.GetRules()
	if rules == nil {
		return "", fmt.Errorf("rules is nil")
	}

	if len(rules) == 0 {
		return "", fmt.Errorf("no rules found")
	}

	for _, rule := range rules {
		ruleID := rule.GetId()
		if ruleID == nil || *ruleID != "Expiration_EndUser_Assignment" {
			continue
		}

		value := UnifiedRoleManagementPolicyExpirationRule{}
		err := unmarshalGraphValue(rule, &value)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal graph value: %w", err)
		}

		isoDuration, err := serialization.ParseISODuration(value.MaximumDuration)
		if err != nil {
			return "", fmt.Errorf("failed to parse ISO duration %q: %w", value.MaximumDuration, err)
		}

		return isoDuration.String(), nil
	}

	return "", fmt.Errorf("no expiration rule found")
}

type RoleAssignmentApprovalStep struct {
	ID            string `json:"id"`
	AssignedToMe  bool   `json:"assignedToMe"`
	ReviewResult  string `json:"reviewResult"`
	Status        string `json:"status"`
	Justification string `json:"justification"`
	ReviewedBy    struct {
		DisplayName       string `json:"displayName"`
		ID                string `json:"id"`
		Mail              string `json:"mail"`
		UserPrincipalName string `json:"userPrincipalName"`
	} `json:"reviewedBy"`
	ReviewedDateTime time.Time `json:"reviewedDateTime"`
}

type RoleAssignmentApprovalSteps []RoleAssignmentApprovalStep

type RoleAssignmentApproval struct {
	ID    string                      `json:"id"`
	Steps RoleAssignmentApprovalSteps `json:"steps"`
}

func (c *Client) PIMEntraRoleAssignmentApprovalByApprovalID(ctx context.Context, approvalID string, justification string, reviewResult ReviewResult) error {
	roleAssignmentApprovalsResponse, err := c.betaClient.RoleManagement().Directory().RoleAssignmentApprovals().ByApprovalId(approvalID).Get(ctx, &rolemanagementbeta.DirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return fmt.Errorf("failed to get assignment approvals: %w", err)
	}

	value := RoleAssignmentApproval{}
	err = unmarshalGraphValue(roleAssignmentApprovalsResponse, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal graph value: %w", err)
	}

	if len(value.Steps) != 1 {
		return fmt.Errorf("expected 1 approval step, got %d", len(value.Steps))
	}

	step := value.Steps[0]
	if step.ReviewResult != "NotReviewed" {
		return fmt.Errorf("expected review result to be NotReviewed, got %q", step.ReviewResult)
	}

	if step.Status != "InProgress" {
		return fmt.Errorf("expected status to be InProgress, got %q", step.Status)
	}

	if !step.AssignedToMe {
		return fmt.Errorf("expected assignedToMe to be true, got %v", step.AssignedToMe)
	}

	if step.ID == "" {
		return fmt.Errorf("expected ID to be non-empty")
	}

	approvalStepID := step.ID
	requestBody := graphmodelsbeta.NewApprovalStep()
	requestBody.SetReviewResult(to.Ptr(reviewResult.String()))
	requestBody.SetJustification(&justification)

	_, err = c.betaClient.RoleManagement().Directory().RoleAssignmentApprovals().ByApprovalId(approvalID).Steps().ByApprovalStepId(approvalStepID).Patch(ctx, requestBody, &rolemanagementbeta.DirectoryRoleAssignmentApprovalsItemStepsApprovalStepItemRequestBuilderPatchRequestConfiguration{})
	if err != nil {
		return fmt.Errorf("failed to patch assignment approval stage: %w", err)
	}

	return nil
}
