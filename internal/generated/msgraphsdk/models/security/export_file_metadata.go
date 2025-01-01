package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExportFileMetadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The downloadUrl property
    downloadUrl *string
    // The fileName property
    fileName *string
    // The OdataType property
    odataType *string
    // The size property
    size *int64
}
// NewExportFileMetadata instantiates a new ExportFileMetadata and sets the default values.
func NewExportFileMetadata()(*ExportFileMetadata) {
    m := &ExportFileMetadata{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExportFileMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExportFileMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExportFileMetadata(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExportFileMetadata) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDownloadUrl gets the downloadUrl property value. The downloadUrl property
// returns a *string when successful
func (m *ExportFileMetadata) GetDownloadUrl()(*string) {
    return m.downloadUrl
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExportFileMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["downloadUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadUrl(val)
        }
        return nil
    }
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
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. The fileName property
// returns a *string when successful
func (m *ExportFileMetadata) GetFileName()(*string) {
    return m.fileName
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ExportFileMetadata) GetOdataType()(*string) {
    return m.odataType
}
// GetSize gets the size property value. The size property
// returns a *int64 when successful
func (m *ExportFileMetadata) GetSize()(*int64) {
    return m.size
}
// Serialize serializes information the current object
func (m *ExportFileMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("downloadUrl", m.GetDownloadUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
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
        err := writer.WriteInt64Value("size", m.GetSize())
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
func (m *ExportFileMetadata) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDownloadUrl sets the downloadUrl property value. The downloadUrl property
func (m *ExportFileMetadata) SetDownloadUrl(value *string)() {
    m.downloadUrl = value
}
// SetFileName sets the fileName property value. The fileName property
func (m *ExportFileMetadata) SetFileName(value *string)() {
    m.fileName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExportFileMetadata) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSize sets the size property value. The size property
func (m *ExportFileMetadata) SetSize(value *int64)() {
    m.size = value
}
type ExportFileMetadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDownloadUrl()(*string)
    GetFileName()(*string)
    GetOdataType()(*string)
    GetSize()(*int64)
    SetDownloadUrl(value *string)()
    SetFileName(value *string)()
    SetOdataType(value *string)()
    SetSize(value *int64)()
}
