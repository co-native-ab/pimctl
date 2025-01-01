package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ObjectDefinitionMetadataEntry struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Possible values are: PropertyNameAccountEnabled, PropertyNameSoftDeleted, IsSoftDeletionSupported, IsSynchronizeAllSupported, ConnectorDataStorageRequired, Extensions, LinkTypeName.
    key *ObjectDefinitionMetadata
    // The OdataType property
    odataType *string
    // Value of the metadata property.
    value *string
}
// NewObjectDefinitionMetadataEntry instantiates a new ObjectDefinitionMetadataEntry and sets the default values.
func NewObjectDefinitionMetadataEntry()(*ObjectDefinitionMetadataEntry) {
    m := &ObjectDefinitionMetadataEntry{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateObjectDefinitionMetadataEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateObjectDefinitionMetadataEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectDefinitionMetadataEntry(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ObjectDefinitionMetadataEntry) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ObjectDefinitionMetadataEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseObjectDefinitionMetadata)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(*ObjectDefinitionMetadata))
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. Possible values are: PropertyNameAccountEnabled, PropertyNameSoftDeleted, IsSoftDeletionSupported, IsSynchronizeAllSupported, ConnectorDataStorageRequired, Extensions, LinkTypeName.
// returns a *ObjectDefinitionMetadata when successful
func (m *ObjectDefinitionMetadataEntry) GetKey()(*ObjectDefinitionMetadata) {
    return m.key
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ObjectDefinitionMetadataEntry) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. Value of the metadata property.
// returns a *string when successful
func (m *ObjectDefinitionMetadataEntry) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ObjectDefinitionMetadataEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetKey() != nil {
        cast := (*m.GetKey()).String()
        err := writer.WriteStringValue("key", &cast)
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
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *ObjectDefinitionMetadataEntry) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKey sets the key property value. Possible values are: PropertyNameAccountEnabled, PropertyNameSoftDeleted, IsSoftDeletionSupported, IsSynchronizeAllSupported, ConnectorDataStorageRequired, Extensions, LinkTypeName.
func (m *ObjectDefinitionMetadataEntry) SetKey(value *ObjectDefinitionMetadata)() {
    m.key = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ObjectDefinitionMetadataEntry) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. Value of the metadata property.
func (m *ObjectDefinitionMetadataEntry) SetValue(value *string)() {
    m.value = value
}
type ObjectDefinitionMetadataEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*ObjectDefinitionMetadata)
    GetOdataType()(*string)
    GetValue()(*string)
    SetKey(value *ObjectDefinitionMetadata)()
    SetOdataType(value *string)()
    SetValue(value *string)()
}