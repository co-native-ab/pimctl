package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AppsInstallationOptionsForMac struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies whether users can install Microsoft 365 apps on their MAC devices. The default value is true.
    isMicrosoft365AppsEnabled *bool
    // Specifies whether users can install Skype for Business on their MAC devices running OS X El Capitan 10.11 or later. The default value is true.
    isSkypeForBusinessEnabled *bool
    // The OdataType property
    odataType *string
}
// NewAppsInstallationOptionsForMac instantiates a new AppsInstallationOptionsForMac and sets the default values.
func NewAppsInstallationOptionsForMac()(*AppsInstallationOptionsForMac) {
    m := &AppsInstallationOptionsForMac{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppsInstallationOptionsForMacFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppsInstallationOptionsForMacFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppsInstallationOptionsForMac(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AppsInstallationOptionsForMac) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppsInstallationOptionsForMac) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isMicrosoft365AppsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMicrosoft365AppsEnabled(val)
        }
        return nil
    }
    res["isSkypeForBusinessEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSkypeForBusinessEnabled(val)
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
// GetIsMicrosoft365AppsEnabled gets the isMicrosoft365AppsEnabled property value. Specifies whether users can install Microsoft 365 apps on their MAC devices. The default value is true.
// returns a *bool when successful
func (m *AppsInstallationOptionsForMac) GetIsMicrosoft365AppsEnabled()(*bool) {
    return m.isMicrosoft365AppsEnabled
}
// GetIsSkypeForBusinessEnabled gets the isSkypeForBusinessEnabled property value. Specifies whether users can install Skype for Business on their MAC devices running OS X El Capitan 10.11 or later. The default value is true.
// returns a *bool when successful
func (m *AppsInstallationOptionsForMac) GetIsSkypeForBusinessEnabled()(*bool) {
    return m.isSkypeForBusinessEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AppsInstallationOptionsForMac) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AppsInstallationOptionsForMac) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isMicrosoft365AppsEnabled", m.GetIsMicrosoft365AppsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSkypeForBusinessEnabled", m.GetIsSkypeForBusinessEnabled())
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
func (m *AppsInstallationOptionsForMac) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsMicrosoft365AppsEnabled sets the isMicrosoft365AppsEnabled property value. Specifies whether users can install Microsoft 365 apps on their MAC devices. The default value is true.
func (m *AppsInstallationOptionsForMac) SetIsMicrosoft365AppsEnabled(value *bool)() {
    m.isMicrosoft365AppsEnabled = value
}
// SetIsSkypeForBusinessEnabled sets the isSkypeForBusinessEnabled property value. Specifies whether users can install Skype for Business on their MAC devices running OS X El Capitan 10.11 or later. The default value is true.
func (m *AppsInstallationOptionsForMac) SetIsSkypeForBusinessEnabled(value *bool)() {
    m.isSkypeForBusinessEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppsInstallationOptionsForMac) SetOdataType(value *string)() {
    m.odataType = value
}
type AppsInstallationOptionsForMacable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsMicrosoft365AppsEnabled()(*bool)
    GetIsSkypeForBusinessEnabled()(*bool)
    GetOdataType()(*string)
    SetIsMicrosoft365AppsEnabled(value *bool)()
    SetIsSkypeForBusinessEnabled(value *bool)()
    SetOdataType(value *string)()
}
