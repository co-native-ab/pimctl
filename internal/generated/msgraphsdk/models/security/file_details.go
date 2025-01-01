package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FileDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the file.
    fileName *string
    // The file path (location) of the file instance.
    filePath *string
    // The publisher of the file.
    filePublisher *string
    // The size of the file in bytes.
    fileSize *int64
    // The certificate authority (CA) that issued the certificate.
    issuer *string
    // The OdataType property
    odataType *string
    // The Sha1 cryptographic hash of the file content.
    sha1 *string
    // The Sha256 cryptographic hash of the file content.
    sha256 *string
    // The signer of the signed file.
    signer *string
}
// NewFileDetails instantiates a new FileDetails and sets the default values.
func NewFileDetails()(*FileDetails) {
    m := &FileDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FileDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FileDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["filePath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePath(val)
        }
        return nil
    }
    res["filePublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePublisher(val)
        }
        return nil
    }
    res["fileSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSize(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["sha1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha1(val)
        }
        return nil
    }
    res["sha256"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha256(val)
        }
        return nil
    }
    res["signer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigner(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. The name of the file.
// returns a *string when successful
func (m *FileDetails) GetFileName()(*string) {
    return m.fileName
}
// GetFilePath gets the filePath property value. The file path (location) of the file instance.
// returns a *string when successful
func (m *FileDetails) GetFilePath()(*string) {
    return m.filePath
}
// GetFilePublisher gets the filePublisher property value. The publisher of the file.
// returns a *string when successful
func (m *FileDetails) GetFilePublisher()(*string) {
    return m.filePublisher
}
// GetFileSize gets the fileSize property value. The size of the file in bytes.
// returns a *int64 when successful
func (m *FileDetails) GetFileSize()(*int64) {
    return m.fileSize
}
// GetIssuer gets the issuer property value. The certificate authority (CA) that issued the certificate.
// returns a *string when successful
func (m *FileDetails) GetIssuer()(*string) {
    return m.issuer
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *FileDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetSha1 gets the sha1 property value. The Sha1 cryptographic hash of the file content.
// returns a *string when successful
func (m *FileDetails) GetSha1()(*string) {
    return m.sha1
}
// GetSha256 gets the sha256 property value. The Sha256 cryptographic hash of the file content.
// returns a *string when successful
func (m *FileDetails) GetSha256()(*string) {
    return m.sha256
}
// GetSigner gets the signer property value. The signer of the signed file.
// returns a *string when successful
func (m *FileDetails) GetSigner()(*string) {
    return m.signer
}
// Serialize serializes information the current object
func (m *FileDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filePath", m.GetFilePath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filePublisher", m.GetFilePublisher())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("fileSize", m.GetFileSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha1", m.GetSha1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha256", m.GetSha256())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signer", m.GetSigner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFileName sets the fileName property value. The name of the file.
func (m *FileDetails) SetFileName(value *string)() {
    m.fileName = value
}
// SetFilePath sets the filePath property value. The file path (location) of the file instance.
func (m *FileDetails) SetFilePath(value *string)() {
    m.filePath = value
}
// SetFilePublisher sets the filePublisher property value. The publisher of the file.
func (m *FileDetails) SetFilePublisher(value *string)() {
    m.filePublisher = value
}
// SetFileSize sets the fileSize property value. The size of the file in bytes.
func (m *FileDetails) SetFileSize(value *int64)() {
    m.fileSize = value
}
// SetIssuer sets the issuer property value. The certificate authority (CA) that issued the certificate.
func (m *FileDetails) SetIssuer(value *string)() {
    m.issuer = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FileDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSha1 sets the sha1 property value. The Sha1 cryptographic hash of the file content.
func (m *FileDetails) SetSha1(value *string)() {
    m.sha1 = value
}
// SetSha256 sets the sha256 property value. The Sha256 cryptographic hash of the file content.
func (m *FileDetails) SetSha256(value *string)() {
    m.sha256 = value
}
// SetSigner sets the signer property value. The signer of the signed file.
func (m *FileDetails) SetSigner(value *string)() {
    m.signer = value
}
type FileDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFileName()(*string)
    GetFilePath()(*string)
    GetFilePublisher()(*string)
    GetFileSize()(*int64)
    GetIssuer()(*string)
    GetOdataType()(*string)
    GetSha1()(*string)
    GetSha256()(*string)
    GetSigner()(*string)
    SetFileName(value *string)()
    SetFilePath(value *string)()
    SetFilePublisher(value *string)()
    SetFileSize(value *int64)()
    SetIssuer(value *string)()
    SetOdataType(value *string)()
    SetSha1(value *string)()
    SetSha256(value *string)()
    SetSigner(value *string)()
}
