package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutOfBoxExperienceSetting the Windows Autopilot Deployment Profile settings used by the device for the out-of-box experience. Supports: $select, $top, $skip. $Search, $orderBy and $filter are not supported.
type OutOfBoxExperienceSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deviceUsageType property
    deviceUsageType *WindowsDeviceUsageType
    // When TRUE, the link that allows user to start over with a different account on company sign-in is hidden. When false, the link that allows user to start over with a different account on company sign-in is available. Default value is FALSE.
    escapeLinkHidden *bool
    // When TRUE, EULA is hidden to the end user during OOBE. When FALSE, EULA is shown to the end user during OOBE. Default value is FALSE.
    eulaHidden *bool
    // When TRUE, the keyboard selection page is hidden to the end user during OOBE if Language and Region are set. When FALSE, the keyboard selection page is skipped during OOBE.
    keyboardSelectionPageSkipped *bool
    // The OdataType property
    odataType *string
    // When TRUE, privacy settings is hidden to the end user during OOBE. When FALSE, privacy settings is shown to the end user during OOBE. Default value is FALSE.
    privacySettingsHidden *bool
    // The userType property
    userType *WindowsUserType
}
// NewOutOfBoxExperienceSetting instantiates a new OutOfBoxExperienceSetting and sets the default values.
func NewOutOfBoxExperienceSetting()(*OutOfBoxExperienceSetting) {
    m := &OutOfBoxExperienceSetting{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOutOfBoxExperienceSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOutOfBoxExperienceSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutOfBoxExperienceSetting(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OutOfBoxExperienceSetting) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceUsageType gets the deviceUsageType property value. The deviceUsageType property
// returns a *WindowsDeviceUsageType when successful
func (m *OutOfBoxExperienceSetting) GetDeviceUsageType()(*WindowsDeviceUsageType) {
    return m.deviceUsageType
}
// GetEscapeLinkHidden gets the escapeLinkHidden property value. When TRUE, the link that allows user to start over with a different account on company sign-in is hidden. When false, the link that allows user to start over with a different account on company sign-in is available. Default value is FALSE.
// returns a *bool when successful
func (m *OutOfBoxExperienceSetting) GetEscapeLinkHidden()(*bool) {
    return m.escapeLinkHidden
}
// GetEulaHidden gets the eulaHidden property value. When TRUE, EULA is hidden to the end user during OOBE. When FALSE, EULA is shown to the end user during OOBE. Default value is FALSE.
// returns a *bool when successful
func (m *OutOfBoxExperienceSetting) GetEulaHidden()(*bool) {
    return m.eulaHidden
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OutOfBoxExperienceSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceUsageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeviceUsageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceUsageType(val.(*WindowsDeviceUsageType))
        }
        return nil
    }
    res["escapeLinkHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEscapeLinkHidden(val)
        }
        return nil
    }
    res["eulaHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEulaHidden(val)
        }
        return nil
    }
    res["keyboardSelectionPageSkipped"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardSelectionPageSkipped(val)
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
    res["privacySettingsHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacySettingsHidden(val)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsUserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val.(*WindowsUserType))
        }
        return nil
    }
    return res
}
// GetKeyboardSelectionPageSkipped gets the keyboardSelectionPageSkipped property value. When TRUE, the keyboard selection page is hidden to the end user during OOBE if Language and Region are set. When FALSE, the keyboard selection page is skipped during OOBE.
// returns a *bool when successful
func (m *OutOfBoxExperienceSetting) GetKeyboardSelectionPageSkipped()(*bool) {
    return m.keyboardSelectionPageSkipped
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *OutOfBoxExperienceSetting) GetOdataType()(*string) {
    return m.odataType
}
// GetPrivacySettingsHidden gets the privacySettingsHidden property value. When TRUE, privacy settings is hidden to the end user during OOBE. When FALSE, privacy settings is shown to the end user during OOBE. Default value is FALSE.
// returns a *bool when successful
func (m *OutOfBoxExperienceSetting) GetPrivacySettingsHidden()(*bool) {
    return m.privacySettingsHidden
}
// GetUserType gets the userType property value. The userType property
// returns a *WindowsUserType when successful
func (m *OutOfBoxExperienceSetting) GetUserType()(*WindowsUserType) {
    return m.userType
}
// Serialize serializes information the current object
func (m *OutOfBoxExperienceSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceUsageType() != nil {
        cast := (*m.GetDeviceUsageType()).String()
        err := writer.WriteStringValue("deviceUsageType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("escapeLinkHidden", m.GetEscapeLinkHidden())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("eulaHidden", m.GetEulaHidden())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("keyboardSelectionPageSkipped", m.GetKeyboardSelectionPageSkipped())
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
        err := writer.WriteBoolValue("privacySettingsHidden", m.GetPrivacySettingsHidden())
        if err != nil {
            return err
        }
    }
    if m.GetUserType() != nil {
        cast := (*m.GetUserType()).String()
        err := writer.WriteStringValue("userType", &cast)
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
func (m *OutOfBoxExperienceSetting) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceUsageType sets the deviceUsageType property value. The deviceUsageType property
func (m *OutOfBoxExperienceSetting) SetDeviceUsageType(value *WindowsDeviceUsageType)() {
    m.deviceUsageType = value
}
// SetEscapeLinkHidden sets the escapeLinkHidden property value. When TRUE, the link that allows user to start over with a different account on company sign-in is hidden. When false, the link that allows user to start over with a different account on company sign-in is available. Default value is FALSE.
func (m *OutOfBoxExperienceSetting) SetEscapeLinkHidden(value *bool)() {
    m.escapeLinkHidden = value
}
// SetEulaHidden sets the eulaHidden property value. When TRUE, EULA is hidden to the end user during OOBE. When FALSE, EULA is shown to the end user during OOBE. Default value is FALSE.
func (m *OutOfBoxExperienceSetting) SetEulaHidden(value *bool)() {
    m.eulaHidden = value
}
// SetKeyboardSelectionPageSkipped sets the keyboardSelectionPageSkipped property value. When TRUE, the keyboard selection page is hidden to the end user during OOBE if Language and Region are set. When FALSE, the keyboard selection page is skipped during OOBE.
func (m *OutOfBoxExperienceSetting) SetKeyboardSelectionPageSkipped(value *bool)() {
    m.keyboardSelectionPageSkipped = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OutOfBoxExperienceSetting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrivacySettingsHidden sets the privacySettingsHidden property value. When TRUE, privacy settings is hidden to the end user during OOBE. When FALSE, privacy settings is shown to the end user during OOBE. Default value is FALSE.
func (m *OutOfBoxExperienceSetting) SetPrivacySettingsHidden(value *bool)() {
    m.privacySettingsHidden = value
}
// SetUserType sets the userType property value. The userType property
func (m *OutOfBoxExperienceSetting) SetUserType(value *WindowsUserType)() {
    m.userType = value
}
type OutOfBoxExperienceSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceUsageType()(*WindowsDeviceUsageType)
    GetEscapeLinkHidden()(*bool)
    GetEulaHidden()(*bool)
    GetKeyboardSelectionPageSkipped()(*bool)
    GetOdataType()(*string)
    GetPrivacySettingsHidden()(*bool)
    GetUserType()(*WindowsUserType)
    SetDeviceUsageType(value *WindowsDeviceUsageType)()
    SetEscapeLinkHidden(value *bool)()
    SetEulaHidden(value *bool)()
    SetKeyboardSelectionPageSkipped(value *bool)()
    SetOdataType(value *string)()
    SetPrivacySettingsHidden(value *bool)()
    SetUserType(value *WindowsUserType)()
}
