package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PrivilegedAccessGroupAssignmentScheduleRequest struct {
    PrivilegedAccessScheduleRequest
    // The identifier of a membership or ownership assignment relationship to the group. Required. The possible values are: owner, member, unknownFutureValue.
    accessId *PrivilegedAccessGroupRelationships
    // When the request activates a membership or ownership assignment in PIM for groups, this object represents the eligibility policy for the group. Otherwise, it is null. Supports $expand.
    activatedUsing PrivilegedAccessGroupEligibilityScheduleable
    // References the group that is the scope of the membership or ownership assignment request through PIM for groups. Supports $expand and $select nested in $expand for select properties like id, displayName, and mail.
    group Groupable
    // The identifier of the group representing the scope of the membership or ownership assignment through PIM for groups. Required.
    groupId *string
    // References the principal that's in the scope of this membership or ownership assignment request through the group that's governed by PIM. Supports $expand and $select nested in $expand for id only.
    principal DirectoryObjectable
    // The identifier of the principal whose membership or ownership assignment to the group is managed through PIM for groups. Supports $filter (eq, ne).
    principalId *string
    // Schedule created by this request. Supports $expand.
    targetSchedule PrivilegedAccessGroupEligibilityScheduleable
    // The identifier of the schedule that's created from the membership or ownership assignment request. Supports $filter (eq, ne).
    targetScheduleId *string
}
// NewPrivilegedAccessGroupAssignmentScheduleRequest instantiates a new PrivilegedAccessGroupAssignmentScheduleRequest and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleRequest()(*PrivilegedAccessGroupAssignmentScheduleRequest) {
    m := &PrivilegedAccessGroupAssignmentScheduleRequest{
        PrivilegedAccessScheduleRequest: *NewPrivilegedAccessScheduleRequest(),
    }
    odataTypeValue := "#microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePrivilegedAccessGroupAssignmentScheduleRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivilegedAccessGroupAssignmentScheduleRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequest(), nil
}
// GetAccessId gets the accessId property value. The identifier of a membership or ownership assignment relationship to the group. Required. The possible values are: owner, member, unknownFutureValue.
// returns a *PrivilegedAccessGroupRelationships when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetAccessId()(*PrivilegedAccessGroupRelationships) {
    return m.accessId
}
// GetActivatedUsing gets the activatedUsing property value. When the request activates a membership or ownership assignment in PIM for groups, this object represents the eligibility policy for the group. Otherwise, it is null. Supports $expand.
// returns a PrivilegedAccessGroupEligibilityScheduleable when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetActivatedUsing()(PrivilegedAccessGroupEligibilityScheduleable) {
    return m.activatedUsing
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["activatedUsing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedUsing(val.(PrivilegedAccessGroupEligibilityScheduleable))
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
// GetGroup gets the group property value. References the group that is the scope of the membership or ownership assignment request through PIM for groups. Supports $expand and $select nested in $expand for select properties like id, displayName, and mail.
// returns a Groupable when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetGroup()(Groupable) {
    return m.group
}
// GetGroupId gets the groupId property value. The identifier of the group representing the scope of the membership or ownership assignment through PIM for groups. Required.
// returns a *string when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetGroupId()(*string) {
    return m.groupId
}
// GetPrincipal gets the principal property value. References the principal that's in the scope of this membership or ownership assignment request through the group that's governed by PIM. Supports $expand and $select nested in $expand for id only.
// returns a DirectoryObjectable when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. The identifier of the principal whose membership or ownership assignment to the group is managed through PIM for groups. Supports $filter (eq, ne).
// returns a *string when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetPrincipalId()(*string) {
    return m.principalId
}
// GetTargetSchedule gets the targetSchedule property value. Schedule created by this request. Supports $expand.
// returns a PrivilegedAccessGroupEligibilityScheduleable when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetTargetSchedule()(PrivilegedAccessGroupEligibilityScheduleable) {
    return m.targetSchedule
}
// GetTargetScheduleId gets the targetScheduleId property value. The identifier of the schedule that's created from the membership or ownership assignment request. Supports $filter (eq, ne).
// returns a *string when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) GetTargetScheduleId()(*string) {
    return m.targetScheduleId
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("activatedUsing", m.GetActivatedUsing())
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
// SetAccessId sets the accessId property value. The identifier of a membership or ownership assignment relationship to the group. Required. The possible values are: owner, member, unknownFutureValue.
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetAccessId(value *PrivilegedAccessGroupRelationships)() {
    m.accessId = value
}
// SetActivatedUsing sets the activatedUsing property value. When the request activates a membership or ownership assignment in PIM for groups, this object represents the eligibility policy for the group. Otherwise, it is null. Supports $expand.
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetActivatedUsing(value PrivilegedAccessGroupEligibilityScheduleable)() {
    m.activatedUsing = value
}
// SetGroup sets the group property value. References the group that is the scope of the membership or ownership assignment request through PIM for groups. Supports $expand and $select nested in $expand for select properties like id, displayName, and mail.
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetGroup(value Groupable)() {
    m.group = value
}
// SetGroupId sets the groupId property value. The identifier of the group representing the scope of the membership or ownership assignment through PIM for groups. Required.
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetGroupId(value *string)() {
    m.groupId = value
}
// SetPrincipal sets the principal property value. References the principal that's in the scope of this membership or ownership assignment request through the group that's governed by PIM. Supports $expand and $select nested in $expand for id only.
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. The identifier of the principal whose membership or ownership assignment to the group is managed through PIM for groups. Supports $filter (eq, ne).
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetPrincipalId(value *string)() {
    m.principalId = value
}
// SetTargetSchedule sets the targetSchedule property value. Schedule created by this request. Supports $expand.
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetTargetSchedule(value PrivilegedAccessGroupEligibilityScheduleable)() {
    m.targetSchedule = value
}
// SetTargetScheduleId sets the targetScheduleId property value. The identifier of the schedule that's created from the membership or ownership assignment request. Supports $filter (eq, ne).
func (m *PrivilegedAccessGroupAssignmentScheduleRequest) SetTargetScheduleId(value *string)() {
    m.targetScheduleId = value
}
type PrivilegedAccessGroupAssignmentScheduleRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrivilegedAccessScheduleRequestable
    GetAccessId()(*PrivilegedAccessGroupRelationships)
    GetActivatedUsing()(PrivilegedAccessGroupEligibilityScheduleable)
    GetGroup()(Groupable)
    GetGroupId()(*string)
    GetPrincipal()(DirectoryObjectable)
    GetPrincipalId()(*string)
    GetTargetSchedule()(PrivilegedAccessGroupEligibilityScheduleable)
    GetTargetScheduleId()(*string)
    SetAccessId(value *PrivilegedAccessGroupRelationships)()
    SetActivatedUsing(value PrivilegedAccessGroupEligibilityScheduleable)()
    SetGroup(value Groupable)()
    SetGroupId(value *string)()
    SetPrincipal(value DirectoryObjectable)()
    SetPrincipalId(value *string)()
    SetTargetSchedule(value PrivilegedAccessGroupEligibilityScheduleable)()
    SetTargetScheduleId(value *string)()
}
