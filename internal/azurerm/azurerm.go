package azurerm

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

func prettyPrintValue[T any](data T) {
	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("failed to pretty print json: %v\n", err)
		return
	}
	fmt.Println(string(json))
}

type Client struct {
	clientFactory    *armauthorization.ClientFactory
	cred             azcore.TokenCredential
	armClientOptions *arm.ClientOptions
}

func NewClient(cred azcore.TokenCredential) (*Client, error) {
	armClientOptions := &arm.ClientOptions{}
	clientFactory, err := armauthorization.NewClientFactory("", cred, armClientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create client factory: %w", err)
	}
	return &Client{
		clientFactory:    clientFactory,
		cred:             cred,
		armClientOptions: armClientOptions,
	}, nil
}

type AzureRoleEligibleAssignment armauthorization.RoleEligibilityScheduleInstance

func (a *AzureRoleEligibleAssignment) Role() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.RoleDefinition == nil || a.Properties.ExpandedProperties.RoleDefinition.DisplayName == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.RoleDefinition.DisplayName
}

func (a *AzureRoleEligibleAssignment) Resource() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Scope == nil || a.Properties.ExpandedProperties.Scope.DisplayName == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.Scope.DisplayName
}

func (a *AzureRoleEligibleAssignment) ResourceType() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Scope == nil || a.Properties.ExpandedProperties.Scope.Type == nil {
		return "UNKNOWN"
	}

	switch *a.Properties.ExpandedProperties.Scope.Type {
	case "managementgroup":
		return "Management group"
	}

	return *a.Properties.ExpandedProperties.Scope.Type
}

func (a *AzureRoleEligibleAssignment) Membership() string {
	if a.Properties == nil || a.Properties.MemberType == nil {
		return "UNKNOWN"
	}
	return string(*a.Properties.MemberType)
}

func (a *AzureRoleEligibleAssignment) Condition() string {
	if a.Properties == nil || a.Properties.Condition == nil {
		return "None"
	}
	return *a.Properties.Condition
}

func (a *AzureRoleEligibleAssignment) EndTime() string {
	if a.Properties == nil || a.Properties.EndDateTime == nil {
		return "Permanent"
	}
	return a.Properties.EndDateTime.Local().Format(time.RFC3339)
}

func (a *AzureRoleEligibleAssignment) RoleDefinitionID() string {
	if a.Properties == nil || a.Properties.RoleDefinitionID == nil {
		return "UNKNOWN"
	}
	return *a.Properties.RoleDefinitionID
}

func (a *AzureRoleEligibleAssignment) Scope() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Scope == nil || a.Properties.ExpandedProperties.Scope.ID == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.Scope.ID
}

type AzureRoleEligibleAssignments []AzureRoleEligibleAssignment

func (c *Client) PIMAzureRoleEligibleAssignments(ctx context.Context) (AzureRoleEligibleAssignments, error) {
	roleEligibilitySchedulesClient := c.clientFactory.NewRoleEligibilityScheduleInstancesClient()

	pager := roleEligibilitySchedulesClient.NewListForScopePager("", &armauthorization.RoleEligibilityScheduleInstancesClientListForScopeOptions{
		Filter: to.Ptr("asTarget()"),
	})

	azureRoleEligibleAssignments := AzureRoleEligibleAssignments{}
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get next page: %w", err)
		}
		for _, v := range page.Value {
			if v == nil {
				return nil, fmt.Errorf("role eligibility schedule instance is nil")
			}

			azureRoleEligibleAssignments = append(azureRoleEligibleAssignments, AzureRoleEligibleAssignment(*v))
		}
	}

	return azureRoleEligibleAssignments, nil
}

type AzureRoleAssignmentRequest armauthorization.RoleAssignmentScheduleRequest

func (a *AzureRoleAssignmentRequest) Role() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.RoleDefinition == nil || a.Properties.ExpandedProperties.RoleDefinition.DisplayName == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.RoleDefinition.DisplayName
}

func (a *AzureRoleAssignmentRequest) Resource() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Scope == nil || a.Properties.ExpandedProperties.Scope.DisplayName == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.Scope.DisplayName
}

func (a *AzureRoleAssignmentRequest) ResourceType() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Scope == nil || a.Properties.ExpandedProperties.Scope.Type == nil {
		return "UNKNOWN"
	}

	switch *a.Properties.ExpandedProperties.Scope.Type {
	case "managementgroup":
		return "Management group"
	}

	return *a.Properties.ExpandedProperties.Scope.Type
}

