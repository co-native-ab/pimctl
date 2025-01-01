package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Community represents a community in Viva Engage that is a central place for conversations,files, events, and updates for people sharing a common interest or goal.
type Community struct {
    Entity
    // The description of the community. The maximum length is 1,024 characters.
    description *string
    // The name of the community. The maximum length is 255 characters.
    displayName *string
    // The Microsoft 365 group that manages the membership of this community.
    group Groupable
    // The ID of the Microsoft 365 group that manages the membership of this community.
    groupId *string
    // The admins of the community. Limited to 100 users. If this property isn't specified when you create the community, the calling user is automatically assigned as the community owner.
    owners []Userable
    // Types of communityPrivacy.
    privacy *CommunityPrivacy
}
// NewCommunity instantiates a new Community and sets the default values.
func NewCommunity()(*Community) {
    m := &Community{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCommunityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommunityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunity(), nil
}
// GetDescription gets the description property value. The description of the community. The maximum length is 1,024 characters.
// returns a *string when successful
func (m *Community) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of the community. The maximum length is 255 characters.
// returns a *string when successful
func (m *Community) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Community) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["owners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetOwners(res)
        }
        return nil
    }
    res["privacy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCommunityPrivacy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacy(val.(*CommunityPrivacy))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The Microsoft 365 group that manages the membership of this community.
// returns a Groupable when successful
func (m *Community) GetGroup()(Groupable) {
    return m.group
}
// GetGroupId gets the groupId property value. The ID of the Microsoft 365 group that manages the membership of this community.
// returns a *string when successful
func (m *Community) GetGroupId()(*string) {
    return m.groupId
}
// GetOwners gets the owners property value. The admins of the community. Limited to 100 users. If this property isn't specified when you create the community, the calling user is automatically assigned as the community owner.
// returns a []Userable when successful
func (m *Community) GetOwners()([]Userable) {
    return m.owners
}
// GetPrivacy gets the privacy property value. Types of communityPrivacy.
// returns a *CommunityPrivacy when successful
func (m *Community) GetPrivacy()(*CommunityPrivacy) {
    return m.privacy
}
// Serialize serializes information the current object
func (m *Community) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetPrivacy() != nil {
        cast := (*m.GetPrivacy()).String()
        err = writer.WriteStringValue("privacy", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description of the community. The maximum length is 1,024 characters.
func (m *Community) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of the community. The maximum length is 255 characters.
func (m *Community) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGroup sets the group property value. The Microsoft 365 group that manages the membership of this community.
func (m *Community) SetGroup(value Groupable)() {
    m.group = value
}
// SetGroupId sets the groupId property value. The ID of the Microsoft 365 group that manages the membership of this community.
func (m *Community) SetGroupId(value *string)() {
    m.groupId = value
}
// SetOwners sets the owners property value. The admins of the community. Limited to 100 users. If this property isn't specified when you create the community, the calling user is automatically assigned as the community owner.
func (m *Community) SetOwners(value []Userable)() {
    m.owners = value
}
// SetPrivacy sets the privacy property value. Types of communityPrivacy.
func (m *Community) SetPrivacy(value *CommunityPrivacy)() {
    m.privacy = value
}
type Communityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetGroup()(Groupable)
    GetGroupId()(*string)
    GetOwners()([]Userable)
    GetPrivacy()(*CommunityPrivacy)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetGroup(value Groupable)()
    SetGroupId(value *string)()
    SetOwners(value []Userable)()
    SetPrivacy(value *CommunityPrivacy)()
}
