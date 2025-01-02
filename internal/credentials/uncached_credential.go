package credentials

import (
	"context"
	"fmt"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity/cache"
)

type uncachedCredential struct {
	cred             tokenGetter
	credCache        azidentity.Cache
	mu               sync.Mutex
	tenantID         string
	clientID         string
	scopes           []string
	authRecord       azidentity.AuthenticationRecord
	credentialMethod CredentialMethod
	profileName      string
}

func (c *uncachedCredential) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return c.cred.GetToken(ctx, options)
}

func (c *uncachedCredential) Authenticate(ctx context.Context, opts *policy.TokenRequestOptions) (azidentity.AuthenticationRecord, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	record, err := c.cred.Authenticate(ctx, opts)
	if err != nil {
		return azidentity.AuthenticationRecord{}, fmt.Errorf("failed to authenticate: %w", err)
	}

	configCache := configCache{
		credentialMethod:       c.credentialMethod,
		azAuthenticationRecord: record,
	}

	err = configCache.Store(c.profileName)
	if err != nil {
		return azidentity.AuthenticationRecord{}, fmt.Errorf("failed to store config cache: %w", err)
	}

	c.authRecord = record
	c.cred, err = newCachedAzureCredential(c.credentialMethod, c.clientID, c.tenantID, c.credCache, c.authRecord)
	if err != nil {
		return azidentity.AuthenticationRecord{}, fmt.Errorf("failed to create azure credential: %w", err)
	}

	return c.authRecord, nil
}

func newUncachedCredential(credentialMethod CredentialMethod, tenantID string, clientID string, scopes []string, profileName string) (*uncachedCredential, error) {
	azCredCache, err := cache.New(&cache.Options{
		Name: getCacheName(profileName),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create cache: %w", err)
	}

	cred, err := newCachedAzureCredential(credentialMethod, clientID, tenantID, azCredCache, azidentity.AuthenticationRecord{})
	if err != nil {
		return nil, fmt.Errorf("failed to create azure credential: %w", err)
	}

	return &uncachedCredential{
		cred:             cred,
		credCache:        azCredCache,
		tenantID:         tenantID,
		clientID:         clientID,
		scopes:           scopes,
		credentialMethod: credentialMethod,
		profileName:      profileName,
	}, nil
}

type cachedAzureCredential struct {
	cred tokenGetter
}

func (c *cachedAzureCredential) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return c.cred.GetToken(ctx, options)
}

func (c *cachedAzureCredential) Authenticate(ctx context.Context, opts *policy.TokenRequestOptions) (azidentity.AuthenticationRecord, error) {
	return c.cred.Authenticate(ctx, opts)
}

func newCachedAzureCredential(credentialMethod CredentialMethod, clientID string, tenantID string, credCache azidentity.Cache, authRecord azidentity.AuthenticationRecord) (*cachedAzureCredential, error) {
	switch credentialMethod {
	case DeviceCodeCredentialMethod:
		cred, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
			ClientID:                       clientID,
			TenantID:                       tenantID,
			Cache:                          credCache,
			AuthenticationRecord:           authRecord,
			DisableAutomaticAuthentication: true,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create device code credential: %w", err)
		}
		return &cachedAzureCredential{cred: cred}, nil
	case InteractiveBrowserCredentialMethod:
		cred, err := azidentity.NewInteractiveBrowserCredential(&azidentity.InteractiveBrowserCredentialOptions{
			ClientID:                       clientID,
			TenantID:                       tenantID,
			Cache:                          credCache,
			AuthenticationRecord:           authRecord,
			DisableAutomaticAuthentication: true,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create interactive browser credential: %w", err)
		}
		return &cachedAzureCredential{cred: cred}, nil
	}
	return nil, fmt.Errorf("unknown credential method %q", credentialMethod)
}
