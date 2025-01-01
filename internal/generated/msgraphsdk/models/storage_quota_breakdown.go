package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StorageQuotaBreakdown struct {
    Entity
    // The displayName property
    displayName *string
    // The manageWebUrl property
    manageWebUrl *string
    // The used property
    used *int64
}
// NewStorageQuotaBreakdown instantiates a new StorageQuotaBreakdown and sets the default values.
func NewStorageQuotaBreakdown()(*StorageQuotaBreakdown) {
    m := &StorageQuotaBreakdown{
        Entity: *NewEntity(),
    }
    return m
}
// CreateStorageQuotaBreakdownFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStorageQuotaBreakdownFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.serviceStorageQuotaBreakdown":
                        return NewServiceStorageQuotaBreakdown(), nil
                }
            }
        }
    }
    return NewStorageQuotaBreakdown(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *StorageQuotaBreakdown) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StorageQuotaBreakdown) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["manageWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManageWebUrl(val)
        }
        return nil
    }
    res["used"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsed(val)
        }
        return nil
    }
    return res
}
// GetManageWebUrl gets the manageWebUrl property value. The manageWebUrl property
// returns a *string when successful
func (m *StorageQuotaBreakdown) GetManageWebUrl()(*string) {
    return m.manageWebUrl
}
// GetUsed gets the used property value. The used property
// returns a *int64 when successful
func (m *StorageQuotaBreakdown) GetUsed()(*int64) {
    return m.used
}
// Serialize serializes information the current object
func (m *StorageQuotaBreakdown) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manageWebUrl", m.GetManageWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("used", m.GetUsed())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *StorageQuotaBreakdown) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetManageWebUrl sets the manageWebUrl property value. The manageWebUrl property
func (m *StorageQuotaBreakdown) SetManageWebUrl(value *string)() {
    m.manageWebUrl = value
}
// SetUsed sets the used property value. The used property
func (m *StorageQuotaBreakdown) SetUsed(value *int64)() {
    m.used = value
}
type StorageQuotaBreakdownable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetManageWebUrl()(*string)
    GetUsed()(*int64)
    SetDisplayName(value *string)()
    SetManageWebUrl(value *string)()
    SetUsed(value *int64)()
}
