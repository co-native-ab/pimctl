package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FileStorageContainerSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether versioning is enabled for items in the container. Optional. Read-write.
    isItemVersioningEnabled *bool
    // Indicates whether Optical Character Recognition (OCR) is enabled for the container. The default value is false. When set to true, OCR extraction is performed for new and updated documents of supported document types, and the extracted fields in the metadata of the document enable end-user search and search-driven solutions. When set to false, existing OCR metadata is not impacted. Optional. Read-write.
    isOcrEnabled *bool
    // The maximum major versions allowed for items in the container. Optional. Read-write.
    itemMajorVersionLimit *int32
    // The OdataType property
    odataType *string
}
// NewFileStorageContainerSettings instantiates a new FileStorageContainerSettings and sets the default values.
func NewFileStorageContainerSettings()(*FileStorageContainerSettings) {
    m := &FileStorageContainerSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileStorageContainerSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainerSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainerSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FileStorageContainerSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FileStorageContainerSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isItemVersioningEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsItemVersioningEnabled(val)
        }
        return nil
    }
    res["isOcrEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOcrEnabled(val)
        }
        return nil
    }
    res["itemMajorVersionLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemMajorVersionLimit(val)
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
    return res
}
// GetIsItemVersioningEnabled gets the isItemVersioningEnabled property value. Indicates whether versioning is enabled for items in the container. Optional. Read-write.
// returns a *bool when successful
func (m *FileStorageContainerSettings) GetIsItemVersioningEnabled()(*bool) {
    return m.isItemVersioningEnabled
}
// GetIsOcrEnabled gets the isOcrEnabled property value. Indicates whether Optical Character Recognition (OCR) is enabled for the container. The default value is false. When set to true, OCR extraction is performed for new and updated documents of supported document types, and the extracted fields in the metadata of the document enable end-user search and search-driven solutions. When set to false, existing OCR metadata is not impacted. Optional. Read-write.
// returns a *bool when successful
func (m *FileStorageContainerSettings) GetIsOcrEnabled()(*bool) {
    return m.isOcrEnabled
}
// GetItemMajorVersionLimit gets the itemMajorVersionLimit property value. The maximum major versions allowed for items in the container. Optional. Read-write.
// returns a *int32 when successful
func (m *FileStorageContainerSettings) GetItemMajorVersionLimit()(*int32) {
    return m.itemMajorVersionLimit
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *FileStorageContainerSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *FileStorageContainerSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isItemVersioningEnabled", m.GetIsItemVersioningEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isOcrEnabled", m.GetIsOcrEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("itemMajorVersionLimit", m.GetItemMajorVersionLimit())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileStorageContainerSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsItemVersioningEnabled sets the isItemVersioningEnabled property value. Indicates whether versioning is enabled for items in the container. Optional. Read-write.
func (m *FileStorageContainerSettings) SetIsItemVersioningEnabled(value *bool)() {
    m.isItemVersioningEnabled = value
}
// SetIsOcrEnabled sets the isOcrEnabled property value. Indicates whether Optical Character Recognition (OCR) is enabled for the container. The default value is false. When set to true, OCR extraction is performed for new and updated documents of supported document types, and the extracted fields in the metadata of the document enable end-user search and search-driven solutions. When set to false, existing OCR metadata is not impacted. Optional. Read-write.
func (m *FileStorageContainerSettings) SetIsOcrEnabled(value *bool)() {
    m.isOcrEnabled = value
}
// SetItemMajorVersionLimit sets the itemMajorVersionLimit property value. The maximum major versions allowed for items in the container. Optional. Read-write.
func (m *FileStorageContainerSettings) SetItemMajorVersionLimit(value *int32)() {
    m.itemMajorVersionLimit = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FileStorageContainerSettings) SetOdataType(value *string)() {
    m.odataType = value
}
type FileStorageContainerSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsItemVersioningEnabled()(*bool)
    GetIsOcrEnabled()(*bool)
    GetItemMajorVersionLimit()(*int32)
    GetOdataType()(*string)
    SetIsItemVersioningEnabled(value *bool)()
    SetIsOcrEnabled(value *bool)()
    SetItemMajorVersionLimit(value *int32)()
    SetOdataType(value *string)()
}
