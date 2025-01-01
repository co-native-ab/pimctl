package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type TriggersRoot struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The retentionEvents property
    retentionEvents []RetentionEventable
}
// NewTriggersRoot instantiates a new TriggersRoot and sets the default values.
func NewTriggersRoot()(*TriggersRoot) {
    m := &TriggersRoot{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateTriggersRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTriggersRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTriggersRoot(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TriggersRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["retentionEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRetentionEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RetentionEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RetentionEventable)
                }
            }
            m.SetRetentionEvents(res)
        }
        return nil
    }
    return res
}
// GetRetentionEvents gets the retentionEvents property value. The retentionEvents property
// returns a []RetentionEventable when successful
func (m *TriggersRoot) GetRetentionEvents()([]RetentionEventable) {
    return m.retentionEvents
}
// Serialize serializes information the current object
func (m *TriggersRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRetentionEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRetentionEvents()))
        for i, v := range m.GetRetentionEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("retentionEvents", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRetentionEvents sets the retentionEvents property value. The retentionEvents property
func (m *TriggersRoot) SetRetentionEvents(value []RetentionEventable)() {
    m.retentionEvents = value
}
type TriggersRootable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRetentionEvents()([]RetentionEventable)
    SetRetentionEvents(value []RetentionEventable)()
}
