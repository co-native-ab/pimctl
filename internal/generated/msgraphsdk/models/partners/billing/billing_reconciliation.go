package billing

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type BillingReconciliation struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The billed property
    billed BilledReconciliationable
}
// NewBillingReconciliation instantiates a new BillingReconciliation and sets the default values.
func NewBillingReconciliation()(*BillingReconciliation) {
    m := &BillingReconciliation{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateBillingReconciliationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBillingReconciliationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBillingReconciliation(), nil
}
// GetBilled gets the billed property value. The billed property
// returns a BilledReconciliationable when successful
func (m *BillingReconciliation) GetBilled()(BilledReconciliationable) {
    return m.billed
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BillingReconciliation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["billed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBilledReconciliationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBilled(val.(BilledReconciliationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *BillingReconciliation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetBilled sets the billed property value. The billed property
func (m *BillingReconciliation) SetBilled(value BilledReconciliationable)() {
    m.billed = value
}
type BillingReconciliationable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBilled()(BilledReconciliationable)
    SetBilled(value BilledReconciliationable)()
}
