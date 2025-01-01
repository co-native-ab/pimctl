package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UnifiedStorageQuota struct {
    Entity
    // The deleted property
    deleted *int64
    // The manageWebUrl property
    manageWebUrl *string
    // The remaining property
    remaining *int64
    // The services property
    services []ServiceStorageQuotaBreakdownable
    // The state property
    state *string
    // The total property
    total *int64
    // The used property
    used *int64
}
// NewUnifiedStorageQuota instantiates a new UnifiedStorageQuota and sets the default values.
func NewUnifiedStorageQuota()(*UnifiedStorageQuota) {
    m := &UnifiedStorageQuota{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedStorageQuotaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUnifiedStorageQuotaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedStorageQuota(), nil
}
// GetDeleted gets the deleted property value. The deleted property
// returns a *int64 when successful
func (m *UnifiedStorageQuota) GetDeleted()(*int64) {
    return m.deleted
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UnifiedStorageQuota) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleted(val)
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
    res["remaining"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemaining(val)
        }
        return nil
    }
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceStorageQuotaBreakdownFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceStorageQuotaBreakdownable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServiceStorageQuotaBreakdownable)
                }
            }
            m.SetServices(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
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
func (m *UnifiedStorageQuota) GetManageWebUrl()(*string) {
    return m.manageWebUrl
}
// GetRemaining gets the remaining property value. The remaining property
// returns a *int64 when successful
func (m *UnifiedStorageQuota) GetRemaining()(*int64) {
    return m.remaining
}
// GetServices gets the services property value. The services property
// returns a []ServiceStorageQuotaBreakdownable when successful
func (m *UnifiedStorageQuota) GetServices()([]ServiceStorageQuotaBreakdownable) {
    return m.services
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *UnifiedStorageQuota) GetState()(*string) {
    return m.state
}
// GetTotal gets the total property value. The total property
// returns a *int64 when successful
func (m *UnifiedStorageQuota) GetTotal()(*int64) {
    return m.total
}
// GetUsed gets the used property value. The used property
// returns a *int64 when successful
func (m *UnifiedStorageQuota) GetUsed()(*int64) {
    return m.used
}
// Serialize serializes information the current object
func (m *UnifiedStorageQuota) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("deleted", m.GetDeleted())
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
        err = writer.WriteInt64Value("remaining", m.GetRemaining())
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServices()))
        for i, v := range m.GetServices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("services", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("total", m.GetTotal())
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
// SetDeleted sets the deleted property value. The deleted property
func (m *UnifiedStorageQuota) SetDeleted(value *int64)() {
    m.deleted = value
}
// SetManageWebUrl sets the manageWebUrl property value. The manageWebUrl property
func (m *UnifiedStorageQuota) SetManageWebUrl(value *string)() {
    m.manageWebUrl = value
}
// SetRemaining sets the remaining property value. The remaining property
func (m *UnifiedStorageQuota) SetRemaining(value *int64)() {
    m.remaining = value
}
// SetServices sets the services property value. The services property
func (m *UnifiedStorageQuota) SetServices(value []ServiceStorageQuotaBreakdownable)() {
    m.services = value
}
// SetState sets the state property value. The state property
func (m *UnifiedStorageQuota) SetState(value *string)() {
    m.state = value
}
// SetTotal sets the total property value. The total property
func (m *UnifiedStorageQuota) SetTotal(value *int64)() {
    m.total = value
}
// SetUsed sets the used property value. The used property
func (m *UnifiedStorageQuota) SetUsed(value *int64)() {
    m.used = value
}
type UnifiedStorageQuotaable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeleted()(*int64)
    GetManageWebUrl()(*string)
    GetRemaining()(*int64)
    GetServices()([]ServiceStorageQuotaBreakdownable)
    GetState()(*string)
    GetTotal()(*int64)
    GetUsed()(*int64)
    SetDeleted(value *int64)()
    SetManageWebUrl(value *string)()
    SetRemaining(value *int64)()
    SetServices(value []ServiceStorageQuotaBreakdownable)()
    SetState(value *string)()
    SetTotal(value *int64)()
    SetUsed(value *int64)()
}
