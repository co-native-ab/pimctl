package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExtensionSchemaProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the strongly typed property defined as part of a schema extension.
    name *string
    // The OdataType property
    odataType *string
    // The type of the property that is defined as part of a schema extension.  Allowed values are Binary, Boolean, DateTime, Integer, or String. For more information, see Supported property data types.
    typeEscaped *string
}
// NewExtensionSchemaProperty instantiates a new ExtensionSchemaProperty and sets the default values.
func NewExtensionSchemaProperty()(*ExtensionSchemaProperty) {
    m := &ExtensionSchemaProperty{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExtensionSchemaPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExtensionSchemaPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExtensionSchemaProperty(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ExtensionSchemaProperty) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExtensionSchemaProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the strongly typed property defined as part of a schema extension.
// returns a *string when successful
func (m *ExtensionSchemaProperty) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ExtensionSchemaProperty) GetOdataType()(*string) {
    return m.odataType
}
// GetTypeEscaped gets the type property value. The type of the property that is defined as part of a schema extension.  Allowed values are Binary, Boolean, DateTime, Integer, or String. For more information, see Supported property data types.
// returns a *string when successful
func (m *ExtensionSchemaProperty) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ExtensionSchemaProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *ExtensionSchemaProperty) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name of the strongly typed property defined as part of a schema extension.
func (m *ExtensionSchemaProperty) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExtensionSchemaProperty) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTypeEscaped sets the type property value. The type of the property that is defined as part of a schema extension.  Allowed values are Binary, Boolean, DateTime, Integer, or String. For more information, see Supported property data types.
func (m *ExtensionSchemaProperty) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type ExtensionSchemaPropertyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetOdataType()(*string)
    GetTypeEscaped()(*string)
    SetName(value *string)()
    SetOdataType(value *string)()
    SetTypeEscaped(value *string)()
}
