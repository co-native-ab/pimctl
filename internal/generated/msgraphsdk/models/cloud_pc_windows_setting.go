package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcWindowsSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Windows language or region tag to use for language pack configuration and localization of the Cloud PC. The default value is en-US, which corresponds to English (United States).
    locale *string
    // The OdataType property
    odataType *string
}
// NewCloudPcWindowsSetting instantiates a new CloudPcWindowsSetting and sets the default values.
func NewCloudPcWindowsSetting()(*CloudPcWindowsSetting) {
    m := &CloudPcWindowsSetting{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcWindowsSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcWindowsSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcWindowsSetting(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudPcWindowsSetting) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcWindowsSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
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
// GetLocale gets the locale property value. The Windows language or region tag to use for language pack configuration and localization of the Cloud PC. The default value is en-US, which corresponds to English (United States).
// returns a *string when successful
func (m *CloudPcWindowsSetting) GetLocale()(*string) {
    return m.locale
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CloudPcWindowsSetting) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CloudPcWindowsSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("locale", m.GetLocale())
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
func (m *CloudPcWindowsSetting) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLocale sets the locale property value. The Windows language or region tag to use for language pack configuration and localization of the Cloud PC. The default value is en-US, which corresponds to English (United States).
func (m *CloudPcWindowsSetting) SetLocale(value *string)() {
    m.locale = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcWindowsSetting) SetOdataType(value *string)() {
    m.odataType = value
}
type CloudPcWindowsSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocale()(*string)
    GetOdataType()(*string)
    SetLocale(value *string)()
    SetOdataType(value *string)()
}
