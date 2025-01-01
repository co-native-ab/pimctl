package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ObjectMappingMetadataEntry struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Possible values are: EscrowBehavior, DisableMonitoringForChanges, OriginalJoiningProperty, Disposition, IsCustomerDefined, ExcludeFromReporting, Unsynchronized.
    key *ObjectMappingMetadata
    // The OdataType property
    odataType *string
    // Value of the metadata property.
    value *string
}
// NewObjectMappingMetadataEntry instantiates a new ObjectMappingMetadataEntry and sets the default values.
func NewObjectMappingMetadataEntry()(*ObjectMappingMetadataEntry) {
    m := &ObjectMappingMetadataEntry{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateObjectMappingMetadataEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateObjectMappingMetadataEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectMappingMetadataEntry(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ObjectMappingMetadataEntry) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ObjectMappingMetadataEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseObjectMappingMetadata)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(*ObjectMappingMetadata))
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
// GetKey gets the key property value. Possible values are: EscrowBehavior, DisableMonitoringForChanges, OriginalJoiningProperty, Disposition, IsCustomerDefined, ExcludeFromReporting, Unsynchronized.
// returns a *ObjectMappingMetadata when successful
func (m *ObjectMappingMetadataEntry) GetKey()(*ObjectMappingMetadata) {
    return m.key
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ObjectMappingMetadataEntry) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. Value of the metadata property.
// returns a *string when successful
func (m *ObjectMappingMetadataEntry) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ObjectMappingMetadataEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ObjectMappingMetadataEntry) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKey sets the key property value. Possible values are: EscrowBehavior, DisableMonitoringForChanges, OriginalJoiningProperty, Disposition, IsCustomerDefined, ExcludeFromReporting, Unsynchronized.
func (m *ObjectMappingMetadataEntry) SetKey(value *ObjectMappingMetadata)() {
    m.key = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ObjectMappingMetadataEntry) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. Value of the metadata property.
func (m *ObjectMappingMetadataEntry) SetValue(value *string)() {
    m.value = value
}
type ObjectMappingMetadataEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*ObjectMappingMetadata)
    GetOdataType()(*string)
    GetValue()(*string)
    SetKey(value *ObjectMappingMetadata)()
    SetOdataType(value *string)()
    SetValue(value *string)()
}
