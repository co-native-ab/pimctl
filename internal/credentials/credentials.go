package credentials

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type Credential struct {
	cred tokenGetter
}

func (c *Credential) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return c.cred.GetToken(ctx, options)
}

func (c *Credential) Authenticate(ctx context.Context, opts *policy.TokenRequestOptions) (azidentity.AuthenticationRecord, error) {
	return c.cred.Authenticate(ctx, opts)
}

type tokenGetter interface {
	GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error)
	Authenticate(ctx context.Context, opts *policy.TokenRequestOptions) (azidentity.AuthenticationRecord, error)
}

func NewUncached(credentialMethod CredentialMethod, tenantID string, clientID string, scopes []string, profileName string) (*Credential, error) {
	cred, err := newUncachedCredential(credentialMethod, tenantID, clientID, scopes, profileName)
	if err != nil {
		return nil, fmt.Errorf("failed to create cached credential: %w", err)
	}
	return &Credential{cred: cred}, nil
}

func NewCached(profileName string) (*Credential, error) {
	cred, err := newCachedCredential(profileName)
	if err != nil {
		return nil, fmt.Errorf("failed to create cached credential: %w", err)
	}
	return &Credential{cred: cred}, nil
}