func (a *AzureRoleAssignmentRequest) RequestType() string {
	if a.Properties == nil || a.Properties.RequestType == nil {
		return "UNKNOWN"
	}
	return string(*a.Properties.RequestType)
}

func (a *AzureRoleAssignmentRequest) Reason() string {
	if a.Properties == nil || a.Properties.Justification == nil {
		return ""
	}
	return *a.Properties.Justification
}

func (a *AzureRoleAssignmentRequest) Condition() string {
	if a.Properties == nil || a.Properties.Condition == nil {
		return "None"
	}
	return *a.Properties.Condition
}

func (a *AzureRoleAssignmentRequest) RequestTime() string {
	if a.Properties == nil || a.Properties.CreatedOn == nil {
		return "UNKNOWN"
	}
	return a.Properties.CreatedOn.Local().Format(time.RFC3339)
}

func (a *AzureRoleAssignmentRequest) StartTime() string {
	if a.Properties == nil || a.Properties.ScheduleInfo == nil || a.Properties.ScheduleInfo.StartDateTime == nil {
		return "UNKNOWN"
	}
	return a.Properties.ScheduleInfo.StartDateTime.Local().Format(time.RFC3339)
}

func (a *AzureRoleAssignmentRequest) RequestStatus() string {
	if a.Properties == nil || a.Properties.Status == nil {
		return "UNKNOWN"
	}
	return string(*a.Properties.Status)
}

func (a *AzureRoleAssignmentRequest) Requestor() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Principal == nil || a.Properties.ExpandedProperties.Principal.DisplayName == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.Principal.DisplayName
}

func (a *AzureRoleAssignmentRequest) EndTime() string {
	if a.Properties == nil || a.Properties.ScheduleInfo == nil || a.Properties.ScheduleInfo.Expiration == nil || a.Properties.ScheduleInfo.Expiration.Type == nil {
		return "UNKNOWN"
	}

	expirationType := *a.Properties.ScheduleInfo.Expiration.Type

	switch expirationType {
	case armauthorization.TypeAfterDateTime:
		if a.Properties.ScheduleInfo.Expiration.EndDateTime == nil {
			return "UNKNOWN"
		}
		return a.Properties.ScheduleInfo.Expiration.EndDateTime.Local().Format(time.RFC3339)
	case armauthorization.TypeAfterDuration:
		if a.Properties.ScheduleInfo.Expiration.Duration == nil || a.Properties.ScheduleInfo.StartDateTime == nil {
			return "UNKNOWN"
		}
		startTime := a.Properties.ScheduleInfo.StartDateTime
		isoDuration, err := serialization.ParseISODuration(*a.Properties.ScheduleInfo.Expiration.Duration)
		if err != nil {
			return "Unable to parse duration"
		}
		duration, err := isoDuration.ToDuration()
		if err != nil {
			return "Unable to convert duration"
		}
		endTime := startTime.Add(duration)
		return endTime.Local().Format(time.RFC3339)
	case armauthorization.TypeNoExpiration:
		return "Permanent"
	default:
		return "Unknown"
	}
}

func (a *AzureRoleAssignmentRequest) Scope() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Scope == nil || a.Properties.ExpandedProperties.Scope.ID == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.Scope.ID
}

type AzureRoleAssignmentRequests []AzureRoleAssignmentRequest

type AzureRoleAssignmentRequestsFilter string

const (
	AzureRoleAssignmentRequestsFilterApprover AzureRoleAssignmentRequestsFilter = "asApprover()"
	AzureRoleAssignmentRequestsFilterTarget   AzureRoleAssignmentRequestsFilter = "asTarget()"
)

func (c *Client) PIMAzureRoleAssignmentRequests(ctx context.Context) (AzureRoleAssignmentRequests, error) {
	return c.pimAzureRoleAssignmentRequests(ctx, AzureRoleAssignmentRequestsFilterTarget)
}

func (c *Client) PIMAzureRoleApprovalRequests(ctx context.Context) (AzureRoleAssignmentRequests, error) {
	return c.pimAzureRoleAssignmentRequests(ctx, AzureRoleAssignmentRequestsFilterApprover)
}

