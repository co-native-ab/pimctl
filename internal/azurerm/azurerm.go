package azurerm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"github.com/lestrrat-go/jwx/jwt"
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

func (c *Client) Me(ctx context.Context, scopes []string, tenantID string) (string, error) {
	token, err := c.cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes:    scopes,
		EnableCAE: true,
		TenantID:  tenantID,
	})

	parsedToken, err := jwt.ParseString(token.Token)
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	upnRaw, ok := parsedToken.Get("upn")
	if !ok {
		return "", fmt.Errorf("failed to get upn from token")
	}

	upn, ok := upnRaw.(string)
	if !ok {
		return "", fmt.Errorf("failed to cast upn to string")
	}

	return upn, nil
}

func (c *Client) PIMAzureRoleEligibleAssignments(ctx context.Context) error {
	roleEligibilitySchedulesClient := c.clientFactory.NewRoleEligibilityScheduleInstancesClient()

	pager := roleEligibilitySchedulesClient.NewListForScopePager("", &armauthorization.RoleEligibilityScheduleInstancesClientListForScopeOptions{
		Filter: to.Ptr("asTarget()"),
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to get next page: %w", err)
		}
		for _, v := range page.Value {
			prettyPrintValue(v)
		}
	}

	return nil
}
