package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type ExternalActivityResult struct {
    ExternalActivity
    // Error information that explains the failure to process an external activity.
    error ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable
}
// NewExternalActivityResult instantiates a new ExternalActivityResult and sets the default values.
func NewExternalActivityResult()(*ExternalActivityResult) {
    m := &ExternalActivityResult{
        ExternalActivity: *NewExternalActivity(),
    }
    return m
}
// CreateExternalActivityResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalActivityResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalActivityResult(), nil
}
// GetError gets the error property value. Error information that explains the failure to process an external activity.
// returns a PublicErrorable when successful
func (m *ExternalActivityResult) GetError()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalActivityResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExternalActivity.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ExternalActivityResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExternalActivity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. Error information that explains the failure to process an external activity.
func (m *ExternalActivityResult) SetError(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable)() {
    m.error = value
}
type ExternalActivityResultable interface {
    ExternalActivityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable)
    SetError(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PublicErrorable)()
}
