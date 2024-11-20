package graph

import (
	"context"
	"fmt"

	"github.com/microsoftgraph/msgraph-sdk-go/groups"
)

type Group struct {
	ID              string   `json:"id"`
	DisplayName     string   `json:"displayName"`
	Description     string   `json:"description"`
	SecurityEnabled bool     `json:"securityEnabled"`
	GroupTypes      []string `json:"groupTypes"`
}

func (g Group) GroupType() string {
	if len(g.GroupTypes) == 0 {
		return "Security"
	}
	return "Microsoft 365"
}

func (c *Client) GroupGetByID(ctx context.Context, id string) (Group, error) {
	res, err := c.client.Groups().ByGroupId(id).Get(ctx, &groups.GroupItemRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return Group{}, fmt.Errorf("failed to get group by ID: %w", err)
	}

	group := Group{}
	err = unmarshalGraphValue(res, &group)
	if err != nil {
		return Group{}, fmt.Errorf("failed to unmarshal graph value: %w", err)
	}

	return group, nil
}
