package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudLogonRequestEvidence struct {
    AlertEvidence
    // The unique identifier for the sign-in request.
    requestId *string
}
// NewCloudLogonRequestEvidence instantiates a new CloudLogonRequestEvidence and sets the default values.
func NewCloudLogonRequestEvidence()(*CloudLogonRequestEvidence) {
    m := &CloudLogonRequestEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.cloudLogonRequestEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudLogonRequestEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudLogonRequestEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudLogonRequestEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudLogonRequestEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["requestId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestId(val)
        }
        return nil
    }
    return res
}
// GetRequestId gets the requestId property value. The unique identifier for the sign-in request.
// returns a *string when successful
func (m *CloudLogonRequestEvidence) GetRequestId()(*string) {
    return m.requestId
}
// Serialize serializes information the current object
func (m *CloudLogonRequestEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("requestId", m.GetRequestId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRequestId sets the requestId property value. The unique identifier for the sign-in request.
func (m *CloudLogonRequestEvidence) SetRequestId(value *string)() {
    m.requestId = value
}
type CloudLogonRequestEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRequestId()(*string)
    SetRequestId(value *string)()
}
