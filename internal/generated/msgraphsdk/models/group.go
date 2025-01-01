package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Group represents a Microsoft Entra group.
type Group struct {
    DirectoryObject
    // The list of users or groups allowed to create posts or calendar events in this group. If this list is nonempty, then only users or groups listed here are allowed to post.
    acceptedSenders []DirectoryObjectable
    // Indicates if people external to the organization can send messages to the group. The default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    allowExternalSenders *bool
    // Represents the app roles granted to a group for an application. Supports $expand.
    appRoleAssignments []AppRoleAssignmentable
    // The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on $select. This property can be updated only in delegated scenarios where the caller requires both the Microsoft Graph permission and a supported administrator role.
    assignedLabels []AssignedLabelable
    // The licenses that are assigned to the group. Returned only on $select. Supports $filter (eq).Read-only.
    assignedLicenses []AssignedLicenseable
    // Indicates if new members added to the group are autosubscribed to receive email notifications. You can set this property in a PATCH request for the group; don't set it in the initial POST request that creates the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    autoSubscribeNewMembers *bool
    // The group's calendar. Read-only.
    calendar Calendarable
    // The calendar view for the calendar. Read-only.
    calendarView []Eventable
    // Describes a classification for the group (such as low, medium, or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith).
    classification *string
    // The group's conversations.
    conversations []Conversationable
    // Timestamp of when the group was created. The value can't be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user (or application) that created the group. NOTE: This property isn't set if the user is an administrator. Read-only.
    createdOnBehalfOf DirectoryObjectable
    // An optional description for the group. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
    description *string
    // The display name for the group. This property is required when a group is created and can't be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderby.
    displayName *string
    // The group's default drive. Read-only.
    drive Driveable
    // The group's drives. Read-only.
    drives []Driveable
    // The group's calendar events.
    events []Eventable
    // Timestamp of when the group is set to expire. It's null for security groups, but for Microsoft 365 groups, it represents when the group is set to expire as defined in the groupLifecyclePolicy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The collection of open extensions defined for the group. Read-only. Nullable.
    extensions []Extensionable
    // The collection of lifecycle policies for this group. Read-only. Nullable.
    groupLifecyclePolicies []GroupLifecyclePolicyable
    // Specifies the group type and its membership. If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or a distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static. Returned by default. Supports $filter (eq, not).
    groupTypes []string
    // Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true). See an example. Supports $filter (eq).
    hasMembersWithLicenseErrors *bool
    // True if the group isn't displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false. The default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    hideFromAddressLists *bool
    // True if the group isn't displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false. The default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    hideFromOutlookClients *bool
    // When a group is associated with a team, this property determines whether the team is in read-only mode.To read this property, use the /group/{groupId}/team endpoint or the Get team API. To update this property, use the archiveTeam and unarchiveTeam APIs.
    isArchived *bool
    // Indicates whether this group can be assigned to a Microsoft Entra role. Optional. This property can only be set while creating the group and is immutable. If set to true, the securityEnabled property must also be set to true, visibility must be Hidden, and the group can't be a dynamic group (that is, groupTypes can't contain DynamicMembership). Only callers with at least the Privileged Role Administrator role can set this property. The caller must also be assigned the RoleManagement.ReadWrite.Directory permission to set this property or update the membership of such groups. For more, see Using a group to manage Microsoft Entra role assignmentsUsing this feature requires a Microsoft Entra ID P1 license. Returned by default. Supports $filter (eq, ne, not).
    isAssignableToRole *bool
    // The isManagementRestricted property
    isManagementRestricted *bool
    // Indicates whether the signed-in user is subscribed to receive email conversations. The default value is true. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    isSubscribedByMail *bool
    // Indicates the status of the group license assignment to all group members. The default value is false. Read-only. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select. Read-only.
    licenseProcessingState LicenseProcessingStateable
    // The SMTP address for the group, for example, 'serviceadmins@contoso.com'. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    mail *string
    // Specifies whether the group is mail-enabled. Required. Returned by default. Supports $filter (eq, ne, not).
    mailEnabled *bool
    // The mail alias for the group, unique for Microsoft 365 groups in the organization. Maximum length is 64 characters. This property can contain only characters in the ASCII character set 0 - 127 except the following characters: @ () / [] ' ; : <> , SPACE. Required. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    mailNickname *string
    // Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObjectable
    // The members of this group, who can be users, devices, other groups, or service principals. Supports the List members, Add member, and Remove member operations. Nullable. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=members($select=id,userPrincipalName,displayName).
    members []DirectoryObjectable
    // The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith).
    membershipRule *string
    // Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by default. Supports $filter (eq, ne, not, in).
    membershipRuleProcessingState *string
    // A list of group members with license errors from this group-based license assignment. Read-only.
    membersWithLicenseErrors []DirectoryObjectable
    // The onenote property
    onenote Onenoteable
    // Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect.Returned by default. Read-only.
    onPremisesDomainName *string
    // Indicates the last time at which the group was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Contains the on-premises netBios name synchronized from the on-premises directory. The property is only populated for customers synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect.Returned by default. Read-only.
    onPremisesNetBiosName *string
    // Errors when using Microsoft synchronization product during provisioning. Returned by default. Supports $filter (eq, not).
    onPremisesProvisioningErrors []OnPremisesProvisioningErrorable
    // Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect.Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith). Read-only.
    onPremisesSamAccountName *string
    // Contains the on-premises security identifier (SID) for the group synchronized from on-premises to the cloud. Read-only. Returned by default. Supports $filter (eq including on null values).
    onPremisesSecurityIdentifier *string
    // true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
    onPremisesSyncEnabled *bool
    // The owners of the group who can be users or service principals. Limited to 100 owners. Nullable. If this property isn't specified when creating a Microsoft 365 group the calling user (admin or non-admin) is automatically assigned as the group owner. A non-admin user can't explicitly add themselves to this collection when they're creating the group. For more information, see the related known issue. For security groups, the admin user isn't automatically added to this collection. For more information, see the related known issue. Supports $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1); Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
    owners []DirectoryObjectable
    // The permissionGrants property
    permissionGrants []ResourceSpecificPermissionGrantable
    // The group's profile photo
    photo ProfilePhotoable
    // The profile photos owned by the group. Read-only. Nullable.
    photos []ProfilePhotoable
    // Entry-point to Planner resource that might exist for a Unified Group.
    planner PlannerGroupable
    // The preferred data location for the Microsoft 365 group. By default, the group inherits the group creator's preferred data location. To set this property, the calling app must be granted the Directory.ReadWrite.All permission and the user be assigned at least one of the following Microsoft Entra roles: User Account Administrator Directory Writer  Exchange Administrator  SharePoint Administrator  For more information about this property, see OneDrive Online Multi-Geo. Nullable. Returned by default.
    preferredDataLocation *string
    // The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example, en-US. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    preferredLanguage *string
    // Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. The any operator is required to filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter (eq, not, ge, le, startsWith, endsWith, /$count eq 0, /$count ne 0).
    proxyAddresses []string
    // The list of users or groups not allowed to create posts or calendar events in this group. Nullable
    rejectedSenders []DirectoryObjectable
    // Timestamp of when the group was last renewed. This value can't be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only.
    renewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, not, in).
    securityEnabled *bool
    // Security identifier of the group, used in Windows scenarios. Read-only. Returned by default.
    securityIdentifier *string
    // Errors published by a federated service describing a nontransient, service-specific error regarding the properties or link from a group object.  Supports $filter (eq, not, for isResolved and serviceInstance).
    serviceProvisioningErrors []ServiceProvisioningErrorable
    // Settings that can govern this group's behavior, like whether members can invite guests to the group. Nullable.
    settings []GroupSettingable
    // The list of SharePoint sites in this group. Access the default site with /sites/root.
    sites []Siteable
    // The team associated with this group.
    team Teamable
    // Specifies a Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange, or Red. Returned by default.
    theme *string
    // The group's conversation threads. Nullable.
    threads []ConversationThreadable
    // The groups that a group is a member of, either directly or through nested membership. Nullable.
    transitiveMemberOf []DirectoryObjectable
    // The direct and transitive members of a group. Nullable.
    transitiveMembers []DirectoryObjectable
    // The unique identifier that can be assigned to a group and used as an alternate key. Immutable. Read-only.
    uniqueName *string
    // Count of conversations that received new posts since the signed-in user last visited the group. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    unseenCount *int32
    // Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or HiddenMembership. HiddenMembership can be set only for Microsoft 365 groups when the groups are created. It can't be updated later. Other values of visibility can be updated after group creation. If visibility value isn't specified during group creation on Microsoft Graph, a security group is created as Private by default, and the Microsoft 365 group is Public. Groups assignable to roles are always Private. To learn more, see group visibility options. Returned by default. Nullable.
    visibility *string
}
// NewGroup instantiates a new Group and sets the default values.
func NewGroup()(*Group) {
    m := &Group{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.group"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroup(), nil
}
// GetAcceptedSenders gets the acceptedSenders property value. The list of users or groups allowed to create posts or calendar events in this group. If this list is nonempty, then only users or groups listed here are allowed to post.
// returns a []DirectoryObjectable when successful
func (m *Group) GetAcceptedSenders()([]DirectoryObjectable) {
    return m.acceptedSenders
}
// GetAllowExternalSenders gets the allowExternalSenders property value. Indicates if people external to the organization can send messages to the group. The default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// returns a *bool when successful
func (m *Group) GetAllowExternalSenders()(*bool) {
    return m.allowExternalSenders
}
// GetAppRoleAssignments gets the appRoleAssignments property value. Represents the app roles granted to a group for an application. Supports $expand.
// returns a []AppRoleAssignmentable when successful
func (m *Group) GetAppRoleAssignments()([]AppRoleAssignmentable) {
    return m.appRoleAssignments
}
// GetAssignedLabels gets the assignedLabels property value. The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on $select. This property can be updated only in delegated scenarios where the caller requires both the Microsoft Graph permission and a supported administrator role.
// returns a []AssignedLabelable when successful
func (m *Group) GetAssignedLabels()([]AssignedLabelable) {
    return m.assignedLabels
}
// GetAssignedLicenses gets the assignedLicenses property value. The licenses that are assigned to the group. Returned only on $select. Supports $filter (eq).Read-only.
// returns a []AssignedLicenseable when successful
func (m *Group) GetAssignedLicenses()([]AssignedLicenseable) {
    return m.assignedLicenses
}
// GetAutoSubscribeNewMembers gets the autoSubscribeNewMembers property value. Indicates if new members added to the group are autosubscribed to receive email notifications. You can set this property in a PATCH request for the group; don't set it in the initial POST request that creates the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// returns a *bool when successful
func (m *Group) GetAutoSubscribeNewMembers()(*bool) {
    return m.autoSubscribeNewMembers
}
// GetCalendar gets the calendar property value. The group's calendar. Read-only.
// returns a Calendarable when successful
func (m *Group) GetCalendar()(Calendarable) {
    return m.calendar
}
// GetCalendarView gets the calendarView property value. The calendar view for the calendar. Read-only.
// returns a []Eventable when successful
func (m *Group) GetCalendarView()([]Eventable) {
    return m.calendarView
}
// GetClassification gets the classification property value. Describes a classification for the group (such as low, medium, or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith).
// returns a *string when successful
func (m *Group) GetClassification()(*string) {
    return m.classification
}
// GetConversations gets the conversations property value. The group's conversations.
// returns a []Conversationable when successful
func (m *Group) GetConversations()([]Conversationable) {
    return m.conversations
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp of when the group was created. The value can't be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only.
// returns a *Time when successful
func (m *Group) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCreatedOnBehalfOf gets the createdOnBehalfOf property value. The user (or application) that created the group. NOTE: This property isn't set if the user is an administrator. Read-only.
// returns a DirectoryObjectable when successful
func (m *Group) GetCreatedOnBehalfOf()(DirectoryObjectable) {
    return m.createdOnBehalfOf
}
// GetDescription gets the description property value. An optional description for the group. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
// returns a *string when successful
func (m *Group) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name for the group. This property is required when a group is created and can't be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderby.
// returns a *string when successful
func (m *Group) GetDisplayName()(*string) {
    return m.displayName
}
// GetDrive gets the drive property value. The group's default drive. Read-only.
// returns a Driveable when successful
func (m *Group) GetDrive()(Driveable) {
    return m.drive
}
// GetDrives gets the drives property value. The group's drives. Read-only.
// returns a []Driveable when successful
func (m *Group) GetDrives()([]Driveable) {
    return m.drives
}
// GetEvents gets the events property value. The group's calendar events.
// returns a []Eventable when successful
func (m *Group) GetEvents()([]Eventable) {
    return m.events
}
// GetExpirationDateTime gets the expirationDateTime property value. Timestamp of when the group is set to expire. It's null for security groups, but for Microsoft 365 groups, it represents when the group is set to expire as defined in the groupLifecyclePolicy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only.
// returns a *Time when successful
func (m *Group) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the group. Read-only. Nullable.
// returns a []Extensionable when successful
func (m *Group) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Group) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["acceptedSenders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAcceptedSenders(res)
        }
        return nil
    }
    res["allowExternalSenders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowExternalSenders(val)
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
    res["assignedLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignedLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedLabelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AssignedLabelable)
                }
            }
            m.SetAssignedLabels(res)
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
    res["autoSubscribeNewMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoSubscribeNewMembers(val)
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
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val)
        }
        return nil
    }
    res["conversations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConversationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Conversationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Conversationable)
                }
            }
            m.SetConversations(res)
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
    res["createdOnBehalfOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedOnBehalfOf(val.(DirectoryObjectable))
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
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
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
    res["groupLifecyclePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupLifecyclePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupLifecyclePolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupLifecyclePolicyable)
                }
            }
            m.SetGroupLifecyclePolicies(res)
        }
        return nil
    }
    res["groupTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetGroupTypes(res)
        }
        return nil
    }
    res["hasMembersWithLicenseErrors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMembersWithLicenseErrors(val)
        }
        return nil
    }
    res["hideFromAddressLists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideFromAddressLists(val)
        }
        return nil
    }
    res["hideFromOutlookClients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideFromOutlookClients(val)
        }
        return nil
    }
    res["isArchived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsArchived(val)
        }
        return nil
    }
    res["isAssignableToRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssignableToRole(val)
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
    res["isSubscribedByMail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSubscribedByMail(val)
        }
        return nil
    }
    res["licenseProcessingState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLicenseProcessingStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseProcessingState(val.(LicenseProcessingStateable))
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
    res["mailEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailEnabled(val)
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
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetMembers(res)
        }
        return nil
    }
    res["membershipRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRule(val)
        }
        return nil
    }
    res["membershipRuleProcessingState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRuleProcessingState(val)
        }
        return nil
    }
    res["membersWithLicenseErrors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetMembersWithLicenseErrors(res)
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
    res["onPremisesNetBiosName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesNetBiosName(val)
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
        val, err := n.GetObjectValue(CreatePlannerGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanner(val.(PlannerGroupable))
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
    res["rejectedSenders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRejectedSenders(res)
        }
        return nil
    }
    res["renewedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenewedDateTime(val)
        }
        return nil
    }
    res["securityEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEnabled(val)
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
        val, err := n.GetCollectionOfObjectValues(CreateGroupSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupSettingable)
                }
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["sites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSites(res)
        }
        return nil
    }
    res["team"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeam(val.(Teamable))
        }
        return nil
    }
    res["theme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTheme(val)
        }
        return nil
    }
    res["threads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConversationThreadFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConversationThreadable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConversationThreadable)
                }
            }
            m.SetThreads(res)
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
    res["transitiveMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTransitiveMembers(res)
        }
        return nil
    }
    res["uniqueName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueName(val)
        }
        return nil
    }
    res["unseenCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnseenCount(val)
        }
        return nil
    }
    res["visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val)
        }
        return nil
    }
    return res
}
// GetGroupLifecyclePolicies gets the groupLifecyclePolicies property value. The collection of lifecycle policies for this group. Read-only. Nullable.
// returns a []GroupLifecyclePolicyable when successful
func (m *Group) GetGroupLifecyclePolicies()([]GroupLifecyclePolicyable) {
    return m.groupLifecyclePolicies
}
// GetGroupTypes gets the groupTypes property value. Specifies the group type and its membership. If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or a distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static. Returned by default. Supports $filter (eq, not).
// returns a []string when successful
func (m *Group) GetGroupTypes()([]string) {
    return m.groupTypes
}
// GetHasMembersWithLicenseErrors gets the hasMembersWithLicenseErrors property value. Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true). See an example. Supports $filter (eq).
// returns a *bool when successful
func (m *Group) GetHasMembersWithLicenseErrors()(*bool) {
    return m.hasMembersWithLicenseErrors
}
// GetHideFromAddressLists gets the hideFromAddressLists property value. True if the group isn't displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false. The default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// returns a *bool when successful
func (m *Group) GetHideFromAddressLists()(*bool) {
    return m.hideFromAddressLists
}
// GetHideFromOutlookClients gets the hideFromOutlookClients property value. True if the group isn't displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false. The default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// returns a *bool when successful
func (m *Group) GetHideFromOutlookClients()(*bool) {
    return m.hideFromOutlookClients
}
// GetIsArchived gets the isArchived property value. When a group is associated with a team, this property determines whether the team is in read-only mode.To read this property, use the /group/{groupId}/team endpoint or the Get team API. To update this property, use the archiveTeam and unarchiveTeam APIs.
// returns a *bool when successful
func (m *Group) GetIsArchived()(*bool) {
    return m.isArchived
}
// GetIsAssignableToRole gets the isAssignableToRole property value. Indicates whether this group can be assigned to a Microsoft Entra role. Optional. This property can only be set while creating the group and is immutable. If set to true, the securityEnabled property must also be set to true, visibility must be Hidden, and the group can't be a dynamic group (that is, groupTypes can't contain DynamicMembership). Only callers with at least the Privileged Role Administrator role can set this property. The caller must also be assigned the RoleManagement.ReadWrite.Directory permission to set this property or update the membership of such groups. For more, see Using a group to manage Microsoft Entra role assignmentsUsing this feature requires a Microsoft Entra ID P1 license. Returned by default. Supports $filter (eq, ne, not).
// returns a *bool when successful
func (m *Group) GetIsAssignableToRole()(*bool) {
    return m.isAssignableToRole
}
// GetIsManagementRestricted gets the isManagementRestricted property value. The isManagementRestricted property
// returns a *bool when successful
func (m *Group) GetIsManagementRestricted()(*bool) {
    return m.isManagementRestricted
}
// GetIsSubscribedByMail gets the isSubscribedByMail property value. Indicates whether the signed-in user is subscribed to receive email conversations. The default value is true. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// returns a *bool when successful
func (m *Group) GetIsSubscribedByMail()(*bool) {
    return m.isSubscribedByMail
}
// GetLicenseProcessingState gets the licenseProcessingState property value. Indicates the status of the group license assignment to all group members. The default value is false. Read-only. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select. Read-only.
// returns a LicenseProcessingStateable when successful
func (m *Group) GetLicenseProcessingState()(LicenseProcessingStateable) {
    return m.licenseProcessingState
}
// GetMail gets the mail property value. The SMTP address for the group, for example, 'serviceadmins@contoso.com'. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *Group) GetMail()(*string) {
    return m.mail
}
// GetMailEnabled gets the mailEnabled property value. Specifies whether the group is mail-enabled. Required. Returned by default. Supports $filter (eq, ne, not).
// returns a *bool when successful
func (m *Group) GetMailEnabled()(*bool) {
    return m.mailEnabled
}
// GetMailNickname gets the mailNickname property value. The mail alias for the group, unique for Microsoft 365 groups in the organization. Maximum length is 64 characters. This property can contain only characters in the ASCII character set 0 - 127 except the following characters: @ () / [] ' ; : <> , SPACE. Required. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *Group) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetMemberOf gets the memberOf property value. Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand.
// returns a []DirectoryObjectable when successful
func (m *Group) GetMemberOf()([]DirectoryObjectable) {
    return m.memberOf
}
// GetMembers gets the members property value. The members of this group, who can be users, devices, other groups, or service principals. Supports the List members, Add member, and Remove member operations. Nullable. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=members($select=id,userPrincipalName,displayName).
// returns a []DirectoryObjectable when successful
func (m *Group) GetMembers()([]DirectoryObjectable) {
    return m.members
}
// GetMembershipRule gets the membershipRule property value. The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith).
// returns a *string when successful
func (m *Group) GetMembershipRule()(*string) {
    return m.membershipRule
}
// GetMembershipRuleProcessingState gets the membershipRuleProcessingState property value. Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by default. Supports $filter (eq, ne, not, in).
// returns a *string when successful
func (m *Group) GetMembershipRuleProcessingState()(*string) {
    return m.membershipRuleProcessingState
}
// GetMembersWithLicenseErrors gets the membersWithLicenseErrors property value. A list of group members with license errors from this group-based license assignment. Read-only.
// returns a []DirectoryObjectable when successful
func (m *Group) GetMembersWithLicenseErrors()([]DirectoryObjectable) {
    return m.membersWithLicenseErrors
}
// GetOnenote gets the onenote property value. The onenote property
// returns a Onenoteable when successful
func (m *Group) GetOnenote()(Onenoteable) {
    return m.onenote
}
// GetOnPremisesDomainName gets the onPremisesDomainName property value. Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect.Returned by default. Read-only.
// returns a *string when successful
func (m *Group) GetOnPremisesDomainName()(*string) {
    return m.onPremisesDomainName
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. Indicates the last time at which the group was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in).
// returns a *Time when successful
func (m *Group) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.onPremisesLastSyncDateTime
}
// GetOnPremisesNetBiosName gets the onPremisesNetBiosName property value. Contains the on-premises netBios name synchronized from the on-premises directory. The property is only populated for customers synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect.Returned by default. Read-only.
// returns a *string when successful
func (m *Group) GetOnPremisesNetBiosName()(*string) {
    return m.onPremisesNetBiosName
}
// GetOnPremisesProvisioningErrors gets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned by default. Supports $filter (eq, not).
// returns a []OnPremisesProvisioningErrorable when successful
func (m *Group) GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningErrorable) {
    return m.onPremisesProvisioningErrors
}
// GetOnPremisesSamAccountName gets the onPremisesSamAccountName property value. Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect.Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith). Read-only.
// returns a *string when successful
func (m *Group) GetOnPremisesSamAccountName()(*string) {
    return m.onPremisesSamAccountName
}
// GetOnPremisesSecurityIdentifier gets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the group synchronized from on-premises to the cloud. Read-only. Returned by default. Supports $filter (eq including on null values).
// returns a *string when successful
func (m *Group) GetOnPremisesSecurityIdentifier()(*string) {
    return m.onPremisesSecurityIdentifier
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
// returns a *bool when successful
func (m *Group) GetOnPremisesSyncEnabled()(*bool) {
    return m.onPremisesSyncEnabled
}
// GetOwners gets the owners property value. The owners of the group who can be users or service principals. Limited to 100 owners. Nullable. If this property isn't specified when creating a Microsoft 365 group the calling user (admin or non-admin) is automatically assigned as the group owner. A non-admin user can't explicitly add themselves to this collection when they're creating the group. For more information, see the related known issue. For security groups, the admin user isn't automatically added to this collection. For more information, see the related known issue. Supports $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1); Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
// returns a []DirectoryObjectable when successful
func (m *Group) GetOwners()([]DirectoryObjectable) {
    return m.owners
}
// GetPermissionGrants gets the permissionGrants property value. The permissionGrants property
// returns a []ResourceSpecificPermissionGrantable when successful
func (m *Group) GetPermissionGrants()([]ResourceSpecificPermissionGrantable) {
    return m.permissionGrants
}
// GetPhoto gets the photo property value. The group's profile photo
// returns a ProfilePhotoable when successful
func (m *Group) GetPhoto()(ProfilePhotoable) {
    return m.photo
}
// GetPhotos gets the photos property value. The profile photos owned by the group. Read-only. Nullable.
// returns a []ProfilePhotoable when successful
func (m *Group) GetPhotos()([]ProfilePhotoable) {
    return m.photos
}
// GetPlanner gets the planner property value. Entry-point to Planner resource that might exist for a Unified Group.
// returns a PlannerGroupable when successful
func (m *Group) GetPlanner()(PlannerGroupable) {
    return m.planner
}
// GetPreferredDataLocation gets the preferredDataLocation property value. The preferred data location for the Microsoft 365 group. By default, the group inherits the group creator's preferred data location. To set this property, the calling app must be granted the Directory.ReadWrite.All permission and the user be assigned at least one of the following Microsoft Entra roles: User Account Administrator Directory Writer  Exchange Administrator  SharePoint Administrator  For more information about this property, see OneDrive Online Multi-Geo. Nullable. Returned by default.
// returns a *string when successful
func (m *Group) GetPreferredDataLocation()(*string) {
    return m.preferredDataLocation
}
// GetPreferredLanguage gets the preferredLanguage property value. The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example, en-US. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
// returns a *string when successful
func (m *Group) GetPreferredLanguage()(*string) {
    return m.preferredLanguage
}
// GetProxyAddresses gets the proxyAddresses property value. Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. The any operator is required to filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter (eq, not, ge, le, startsWith, endsWith, /$count eq 0, /$count ne 0).
// returns a []string when successful
func (m *Group) GetProxyAddresses()([]string) {
    return m.proxyAddresses
}
// GetRejectedSenders gets the rejectedSenders property value. The list of users or groups not allowed to create posts or calendar events in this group. Nullable
// returns a []DirectoryObjectable when successful
func (m *Group) GetRejectedSenders()([]DirectoryObjectable) {
    return m.rejectedSenders
}
// GetRenewedDateTime gets the renewedDateTime property value. Timestamp of when the group was last renewed. This value can't be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only.
// returns a *Time when successful
func (m *Group) GetRenewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.renewedDateTime
}
// GetSecurityEnabled gets the securityEnabled property value. Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, not, in).
// returns a *bool when successful
func (m *Group) GetSecurityEnabled()(*bool) {
    return m.securityEnabled
}
// GetSecurityIdentifier gets the securityIdentifier property value. Security identifier of the group, used in Windows scenarios. Read-only. Returned by default.
// returns a *string when successful
func (m *Group) GetSecurityIdentifier()(*string) {
    return m.securityIdentifier
}
// GetServiceProvisioningErrors gets the serviceProvisioningErrors property value. Errors published by a federated service describing a nontransient, service-specific error regarding the properties or link from a group object.  Supports $filter (eq, not, for isResolved and serviceInstance).
// returns a []ServiceProvisioningErrorable when successful
func (m *Group) GetServiceProvisioningErrors()([]ServiceProvisioningErrorable) {
    return m.serviceProvisioningErrors
}
// GetSettings gets the settings property value. Settings that can govern this group's behavior, like whether members can invite guests to the group. Nullable.
// returns a []GroupSettingable when successful
func (m *Group) GetSettings()([]GroupSettingable) {
    return m.settings
}
// GetSites gets the sites property value. The list of SharePoint sites in this group. Access the default site with /sites/root.
// returns a []Siteable when successful
func (m *Group) GetSites()([]Siteable) {
    return m.sites
}
// GetTeam gets the team property value. The team associated with this group.
// returns a Teamable when successful
func (m *Group) GetTeam()(Teamable) {
    return m.team
}
// GetTheme gets the theme property value. Specifies a Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange, or Red. Returned by default.
// returns a *string when successful
func (m *Group) GetTheme()(*string) {
    return m.theme
}
// GetThreads gets the threads property value. The group's conversation threads. Nullable.
// returns a []ConversationThreadable when successful
func (m *Group) GetThreads()([]ConversationThreadable) {
    return m.threads
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. The groups that a group is a member of, either directly or through nested membership. Nullable.
// returns a []DirectoryObjectable when successful
func (m *Group) GetTransitiveMemberOf()([]DirectoryObjectable) {
    return m.transitiveMemberOf
}
// GetTransitiveMembers gets the transitiveMembers property value. The direct and transitive members of a group. Nullable.
// returns a []DirectoryObjectable when successful
func (m *Group) GetTransitiveMembers()([]DirectoryObjectable) {
    return m.transitiveMembers
}
// GetUniqueName gets the uniqueName property value. The unique identifier that can be assigned to a group and used as an alternate key. Immutable. Read-only.
// returns a *string when successful
func (m *Group) GetUniqueName()(*string) {
    return m.uniqueName
}
// GetUnseenCount gets the unseenCount property value. Count of conversations that received new posts since the signed-in user last visited the group. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// returns a *int32 when successful
func (m *Group) GetUnseenCount()(*int32) {
    return m.unseenCount
}
// GetVisibility gets the visibility property value. Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or HiddenMembership. HiddenMembership can be set only for Microsoft 365 groups when the groups are created. It can't be updated later. Other values of visibility can be updated after group creation. If visibility value isn't specified during group creation on Microsoft Graph, a security group is created as Private by default, and the Microsoft 365 group is Public. Groups assignable to roles are always Private. To learn more, see group visibility options. Returned by default. Nullable.
// returns a *string when successful
func (m *Group) GetVisibility()(*string) {
    return m.visibility
}
// Serialize serializes information the current object
func (m *Group) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAcceptedSenders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAcceptedSenders()))
        for i, v := range m.GetAcceptedSenders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("acceptedSenders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowExternalSenders", m.GetAllowExternalSenders())
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
    if m.GetAssignedLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedLabels()))
        for i, v := range m.GetAssignedLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignedLabels", cast)
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
    {
        err = writer.WriteBoolValue("autoSubscribeNewMembers", m.GetAutoSubscribeNewMembers())
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
    {
        err = writer.WriteStringValue("classification", m.GetClassification())
        if err != nil {
            return err
        }
    }
    if m.GetConversations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConversations()))
        for i, v := range m.GetConversations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("conversations", cast)
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
    {
        err = writer.WriteObjectValue("createdOnBehalfOf", m.GetCreatedOnBehalfOf())
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
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
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
    if m.GetGroupLifecyclePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupLifecyclePolicies()))
        for i, v := range m.GetGroupLifecyclePolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groupLifecyclePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupTypes() != nil {
        err = writer.WriteCollectionOfStringValues("groupTypes", m.GetGroupTypes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasMembersWithLicenseErrors", m.GetHasMembersWithLicenseErrors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hideFromAddressLists", m.GetHideFromAddressLists())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hideFromOutlookClients", m.GetHideFromOutlookClients())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isArchived", m.GetIsArchived())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAssignableToRole", m.GetIsAssignableToRole())
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
        err = writer.WriteBoolValue("isSubscribedByMail", m.GetIsSubscribedByMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("licenseProcessingState", m.GetLicenseProcessingState())
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
        err = writer.WriteBoolValue("mailEnabled", m.GetMailEnabled())
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
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("membershipRule", m.GetMembershipRule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("membershipRuleProcessingState", m.GetMembershipRuleProcessingState())
        if err != nil {
            return err
        }
    }
    if m.GetMembersWithLicenseErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembersWithLicenseErrors()))
        for i, v := range m.GetMembersWithLicenseErrors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("membersWithLicenseErrors", cast)
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
    {
        err = writer.WriteStringValue("onPremisesDomainName", m.GetOnPremisesDomainName())
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
    {
        err = writer.WriteStringValue("onPremisesNetBiosName", m.GetOnPremisesNetBiosName())
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
    if m.GetProxyAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("proxyAddresses", m.GetProxyAddresses())
        if err != nil {
            return err
        }
    }
    if m.GetRejectedSenders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRejectedSenders()))
        for i, v := range m.GetRejectedSenders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("rejectedSenders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("renewedDateTime", m.GetRenewedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityEnabled", m.GetSecurityEnabled())
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
    if m.GetSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSites()))
        for i, v := range m.GetSites() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sites", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("team", m.GetTeam())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("theme", m.GetTheme())
        if err != nil {
            return err
        }
    }
    if m.GetThreads() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetThreads()))
        for i, v := range m.GetThreads() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("threads", cast)
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
    if m.GetTransitiveMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransitiveMembers()))
        for i, v := range m.GetTransitiveMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uniqueName", m.GetUniqueName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unseenCount", m.GetUnseenCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("visibility", m.GetVisibility())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedSenders sets the acceptedSenders property value. The list of users or groups allowed to create posts or calendar events in this group. If this list is nonempty, then only users or groups listed here are allowed to post.