func (c *Client) pimAzureRoleAssignmentRequests(ctx context.Context, filter AzureRoleAssignmentRequestsFilter) (AzureRoleAssignmentRequests, error) {
	roleAssignmentScheduleRequestsClient := c.clientFactory.NewRoleAssignmentScheduleRequestsClient()

	pager := roleAssignmentScheduleRequestsClient.NewListForScopePager("", &armauthorization.RoleAssignmentScheduleRequestsClientListForScopeOptions{
		Filter: to.Ptr(string(filter)),
	})

	azureRoleAssignmentRequests := AzureRoleAssignmentRequests{}
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get next page: %w", err)
		}
		for _, v := range page.Value {
			if v == nil {
				return nil, fmt.Errorf("role assignment schedule request is nil")
			}

			azureRoleAssignmentRequests = append(azureRoleAssignmentRequests, AzureRoleAssignmentRequest(*v))
		}
	}

	return azureRoleAssignmentRequests, nil
}

type AzureRoleActiveAssignment armauthorization.RoleAssignmentScheduleInstance

func (a *AzureRoleActiveAssignment) Role() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.RoleDefinition == nil || a.Properties.ExpandedProperties.RoleDefinition.DisplayName == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.RoleDefinition.DisplayName
}

func (a *AzureRoleActiveAssignment) Resource() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Scope == nil || a.Properties.ExpandedProperties.Scope.DisplayName == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.Scope.DisplayName
}

func (a *AzureRoleActiveAssignment) ResourceType() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Scope == nil || a.Properties.ExpandedProperties.Scope.Type == nil {
		return "UNKNOWN"
	}

	switch *a.Properties.ExpandedProperties.Scope.Type {
	case "managementgroup":
		return "Management group"
	}

	return *a.Properties.ExpandedProperties.Scope.Type
}

func (a *AzureRoleActiveAssignment) Membership() string {
	if a.Properties == nil || a.Properties.MemberType == nil {
		return "UNKNOWN"
	}
	return string(*a.Properties.MemberType)
}

func (a *AzureRoleActiveAssignment) Condition() string {
	if a.Properties == nil || a.Properties.Condition == nil {
		return "None"
	}
	return *a.Properties.Condition
}

func (a *AzureRoleActiveAssignment) State() string {
	if a.Properties == nil || a.Properties.Status == nil {
		return "UNKNOWN"
	}
	return string(*a.Properties.Status)
}

func (a *AzureRoleActiveAssignment) EndTime() string {
	if a.Properties == nil || a.Properties.EndDateTime == nil {
		return "UNKNOWN"
	}
	return a.Properties.EndDateTime.Local().Format(time.RFC3339)
}

func (a *AzureRoleActiveAssignment) Scope() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Scope == nil || a.Properties.ExpandedProperties.Scope.ID == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.Scope.ID
}

type AzureRoleActiveAssignments []AzureRoleActiveAssignment

func (c *Client) PIMAzureRoleActiveAssignments(ctx context.Context) (AzureRoleActiveAssignments, error) {
	// NOTE: Listing an empty scope roleAssignmentScheduleInstancesClient.NewListForScopePager("", &armauthorization.RoleAssignmentScheduleInstancesClientListForScopeOptions{})
	//       doesn't work and returns an empty list. So, we need to list for a specific scopes for now. It seems to be the same behavior in the portal.
	roleEligibleAssignments, err := c.PIMAzureRoleEligibleAssignments(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get eligible assignments: %w", err)
	}

	scopes := map[string]struct{}{}
	for _, v := range roleEligibleAssignments {
		if v.Properties == nil || v.Properties.ExpandedProperties == nil || v.Properties.ExpandedProperties.Scope == nil || v.Properties.ExpandedProperties.Scope.ID == nil {
			return nil, fmt.Errorf("scope for role eligibible assignment is nil")
		}
		scopes[*v.Properties.ExpandedProperties.Scope.ID] = struct{}{}
	}

	if len(scopes) == 0 {
		return nil, nil
	}

	roleAssignmentScheduleInstancesClient := c.clientFactory.NewRoleAssignmentScheduleInstancesClient()

	azureRoleActiveAssignments := AzureRoleActiveAssignments{}
	for scope := range scopes {
		pager := roleAssignmentScheduleInstancesClient.NewListForScopePager(scope, &armauthorization.RoleAssignmentScheduleInstancesClientListForScopeOptions{
			Filter: to.Ptr("asTarget()"),
		})

		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get next page: %w", err)
			}
			for _, v := range page.Value {
				if v == nil {
					return nil, fmt.Errorf("role assignment is nil")
				}

				if v.Properties == nil || v.Properties.PrincipalType == nil {
					return nil, fmt.Errorf("principal type is nil")
				}

				if *v.Properties.PrincipalType != armauthorization.PrincipalTypeUser {
					continue
				}

				azureRoleActiveAssignments = append(azureRoleActiveAssignments, AzureRoleActiveAssignment(*v))
			}
		}
	}

	return azureRoleActiveAssignments, nil
}

