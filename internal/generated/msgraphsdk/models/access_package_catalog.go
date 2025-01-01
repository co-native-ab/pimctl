package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageCatalog struct {
    Entity
    // The access packages in this catalog. Read-only. Nullable.
    accessPackages []AccessPackageable
    // Whether the catalog is created by a user or entitlement management. The possible values are: userManaged, serviceDefault, serviceManaged, unknownFutureValue.
    catalogType *AccessPackageCatalogType
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customWorkflowExtensions property
    customWorkflowExtensions []CustomCalloutExtensionable
    // The description of the access package catalog.
    description *string
    // The display name of the access package catalog.
    displayName *string
    // Whether the access packages in this catalog can be requested by users outside of the tenant.
    isExternallyVisible *bool
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The resourceRoles property
    resourceRoles []AccessPackageResourceRoleable
    // Access package resources in this catalog.
    resources []AccessPackageResourceable
    // The resourceScopes property
    resourceScopes []AccessPackageResourceScopeable
    // Has the value published if the access packages are available for management. The possible values are: unpublished, published, unknownFutureValue.
    state *AccessPackageCatalogState
}
// NewAccessPackageCatalog instantiates a new AccessPackageCatalog and sets the default values.
func NewAccessPackageCatalog()(*AccessPackageCatalog) {
    m := &AccessPackageCatalog{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageCatalogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageCatalogFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageCatalog(), nil
}
// GetAccessPackages gets the accessPackages property value. The access packages in this catalog. Read-only. Nullable.
// returns a []AccessPackageable when successful
func (m *AccessPackageCatalog) GetAccessPackages()([]AccessPackageable) {
    return m.accessPackages
}
// GetCatalogType gets the catalogType property value. Whether the catalog is created by a user or entitlement management. The possible values are: userManaged, serviceDefault, serviceManaged, unknownFutureValue.
// returns a *AccessPackageCatalogType when successful
func (m *AccessPackageCatalog) GetCatalogType()(*AccessPackageCatalogType) {
    return m.catalogType
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *AccessPackageCatalog) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomWorkflowExtensions gets the customWorkflowExtensions property value. The customWorkflowExtensions property
// returns a []CustomCalloutExtensionable when successful
func (m *AccessPackageCatalog) GetCustomWorkflowExtensions()([]CustomCalloutExtensionable) {
    return m.customWorkflowExtensions
}
// GetDescription gets the description property value. The description of the access package catalog.
// returns a *string when successful
func (m *AccessPackageCatalog) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the access package catalog.
// returns a *string when successful
func (m *AccessPackageCatalog) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageCatalog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageable)
                }
            }
            m.SetAccessPackages(res)
        }
        return nil
    }
    res["catalogType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCatalogType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogType(val.(*AccessPackageCatalogType))
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
    res["customWorkflowExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomCalloutExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomCalloutExtensionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomCalloutExtensionable)
                }
            }
            m.SetCustomWorkflowExtensions(res)
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
    res["isExternallyVisible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternallyVisible(val)
        }
        return nil
    }
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["resourceRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceRoleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceRoleable)
                }
            }
            m.SetResourceRoles(res)
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceable)
                }
            }
            m.SetResources(res)
        }
        return nil
    }
    res["resourceScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageResourceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageResourceScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageResourceScopeable)
                }
            }
            m.SetResourceScopes(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCatalogState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AccessPackageCatalogState))
        }
        return nil
    }
    return res
}
// GetIsExternallyVisible gets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
// returns a *bool when successful
func (m *AccessPackageCatalog) GetIsExternallyVisible()(*bool) {
    return m.isExternallyVisible
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *AccessPackageCatalog) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetResourceRoles gets the resourceRoles property value. The resourceRoles property
// returns a []AccessPackageResourceRoleable when successful
func (m *AccessPackageCatalog) GetResourceRoles()([]AccessPackageResourceRoleable) {
    return m.resourceRoles
}
// GetResources gets the resources property value. Access package resources in this catalog.
// returns a []AccessPackageResourceable when successful
func (m *AccessPackageCatalog) GetResources()([]AccessPackageResourceable) {
    return m.resources
}
// GetResourceScopes gets the resourceScopes property value. The resourceScopes property
// returns a []AccessPackageResourceScopeable when successful
func (m *AccessPackageCatalog) GetResourceScopes()([]AccessPackageResourceScopeable) {
    return m.resourceScopes
}
// GetState gets the state property value. Has the value published if the access packages are available for management. The possible values are: unpublished, published, unknownFutureValue.
// returns a *AccessPackageCatalogState when successful
func (m *AccessPackageCatalog) GetState()(*AccessPackageCatalogState) {
    return m.state
}
// Serialize serializes information the current object
func (m *AccessPackageCatalog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackages()))
        for i, v := range m.GetAccessPackages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCatalogType() != nil {
        cast := (*m.GetCatalogType()).String()
        err = writer.WriteStringValue("catalogType", &cast)
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
    if m.GetCustomWorkflowExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomWorkflowExtensions()))
        for i, v := range m.GetCustomWorkflowExtensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customWorkflowExtensions", cast)
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
        err = writer.WriteBoolValue("isExternallyVisible", m.GetIsExternallyVisible())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetResourceRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceRoles()))
        for i, v := range m.GetResourceRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceRoles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceScopes()))
        for i, v := range m.GetResourceScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackages sets the accessPackages property value. The access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageCatalog) SetAccessPackages(value []AccessPackageable)() {
    m.accessPackages = value
}
// SetCatalogType sets the catalogType property value. Whether the catalog is created by a user or entitlement management. The possible values are: userManaged, serviceDefault, serviceManaged, unknownFutureValue.
func (m *AccessPackageCatalog) SetCatalogType(value *AccessPackageCatalogType)() {
    m.catalogType = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomWorkflowExtensions sets the customWorkflowExtensions property value. The customWorkflowExtensions property
func (m *AccessPackageCatalog) SetCustomWorkflowExtensions(value []CustomCalloutExtensionable)() {
    m.customWorkflowExtensions = value
}
// SetDescription sets the description property value. The description of the access package catalog.
func (m *AccessPackageCatalog) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the access package catalog.
func (m *AccessPackageCatalog) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsExternallyVisible sets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
func (m *AccessPackageCatalog) SetIsExternallyVisible(value *bool)() {
    m.isExternallyVisible = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetResourceRoles sets the resourceRoles property value. The resourceRoles property
func (m *AccessPackageCatalog) SetResourceRoles(value []AccessPackageResourceRoleable)() {
    m.resourceRoles = value
}
// SetResources sets the resources property value. Access package resources in this catalog.
func (m *AccessPackageCatalog) SetResources(value []AccessPackageResourceable)() {
    m.resources = value
}
// SetResourceScopes sets the resourceScopes property value. The resourceScopes property
func (m *AccessPackageCatalog) SetResourceScopes(value []AccessPackageResourceScopeable)() {
    m.resourceScopes = value
}
// SetState sets the state property value. Has the value published if the access packages are available for management. The possible values are: unpublished, published, unknownFutureValue.
func (m *AccessPackageCatalog) SetState(value *AccessPackageCatalogState)() {
    m.state = value
}
type AccessPackageCatalogable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackages()([]AccessPackageable)
    GetCatalogType()(*AccessPackageCatalogType)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomWorkflowExtensions()([]CustomCalloutExtensionable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsExternallyVisible()(*bool)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResourceRoles()([]AccessPackageResourceRoleable)
    GetResources()([]AccessPackageResourceable)
    GetResourceScopes()([]AccessPackageResourceScopeable)
    GetState()(*AccessPackageCatalogState)
    SetAccessPackages(value []AccessPackageable)()
    SetCatalogType(value *AccessPackageCatalogType)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomWorkflowExtensions(value []CustomCalloutExtensionable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsExternallyVisible(value *bool)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResourceRoles(value []AccessPackageResourceRoleable)()
    SetResources(value []AccessPackageResourceable)()
    SetResourceScopes(value []AccessPackageResourceScopeable)()
    SetState(value *AccessPackageCatalogState)()
}
