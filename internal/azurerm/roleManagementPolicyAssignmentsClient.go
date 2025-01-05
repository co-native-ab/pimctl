// This copy of the client is created to make it possible to use $filter with RoleManagementPolicyAssignmentsClient NewListForScopePager() method as well as return EffectiveRules.
//
// Original file is located at: https://github.com/Azure/azure-sdk-for-go/blob/sdk/resourcemanager/authorization/armauthorization/v3.0.0-beta.2/sdk/resourcemanager/authorization/armauthorization/rolemanagementpolicyassignments_client.go
//
// Original file header:
// ---
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// ---
//
// Original license:
// ---
// The MIT License (MIT)

// Copyright (c) Microsoft Corporation.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
// ---

package azurerm

import (
	"context"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

const (
	moduleName    = "github.com/co-native-ab/pimctl"
	moduleVersion = "v0.0.0"
)

type roleManagementPolicyAssignmentsClient struct {
	internal *arm.Client
}

func newRoleManagementPolicyAssignmentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*roleManagementPolicyAssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &roleManagementPolicyAssignmentsClient{
		internal: cl,
	}
	return client, nil
}

type roleManagementPolicyAssignmentsClientListForScopeResponse struct {
	roleManagementPolicyAssignmentListResult
}

type roleManagementPolicyAssignmentListResult struct {
	NextLink *string
	Value    []*roleManagementPolicyAssignment
}

type roleManagementPolicyAssignment struct {
	Properties *roleManagementPolicyAssignmentProperties
	ID         *string
	Name       *string
	Type       *string
}

type roleManagementPolicyAssignmentPropertiesEffectiveRule struct {
	MaximumDuration *string
	ID              *string
}

type roleManagementPolicyAssignmentProperties struct {
	PolicyID                   *string
	RoleDefinitionID           *string
	Scope                      *string
	PolicyAssignmentProperties *armauthorization.PolicyAssignmentProperties
	// NOTE: Added EffectiveRules field
	EffectiveRules []*roleManagementPolicyAssignmentPropertiesEffectiveRule
}

type roleManagementPolicyAssignmentsClientListForScopeOptions struct {
	Filter *string
}

func (client *roleManagementPolicyAssignmentsClient) NewListForScopePager(scope string, options *roleManagementPolicyAssignmentsClientListForScopeOptions) *runtime.Pager[roleManagementPolicyAssignmentsClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[roleManagementPolicyAssignmentsClientListForScopeResponse]{
		More: func(page roleManagementPolicyAssignmentsClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *roleManagementPolicyAssignmentsClientListForScopeResponse) (roleManagementPolicyAssignmentsClientListForScopeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RoleManagementPolicyAssignmentsClient.NewListForScopePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listForScopeCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return roleManagementPolicyAssignmentsClientListForScopeResponse{}, err
			}
			return client.listForScopeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

func (client *roleManagementPolicyAssignmentsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *roleManagementPolicyAssignmentsClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	// NOTE: Added $filter query parameter
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01") // NOTE: Removed -preview prefix
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

func (client *roleManagementPolicyAssignmentsClient) listForScopeHandleResponse(resp *http.Response) (roleManagementPolicyAssignmentsClientListForScopeResponse, error) {
	result := roleManagementPolicyAssignmentsClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.roleManagementPolicyAssignmentListResult); err != nil {
		return roleManagementPolicyAssignmentsClientListForScopeResponse{}, err
	}
	return result, nil
}
