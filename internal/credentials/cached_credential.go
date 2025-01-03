package credentials

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity/cache"
)

type cachedCredential struct {
	cred tokenGetter
}

func (c *cachedCredential) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return c.cred.GetToken(ctx, options)
}

func (c *cachedCredential) Authenticate(ctx context.Context, opts *policy.TokenRequestOptions) (azidentity.AuthenticationRecord, error) {
	return c.cred.Authenticate(ctx, opts)
}

func newCachedCredential(profileName string) (*cachedCredential, error) {
	credCache, err := cache.New(&cache.Options{
		Name: getCacheName(profileName),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create cache: %w", err)
	}

	configCache, err := getCacheConfig(profileName)
	if err != nil {
		return nil, fmt.Errorf("failed to get cache config: %w", err)
	}

	cred, err := newCachedAzureCredential(configCache.credentialMethod, configCache.azAuthenticationRecord.ClientID, configCache.azAuthenticationRecord.TenantID, credCache, configCache.azAuthenticationRecord)
	if err != nil {
		return nil, fmt.Errorf("failed to create azure credential: %w", err)
	}

	return &cachedCredential{
		cred: cred,
	}, nil
}
