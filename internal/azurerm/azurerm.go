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
