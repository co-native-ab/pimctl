package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppAutoUpdateSettings contains properties used to perform the auto-update of an application.
type Win32LobAppAutoUpdateSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Contains value for auto-update superseded apps.
    autoUpdateSupersededAppsState *Win32LobAutoUpdateSupersededAppsState
    // The OdataType property
    odataType *string
}
// NewWin32LobAppAutoUpdateSettings instantiates a new Win32LobAppAutoUpdateSettings and sets the default values.
func NewWin32LobAppAutoUpdateSettings()(*Win32LobAppAutoUpdateSettings) {
    m := &Win32LobAppAutoUpdateSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWin32LobAppAutoUpdateSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWin32LobAppAutoUpdateSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppAutoUpdateSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Win32LobAppAutoUpdateSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAutoUpdateSupersededAppsState gets the autoUpdateSupersededAppsState property value. Contains value for auto-update superseded apps.
// returns a *Win32LobAutoUpdateSupersededAppsState when successful
func (m *Win32LobAppAutoUpdateSettings) GetAutoUpdateSupersededAppsState()(*Win32LobAutoUpdateSupersededAppsState) {
    return m.autoUpdateSupersededAppsState
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Win32LobAppAutoUpdateSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["autoUpdateSupersededAppsState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAutoUpdateSupersededAppsState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoUpdateSupersededAppsState(val.(*Win32LobAutoUpdateSupersededAppsState))
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
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *Win32LobAppAutoUpdateSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *Win32LobAppAutoUpdateSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAutoUpdateSupersededAppsState() != nil {
        cast := (*m.GetAutoUpdateSupersededAppsState()).String()
        err := writer.WriteStringValue("autoUpdateSupersededAppsState", &cast)
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
func (m *Win32LobAppAutoUpdateSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAutoUpdateSupersededAppsState sets the autoUpdateSupersededAppsState property value. Contains value for auto-update superseded apps.
func (m *Win32LobAppAutoUpdateSettings) SetAutoUpdateSupersededAppsState(value *Win32LobAutoUpdateSupersededAppsState)() {
    m.autoUpdateSupersededAppsState = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Win32LobAppAutoUpdateSettings) SetOdataType(value *string)() {
    m.odataType = value
}
type Win32LobAppAutoUpdateSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoUpdateSupersededAppsState()(*Win32LobAutoUpdateSupersededAppsState)
    GetOdataType()(*string)
    SetAutoUpdateSupersededAppsState(value *Win32LobAutoUpdateSupersededAppsState)()
    SetOdataType(value *string)()
}
