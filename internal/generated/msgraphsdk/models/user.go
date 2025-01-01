package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// User represents a Microsoft Entra user account.
type User struct {
    DirectoryObject
    // A freeform text entry field for the user to describe themselves. Returned only on $select.
    aboutMe *string
    // true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter (eq, ne, not, and in).
    accountEnabled *bool
    // The user's activities across devices. Read-only. Nullable.
    activities []UserActivityable
    // Sets the age group of the user. Allowed values: null, Minor, NotAdult, and Adult. For more information, see legal age group property definitions. Returned only on $select. Supports $filter (eq, ne, not, and in).
    ageGroup *string
    // The user's terms of use acceptance statuses. Read-only. Nullable.
    agreementAcceptances []AgreementAcceptanceable
    // Represents the app roles a user is granted for an application. Supports $expand.
    appRoleAssignments []AppRoleAssignmentable
    // The licenses that are assigned to the user, including inherited (group-based) licenses. This property doesn't differentiate between directly assigned and inherited licenses. Use the licenseAssignmentStates property to identify the directly assigned and inherited licenses. Not nullable. Returned only on $select. Supports $filter (eq, not, /$count eq 0, /$count ne 0).
    assignedLicenses []AssignedLicenseable
    // The plans that are assigned to the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq and not).
    assignedPlans []AssignedPlanable
    // The authentication methods that are supported for the user.
    authentication Authenticationable
    // The authorizationInfo property
    authorizationInfo AuthorizationInfoable
    // The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z. Returned only on $select.
    birthday *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The telephone numbers for the user. NOTE: Although it's a string collection, only one number can be set for this property. Read-only for users synced from the on-premises directory. Returned by default. Supports $filter (eq, not, ge, le, startsWith).
    businessPhones []string
    // The user's primary calendar. Read-only.
    calendar Calendarable
    // The user's calendar groups. Read-only. Nullable.
    calendarGroups []CalendarGroupable
    // The user's calendars. Read-only. Nullable.
    calendars []Calendarable
    // The calendar view for the calendar. Read-only. Nullable.
    calendarView []Eventable
    // The chats property
    chats []Chatable
    // The city where the user is located. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    city *string
    // The cloudClipboard property
    cloudClipboard CloudClipboardRootable
    // The name of the company that the user is associated with. This property can be useful for describing the company that a guest comes from. The maximum length is 64 characters.Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    companyName *string
    // Sets whether consent was obtained for minors. Allowed values: null, Granted, Denied, and NotRequired. For more information, see legal age group property definitions. Returned only on $select. Supports $filter (eq, ne, not, and in).
    consentProvidedForMinor *string
    // The user's contacts folders. Read-only. Nullable.
    contactFolders []ContactFolderable
    // The user's contacts. Read-only. Nullable.
    contacts []Contactable
    // The country/region where the user is located; for example, US or UK. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    country *string
    // The date and time the user was created, in ISO 8601 format and UTC. The value can't be modified and is automatically populated when the entity is created. Nullable. For on-premises users, the value represents when they were first created in Microsoft Entra ID. Property is null for some users created before June 2018 and on-premises users that were synced to Microsoft Entra ID before June 2018. Read-only. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Directory objects that the user created. Read-only. Nullable.
    createdObjects []DirectoryObjectable
    // Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by a guest signing up through a link that is part of a user flow (SelfServiceSignUp). Read-only.Returned only on $select. Supports $filter (eq, ne, not, in).
    creationType *string
    // An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith). The filter value is case-sensitive. To read this property, the calling app must be assigned the CustomSecAttributeAssignment.Read.All permission. To write this property, the calling app must be assigned the CustomSecAttributeAssignment.ReadWrite.All permissions. To read or write this property in delegated scenarios, the admin must be assigned the Attribute Assignment Administrator role.
    customSecurityAttributes CustomSecurityAttributeValueable
    // The name of the department in which the user works. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, and eq on null values).
    department *string
    // The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
    deviceEnrollmentLimit *int32
    // The list of troubleshooting events for this user.
    deviceManagementTroubleshootingEvents []DeviceManagementTroubleshootingEventable
    // The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
    directReports []DirectoryObjectable
    // The name displayed in the address book for the user. This value is usually the combination of the user's first name, middle initial, and family name. This property is required when a user is created and it can't be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values), $orderby, and $search.
    displayName *string
    // The user's OneDrive. Read-only.
    drive Driveable
    // A collection of drives available for this user. Read-only.
    drives []Driveable
    // The employeeExperience property
    employeeExperience EmployeeExperienceUserable
    // The date and time when the user was hired or will start work in a future hire. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
    employeeHireDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The employee identifier assigned to the user by the organization. The maximum length is 16 characters. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    employeeId *string
    // The date and time when the user left or will leave the organization. To read this property, the calling app must be assigned the User-LifeCycleInfo.Read.All permission. To write this property, the calling app must be assigned the User.Read.All and User-LifeCycleInfo.ReadWrite.All permissions. To read this property in delegated scenarios, the admin needs at least one of the following Microsoft Entra roles: Lifecycle Workflows Administrator (least privilege), Global Reader. To write this property in delegated scenarios, the admin needs the Global Administrator role. Supports $filter (eq, ne, not , ge, le, in). For more information, see Configure the employeeLeaveDateTime property for a user.
    employeeLeaveDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Represents organization data (for example, division and costCenter) associated with a user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
    employeeOrgData EmployeeOrgDataable
    // Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith).
    employeeType *string
    // The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
    events []Eventable
    // The collection of open extensions defined for the user. Read-only. Supports $expand. Nullable.
    extensions []Extensionable
    // For a guest invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only on $select. Supports $filter (eq, ne, not , in).
    externalUserState *string
    // Shows the timestamp for the latest change to the externalUserState property. Returned only on $select. Supports $filter (eq, ne, not , in).
    externalUserStateChangeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The fax number of the user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    faxNumber *string
    // The followedSites property
    followedSites []Siteable
    // The given name (first name) of the user. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    givenName *string
    // The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z. Returned only on $select.  Note: This property is specific to SharePoint in Microsoft 365. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
    hireDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Represents the identities that can be used to sign in to this user account. Microsoft (also known as a local account), organizations, or social identity providers such as Facebook, Google, and Microsoft can provide identity and tie it to a user account. It might contain multiple items with the same signInType value. Returned only on $select.  Supports $filter (eq) with limitations.
    identities []ObjectIdentityable
    // The instant message voice-over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith).
    imAddresses []string
    // Relevance classification of the user's messages based on explicit designations that override inferred relevance or importance.
    inferenceClassification InferenceClassificationable
    // Represents relationships between a user and items such as OneDrive for work or school documents, calculated using advanced analytics and machine learning techniques. Read-only. Nullable.
    insights ItemInsightsable
    // A list for the user to describe their interests. Returned only on $select.
    interests []string
    // The isManagementRestricted property
    isManagementRestricted *bool
    // Don't use – reserved for future use.
    isResourceAccount *bool
    // The user's job title. Maximum length is 128 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    jobTitle *string
    // The joinedTeams property
    joinedTeams []Teamable
    // The time when this Microsoft Entra user last changed their password or when their password was created, whichever date the latest action was performed. The date and time information uses ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
    lastPasswordChangeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, Undefined,  MinorWithOutParentalConsent, MinorWithParentalConsent, MinorNoParentalConsentRequired, NotAdult, and Adult. For more information, see legal age group property definitions. Returned only on $select.
    legalAgeGroupClassification *string
    // State of license assignments for this user. Also indicates licenses that are directly assigned or the user inherited through group memberships. Read-only. Returned only on $select.
    licenseAssignmentStates []LicenseAssignmentStateable
    // A collection of this user's license details. Read-only.
    licenseDetails []LicenseDetailsable
    // The SMTP address for the user, for example, jeff@contoso.com. Changes to this property update the user's proxyAddresses collection to include the value as an SMTP address. This property can't contain accent characters.  NOTE: We don't recommend updating this property for Azure AD B2C user profiles. Use the otherMails property instead. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith, and eq on null values).
    mail *string
    // Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. Returned only on $select.
    mailboxSettings MailboxSettingsable
    // The user's mail folders. Read-only. Nullable.
    mailFolders []MailFolderable
    // The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    mailNickname *string
    // Zero or more managed app registrations that belong to the user.
    managedAppRegistrations []ManagedAppRegistrationable
    // The managed devices associated with the user.
    managedDevices []ManagedDeviceable
    // The user or contact that is this user's manager. Read-only. Supports $expand.
    manager DirectoryObjectable
    // The groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObjectable
    // The messages in a mailbox or folder. Read-only. Nullable.
    messages []Messageable
    // The primary cellular telephone number for the user. Read-only for users synced from the on-premises directory. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values) and $search.
    mobilePhone *string
    // The URL for the user's site. Returned only on $select.
    mySite *string
    // The oauth2PermissionGrants property
    oauth2PermissionGrants []OAuth2PermissionGrantable
    // The office location in the user's place of business. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    officeLocation *string
    // The onenote property
    onenote Onenoteable
    // Information about a meeting, including the URL used to join a meeting, the attendees list, and the description.
    onlineMeetings []OnlineMeetingable
    // Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select.
    onPremisesDistinguishedName *string
    // Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select.
    onPremisesDomainName *string
    // Contains extensionAttributes1-15 for the user. These extension attributes are also known as Exchange custom attributes 1-15. Each attribute can store up to 1024 characters. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties can be set during the creation or update of a user object.  For a cloud-only user previously synced from on-premises Active Directory, these properties are read-only in Microsoft Graph but can be fully managed through the Exchange Admin Center or the Exchange Online V2 module in PowerShell. Returned only on $select. Supports $filter (eq, ne, not, in).
    onPremisesExtensionAttributes OnPremisesExtensionAttributesable
    // This property is used to associate an on-premises Active Directory user account to their Microsoft Entra user object. This property must be specified when creating a new user account in the Graph if you're using a federated domain for the user's userPrincipalName (UPN) property. NOTE: The $ and _ characters can't be used when specifying this property. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).
    onPremisesImmutableId *string
    // Indicates the last time at which the object was synced with the on-premises directory; for example: 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Errors when using Microsoft synchronization product during provisioning. Returned only on $select. Supports $filter (eq, not, ge, le).
    onPremisesProvisioningErrors []OnPremisesProvisioningErrorable
    // Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
    onPremisesSamAccountName *string
    // Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Returned only on $select. Supports $filter (eq including on null values).
    onPremisesSecurityIdentifier *string
    // true if this user object is currently being synced from an on-premises Active Directory (AD); otherwise the user isn't being synced and can be managed in Microsoft Entra ID. Read-only. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values).
    onPremisesSyncEnabled *bool
    // Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
    onPremisesUserPrincipalName *string
    // A list of other email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com']. NOTE: This property can't contain accent characters. Returned only on $select. Supports $filter (eq, not, ge, le, in, startsWith, endsWith, /$count eq 0, /$count ne 0).
    otherMails []string
    // The outlook property
    outlook OutlookUserable
    // Devices the user owns. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
    ownedDevices []DirectoryObjectable
    // Directory objects the user owns. Read-only. Nullable. Supports $expand, $select nested in $expand, and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
    ownedObjects []DirectoryObjectable
    // Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two might be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. Returned only on $select. For more information on the default password policies, see Microsoft Entra password policies. Supports $filter (ne, not, and eq on null values).
    passwordPolicies *string
    // Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). To update this property:  In delegated access, the calling app must be assigned the User-PasswordProfile.ReadWrite.All delegated permission on behalf of the signed-in user.  In application-only access, the calling app must be assigned the User-PasswordProfile.ReadWrite.All application permission and at least the User Administrator Microsoft Entra role.
    passwordProfile PasswordProfileable
    // A list for the user to enumerate their past projects. Returned only on $select.
    pastProjects []string
    // People that are relevant to the user. Read-only. Nullable.
    people []Personable
    // List all resource-specific permission grants of a user.
    permissionGrants []ResourceSpecificPermissionGrantable
    // The user's profile photo. Read-only.
    photo ProfilePhotoable
    // The collection of the user's profile photos in different sizes. Read-only.
    photos []ProfilePhotoable
    // Entry-point to the Planner resource that might exist for a user. Read-only.
    planner PlannerUserable
    // The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    postalCode *string
    // The preferred data location for the user. For more information, see OneDrive Online Multi-Geo.
    preferredDataLocation *string
    // The preferred language for the user. The preferred language format is based on RFC 4646. The name is a combination of an ISO 639 two-letter lowercase culture code associated with the language, and an ISO 3166 two-letter uppercase subculture code associated with the country or region. Example: 'en-US', or 'es-ES'. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values)
    preferredLanguage *string
    // The preferred name for the user. Not Supported. This attribute returns an empty string.Returned only on $select.
    preferredName *string
    // The presence property
    presence Presenceable
    // The print property
    print UserPrintable
    // The plans that are provisioned for the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le).
    provisionedPlans []ProvisionedPlanable
    // For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. Changes to the mail property update this collection to include the value as an SMTP address. For more information, see mail and proxyAddresses properties. The proxy address prefixed with SMTP (capitalized) is the primary proxy address, while those addresses prefixed with smtp are the secondary proxy addresses. For Azure AD B2C accounts, this property has a limit of 10 unique addresses. Read-only in Microsoft Graph; you can update this property only through the Microsoft 365 admin center. Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith, endsWith, /$count eq 0, /$count ne 0).
    proxyAddresses []string
    // Devices that are registered for the user. Read-only. Nullable. Supports $expand and returns up to 100 objects.
    registeredDevices []DirectoryObjectable
    // A list for the user to enumerate their responsibilities. Returned only on $select.
    responsibilities []string
    // A list for the user to enumerate the schools they attended. Returned only on $select.
    schools []string
    // The scopedRoleMemberOf property
    scopedRoleMemberOf []ScopedRoleMembershipable
    // Security identifier (SID) of the user, used in Windows scenarios. Read-only. Returned by default. Supports $select and $filter (eq, not, ge, le, startsWith).
    securityIdentifier *string
    // Errors published by a federated service describing a nontransient, service-specific error regarding the properties or link from a user object.  Supports $filter (eq, not, for isResolved and serviceInstance).
    serviceProvisioningErrors []ServiceProvisioningErrorable
    // The settings property
    settings UserSettingsable
    // Do not use in Microsoft Graph. Manage this property through the Microsoft 365 admin center instead. Represents whether the user should be included in the Outlook global address list. See Known issue.
    showInAddressList *bool
    // Get the last signed-in date and request ID of the sign-in for a given user. Read-only.Returned only on $select. Supports $filter (eq, ne, not, ge, le) but not with any other filterable properties. Note: Details for this property require a Microsoft Entra ID P1 or P2 license and the AuditLog.Read.All permission.This property isn't returned for a user who never signed in or last signed in before April 2020.
    signInActivity SignInActivityable
    // Any refresh tokens or session tokens (session cookies) issued before this time are invalid. Applications get an error when using an invalid refresh or session token to acquire a delegated access token (to access APIs such as Microsoft Graph). If this happens, the application needs to acquire a new refresh token by requesting the authorized endpoint. Read-only. Use revokeSignInSessions to reset. Returned only on $select.
    signInSessionsValidFromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A list for the user to enumerate their skills. Returned only on $select.
    skills []string
    // The identifier that relates the user to the working time schedule triggers. Read-Only. Nullable
    solutions UserSolutionRootable
    // The users and groups responsible for this guest's privileges in the tenant and keeping the guest's information and access updated. (HTTP Methods: GET, POST, DELETE.). Supports $expand.
    sponsors []DirectoryObjectable
    // The state or province in the user's address. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    state *string
    // The street address of the user's place of business. Maximum length is 1,024 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    streetAddress *string
    // The user's surname (family name or last name). Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    surname *string
    // A container for Microsoft Teams features available for the user. Read-only. Nullable.
    teamwork UserTeamworkable
    // Represents the To Do services available to a user.
    todo Todoable
    // The groups, including nested groups, and directory roles that a user is a member of. Nullable.
    transitiveMemberOf []DirectoryObjectable
    // A two-letter country code (ISO standard 3166). Required for users that are assigned licenses due to legal requirements to check for availability of services in countries. Examples include: US, JP, and GB. Not nullable. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    usageLocation *string
    // The user principal name (UPN) of the user. The UPN is an Internet-style sign-in name for the user based on the Internet standard RFC 822. By convention, this value should map to the user's email name. The general format is alias@domain, where the domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property can't contain accent characters. Only the following characters are allowed A - Z, a - z, 0 - 9, ' . - _ ! # ^ ~. For the complete list of allowed characters, see username policies. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith) and $orderby.
    userPrincipalName *string
    // A string value that can be used to classify user types in your directory. The possible values are Member and Guest. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). NOTE: For more information about the permissions for members and guests, see What are the default user permissions in Microsoft Entra ID?
    userType *string
}
// NewUser instantiates a new User and sets the default values.
func NewUser()(*User) {
    m := &User{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.user"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUser(), nil
}
// GetAboutMe gets the aboutMe property value. A freeform text entry field for the user to describe themselves. Returned only on $select.
// returns a *string when successful
func (m *User) GetAboutMe()(*string) {
    return m.aboutMe
}
// GetAccountEnabled gets the accountEnabled property value. true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter (eq, ne, not, and in).
// returns a *bool when successful
func (m *User) GetAccountEnabled()(*bool) {
    return m.accountEnabled
}
// GetActivities gets the activities property value. The user's activities across devices. Read-only. Nullable.
// returns a []UserActivityable when successful
func (m *User) GetActivities()([]UserActivityable) {
    return m.activities
}
// GetAgeGroup gets the ageGroup property value. Sets the age group of the user. Allowed values: null, Minor, NotAdult, and Adult. For more information, see legal age group property definitions. Returned only on $select. Supports $filter (eq, ne, not, and in).
// returns a *string when successful
func (m *User) GetAgeGroup()(*string) {
    return m.ageGroup
}
// GetAgreementAcceptances gets the agreementAcceptances property value. The user's terms of use acceptance statuses. Read-only. Nullable.
// returns a []AgreementAcceptanceable when successful
func (m *User) GetAgreementAcceptances()([]AgreementAcceptanceable) {
    return m.agreementAcceptances
}
// GetAppRoleAssignments gets the appRoleAssignments property value. Represents the app roles a user is granted for an application. Supports $expand.
// returns a []AppRoleAssignmentable when successful
func (m *User) GetAppRoleAssignments()([]AppRoleAssignmentable) {
    return m.appRoleAssignments
}
// GetAssignedLicenses gets the assignedLicenses property value. The licenses that are assigned to the user, including inherited (group-based) licenses. This property doesn't differentiate between directly assigned and inherited licenses. Use the licenseAssignmentStates property to identify the directly assigned and inherited licenses. Not nullable. Returned only on $select. Supports $filter (eq, not, /$count eq 0, /$count ne 0).
// returns a []AssignedLicenseable when successful
func (m *User) GetAssignedLicenses()([]AssignedLicenseable) {
    return m.assignedLicenses
}
// GetAssignedPlans gets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq and not).
// returns a []AssignedPlanable when successful
func (m *User) GetAssignedPlans()([]AssignedPlanable) {
    return m.assignedPlans
}
// GetAuthentication gets the authentication property value. The authentication methods that are supported for the user.
// returns a Authenticationable when successful
func (m *User) GetAuthentication()(Authenticationable) {
    return m.authentication
}
// GetAuthorizationInfo gets the authorizationInfo property value. The authorizationInfo property
// returns a AuthorizationInfoable when successful
func (m *User) GetAuthorizationInfo()(AuthorizationInfoable) {
    return m.authorizationInfo
}
// GetBirthday gets the birthday property value. The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z. Returned only on $select.
// returns a *Time when successful
func (m *User) GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.birthday
}
// GetBusinessPhones gets the businessPhones property value. The telephone numbers for the user. NOTE: Although it's a string collection, only one number can be set for this property. Read-only for users synced from the on-premises directory. Returned by default. Supports $filter (eq, not, ge, le, startsWith).
// returns a []string when successful
func (m *User) GetBusinessPhones()([]string) {
    return m.businessPhones
}
// GetCalendar gets the calendar property value. The user's primary calendar. Read-only.
// returns a Calendarable when successful
func (m *User) GetCalendar()(Calendarable) {
    return m.calendar
}
// GetCalendarGroups gets the calendarGroups property value. The user's calendar groups. Read-only. Nullable.
// returns a []CalendarGroupable when successful
func (m *User) GetCalendarGroups()([]CalendarGroupable) {
    return m.calendarGroups
}
// GetCalendars gets the calendars property value. The user's calendars. Read-only. Nullable.
// returns a []Calendarable when successful
func (m *User) GetCalendars()([]Calendarable) {
    return m.calendars
}
// GetCalendarView gets the calendarView property value. The calendar view for the calendar. Read-only. Nullable.
// returns a []Eventable when successful
func (m *User) GetCalendarView()([]Eventable) {
    return m.calendarView
}
// GetChats gets the chats property value. The chats property
// returns a []Chatable when successful
func (m *User) GetChats()([]Chatable) {
    return m.chats
}
// GetCity gets the city property value. The city where the user is located. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetCity()(*string) {
    return m.city
}
// GetCloudClipboard gets the cloudClipboard property value. The cloudClipboard property
// returns a CloudClipboardRootable when successful
func (m *User) GetCloudClipboard()(CloudClipboardRootable) {
    return m.cloudClipboard
}
// GetCompanyName gets the companyName property value. The name of the company that the user is associated with. This property can be useful for describing the company that a guest comes from. The maximum length is 64 characters.Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetCompanyName()(*string) {
    return m.companyName
}
// GetConsentProvidedForMinor gets the consentProvidedForMinor property value. Sets whether consent was obtained for minors. Allowed values: null, Granted, Denied, and NotRequired. For more information, see legal age group property definitions. Returned only on $select. Supports $filter (eq, ne, not, and in).
// returns a *string when successful
func (m *User) GetConsentProvidedForMinor()(*string) {
    return m.consentProvidedForMinor
}
// GetContactFolders gets the contactFolders property value. The user's contacts folders. Read-only. Nullable.
// returns a []ContactFolderable when successful
func (m *User) GetContactFolders()([]ContactFolderable) {
    return m.contactFolders
}
// GetContacts gets the contacts property value. The user's contacts. Read-only. Nullable.
// returns a []Contactable when successful
func (m *User) GetContacts()([]Contactable) {
    return m.contacts
}
// GetCountry gets the country property value. The country/region where the user is located; for example, US or UK. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetCountry()(*string) {
    return m.country
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the user was created, in ISO 8601 format and UTC. The value can't be modified and is automatically populated when the entity is created. Nullable. For on-premises users, the value represents when they were first created in Microsoft Entra ID. Property is null for some users created before June 2018 and on-premises users that were synced to Microsoft Entra ID before June 2018. Read-only. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
// returns a *Time when successful
func (m *User) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCreatedObjects gets the createdObjects property value. Directory objects that the user created. Read-only. Nullable.
// returns a []DirectoryObjectable when successful
func (m *User) GetCreatedObjects()([]DirectoryObjectable) {
    return m.createdObjects
}
// GetCreationType gets the creationType property value. Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by a guest signing up through a link that is part of a user flow (SelfServiceSignUp). Read-only.Returned only on $select. Supports $filter (eq, ne, not, in).
// returns a *string when successful
func (m *User) GetCreationType()(*string) {
    return m.creationType
}
// GetCustomSecurityAttributes gets the customSecurityAttributes property value. An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith). The filter value is case-sensitive. To read this property, the calling app must be assigned the CustomSecAttributeAssignment.Read.All permission. To write this property, the calling app must be assigned the CustomSecAttributeAssignment.ReadWrite.All permissions. To read or write this property in delegated scenarios, the admin must be assigned the Attribute Assignment Administrator role.
// returns a CustomSecurityAttributeValueable when successful
func (m *User) GetCustomSecurityAttributes()(CustomSecurityAttributeValueable) {
    return m.customSecurityAttributes
}
// GetDepartment gets the department property value. The name of the department in which the user works. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, and eq on null values).
// returns a *string when successful
func (m *User) GetDepartment()(*string) {
    return m.department
}
// GetDeviceEnrollmentLimit gets the deviceEnrollmentLimit property value. The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
// returns a *int32 when successful
func (m *User) GetDeviceEnrollmentLimit()(*int32) {
    return m.deviceEnrollmentLimit
}
// GetDeviceManagementTroubleshootingEvents gets the deviceManagementTroubleshootingEvents property value. The list of troubleshooting events for this user.
// returns a []DeviceManagementTroubleshootingEventable when successful
func (m *User) GetDeviceManagementTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable) {
    return m.deviceManagementTroubleshootingEvents
}
// GetDirectReports gets the directReports property value. The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
// returns a []DirectoryObjectable when successful
func (m *User) GetDirectReports()([]DirectoryObjectable) {
    return m.directReports
}
// GetDisplayName gets the displayName property value. The name displayed in the address book for the user. This value is usually the combination of the user's first name, middle initial, and family name. This property is required when a user is created and it can't be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values), $orderby, and $search.
// returns a *string when successful
func (m *User) GetDisplayName()(*string) {
    return m.displayName
}
// GetDrive gets the drive property value. The user's OneDrive. Read-only.
// returns a Driveable when successful
func (m *User) GetDrive()(Driveable) {
    return m.drive
}
// GetDrives gets the drives property value. A collection of drives available for this user. Read-only.
// returns a []Driveable when successful
func (m *User) GetDrives()([]Driveable) {
    return m.drives
}
// GetEmployeeExperience gets the employeeExperience property value. The employeeExperience property
// returns a EmployeeExperienceUserable when successful
func (m *User) GetEmployeeExperience()(EmployeeExperienceUserable) {
    return m.employeeExperience
}
// GetEmployeeHireDate gets the employeeHireDate property value. The date and time when the user was hired or will start work in a future hire. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
// returns a *Time when successful
func (m *User) GetEmployeeHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.employeeHireDate
}
// GetEmployeeId gets the employeeId property value. The employee identifier assigned to the user by the organization. The maximum length is 16 characters. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetEmployeeId()(*string) {
    return m.employeeId
}
// GetEmployeeLeaveDateTime gets the employeeLeaveDateTime property value. The date and time when the user left or will leave the organization. To read this property, the calling app must be assigned the User-LifeCycleInfo.Read.All permission. To write this property, the calling app must be assigned the User.Read.All and User-LifeCycleInfo.ReadWrite.All permissions. To read this property in delegated scenarios, the admin needs at least one of the following Microsoft Entra roles: Lifecycle Workflows Administrator (least privilege), Global Reader. To write this property in delegated scenarios, the admin needs the Global Administrator role. Supports $filter (eq, ne, not , ge, le, in). For more information, see Configure the employeeLeaveDateTime property for a user.
// returns a *Time when successful
func (m *User) GetEmployeeLeaveDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.employeeLeaveDateTime
}
// GetEmployeeOrgData gets the employeeOrgData property value. Represents organization data (for example, division and costCenter) associated with a user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
// returns a EmployeeOrgDataable when successful
func (m *User) GetEmployeeOrgData()(EmployeeOrgDataable) {
    return m.employeeOrgData
}
// GetEmployeeType gets the employeeType property value. Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith).
// returns a *string when successful
func (m *User) GetEmployeeType()(*string) {
    return m.employeeType
}
// GetEvents gets the events property value. The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
// returns a []Eventable when successful
func (m *User) GetEvents()([]Eventable) {
    return m.events
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the user. Read-only. Supports $expand. Nullable.
// returns a []Extensionable when successful
func (m *User) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetExternalUserState gets the externalUserState property value. For a guest invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only on $select. Supports $filter (eq, ne, not , in).
// returns a *string when successful
func (m *User) GetExternalUserState()(*string) {
    return m.externalUserState
}
// GetExternalUserStateChangeDateTime gets the externalUserStateChangeDateTime property value. Shows the timestamp for the latest change to the externalUserState property. Returned only on $select. Supports $filter (eq, ne, not , in).
// returns a *Time when successful
func (m *User) GetExternalUserStateChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.externalUserStateChangeDateTime
}
// GetFaxNumber gets the faxNumber property value. The fax number of the user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetFaxNumber()(*string) {
    return m.faxNumber
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *User) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["aboutMe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAboutMe(val)
        }
        return nil
    }
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
    res["activities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserActivityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserActivityable)
                }
            }
            m.SetActivities(res)
        }
        return nil
    }
    res["ageGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgeGroup(val)
        }
        return nil
    }
    res["agreementAcceptances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgreementAcceptanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementAcceptanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AgreementAcceptanceable)
                }
            }
            m.SetAgreementAcceptances(res)
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
    res["assignedLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignedLicenseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedLicenseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AssignedLicenseable)
                }
            }
            m.SetAssignedLicenses(res)
        }
        return nil
    }
    res["assignedPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignedPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedPlanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AssignedPlanable)
                }
            }
            m.SetAssignedPlans(res)
        }
        return nil
    }
    res["authentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthentication(val.(Authenticationable))
        }
        return nil
    }
    res["authorizationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationInfo(val.(AuthorizationInfoable))
        }
        return nil
    }
    res["birthday"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthday(val)
        }
        return nil
    }
    res["businessPhones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBusinessPhones(res)
        }
        return nil
    }
    res["calendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCalendarFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendar(val.(Calendarable))
        }
        return nil
    }
    res["calendarGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCalendarGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CalendarGroupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CalendarGroupable)
                }
            }
            m.SetCalendarGroups(res)
        }
        return nil
    }
    res["calendars"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCalendarFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Calendarable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Calendarable)
                }
            }
            m.SetCalendars(res)
        }
        return nil
    }
    res["calendarView"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Eventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Eventable)
                }
            }
            m.SetCalendarView(res)
        }
        return nil
    }
    res["chats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Chatable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Chatable)
                }
            }
            m.SetChats(res)
        }
        return nil
    }
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["cloudClipboard"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudClipboardRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudClipboard(val.(CloudClipboardRootable))
        }
        return nil
    }
    res["companyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyName(val)
        }
        return nil
    }
    res["consentProvidedForMinor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsentProvidedForMinor(val)
        }
        return nil
    }
    res["contactFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContactFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContactFolderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ContactFolderable)
                }
            }
            m.SetContactFolders(res)
        }
        return nil
    }
    res["contacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Contactable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Contactable)
                }
            }
            m.SetContacts(res)
        }
        return nil
    }
    res["country"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["creationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationType(val)
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
    res["department"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
        }
        return nil
    }
    res["deviceEnrollmentLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceEnrollmentLimit(val)
        }
        return nil
    }
    res["deviceManagementTroubleshootingEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTroubleshootingEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTroubleshootingEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementTroubleshootingEventable)
                }
            }
            m.SetDeviceManagementTroubleshootingEvents(res)
        }
        return nil
    }
    res["directReports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetDirectReports(res)
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
    res["drive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrive(val.(Driveable))
        }
        return nil
    }
    res["drives"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Driveable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Driveable)
                }
            }
            m.SetDrives(res)
        }
        return nil
    }
    res["employeeExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmployeeExperienceUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeExperience(val.(EmployeeExperienceUserable))
        }
        return nil
    }
    res["employeeHireDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeHireDate(val)
        }
        return nil
    }
    res["employeeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeId(val)
        }
        return nil
    }
    res["employeeLeaveDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeLeaveDateTime(val)
        }
        return nil
    }
    res["employeeOrgData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmployeeOrgDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeOrgData(val.(EmployeeOrgDataable))
        }
        return nil
    }
    res["employeeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeType(val)
        }
        return nil
    }
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Eventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Eventable)
                }
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Extensionable)
                }
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["externalUserState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUserState(val)
        }
        return nil
    }
    res["externalUserStateChangeDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUserStateChangeDateTime(val)
        }
        return nil
    }
    res["faxNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaxNumber(val)
        }
        return nil
    }
    res["followedSites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Siteable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Siteable)
                }
            }
            m.SetFollowedSites(res)
        }
        return nil
    }
    res["givenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["hireDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHireDate(val)
        }
        return nil
    }
    res["identities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateObjectIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ObjectIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ObjectIdentityable)
                }
            }
            m.SetIdentities(res)
        }
        return nil
    }
    res["imAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetImAddresses(res)
        }
        return nil
    }
    res["inferenceClassification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInferenceClassificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInferenceClassification(val.(InferenceClassificationable))
        }
        return nil
    }
    res["insights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemInsightsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsights(val.(ItemInsightsable))
        }
        return nil
    }
    res["interests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetInterests(res)
        }
        return nil
    }
    res["isManagementRestricted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsManagementRestricted(val)
        }
        return nil
    }
    res["isResourceAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsResourceAccount(val)
        }
        return nil
    }
    res["jobTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
        }
        return nil
    }
    res["joinedTeams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Teamable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Teamable)
                }
            }
            m.SetJoinedTeams(res)
        }
        return nil
    }
    res["lastPasswordChangeDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastPasswordChangeDateTime(val)
        }
        return nil
    }
    res["legalAgeGroupClassification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegalAgeGroupClassification(val)
        }
        return nil
    }
    res["licenseAssignmentStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLicenseAssignmentStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LicenseAssignmentStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LicenseAssignmentStateable)
                }
            }
            m.SetLicenseAssignmentStates(res)
        }
        return nil
    }
    res["licenseDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLicenseDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LicenseDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LicenseDetailsable)
                }
            }
            m.SetLicenseDetails(res)
        }
        return nil
    }
    res["mail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMail(val)
        }
        return nil
    }
    res["mailboxSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMailboxSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxSettings(val.(MailboxSettingsable))
        }
        return nil
    }
    res["mailFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMailFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MailFolderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MailFolderable)
                }
            }
            m.SetMailFolders(res)
        }
        return nil
    }
    res["mailNickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNickname(val)
        }
        return nil
    }
    res["managedAppRegistrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppRegistrationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedAppRegistrationable)
                }
            }
            m.SetManagedAppRegistrations(res)
        }
        return nil
    }
    res["managedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedDeviceable)
                }
            }
            m.SetManagedDevices(res)
        }
        return nil
    }
    res["manager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val.(DirectoryObjectable))
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
    res["messages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Messageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Messageable)
                }
            }
            m.SetMessages(res)
        }
        return nil
    }
    res["mobilePhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
        }
        return nil
    }
    res["mySite"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMySite(val)
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
    res["officeLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["onenote"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnenoteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnenote(val.(Onenoteable))
        }
        return nil
    }
    res["onlineMeetings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnlineMeetingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnlineMeetingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OnlineMeetingable)
                }
            }
            m.SetOnlineMeetings(res)
        }
        return nil
    }
    res["onPremisesDistinguishedName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesDistinguishedName(val)
        }
        return nil
    }
    res["onPremisesDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesDomainName(val)
        }
        return nil
    }
    res["onPremisesExtensionAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesExtensionAttributesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesExtensionAttributes(val.(OnPremisesExtensionAttributesable))
        }
        return nil
    }
    res["onPremisesImmutableId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesImmutableId(val)
        }
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesProvisioningErrors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesProvisioningErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesProvisioningErrorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OnPremisesProvisioningErrorable)
                }
            }
            m.SetOnPremisesProvisioningErrors(res)
        }
        return nil
    }
    res["onPremisesSamAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSamAccountName(val)
        }
        return nil
    }
    res["onPremisesSecurityIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSecurityIdentifier(val)
        }
        return nil
    }
    res["onPremisesSyncEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSyncEnabled(val)
        }
        return nil
    }
    res["onPremisesUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesUserPrincipalName(val)
        }
        return nil
    }
    res["otherMails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetOtherMails(res)
        }
        return nil
    }
    res["outlook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutlookUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutlook(val.(OutlookUserable))
        }
        return nil
    }
    res["ownedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetOwnedDevices(res)
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
    res["passwordPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPolicies(val)
        }
        return nil
    }
    res["passwordProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePasswordProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordProfile(val.(PasswordProfileable))
        }
        return nil
    }
    res["pastProjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPastProjects(res)
        }
        return nil
    }
    res["people"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Personable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Personable)
                }
            }
            m.SetPeople(res)
        }
        return nil
    }
    res["permissionGrants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceSpecificPermissionGrantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceSpecificPermissionGrantable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ResourceSpecificPermissionGrantable)
                }
            }
            m.SetPermissionGrants(res)
        }
        return nil
    }
    res["photo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfilePhotoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(ProfilePhotoable))
        }
        return nil
    }
    res["photos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProfilePhotoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfilePhotoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProfilePhotoable)
                }
            }
            m.SetPhotos(res)
        }
        return nil
    }
    res["planner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanner(val.(PlannerUserable))
        }
        return nil
    }
    res["postalCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostalCode(val)
        }
        return nil
    }
    res["preferredDataLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredDataLocation(val)
        }
        return nil
    }
    res["preferredLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguage(val)
        }
        return nil
    }
    res["preferredName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredName(val)
        }
        return nil
    }
    res["presence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePresenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPresence(val.(Presenceable))
        }
        return nil
    }
    res["print"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserPrintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrint(val.(UserPrintable))
        }
        return nil
    }
    res["provisionedPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisionedPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisionedPlanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProvisionedPlanable)
                }
            }
            m.SetProvisionedPlans(res)
        }
        return nil
    }
    res["proxyAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetProxyAddresses(res)
        }
        return nil
    }
    res["registeredDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRegisteredDevices(res)
        }
        return nil
    }
    res["responsibilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetResponsibilities(res)
        }
        return nil
    }
    res["schools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSchools(res)
        }
        return nil
    }
    res["scopedRoleMemberOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateScopedRoleMembershipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScopedRoleMembershipable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ScopedRoleMembershipable)
                }
            }
            m.SetScopedRoleMemberOf(res)
        }
        return nil
    }
    res["securityIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityIdentifier(val)
        }
        return nil
    }
    res["serviceProvisioningErrors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceProvisioningErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceProvisioningErrorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServiceProvisioningErrorable)
                }
            }
            m.SetServiceProvisioningErrors(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(UserSettingsable))
        }
        return nil
    }
    res["showInAddressList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowInAddressList(val)
        }
        return nil
    }
    res["signInActivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInActivity(val.(SignInActivityable))
        }
        return nil
    }
    res["signInSessionsValidFromDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInSessionsValidFromDateTime(val)
        }
        return nil
    }
    res["skills"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSkills(res)
        }
        return nil
    }
    res["solutions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserSolutionRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSolutions(val.(UserSolutionRootable))
        }
        return nil
    }
    res["sponsors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSponsors(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["streetAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreetAddress(val)
        }
        return nil
    }
    res["surname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["teamwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserTeamworkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamwork(val.(UserTeamworkable))
        }
        return nil
    }
    res["todo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTodoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTodo(val.(Todoable))
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
    res["usageLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageLocation(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val)
        }
        return nil
    }
    return res
}
// GetFollowedSites gets the followedSites property value. The followedSites property
// returns a []Siteable when successful
func (m *User) GetFollowedSites()([]Siteable) {
    return m.followedSites
}
// GetGivenName gets the givenName property value. The given name (first name) of the user. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetGivenName()(*string) {
    return m.givenName
}
// GetHireDate gets the hireDate property value. The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z. Returned only on $select.  Note: This property is specific to SharePoint in Microsoft 365. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
// returns a *Time when successful
func (m *User) GetHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.hireDate
}
// GetIdentities gets the identities property value. Represents the identities that can be used to sign in to this user account. Microsoft (also known as a local account), organizations, or social identity providers such as Facebook, Google, and Microsoft can provide identity and tie it to a user account. It might contain multiple items with the same signInType value. Returned only on $select.  Supports $filter (eq) with limitations.
// returns a []ObjectIdentityable when successful
func (m *User) GetIdentities()([]ObjectIdentityable) {
    return m.identities
}
// GetImAddresses gets the imAddresses property value. The instant message voice-over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith).
// returns a []string when successful
func (m *User) GetImAddresses()([]string) {
    return m.imAddresses
}
// GetInferenceClassification gets the inferenceClassification property value. Relevance classification of the user's messages based on explicit designations that override inferred relevance or importance.
// returns a InferenceClassificationable when successful
func (m *User) GetInferenceClassification()(InferenceClassificationable) {
    return m.inferenceClassification
}
// GetInsights gets the insights property value. Represents relationships between a user and items such as OneDrive for work or school documents, calculated using advanced analytics and machine learning techniques. Read-only. Nullable.
// returns a ItemInsightsable when successful
func (m *User) GetInsights()(ItemInsightsable) {
    return m.insights
}
// GetInterests gets the interests property value. A list for the user to describe their interests. Returned only on $select.
// returns a []string when successful
func (m *User) GetInterests()([]string) {
    return m.interests
}
// GetIsManagementRestricted gets the isManagementRestricted property value. The isManagementRestricted property
// returns a *bool when successful
func (m *User) GetIsManagementRestricted()(*bool) {
    return m.isManagementRestricted
}
// GetIsResourceAccount gets the isResourceAccount property value. Don't use – reserved for future use.
// returns a *bool when successful
func (m *User) GetIsResourceAccount()(*bool) {
    return m.isResourceAccount
}
// GetJobTitle gets the jobTitle property value. The user's job title. Maximum length is 128 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetJobTitle()(*string) {
    return m.jobTitle
}
// GetJoinedTeams gets the joinedTeams property value. The joinedTeams property
// returns a []Teamable when successful
func (m *User) GetJoinedTeams()([]Teamable) {
    return m.joinedTeams
}
// GetLastPasswordChangeDateTime gets the lastPasswordChangeDateTime property value. The time when this Microsoft Entra user last changed their password or when their password was created, whichever date the latest action was performed. The date and time information uses ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
// returns a *Time when successful
func (m *User) GetLastPasswordChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastPasswordChangeDateTime
}
// GetLegalAgeGroupClassification gets the legalAgeGroupClassification property value. Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, Undefined,  MinorWithOutParentalConsent, MinorWithParentalConsent, MinorNoParentalConsentRequired, NotAdult, and Adult. For more information, see legal age group property definitions. Returned only on $select.
// returns a *string when successful
func (m *User) GetLegalAgeGroupClassification()(*string) {
    return m.legalAgeGroupClassification
}
// GetLicenseAssignmentStates gets the licenseAssignmentStates property value. State of license assignments for this user. Also indicates licenses that are directly assigned or the user inherited through group memberships. Read-only. Returned only on $select.
// returns a []LicenseAssignmentStateable when successful
func (m *User) GetLicenseAssignmentStates()([]LicenseAssignmentStateable) {
    return m.licenseAssignmentStates
}
// GetLicenseDetails gets the licenseDetails property value. A collection of this user's license details. Read-only.
// returns a []LicenseDetailsable when successful
func (m *User) GetLicenseDetails()([]LicenseDetailsable) {
    return m.licenseDetails
}
// GetMail gets the mail property value. The SMTP address for the user, for example, jeff@contoso.com. Changes to this property update the user's proxyAddresses collection to include the value as an SMTP address. This property can't contain accent characters.  NOTE: We don't recommend updating this property for Azure AD B2C user profiles. Use the otherMails property instead. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetMail()(*string) {
    return m.mail
}
// GetMailboxSettings gets the mailboxSettings property value. Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. Returned only on $select.
// returns a MailboxSettingsable when successful
func (m *User) GetMailboxSettings()(MailboxSettingsable) {
    return m.mailboxSettings
}
// GetMailFolders gets the mailFolders property value. The user's mail folders. Read-only. Nullable.
// returns a []MailFolderable when successful
func (m *User) GetMailFolders()([]MailFolderable) {
    return m.mailFolders
}
// GetMailNickname gets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetManagedAppRegistrations gets the managedAppRegistrations property value. Zero or more managed app registrations that belong to the user.
// returns a []ManagedAppRegistrationable when successful
func (m *User) GetManagedAppRegistrations()([]ManagedAppRegistrationable) {
    return m.managedAppRegistrations
}
// GetManagedDevices gets the managedDevices property value. The managed devices associated with the user.
// returns a []ManagedDeviceable when successful
func (m *User) GetManagedDevices()([]ManagedDeviceable) {
    return m.managedDevices
}
// GetManager gets the manager property value. The user or contact that is this user's manager. Read-only. Supports $expand.
// returns a DirectoryObjectable when successful
func (m *User) GetManager()(DirectoryObjectable) {
    return m.manager
}
// GetMemberOf gets the memberOf property value. The groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
// returns a []DirectoryObjectable when successful
func (m *User) GetMemberOf()([]DirectoryObjectable) {
    return m.memberOf
}
// GetMessages gets the messages property value. The messages in a mailbox or folder. Read-only. Nullable.
// returns a []Messageable when successful
func (m *User) GetMessages()([]Messageable) {
    return m.messages
}
// GetMobilePhone gets the mobilePhone property value. The primary cellular telephone number for the user. Read-only for users synced from the on-premises directory. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values) and $search.
// returns a *string when successful
func (m *User) GetMobilePhone()(*string) {
    return m.mobilePhone
}
// GetMySite gets the mySite property value. The URL for the user's site. Returned only on $select.
// returns a *string when successful
func (m *User) GetMySite()(*string) {
    return m.mySite
}
// GetOauth2PermissionGrants gets the oauth2PermissionGrants property value. The oauth2PermissionGrants property
// returns a []OAuth2PermissionGrantable when successful
func (m *User) GetOauth2PermissionGrants()([]OAuth2PermissionGrantable) {
    return m.oauth2PermissionGrants
}
// GetOfficeLocation gets the officeLocation property value. The office location in the user's place of business. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetOfficeLocation()(*string) {
    return m.officeLocation
}
// GetOnenote gets the onenote property value. The onenote property
// returns a Onenoteable when successful
func (m *User) GetOnenote()(Onenoteable) {
    return m.onenote
}
// GetOnlineMeetings gets the onlineMeetings property value. Information about a meeting, including the URL used to join a meeting, the attendees list, and the description.
// returns a []OnlineMeetingable when successful
func (m *User) GetOnlineMeetings()([]OnlineMeetingable) {
    return m.onlineMeetings
}
// GetOnPremisesDistinguishedName gets the onPremisesDistinguishedName property value. Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select.
// returns a *string when successful
func (m *User) GetOnPremisesDistinguishedName()(*string) {
    return m.onPremisesDistinguishedName
}
// GetOnPremisesDomainName gets the onPremisesDomainName property value. Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select.
// returns a *string when successful
func (m *User) GetOnPremisesDomainName()(*string) {
    return m.onPremisesDomainName
}
// GetOnPremisesExtensionAttributes gets the onPremisesExtensionAttributes property value. Contains extensionAttributes1-15 for the user. These extension attributes are also known as Exchange custom attributes 1-15. Each attribute can store up to 1024 characters. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties can be set during the creation or update of a user object.  For a cloud-only user previously synced from on-premises Active Directory, these properties are read-only in Microsoft Graph but can be fully managed through the Exchange Admin Center or the Exchange Online V2 module in PowerShell. Returned only on $select. Supports $filter (eq, ne, not, in).
// returns a OnPremisesExtensionAttributesable when successful
func (m *User) GetOnPremisesExtensionAttributes()(OnPremisesExtensionAttributesable) {
    return m.onPremisesExtensionAttributes
}
// GetOnPremisesImmutableId gets the onPremisesImmutableId property value. This property is used to associate an on-premises Active Directory user account to their Microsoft Entra user object. This property must be specified when creating a new user account in the Graph if you're using a federated domain for the user's userPrincipalName (UPN) property. NOTE: The $ and _ characters can't be used when specifying this property. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).
// returns a *string when successful
func (m *User) GetOnPremisesImmutableId()(*string) {
    return m.onPremisesImmutableId
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. Indicates the last time at which the object was synced with the on-premises directory; for example: 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).
// returns a *Time when successful
func (m *User) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.onPremisesLastSyncDateTime
}
// GetOnPremisesProvisioningErrors gets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned only on $select. Supports $filter (eq, not, ge, le).
// returns a []OnPremisesProvisioningErrorable when successful
func (m *User) GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningErrorable) {
    return m.onPremisesProvisioningErrors
}
// GetOnPremisesSamAccountName gets the onPremisesSamAccountName property value. Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
// returns a *string when successful
func (m *User) GetOnPremisesSamAccountName()(*string) {
    return m.onPremisesSamAccountName
}
// GetOnPremisesSecurityIdentifier gets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Returned only on $select. Supports $filter (eq including on null values).
// returns a *string when successful
func (m *User) GetOnPremisesSecurityIdentifier()(*string) {
    return m.onPremisesSecurityIdentifier
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this user object is currently being synced from an on-premises Active Directory (AD); otherwise the user isn't being synced and can be managed in Microsoft Entra ID. Read-only. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values).
// returns a *bool when successful
func (m *User) GetOnPremisesSyncEnabled()(*bool) {
    return m.onPremisesSyncEnabled
}
// GetOnPremisesUserPrincipalName gets the onPremisesUserPrincipalName property value. Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
// returns a *string when successful
func (m *User) GetOnPremisesUserPrincipalName()(*string) {
    return m.onPremisesUserPrincipalName
}
// GetOtherMails gets the otherMails property value. A list of other email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com']. NOTE: This property can't contain accent characters. Returned only on $select. Supports $filter (eq, not, ge, le, in, startsWith, endsWith, /$count eq 0, /$count ne 0).
// returns a []string when successful
func (m *User) GetOtherMails()([]string) {
    return m.otherMails
}
// GetOutlook gets the outlook property value. The outlook property
// returns a OutlookUserable when successful
func (m *User) GetOutlook()(OutlookUserable) {
    return m.outlook
}
// GetOwnedDevices gets the ownedDevices property value. Devices the user owns. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
// returns a []DirectoryObjectable when successful
func (m *User) GetOwnedDevices()([]DirectoryObjectable) {
    return m.ownedDevices
}
// GetOwnedObjects gets the ownedObjects property value. Directory objects the user owns. Read-only. Nullable. Supports $expand, $select nested in $expand, and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
// returns a []DirectoryObjectable when successful
func (m *User) GetOwnedObjects()([]DirectoryObjectable) {
    return m.ownedObjects
}
// GetPasswordPolicies gets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two might be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. Returned only on $select. For more information on the default password policies, see Microsoft Entra password policies. Supports $filter (ne, not, and eq on null values).
// returns a *string when successful
func (m *User) GetPasswordPolicies()(*string) {
    return m.passwordPolicies
}
// GetPasswordProfile gets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). To update this property:  In delegated access, the calling app must be assigned the User-PasswordProfile.ReadWrite.All delegated permission on behalf of the signed-in user.  In application-only access, the calling app must be assigned the User-PasswordProfile.ReadWrite.All application permission and at least the User Administrator Microsoft Entra role.
// returns a PasswordProfileable when successful
func (m *User) GetPasswordProfile()(PasswordProfileable) {
    return m.passwordProfile
}
// GetPastProjects gets the pastProjects property value. A list for the user to enumerate their past projects. Returned only on $select.
// returns a []string when successful
func (m *User) GetPastProjects()([]string) {
    return m.pastProjects
}
// GetPeople gets the people property value. People that are relevant to the user. Read-only. Nullable.
// returns a []Personable when successful
func (m *User) GetPeople()([]Personable) {
    return m.people
}
// GetPermissionGrants gets the permissionGrants property value. List all resource-specific permission grants of a user.
// returns a []ResourceSpecificPermissionGrantable when successful
func (m *User) GetPermissionGrants()([]ResourceSpecificPermissionGrantable) {
    return m.permissionGrants
}
// GetPhoto gets the photo property value. The user's profile photo. Read-only.
// returns a ProfilePhotoable when successful
func (m *User) GetPhoto()(ProfilePhotoable) {
    return m.photo
}
// GetPhotos gets the photos property value. The collection of the user's profile photos in different sizes. Read-only.
// returns a []ProfilePhotoable when successful
func (m *User) GetPhotos()([]ProfilePhotoable) {
    return m.photos
}
// GetPlanner gets the planner property value. Entry-point to the Planner resource that might exist for a user. Read-only.
// returns a PlannerUserable when successful
func (m *User) GetPlanner()(PlannerUserable) {
    return m.planner
}
// GetPostalCode gets the postalCode property value. The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetPostalCode()(*string) {
    return m.postalCode
}
// GetPreferredDataLocation gets the preferredDataLocation property value. The preferred data location for the user. For more information, see OneDrive Online Multi-Geo.
// returns a *string when successful
func (m *User) GetPreferredDataLocation()(*string) {
    return m.preferredDataLocation
}
// GetPreferredLanguage gets the preferredLanguage property value. The preferred language for the user. The preferred language format is based on RFC 4646. The name is a combination of an ISO 639 two-letter lowercase culture code associated with the language, and an ISO 3166 two-letter uppercase subculture code associated with the country or region. Example: 'en-US', or 'es-ES'. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values)
// returns a *string when successful
func (m *User) GetPreferredLanguage()(*string) {
    return m.preferredLanguage
}
// GetPreferredName gets the preferredName property value. The preferred name for the user. Not Supported. This attribute returns an empty string.Returned only on $select.
// returns a *string when successful
func (m *User) GetPreferredName()(*string) {
    return m.preferredName
}
// GetPresence gets the presence property value. The presence property
// returns a Presenceable when successful
func (m *User) GetPresence()(Presenceable) {
    return m.presence
}
// GetPrint gets the print property value. The print property
// returns a UserPrintable when successful
func (m *User) GetPrint()(UserPrintable) {
    return m.print
}
// GetProvisionedPlans gets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le).
// returns a []ProvisionedPlanable when successful
func (m *User) GetProvisionedPlans()([]ProvisionedPlanable) {
    return m.provisionedPlans
}
// GetProxyAddresses gets the proxyAddresses property value. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. Changes to the mail property update this collection to include the value as an SMTP address. For more information, see mail and proxyAddresses properties. The proxy address prefixed with SMTP (capitalized) is the primary proxy address, while those addresses prefixed with smtp are the secondary proxy addresses. For Azure AD B2C accounts, this property has a limit of 10 unique addresses. Read-only in Microsoft Graph; you can update this property only through the Microsoft 365 admin center. Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith, endsWith, /$count eq 0, /$count ne 0).
// returns a []string when successful
func (m *User) GetProxyAddresses()([]string) {
    return m.proxyAddresses
}
// GetRegisteredDevices gets the registeredDevices property value. Devices that are registered for the user. Read-only. Nullable. Supports $expand and returns up to 100 objects.
// returns a []DirectoryObjectable when successful
func (m *User) GetRegisteredDevices()([]DirectoryObjectable) {
    return m.registeredDevices
}
// GetResponsibilities gets the responsibilities property value. A list for the user to enumerate their responsibilities. Returned only on $select.
// returns a []string when successful
func (m *User) GetResponsibilities()([]string) {
    return m.responsibilities
}
// GetSchools gets the schools property value. A list for the user to enumerate the schools they attended. Returned only on $select.
// returns a []string when successful
func (m *User) GetSchools()([]string) {
    return m.schools
}
// GetScopedRoleMemberOf gets the scopedRoleMemberOf property value. The scopedRoleMemberOf property
// returns a []ScopedRoleMembershipable when successful
func (m *User) GetScopedRoleMemberOf()([]ScopedRoleMembershipable) {
    return m.scopedRoleMemberOf
}
// GetSecurityIdentifier gets the securityIdentifier property value. Security identifier (SID) of the user, used in Windows scenarios. Read-only. Returned by default. Supports $select and $filter (eq, not, ge, le, startsWith).
// returns a *string when successful
func (m *User) GetSecurityIdentifier()(*string) {
    return m.securityIdentifier
}
// GetServiceProvisioningErrors gets the serviceProvisioningErrors property value. Errors published by a federated service describing a nontransient, service-specific error regarding the properties or link from a user object.  Supports $filter (eq, not, for isResolved and serviceInstance).
// returns a []ServiceProvisioningErrorable when successful
func (m *User) GetServiceProvisioningErrors()([]ServiceProvisioningErrorable) {
    return m.serviceProvisioningErrors
}
// GetSettings gets the settings property value. The settings property
// returns a UserSettingsable when successful
func (m *User) GetSettings()(UserSettingsable) {
    return m.settings
}
// GetShowInAddressList gets the showInAddressList property value. Do not use in Microsoft Graph. Manage this property through the Microsoft 365 admin center instead. Represents whether the user should be included in the Outlook global address list. See Known issue.
// returns a *bool when successful
func (m *User) GetShowInAddressList()(*bool) {
    return m.showInAddressList
}
// GetSignInActivity gets the signInActivity property value. Get the last signed-in date and request ID of the sign-in for a given user. Read-only.Returned only on $select. Supports $filter (eq, ne, not, ge, le) but not with any other filterable properties. Note: Details for this property require a Microsoft Entra ID P1 or P2 license and the AuditLog.Read.All permission.This property isn't returned for a user who never signed in or last signed in before April 2020.
// returns a SignInActivityable when successful
func (m *User) GetSignInActivity()(SignInActivityable) {
    return m.signInActivity
}
// GetSignInSessionsValidFromDateTime gets the signInSessionsValidFromDateTime property value. Any refresh tokens or session tokens (session cookies) issued before this time are invalid. Applications get an error when using an invalid refresh or session token to acquire a delegated access token (to access APIs such as Microsoft Graph). If this happens, the application needs to acquire a new refresh token by requesting the authorized endpoint. Read-only. Use revokeSignInSessions to reset. Returned only on $select.
// returns a *Time when successful
func (m *User) GetSignInSessionsValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.signInSessionsValidFromDateTime
}
// GetSkills gets the skills property value. A list for the user to enumerate their skills. Returned only on $select.
// returns a []string when successful
func (m *User) GetSkills()([]string) {
    return m.skills
}
// GetSolutions gets the solutions property value. The identifier that relates the user to the working time schedule triggers. Read-Only. Nullable
// returns a UserSolutionRootable when successful
func (m *User) GetSolutions()(UserSolutionRootable) {
    return m.solutions
}
// GetSponsors gets the sponsors property value. The users and groups responsible for this guest's privileges in the tenant and keeping the guest's information and access updated. (HTTP Methods: GET, POST, DELETE.). Supports $expand.
// returns a []DirectoryObjectable when successful
func (m *User) GetSponsors()([]DirectoryObjectable) {
    return m.sponsors
}
// GetState gets the state property value. The state or province in the user's address. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetState()(*string) {
    return m.state
}
// GetStreetAddress gets the streetAddress property value. The street address of the user's place of business. Maximum length is 1,024 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetStreetAddress()(*string) {
    return m.streetAddress
}
// GetSurname gets the surname property value. The user's surname (family name or last name). Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetSurname()(*string) {
    return m.surname
}
// GetTeamwork gets the teamwork property value. A container for Microsoft Teams features available for the user. Read-only. Nullable.
// returns a UserTeamworkable when successful
func (m *User) GetTeamwork()(UserTeamworkable) {
    return m.teamwork
}
// GetTodo gets the todo property value. Represents the To Do services available to a user.
// returns a Todoable when successful
func (m *User) GetTodo()(Todoable) {
    return m.todo
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. The groups, including nested groups, and directory roles that a user is a member of. Nullable.
// returns a []DirectoryObjectable when successful
func (m *User) GetTransitiveMemberOf()([]DirectoryObjectable) {
    return m.transitiveMemberOf
}
// GetUsageLocation gets the usageLocation property value. A two-letter country code (ISO standard 3166). Required for users that are assigned licenses due to legal requirements to check for availability of services in countries. Examples include: US, JP, and GB. Not nullable. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *User) GetUsageLocation()(*string) {
    return m.usageLocation
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style sign-in name for the user based on the Internet standard RFC 822. By convention, this value should map to the user's email name. The general format is alias@domain, where the domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property can't contain accent characters. Only the following characters are allowed A - Z, a - z, 0 - 9, ' . - _ ! # ^ ~. For the complete list of allowed characters, see username policies. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith) and $orderby.
// returns a *string when successful
func (m *User) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUserType gets the userType property value. A string value that can be used to classify user types in your directory. The possible values are Member and Guest. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). NOTE: For more information about the permissions for members and guests, see What are the default user permissions in Microsoft Entra ID?
// returns a *string when successful
func (m *User) GetUserType()(*string) {
    return m.userType
}
// Serialize serializes information the current object
func (m *User) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("aboutMe", m.GetAboutMe())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivities()))
        for i, v := range m.GetActivities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("activities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ageGroup", m.GetAgeGroup())
        if err != nil {
            return err
        }
    }
    if m.GetAgreementAcceptances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAgreementAcceptances()))
        for i, v := range m.GetAgreementAcceptances() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("agreementAcceptances", cast)
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
    if m.GetAssignedLicenses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedLicenses()))
        for i, v := range m.GetAssignedLicenses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignedLicenses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignedPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedPlans()))
        for i, v := range m.GetAssignedPlans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignedPlans", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authentication", m.GetAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authorizationInfo", m.GetAuthorizationInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("birthday", m.GetBirthday())
        if err != nil {
            return err
        }
    }
    if m.GetBusinessPhones() != nil {
        err = writer.WriteCollectionOfStringValues("businessPhones", m.GetBusinessPhones())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("calendar", m.GetCalendar())
        if err != nil {
            return err
        }
    }
    if m.GetCalendarGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCalendarGroups()))
        for i, v := range m.GetCalendarGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("calendarGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCalendars() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCalendars()))
        for i, v := range m.GetCalendars() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("calendars", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCalendarView() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCalendarView()))
        for i, v := range m.GetCalendarView() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
        if err != nil {
            return err
        }
    }
    if m.GetChats() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChats()))
        for i, v := range m.GetChats() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("chats", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("cloudClipboard", m.GetCloudClipboard())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyName", m.GetCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("consentProvidedForMinor", m.GetConsentProvidedForMinor())
        if err != nil {
            return err
        }
    }
    if m.GetContactFolders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContactFolders()))
        for i, v := range m.GetContactFolders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("contactFolders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContacts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContacts()))
        for i, v := range m.GetContacts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("contacts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("country", m.GetCountry())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err = writer.WriteStringValue("creationType", m.GetCreationType())
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
    {
        err = writer.WriteStringValue("department", m.GetDepartment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceEnrollmentLimit", m.GetDeviceEnrollmentLimit())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementTroubleshootingEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceManagementTroubleshootingEvents()))
        for i, v := range m.GetDeviceManagementTroubleshootingEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceManagementTroubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDirectReports() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDirectReports()))
        for i, v := range m.GetDirectReports() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("directReports", cast)
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
    {
        err = writer.WriteObjectValue("drive", m.GetDrive())
        if err != nil {
            return err
        }
    }
    if m.GetDrives() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDrives()))
        for i, v := range m.GetDrives() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("drives", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("employeeExperience", m.GetEmployeeExperience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("employeeHireDate", m.GetEmployeeHireDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("employeeId", m.GetEmployeeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("employeeLeaveDateTime", m.GetEmployeeLeaveDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("employeeOrgData", m.GetEmployeeOrgData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("employeeType", m.GetEmployeeType())
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalUserState", m.GetExternalUserState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("externalUserStateChangeDateTime", m.GetExternalUserStateChangeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("faxNumber", m.GetFaxNumber())
        if err != nil {
            return err
        }
    }
    if m.GetFollowedSites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFollowedSites()))
        for i, v := range m.GetFollowedSites() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("followedSites", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("hireDate", m.GetHireDate())
        if err != nil {
            return err
        }
    }
    if m.GetIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIdentities()))
        for i, v := range m.GetIdentities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("identities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("imAddresses", m.GetImAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inferenceClassification", m.GetInferenceClassification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("insights", m.GetInsights())
        if err != nil {
            return err
        }
    }
    if m.GetInterests() != nil {
        err = writer.WriteCollectionOfStringValues("interests", m.GetInterests())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isManagementRestricted", m.GetIsManagementRestricted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isResourceAccount", m.GetIsResourceAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    if m.GetJoinedTeams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJoinedTeams()))
        for i, v := range m.GetJoinedTeams() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("joinedTeams", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastPasswordChangeDateTime", m.GetLastPasswordChangeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("legalAgeGroupClassification", m.GetLegalAgeGroupClassification())
        if err != nil {
            return err
        }
    }
    if m.GetLicenseAssignmentStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLicenseAssignmentStates()))
        for i, v := range m.GetLicenseAssignmentStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("licenseAssignmentStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLicenseDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLicenseDetails()))
        for i, v := range m.GetLicenseDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("licenseDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mail", m.GetMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mailboxSettings", m.GetMailboxSettings())
        if err != nil {
            return err
        }
    }
    if m.GetMailFolders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMailFolders()))
        for i, v := range m.GetMailFolders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mailFolders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppRegistrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedAppRegistrations()))
        for i, v := range m.GetManagedAppRegistrations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("managedAppRegistrations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDevices()))
        for i, v := range m.GetManagedDevices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("managedDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("manager", m.GetManager())
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
    if m.GetMessages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobilePhone", m.GetMobilePhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mySite", m.GetMySite())
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
    {
        err = writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onenote", m.GetOnenote())
        if err != nil {
            return err
        }
    }
    if m.GetOnlineMeetings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnlineMeetings()))
        for i, v := range m.GetOnlineMeetings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("onlineMeetings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesDistinguishedName", m.GetOnPremisesDistinguishedName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesDomainName", m.GetOnPremisesDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onPremisesExtensionAttributes", m.GetOnPremisesExtensionAttributes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesImmutableId", m.GetOnPremisesImmutableId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOnPremisesProvisioningErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnPremisesProvisioningErrors()))
        for i, v := range m.GetOnPremisesProvisioningErrors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("onPremisesProvisioningErrors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSamAccountName", m.GetOnPremisesSamAccountName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSecurityIdentifier", m.GetOnPremisesSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesUserPrincipalName", m.GetOnPremisesUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetOtherMails() != nil {
        err = writer.WriteCollectionOfStringValues("otherMails", m.GetOtherMails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("outlook", m.GetOutlook())
        if err != nil {
            return err
        }
    }
    if m.GetOwnedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOwnedDevices()))
        for i, v := range m.GetOwnedDevices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("ownedDevices", cast)
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
    {
        err = writer.WriteStringValue("passwordPolicies", m.GetPasswordPolicies())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("passwordProfile", m.GetPasswordProfile())
        if err != nil {
            return err
        }
    }
    if m.GetPastProjects() != nil {
        err = writer.WriteCollectionOfStringValues("pastProjects", m.GetPastProjects())
        if err != nil {
            return err
        }
    }
    if m.GetPeople() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPeople()))
        for i, v := range m.GetPeople() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("people", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPermissionGrants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissionGrants()))
        for i, v := range m.GetPermissionGrants() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("permissionGrants", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    if m.GetPhotos() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhotos()))
        for i, v := range m.GetPhotos() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("photos", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("planner", m.GetPlanner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postalCode", m.GetPostalCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredDataLocation", m.GetPreferredDataLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredLanguage", m.GetPreferredLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredName", m.GetPreferredName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("presence", m.GetPresence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("print", m.GetPrint())
        if err != nil {
            return err
        }
    }
    if m.GetProvisionedPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProvisionedPlans()))
        for i, v := range m.GetProvisionedPlans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("provisionedPlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProxyAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("proxyAddresses", m.GetProxyAddresses())
        if err != nil {
            return err
        }
    }
    if m.GetRegisteredDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegisteredDevices()))
        for i, v := range m.GetRegisteredDevices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("registeredDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResponsibilities() != nil {
        err = writer.WriteCollectionOfStringValues("responsibilities", m.GetResponsibilities())
        if err != nil {
            return err
        }
    }
    if m.GetSchools() != nil {
        err = writer.WriteCollectionOfStringValues("schools", m.GetSchools())
        if err != nil {
            return err
        }
    }
    if m.GetScopedRoleMemberOf() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScopedRoleMemberOf()))
        for i, v := range m.GetScopedRoleMemberOf() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("scopedRoleMemberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("securityIdentifier", m.GetSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetServiceProvisioningErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServiceProvisioningErrors()))
        for i, v := range m.GetServiceProvisioningErrors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("serviceProvisioningErrors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showInAddressList", m.GetShowInAddressList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("signInActivity", m.GetSignInActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("signInSessionsValidFromDateTime", m.GetSignInSessionsValidFromDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSkills() != nil {
        err = writer.WriteCollectionOfStringValues("skills", m.GetSkills())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("solutions", m.GetSolutions())
        if err != nil {
            return err
        }
    }
    if m.GetSponsors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSponsors()))
        for i, v := range m.GetSponsors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sponsors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("streetAddress", m.GetStreetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamwork", m.GetTeamwork())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("todo", m.GetTodo())
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
        err = writer.WriteStringValue("usageLocation", m.GetUsageLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userType", m.GetUserType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAboutMe sets the aboutMe property value. A freeform text entry field for the user to describe themselves. Returned only on $select.
func (m *User) SetAboutMe(value *string)() {
    m.aboutMe = value
}
// SetAccountEnabled sets the accountEnabled property value. true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter (eq, ne, not, and in).
func (m *User) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// SetActivities sets the activities property value. The user's activities across devices. Read-only. Nullable.
func (m *User) SetActivities(value []UserActivityable)() {
    m.activities = value
}
// SetAgeGroup sets the ageGroup property value. Sets the age group of the user. Allowed values: null, Minor, NotAdult, and Adult. For more information, see legal age group property definitions. Returned only on $select. Supports $filter (eq, ne, not, and in).
func (m *User) SetAgeGroup(value *string)() {
    m.ageGroup = value
}
// SetAgreementAcceptances sets the agreementAcceptances property value. The user's terms of use acceptance statuses. Read-only. Nullable.
func (m *User) SetAgreementAcceptances(value []AgreementAcceptanceable)() {
    m.agreementAcceptances = value
}
// SetAppRoleAssignments sets the appRoleAssignments property value. Represents the app roles a user is granted for an application. Supports $expand.
func (m *User) SetAppRoleAssignments(value []AppRoleAssignmentable)() {
    m.appRoleAssignments = value
}
// SetAssignedLicenses sets the assignedLicenses property value. The licenses that are assigned to the user, including inherited (group-based) licenses. This property doesn't differentiate between directly assigned and inherited licenses. Use the licenseAssignmentStates property to identify the directly assigned and inherited licenses. Not nullable. Returned only on $select. Supports $filter (eq, not, /$count eq 0, /$count ne 0).
func (m *User) SetAssignedLicenses(value []AssignedLicenseable)() {
    m.assignedLicenses = value
}
// SetAssignedPlans sets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq and not).
func (m *User) SetAssignedPlans(value []AssignedPlanable)() {
    m.assignedPlans = value
}
// SetAuthentication sets the authentication property value. The authentication methods that are supported for the user.
func (m *User) SetAuthentication(value Authenticationable)() {
    m.authentication = value
}
// SetAuthorizationInfo sets the authorizationInfo property value. The authorizationInfo property
func (m *User) SetAuthorizationInfo(value AuthorizationInfoable)() {
    m.authorizationInfo = value
}
// SetBirthday sets the birthday property value. The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z. Returned only on $select.
func (m *User) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.birthday = value
}
// SetBusinessPhones sets the businessPhones property value. The telephone numbers for the user. NOTE: Although it's a string collection, only one number can be set for this property. Read-only for users synced from the on-premises directory. Returned by default. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) SetBusinessPhones(value []string)() {
    m.businessPhones = value
}
// SetCalendar sets the calendar property value. The user's primary calendar. Read-only.
func (m *User) SetCalendar(value Calendarable)() {
    m.calendar = value
}
// SetCalendarGroups sets the calendarGroups property value. The user's calendar groups. Read-only. Nullable.
func (m *User) SetCalendarGroups(value []CalendarGroupable)() {
    m.calendarGroups = value
}
// SetCalendars sets the calendars property value. The user's calendars. Read-only. Nullable.
func (m *User) SetCalendars(value []Calendarable)() {
    m.calendars = value
}
// SetCalendarView sets the calendarView property value. The calendar view for the calendar. Read-only. Nullable.
func (m *User) SetCalendarView(value []Eventable)() {
    m.calendarView = value
}
// SetChats sets the chats property value. The chats property
func (m *User) SetChats(value []Chatable)() {
    m.chats = value
}
// SetCity sets the city property value. The city where the user is located. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetCity(value *string)() {
    m.city = value
}
// SetCloudClipboard sets the cloudClipboard property value. The cloudClipboard property
func (m *User) SetCloudClipboard(value CloudClipboardRootable)() {
    m.cloudClipboard = value
}
// SetCompanyName sets the companyName property value. The name of the company that the user is associated with. This property can be useful for describing the company that a guest comes from. The maximum length is 64 characters.Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetCompanyName(value *string)() {
    m.companyName = value
}
// SetConsentProvidedForMinor sets the consentProvidedForMinor property value. Sets whether consent was obtained for minors. Allowed values: null, Granted, Denied, and NotRequired. For more information, see legal age group property definitions. Returned only on $select. Supports $filter (eq, ne, not, and in).
func (m *User) SetConsentProvidedForMinor(value *string)() {
    m.consentProvidedForMinor = value
}
// SetContactFolders sets the contactFolders property value. The user's contacts folders. Read-only. Nullable.
func (m *User) SetContactFolders(value []ContactFolderable)() {
    m.contactFolders = value
}
// SetContacts sets the contacts property value. The user's contacts. Read-only. Nullable.
func (m *User) SetContacts(value []Contactable)() {
    m.contacts = value
}
// SetCountry sets the country property value. The country/region where the user is located; for example, US or UK. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetCountry(value *string)() {
    m.country = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the user was created, in ISO 8601 format and UTC. The value can't be modified and is automatically populated when the entity is created. Nullable. For on-premises users, the value represents when they were first created in Microsoft Entra ID. Property is null for some users created before June 2018 and on-premises users that were synced to Microsoft Entra ID before June 2018. Read-only. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreatedObjects sets the createdObjects property value. Directory objects that the user created. Read-only. Nullable.
