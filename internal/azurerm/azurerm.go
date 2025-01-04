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
	clientFactory *armauthorization.ClientFactory
	cred          azcore.TokenCredential
}

func NewClient(cred azcore.TokenCredential) (*Client, error) {
	clientFactory, err := armauthorization.NewClientFactory("", cred, &arm.ClientOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create client factory: %w", err)
	}
	return &Client{
		clientFactory: clientFactory,
		cred:          cred,
	}, nil
}

type AzureRoleEligibleAssignment armauthorization.RoleEligibilityScheduleInstance

func (a *AzureRoleEligibleAssignment) Role() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.RoleDefinition == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.RoleDefinition.DisplayName
}

func (a *AzureRoleEligibleAssignment) Resource() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Scope == nil {
		return "UNKNOWN"
	}
	return *a.Properties.ExpandedProperties.Scope.DisplayName
}

func (a *AzureRoleEligibleAssignment) ResourceType() string {
	if a.Properties == nil || a.Properties.ExpandedProperties == nil || a.Properties.ExpandedProperties.Scope == nil {
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

// Role
// Resource
// Request type
// Reason
// Condition
// Request time
// Start time
// Request status

// Owner
// M365x96494877
// Member add
// asd
// None
// 1/4/2025, 3:21:37 PM
// 1/4/2025, 3:21:36 PM
// Pending approval​

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
