package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AttributeMappingParameterSchema struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The given parameter can be provided multiple times (for example, multiple input strings in the Concatenate(string,string,...) function).
    allowMultipleOccurrences *bool
    // Parameter name.
    name *string
    // The OdataType property
    odataType *string
    // true if the parameter is required; otherwise false.
    required *bool
    // The type property
    typeEscaped *AttributeType
}
// NewAttributeMappingParameterSchema instantiates a new AttributeMappingParameterSchema and sets the default values.
func NewAttributeMappingParameterSchema()(*AttributeMappingParameterSchema) {
    m := &AttributeMappingParameterSchema{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAttributeMappingParameterSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAttributeMappingParameterSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeMappingParameterSchema(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AttributeMappingParameterSchema) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowMultipleOccurrences gets the allowMultipleOccurrences property value. The given parameter can be provided multiple times (for example, multiple input strings in the Concatenate(string,string,...) function).
// returns a *bool when successful
func (m *AttributeMappingParameterSchema) GetAllowMultipleOccurrences()(*bool) {
    return m.allowMultipleOccurrences
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AttributeMappingParameterSchema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowMultipleOccurrences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMultipleOccurrences(val)
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
    res["required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*AttributeType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Parameter name.
// returns a *string when successful
func (m *AttributeMappingParameterSchema) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AttributeMappingParameterSchema) GetOdataType()(*string) {
    return m.odataType
}
// GetRequired gets the required property value. true if the parameter is required; otherwise false.
// returns a *bool when successful
func (m *AttributeMappingParameterSchema) GetRequired()(*bool) {
    return m.required
}
// GetTypeEscaped gets the type property value. The type property
// returns a *AttributeType when successful
func (m *AttributeMappingParameterSchema) GetTypeEscaped()(*AttributeType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *AttributeMappingParameterSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleOccurrences", m.GetAllowMultipleOccurrences())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *AttributeMappingParameterSchema) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowMultipleOccurrences sets the allowMultipleOccurrences property value. The given parameter can be provided multiple times (for example, multiple input strings in the Concatenate(string,string,...) function).
func (m *AttributeMappingParameterSchema) SetAllowMultipleOccurrences(value *bool)() {
    m.allowMultipleOccurrences = value
}
// SetName sets the name property value. Parameter name.
func (m *AttributeMappingParameterSchema) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttributeMappingParameterSchema) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRequired sets the required property value. true if the parameter is required; otherwise false.
func (m *AttributeMappingParameterSchema) SetRequired(value *bool)() {
    m.required = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *AttributeMappingParameterSchema) SetTypeEscaped(value *AttributeType)() {
    m.typeEscaped = value
}
type AttributeMappingParameterSchemaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowMultipleOccurrences()(*bool)
    GetName()(*string)
    GetOdataType()(*string)
    GetRequired()(*bool)
    GetTypeEscaped()(*AttributeType)
    SetAllowMultipleOccurrences(value *bool)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetRequired(value *bool)()
    SetTypeEscaped(value *AttributeType)()
}
