package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AppsInstallationOptionsForWindows struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies whether users can install Microsoft 365 apps, including Skype for Business, on their Windows devices. The default value is true.
    isMicrosoft365AppsEnabled *bool
    // Specifies whether users can install Microsoft Project on their Windows devices. The default value is true.
    isProjectEnabled *bool
    // Specifies whether users can install Skype for Business (standalone) on their Windows devices. The default value is true.
    isSkypeForBusinessEnabled *bool
    // Specifies whether users can install Visio on their Windows devices. The default value is true.
    isVisioEnabled *bool
    // The OdataType property
    odataType *string
}
// NewAppsInstallationOptionsForWindows instantiates a new AppsInstallationOptionsForWindows and sets the default values.
func NewAppsInstallationOptionsForWindows()(*AppsInstallationOptionsForWindows) {
    m := &AppsInstallationOptionsForWindows{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppsInstallationOptionsForWindowsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppsInstallationOptionsForWindowsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppsInstallationOptionsForWindows(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AppsInstallationOptionsForWindows) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppsInstallationOptionsForWindows) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["isProjectEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsProjectEnabled(val)
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
    res["isVisioEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVisioEnabled(val)
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
// GetIsMicrosoft365AppsEnabled gets the isMicrosoft365AppsEnabled property value. Specifies whether users can install Microsoft 365 apps, including Skype for Business, on their Windows devices. The default value is true.
// returns a *bool when successful
func (m *AppsInstallationOptionsForWindows) GetIsMicrosoft365AppsEnabled()(*bool) {
    return m.isMicrosoft365AppsEnabled
}
// GetIsProjectEnabled gets the isProjectEnabled property value. Specifies whether users can install Microsoft Project on their Windows devices. The default value is true.
// returns a *bool when successful
func (m *AppsInstallationOptionsForWindows) GetIsProjectEnabled()(*bool) {
    return m.isProjectEnabled
}
// GetIsSkypeForBusinessEnabled gets the isSkypeForBusinessEnabled property value. Specifies whether users can install Skype for Business (standalone) on their Windows devices. The default value is true.
// returns a *bool when successful
func (m *AppsInstallationOptionsForWindows) GetIsSkypeForBusinessEnabled()(*bool) {
    return m.isSkypeForBusinessEnabled
}
// GetIsVisioEnabled gets the isVisioEnabled property value. Specifies whether users can install Visio on their Windows devices. The default value is true.
// returns a *bool when successful
func (m *AppsInstallationOptionsForWindows) GetIsVisioEnabled()(*bool) {
    return m.isVisioEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AppsInstallationOptionsForWindows) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AppsInstallationOptionsForWindows) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isMicrosoft365AppsEnabled", m.GetIsMicrosoft365AppsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isProjectEnabled", m.GetIsProjectEnabled())
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
        err := writer.WriteBoolValue("isVisioEnabled", m.GetIsVisioEnabled())
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
func (m *AppsInstallationOptionsForWindows) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsMicrosoft365AppsEnabled sets the isMicrosoft365AppsEnabled property value. Specifies whether users can install Microsoft 365 apps, including Skype for Business, on their Windows devices. The default value is true.
func (m *AppsInstallationOptionsForWindows) SetIsMicrosoft365AppsEnabled(value *bool)() {
    m.isMicrosoft365AppsEnabled = value
}
// SetIsProjectEnabled sets the isProjectEnabled property value. Specifies whether users can install Microsoft Project on their Windows devices. The default value is true.
func (m *AppsInstallationOptionsForWindows) SetIsProjectEnabled(value *bool)() {
    m.isProjectEnabled = value
}
// SetIsSkypeForBusinessEnabled sets the isSkypeForBusinessEnabled property value. Specifies whether users can install Skype for Business (standalone) on their Windows devices. The default value is true.
func (m *AppsInstallationOptionsForWindows) SetIsSkypeForBusinessEnabled(value *bool)() {
    m.isSkypeForBusinessEnabled = value
}
// SetIsVisioEnabled sets the isVisioEnabled property value. Specifies whether users can install Visio on their Windows devices. The default value is true.
func (m *AppsInstallationOptionsForWindows) SetIsVisioEnabled(value *bool)() {
    m.isVisioEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppsInstallationOptionsForWindows) SetOdataType(value *string)() {
    m.odataType = value
}
type AppsInstallationOptionsForWindowsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsMicrosoft365AppsEnabled()(*bool)
    GetIsProjectEnabled()(*bool)
    GetIsSkypeForBusinessEnabled()(*bool)
    GetIsVisioEnabled()(*bool)
    GetOdataType()(*string)
    SetIsMicrosoft365AppsEnabled(value *bool)()
    SetIsProjectEnabled(value *bool)()
    SetIsSkypeForBusinessEnabled(value *bool)()
    SetIsVisioEnabled(value *bool)()
    SetOdataType(value *string)()
}