func (c *Client) PIMAzureRoleAssignmentScheduleRequest(ctx context.Context, principalID string, roleDefinitionID string, justification string, startDateTime time.Time, durationHours string, scope string) (string, error) {
	roleAssignmentScheduleRequestsClient := c.clientFactory.NewRoleAssignmentScheduleRequestsClient()

	roleAssignmentScheduleRequestNameUUID, err := newUUID()
	if err != nil {
		return "", fmt.Errorf("failed to create uuid: %w", err)
	}
	roleAssignmentScheduleRequestName := roleAssignmentScheduleRequestNameUUID.String()
	parameters := armauthorization.RoleAssignmentScheduleRequest{
		Properties: &armauthorization.RoleAssignmentScheduleRequestProperties{
			Justification:    &justification,
			PrincipalID:      &principalID,
			RequestType:      to.Ptr(armauthorization.RequestTypeSelfActivate),
			RoleDefinitionID: &roleDefinitionID,
			ScheduleInfo: &armauthorization.RoleAssignmentScheduleRequestPropertiesScheduleInfo{
				Expiration: &armauthorization.RoleAssignmentScheduleRequestPropertiesScheduleInfoExpiration{
					Type:     to.Ptr(armauthorization.TypeAfterDuration),
					Duration: &durationHours,
				},
				StartDateTime: &startDateTime,
			},
		},
	}
	res, err := roleAssignmentScheduleRequestsClient.Create(ctx, scope, roleAssignmentScheduleRequestName, parameters, &armauthorization.RoleAssignmentScheduleRequestsClientCreateOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to create role assignment schedule request: %w", err)
	}

	if res.Properties == nil || res.Properties.Status == nil {
		return "", fmt.Errorf("role assignment schedule request status is nil")
	}

	return string(*res.Properties.Status), nil
}

func (c *Client) PIMAzureRoleGetMaximumExpirationByRoleID(ctx context.Context, scope string, roleDefinitionID string) (string, error) {
	roleManagementPolicyAssignmentsClient, err := newRoleManagementPolicyAssignmentsClient(c.cred, c.armClientOptions)
	if err != nil {
		return "", fmt.Errorf("failed to create role management policy assignments client: %w", err)
	}

	pager := roleManagementPolicyAssignmentsClient.NewListForScopePager(scope, &roleManagementPolicyAssignmentsClientListForScopeOptions{
		Filter: to.Ptr(fmt.Sprintf("roleDefinitionId eq '%s'", roleDefinitionID)),
	})

	value := &roleManagementPolicyAssignment{}
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return "", fmt.Errorf("failed to get next page: %w", err)
		}

		if len(page.Value) != 1 {
			return "", fmt.Errorf("expected 1 role management policy assignment but got %d", len(page.Value))
		}

		value = page.Value[0]
	}

	if value == nil || value.Properties == nil || value.Properties.EffectiveRules == nil {
		return "", fmt.Errorf("role management policy assignment is nil")
	}

	if len(value.Properties.EffectiveRules) == 0 {
		return "", fmt.Errorf("no effective rules")
	}

	for _, v := range value.Properties.EffectiveRules {
		if v == nil || v.ID == nil || *v.ID != "Expiration_EndUser_Assignment" {
			continue
		}

		if v.MaximumDuration == nil {
			return "", fmt.Errorf("maximum duration is nil")
		}

		isoDuration, err := serialization.ParseISODuration(*v.MaximumDuration)
		if err != nil {
			return "", fmt.Errorf("failed to parse duration: %w", err)
		}

		return isoDuration.String(), nil
	}

	return "", fmt.Errorf("no expiration rule found")
}
