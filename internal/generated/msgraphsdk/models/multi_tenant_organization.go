package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MultiTenantOrganization struct {
    Entity
    // Date when multitenant organization was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description of the multitenant organization.
    description *string
    // Display name of the multitenant organization.
    displayName *string
    // Defines the status of a tenant joining a multitenant organization.
    joinRequest MultiTenantOrganizationJoinRequestRecordable
    // State of the multitenant organization. The possible values are: active, inactive, unknownFutureValue. active indicates the multitenant organization is created. inactive indicates the multitenant organization isn't created. Read-only.
    state *MultiTenantOrganizationState
    // Defines tenants added to a multitenant organization.
    tenants []MultiTenantOrganizationMemberable
}
// NewMultiTenantOrganization instantiates a new MultiTenantOrganization and sets the default values.
func NewMultiTenantOrganization()(*MultiTenantOrganization) {
    m := &MultiTenantOrganization{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMultiTenantOrganizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMultiTenantOrganizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMultiTenantOrganization(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Date when multitenant organization was created. Read-only.
// returns a *Time when successful
func (m *MultiTenantOrganization) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description of the multitenant organization.
// returns a *string when successful
func (m *MultiTenantOrganization) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name of the multitenant organization.
// returns a *string when successful
func (m *MultiTenantOrganization) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MultiTenantOrganization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["joinRequest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMultiTenantOrganizationJoinRequestRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinRequest(val.(MultiTenantOrganizationJoinRequestRecordable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiTenantOrganizationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*MultiTenantOrganizationState))
        }
        return nil
    }
    res["tenants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMultiTenantOrganizationMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiTenantOrganizationMemberable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MultiTenantOrganizationMemberable)
                }
            }
            m.SetTenants(res)
        }
        return nil
    }
    return res
}
// GetJoinRequest gets the joinRequest property value. Defines the status of a tenant joining a multitenant organization.
// returns a MultiTenantOrganizationJoinRequestRecordable when successful
func (m *MultiTenantOrganization) GetJoinRequest()(MultiTenantOrganizationJoinRequestRecordable) {
    return m.joinRequest
}
// GetState gets the state property value. State of the multitenant organization. The possible values are: active, inactive, unknownFutureValue. active indicates the multitenant organization is created. inactive indicates the multitenant organization isn't created. Read-only.
// returns a *MultiTenantOrganizationState when successful
func (m *MultiTenantOrganization) GetState()(*MultiTenantOrganizationState) {
    return m.state
}
// GetTenants gets the tenants property value. Defines tenants added to a multitenant organization.
// returns a []MultiTenantOrganizationMemberable when successful
func (m *MultiTenantOrganization) GetTenants()([]MultiTenantOrganizationMemberable) {
    return m.tenants
}
// Serialize serializes information the current object
func (m *MultiTenantOrganization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err = writer.WriteObjectValue("joinRequest", m.GetJoinRequest())
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
    if m.GetTenants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenants()))
        for i, v := range m.GetTenants() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("tenants", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Date when multitenant organization was created. Read-only.
func (m *MultiTenantOrganization) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description of the multitenant organization.
func (m *MultiTenantOrganization) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name of the multitenant organization.
func (m *MultiTenantOrganization) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetJoinRequest sets the joinRequest property value. Defines the status of a tenant joining a multitenant organization.
func (m *MultiTenantOrganization) SetJoinRequest(value MultiTenantOrganizationJoinRequestRecordable)() {
    m.joinRequest = value
}
// SetState sets the state property value. State of the multitenant organization. The possible values are: active, inactive, unknownFutureValue. active indicates the multitenant organization is created. inactive indicates the multitenant organization isn't created. Read-only.
func (m *MultiTenantOrganization) SetState(value *MultiTenantOrganizationState)() {
    m.state = value
}
// SetTenants sets the tenants property value. Defines tenants added to a multitenant organization.
func (m *MultiTenantOrganization) SetTenants(value []MultiTenantOrganizationMemberable)() {
    m.tenants = value
}
type MultiTenantOrganizationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetJoinRequest()(MultiTenantOrganizationJoinRequestRecordable)
    GetState()(*MultiTenantOrganizationState)
    GetTenants()([]MultiTenantOrganizationMemberable)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetJoinRequest(value MultiTenantOrganizationJoinRequestRecordable)()
    SetState(value *MultiTenantOrganizationState)()
    SetTenants(value []MultiTenantOrganizationMemberable)()
}
