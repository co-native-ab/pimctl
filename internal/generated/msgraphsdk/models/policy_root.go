package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PolicyRoot struct {
    Entity
    // The policy that controls the idle time out for web sessions for applications.
    activityBasedTimeoutPolicies []ActivityBasedTimeoutPolicyable
    // The policy by which consent requests are created and managed for the entire tenant.
    adminConsentRequestPolicy AdminConsentRequestPolicyable
    // The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
    appManagementPolicies []AppManagementPolicyable
    // The policy configuration of the self-service sign-up experience of external users.
    authenticationFlowsPolicy AuthenticationFlowsPolicyable
    // The authentication methods and the users that are allowed to use them to sign in and perform multifactor authentication (MFA) in Microsoft Entra ID.
    authenticationMethodsPolicy AuthenticationMethodsPolicyable
    // The authentication method combinations that are to be used in scenarios defined by Microsoft Entra Conditional Access.
    authenticationStrengthPolicies []AuthenticationStrengthPolicyable
    // The policy that controls Microsoft Entra authorization settings.
    authorizationPolicy AuthorizationPolicyable
    // The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
    claimsMappingPolicies []ClaimsMappingPolicyable
    // The custom rules that define an access scenario.
    conditionalAccessPolicies []ConditionalAccessPolicyable
    // The custom rules that define an access scenario when interacting with external Microsoft Entra tenants.
    crossTenantAccessPolicy CrossTenantAccessPolicyable
    // The tenant-wide policy that enforces app management restrictions for all applications and service principals.
    defaultAppManagementPolicy TenantAppManagementPolicyable
    // The deviceRegistrationPolicy property
    deviceRegistrationPolicy DeviceRegistrationPolicyable
    // The feature rollout policy associated with a directory object.
    featureRolloutPolicies []FeatureRolloutPolicyable
    // The policy to control Microsoft Entra authentication behavior for federated users.
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicyable
    // The policy that represents the security defaults that protect against common attacks.
    identitySecurityDefaultsEnforcementPolicy IdentitySecurityDefaultsEnforcementPolicyable
    // The policy that specifies the conditions under which consent can be granted.
    permissionGrantPolicies []PermissionGrantPolicyable
    // Specifies the various policies associated with scopes and roles.
    roleManagementPolicies []UnifiedRoleManagementPolicyable
    // The assignment of a role management policy to a role definition object.
    roleManagementPolicyAssignments []UnifiedRoleManagementPolicyAssignmentable
    // The policy that specifies the characteristics of SAML tokens issued by Microsoft Entra ID.
    tokenIssuancePolicies []TokenIssuancePolicyable
    // The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Microsoft Entra ID.
    tokenLifetimePolicies []TokenLifetimePolicyable
}
// NewPolicyRoot instantiates a new PolicyRoot and sets the default values.
func NewPolicyRoot()(*PolicyRoot) {
    m := &PolicyRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePolicyRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePolicyRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicyRoot(), nil
}
// GetActivityBasedTimeoutPolicies gets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
// returns a []ActivityBasedTimeoutPolicyable when successful
func (m *PolicyRoot) GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicyable) {
    return m.activityBasedTimeoutPolicies
}
// GetAdminConsentRequestPolicy gets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
// returns a AdminConsentRequestPolicyable when successful
func (m *PolicyRoot) GetAdminConsentRequestPolicy()(AdminConsentRequestPolicyable) {
    return m.adminConsentRequestPolicy
}
// GetAppManagementPolicies gets the appManagementPolicies property value. The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
// returns a []AppManagementPolicyable when successful
func (m *PolicyRoot) GetAppManagementPolicies()([]AppManagementPolicyable) {
    return m.appManagementPolicies
}
// GetAuthenticationFlowsPolicy gets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
// returns a AuthenticationFlowsPolicyable when successful
func (m *PolicyRoot) GetAuthenticationFlowsPolicy()(AuthenticationFlowsPolicyable) {
    return m.authenticationFlowsPolicy
}
// GetAuthenticationMethodsPolicy gets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multifactor authentication (MFA) in Microsoft Entra ID.
// returns a AuthenticationMethodsPolicyable when successful
func (m *PolicyRoot) GetAuthenticationMethodsPolicy()(AuthenticationMethodsPolicyable) {
    return m.authenticationMethodsPolicy
}
// GetAuthenticationStrengthPolicies gets the authenticationStrengthPolicies property value. The authentication method combinations that are to be used in scenarios defined by Microsoft Entra Conditional Access.
// returns a []AuthenticationStrengthPolicyable when successful
func (m *PolicyRoot) GetAuthenticationStrengthPolicies()([]AuthenticationStrengthPolicyable) {
    return m.authenticationStrengthPolicies
}
// GetAuthorizationPolicy gets the authorizationPolicy property value. The policy that controls Microsoft Entra authorization settings.
// returns a AuthorizationPolicyable when successful
func (m *PolicyRoot) GetAuthorizationPolicy()(AuthorizationPolicyable) {
    return m.authorizationPolicy
}
// GetClaimsMappingPolicies gets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
// returns a []ClaimsMappingPolicyable when successful
func (m *PolicyRoot) GetClaimsMappingPolicies()([]ClaimsMappingPolicyable) {
    return m.claimsMappingPolicies
}
// GetConditionalAccessPolicies gets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
// returns a []ConditionalAccessPolicyable when successful
func (m *PolicyRoot) GetConditionalAccessPolicies()([]ConditionalAccessPolicyable) {
    return m.conditionalAccessPolicies
}
// GetCrossTenantAccessPolicy gets the crossTenantAccessPolicy property value. The custom rules that define an access scenario when interacting with external Microsoft Entra tenants.
// returns a CrossTenantAccessPolicyable when successful
func (m *PolicyRoot) GetCrossTenantAccessPolicy()(CrossTenantAccessPolicyable) {
    return m.crossTenantAccessPolicy
}
// GetDefaultAppManagementPolicy gets the defaultAppManagementPolicy property value. The tenant-wide policy that enforces app management restrictions for all applications and service principals.
// returns a TenantAppManagementPolicyable when successful
func (m *PolicyRoot) GetDefaultAppManagementPolicy()(TenantAppManagementPolicyable) {
    return m.defaultAppManagementPolicy
}
// GetDeviceRegistrationPolicy gets the deviceRegistrationPolicy property value. The deviceRegistrationPolicy property
// returns a DeviceRegistrationPolicyable when successful
func (m *PolicyRoot) GetDeviceRegistrationPolicy()(DeviceRegistrationPolicyable) {
    return m.deviceRegistrationPolicy
}
// GetFeatureRolloutPolicies gets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
// returns a []FeatureRolloutPolicyable when successful
func (m *PolicyRoot) GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable) {
    return m.featureRolloutPolicies
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PolicyRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityBasedTimeoutPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActivityBasedTimeoutPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActivityBasedTimeoutPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ActivityBasedTimeoutPolicyable)
                }
            }
            m.SetActivityBasedTimeoutPolicies(res)
        }
        return nil
    }
    res["adminConsentRequestPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminConsentRequestPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConsentRequestPolicy(val.(AdminConsentRequestPolicyable))
        }
        return nil
    }
    res["appManagementPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppManagementPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppManagementPolicyable)
                }
            }
            m.SetAppManagementPolicies(res)
        }
        return nil
    }
    res["authenticationFlowsPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationFlowsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationFlowsPolicy(val.(AuthenticationFlowsPolicyable))
        }
        return nil
    }
    res["authenticationMethodsPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethodsPolicy(val.(AuthenticationMethodsPolicyable))
        }
        return nil
    }
    res["authenticationStrengthPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationStrengthPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationStrengthPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationStrengthPolicyable)
                }
            }
            m.SetAuthenticationStrengthPolicies(res)
        }
        return nil
    }
    res["authorizationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationPolicy(val.(AuthorizationPolicyable))
        }
        return nil
    }
    res["claimsMappingPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClaimsMappingPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClaimsMappingPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ClaimsMappingPolicyable)
                }
            }
            m.SetClaimsMappingPolicies(res)
        }
        return nil
    }
    res["conditionalAccessPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionalAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConditionalAccessPolicyable)
                }
            }
            m.SetConditionalAccessPolicies(res)
        }
        return nil
    }
    res["crossTenantAccessPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossTenantAccessPolicy(val.(CrossTenantAccessPolicyable))
        }
        return nil
    }
    res["defaultAppManagementPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTenantAppManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultAppManagementPolicy(val.(TenantAppManagementPolicyable))
        }
        return nil
    }
    res["deviceRegistrationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceRegistrationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRegistrationPolicy(val.(DeviceRegistrationPolicyable))
        }
        return nil
    }
    res["featureRolloutPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFeatureRolloutPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FeatureRolloutPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FeatureRolloutPolicyable)
                }
            }
            m.SetFeatureRolloutPolicies(res)
        }
        return nil
    }
    res["homeRealmDiscoveryPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HomeRealmDiscoveryPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HomeRealmDiscoveryPolicyable)
                }
            }
            m.SetHomeRealmDiscoveryPolicies(res)
        }
        return nil
    }
    res["identitySecurityDefaultsEnforcementPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentitySecurityDefaultsEnforcementPolicy(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["permissionGrantPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionGrantPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionGrantPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PermissionGrantPolicyable)
                }
            }
            m.SetPermissionGrantPolicies(res)
        }
        return nil
    }
    res["roleManagementPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleManagementPolicyable)
                }
            }
            m.SetRoleManagementPolicies(res)
        }
        return nil
    }
    res["roleManagementPolicyAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleManagementPolicyAssignmentable)
                }
            }
            m.SetRoleManagementPolicyAssignments(res)
        }
        return nil
    }
    res["tokenIssuancePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenIssuancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenIssuancePolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TokenIssuancePolicyable)
                }
            }
            m.SetTokenIssuancePolicies(res)
        }
        return nil
    }
    res["tokenLifetimePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenLifetimePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenLifetimePolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TokenLifetimePolicyable)
                }
            }
            m.SetTokenLifetimePolicies(res)
        }
        return nil
    }
    return res
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The policy to control Microsoft Entra authentication behavior for federated users.
// returns a []HomeRealmDiscoveryPolicyable when successful
func (m *PolicyRoot) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable) {
    return m.homeRealmDiscoveryPolicies
}
// GetIdentitySecurityDefaultsEnforcementPolicy gets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
// returns a IdentitySecurityDefaultsEnforcementPolicyable when successful
func (m *PolicyRoot) GetIdentitySecurityDefaultsEnforcementPolicy()(IdentitySecurityDefaultsEnforcementPolicyable) {
    return m.identitySecurityDefaultsEnforcementPolicy
}
// GetPermissionGrantPolicies gets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
// returns a []PermissionGrantPolicyable when successful
func (m *PolicyRoot) GetPermissionGrantPolicies()([]PermissionGrantPolicyable) {
    return m.permissionGrantPolicies
}
// GetRoleManagementPolicies gets the roleManagementPolicies property value. Specifies the various policies associated with scopes and roles.
// returns a []UnifiedRoleManagementPolicyable when successful
func (m *PolicyRoot) GetRoleManagementPolicies()([]UnifiedRoleManagementPolicyable) {
    return m.roleManagementPolicies
}
// GetRoleManagementPolicyAssignments gets the roleManagementPolicyAssignments property value. The assignment of a role management policy to a role definition object.
// returns a []UnifiedRoleManagementPolicyAssignmentable when successful
func (m *PolicyRoot) GetRoleManagementPolicyAssignments()([]UnifiedRoleManagementPolicyAssignmentable) {
    return m.roleManagementPolicyAssignments
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Microsoft Entra ID.
// returns a []TokenIssuancePolicyable when successful
func (m *PolicyRoot) GetTokenIssuancePolicies()([]TokenIssuancePolicyable) {
    return m.tokenIssuancePolicies
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Microsoft Entra ID.
// returns a []TokenLifetimePolicyable when successful
func (m *PolicyRoot) GetTokenLifetimePolicies()([]TokenLifetimePolicyable) {
    return m.tokenLifetimePolicies
}
// Serialize serializes information the current object
func (m *PolicyRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivityBasedTimeoutPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivityBasedTimeoutPolicies()))
        for i, v := range m.GetActivityBasedTimeoutPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("activityBasedTimeoutPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("adminConsentRequestPolicy", m.GetAdminConsentRequestPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetAppManagementPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppManagementPolicies()))
        for i, v := range m.GetAppManagementPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationFlowsPolicy", m.GetAuthenticationFlowsPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationMethodsPolicy", m.GetAuthenticationMethodsPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationStrengthPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationStrengthPolicies()))
        for i, v := range m.GetAuthenticationStrengthPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("authenticationStrengthPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authorizationPolicy", m.GetAuthorizationPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetClaimsMappingPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClaimsMappingPolicies()))
        for i, v := range m.GetClaimsMappingPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("claimsMappingPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConditionalAccessPolicies()))
        for i, v := range m.GetConditionalAccessPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("conditionalAccessPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("crossTenantAccessPolicy", m.GetCrossTenantAccessPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultAppManagementPolicy", m.GetDefaultAppManagementPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceRegistrationPolicy", m.GetDeviceRegistrationPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetFeatureRolloutPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFeatureRolloutPolicies()))
        for i, v := range m.GetFeatureRolloutPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHomeRealmDiscoveryPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHomeRealmDiscoveryPolicies()))
        for i, v := range m.GetHomeRealmDiscoveryPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("homeRealmDiscoveryPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identitySecurityDefaultsEnforcementPolicy", m.GetIdentitySecurityDefaultsEnforcementPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionGrantPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissionGrantPolicies()))
        for i, v := range m.GetPermissionGrantPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("permissionGrantPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleManagementPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleManagementPolicies()))
        for i, v := range m.GetRoleManagementPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleManagementPolicyAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleManagementPolicyAssignments()))
        for i, v := range m.GetRoleManagementPolicyAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleManagementPolicyAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenIssuancePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTokenIssuancePolicies()))
        for i, v := range m.GetTokenIssuancePolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenLifetimePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTokenLifetimePolicies()))
        for i, v := range m.GetTokenLifetimePolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("tokenLifetimePolicies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityBasedTimeoutPolicies sets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
func (m *PolicyRoot) SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicyable)() {
    m.activityBasedTimeoutPolicies = value
}
// SetAdminConsentRequestPolicy sets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *PolicyRoot) SetAdminConsentRequestPolicy(value AdminConsentRequestPolicyable)() {
    m.adminConsentRequestPolicy = value
}
// SetAppManagementPolicies sets the appManagementPolicies property value. The policies that enforce app management restrictions for specific applications and service principals, overriding the defaultAppManagementPolicy.
func (m *PolicyRoot) SetAppManagementPolicies(value []AppManagementPolicyable)() {
    m.appManagementPolicies = value
}
// SetAuthenticationFlowsPolicy sets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
func (m *PolicyRoot) SetAuthenticationFlowsPolicy(value AuthenticationFlowsPolicyable)() {
    m.authenticationFlowsPolicy = value
}
// SetAuthenticationMethodsPolicy sets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multifactor authentication (MFA) in Microsoft Entra ID.
func (m *PolicyRoot) SetAuthenticationMethodsPolicy(value AuthenticationMethodsPolicyable)() {
    m.authenticationMethodsPolicy = value
}
// SetAuthenticationStrengthPolicies sets the authenticationStrengthPolicies property value. The authentication method combinations that are to be used in scenarios defined by Microsoft Entra Conditional Access.
func (m *PolicyRoot) SetAuthenticationStrengthPolicies(value []AuthenticationStrengthPolicyable)() {
    m.authenticationStrengthPolicies = value
}
// SetAuthorizationPolicy sets the authorizationPolicy property value. The policy that controls Microsoft Entra authorization settings.
func (m *PolicyRoot) SetAuthorizationPolicy(value AuthorizationPolicyable)() {
    m.authorizationPolicy = value
}
// SetClaimsMappingPolicies sets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *PolicyRoot) SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)() {
    m.claimsMappingPolicies = value
}
// SetConditionalAccessPolicies sets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *PolicyRoot) SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)() {
    m.conditionalAccessPolicies = value
}
// SetCrossTenantAccessPolicy sets the crossTenantAccessPolicy property value. The custom rules that define an access scenario when interacting with external Microsoft Entra tenants.
func (m *PolicyRoot) SetCrossTenantAccessPolicy(value CrossTenantAccessPolicyable)() {
    m.crossTenantAccessPolicy = value
}
// SetDefaultAppManagementPolicy sets the defaultAppManagementPolicy property value. The tenant-wide policy that enforces app management restrictions for all applications and service principals.
func (m *PolicyRoot) SetDefaultAppManagementPolicy(value TenantAppManagementPolicyable)() {
    m.defaultAppManagementPolicy = value
}
// SetDeviceRegistrationPolicy sets the deviceRegistrationPolicy property value. The deviceRegistrationPolicy property
func (m *PolicyRoot) SetDeviceRegistrationPolicy(value DeviceRegistrationPolicyable)() {
    m.deviceRegistrationPolicy = value
}
// SetFeatureRolloutPolicies sets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *PolicyRoot) SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)() {
    m.featureRolloutPolicies = value
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The policy to control Microsoft Entra authentication behavior for federated users.
func (m *PolicyRoot) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)() {
    m.homeRealmDiscoveryPolicies = value
}
// SetIdentitySecurityDefaultsEnforcementPolicy sets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *PolicyRoot) SetIdentitySecurityDefaultsEnforcementPolicy(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    m.identitySecurityDefaultsEnforcementPolicy = value
}
// SetPermissionGrantPolicies sets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *PolicyRoot) SetPermissionGrantPolicies(value []PermissionGrantPolicyable)() {
    m.permissionGrantPolicies = value
}
// SetRoleManagementPolicies sets the roleManagementPolicies property value. Specifies the various policies associated with scopes and roles.
func (m *PolicyRoot) SetRoleManagementPolicies(value []UnifiedRoleManagementPolicyable)() {
    m.roleManagementPolicies = value
}
// SetRoleManagementPolicyAssignments sets the roleManagementPolicyAssignments property value. The assignment of a role management policy to a role definition object.
func (m *PolicyRoot) SetRoleManagementPolicyAssignments(value []UnifiedRoleManagementPolicyAssignmentable)() {
    m.roleManagementPolicyAssignments = value
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Microsoft Entra ID.
func (m *PolicyRoot) SetTokenIssuancePolicies(value []TokenIssuancePolicyable)() {
    m.tokenIssuancePolicies = value
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Microsoft Entra ID.
func (m *PolicyRoot) SetTokenLifetimePolicies(value []TokenLifetimePolicyable)() {
    m.tokenLifetimePolicies = value
}
type PolicyRootable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicyable)
    GetAdminConsentRequestPolicy()(AdminConsentRequestPolicyable)
    GetAppManagementPolicies()([]AppManagementPolicyable)
    GetAuthenticationFlowsPolicy()(AuthenticationFlowsPolicyable)
    GetAuthenticationMethodsPolicy()(AuthenticationMethodsPolicyable)
    GetAuthenticationStrengthPolicies()([]AuthenticationStrengthPolicyable)
    GetAuthorizationPolicy()(AuthorizationPolicyable)
    GetClaimsMappingPolicies()([]ClaimsMappingPolicyable)
    GetConditionalAccessPolicies()([]ConditionalAccessPolicyable)
    GetCrossTenantAccessPolicy()(CrossTenantAccessPolicyable)
    GetDefaultAppManagementPolicy()(TenantAppManagementPolicyable)
    GetDeviceRegistrationPolicy()(DeviceRegistrationPolicyable)
    GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable)
    GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable)
    GetIdentitySecurityDefaultsEnforcementPolicy()(IdentitySecurityDefaultsEnforcementPolicyable)
    GetPermissionGrantPolicies()([]PermissionGrantPolicyable)
    GetRoleManagementPolicies()([]UnifiedRoleManagementPolicyable)
    GetRoleManagementPolicyAssignments()([]UnifiedRoleManagementPolicyAssignmentable)
    GetTokenIssuancePolicies()([]TokenIssuancePolicyable)
    GetTokenLifetimePolicies()([]TokenLifetimePolicyable)
    SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicyable)()
    SetAdminConsentRequestPolicy(value AdminConsentRequestPolicyable)()
    SetAppManagementPolicies(value []AppManagementPolicyable)()
    SetAuthenticationFlowsPolicy(value AuthenticationFlowsPolicyable)()
    SetAuthenticationMethodsPolicy(value AuthenticationMethodsPolicyable)()
    SetAuthenticationStrengthPolicies(value []AuthenticationStrengthPolicyable)()
    SetAuthorizationPolicy(value AuthorizationPolicyable)()
    SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)()
    SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)()
    SetCrossTenantAccessPolicy(value CrossTenantAccessPolicyable)()
    SetDefaultAppManagementPolicy(value TenantAppManagementPolicyable)()
    SetDeviceRegistrationPolicy(value DeviceRegistrationPolicyable)()
    SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)()
    SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)()
    SetIdentitySecurityDefaultsEnforcementPolicy(value IdentitySecurityDefaultsEnforcementPolicyable)()
    SetPermissionGrantPolicies(value []PermissionGrantPolicyable)()
    SetRoleManagementPolicies(value []UnifiedRoleManagementPolicyable)()
    SetRoleManagementPolicyAssignments(value []UnifiedRoleManagementPolicyAssignmentable)()
    SetTokenIssuancePolicies(value []TokenIssuancePolicyable)()
    SetTokenLifetimePolicies(value []TokenLifetimePolicyable)()
}