func (m *User) SetCreatedObjects(value []DirectoryObjectable)() {
    m.createdObjects = value
}
// SetCreationType sets the creationType property value. Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by a guest signing up through a link that is part of a user flow (SelfServiceSignUp). Read-only.Returned only on $select. Supports $filter (eq, ne, not, in).
func (m *User) SetCreationType(value *string)() {
    m.creationType = value
}
// SetCustomSecurityAttributes sets the customSecurityAttributes property value. An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith). The filter value is case-sensitive. To read this property, the calling app must be assigned the CustomSecAttributeAssignment.Read.All permission. To write this property, the calling app must be assigned the CustomSecAttributeAssignment.ReadWrite.All permissions. To read or write this property in delegated scenarios, the admin must be assigned the Attribute Assignment Administrator role.
func (m *User) SetCustomSecurityAttributes(value CustomSecurityAttributeValueable)() {
    m.customSecurityAttributes = value
}
// SetDepartment sets the department property value. The name of the department in which the user works. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, and eq on null values).
func (m *User) SetDepartment(value *string)() {
    m.department = value
}
// SetDeviceEnrollmentLimit sets the deviceEnrollmentLimit property value. The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
func (m *User) SetDeviceEnrollmentLimit(value *int32)() {
    m.deviceEnrollmentLimit = value
}
// SetDeviceManagementTroubleshootingEvents sets the deviceManagementTroubleshootingEvents property value. The list of troubleshooting events for this user.
func (m *User) SetDeviceManagementTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)() {
    m.deviceManagementTroubleshootingEvents = value
}
// SetDirectReports sets the directReports property value. The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
func (m *User) SetDirectReports(value []DirectoryObjectable)() {
    m.directReports = value
}
// SetDisplayName sets the displayName property value. The name displayed in the address book for the user. This value is usually the combination of the user's first name, middle initial, and family name. This property is required when a user is created and it can't be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values), $orderby, and $search.
func (m *User) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDrive sets the drive property value. The user's OneDrive. Read-only.
func (m *User) SetDrive(value Driveable)() {
    m.drive = value
}
// SetDrives sets the drives property value. A collection of drives available for this user. Read-only.
func (m *User) SetDrives(value []Driveable)() {
    m.drives = value
}
// SetEmployeeExperience sets the employeeExperience property value. The employeeExperience property
func (m *User) SetEmployeeExperience(value EmployeeExperienceUserable)() {
    m.employeeExperience = value
}
// SetEmployeeHireDate sets the employeeHireDate property value. The date and time when the user was hired or will start work in a future hire. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) SetEmployeeHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.employeeHireDate = value
}
// SetEmployeeId sets the employeeId property value. The employee identifier assigned to the user by the organization. The maximum length is 16 characters. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetEmployeeId(value *string)() {
    m.employeeId = value
}
// SetEmployeeLeaveDateTime sets the employeeLeaveDateTime property value. The date and time when the user left or will leave the organization. To read this property, the calling app must be assigned the User-LifeCycleInfo.Read.All permission. To write this property, the calling app must be assigned the User.Read.All and User-LifeCycleInfo.ReadWrite.All permissions. To read this property in delegated scenarios, the admin needs at least one of the following Microsoft Entra roles: Lifecycle Workflows Administrator (least privilege), Global Reader. To write this property in delegated scenarios, the admin needs the Global Administrator role. Supports $filter (eq, ne, not , ge, le, in). For more information, see Configure the employeeLeaveDateTime property for a user.
func (m *User) SetEmployeeLeaveDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.employeeLeaveDateTime = value
}
// SetEmployeeOrgData sets the employeeOrgData property value. Represents organization data (for example, division and costCenter) associated with a user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) SetEmployeeOrgData(value EmployeeOrgDataable)() {
    m.employeeOrgData = value
}
// SetEmployeeType sets the employeeType property value. Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith).
func (m *User) SetEmployeeType(value *string)() {
    m.employeeType = value
}
// SetEvents sets the events property value. The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *User) SetEvents(value []Eventable)() {
    m.events = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the user. Read-only. Supports $expand. Nullable.
