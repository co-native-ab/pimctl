package cmdhelper

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/co-native-ab/pimctl/internal/azurerm"
	"github.com/co-native-ab/pimctl/internal/credentials"
	"github.com/co-native-ab/pimctl/internal/graph"
	"github.com/spf13/viper"
)

func NewUncachedCredential(credentialMethod credentials.CredentialMethod, tenantID string, clientID string, scopes []string) (*credentials.Credential, error) {
	profileName := viper.GetString("profile")
	return credentials.NewUncached(credentialMethod, tenantID, clientID, scopes, profileName)
}

func NewCachedCredential() (*credentials.Credential, error) {
	profileName := viper.GetString("profile")
	return credentials.NewCached(profileName)
}

func ClearCachedCredential() error {
	profileName := viper.GetString("profile")
	return credentials.ClearCache(profileName)
}

func NewGraphClientWithCachedCredential() (*graph.Client, error) {
	cred, err := NewCachedCredential()
	if err != nil {
		return nil, fmt.Errorf("failed to create cached credential: %w", err)
	}

	graphClient, err := graph.NewClient(cred, []string{credentials.MicrosoftGraphScope})
	if err != nil {
		return nil, fmt.Errorf("failed to create graph client: %w", err)
	}

	return graphClient, nil
}

func NewAzurermClientWithCachedCredential() (*azurerm.Client, error) {
	cred, err := NewCachedCredential()
	if err != nil {
		return nil, fmt.Errorf("failed to create cached credential: %w", err)
	}

	azurermClient, err := azurerm.NewClient(cred)
	if err != nil {
		return nil, fmt.Errorf("failed to create azurerm client: %w", err)
	}

	return azurermClient, nil
}

func GetCachedCredentialAuthenticationRecord() (azidentity.AuthenticationRecord, error) {
	profileName := viper.GetString("profile")
	return credentials.GetCachedAuthenticationRecord(profileName)
}

func PIMGroupAssignmentScheduleRequest(ctx context.Context, graphClient *graph.Client, flagsDuration int, groupID string, justification string) (string, error) {
	me, err := graphClient.Me(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get me: %w", err)
	}

	if me.ID == "" {
		return "", fmt.Errorf("could not get user principal id")
	}

	duration := ""
	if flagsDuration != 0 {
		duration = fmt.Sprintf("PT%dH", flagsDuration)
	} else {
		duration, err = graphClient.PIMGroupGetMaximumExpirationByGroupID(ctx, groupID)
		if err != nil {
			return "", fmt.Errorf("failed to get group role management policy: %w", err)
		}
	}

	status, err := graphClient.PIMGroupAssignmentScheduleRequest(ctx, me.ID, groupID, justification, time.Now(), duration)
	if err != nil {
		return "", fmt.Errorf("failed to create group assignment requests: %w", err)
	}

	return status, nil
}

func PIMEntraRoleAssignmentScheduleRequest(ctx context.Context, graphClient *graph.Client, flagsDuration int, entraRoleID string, justification string, entraRoleScopeID string) (string, error) {
	me, err := graphClient.Me(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get me: %w", err)
	}

	if me.ID == "" {
		return "", fmt.Errorf("could not get user principal id")
	}

	duration := ""
	if flagsDuration != 0 {
		duration = fmt.Sprintf("PT%dH", flagsDuration)
	} else {
		duration, err = graphClient.PIMEntraRoleGetMaximumExpirationByRoleID(ctx, entraRoleID, entraRoleScopeID)
		if err != nil {
			return "", fmt.Errorf("failed to get entra role role management policy: %w", err)
		}
	}

	status, err := graphClient.PIMEntraRoleAssignmentScheduleRequest(ctx, me.ID, entraRoleID, justification, time.Now(), duration, entraRoleScopeID)
	if err != nil {
		return "", fmt.Errorf("failed to create entra role assignment requests: %w", err)
	}

	return status, nil
}

func PIMAzureRoleAssignmentScheduleRequest(ctx context.Context, azurermClient *azurerm.Client, azureRoleEligibleAssignment azurerm.AzureRoleEligibleAssignment, flagsDuration int, justification string) (string, error) {
	duration := ""
	if flagsDuration != 0 {
		duration = fmt.Sprintf("PT%dH", flagsDuration)
	} else {
		var err error
		duration, err = azurermClient.PIMAzureRoleGetMaximumExpirationByRoleID(ctx, *azureRoleEligibleAssignment.Properties.ExpandedProperties.Scope.ID, *azureRoleEligibleAssignment.Properties.RoleDefinitionID)
		if err != nil {
			return "", fmt.Errorf("failed to get entra role role management policy: %w", err)
		}
	}

	status, err := azurermClient.PIMAzureRoleAssignmentScheduleRequest(ctx, *azureRoleEligibleAssignment.Properties.ExpandedProperties.Principal.ID, *azureRoleEligibleAssignment.Properties.RoleDefinitionID, justification, time.Now(), duration, *azureRoleEligibleAssignment.Properties.ExpandedProperties.Scope.ID)
	if err != nil {
		return "", fmt.Errorf("failed to create entra role assignment requests: %w", err)
	}

	return status, nil
}
