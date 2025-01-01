package billing

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type Billing struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // Represents metadata for the exported data.
    manifests []Manifestable
    // Represents an operation to export the billing data of a partner.
    operations []Operationable
    // The reconciliation property
    reconciliation BillingReconciliationable
    // The usage property
    usage AzureUsageable
}
// NewBilling instantiates a new Billing and sets the default values.
func NewBilling()(*Billing) {
    m := &Billing{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateBillingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBillingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBilling(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Billing) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["manifests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManifestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Manifestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Manifestable)
                }
            }
            m.SetManifests(res)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Operationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Operationable)
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["reconciliation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBillingReconciliationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReconciliation(val.(BillingReconciliationable))
        }
        return nil
    }
    res["usage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAzureUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsage(val.(AzureUsageable))
        }
        return nil
    }
    return res
}
// GetManifests gets the manifests property value. Represents metadata for the exported data.
// returns a []Manifestable when successful
func (m *Billing) GetManifests()([]Manifestable) {
    return m.manifests
}
// GetOperations gets the operations property value. Represents an operation to export the billing data of a partner.
// returns a []Operationable when successful
func (m *Billing) GetOperations()([]Operationable) {
    return m.operations
}
// GetReconciliation gets the reconciliation property value. The reconciliation property
// returns a BillingReconciliationable when successful
func (m *Billing) GetReconciliation()(BillingReconciliationable) {
    return m.reconciliation
}
// GetUsage gets the usage property value. The usage property
// returns a AzureUsageable when successful
func (m *Billing) GetUsage()(AzureUsageable) {
    return m.usage
}
// Serialize serializes information the current object
func (m *Billing) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetManifests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManifests()))
        for i, v := range m.GetManifests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("manifests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reconciliation", m.GetReconciliation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("usage", m.GetUsage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetManifests sets the manifests property value. Represents metadata for the exported data.
func (m *Billing) SetManifests(value []Manifestable)() {
    m.manifests = value
}
// SetOperations sets the operations property value. Represents an operation to export the billing data of a partner.
func (m *Billing) SetOperations(value []Operationable)() {
    m.operations = value
}
// SetReconciliation sets the reconciliation property value. The reconciliation property
func (m *Billing) SetReconciliation(value BillingReconciliationable)() {
    m.reconciliation = value
}
// SetUsage sets the usage property value. The usage property
func (m *Billing) SetUsage(value AzureUsageable)() {
    m.usage = value
}
type Billingable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManifests()([]Manifestable)
    GetOperations()([]Operationable)
    GetReconciliation()(BillingReconciliationable)
    GetUsage()(AzureUsageable)
    SetManifests(value []Manifestable)()
    SetOperations(value []Operationable)()
    SetReconciliation(value BillingReconciliationable)()
    SetUsage(value AzureUsageable)()
}
