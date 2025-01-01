package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PrivilegedAccessGroupEligibilityScheduleRequest struct {
    PrivilegedAccessScheduleRequest
    // The identifier of membership or ownership eligibility relationship to the group. Required. The possible values are: owner, member, unknownFutureValue.
    accessId *PrivilegedAccessGroupRelationships
    // References the group that is the scope of the membership or ownership eligibility request through PIM for groups. Supports $expand and $select nested in $expand for select properties like id, displayName, and mail.
    group Groupable
    // The identifier of the group representing the scope of the membership and ownership eligibility through PIM for groups. Required.
    groupId *string
    // References the principal that's in the scope of the membership or ownership eligibility request through the group that's governed by PIM. Supports $expand and $select nested in $expand for id only.
    principal DirectoryObjectable
    // The identifier of the principal whose membership or ownership eligibility to the group is managed through PIM for groups. Required.
    principalId *string
    // Schedule created by this request.
    targetSchedule PrivilegedAccessGroupEligibilityScheduleable
    // The identifier of the schedule that's created from the eligibility request. Optional.
    targetScheduleId *string
}
// NewPrivilegedAccessGroupEligibilityScheduleRequest instantiates a new PrivilegedAccessGroupEligibilityScheduleRequest and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleRequest()(*PrivilegedAccessGroupEligibilityScheduleRequest) {
    m := &PrivilegedAccessGroupEligibilityScheduleRequest{
        PrivilegedAccessScheduleRequest: *NewPrivilegedAccessScheduleRequest(),
    }
    odataTypeValue := "#microsoft.graph.privilegedAccessGroupEligibilityScheduleRequest"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePrivilegedAccessGroupEligibilityScheduleRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivilegedAccessGroupEligibilityScheduleRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequest(), nil
}
// GetAccessId gets the accessId property value. The identifier of membership or ownership eligibility relationship to the group. Required. The possible values are: owner, member, unknownFutureValue.
// returns a *PrivilegedAccessGroupRelationships when successful
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetAccessId()(*PrivilegedAccessGroupRelationships) {
    return m.accessId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrivilegedAccessScheduleRequest.GetFieldDeserializers()
    res["accessId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivilegedAccessGroupRelationships)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessId(val.(*PrivilegedAccessGroupRelationships))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(Groupable))
        }
        return nil
    }
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["principal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipal(val.(DirectoryObjectable))
        }
        return nil
    }
    res["principalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    res["targetSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetSchedule(val.(PrivilegedAccessGroupEligibilityScheduleable))
        }
        return nil
    }
    res["targetScheduleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetScheduleId(val)
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. References the group that is the scope of the membership or ownership eligibility request through PIM for groups. Supports $expand and $select nested in $expand for select properties like id, displayName, and mail.
// returns a Groupable when successful
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetGroup()(Groupable) {
    return m.group
}
// GetGroupId gets the groupId property value. The identifier of the group representing the scope of the membership and ownership eligibility through PIM for groups. Required.
// returns a *string when successful
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetGroupId()(*string) {
    return m.groupId
}
// GetPrincipal gets the principal property value. References the principal that's in the scope of the membership or ownership eligibility request through the group that's governed by PIM. Supports $expand and $select nested in $expand for id only.
// returns a DirectoryObjectable when successful
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. The identifier of the principal whose membership or ownership eligibility to the group is managed through PIM for groups. Required.
// returns a *string when successful
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetPrincipalId()(*string) {
    return m.principalId
}
// GetTargetSchedule gets the targetSchedule property value. Schedule created by this request.
// returns a PrivilegedAccessGroupEligibilityScheduleable when successful
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetTargetSchedule()(PrivilegedAccessGroupEligibilityScheduleable) {
    return m.targetSchedule
}
// GetTargetScheduleId gets the targetScheduleId property value. The identifier of the schedule that's created from the eligibility request. Optional.
// returns a *string when successful
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetTargetScheduleId()(*string) {
    return m.targetScheduleId
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrivilegedAccessScheduleRequest.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessId() != nil {
        cast := (*m.GetAccessId()).String()
        err = writer.WriteStringValue("accessId", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principal", m.GetPrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetSchedule", m.GetTargetSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetScheduleId", m.GetTargetScheduleId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessId sets the accessId property value. The identifier of membership or ownership eligibility relationship to the group. Required. The possible values are: owner, member, unknownFutureValue.
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetAccessId(value *PrivilegedAccessGroupRelationships)() {
    m.accessId = value
}
// SetGroup sets the group property value. References the group that is the scope of the membership or ownership eligibility request through PIM for groups. Supports $expand and $select nested in $expand for select properties like id, displayName, and mail.
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetGroup(value Groupable)() {
    m.group = value
}
// SetGroupId sets the groupId property value. The identifier of the group representing the scope of the membership and ownership eligibility through PIM for groups. Required.
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetGroupId(value *string)() {
    m.groupId = value
}
// SetPrincipal sets the principal property value. References the principal that's in the scope of the membership or ownership eligibility request through the group that's governed by PIM. Supports $expand and $select nested in $expand for id only.
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. The identifier of the principal whose membership or ownership eligibility to the group is managed through PIM for groups. Required.
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetPrincipalId(value *string)() {
    m.principalId = value
}
// SetTargetSchedule sets the targetSchedule property value. Schedule created by this request.
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetTargetSchedule(value PrivilegedAccessGroupEligibilityScheduleable)() {
    m.targetSchedule = value
}
// SetTargetScheduleId sets the targetScheduleId property value. The identifier of the schedule that's created from the eligibility request. Optional.
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetTargetScheduleId(value *string)() {
    m.targetScheduleId = value
}
type PrivilegedAccessGroupEligibilityScheduleRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrivilegedAccessScheduleRequestable
    GetAccessId()(*PrivilegedAccessGroupRelationships)
    GetGroup()(Groupable)
    GetGroupId()(*string)
    GetPrincipal()(DirectoryObjectable)
    GetPrincipalId()(*string)
    GetTargetSchedule()(PrivilegedAccessGroupEligibilityScheduleable)
    GetTargetScheduleId()(*string)
    SetAccessId(value *PrivilegedAccessGroupRelationships)()
    SetGroup(value Groupable)()
    SetGroupId(value *string)()
    SetPrincipal(value DirectoryObjectable)()
    SetPrincipalId(value *string)()
    SetTargetSchedule(value PrivilegedAccessGroupEligibilityScheduleable)()
    SetTargetScheduleId(value *string)()
}
