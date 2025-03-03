package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrgContact struct {
    DirectoryObject
    // Postal addresses for this organizational contact. For now a contact can only have one physical address.
    addresses []PhysicalOfficeAddressable
    // Name of the company that this organizational contact belongs to.  Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
    companyName *string
    // The name for the department in which the contact works.  Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
    department *string
    // The contact's direct reports. (The users and contacts that have their manager property set to this contact.)  Read-only. Nullable. Supports $expand.
    directReports []DirectoryObjectable
    // Display name for this organizational contact. Maximum length is 256 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values), $search, and $orderby.
    displayName *string
    // First name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
    givenName *string
    // Job title for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
    jobTitle *string
    // The SMTP address for the contact, for example, 'jeff@contoso.com'. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
    mail *string
    // Email alias (portion of email address pre-pending the @ symbol) for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
    mailNickname *string
    // The user or contact that is this contact's manager. Read-only. Supports $expand and $filter (eq) by id.
    manager DirectoryObjectable
    // Groups that this contact is a member of. Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObjectable
    // Date and time when this organizational contact was last synchronized from on-premises AD. This date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, not, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of any synchronization provisioning errors for this organizational contact. Supports $filter (eq, not for category and propertyCausingError), /$count eq 0, /$count ne 0.
    onPremisesProvisioningErrors []OnPremisesProvisioningErrorable
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced and now mastered in Exchange; null if this object has never been synced from an on-premises directory (default).   Supports $filter (eq, ne, not, in, and eq for null values).
    onPremisesSyncEnabled *bool
    // List of phones for this organizational contact. Phone types can be mobile, business, and businessFax. Only one of each type can ever be present in the collection.
    phones []Phoneable
    // For example: 'SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com'. The any operator is required for filter expressions on multi-valued properties. Supports $filter (eq, not, ge, le, startsWith, /$count eq 0, /$count ne 0).
    proxyAddresses []string
    // Errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from an organizational contact object .  Supports $filter (eq, not, for isResolved and serviceInstance).
    serviceProvisioningErrors []ServiceProvisioningErrorable
    // Last name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
    surname *string
    // Groups that this contact is a member of, including groups that the contact is nested under. Read-only. Nullable.
    transitiveMemberOf []DirectoryObjectable
}
// NewOrgContact instantiates a new OrgContact and sets the default values.
func NewOrgContact()(*OrgContact) {
    m := &OrgContact{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.orgContact"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOrgContactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrgContactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrgContact(), nil
}
// GetAddresses gets the addresses property value. Postal addresses for this organizational contact. For now a contact can only have one physical address.
// returns a []PhysicalOfficeAddressable when successful
func (m *OrgContact) GetAddresses()([]PhysicalOfficeAddressable) {
    return m.addresses
}
// GetCompanyName gets the companyName property value. Name of the company that this organizational contact belongs to.  Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
// returns a *string when successful
func (m *OrgContact) GetCompanyName()(*string) {
    return m.companyName
}
// GetDepartment gets the department property value. The name for the department in which the contact works.  Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
// returns a *string when successful
func (m *OrgContact) GetDepartment()(*string) {
    return m.department
}
// GetDirectReports gets the directReports property value. The contact's direct reports. (The users and contacts that have their manager property set to this contact.)  Read-only. Nullable. Supports $expand.
// returns a []DirectoryObjectable when successful
func (m *OrgContact) GetDirectReports()([]DirectoryObjectable) {
    return m.directReports
}
// GetDisplayName gets the displayName property value. Display name for this organizational contact. Maximum length is 256 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values), $search, and $orderby.
// returns a *string when successful
func (m *OrgContact) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrgContact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhysicalOfficeAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PhysicalOfficeAddressable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PhysicalOfficeAddressable)
                }
            }
            m.SetAddresses(res)
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
    res["phones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Phoneable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Phoneable)
                }
            }
            m.SetPhones(res)
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
    return res
}
// GetGivenName gets the givenName property value. First name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
// returns a *string when successful
func (m *OrgContact) GetGivenName()(*string) {
    return m.givenName
}
// GetJobTitle gets the jobTitle property value. Job title for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
// returns a *string when successful
func (m *OrgContact) GetJobTitle()(*string) {
    return m.jobTitle
}
// GetMail gets the mail property value. The SMTP address for the contact, for example, 'jeff@contoso.com'. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
// returns a *string when successful
func (m *OrgContact) GetMail()(*string) {
    return m.mail
}
// GetMailNickname gets the mailNickname property value. Email alias (portion of email address pre-pending the @ symbol) for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
// returns a *string when successful
func (m *OrgContact) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetManager gets the manager property value. The user or contact that is this contact's manager. Read-only. Supports $expand and $filter (eq) by id.
// returns a DirectoryObjectable when successful
func (m *OrgContact) GetManager()(DirectoryObjectable) {
    return m.manager
}
// GetMemberOf gets the memberOf property value. Groups that this contact is a member of. Read-only. Nullable. Supports $expand.
// returns a []DirectoryObjectable when successful
func (m *OrgContact) GetMemberOf()([]DirectoryObjectable) {
    return m.memberOf
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. Date and time when this organizational contact was last synchronized from on-premises AD. This date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, not, ge, le, in).
// returns a *Time when successful
func (m *OrgContact) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.onPremisesLastSyncDateTime
}
// GetOnPremisesProvisioningErrors gets the onPremisesProvisioningErrors property value. List of any synchronization provisioning errors for this organizational contact. Supports $filter (eq, not for category and propertyCausingError), /$count eq 0, /$count ne 0.
// returns a []OnPremisesProvisioningErrorable when successful
func (m *OrgContact) GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningErrorable) {
    return m.onPremisesProvisioningErrors
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced and now mastered in Exchange; null if this object has never been synced from an on-premises directory (default).   Supports $filter (eq, ne, not, in, and eq for null values).
// returns a *bool when successful
func (m *OrgContact) GetOnPremisesSyncEnabled()(*bool) {
    return m.onPremisesSyncEnabled
}
// GetPhones gets the phones property value. List of phones for this organizational contact. Phone types can be mobile, business, and businessFax. Only one of each type can ever be present in the collection.
// returns a []Phoneable when successful
func (m *OrgContact) GetPhones()([]Phoneable) {
    return m.phones
}
// GetProxyAddresses gets the proxyAddresses property value. For example: 'SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com'. The any operator is required for filter expressions on multi-valued properties. Supports $filter (eq, not, ge, le, startsWith, /$count eq 0, /$count ne 0).
// returns a []string when successful
func (m *OrgContact) GetProxyAddresses()([]string) {
    return m.proxyAddresses
}
// GetServiceProvisioningErrors gets the serviceProvisioningErrors property value. Errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from an organizational contact object .  Supports $filter (eq, not, for isResolved and serviceInstance).
// returns a []ServiceProvisioningErrorable when successful
func (m *OrgContact) GetServiceProvisioningErrors()([]ServiceProvisioningErrorable) {
    return m.serviceProvisioningErrors
}
// GetSurname gets the surname property value. Last name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
// returns a *string when successful
func (m *OrgContact) GetSurname()(*string) {
    return m.surname
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. Groups that this contact is a member of, including groups that the contact is nested under. Read-only. Nullable.
// returns a []DirectoryObjectable when successful
func (m *OrgContact) GetTransitiveMemberOf()([]DirectoryObjectable) {
    return m.transitiveMemberOf
}
// Serialize serializes information the current object
func (m *OrgContact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddresses()))
        for i, v := range m.GetAddresses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("addresses", cast)
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
        err = writer.WriteStringValue("department", m.GetDepartment())
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
        err = writer.WriteStringValue("givenName", m.GetGivenName())
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
    {
        err = writer.WriteStringValue("mail", m.GetMail())
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
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetPhones() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
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
        err = writer.WriteStringValue("surname", m.GetSurname())
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
    return nil
}
// SetAddresses sets the addresses property value. Postal addresses for this organizational contact. For now a contact can only have one physical address.
func (m *OrgContact) SetAddresses(value []PhysicalOfficeAddressable)() {
    m.addresses = value
}
// SetCompanyName sets the companyName property value. Name of the company that this organizational contact belongs to.  Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
func (m *OrgContact) SetCompanyName(value *string)() {
    m.companyName = value
}
// SetDepartment sets the department property value. The name for the department in which the contact works.  Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
func (m *OrgContact) SetDepartment(value *string)() {
    m.department = value
}
// SetDirectReports sets the directReports property value. The contact's direct reports. (The users and contacts that have their manager property set to this contact.)  Read-only. Nullable. Supports $expand.
func (m *OrgContact) SetDirectReports(value []DirectoryObjectable)() {
    m.directReports = value
}
// SetDisplayName sets the displayName property value. Display name for this organizational contact. Maximum length is 256 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values), $search, and $orderby.
func (m *OrgContact) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGivenName sets the givenName property value. First name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
func (m *OrgContact) SetGivenName(value *string)() {
    m.givenName = value
}
// SetJobTitle sets the jobTitle property value. Job title for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
func (m *OrgContact) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// SetMail sets the mail property value. The SMTP address for the contact, for example, 'jeff@contoso.com'. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
func (m *OrgContact) SetMail(value *string)() {
    m.mail = value
}
// SetMailNickname sets the mailNickname property value. Email alias (portion of email address pre-pending the @ symbol) for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
func (m *OrgContact) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetManager sets the manager property value. The user or contact that is this contact's manager. Read-only. Supports $expand and $filter (eq) by id.
func (m *OrgContact) SetManager(value DirectoryObjectable)() {
    m.manager = value
}
// SetMemberOf sets the memberOf property value. Groups that this contact is a member of. Read-only. Nullable. Supports $expand.
func (m *OrgContact) SetMemberOf(value []DirectoryObjectable)() {
    m.memberOf = value
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. Date and time when this organizational contact was last synchronized from on-premises AD. This date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, not, ge, le, in).
func (m *OrgContact) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// SetOnPremisesProvisioningErrors sets the onPremisesProvisioningErrors property value. List of any synchronization provisioning errors for this organizational contact. Supports $filter (eq, not for category and propertyCausingError), /$count eq 0, /$count ne 0.
func (m *OrgContact) SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningErrorable)() {
    m.onPremisesProvisioningErrors = value
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced and now mastered in Exchange; null if this object has never been synced from an on-premises directory (default).   Supports $filter (eq, ne, not, in, and eq for null values).
func (m *OrgContact) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// SetPhones sets the phones property value. List of phones for this organizational contact. Phone types can be mobile, business, and businessFax. Only one of each type can ever be present in the collection.
func (m *OrgContact) SetPhones(value []Phoneable)() {
    m.phones = value
}
// SetProxyAddresses sets the proxyAddresses property value. For example: 'SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com'. The any operator is required for filter expressions on multi-valued properties. Supports $filter (eq, not, ge, le, startsWith, /$count eq 0, /$count ne 0).
func (m *OrgContact) SetProxyAddresses(value []string)() {
    m.proxyAddresses = value
}
// SetServiceProvisioningErrors sets the serviceProvisioningErrors property value. Errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from an organizational contact object .  Supports $filter (eq, not, for isResolved and serviceInstance).
func (m *OrgContact) SetServiceProvisioningErrors(value []ServiceProvisioningErrorable)() {
    m.serviceProvisioningErrors = value
}
// SetSurname sets the surname property value. Last name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null values).
func (m *OrgContact) SetSurname(value *string)() {
    m.surname = value
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. Groups that this contact is a member of, including groups that the contact is nested under. Read-only. Nullable.
func (m *OrgContact) SetTransitiveMemberOf(value []DirectoryObjectable)() {
    m.transitiveMemberOf = value
}
type OrgContactable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddresses()([]PhysicalOfficeAddressable)
    GetCompanyName()(*string)
    GetDepartment()(*string)
    GetDirectReports()([]DirectoryObjectable)
    GetDisplayName()(*string)
    GetGivenName()(*string)
    GetJobTitle()(*string)
    GetMail()(*string)
    GetMailNickname()(*string)
    GetManager()(DirectoryObjectable)
    GetMemberOf()([]DirectoryObjectable)
    GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningErrorable)
    GetOnPremisesSyncEnabled()(*bool)
    GetPhones()([]Phoneable)
    GetProxyAddresses()([]string)
    GetServiceProvisioningErrors()([]ServiceProvisioningErrorable)
    GetSurname()(*string)
    GetTransitiveMemberOf()([]DirectoryObjectable)
    SetAddresses(value []PhysicalOfficeAddressable)()
    SetCompanyName(value *string)()
    SetDepartment(value *string)()
    SetDirectReports(value []DirectoryObjectable)()
    SetDisplayName(value *string)()
    SetGivenName(value *string)()
    SetJobTitle(value *string)()
    SetMail(value *string)()
    SetMailNickname(value *string)()
    SetManager(value DirectoryObjectable)()
    SetMemberOf(value []DirectoryObjectable)()
    SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningErrorable)()
    SetOnPremisesSyncEnabled(value *bool)()
    SetPhones(value []Phoneable)()
    SetProxyAddresses(value []string)()
    SetServiceProvisioningErrors(value []ServiceProvisioningErrorable)()
    SetSurname(value *string)()
    SetTransitiveMemberOf(value []DirectoryObjectable)()
}