func (m *User) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetExternalUserState sets the externalUserState property value. For a guest invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only on $select. Supports $filter (eq, ne, not , in).
func (m *User) SetExternalUserState(value *string)() {
    m.externalUserState = value
}
// SetExternalUserStateChangeDateTime sets the externalUserStateChangeDateTime property value. Shows the timestamp for the latest change to the externalUserState property. Returned only on $select. Supports $filter (eq, ne, not , in).
func (m *User) SetExternalUserStateChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.externalUserStateChangeDateTime = value
}
// SetFaxNumber sets the faxNumber property value. The fax number of the user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetFaxNumber(value *string)() {
    m.faxNumber = value
}
// SetFollowedSites sets the followedSites property value. The followedSites property
func (m *User) SetFollowedSites(value []Siteable)() {
    m.followedSites = value
}
// SetGivenName sets the givenName property value. The given name (first name) of the user. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetGivenName(value *string)() {
    m.givenName = value
}
// SetHireDate sets the hireDate property value. The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z. Returned only on $select.  Note: This property is specific to SharePoint in Microsoft 365. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
func (m *User) SetHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.hireDate = value
}
// SetIdentities sets the identities property value. Represents the identities that can be used to sign in to this user account. Microsoft (also known as a local account), organizations, or social identity providers such as Facebook, Google, and Microsoft can provide identity and tie it to a user account. It might contain multiple items with the same signInType value. Returned only on $select.  Supports $filter (eq) with limitations.
func (m *User) SetIdentities(value []ObjectIdentityable)() {
    m.identities = value
}
// SetImAddresses sets the imAddresses property value. The instant message voice-over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) SetImAddresses(value []string)() {
    m.imAddresses = value
}
// SetInferenceClassification sets the inferenceClassification property value. Relevance classification of the user's messages based on explicit designations that override inferred relevance or importance.
func (m *User) SetInferenceClassification(value InferenceClassificationable)() {
    m.inferenceClassification = value
}
// SetInsights sets the insights property value. Represents relationships between a user and items such as OneDrive for work or school documents, calculated using advanced analytics and machine learning techniques. Read-only. Nullable.
func (m *User) SetInsights(value ItemInsightsable)() {
    m.insights = value
}
// SetInterests sets the interests property value. A list for the user to describe their interests. Returned only on $select.
func (m *User) SetInterests(value []string)() {
    m.interests = value
}
// SetIsManagementRestricted sets the isManagementRestricted property value. The isManagementRestricted property
func (m *User) SetIsManagementRestricted(value *bool)() {
    m.isManagementRestricted = value
}
// SetIsResourceAccount sets the isResourceAccount property value. Don't use – reserved for future use.
func (m *User) SetIsResourceAccount(value *bool)() {
    m.isResourceAccount = value
}
// SetJobTitle sets the jobTitle property value. The user's job title. Maximum length is 128 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// SetJoinedTeams sets the joinedTeams property value. The joinedTeams property
func (m *User) SetJoinedTeams(value []Teamable)() {
    m.joinedTeams = value
}
// SetLastPasswordChangeDateTime sets the lastPasswordChangeDateTime property value. The time when this Microsoft Entra user last changed their password or when their password was created, whichever date the latest action was performed. The date and time information uses ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
func (m *User) SetLastPasswordChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastPasswordChangeDateTime = value
}
// SetLegalAgeGroupClassification sets the legalAgeGroupClassification property value. Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, Undefined,  MinorWithOutParentalConsent, MinorWithParentalConsent, MinorNoParentalConsentRequired, NotAdult, and Adult. For more information, see legal age group property definitions. Returned only on $select.
func (m *User) SetLegalAgeGroupClassification(value *string)() {
    m.legalAgeGroupClassification = value
}
// SetLicenseAssignmentStates sets the licenseAssignmentStates property value. State of license assignments for this user. Also indicates licenses that are directly assigned or the user inherited through group memberships. Read-only. Returned only on $select.
func (m *User) SetLicenseAssignmentStates(value []LicenseAssignmentStateable)() {
    m.licenseAssignmentStates = value
}
// SetLicenseDetails sets the licenseDetails property value. A collection of this user's license details. Read-only.
func (m *User) SetLicenseDetails(value []LicenseDetailsable)() {
    m.licenseDetails = value
}
// SetMail sets the mail property value. The SMTP address for the user, for example, jeff@contoso.com. Changes to this property update the user's proxyAddresses collection to include the value as an SMTP address. This property can't contain accent characters.  NOTE: We don't recommend updating this property for Azure AD B2C user profiles. Use the otherMails property instead. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith, and eq on null values).
func (m *User) SetMail(value *string)() {
    m.mail = value
}
// SetMailboxSettings sets the mailboxSettings property value. Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. Returned only on $select.
func (m *User) SetMailboxSettings(value MailboxSettingsable)() {
    m.mailboxSettings = value
}
// SetMailFolders sets the mailFolders property value. The user's mail folders. Read-only. Nullable.
func (m *User) SetMailFolders(value []MailFolderable)() {
    m.mailFolders = value
}
// SetMailNickname sets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetManagedAppRegistrations sets the managedAppRegistrations property value. Zero or more managed app registrations that belong to the user.
func (m *User) SetManagedAppRegistrations(value []ManagedAppRegistrationable)() {
    m.managedAppRegistrations = value
}
// SetManagedDevices sets the managedDevices property value. The managed devices associated with the user.
func (m *User) SetManagedDevices(value []ManagedDeviceable)() {
    m.managedDevices = value
}
// SetManager sets the manager property value. The user or contact that is this user's manager. Read-only. Supports $expand.
func (m *User) SetManager(value DirectoryObjectable)() {
    m.manager = value
}
// SetMemberOf sets the memberOf property value. The groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
func (m *User) SetMemberOf(value []DirectoryObjectable)() {
    m.memberOf = value
}
// SetMessages sets the messages property value. The messages in a mailbox or folder. Read-only. Nullable.
func (m *User) SetMessages(value []Messageable)() {
    m.messages = value
}
// SetMobilePhone sets the mobilePhone property value. The primary cellular telephone number for the user. Read-only for users synced from the on-premises directory. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values) and $search.
func (m *User) SetMobilePhone(value *string)() {
    m.mobilePhone = value
}
// SetMySite sets the mySite property value. The URL for the user's site. Returned only on $select.
func (m *User) SetMySite(value *string)() {
    m.mySite = value
}
// SetOauth2PermissionGrants sets the oauth2PermissionGrants property value. The oauth2PermissionGrants property
func (m *User) SetOauth2PermissionGrants(value []OAuth2PermissionGrantable)() {
    m.oauth2PermissionGrants = value
}
// SetOfficeLocation sets the officeLocation property value. The office location in the user's place of business. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// SetOnenote sets the onenote property value. The onenote property
func (m *User) SetOnenote(value Onenoteable)() {
    m.onenote = value
}
// SetOnlineMeetings sets the onlineMeetings property value. Information about a meeting, including the URL used to join a meeting, the attendees list, and the description.
func (m *User) SetOnlineMeetings(value []OnlineMeetingable)() {
    m.onlineMeetings = value
}
// SetOnPremisesDistinguishedName sets the onPremisesDistinguishedName property value. Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select.
func (m *User) SetOnPremisesDistinguishedName(value *string)() {
    m.onPremisesDistinguishedName = value
}
// SetOnPremisesDomainName sets the onPremisesDomainName property value. Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select.
func (m *User) SetOnPremisesDomainName(value *string)() {
    m.onPremisesDomainName = value
}
// SetOnPremisesExtensionAttributes sets the onPremisesExtensionAttributes property value. Contains extensionAttributes1-15 for the user. These extension attributes are also known as Exchange custom attributes 1-15. Each attribute can store up to 1024 characters. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties can be set during the creation or update of a user object.  For a cloud-only user previously synced from on-premises Active Directory, these properties are read-only in Microsoft Graph but can be fully managed through the Exchange Admin Center or the Exchange Online V2 module in PowerShell. Returned only on $select. Supports $filter (eq, ne, not, in).
func (m *User) SetOnPremisesExtensionAttributes(value OnPremisesExtensionAttributesable)() {
    m.onPremisesExtensionAttributes = value
}
// SetOnPremisesImmutableId sets the onPremisesImmutableId property value. This property is used to associate an on-premises Active Directory user account to their Microsoft Entra user object. This property must be specified when creating a new user account in the Graph if you're using a federated domain for the user's userPrincipalName (UPN) property. NOTE: The $ and _ characters can't be used when specifying this property. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).
func (m *User) SetOnPremisesImmutableId(value *string)() {
    m.onPremisesImmutableId = value
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. Indicates the last time at which the object was synced with the on-premises directory; for example: 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).
func (m *User) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// SetOnPremisesProvisioningErrors sets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned only on $select. Supports $filter (eq, not, ge, le).
func (m *User) SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningErrorable)() {
    m.onPremisesProvisioningErrors = value
}
// SetOnPremisesSamAccountName sets the onPremisesSamAccountName property value. Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
func (m *User) SetOnPremisesSamAccountName(value *string)() {
    m.onPremisesSamAccountName = value
}
// SetOnPremisesSecurityIdentifier sets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Returned only on $select. Supports $filter (eq including on null values).
func (m *User) SetOnPremisesSecurityIdentifier(value *string)() {
    m.onPremisesSecurityIdentifier = value
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this user object is currently being synced from an on-premises Active Directory (AD); otherwise the user isn't being synced and can be managed in Microsoft Entra ID. Read-only. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *User) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// SetOnPremisesUserPrincipalName sets the onPremisesUserPrincipalName property value. Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
func (m *User) SetOnPremisesUserPrincipalName(value *string)() {
    m.onPremisesUserPrincipalName = value
}
// SetOtherMails sets the otherMails property value. A list of other email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com']. NOTE: This property can't contain accent characters. Returned only on $select. Supports $filter (eq, not, ge, le, in, startsWith, endsWith, /$count eq 0, /$count ne 0).
func (m *User) SetOtherMails(value []string)() {
    m.otherMails = value
}
// SetOutlook sets the outlook property value. The outlook property
func (m *User) SetOutlook(value OutlookUserable)() {
    m.outlook = value
}
// SetOwnedDevices sets the ownedDevices property value. Devices the user owns. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
func (m *User) SetOwnedDevices(value []DirectoryObjectable)() {
    m.ownedDevices = value
}
// SetOwnedObjects sets the ownedObjects property value. Directory objects the user owns. Read-only. Nullable. Supports $expand, $select nested in $expand, and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
func (m *User) SetOwnedObjects(value []DirectoryObjectable)() {
    m.ownedObjects = value
}
// SetPasswordPolicies sets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two might be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. Returned only on $select. For more information on the default password policies, see Microsoft Entra password policies. Supports $filter (ne, not, and eq on null values).
func (m *User) SetPasswordPolicies(value *string)() {
    m.passwordPolicies = value
}
// SetPasswordProfile sets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). To update this property:  In delegated access, the calling app must be assigned the User-PasswordProfile.ReadWrite.All delegated permission on behalf of the signed-in user.  In application-only access, the calling app must be assigned the User-PasswordProfile.ReadWrite.All application permission and at least the User Administrator Microsoft Entra role.
func (m *User) SetPasswordProfile(value PasswordProfileable)() {
    m.passwordProfile = value
}
// SetPastProjects sets the pastProjects property value. A list for the user to enumerate their past projects. Returned only on $select.
func (m *User) SetPastProjects(value []string)() {
    m.pastProjects = value
}
// SetPeople sets the people property value. People that are relevant to the user. Read-only. Nullable.
func (m *User) SetPeople(value []Personable)() {
    m.people = value
}
// SetPermissionGrants sets the permissionGrants property value. List all resource-specific permission grants of a user.
func (m *User) SetPermissionGrants(value []ResourceSpecificPermissionGrantable)() {
    m.permissionGrants = value
}
// SetPhoto sets the photo property value. The user's profile photo. Read-only.
func (m *User) SetPhoto(value ProfilePhotoable)() {
    m.photo = value
}
// SetPhotos sets the photos property value. The collection of the user's profile photos in different sizes. Read-only.
func (m *User) SetPhotos(value []ProfilePhotoable)() {
    m.photos = value
}
// SetPlanner sets the planner property value. Entry-point to the Planner resource that might exist for a user. Read-only.
func (m *User) SetPlanner(value PlannerUserable)() {
    m.planner = value
}
// SetPostalCode sets the postalCode property value. The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetPostalCode(value *string)() {
    m.postalCode = value
}
// SetPreferredDataLocation sets the preferredDataLocation property value. The preferred data location for the user. For more information, see OneDrive Online Multi-Geo.
func (m *User) SetPreferredDataLocation(value *string)() {
    m.preferredDataLocation = value
}
// SetPreferredLanguage sets the preferredLanguage property value. The preferred language for the user. The preferred language format is based on RFC 4646. The name is a combination of an ISO 639 two-letter lowercase culture code associated with the language, and an ISO 3166 two-letter uppercase subculture code associated with the country or region. Example: 'en-US', or 'es-ES'. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values)
func (m *User) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
// SetPreferredName sets the preferredName property value. The preferred name for the user. Not Supported. This attribute returns an empty string.Returned only on $select.
func (m *User) SetPreferredName(value *string)() {
    m.preferredName = value
}
// SetPresence sets the presence property value. The presence property
func (m *User) SetPresence(value Presenceable)() {
    m.presence = value
}
// SetPrint sets the print property value. The print property
func (m *User) SetPrint(value UserPrintable)() {
    m.print = value
}
// SetProvisionedPlans sets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le).
func (m *User) SetProvisionedPlans(value []ProvisionedPlanable)() {
    m.provisionedPlans = value
}
// SetProxyAddresses sets the proxyAddresses property value. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. Changes to the mail property update this collection to include the value as an SMTP address. For more information, see mail and proxyAddresses properties. The proxy address prefixed with SMTP (capitalized) is the primary proxy address, while those addresses prefixed with smtp are the secondary proxy addresses. For Azure AD B2C accounts, this property has a limit of 10 unique addresses. Read-only in Microsoft Graph; you can update this property only through the Microsoft 365 admin center. Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith, endsWith, /$count eq 0, /$count ne 0).
func (m *User) SetProxyAddresses(value []string)() {
    m.proxyAddresses = value
}
// SetRegisteredDevices sets the registeredDevices property value. Devices that are registered for the user. Read-only. Nullable. Supports $expand and returns up to 100 objects.
func (m *User) SetRegisteredDevices(value []DirectoryObjectable)() {
    m.registeredDevices = value
}
// SetResponsibilities sets the responsibilities property value. A list for the user to enumerate their responsibilities. Returned only on $select.
func (m *User) SetResponsibilities(value []string)() {
    m.responsibilities = value
}
// SetSchools sets the schools property value. A list for the user to enumerate the schools they attended. Returned only on $select.
func (m *User) SetSchools(value []string)() {
    m.schools = value
}
// SetScopedRoleMemberOf sets the scopedRoleMemberOf property value. The scopedRoleMemberOf property
func (m *User) SetScopedRoleMemberOf(value []ScopedRoleMembershipable)() {
    m.scopedRoleMemberOf = value
}
// SetSecurityIdentifier sets the securityIdentifier property value. Security identifier (SID) of the user, used in Windows scenarios. Read-only. Returned by default. Supports $select and $filter (eq, not, ge, le, startsWith).
func (m *User) SetSecurityIdentifier(value *string)() {
    m.securityIdentifier = value
}
// SetServiceProvisioningErrors sets the serviceProvisioningErrors property value. Errors published by a federated service describing a nontransient, service-specific error regarding the properties or link from a user object.  Supports $filter (eq, not, for isResolved and serviceInstance).
func (m *User) SetServiceProvisioningErrors(value []ServiceProvisioningErrorable)() {
    m.serviceProvisioningErrors = value
}
// SetSettings sets the settings property value. The settings property
func (m *User) SetSettings(value UserSettingsable)() {
    m.settings = value
}
// SetShowInAddressList sets the showInAddressList property value. Do not use in Microsoft Graph. Manage this property through the Microsoft 365 admin center instead. Represents whether the user should be included in the Outlook global address list. See Known issue.
func (m *User) SetShowInAddressList(value *bool)() {
    m.showInAddressList = value
}
// SetSignInActivity sets the signInActivity property value. Get the last signed-in date and request ID of the sign-in for a given user. Read-only.Returned only on $select. Supports $filter (eq, ne, not, ge, le) but not with any other filterable properties. Note: Details for this property require a Microsoft Entra ID P1 or P2 license and the AuditLog.Read.All permission.This property isn't returned for a user who never signed in or last signed in before April 2020.
func (m *User) SetSignInActivity(value SignInActivityable)() {
    m.signInActivity = value
}
// SetSignInSessionsValidFromDateTime sets the signInSessionsValidFromDateTime property value. Any refresh tokens or session tokens (session cookies) issued before this time are invalid. Applications get an error when using an invalid refresh or session token to acquire a delegated access token (to access APIs such as Microsoft Graph). If this happens, the application needs to acquire a new refresh token by requesting the authorized endpoint. Read-only. Use revokeSignInSessions to reset. Returned only on $select.
func (m *User) SetSignInSessionsValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.signInSessionsValidFromDateTime = value
}
// SetSkills sets the skills property value. A list for the user to enumerate their skills. Returned only on $select.
func (m *User) SetSkills(value []string)() {
    m.skills = value
}
// SetSolutions sets the solutions property value. The identifier that relates the user to the working time schedule triggers. Read-Only. Nullable
func (m *User) SetSolutions(value UserSolutionRootable)() {
    m.solutions = value
}
// SetSponsors sets the sponsors property value. The users and groups responsible for this guest's privileges in the tenant and keeping the guest's information and access updated. (HTTP Methods: GET, POST, DELETE.). Supports $expand.
func (m *User) SetSponsors(value []DirectoryObjectable)() {
    m.sponsors = value
}
// SetState sets the state property value. The state or province in the user's address. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetState(value *string)() {
    m.state = value
}
// SetStreetAddress sets the streetAddress property value. The street address of the user's place of business. Maximum length is 1,024 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetStreetAddress(value *string)() {
    m.streetAddress = value
}
// SetSurname sets the surname property value. The user's surname (family name or last name). Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetSurname(value *string)() {
    m.surname = value
}
// SetTeamwork sets the teamwork property value. A container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *User) SetTeamwork(value UserTeamworkable)() {
    m.teamwork = value
}
// SetTodo sets the todo property value. Represents the To Do services available to a user.
func (m *User) SetTodo(value Todoable)() {
    m.todo = value
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. The groups, including nested groups, and directory roles that a user is a member of. Nullable.
func (m *User) SetTransitiveMemberOf(value []DirectoryObjectable)() {
    m.transitiveMemberOf = value
}
// SetUsageLocation sets the usageLocation property value. A two-letter country code (ISO standard 3166). Required for users that are assigned licenses due to legal requirements to check for availability of services in countries. Examples include: US, JP, and GB. Not nullable. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetUsageLocation(value *string)() {
    m.usageLocation = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style sign-in name for the user based on the Internet standard RFC 822. By convention, this value should map to the user's email name. The general format is alias@domain, where the domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property can't contain accent characters. Only the following characters are allowed A - Z, a - z, 0 - 9, ' . - _ ! # ^ ~. For the complete list of allowed characters, see username policies. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith) and $orderby.
func (m *User) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUserType sets the userType property value. A string value that can be used to classify user types in your directory. The possible values are Member and Guest. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). NOTE: For more information about the permissions for members and guests, see What are the default user permissions in Microsoft Entra ID?
func (m *User) SetUserType(value *string)() {
    m.userType = value
}
type Userable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAboutMe()(*string)
    GetAccountEnabled()(*bool)
    GetActivities()([]UserActivityable)
    GetAgeGroup()(*string)
    GetAgreementAcceptances()([]AgreementAcceptanceable)
    GetAppRoleAssignments()([]AppRoleAssignmentable)
    GetAssignedLicenses()([]AssignedLicenseable)
    GetAssignedPlans()([]AssignedPlanable)
    GetAuthentication()(Authenticationable)
    GetAuthorizationInfo()(AuthorizationInfoable)
    GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetBusinessPhones()([]string)
    GetCalendar()(Calendarable)
    GetCalendarGroups()([]CalendarGroupable)
    GetCalendars()([]Calendarable)
    GetCalendarView()([]Eventable)
    GetChats()([]Chatable)
    GetCity()(*string)
    GetCloudClipboard()(CloudClipboardRootable)
    GetCompanyName()(*string)
    GetConsentProvidedForMinor()(*string)
    GetContactFolders()([]ContactFolderable)
    GetContacts()([]Contactable)
    GetCountry()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedObjects()([]DirectoryObjectable)
    GetCreationType()(*string)
    GetCustomSecurityAttributes()(CustomSecurityAttributeValueable)
    GetDepartment()(*string)
    GetDeviceEnrollmentLimit()(*int32)
    GetDeviceManagementTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable)
    GetDirectReports()([]DirectoryObjectable)
    GetDisplayName()(*string)
    GetDrive()(Driveable)
    GetDrives()([]Driveable)
    GetEmployeeExperience()(EmployeeExperienceUserable)
    GetEmployeeHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmployeeId()(*string)
    GetEmployeeLeaveDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmployeeOrgData()(EmployeeOrgDataable)
    GetEmployeeType()(*string)
    GetEvents()([]Eventable)
    GetExtensions()([]Extensionable)
    GetExternalUserState()(*string)
    GetExternalUserStateChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFaxNumber()(*string)
    GetFollowedSites()([]Siteable)
    GetGivenName()(*string)
    GetHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIdentities()([]ObjectIdentityable)
    GetImAddresses()([]string)
    GetInferenceClassification()(InferenceClassificationable)
    GetInsights()(ItemInsightsable)
    GetInterests()([]string)
    GetIsManagementRestricted()(*bool)
    GetIsResourceAccount()(*bool)
    GetJobTitle()(*string)
    GetJoinedTeams()([]Teamable)
    GetLastPasswordChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLegalAgeGroupClassification()(*string)
    GetLicenseAssignmentStates()([]LicenseAssignmentStateable)
    GetLicenseDetails()([]LicenseDetailsable)
    GetMail()(*string)
    GetMailboxSettings()(MailboxSettingsable)
    GetMailFolders()([]MailFolderable)
    GetMailNickname()(*string)
    GetManagedAppRegistrations()([]ManagedAppRegistrationable)
    GetManagedDevices()([]ManagedDeviceable)
    GetManager()(DirectoryObjectable)
    GetMemberOf()([]DirectoryObjectable)
    GetMessages()([]Messageable)
    GetMobilePhone()(*string)
    GetMySite()(*string)
    GetOauth2PermissionGrants()([]OAuth2PermissionGrantable)
    GetOfficeLocation()(*string)
    GetOnenote()(Onenoteable)
    GetOnlineMeetings()([]OnlineMeetingable)
    GetOnPremisesDistinguishedName()(*string)
    GetOnPremisesDomainName()(*string)
    GetOnPremisesExtensionAttributes()(OnPremisesExtensionAttributesable)
    GetOnPremisesImmutableId()(*string)
    GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningErrorable)
    GetOnPremisesSamAccountName()(*string)
    GetOnPremisesSecurityIdentifier()(*string)
    GetOnPremisesSyncEnabled()(*bool)
    GetOnPremisesUserPrincipalName()(*string)
    GetOtherMails()([]string)
    GetOutlook()(OutlookUserable)
    GetOwnedDevices()([]DirectoryObjectable)
    GetOwnedObjects()([]DirectoryObjectable)
    GetPasswordPolicies()(*string)
    GetPasswordProfile()(PasswordProfileable)
    GetPastProjects()([]string)
    GetPeople()([]Personable)
    GetPermissionGrants()([]ResourceSpecificPermissionGrantable)
    GetPhoto()(ProfilePhotoable)
    GetPhotos()([]ProfilePhotoable)
    GetPlanner()(PlannerUserable)
    GetPostalCode()(*string)
    GetPreferredDataLocation()(*string)
    GetPreferredLanguage()(*string)
    GetPreferredName()(*string)
    GetPresence()(Presenceable)
    GetPrint()(UserPrintable)
    GetProvisionedPlans()([]ProvisionedPlanable)
    GetProxyAddresses()([]string)
    GetRegisteredDevices()([]DirectoryObjectable)
    GetResponsibilities()([]string)
    GetSchools()([]string)
    GetScopedRoleMemberOf()([]ScopedRoleMembershipable)
    GetSecurityIdentifier()(*string)
    GetServiceProvisioningErrors()([]ServiceProvisioningErrorable)
    GetSettings()(UserSettingsable)
    GetShowInAddressList()(*bool)
    GetSignInActivity()(SignInActivityable)
    GetSignInSessionsValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSkills()([]string)
    GetSolutions()(UserSolutionRootable)
    GetSponsors()([]DirectoryObjectable)
    GetState()(*string)
    GetStreetAddress()(*string)
    GetSurname()(*string)
    GetTeamwork()(UserTeamworkable)
    GetTodo()(Todoable)
    GetTransitiveMemberOf()([]DirectoryObjectable)
    GetUsageLocation()(*string)
    GetUserPrincipalName()(*string)
    GetUserType()(*string)
    SetAboutMe(value *string)()
    SetAccountEnabled(value *bool)()
    SetActivities(value []UserActivityable)()
    SetAgeGroup(value *string)()
    SetAgreementAcceptances(value []AgreementAcceptanceable)()
    SetAppRoleAssignments(value []AppRoleAssignmentable)()
    SetAssignedLicenses(value []AssignedLicenseable)()
    SetAssignedPlans(value []AssignedPlanable)()
    SetAuthentication(value Authenticationable)()
    SetAuthorizationInfo(value AuthorizationInfoable)()
    SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetBusinessPhones(value []string)()
    SetCalendar(value Calendarable)()
    SetCalendarGroups(value []CalendarGroupable)()
    SetCalendars(value []Calendarable)()
    SetCalendarView(value []Eventable)()
    SetChats(value []Chatable)()
    SetCity(value *string)()
    SetCloudClipboard(value CloudClipboardRootable)()
    SetCompanyName(value *string)()
    SetConsentProvidedForMinor(value *string)()
    SetContactFolders(value []ContactFolderable)()
    SetContacts(value []Contactable)()
    SetCountry(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedObjects(value []DirectoryObjectable)()
    SetCreationType(value *string)()
    SetCustomSecurityAttributes(value CustomSecurityAttributeValueable)()
    SetDepartment(value *string)()
    SetDeviceEnrollmentLimit(value *int32)()
    SetDeviceManagementTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)()
    SetDirectReports(value []DirectoryObjectable)()
    SetDisplayName(value *string)()
    SetDrive(value Driveable)()
    SetDrives(value []Driveable)()
    SetEmployeeExperience(value EmployeeExperienceUserable)()
    SetEmployeeHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmployeeId(value *string)()
    SetEmployeeLeaveDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmployeeOrgData(value EmployeeOrgDataable)()
    SetEmployeeType(value *string)()
    SetEvents(value []Eventable)()
    SetExtensions(value []Extensionable)()
    SetExternalUserState(value *string)()
    SetExternalUserStateChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFaxNumber(value *string)()
    SetFollowedSites(value []Siteable)()
    SetGivenName(value *string)()
    SetHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIdentities(value []ObjectIdentityable)()
    SetImAddresses(value []string)()
    SetInferenceClassification(value InferenceClassificationable)()
    SetInsights(value ItemInsightsable)()
    SetInterests(value []string)()
    SetIsManagementRestricted(value *bool)()
    SetIsResourceAccount(value *bool)()
    SetJobTitle(value *string)()
    SetJoinedTeams(value []Teamable)()
    SetLastPasswordChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLegalAgeGroupClassification(value *string)()
    SetLicenseAssignmentStates(value []LicenseAssignmentStateable)()
    SetLicenseDetails(value []LicenseDetailsable)()
    SetMail(value *string)()
    SetMailboxSettings(value MailboxSettingsable)()
    SetMailFolders(value []MailFolderable)()
    SetMailNickname(value *string)()
    SetManagedAppRegistrations(value []ManagedAppRegistrationable)()
    SetManagedDevices(value []ManagedDeviceable)()
    SetManager(value DirectoryObjectable)()
    SetMemberOf(value []DirectoryObjectable)()
    SetMessages(value []Messageable)()
    SetMobilePhone(value *string)()
    SetMySite(value *string)()
    SetOauth2PermissionGrants(value []OAuth2PermissionGrantable)()
    SetOfficeLocation(value *string)()
    SetOnenote(value Onenoteable)()
    SetOnlineMeetings(value []OnlineMeetingable)()
    SetOnPremisesDistinguishedName(value *string)()
    SetOnPremisesDomainName(value *string)()
    SetOnPremisesExtensionAttributes(value OnPremisesExtensionAttributesable)()
    SetOnPremisesImmutableId(value *string)()
    SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningErrorable)()
    SetOnPremisesSamAccountName(value *string)()
    SetOnPremisesSecurityIdentifier(value *string)()
    SetOnPremisesSyncEnabled(value *bool)()
    SetOnPremisesUserPrincipalName(value *string)()
    SetOtherMails(value []string)()
    SetOutlook(value OutlookUserable)()
    SetOwnedDevices(value []DirectoryObjectable)()
    SetOwnedObjects(value []DirectoryObjectable)()
    SetPasswordPolicies(value *string)()
    SetPasswordProfile(value PasswordProfileable)()
    SetPastProjects(value []string)()
    SetPeople(value []Personable)()
    SetPermissionGrants(value []ResourceSpecificPermissionGrantable)()
    SetPhoto(value ProfilePhotoable)()
    SetPhotos(value []ProfilePhotoable)()
    SetPlanner(value PlannerUserable)()
    SetPostalCode(value *string)()
    SetPreferredDataLocation(value *string)()
    SetPreferredLanguage(value *string)()
    SetPreferredName(value *string)()
    SetPresence(value Presenceable)()
    SetPrint(value UserPrintable)()
    SetProvisionedPlans(value []ProvisionedPlanable)()
    SetProxyAddresses(value []string)()
    SetRegisteredDevices(value []DirectoryObjectable)()
    SetResponsibilities(value []string)()
    SetSchools(value []string)()
    SetScopedRoleMemberOf(value []ScopedRoleMembershipable)()
    SetSecurityIdentifier(value *string)()
    SetServiceProvisioningErrors(value []ServiceProvisioningErrorable)()
    SetSettings(value UserSettingsable)()
    SetShowInAddressList(value *bool)()
    SetSignInActivity(value SignInActivityable)()
    SetSignInSessionsValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSkills(value []string)()
    SetSolutions(value UserSolutionRootable)()
    SetSponsors(value []DirectoryObjectable)()
    SetState(value *string)()
    SetStreetAddress(value *string)()
    SetSurname(value *string)()
    SetTeamwork(value UserTeamworkable)()
    SetTodo(value Todoable)()
    SetTransitiveMemberOf(value []DirectoryObjectable)()
    SetUsageLocation(value *string)()
    SetUserPrincipalName(value *string)()
    SetUserType(value *string)()
}
