package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BlobContainerEvidence struct {
    AlertEvidence
    // The name of the blob container.
    name *string
    // The storage which the blob container belongs to.
    storageResource AzureResourceEvidenceable
    // The full URL representation of the blob container.
    url *string
}
// NewBlobContainerEvidence instantiates a new BlobContainerEvidence and sets the default values.
func NewBlobContainerEvidence()(*BlobContainerEvidence) {
    m := &BlobContainerEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.blobContainerEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateBlobContainerEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlobContainerEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlobContainerEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlobContainerEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["storageResource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAzureResourceEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageResource(val.(AzureResourceEvidenceable))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the blob container.
// returns a *string when successful
func (m *BlobContainerEvidence) GetName()(*string) {
    return m.name
}
// GetStorageResource gets the storageResource property value. The storage which the blob container belongs to.
// returns a AzureResourceEvidenceable when successful
func (m *BlobContainerEvidence) GetStorageResource()(AzureResourceEvidenceable) {
    return m.storageResource
}
// GetUrl gets the url property value. The full URL representation of the blob container.
// returns a *string when successful
func (m *BlobContainerEvidence) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *BlobContainerEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("storageResource", m.GetStorageResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name of the blob container.
func (m *BlobContainerEvidence) SetName(value *string)() {
    m.name = value
}
// SetStorageResource sets the storageResource property value. The storage which the blob container belongs to.
func (m *BlobContainerEvidence) SetStorageResource(value AzureResourceEvidenceable)() {
    m.storageResource = value
}
// SetUrl sets the url property value. The full URL representation of the blob container.
func (m *BlobContainerEvidence) SetUrl(value *string)() {
    m.url = value
}
type BlobContainerEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetStorageResource()(AzureResourceEvidenceable)
    GetUrl()(*string)
    SetName(value *string)()
    SetStorageResource(value AzureResourceEvidenceable)()
    SetUrl(value *string)()
}
