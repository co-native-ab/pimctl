package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingTemplateValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Default value for the setting.
    defaultValue *string
    // Description of the setting.
    description *string
    // Name of the setting.
    name *string
    // The OdataType property
    odataType *string
    // Type of the setting.
    typeEscaped *string
}
// NewSettingTemplateValue instantiates a new SettingTemplateValue and sets the default values.
func NewSettingTemplateValue()(*SettingTemplateValue) {
    m := &SettingTemplateValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSettingTemplateValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingTemplateValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSettingTemplateValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingTemplateValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultValue gets the defaultValue property value. Default value for the setting.
// returns a *string when successful
func (m *SettingTemplateValue) GetDefaultValue()(*string) {
    return m.defaultValue
}
// GetDescription gets the description property value. Description of the setting.
// returns a *string when successful
func (m *SettingTemplateValue) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingTemplateValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
// GetName gets the name property value. Name of the setting.
// returns a *string when successful
func (m *SettingTemplateValue) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SettingTemplateValue) GetOdataType()(*string) {
    return m.odataType
}
// GetTypeEscaped gets the type property value. Type of the setting.
// returns a *string when successful
func (m *SettingTemplateValue) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *SettingTemplateValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
func (m *SettingTemplateValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultValue sets the defaultValue property value. Default value for the setting.
func (m *SettingTemplateValue) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
// SetDescription sets the description property value. Description of the setting.
func (m *SettingTemplateValue) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. Name of the setting.
func (m *SettingTemplateValue) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SettingTemplateValue) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTypeEscaped sets the type property value. Type of the setting.
func (m *SettingTemplateValue) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type SettingTemplateValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(*string)
    GetDescription()(*string)
    GetName()(*string)
    GetOdataType()(*string)
    GetTypeEscaped()(*string)
    SetDefaultValue(value *string)()
    SetDescription(value *string)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetTypeEscaped(value *string)()
}
