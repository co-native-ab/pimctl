package billing

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type AzureUsage struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The billed property
    billed BilledUsageable
    // The unbilled property
    unbilled UnbilledUsageable
}
// NewAzureUsage instantiates a new AzureUsage and sets the default values.
func NewAzureUsage()(*AzureUsage) {
    m := &AzureUsage{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateAzureUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAzureUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureUsage(), nil
}
// GetBilled gets the billed property value. The billed property
// returns a BilledUsageable when successful
func (m *AzureUsage) GetBilled()(BilledUsageable) {
    return m.billed
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AzureUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["billed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBilledUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBilled(val.(BilledUsageable))
        }
        return nil
    }
    res["unbilled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnbilledUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnbilled(val.(UnbilledUsageable))
        }
        return nil
    }
    return res
}
// GetUnbilled gets the unbilled property value. The unbilled property
// returns a UnbilledUsageable when successful
func (m *AzureUsage) GetUnbilled()(UnbilledUsageable) {
    return m.unbilled
}
// Serialize serializes information the current object
func (m *AzureUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("billed", m.GetBilled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("unbilled", m.GetUnbilled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBilled sets the billed property value. The billed property
func (m *AzureUsage) SetBilled(value BilledUsageable)() {
    m.billed = value
}
// SetUnbilled sets the unbilled property value. The unbilled property
func (m *AzureUsage) SetUnbilled(value UnbilledUsageable)() {
    m.unbilled = value
}
type AzureUsageable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBilled()(BilledUsageable)
    GetUnbilled()(UnbilledUsageable)
    SetBilled(value BilledUsageable)()
    SetUnbilled(value UnbilledUsageable)()
}
