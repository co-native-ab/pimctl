package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BlobEvidence struct {
    AlertEvidence
    // The container which the blob belongs to.
    blobContainer BlobContainerEvidenceable
    // The Etag associated with this blob.
    etag *string
    // The file hashes associated with this blob.
    fileHashes []FileHashable
    // The name of the blob.
    name *string
    // The full URL representation of the blob.
    url *string
}
// NewBlobEvidence instantiates a new BlobEvidence and sets the default values.
func NewBlobEvidence()(*BlobEvidence) {
    m := &BlobEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.blobEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateBlobEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlobEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlobEvidence(), nil
}
// GetBlobContainer gets the blobContainer property value. The container which the blob belongs to.
// returns a BlobContainerEvidenceable when successful
func (m *BlobEvidence) GetBlobContainer()(BlobContainerEvidenceable) {
    return m.blobContainer
}
// GetEtag gets the etag property value. The Etag associated with this blob.
// returns a *string when successful
func (m *BlobEvidence) GetEtag()(*string) {
    return m.etag
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlobEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["blobContainer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBlobContainerEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlobContainer(val.(BlobContainerEvidenceable))
        }
        return nil
    }
    res["etag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEtag(val)
        }
        return nil
    }
    res["fileHashes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFileHashFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileHashable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FileHashable)
                }
            }
            m.SetFileHashes(res)
        }
        return nil
    }
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
// GetFileHashes gets the fileHashes property value. The file hashes associated with this blob.
// returns a []FileHashable when successful
func (m *BlobEvidence) GetFileHashes()([]FileHashable) {
    return m.fileHashes
}
// GetName gets the name property value. The name of the blob.
// returns a *string when successful
func (m *BlobEvidence) GetName()(*string) {
    return m.name
}
// GetUrl gets the url property value. The full URL representation of the blob.
// returns a *string when successful
func (m *BlobEvidence) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *BlobEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("blobContainer", m.GetBlobContainer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("etag", m.GetEtag())
        if err != nil {
            return err
        }
    }
    if m.GetFileHashes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFileHashes()))
        for i, v := range m.GetFileHashes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("fileHashes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
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
// SetBlobContainer sets the blobContainer property value. The container which the blob belongs to.
func (m *BlobEvidence) SetBlobContainer(value BlobContainerEvidenceable)() {
    m.blobContainer = value
}
// SetEtag sets the etag property value. The Etag associated with this blob.
func (m *BlobEvidence) SetEtag(value *string)() {
    m.etag = value
}
// SetFileHashes sets the fileHashes property value. The file hashes associated with this blob.
func (m *BlobEvidence) SetFileHashes(value []FileHashable)() {
    m.fileHashes = value
}
// SetName sets the name property value. The name of the blob.
func (m *BlobEvidence) SetName(value *string)() {
    m.name = value
}
// SetUrl sets the url property value. The full URL representation of the blob.
func (m *BlobEvidence) SetUrl(value *string)() {
    m.url = value
}
type BlobEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlobContainer()(BlobContainerEvidenceable)
    GetEtag()(*string)
    GetFileHashes()([]FileHashable)
    GetName()(*string)
    GetUrl()(*string)
    SetBlobContainer(value BlobContainerEvidenceable)()
    SetEtag(value *string)()
    SetFileHashes(value []FileHashable)()
    SetName(value *string)()
    SetUrl(value *string)()
}
