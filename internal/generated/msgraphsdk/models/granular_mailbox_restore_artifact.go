package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GranularMailboxRestoreArtifact struct {
    MailboxRestoreArtifact
    // .
    artifactCount *int32
    // .
    searchResponseId *string
}
// NewGranularMailboxRestoreArtifact instantiates a new GranularMailboxRestoreArtifact and sets the default values.
func NewGranularMailboxRestoreArtifact()(*GranularMailboxRestoreArtifact) {
    m := &GranularMailboxRestoreArtifact{
        MailboxRestoreArtifact: *NewMailboxRestoreArtifact(),
    }
    return m
}
// CreateGranularMailboxRestoreArtifactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGranularMailboxRestoreArtifactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGranularMailboxRestoreArtifact(), nil
}
// GetArtifactCount gets the artifactCount property value. .
// returns a *int32 when successful
func (m *GranularMailboxRestoreArtifact) GetArtifactCount()(*int32) {
    return m.artifactCount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GranularMailboxRestoreArtifact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MailboxRestoreArtifact.GetFieldDeserializers()
    res["artifactCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArtifactCount(val)
        }
        return nil
    }
    res["searchResponseId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchResponseId(val)
        }
        return nil
    }
    return res
}
// GetSearchResponseId gets the searchResponseId property value. .
// returns a *string when successful
func (m *GranularMailboxRestoreArtifact) GetSearchResponseId()(*string) {
    return m.searchResponseId
}
// Serialize serializes information the current object
func (m *GranularMailboxRestoreArtifact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MailboxRestoreArtifact.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("artifactCount", m.GetArtifactCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("searchResponseId", m.GetSearchResponseId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArtifactCount sets the artifactCount property value. .
func (m *GranularMailboxRestoreArtifact) SetArtifactCount(value *int32)() {
    m.artifactCount = value
}
// SetSearchResponseId sets the searchResponseId property value. .
func (m *GranularMailboxRestoreArtifact) SetSearchResponseId(value *string)() {
    m.searchResponseId = value
}
type GranularMailboxRestoreArtifactable interface {
    MailboxRestoreArtifactable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArtifactCount()(*int32)
    GetSearchResponseId()(*string)
    SetArtifactCount(value *int32)()
    SetSearchResponseId(value *string)()
}
