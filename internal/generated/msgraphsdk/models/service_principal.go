package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServicePrincipal struct {
    DirectoryObject
    // true if the service principal account is enabled; otherwise, false. If set to false, then no users are able to sign in to this app, even if they're assigned to it. Supports $filter (eq, ne, not, in).
    accountEnabled *bool
    // Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This lets services like Microsoft 365 call the application in the context of a document the user is working on.
    addIns []AddInable
    // Used to retrieve service principals by subscription, identify resource group and full resource IDs for managed identities. Supports $filter (eq, not, ge, le, startsWith).
    alternativeNames []string
    // The description exposed by the associated application.
    appDescription *string
    // The display name exposed by the associated application. Maximum length is 256 characters.
    appDisplayName *string
    // The unique identifier for the associated application (its appId property). Alternate key. Supports $filter (eq, ne, not, in, startsWith).
    appId *string
    // Unique identifier of the applicationTemplate. Supports $filter (eq, not, ne). Read-only. null if the service principal wasn't created from an application template.
    applicationTemplateId *string
    // The appManagementPolicy applied to this application.
    appManagementPolicies []AppManagementPolicyable
    // Contains the tenant ID where the application is registered. This is applicable only to service principals backed by applications. Supports $filter (eq, ne, NOT, ge, le).
    appOwnerOrganizationId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // App role assignments for this app or service, granted to users, groups, and other service principals. Supports $expand.
    appRoleAssignedTo []AppRoleAssignmentable
    // Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
    appRoleAssignmentRequired *bool
    // App role assignment for another app or service, granted to this service principal. Supports $expand.
    appRoleAssignments []AppRoleAssignmentable
    // The roles exposed by the application that's linked to this service principal. For more information, see the appRoles property definition on the application entity. Not nullable.
    appRoles []AppRoleable
    // The claimsMappingPolicies assigned to this service principal. Supports $expand.
    claimsMappingPolicies []ClaimsMappingPolicyable
    // Directory objects created by this service principal. Read-only. Nullable.
    createdObjects []DirectoryObjectable
    // An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith). Filter value is case sensitive. To read this property, the calling app must be assigned the CustomSecAttributeAssignment.Read.All permission. To write this property, the calling app must be assigned the CustomSecAttributeAssignment.ReadWrite.All permissions. To read or write this property in delegated scenarios, the admin must be assigned the Attribute Assignment Administrator role.
    customSecurityAttributes CustomSecurityAttributeValueable
    // The delegatedPermissionClassifications property
    delegatedPermissionClassifications []DelegatedPermissionClassificationable
    // Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps displays the application description in this field. The maximum allowed size is 1,024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
    description *string
    // Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
    disabledByMicrosoftStatus *string
    // The display name for the service principal. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderby.
    displayName *string
    // The endpoints property
    endpoints []Endpointable
    // Federated identities for a specific type of service principal - managed identity. Supports $expand and $filter (/$count eq 0, /$count ne 0).
    federatedIdentityCredentials []FederatedIdentityCredentialable
    // Home page or landing page of the application.
    homepage *string
    // The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicyable
    // Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Microsoft Entra apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
    info InformationalUrlable
    // The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, not, ge, le).
    keyCredentials []KeyCredentialable
    // Specifies the URL where the service provider redirects the user to Microsoft Entra ID to authenticate. Microsoft Entra ID uses the URL to launch the application from Microsoft 365 or the Microsoft Entra My Apps. When blank, Microsoft Entra ID performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Microsoft Entra My Apps, or the Microsoft Entra SSO URL.
    loginUrl *string
    // Specifies the URL that the Microsoft's authorization service uses to sign out a user using OpenID Connect front-channel, back-channel, or SAML sign out protocols.
    logoutUrl *string
    // Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObjectable
    // Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1,024 characters.
    notes *string
    // Specifies the list of email addresses where Microsoft Entra ID sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Microsoft Entra Gallery applications.
    notificationEmailAddresses []string
    // Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
    oauth2PermissionGrants []OAuth2PermissionGrantable
    // The delegated permissions exposed by the application. For more information, see the oauth2PermissionScopes property on the application entity's api property. Not nullable.
    oauth2PermissionScopes []PermissionScopeable
    // Directory objects that this service principal owns. Read-only. Nullable. Supports $expand, $select nested in $expand, and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
    ownedObjects []DirectoryObjectable
    // Directory objects that are owners of this servicePrincipal. The owners are a set of nonadmin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand, $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1), and $select nested in $expand.
    owners []DirectoryObjectable
    // The collection of password credentials associated with the application. Not nullable.
    passwordCredentials []PasswordCredentialable
    // Specifies the single sign-on mode configured for this application. Microsoft Entra ID uses the preferred single sign-on mode to launch the application from Microsoft 365 or the My Apps portal. The supported values are password, saml, notSupported, and oidc. Note: This field might be null for older SAML apps and for OIDC applications where it isn't set automatically.
    preferredSingleSignOnMode *string
    // This property can be used on SAML applications (apps that have preferredSingleSignOnMode set to saml) to control which certificate is used to sign the SAML responses. For applications that aren't SAML, don't write or otherwise rely on this property.
    preferredTokenSigningKeyThumbprint *string
    // The remoteDesktopSecurityConfiguration object applied to this service principal. Supports $filter (eq) for isRemoteDesktopProtocolEnabled property.
    remoteDesktopSecurityConfiguration RemoteDesktopSecurityConfigurationable
    // The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
    replyUrls []string
    // The resource-specific application permissions exposed by this application. Currently, resource-specific permissions are only supported for Teams apps accessing to specific chats and teams using Microsoft Graph. Read-only.
    resourceSpecificApplicationPermissions []ResourceSpecificPermissionable
    // The collection for settings related to saml single sign-on.
    samlSingleSignOnSettings SamlSingleSignOnSettingsable
    // Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Microsoft Entra ID. For example,Client apps can specify a resource URI that is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, not, ge, le, startsWith).
    servicePrincipalNames []string
    // Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Microsoft Entra ID internally. The servicePrincipalType property can be set to three different values: Application - A service principal that represents an application or service. The appId property identifies the associated app registration, and matches the appId of an application, possibly from a different tenant. If the associated app registration is missing, tokens aren't issued for the service principal.ManagedIdentity - A service principal that represents a managed identity. Service principals representing managed identities can be granted access and permissions, but can't be updated or modified directly.Legacy - A service principal that represents an app created before app registrations, or through legacy experiences. A legacy service principal can have credentials, service principal names, reply URLs, and other properties that are editable by an authorized user, but doesn't have an associated app registration. The appId value doesn't associate the service principal with an app registration. The service principal can only be used in the tenant where it was created.SocialIdp - For internal use.
    servicePrincipalType *string
    // Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization's Microsoft Entra tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization's Microsoft Entra tenant (multitenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization's Microsoft Entra tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
    signInAudience *string
    // Represents the capability for Microsoft Entra identity synchronization through the Microsoft Graph API.
    synchronization Synchronizationable
    // Custom strings that can be used to categorize and identify the service principal. Not nullable. The value is the union of strings set here and on the associated application entity's tags property.Supports $filter (eq, not, ge, le, startsWith).
    tags []string
    // Specifies the keyId of a public key from the keyCredentials collection. When configured, Microsoft Entra ID issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
    tokenEncryptionKeyId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The tokenIssuancePolicies assigned to this service principal.
    tokenIssuancePolicies []TokenIssuancePolicyable
    // The tokenLifetimePolicies assigned to this service principal.
    tokenLifetimePolicies []TokenLifetimePolicyable
    // The transitiveMemberOf property
    transitiveMemberOf []DirectoryObjectable
    // Specifies the verified publisher of the application that's linked to this service principal.
    verifiedPublisher VerifiedPublisherable
}
// NewServicePrincipal instantiates a new ServicePrincipal and sets the default values.
func NewServicePrincipal()(*ServicePrincipal) {
    m := &ServicePrincipal{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.servicePrincipal"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateServicePrincipalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServicePrincipalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipal(), nil
}
// GetAccountEnabled gets the accountEnabled property value. true if the service principal account is enabled; otherwise, false. If set to false, then no users are able to sign in to this app, even if they're assigned to it. Supports $filter (eq, ne, not, in).
// returns a *bool when successful
func (m *ServicePrincipal) GetAccountEnabled()(*bool) {
    return m.accountEnabled
}
// GetAddIns gets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This lets services like Microsoft 365 call the application in the context of a document the user is working on.
// returns a []AddInable when successful
func (m *ServicePrincipal) GetAddIns()([]AddInable) {
    return m.addIns
}
// GetAlternativeNames gets the alternativeNames property value. Used to retrieve service principals by subscription, identify resource group and full resource IDs for managed identities. Supports $filter (eq, not, ge, le, startsWith).
// returns a []string when successful
func (m *ServicePrincipal) GetAlternativeNames()([]string) {
    return m.alternativeNames
}
// GetAppDescription gets the appDescription property value. The description exposed by the associated application.
// returns a *string when successful
func (m *ServicePrincipal) GetAppDescription()(*string) {
    return m.appDescription
}
// GetAppDisplayName gets the appDisplayName property value. The display name exposed by the associated application. Maximum length is 256 characters.
// returns a *string when successful
func (m *ServicePrincipal) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppId gets the appId property value. The unique identifier for the associated application (its appId property). Alternate key. Supports $filter (eq, ne, not, in, startsWith).
// returns a *string when successful
func (m *ServicePrincipal) GetAppId()(*string) {
    return m.appId
}
// GetApplicationTemplateId gets the applicationTemplateId property value. Unique identifier of the applicationTemplate. Supports $filter (eq, not, ne). Read-only. null if the service principal wasn't created from an application template.
// returns a *string when successful
func (m *ServicePrincipal) GetApplicationTemplateId()(*string) {
    return m.applicationTemplateId
}
// GetAppManagementPolicies gets the appManagementPolicies property value. The appManagementPolicy applied to this application.
// returns a []AppManagementPolicyable when successful
func (m *ServicePrincipal) GetAppManagementPolicies()([]AppManagementPolicyable) {
    return m.appManagementPolicies
}
// GetAppOwnerOrganizationId gets the appOwnerOrganizationId property value. Contains the tenant ID where the application is registered. This is applicable only to service principals backed by applications. Supports $filter (eq, ne, NOT, ge, le).
// returns a *UUID when successful
func (m *ServicePrincipal) GetAppOwnerOrganizationId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.appOwnerOrganizationId
}
// GetAppRoleAssignedTo gets the appRoleAssignedTo property value. App role assignments for this app or service, granted to users, groups, and other service principals. Supports $expand.
// returns a []AppRoleAssignmentable when successful
func (m *ServicePrincipal) GetAppRoleAssignedTo()([]AppRoleAssignmentable) {
    return m.appRoleAssignedTo
}
// GetAppRoleAssignmentRequired gets the appRoleAssignmentRequired property value. Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
// returns a *bool when successful
func (m *ServicePrincipal) GetAppRoleAssignmentRequired()(*bool) {
    return m.appRoleAssignmentRequired
}
// GetAppRoleAssignments gets the appRoleAssignments property value. App role assignment for another app or service, granted to this service principal. Supports $expand.
// returns a []AppRoleAssignmentable when successful
func (m *ServicePrincipal) GetAppRoleAssignments()([]AppRoleAssignmentable) {
    return m.appRoleAssignments
}
// GetAppRoles gets the appRoles property value. The roles exposed by the application that's linked to this service principal. For more information, see the appRoles property definition on the application entity. Not nullable.
// returns a []AppRoleable when successful
func (m *ServicePrincipal) GetAppRoles()([]AppRoleable) {
    return m.appRoles
}
// GetClaimsMappingPolicies gets the claimsMappingPolicies property value. The claimsMappingPolicies assigned to this service principal. Supports $expand.
// returns a []ClaimsMappingPolicyable when successful
func (m *ServicePrincipal) GetClaimsMappingPolicies()([]ClaimsMappingPolicyable) {
    return m.claimsMappingPolicies
}
// GetCreatedObjects gets the createdObjects property value. Directory objects created by this service principal. Read-only. Nullable.
// returns a []DirectoryObjectable when successful
func (m *ServicePrincipal) GetCreatedObjects()([]DirectoryObjectable) {
    return m.createdObjects
}
// GetCustomSecurityAttributes gets the customSecurityAttributes property value. An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith). Filter value is case sensitive. To read this property, the calling app must be assigned the CustomSecAttributeAssignment.Read.All permission. To write this property, the calling app must be assigned the CustomSecAttributeAssignment.ReadWrite.All permissions. To read or write this property in delegated scenarios, the admin must be assigned the Attribute Assignment Administrator role.
// returns a CustomSecurityAttributeValueable when successful
func (m *ServicePrincipal) GetCustomSecurityAttributes()(CustomSecurityAttributeValueable) {
    return m.customSecurityAttributes
}
// GetDelegatedPermissionClassifications gets the delegatedPermissionClassifications property value. The delegatedPermissionClassifications property
// returns a []DelegatedPermissionClassificationable when successful
func (m *ServicePrincipal) GetDelegatedPermissionClassifications()([]DelegatedPermissionClassificationable) {
    return m.delegatedPermissionClassifications
}
// GetDescription gets the description property value. Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps displays the application description in this field. The maximum allowed size is 1,024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
// returns a *string when successful
func (m *ServicePrincipal) GetDescription()(*string) {
    return m.description
}
// GetDisabledByMicrosoftStatus gets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
// returns a *string when successful
func (m *ServicePrincipal) GetDisabledByMicrosoftStatus()(*string) {
    return m.disabledByMicrosoftStatus
}
// GetDisplayName gets the displayName property value. The display name for the service principal. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderby.
// returns a *string when successful
func (m *ServicePrincipal) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndpoints gets the endpoints property value. The endpoints property
// returns a []Endpointable when successful
func (m *ServicePrincipal) GetEndpoints()([]Endpointable) {
    return m.endpoints
}
// GetFederatedIdentityCredentials gets the federatedIdentityCredentials property value. Federated identities for a specific type of service principal - managed identity. Supports $expand and $filter (/$count eq 0, /$count ne 0).
// returns a []FederatedIdentityCredentialable when successful
func (m *ServicePrincipal) GetFederatedIdentityCredentials()([]FederatedIdentityCredentialable) {
    return m.federatedIdentityCredentials
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServicePrincipal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["accountEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountEnabled(val)
        }
        return nil
    }
    res["addIns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAddInFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AddInable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AddInable)
                }
            }
            m.SetAddIns(res)
        }
        return nil
    }
    res["alternativeNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAlternativeNames(res)
        }
        return nil
    }
    res["appDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDescription(val)
        }
        return nil
    }
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["applicationTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationTemplateId(val)
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
    res["appOwnerOrganizationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppOwnerOrganizationId(val)
        }
        return nil
    }
    res["appRoleAssignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppRoleAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppRoleAssignmentable)
                }
            }
            m.SetAppRoleAssignedTo(res)
        }
        return nil
    }
    res["appRoleAssignmentRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppRoleAssignmentRequired(val)
        }
        return nil
    }
    res["appRoleAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppRoleAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppRoleAssignmentable)
                }
            }
            m.SetAppRoleAssignments(res)
        }
        return nil
    }
    res["appRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppRoleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppRoleable)
                }
            }
            m.SetAppRoles(res)
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
    res["createdObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
            }
            m.SetCreatedObjects(res)
        }
        return nil
    }
    res["customSecurityAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomSecurityAttributeValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomSecurityAttributes(val.(CustomSecurityAttributeValueable))
        }
        return nil
    }
    res["delegatedPermissionClassifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedPermissionClassificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedPermissionClassificationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DelegatedPermissionClassificationable)
                }
            }
            m.SetDelegatedPermissionClassifications(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["disabledByMicrosoftStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabledByMicrosoftStatus(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endpoints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEndpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Endpointable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Endpointable)
                }
            }
            m.SetEndpoints(res)
        }
        return nil
    }
    res["federatedIdentityCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFederatedIdentityCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FederatedIdentityCredentialable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FederatedIdentityCredentialable)
                }
            }
            m.SetFederatedIdentityCredentials(res)
        }
        return nil
    }
    res["homepage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomepage(val)
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
    res["info"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInformationalUrlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfo(val.(InformationalUrlable))
        }
        return nil
    }
    res["keyCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyCredentialable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyCredentialable)
                }
            }
            m.SetKeyCredentials(res)
        }
        return nil
    }
    res["loginUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginUrl(val)
        }
        return nil
    }
    res["logoutUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogoutUrl(val)
        }
        return nil
    }
    res["memberOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
            }
            m.SetMemberOf(res)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["notificationEmailAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetNotificationEmailAddresses(res)
        }
        return nil
    }
    res["oauth2PermissionGrants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOAuth2PermissionGrantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OAuth2PermissionGrantable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OAuth2PermissionGrantable)
                }
            }
            m.SetOauth2PermissionGrants(res)
        }
        return nil
    }
    res["oauth2PermissionScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PermissionScopeable)
                }
            }
            m.SetOauth2PermissionScopes(res)
        }
        return nil
    }
    res["ownedObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
            }
            m.SetOwnedObjects(res)
        }
        return nil
    }
    res["owners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
            }
            m.SetOwners(res)
        }
        return nil
    }
    res["passwordCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePasswordCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordCredentialable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PasswordCredentialable)
                }
            }
            m.SetPasswordCredentials(res)
        }
        return nil
    }
    res["preferredSingleSignOnMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredSingleSignOnMode(val)
        }
        return nil
    }
    res["preferredTokenSigningKeyThumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredTokenSigningKeyThumbprint(val)
        }
        return nil
    }
    res["remoteDesktopSecurityConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRemoteDesktopSecurityConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteDesktopSecurityConfiguration(val.(RemoteDesktopSecurityConfigurationable))
        }
        return nil
    }
    res["replyUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetReplyUrls(res)
        }
        return nil
    }
    res["resourceSpecificApplicationPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceSpecificPermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceSpecificPermissionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ResourceSpecificPermissionable)
                }
            }
            m.SetResourceSpecificApplicationPermissions(res)
        }
        return nil
    }
    res["samlSingleSignOnSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSamlSingleSignOnSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSamlSingleSignOnSettings(val.(SamlSingleSignOnSettingsable))
        }
        return nil
    }
    res["servicePrincipalNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetServicePrincipalNames(res)
        }
        return nil
    }
    res["servicePrincipalType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalType(val)
        }
        return nil
    }
    res["signInAudience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInAudience(val)
        }
        return nil
    }
    res["synchronization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronization(val.(Synchronizationable))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTags(res)
        }
        return nil
    }
    res["tokenEncryptionKeyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenEncryptionKeyId(val)
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
    res["transitiveMemberOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
            }
            m.SetTransitiveMemberOf(res)
        }
        return nil
    }
    res["verifiedPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVerifiedPublisherFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedPublisher(val.(VerifiedPublisherable))
        }
        return nil
    }
    return res
}
// GetHomepage gets the homepage property value. Home page or landing page of the application.
// returns a *string when successful
func (m *ServicePrincipal) GetHomepage()(*string) {
    return m.homepage
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
// returns a []HomeRealmDiscoveryPolicyable when successful
func (m *ServicePrincipal) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable) {
    return m.homeRealmDiscoveryPolicies
}
// GetInfo gets the info property value. Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Microsoft Entra apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
// returns a InformationalUrlable when successful
func (m *ServicePrincipal) GetInfo()(InformationalUrlable) {
    return m.info
}
// GetKeyCredentials gets the keyCredentials property value. The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, not, ge, le).
// returns a []KeyCredentialable when successful
func (m *ServicePrincipal) GetKeyCredentials()([]KeyCredentialable) {
    return m.keyCredentials
}
// GetLoginUrl gets the loginUrl property value. Specifies the URL where the service provider redirects the user to Microsoft Entra ID to authenticate. Microsoft Entra ID uses the URL to launch the application from Microsoft 365 or the Microsoft Entra My Apps. When blank, Microsoft Entra ID performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Microsoft Entra My Apps, or the Microsoft Entra SSO URL.
// returns a *string when successful
func (m *ServicePrincipal) GetLoginUrl()(*string) {
    return m.loginUrl
}
// GetLogoutUrl gets the logoutUrl property value. Specifies the URL that the Microsoft's authorization service uses to sign out a user using OpenID Connect front-channel, back-channel, or SAML sign out protocols.
// returns a *string when successful
func (m *ServicePrincipal) GetLogoutUrl()(*string) {
    return m.logoutUrl
}
// GetMemberOf gets the memberOf property value. Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
// returns a []DirectoryObjectable when successful
func (m *ServicePrincipal) GetMemberOf()([]DirectoryObjectable) {
    return m.memberOf
}
// GetNotes gets the notes property value. Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1,024 characters.
// returns a *string when successful
func (m *ServicePrincipal) GetNotes()(*string) {
    return m.notes
}
// GetNotificationEmailAddresses gets the notificationEmailAddresses property value. Specifies the list of email addresses where Microsoft Entra ID sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Microsoft Entra Gallery applications.
// returns a []string when successful
func (m *ServicePrincipal) GetNotificationEmailAddresses()([]string) {
    return m.notificationEmailAddresses
}
// GetOauth2PermissionGrants gets the oauth2PermissionGrants property value. Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
// returns a []OAuth2PermissionGrantable when successful
func (m *ServicePrincipal) GetOauth2PermissionGrants()([]OAuth2PermissionGrantable) {
    return m.oauth2PermissionGrants
}
// GetOauth2PermissionScopes gets the oauth2PermissionScopes property value. The delegated permissions exposed by the application. For more information, see the oauth2PermissionScopes property on the application entity's api property. Not nullable.
// returns a []PermissionScopeable when successful
func (m *ServicePrincipal) GetOauth2PermissionScopes()([]PermissionScopeable) {
    return m.oauth2PermissionScopes
}
// GetOwnedObjects gets the ownedObjects property value. Directory objects that this service principal owns. Read-only. Nullable. Supports $expand, $select nested in $expand, and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
// returns a []DirectoryObjectable when successful
func (m *ServicePrincipal) GetOwnedObjects()([]DirectoryObjectable) {
    return m.ownedObjects
}
// GetOwners gets the owners property value. Directory objects that are owners of this servicePrincipal. The owners are a set of nonadmin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand, $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1), and $select nested in $expand.
// returns a []DirectoryObjectable when successful
func (m *ServicePrincipal) GetOwners()([]DirectoryObjectable) {
    return m.owners
}
// GetPasswordCredentials gets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
// returns a []PasswordCredentialable when successful
func (m *ServicePrincipal) GetPasswordCredentials()([]PasswordCredentialable) {
    return m.passwordCredentials
}
// GetPreferredSingleSignOnMode gets the preferredSingleSignOnMode property value. Specifies the single sign-on mode configured for this application. Microsoft Entra ID uses the preferred single sign-on mode to launch the application from Microsoft 365 or the My Apps portal. The supported values are password, saml, notSupported, and oidc. Note: This field might be null for older SAML apps and for OIDC applications where it isn't set automatically.
// returns a *string when successful
func (m *ServicePrincipal) GetPreferredSingleSignOnMode()(*string) {
    return m.preferredSingleSignOnMode
}
// GetPreferredTokenSigningKeyThumbprint gets the preferredTokenSigningKeyThumbprint property value. This property can be used on SAML applications (apps that have preferredSingleSignOnMode set to saml) to control which certificate is used to sign the SAML responses. For applications that aren't SAML, don't write or otherwise rely on this property.
// returns a *string when successful
func (m *ServicePrincipal) GetPreferredTokenSigningKeyThumbprint()(*string) {
    return m.preferredTokenSigningKeyThumbprint
}
// GetRemoteDesktopSecurityConfiguration gets the remoteDesktopSecurityConfiguration property value. The remoteDesktopSecurityConfiguration object applied to this service principal. Supports $filter (eq) for isRemoteDesktopProtocolEnabled property.
// returns a RemoteDesktopSecurityConfigurationable when successful
func (m *ServicePrincipal) GetRemoteDesktopSecurityConfiguration()(RemoteDesktopSecurityConfigurationable) {
    return m.remoteDesktopSecurityConfiguration
}
// GetReplyUrls gets the replyUrls property value. The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
// returns a []string when successful
func (m *ServicePrincipal) GetReplyUrls()([]string) {
    return m.replyUrls
}
// GetResourceSpecificApplicationPermissions gets the resourceSpecificApplicationPermissions property value. The resource-specific application permissions exposed by this application. Currently, resource-specific permissions are only supported for Teams apps accessing to specific chats and teams using Microsoft Graph. Read-only.
// returns a []ResourceSpecificPermissionable when successful
func (m *ServicePrincipal) GetResourceSpecificApplicationPermissions()([]ResourceSpecificPermissionable) {
    return m.resourceSpecificApplicationPermissions
}
// GetSamlSingleSignOnSettings gets the samlSingleSignOnSettings property value. The collection for settings related to saml single sign-on.
// returns a SamlSingleSignOnSettingsable when successful
func (m *ServicePrincipal) GetSamlSingleSignOnSettings()(SamlSingleSignOnSettingsable) {
    return m.samlSingleSignOnSettings
}
// GetServicePrincipalNames gets the servicePrincipalNames property value. Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Microsoft Entra ID. For example,Client apps can specify a resource URI that is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, not, ge, le, startsWith).
// returns a []string when successful
func (m *ServicePrincipal) GetServicePrincipalNames()([]string) {
    return m.servicePrincipalNames
}
// GetServicePrincipalType gets the servicePrincipalType property value. Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Microsoft Entra ID internally. The servicePrincipalType property can be set to three different values: Application - A service principal that represents an application or service. The appId property identifies the associated app registration, and matches the appId of an application, possibly from a different tenant. If the associated app registration is missing, tokens aren't issued for the service principal.ManagedIdentity - A service principal that represents a managed identity. Service principals representing managed identities can be granted access and permissions, but can't be updated or modified directly.Legacy - A service principal that represents an app created before app registrations, or through legacy experiences. A legacy service principal can have credentials, service principal names, reply URLs, and other properties that are editable by an authorized user, but doesn't have an associated app registration. The appId value doesn't associate the service principal with an app registration. The service principal can only be used in the tenant where it was created.SocialIdp - For internal use.
// returns a *string when successful
func (m *ServicePrincipal) GetServicePrincipalType()(*string) {
    return m.servicePrincipalType
}
// GetSignInAudience gets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization's Microsoft Entra tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization's Microsoft Entra tenant (multitenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization's Microsoft Entra tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
// returns a *string when successful
func (m *ServicePrincipal) GetSignInAudience()(*string) {
    return m.signInAudience
}
// GetSynchronization gets the synchronization property value. Represents the capability for Microsoft Entra identity synchronization through the Microsoft Graph API.
// returns a Synchronizationable when successful
func (m *ServicePrincipal) GetSynchronization()(Synchronizationable) {
    return m.synchronization
}
// GetTags gets the tags property value. Custom strings that can be used to categorize and identify the service principal. Not nullable. The value is the union of strings set here and on the associated application entity's tags property.Supports $filter (eq, not, ge, le, startsWith).
// returns a []string when successful
func (m *ServicePrincipal) GetTags()([]string) {
    return m.tags
}
// GetTokenEncryptionKeyId gets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Microsoft Entra ID issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
// returns a *UUID when successful
func (m *ServicePrincipal) GetTokenEncryptionKeyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.tokenEncryptionKeyId
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The tokenIssuancePolicies assigned to this service principal.
// returns a []TokenIssuancePolicyable when successful
func (m *ServicePrincipal) GetTokenIssuancePolicies()([]TokenIssuancePolicyable) {
    return m.tokenIssuancePolicies
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this service principal.
// returns a []TokenLifetimePolicyable when successful
func (m *ServicePrincipal) GetTokenLifetimePolicies()([]TokenLifetimePolicyable) {
    return m.tokenLifetimePolicies
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. The transitiveMemberOf property
// returns a []DirectoryObjectable when successful
func (m *ServicePrincipal) GetTransitiveMemberOf()([]DirectoryObjectable) {
    return m.transitiveMemberOf
}
// GetVerifiedPublisher gets the verifiedPublisher property value. Specifies the verified publisher of the application that's linked to this service principal.
// returns a VerifiedPublisherable when successful
func (m *ServicePrincipal) GetVerifiedPublisher()(VerifiedPublisherable) {
    return m.verifiedPublisher
}
// Serialize serializes information the current object
func (m *ServicePrincipal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetAddIns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddIns()))
        for i, v := range m.GetAddIns() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("addIns", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAlternativeNames() != nil {
        err = writer.WriteCollectionOfStringValues("alternativeNames", m.GetAlternativeNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDescription", m.GetAppDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("applicationTemplateId", m.GetApplicationTemplateId())
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
        err = writer.WriteUUIDValue("appOwnerOrganizationId", m.GetAppOwnerOrganizationId())
        if err != nil {
            return err
        }
    }
    if m.GetAppRoleAssignedTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppRoleAssignedTo()))
        for i, v := range m.GetAppRoleAssignedTo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appRoleAssignedTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appRoleAssignmentRequired", m.GetAppRoleAssignmentRequired())
        if err != nil {
            return err
        }
    }
    if m.GetAppRoleAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppRoleAssignments()))
        for i, v := range m.GetAppRoleAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appRoleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppRoles()))
        for i, v := range m.GetAppRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appRoles", cast)
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
    if m.GetCreatedObjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCreatedObjects()))
        for i, v := range m.GetCreatedObjects() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("createdObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("customSecurityAttributes", m.GetCustomSecurityAttributes())
        if err != nil {
            return err
        }
    }
    if m.GetDelegatedPermissionClassifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDelegatedPermissionClassifications()))
        for i, v := range m.GetDelegatedPermissionClassifications() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("delegatedPermissionClassifications", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("disabledByMicrosoftStatus", m.GetDisabledByMicrosoftStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEndpoints() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEndpoints()))
        for i, v := range m.GetEndpoints() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("endpoints", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFederatedIdentityCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFederatedIdentityCredentials()))
        for i, v := range m.GetFederatedIdentityCredentials() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("federatedIdentityCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("homepage", m.GetHomepage())
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
        err = writer.WriteObjectValue("info", m.GetInfo())
        if err != nil {
            return err
        }
    }
    if m.GetKeyCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKeyCredentials()))
        for i, v := range m.GetKeyCredentials() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("keyCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("loginUrl", m.GetLoginUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("logoutUrl", m.GetLogoutUrl())
        if err != nil {
            return err
        }
    }
    if m.GetMemberOf() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationEmailAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("notificationEmailAddresses", m.GetNotificationEmailAddresses())
        if err != nil {
            return err
        }
    }
    if m.GetOauth2PermissionGrants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOauth2PermissionGrants()))
        for i, v := range m.GetOauth2PermissionGrants() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("oauth2PermissionGrants", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOauth2PermissionScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOauth2PermissionScopes()))
        for i, v := range m.GetOauth2PermissionScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("oauth2PermissionScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwnedObjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOwnedObjects()))
        for i, v := range m.GetOwnedObjects() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("ownedObjects", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOwners()))
        for i, v := range m.GetOwners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("owners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPasswordCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPasswordCredentials()))
        for i, v := range m.GetPasswordCredentials() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("passwordCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredSingleSignOnMode", m.GetPreferredSingleSignOnMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredTokenSigningKeyThumbprint", m.GetPreferredTokenSigningKeyThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("remoteDesktopSecurityConfiguration", m.GetRemoteDesktopSecurityConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetReplyUrls() != nil {
        err = writer.WriteCollectionOfStringValues("replyUrls", m.GetReplyUrls())
        if err != nil {
            return err
        }
    }
    if m.GetResourceSpecificApplicationPermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceSpecificApplicationPermissions()))
        for i, v := range m.GetResourceSpecificApplicationPermissions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceSpecificApplicationPermissions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("samlSingleSignOnSettings", m.GetSamlSingleSignOnSettings())
        if err != nil {
            return err
        }
    }
    if m.GetServicePrincipalNames() != nil {
        err = writer.WriteCollectionOfStringValues("servicePrincipalNames", m.GetServicePrincipalNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalType", m.GetServicePrincipalType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInAudience", m.GetSignInAudience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("synchronization", m.GetSynchronization())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("tokenEncryptionKeyId", m.GetTokenEncryptionKeyId())
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
    if m.GetTransitiveMemberOf() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("verifiedPublisher", m.GetVerifiedPublisher())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountEnabled sets the accountEnabled property value. true if the service principal account is enabled; otherwise, false. If set to false, then no users are able to sign in to this app, even if they're assigned to it. Supports $filter (eq, ne, not, in).
func (m *ServicePrincipal) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// SetAddIns sets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This lets services like Microsoft 365 call the application in the context of a document the user is working on.
func (m *ServicePrincipal) SetAddIns(value []AddInable)() {
    m.addIns = value
}
// SetAlternativeNames sets the alternativeNames property value. Used to retrieve service principals by subscription, identify resource group and full resource IDs for managed identities. Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) SetAlternativeNames(value []string)() {
    m.alternativeNames = value
}
// SetAppDescription sets the appDescription property value. The description exposed by the associated application.
func (m *ServicePrincipal) SetAppDescription(value *string)() {
    m.appDescription = value
}
// SetAppDisplayName sets the appDisplayName property value. The display name exposed by the associated application. Maximum length is 256 characters.
func (m *ServicePrincipal) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppId sets the appId property value. The unique identifier for the associated application (its appId property). Alternate key. Supports $filter (eq, ne, not, in, startsWith).
func (m *ServicePrincipal) SetAppId(value *string)() {
    m.appId = value
}
// SetApplicationTemplateId sets the applicationTemplateId property value. Unique identifier of the applicationTemplate. Supports $filter (eq, not, ne). Read-only. null if the service principal wasn't created from an application template.
func (m *ServicePrincipal) SetApplicationTemplateId(value *string)() {
    m.applicationTemplateId = value
}
// SetAppManagementPolicies sets the appManagementPolicies property value. The appManagementPolicy applied to this application.
func (m *ServicePrincipal) SetAppManagementPolicies(value []AppManagementPolicyable)() {
    m.appManagementPolicies = value
}
// SetAppOwnerOrganizationId sets the appOwnerOrganizationId property value. Contains the tenant ID where the application is registered. This is applicable only to service principals backed by applications. Supports $filter (eq, ne, NOT, ge, le).
func (m *ServicePrincipal) SetAppOwnerOrganizationId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.appOwnerOrganizationId = value
}
// SetAppRoleAssignedTo sets the appRoleAssignedTo property value. App role assignments for this app or service, granted to users, groups, and other service principals. Supports $expand.
func (m *ServicePrincipal) SetAppRoleAssignedTo(value []AppRoleAssignmentable)() {
    m.appRoleAssignedTo = value
}
// SetAppRoleAssignmentRequired sets the appRoleAssignmentRequired property value. Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
func (m *ServicePrincipal) SetAppRoleAssignmentRequired(value *bool)() {
    m.appRoleAssignmentRequired = value
}
// SetAppRoleAssignments sets the appRoleAssignments property value. App role assignment for another app or service, granted to this service principal. Supports $expand.
func (m *ServicePrincipal) SetAppRoleAssignments(value []AppRoleAssignmentable)() {
    m.appRoleAssignments = value
}
// SetAppRoles sets the appRoles property value. The roles exposed by the application that's linked to this service principal. For more information, see the appRoles property definition on the application entity. Not nullable.
func (m *ServicePrincipal) SetAppRoles(value []AppRoleable)() {
    m.appRoles = value
}
// SetClaimsMappingPolicies sets the claimsMappingPolicies property value. The claimsMappingPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)() {
    m.claimsMappingPolicies = value
}
// SetCreatedObjects sets the createdObjects property value. Directory objects created by this service principal. Read-only. Nullable.
func (m *ServicePrincipal) SetCreatedObjects(value []DirectoryObjectable)() {
    m.createdObjects = value
}
// SetCustomSecurityAttributes sets the customSecurityAttributes property value. An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith). Filter value is case sensitive. To read this property, the calling app must be assigned the CustomSecAttributeAssignment.Read.All permission. To write this property, the calling app must be assigned the CustomSecAttributeAssignment.ReadWrite.All permissions. To read or write this property in delegated scenarios, the admin must be assigned the Attribute Assignment Administrator role.
func (m *ServicePrincipal) SetCustomSecurityAttributes(value CustomSecurityAttributeValueable)() {
    m.customSecurityAttributes = value
}
// SetDelegatedPermissionClassifications sets the delegatedPermissionClassifications property value. The delegatedPermissionClassifications property
func (m *ServicePrincipal) SetDelegatedPermissionClassifications(value []DelegatedPermissionClassificationable)() {
    m.delegatedPermissionClassifications = value
}
// SetDescription sets the description property value. Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps displays the application description in this field. The maximum allowed size is 1,024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *ServicePrincipal) SetDescription(value *string)() {
    m.description = value
}
// SetDisabledByMicrosoftStatus sets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
func (m *ServicePrincipal) SetDisabledByMicrosoftStatus(value *string)() {
    m.disabledByMicrosoftStatus = value
}
// SetDisplayName sets the displayName property value. The display name for the service principal. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderby.
func (m *ServicePrincipal) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndpoints sets the endpoints property value. The endpoints property
func (m *ServicePrincipal) SetEndpoints(value []Endpointable)() {
    m.endpoints = value
}
// SetFederatedIdentityCredentials sets the federatedIdentityCredentials property value. Federated identities for a specific type of service principal - managed identity. Supports $expand and $filter (/$count eq 0, /$count ne 0).
func (m *ServicePrincipal) SetFederatedIdentityCredentials(value []FederatedIdentityCredentialable)() {
    m.federatedIdentityCredentials = value
}
// SetHomepage sets the homepage property value. Home page or landing page of the application.
func (m *ServicePrincipal) SetHomepage(value *string)() {
    m.homepage = value
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)() {
    m.homeRealmDiscoveryPolicies = value
}
// SetInfo sets the info property value. Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Microsoft Entra apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
func (m *ServicePrincipal) SetInfo(value InformationalUrlable)() {
    m.info = value
}
// SetKeyCredentials sets the keyCredentials property value. The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, not, ge, le).
func (m *ServicePrincipal) SetKeyCredentials(value []KeyCredentialable)() {
    m.keyCredentials = value
}
// SetLoginUrl sets the loginUrl property value. Specifies the URL where the service provider redirects the user to Microsoft Entra ID to authenticate. Microsoft Entra ID uses the URL to launch the application from Microsoft 365 or the Microsoft Entra My Apps. When blank, Microsoft Entra ID performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Microsoft Entra My Apps, or the Microsoft Entra SSO URL.
func (m *ServicePrincipal) SetLoginUrl(value *string)() {
    m.loginUrl = value
}
// SetLogoutUrl sets the logoutUrl property value. Specifies the URL that the Microsoft's authorization service uses to sign out a user using OpenID Connect front-channel, back-channel, or SAML sign out protocols.
func (m *ServicePrincipal) SetLogoutUrl(value *string)() {
    m.logoutUrl = value
}
// SetMemberOf sets the memberOf property value. Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) SetMemberOf(value []DirectoryObjectable)() {
    m.memberOf = value
}
// SetNotes sets the notes property value. Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1,024 characters.
func (m *ServicePrincipal) SetNotes(value *string)() {
    m.notes = value
}
// SetNotificationEmailAddresses sets the notificationEmailAddresses property value. Specifies the list of email addresses where Microsoft Entra ID sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Microsoft Entra Gallery applications.
func (m *ServicePrincipal) SetNotificationEmailAddresses(value []string)() {
    m.notificationEmailAddresses = value
}
// SetOauth2PermissionGrants sets the oauth2PermissionGrants property value. Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
func (m *ServicePrincipal) SetOauth2PermissionGrants(value []OAuth2PermissionGrantable)() {
    m.oauth2PermissionGrants = value
}
// SetOauth2PermissionScopes sets the oauth2PermissionScopes property value. The delegated permissions exposed by the application. For more information, see the oauth2PermissionScopes property on the application entity's api property. Not nullable.
func (m *ServicePrincipal) SetOauth2PermissionScopes(value []PermissionScopeable)() {
    m.oauth2PermissionScopes = value
}
// SetOwnedObjects sets the ownedObjects property value. Directory objects that this service principal owns. Read-only. Nullable. Supports $expand, $select nested in $expand, and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
func (m *ServicePrincipal) SetOwnedObjects(value []DirectoryObjectable)() {
    m.ownedObjects = value
}
// SetOwners sets the owners property value. Directory objects that are owners of this servicePrincipal. The owners are a set of nonadmin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand, $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1), and $select nested in $expand.
func (m *ServicePrincipal) SetOwners(value []DirectoryObjectable)() {
    m.owners = value
}
// SetPasswordCredentials sets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
func (m *ServicePrincipal) SetPasswordCredentials(value []PasswordCredentialable)() {
    m.passwordCredentials = value
}
// SetPreferredSingleSignOnMode sets the preferredSingleSignOnMode property value. Specifies the single sign-on mode configured for this application. Microsoft Entra ID uses the preferred single sign-on mode to launch the application from Microsoft 365 or the My Apps portal. The supported values are password, saml, notSupported, and oidc. Note: This field might be null for older SAML apps and for OIDC applications where it isn't set automatically.
func (m *ServicePrincipal) SetPreferredSingleSignOnMode(value *string)() {
    m.preferredSingleSignOnMode = value
}
// SetPreferredTokenSigningKeyThumbprint sets the preferredTokenSigningKeyThumbprint property value. This property can be used on SAML applications (apps that have preferredSingleSignOnMode set to saml) to control which certificate is used to sign the SAML responses. For applications that aren't SAML, don't write or otherwise rely on this property.
func (m *ServicePrincipal) SetPreferredTokenSigningKeyThumbprint(value *string)() {
    m.preferredTokenSigningKeyThumbprint = value
}
// SetRemoteDesktopSecurityConfiguration sets the remoteDesktopSecurityConfiguration property value. The remoteDesktopSecurityConfiguration object applied to this service principal. Supports $filter (eq) for isRemoteDesktopProtocolEnabled property.
func (m *ServicePrincipal) SetRemoteDesktopSecurityConfiguration(value RemoteDesktopSecurityConfigurationable)() {
    m.remoteDesktopSecurityConfiguration = value
}
// SetReplyUrls sets the replyUrls property value. The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
func (m *ServicePrincipal) SetReplyUrls(value []string)() {
    m.replyUrls = value
}
// SetResourceSpecificApplicationPermissions sets the resourceSpecificApplicationPermissions property value. The resource-specific application permissions exposed by this application. Currently, resource-specific permissions are only supported for Teams apps accessing to specific chats and teams using Microsoft Graph. Read-only.
func (m *ServicePrincipal) SetResourceSpecificApplicationPermissions(value []ResourceSpecificPermissionable)() {
    m.resourceSpecificApplicationPermissions = value
}
// SetSamlSingleSignOnSettings sets the samlSingleSignOnSettings property value. The collection for settings related to saml single sign-on.
func (m *ServicePrincipal) SetSamlSingleSignOnSettings(value SamlSingleSignOnSettingsable)() {
    m.samlSingleSignOnSettings = value
}
// SetServicePrincipalNames sets the servicePrincipalNames property value. Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Microsoft Entra ID. For example,Client apps can specify a resource URI that is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) SetServicePrincipalNames(value []string)() {
    m.servicePrincipalNames = value
}
// SetServicePrincipalType sets the servicePrincipalType property value. Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Microsoft Entra ID internally. The servicePrincipalType property can be set to three different values: Application - A service principal that represents an application or service. The appId property identifies the associated app registration, and matches the appId of an application, possibly from a different tenant. If the associated app registration is missing, tokens aren't issued for the service principal.ManagedIdentity - A service principal that represents a managed identity. Service principals representing managed identities can be granted access and permissions, but can't be updated or modified directly.Legacy - A service principal that represents an app created before app registrations, or through legacy experiences. A legacy service principal can have credentials, service principal names, reply URLs, and other properties that are editable by an authorized user, but doesn't have an associated app registration. The appId value doesn't associate the service principal with an app registration. The service principal can only be used in the tenant where it was created.SocialIdp - For internal use.
func (m *ServicePrincipal) SetServicePrincipalType(value *string)() {
    m.servicePrincipalType = value
}
// SetSignInAudience sets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization's Microsoft Entra tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization's Microsoft Entra tenant (multitenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization's Microsoft Entra tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
func (m *ServicePrincipal) SetSignInAudience(value *string)() {
    m.signInAudience = value
}
// SetSynchronization sets the synchronization property value. Represents the capability for Microsoft Entra identity synchronization through the Microsoft Graph API.
func (m *ServicePrincipal) SetSynchronization(value Synchronizationable)() {
    m.synchronization = value
}
// SetTags sets the tags property value. Custom strings that can be used to categorize and identify the service principal. Not nullable. The value is the union of strings set here and on the associated application entity's tags property.Supports $filter (eq, not, ge, le, startsWith).
func (m *ServicePrincipal) SetTags(value []string)() {
    m.tags = value
}
// SetTokenEncryptionKeyId sets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Microsoft Entra ID issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *ServicePrincipal) SetTokenEncryptionKeyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.tokenEncryptionKeyId = value
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The tokenIssuancePolicies assigned to this service principal.
func (m *ServicePrincipal) SetTokenIssuancePolicies(value []TokenIssuancePolicyable)() {
    m.tokenIssuancePolicies = value
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this service principal.
func (m *ServicePrincipal) SetTokenLifetimePolicies(value []TokenLifetimePolicyable)() {
    m.tokenLifetimePolicies = value
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. The transitiveMemberOf property
func (m *ServicePrincipal) SetTransitiveMemberOf(value []DirectoryObjectable)() {
    m.transitiveMemberOf = value
}
// SetVerifiedPublisher sets the verifiedPublisher property value. Specifies the verified publisher of the application that's linked to this service principal.
func (m *ServicePrincipal) SetVerifiedPublisher(value VerifiedPublisherable)() {
    m.verifiedPublisher = value
}
type ServicePrincipalable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountEnabled()(*bool)
    GetAddIns()([]AddInable)
    GetAlternativeNames()([]string)
    GetAppDescription()(*string)
    GetAppDisplayName()(*string)
    GetAppId()(*string)
    GetApplicationTemplateId()(*string)
    GetAppManagementPolicies()([]AppManagementPolicyable)
    GetAppOwnerOrganizationId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetAppRoleAssignedTo()([]AppRoleAssignmentable)
    GetAppRoleAssignmentRequired()(*bool)
    GetAppRoleAssignments()([]AppRoleAssignmentable)
    GetAppRoles()([]AppRoleable)
    GetClaimsMappingPolicies()([]ClaimsMappingPolicyable)
    GetCreatedObjects()([]DirectoryObjectable)
    GetCustomSecurityAttributes()(CustomSecurityAttributeValueable)
    GetDelegatedPermissionClassifications()([]DelegatedPermissionClassificationable)
    GetDescription()(*string)
    GetDisabledByMicrosoftStatus()(*string)
    GetDisplayName()(*string)
    GetEndpoints()([]Endpointable)
    GetFederatedIdentityCredentials()([]FederatedIdentityCredentialable)
    GetHomepage()(*string)
    GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable)
    GetInfo()(InformationalUrlable)
    GetKeyCredentials()([]KeyCredentialable)
    GetLoginUrl()(*string)
    GetLogoutUrl()(*string)
    GetMemberOf()([]DirectoryObjectable)
    GetNotes()(*string)
    GetNotificationEmailAddresses()([]string)
    GetOauth2PermissionGrants()([]OAuth2PermissionGrantable)
    GetOauth2PermissionScopes()([]PermissionScopeable)
    GetOwnedObjects()([]DirectoryObjectable)
    GetOwners()([]DirectoryObjectable)
    GetPasswordCredentials()([]PasswordCredentialable)
    GetPreferredSingleSignOnMode()(*string)
    GetPreferredTokenSigningKeyThumbprint()(*string)
    GetRemoteDesktopSecurityConfiguration()(RemoteDesktopSecurityConfigurationable)
    GetReplyUrls()([]string)
    GetResourceSpecificApplicationPermissions()([]ResourceSpecificPermissionable)
    GetSamlSingleSignOnSettings()(SamlSingleSignOnSettingsable)
    GetServicePrincipalNames()([]string)
    GetServicePrincipalType()(*string)
    GetSignInAudience()(*string)
    GetSynchronization()(Synchronizationable)
    GetTags()([]string)
    GetTokenEncryptionKeyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetTokenIssuancePolicies()([]TokenIssuancePolicyable)
    GetTokenLifetimePolicies()([]TokenLifetimePolicyable)
    GetTransitiveMemberOf()([]DirectoryObjectable)
    GetVerifiedPublisher()(VerifiedPublisherable)
    SetAccountEnabled(value *bool)()
    SetAddIns(value []AddInable)()
    SetAlternativeNames(value []string)()
    SetAppDescription(value *string)()
    SetAppDisplayName(value *string)()
    SetAppId(value *string)()
    SetApplicationTemplateId(value *string)()
    SetAppManagementPolicies(value []AppManagementPolicyable)()
    SetAppOwnerOrganizationId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetAppRoleAssignedTo(value []AppRoleAssignmentable)()
    SetAppRoleAssignmentRequired(value *bool)()
    SetAppRoleAssignments(value []AppRoleAssignmentable)()
    SetAppRoles(value []AppRoleable)()
    SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)()
    SetCreatedObjects(value []DirectoryObjectable)()
    SetCustomSecurityAttributes(value CustomSecurityAttributeValueable)()
    SetDelegatedPermissionClassifications(value []DelegatedPermissionClassificationable)()
    SetDescription(value *string)()
    SetDisabledByMicrosoftStatus(value *string)()
    SetDisplayName(value *string)()
    SetEndpoints(value []Endpointable)()
    SetFederatedIdentityCredentials(value []FederatedIdentityCredentialable)()
    SetHomepage(value *string)()
    SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)()
    SetInfo(value InformationalUrlable)()
    SetKeyCredentials(value []KeyCredentialable)()
    SetLoginUrl(value *string)()
    SetLogoutUrl(value *string)()
    SetMemberOf(value []DirectoryObjectable)()
    SetNotes(value *string)()
    SetNotificationEmailAddresses(value []string)()
    SetOauth2PermissionGrants(value []OAuth2PermissionGrantable)()
    SetOauth2PermissionScopes(value []PermissionScopeable)()
    SetOwnedObjects(value []DirectoryObjectable)()
    SetOwners(value []DirectoryObjectable)()
    SetPasswordCredentials(value []PasswordCredentialable)()
    SetPreferredSingleSignOnMode(value *string)()
    SetPreferredTokenSigningKeyThumbprint(value *string)()
    SetRemoteDesktopSecurityConfiguration(value RemoteDesktopSecurityConfigurationable)()
    SetReplyUrls(value []string)()
    SetResourceSpecificApplicationPermissions(value []ResourceSpecificPermissionable)()
    SetSamlSingleSignOnSettings(value SamlSingleSignOnSettingsable)()
    SetServicePrincipalNames(value []string)()
    SetServicePrincipalType(value *string)()
    SetSignInAudience(value *string)()
    SetSynchronization(value Synchronizationable)()
    SetTags(value []string)()
    SetTokenEncryptionKeyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetTokenIssuancePolicies(value []TokenIssuancePolicyable)()
    SetTokenLifetimePolicies(value []TokenLifetimePolicyable)()
    SetTransitiveMemberOf(value []DirectoryObjectable)()
    SetVerifiedPublisher(value VerifiedPublisherable)()
}
