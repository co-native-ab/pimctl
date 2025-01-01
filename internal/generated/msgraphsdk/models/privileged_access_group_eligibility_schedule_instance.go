package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PrivilegedAccessGroupEligibilityScheduleInstance struct {
    PrivilegedAccessScheduleInstance
    // The identifier of the membership or ownership eligibility relationship to the group. Required. The possible values are: owner, member. Supports $filter (eq).
    accessId *PrivilegedAccessGroupRelationships
    // The identifier of the privilegedAccessGroupEligibilitySchedule from which this instance was created. Required. Supports $filter (eq, ne).
    eligibilityScheduleId *string
    // References the group that is the scope of the membership or ownership eligibility through PIM for groups. Supports $expand.
    group Groupable
    // The identifier of the group representing the scope of the membership or ownership eligibility through PIM for groups. Required. Supports $filter (eq).
    groupId *string
    // Indicates whether the assignment is derived from a group assignment. It can further imply whether the calling principal can manage the assignment schedule. Required. The possible values are: direct, group, unknownFutureValue. Supports $filter (eq).
    memberType *PrivilegedAccessGroupMemberType
    // References the principal that's in the scope of the membership or ownership eligibility request through the group that's governed by PIM. Supports $expand.
    principal DirectoryObjectable
    // The identifier of the principal whose membership or ownership eligibility to the group is managed through PIM for groups. Required. Supports $filter (eq).
    principalId *string
}
// NewPrivilegedAccessGroupEligibilityScheduleInstance instantiates a new PrivilegedAccessGroupEligibilityScheduleInstance and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleInstance()(*PrivilegedAccessGroupEligibilityScheduleInstance) {
    m := &PrivilegedAccessGroupEligibilityScheduleInstance{
        PrivilegedAccessScheduleInstance: *NewPrivilegedAccessScheduleInstance(),
    }
    odataTypeValue := "#microsoft.graph.privilegedAccessGroupEligibilityScheduleInstance"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePrivilegedAccessGroupEligibilityScheduleInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivilegedAccessGroupEligibilityScheduleInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupEligibilityScheduleInstance(), nil
}
// GetAccessId gets the accessId property value. The identifier of the membership or ownership eligibility relationship to the group. Required. The possible values are: owner, member. Supports $filter (eq).
// returns a *PrivilegedAccessGroupRelationships when successful
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetAccessId()(*PrivilegedAccessGroupRelationships) {
    return m.accessId
}
// GetEligibilityScheduleId gets the eligibilityScheduleId property value. The identifier of the privilegedAccessGroupEligibilitySchedule from which this instance was created. Required. Supports $filter (eq, ne).
// returns a *string when successful
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetEligibilityScheduleId()(*string) {
    return m.eligibilityScheduleId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrivilegedAccessScheduleInstance.GetFieldDeserializers()
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
    res["eligibilityScheduleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEligibilityScheduleId(val)
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
    res["memberType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivilegedAccessGroupMemberType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberType(val.(*PrivilegedAccessGroupMemberType))
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
    return res
}
// GetGroup gets the group property value. References the group that is the scope of the membership or ownership eligibility through PIM for groups. Supports $expand.
// returns a Groupable when successful
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetGroup()(Groupable) {
    return m.group
}
// GetGroupId gets the groupId property value. The identifier of the group representing the scope of the membership or ownership eligibility through PIM for groups. Required. Supports $filter (eq).
// returns a *string when successful
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetGroupId()(*string) {
    return m.groupId
}
// GetMemberType gets the memberType property value. Indicates whether the assignment is derived from a group assignment. It can further imply whether the calling principal can manage the assignment schedule. Required. The possible values are: direct, group, unknownFutureValue. Supports $filter (eq).
// returns a *PrivilegedAccessGroupMemberType when successful
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetMemberType()(*PrivilegedAccessGroupMemberType) {
    return m.memberType
}
// GetPrincipal gets the principal property value. References the principal that's in the scope of the membership or ownership eligibility request through the group that's governed by PIM. Supports $expand.
// returns a DirectoryObjectable when successful
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. The identifier of the principal whose membership or ownership eligibility to the group is managed through PIM for groups. Required. Supports $filter (eq).
// returns a *string when successful
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) GetPrincipalId()(*string) {
    return m.principalId
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrivilegedAccessScheduleInstance.Serialize(writer)
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
        err = writer.WriteStringValue("eligibilityScheduleId", m.GetEligibilityScheduleId())
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
    if m.GetMemberType() != nil {
        cast := (*m.GetMemberType()).String()
        err = writer.WriteStringValue("memberType", &cast)
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
    return nil
}
// SetAccessId sets the accessId property value. The identifier of the membership or ownership eligibility relationship to the group. Required. The possible values are: owner, member. Supports $filter (eq).
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetAccessId(value *PrivilegedAccessGroupRelationships)() {
    m.accessId = value
}
// SetEligibilityScheduleId sets the eligibilityScheduleId property value. The identifier of the privilegedAccessGroupEligibilitySchedule from which this instance was created. Required. Supports $filter (eq, ne).
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetEligibilityScheduleId(value *string)() {
    m.eligibilityScheduleId = value
}
// SetGroup sets the group property value. References the group that is the scope of the membership or ownership eligibility through PIM for groups. Supports $expand.
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetGroup(value Groupable)() {
    m.group = value
}
// SetGroupId sets the groupId property value. The identifier of the group representing the scope of the membership or ownership eligibility through PIM for groups. Required. Supports $filter (eq).
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetGroupId(value *string)() {
    m.groupId = value
}
// SetMemberType sets the memberType property value. Indicates whether the assignment is derived from a group assignment. It can further imply whether the calling principal can manage the assignment schedule. Required. The possible values are: direct, group, unknownFutureValue. Supports $filter (eq).
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetMemberType(value *PrivilegedAccessGroupMemberType)() {
    m.memberType = value
}
// SetPrincipal sets the principal property value. References the principal that's in the scope of the membership or ownership eligibility request through the group that's governed by PIM. Supports $expand.
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. The identifier of the principal whose membership or ownership eligibility to the group is managed through PIM for groups. Required. Supports $filter (eq).
func (m *PrivilegedAccessGroupEligibilityScheduleInstance) SetPrincipalId(value *string)() {
    m.principalId = value
}
type PrivilegedAccessGroupEligibilityScheduleInstanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrivilegedAccessScheduleInstanceable
    GetAccessId()(*PrivilegedAccessGroupRelationships)
    GetEligibilityScheduleId()(*string)
    GetGroup()(Groupable)
    GetGroupId()(*string)
    GetMemberType()(*PrivilegedAccessGroupMemberType)
    GetPrincipal()(DirectoryObjectable)
    GetPrincipalId()(*string)
    SetAccessId(value *PrivilegedAccessGroupRelationships)()
    SetEligibilityScheduleId(value *string)()
    SetGroup(value Groupable)()
    SetGroupId(value *string)()
    SetMemberType(value *PrivilegedAccessGroupMemberType)()
    SetPrincipal(value DirectoryObjectable)()
    SetPrincipalId(value *string)()
}
