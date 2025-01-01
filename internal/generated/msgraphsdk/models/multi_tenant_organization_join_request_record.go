package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MultiTenantOrganizationJoinRequestRecord struct {
    Entity
    // Tenant ID of the Microsoft Entra tenant that added a tenant to the multitenant organization. To reset a failed join request, set addedByTenantId to 00000000-0000-0000-0000-000000000000. Required.
    addedByTenantId *string
    // State of the tenant in the multitenant organization. The possible values are: pending, active, removed, unknownFutureValue. Tenants in the pending state must join the multitenant organization to participate in the multitenant organization. Tenants in the active state can participate in the multitenant organization. Tenants in the removed state are in the process of being removed from the multitenant organization. Read-only.
    memberState *MultiTenantOrganizationMemberState
    // Role of the tenant in the multitenant organization. The possible values are: owner, member (default), unknownFutureValue. Tenants with the owner role can manage the multitenant organization. There can be multiple tenants with the owner role in a multitenant organization. Tenants with the member role can participate in a multitenant organization.
    role *MultiTenantOrganizationMemberRole
    // Details of the processing status for a tenant joining a multitenant organization. Read-only.
    transitionDetails MultiTenantOrganizationJoinRequestTransitionDetailsable
}
// NewMultiTenantOrganizationJoinRequestRecord instantiates a new MultiTenantOrganizationJoinRequestRecord and sets the default values.
func NewMultiTenantOrganizationJoinRequestRecord()(*MultiTenantOrganizationJoinRequestRecord) {
    m := &MultiTenantOrganizationJoinRequestRecord{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMultiTenantOrganizationJoinRequestRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMultiTenantOrganizationJoinRequestRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMultiTenantOrganizationJoinRequestRecord(), nil
}
// GetAddedByTenantId gets the addedByTenantId property value. Tenant ID of the Microsoft Entra tenant that added a tenant to the multitenant organization. To reset a failed join request, set addedByTenantId to 00000000-0000-0000-0000-000000000000. Required.
// returns a *string when successful
func (m *MultiTenantOrganizationJoinRequestRecord) GetAddedByTenantId()(*string) {
    return m.addedByTenantId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MultiTenantOrganizationJoinRequestRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addedByTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedByTenantId(val)
        }
        return nil
    }
    res["memberState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiTenantOrganizationMemberState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberState(val.(*MultiTenantOrganizationMemberState))
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiTenantOrganizationMemberRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*MultiTenantOrganizationMemberRole))
        }
        return nil
    }
    res["transitionDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMultiTenantOrganizationJoinRequestTransitionDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransitionDetails(val.(MultiTenantOrganizationJoinRequestTransitionDetailsable))
        }
        return nil
    }
    return res
}
// GetMemberState gets the memberState property value. State of the tenant in the multitenant organization. The possible values are: pending, active, removed, unknownFutureValue. Tenants in the pending state must join the multitenant organization to participate in the multitenant organization. Tenants in the active state can participate in the multitenant organization. Tenants in the removed state are in the process of being removed from the multitenant organization. Read-only.
// returns a *MultiTenantOrganizationMemberState when successful
func (m *MultiTenantOrganizationJoinRequestRecord) GetMemberState()(*MultiTenantOrganizationMemberState) {
    return m.memberState
}
// GetRole gets the role property value. Role of the tenant in the multitenant organization. The possible values are: owner, member (default), unknownFutureValue. Tenants with the owner role can manage the multitenant organization. There can be multiple tenants with the owner role in a multitenant organization. Tenants with the member role can participate in a multitenant organization.
// returns a *MultiTenantOrganizationMemberRole when successful
func (m *MultiTenantOrganizationJoinRequestRecord) GetRole()(*MultiTenantOrganizationMemberRole) {
    return m.role
}
// GetTransitionDetails gets the transitionDetails property value. Details of the processing status for a tenant joining a multitenant organization. Read-only.
// returns a MultiTenantOrganizationJoinRequestTransitionDetailsable when successful
func (m *MultiTenantOrganizationJoinRequestRecord) GetTransitionDetails()(MultiTenantOrganizationJoinRequestTransitionDetailsable) {
    return m.transitionDetails
}
// Serialize serializes information the current object
func (m *MultiTenantOrganizationJoinRequestRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("addedByTenantId", m.GetAddedByTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetMemberState() != nil {
        cast := (*m.GetMemberState()).String()
        err = writer.WriteStringValue("memberState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err = writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("transitionDetails", m.GetTransitionDetails())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddedByTenantId sets the addedByTenantId property value. Tenant ID of the Microsoft Entra tenant that added a tenant to the multitenant organization. To reset a failed join request, set addedByTenantId to 00000000-0000-0000-0000-000000000000. Required.
func (m *MultiTenantOrganizationJoinRequestRecord) SetAddedByTenantId(value *string)() {
    m.addedByTenantId = value
}
// SetMemberState sets the memberState property value. State of the tenant in the multitenant organization. The possible values are: pending, active, removed, unknownFutureValue. Tenants in the pending state must join the multitenant organization to participate in the multitenant organization. Tenants in the active state can participate in the multitenant organization. Tenants in the removed state are in the process of being removed from the multitenant organization. Read-only.
func (m *MultiTenantOrganizationJoinRequestRecord) SetMemberState(value *MultiTenantOrganizationMemberState)() {
    m.memberState = value
}
// SetRole sets the role property value. Role of the tenant in the multitenant organization. The possible values are: owner, member (default), unknownFutureValue. Tenants with the owner role can manage the multitenant organization. There can be multiple tenants with the owner role in a multitenant organization. Tenants with the member role can participate in a multitenant organization.
func (m *MultiTenantOrganizationJoinRequestRecord) SetRole(value *MultiTenantOrganizationMemberRole)() {
    m.role = value
}
// SetTransitionDetails sets the transitionDetails property value. Details of the processing status for a tenant joining a multitenant organization. Read-only.
func (m *MultiTenantOrganizationJoinRequestRecord) SetTransitionDetails(value MultiTenantOrganizationJoinRequestTransitionDetailsable)() {
    m.transitionDetails = value
}
type MultiTenantOrganizationJoinRequestRecordable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedByTenantId()(*string)
    GetMemberState()(*MultiTenantOrganizationMemberState)
    GetRole()(*MultiTenantOrganizationMemberRole)
    GetTransitionDetails()(MultiTenantOrganizationJoinRequestTransitionDetailsable)
    SetAddedByTenantId(value *string)()
    SetMemberState(value *MultiTenantOrganizationMemberState)()
    SetRole(value *MultiTenantOrganizationMemberRole)()
    SetTransitionDetails(value MultiTenantOrganizationJoinRequestTransitionDetailsable)()
}
