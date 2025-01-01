package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AttributeDefinitionMetadataEntry struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Possible values are: BaseAttributeName, ComplexObjectDefinition, IsContainer, IsCustomerDefined, IsDomainQualified, LinkPropertyNames, LinkTypeName, MaximumLength, ReferencedProperty.
    key *AttributeDefinitionMetadata
    // The OdataType property
    odataType *string
    // Value of the metadata property.
    value *string
}
// NewAttributeDefinitionMetadataEntry instantiates a new AttributeDefinitionMetadataEntry and sets the default values.
func NewAttributeDefinitionMetadataEntry()(*AttributeDefinitionMetadataEntry) {
    m := &AttributeDefinitionMetadataEntry{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAttributeDefinitionMetadataEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAttributeDefinitionMetadataEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeDefinitionMetadataEntry(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AttributeDefinitionMetadataEntry) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AttributeDefinitionMetadataEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeDefinitionMetadata)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(*AttributeDefinitionMetadata))
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
// GetKey gets the key property value. Possible values are: BaseAttributeName, ComplexObjectDefinition, IsContainer, IsCustomerDefined, IsDomainQualified, LinkPropertyNames, LinkTypeName, MaximumLength, ReferencedProperty.
// returns a *AttributeDefinitionMetadata when successful
func (m *AttributeDefinitionMetadataEntry) GetKey()(*AttributeDefinitionMetadata) {
    return m.key
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AttributeDefinitionMetadataEntry) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. Value of the metadata property.
// returns a *string when successful
func (m *AttributeDefinitionMetadataEntry) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *AttributeDefinitionMetadataEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AttributeDefinitionMetadataEntry) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKey sets the key property value. Possible values are: BaseAttributeName, ComplexObjectDefinition, IsContainer, IsCustomerDefined, IsDomainQualified, LinkPropertyNames, LinkTypeName, MaximumLength, ReferencedProperty.
func (m *AttributeDefinitionMetadataEntry) SetKey(value *AttributeDefinitionMetadata)() {
    m.key = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttributeDefinitionMetadataEntry) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. Value of the metadata property.
func (m *AttributeDefinitionMetadataEntry) SetValue(value *string)() {
    m.value = value
}
type AttributeDefinitionMetadataEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*AttributeDefinitionMetadata)
    GetOdataType()(*string)
    GetValue()(*string)
    SetKey(value *AttributeDefinitionMetadata)()
    SetOdataType(value *string)()
    SetValue(value *string)()
}
