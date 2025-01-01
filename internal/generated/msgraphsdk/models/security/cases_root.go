package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type CasesRoot struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The ediscoveryCases property
    ediscoveryCases []EdiscoveryCaseable
}
// NewCasesRoot instantiates a new CasesRoot and sets the default values.
func NewCasesRoot()(*CasesRoot) {
    m := &CasesRoot{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateCasesRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCasesRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesRoot(), nil
}
// GetEdiscoveryCases gets the ediscoveryCases property value. The ediscoveryCases property
// returns a []EdiscoveryCaseable when successful
func (m *CasesRoot) GetEdiscoveryCases()([]EdiscoveryCaseable) {
    return m.ediscoveryCases
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CasesRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["ediscoveryCases"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryCaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryCaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EdiscoveryCaseable)
                }
            }
            m.SetEdiscoveryCases(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CasesRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEdiscoveryCases() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEdiscoveryCases()))
        for i, v := range m.GetEdiscoveryCases() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("ediscoveryCases", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEdiscoveryCases sets the ediscoveryCases property value. The ediscoveryCases property
func (m *CasesRoot) SetEdiscoveryCases(value []EdiscoveryCaseable)() {
    m.ediscoveryCases = value
}
type CasesRootable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEdiscoveryCases()([]EdiscoveryCaseable)
    SetEdiscoveryCases(value []EdiscoveryCaseable)()
}