func (m *Group) SetAcceptedSenders(value []DirectoryObjectable)() {
    m.acceptedSenders = value
}
// SetAllowExternalSenders sets the allowExternalSenders property value. Indicates if people external to the organization can send messages to the group. The default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetAllowExternalSenders(value *bool)() {
    m.allowExternalSenders = value
}
// SetAppRoleAssignments sets the appRoleAssignments property value. Represents the app roles granted to a group for an application. Supports $expand.
func (m *Group) SetAppRoleAssignments(value []AppRoleAssignmentable)() {
    m.appRoleAssignments = value
}
// SetAssignedLabels sets the assignedLabels property value. The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on $select. This property can be updated only in delegated scenarios where the caller requires both the Microsoft Graph permission and a supported administrator role.
func (m *Group) SetAssignedLabels(value []AssignedLabelable)() {
    m.assignedLabels = value
}
// SetAssignedLicenses sets the assignedLicenses property value. The licenses that are assigned to the group. Returned only on $select. Supports $filter (eq).Read-only.
func (m *Group) SetAssignedLicenses(value []AssignedLicenseable)() {
    m.assignedLicenses = value
}
// SetAutoSubscribeNewMembers sets the autoSubscribeNewMembers property value. Indicates if new members added to the group are autosubscribed to receive email notifications. You can set this property in a PATCH request for the group; don't set it in the initial POST request that creates the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetAutoSubscribeNewMembers(value *bool)() {
    m.autoSubscribeNewMembers = value
}
// SetCalendar sets the calendar property value. The group's calendar. Read-only.
func (m *Group) SetCalendar(value Calendarable)() {
    m.calendar = value
}
// SetCalendarView sets the calendarView property value. The calendar view for the calendar. Read-only.
func (m *Group) SetCalendarView(value []Eventable)() {
    m.calendarView = value
}
// SetClassification sets the classification property value. Describes a classification for the group (such as low, medium, or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith).
func (m *Group) SetClassification(value *string)() {
    m.classification = value
}
// SetConversations sets the conversations property value. The group's conversations.
func (m *Group) SetConversations(value []Conversationable)() {
    m.conversations = value
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp of when the group was created. The value can't be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only.
func (m *Group) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreatedOnBehalfOf sets the createdOnBehalfOf property value. The user (or application) that created the group. NOTE: This property isn't set if the user is an administrator. Read-only.
func (m *Group) SetCreatedOnBehalfOf(value DirectoryObjectable)() {
    m.createdOnBehalfOf = value
}
// SetDescription sets the description property value. An optional description for the group. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *Group) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for the group. This property is required when a group is created and can't be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderby.
func (m *Group) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDrive sets the drive property value. The group's default drive. Read-only.
func (m *Group) SetDrive(value Driveable)() {
    m.drive = value
}
// SetDrives sets the drives property value. The group's drives. Read-only.
func (m *Group) SetDrives(value []Driveable)() {
    m.drives = value
}
// SetEvents sets the events property value. The group's calendar events.
func (m *Group) SetEvents(value []Eventable)() {
    m.events = value
}
// SetExpirationDateTime sets the expirationDateTime property value. Timestamp of when the group is set to expire. It's null for security groups, but for Microsoft 365 groups, it represents when the group is set to expire as defined in the groupLifecyclePolicy. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only.
func (m *Group) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the group. Read-only. Nullable.
func (m *Group) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetGroupLifecyclePolicies sets the groupLifecyclePolicies property value. The collection of lifecycle policies for this group. Read-only. Nullable.
func (m *Group) SetGroupLifecyclePolicies(value []GroupLifecyclePolicyable)() {
    m.groupLifecyclePolicies = value
}
// SetGroupTypes sets the groupTypes property value. Specifies the group type and its membership. If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or a distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static. Returned by default. Supports $filter (eq, not).
func (m *Group) SetGroupTypes(value []string)() {
    m.groupTypes = value
}
// SetHasMembersWithLicenseErrors sets the hasMembersWithLicenseErrors property value. Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true). See an example. Supports $filter (eq).
func (m *Group) SetHasMembersWithLicenseErrors(value *bool)() {
    m.hasMembersWithLicenseErrors = value
}
// SetHideFromAddressLists sets the hideFromAddressLists property value. True if the group isn't displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false. The default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetHideFromAddressLists(value *bool)() {
    m.hideFromAddressLists = value
}
// SetHideFromOutlookClients sets the hideFromOutlookClients property value. True if the group isn't displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false. The default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetHideFromOutlookClients(value *bool)() {
    m.hideFromOutlookClients = value
}
// SetIsArchived sets the isArchived property value. When a group is associated with a team, this property determines whether the team is in read-only mode.To read this property, use the /group/{groupId}/team endpoint or the Get team API. To update this property, use the archiveTeam and unarchiveTeam APIs.
func (m *Group) SetIsArchived(value *bool)() {
    m.isArchived = value
}
// SetIsAssignableToRole sets the isAssignableToRole property value. Indicates whether this group can be assigned to a Microsoft Entra role. Optional. This property can only be set while creating the group and is immutable. If set to true, the securityEnabled property must also be set to true, visibility must be Hidden, and the group can't be a dynamic group (that is, groupTypes can't contain DynamicMembership). Only callers with at least the Privileged Role Administrator role can set this property. The caller must also be assigned the RoleManagement.ReadWrite.Directory permission to set this property or update the membership of such groups. For more, see Using a group to manage Microsoft Entra role assignmentsUsing this feature requires a Microsoft Entra ID P1 license. Returned by default. Supports $filter (eq, ne, not).
func (m *Group) SetIsAssignableToRole(value *bool)() {
    m.isAssignableToRole = value
}
// SetIsManagementRestricted sets the isManagementRestricted property value. The isManagementRestricted property
func (m *Group) SetIsManagementRestricted(value *bool)() {
    m.isManagementRestricted = value
}
// SetIsSubscribedByMail sets the isSubscribedByMail property value. Indicates whether the signed-in user is subscribed to receive email conversations. The default value is true. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetIsSubscribedByMail(value *bool)() {
    m.isSubscribedByMail = value
}
// SetLicenseProcessingState sets the licenseProcessingState property value. Indicates the status of the group license assignment to all group members. The default value is false. Read-only. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select. Read-only.
func (m *Group) SetLicenseProcessingState(value LicenseProcessingStateable)() {
    m.licenseProcessingState = value
}
// SetMail sets the mail property value. The SMTP address for the group, for example, 'serviceadmins@contoso.com'. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *Group) SetMail(value *string)() {
    m.mail = value
}
// SetMailEnabled sets the mailEnabled property value. Specifies whether the group is mail-enabled. Required. Returned by default. Supports $filter (eq, ne, not).
func (m *Group) SetMailEnabled(value *bool)() {
    m.mailEnabled = value
}
// SetMailNickname sets the mailNickname property value. The mail alias for the group, unique for Microsoft 365 groups in the organization. Maximum length is 64 characters. This property can contain only characters in the ASCII character set 0 - 127 except the following characters: @ () / [] ' ; : <> , SPACE. Required. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *Group) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetMemberOf sets the memberOf property value. Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand.
func (m *Group) SetMemberOf(value []DirectoryObjectable)() {
    m.memberOf = value
}
// SetMembers sets the members property value. The members of this group, who can be users, devices, other groups, or service principals. Supports the List members, Add member, and Remove member operations. Nullable. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=members($select=id,userPrincipalName,displayName).
func (m *Group) SetMembers(value []DirectoryObjectable)() {
    m.members = value
}
// SetMembershipRule sets the membershipRule property value. The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith).
func (m *Group) SetMembershipRule(value *string)() {
    m.membershipRule = value
}
// SetMembershipRuleProcessingState sets the membershipRuleProcessingState property value. Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by default. Supports $filter (eq, ne, not, in).
func (m *Group) SetMembershipRuleProcessingState(value *string)() {
    m.membershipRuleProcessingState = value
}
// SetMembersWithLicenseErrors sets the membersWithLicenseErrors property value. A list of group members with license errors from this group-based license assignment. Read-only.
func (m *Group) SetMembersWithLicenseErrors(value []DirectoryObjectable)() {
    m.membersWithLicenseErrors = value
}
// SetOnenote sets the onenote property value. The onenote property
func (m *Group) SetOnenote(value Onenoteable)() {
    m.onenote = value
}
// SetOnPremisesDomainName sets the onPremisesDomainName property value. Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect.Returned by default. Read-only.
func (m *Group) SetOnPremisesDomainName(value *string)() {
    m.onPremisesDomainName = value
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. Indicates the last time at which the group was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in).
func (m *Group) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// SetOnPremisesNetBiosName sets the onPremisesNetBiosName property value. Contains the on-premises netBios name synchronized from the on-premises directory. The property is only populated for customers synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect.Returned by default. Read-only.
func (m *Group) SetOnPremisesNetBiosName(value *string)() {
    m.onPremisesNetBiosName = value
}
// SetOnPremisesProvisioningErrors sets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned by default. Supports $filter (eq, not).
func (m *Group) SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningErrorable)() {
    m.onPremisesProvisioningErrors = value
}
// SetOnPremisesSamAccountName sets the onPremisesSamAccountName property value. Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers synchronizing their on-premises directory to Microsoft Entra ID via Microsoft Entra Connect.Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith). Read-only.
func (m *Group) SetOnPremisesSamAccountName(value *string)() {
    m.onPremisesSamAccountName = value
}
// SetOnPremisesSecurityIdentifier sets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the group synchronized from on-premises to the cloud. Read-only. Returned by default. Supports $filter (eq including on null values).
func (m *Group) SetOnPremisesSecurityIdentifier(value *string)() {
    m.onPremisesSecurityIdentifier = value
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *Group) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// SetOwners sets the owners property value. The owners of the group who can be users or service principals. Limited to 100 owners. Nullable. If this property isn't specified when creating a Microsoft 365 group the calling user (admin or non-admin) is automatically assigned as the group owner. A non-admin user can't explicitly add themselves to this collection when they're creating the group. For more information, see the related known issue. For security groups, the admin user isn't automatically added to this collection. For more information, see the related known issue. Supports $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1); Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
func (m *Group) SetOwners(value []DirectoryObjectable)() {
    m.owners = value
}
// SetPermissionGrants sets the permissionGrants property value. The permissionGrants property
func (m *Group) SetPermissionGrants(value []ResourceSpecificPermissionGrantable)() {
    m.permissionGrants = value
}
// SetPhoto sets the photo property value. The group's profile photo
func (m *Group) SetPhoto(value ProfilePhotoable)() {
    m.photo = value
}
// SetPhotos sets the photos property value. The profile photos owned by the group. Read-only. Nullable.
func (m *Group) SetPhotos(value []ProfilePhotoable)() {
    m.photos = value
}
// SetPlanner sets the planner property value. Entry-point to Planner resource that might exist for a Unified Group.
func (m *Group) SetPlanner(value PlannerGroupable)() {
    m.planner = value
}
// SetPreferredDataLocation sets the preferredDataLocation property value. The preferred data location for the Microsoft 365 group. By default, the group inherits the group creator's preferred data location. To set this property, the calling app must be granted the Directory.ReadWrite.All permission and the user be assigned at least one of the following Microsoft Entra roles: User Account Administrator Directory Writer  Exchange Administrator  SharePoint Administrator  For more information about this property, see OneDrive Online Multi-Geo. Nullable. Returned by default.
func (m *Group) SetPreferredDataLocation(value *string)() {
    m.preferredDataLocation = value
}
// SetPreferredLanguage sets the preferredLanguage property value. The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example, en-US. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *Group) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
// SetProxyAddresses sets the proxyAddresses property value. Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. The any operator is required to filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter (eq, not, ge, le, startsWith, endsWith, /$count eq 0, /$count ne 0).
func (m *Group) SetProxyAddresses(value []string)() {
    m.proxyAddresses = value
}
// SetRejectedSenders sets the rejectedSenders property value. The list of users or groups not allowed to create posts or calendar events in this group. Nullable
func (m *Group) SetRejectedSenders(value []DirectoryObjectable)() {
    m.rejectedSenders = value
}
// SetRenewedDateTime sets the renewedDateTime property value. Timestamp of when the group was last renewed. This value can't be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on January 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only.
func (m *Group) SetRenewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.renewedDateTime = value
}
// SetSecurityEnabled sets the securityEnabled property value. Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, not, in).
func (m *Group) SetSecurityEnabled(value *bool)() {
    m.securityEnabled = value
}
// SetSecurityIdentifier sets the securityIdentifier property value. Security identifier of the group, used in Windows scenarios. Read-only. Returned by default.
func (m *Group) SetSecurityIdentifier(value *string)() {
    m.securityIdentifier = value
}
// SetServiceProvisioningErrors sets the serviceProvisioningErrors property value. Errors published by a federated service describing a nontransient, service-specific error regarding the properties or link from a group object.  Supports $filter (eq, not, for isResolved and serviceInstance).
func (m *Group) SetServiceProvisioningErrors(value []ServiceProvisioningErrorable)() {
    m.serviceProvisioningErrors = value
}
// SetSettings sets the settings property value. Settings that can govern this group's behavior, like whether members can invite guests to the group. Nullable.
func (m *Group) SetSettings(value []GroupSettingable)() {
    m.settings = value
}
// SetSites sets the sites property value. The list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *Group) SetSites(value []Siteable)() {
    m.sites = value
}
// SetTeam sets the team property value. The team associated with this group.
func (m *Group) SetTeam(value Teamable)() {
    m.team = value
}
// SetTheme sets the theme property value. Specifies a Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange, or Red. Returned by default.
func (m *Group) SetTheme(value *string)() {
    m.theme = value
}
// SetThreads sets the threads property value. The group's conversation threads. Nullable.
func (m *Group) SetThreads(value []ConversationThreadable)() {
    m.threads = value
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. The groups that a group is a member of, either directly or through nested membership. Nullable.
func (m *Group) SetTransitiveMemberOf(value []DirectoryObjectable)() {
    m.transitiveMemberOf = value
}
// SetTransitiveMembers sets the transitiveMembers property value. The direct and transitive members of a group. Nullable.
func (m *Group) SetTransitiveMembers(value []DirectoryObjectable)() {
    m.transitiveMembers = value
}
// SetUniqueName sets the uniqueName property value. The unique identifier that can be assigned to a group and used as an alternate key. Immutable. Read-only.
func (m *Group) SetUniqueName(value *string)() {
    m.uniqueName = value
}
// SetUnseenCount sets the unseenCount property value. Count of conversations that received new posts since the signed-in user last visited the group. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetUnseenCount(value *int32)() {
    m.unseenCount = value
}
// SetVisibility sets the visibility property value. Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or HiddenMembership. HiddenMembership can be set only for Microsoft 365 groups when the groups are created. It can't be updated later. Other values of visibility can be updated after group creation. If visibility value isn't specified during group creation on Microsoft Graph, a security group is created as Private by default, and the Microsoft 365 group is Public. Groups assignable to roles are always Private. To learn more, see group visibility options. Returned by default. Nullable.
func (m *Group) SetVisibility(value *string)() {
    m.visibility = value
}
type Groupable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptedSenders()([]DirectoryObjectable)
    GetAllowExternalSenders()(*bool)
    GetAppRoleAssignments()([]AppRoleAssignmentable)
    GetAssignedLabels()([]AssignedLabelable)
    GetAssignedLicenses()([]AssignedLicenseable)
    GetAutoSubscribeNewMembers()(*bool)
    GetCalendar()(Calendarable)
    GetCalendarView()([]Eventable)
    GetClassification()(*string)
    GetConversations()([]Conversationable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedOnBehalfOf()(DirectoryObjectable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDrive()(Driveable)
    GetDrives()([]Driveable)
    GetEvents()([]Eventable)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExtensions()([]Extensionable)
    GetGroupLifecyclePolicies()([]GroupLifecyclePolicyable)
    GetGroupTypes()([]string)
    GetHasMembersWithLicenseErrors()(*bool)
    GetHideFromAddressLists()(*bool)
    GetHideFromOutlookClients()(*bool)
    GetIsArchived()(*bool)
    GetIsAssignableToRole()(*bool)
    GetIsManagementRestricted()(*bool)
    GetIsSubscribedByMail()(*bool)
    GetLicenseProcessingState()(LicenseProcessingStateable)
    GetMail()(*string)
    GetMailEnabled()(*bool)
    GetMailNickname()(*string)
    GetMemberOf()([]DirectoryObjectable)
    GetMembers()([]DirectoryObjectable)
    GetMembershipRule()(*string)
    GetMembershipRuleProcessingState()(*string)
    GetMembersWithLicenseErrors()([]DirectoryObjectable)
    GetOnenote()(Onenoteable)
    GetOnPremisesDomainName()(*string)
    GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesNetBiosName()(*string)
    GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningErrorable)
    GetOnPremisesSamAccountName()(*string)
    GetOnPremisesSecurityIdentifier()(*string)
    GetOnPremisesSyncEnabled()(*bool)
    GetOwners()([]DirectoryObjectable)
    GetPermissionGrants()([]ResourceSpecificPermissionGrantable)
    GetPhoto()(ProfilePhotoable)
    GetPhotos()([]ProfilePhotoable)
    GetPlanner()(PlannerGroupable)
    GetPreferredDataLocation()(*string)
    GetPreferredLanguage()(*string)
    GetProxyAddresses()([]string)
    GetRejectedSenders()([]DirectoryObjectable)
    GetRenewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSecurityEnabled()(*bool)
    GetSecurityIdentifier()(*string)
    GetServiceProvisioningErrors()([]ServiceProvisioningErrorable)
    GetSettings()([]GroupSettingable)
    GetSites()([]Siteable)
    GetTeam()(Teamable)
    GetTheme()(*string)
    GetThreads()([]ConversationThreadable)
    GetTransitiveMemberOf()([]DirectoryObjectable)
    GetTransitiveMembers()([]DirectoryObjectable)
    GetUniqueName()(*string)
    GetUnseenCount()(*int32)
    GetVisibility()(*string)
    SetAcceptedSenders(value []DirectoryObjectable)()
    SetAllowExternalSenders(value *bool)()
    SetAppRoleAssignments(value []AppRoleAssignmentable)()
    SetAssignedLabels(value []AssignedLabelable)()
    SetAssignedLicenses(value []AssignedLicenseable)()
    SetAutoSubscribeNewMembers(value *bool)()
    SetCalendar(value Calendarable)()
    SetCalendarView(value []Eventable)()
    SetClassification(value *string)()
    SetConversations(value []Conversationable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedOnBehalfOf(value DirectoryObjectable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDrive(value Driveable)()
    SetDrives(value []Driveable)()
    SetEvents(value []Eventable)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExtensions(value []Extensionable)()
    SetGroupLifecyclePolicies(value []GroupLifecyclePolicyable)()
    SetGroupTypes(value []string)()
    SetHasMembersWithLicenseErrors(value *bool)()
    SetHideFromAddressLists(value *bool)()
    SetHideFromOutlookClients(value *bool)()
    SetIsArchived(value *bool)()
    SetIsAssignableToRole(value *bool)()
    SetIsManagementRestricted(value *bool)()
    SetIsSubscribedByMail(value *bool)()
    SetLicenseProcessingState(value LicenseProcessingStateable)()
    SetMail(value *string)()
    SetMailEnabled(value *bool)()
    SetMailNickname(value *string)()
    SetMemberOf(value []DirectoryObjectable)()
    SetMembers(value []DirectoryObjectable)()
    SetMembershipRule(value *string)()
    SetMembershipRuleProcessingState(value *string)()
    SetMembersWithLicenseErrors(value []DirectoryObjectable)()
    SetOnenote(value Onenoteable)()
    SetOnPremisesDomainName(value *string)()
    SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesNetBiosName(value *string)()
    SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningErrorable)()
    SetOnPremisesSamAccountName(value *string)()
    SetOnPremisesSecurityIdentifier(value *string)()
    SetOnPremisesSyncEnabled(value *bool)()
    SetOwners(value []DirectoryObjectable)()
    SetPermissionGrants(value []ResourceSpecificPermissionGrantable)()
    SetPhoto(value ProfilePhotoable)()
    SetPhotos(value []ProfilePhotoable)()
    SetPlanner(value PlannerGroupable)()
    SetPreferredDataLocation(value *string)()
    SetPreferredLanguage(value *string)()
    SetProxyAddresses(value []string)()
    SetRejectedSenders(value []DirectoryObjectable)()
    SetRenewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSecurityEnabled(value *bool)()
    SetSecurityIdentifier(value *string)()
    SetServiceProvisioningErrors(value []ServiceProvisioningErrorable)()
    SetSettings(value []GroupSettingable)()
    SetSites(value []Siteable)()
    SetTeam(value Teamable)()
    SetTheme(value *string)()
    SetThreads(value []ConversationThreadable)()
    SetTransitiveMemberOf(value []DirectoryObjectable)()
    SetTransitiveMembers(value []DirectoryObjectable)()
    SetUniqueName(value *string)()
    SetUnseenCount(value *int32)()
    SetVisibility(value *string)()
}
