package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserEvidence struct {
    AlertEvidence
    // The stream property
    stream Streamable
    // The user account details.
    userAccount UserAccountable
}
// NewUserEvidence instantiates a new UserEvidence and sets the default values.
func NewUserEvidence()(*UserEvidence) {
    m := &UserEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.userEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUserEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["stream"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStreamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStream(val.(Streamable))
        }
        return nil
    }
    res["userAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccount(val.(UserAccountable))
        }
        return nil
    }
    return res
}
// GetStream gets the stream property value. The stream property
// returns a Streamable when successful
func (m *UserEvidence) GetStream()(Streamable) {
    return m.stream
}
// GetUserAccount gets the userAccount property value. The user account details.
// returns a UserAccountable when successful
func (m *UserEvidence) GetUserAccount()(UserAccountable) {
    return m.userAccount
}
// Serialize serializes information the current object
func (m *UserEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("stream", m.GetStream())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userAccount", m.GetUserAccount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStream sets the stream property value. The stream property
func (m *UserEvidence) SetStream(value Streamable)() {
    m.stream = value
}
// SetUserAccount sets the userAccount property value. The user account details.
func (m *UserEvidence) SetUserAccount(value UserAccountable)() {
    m.userAccount = value
}
type UserEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStream()(Streamable)
    GetUserAccount()(UserAccountable)
    SetStream(value Streamable)()
    SetUserAccount(value UserAccountable)()
}
