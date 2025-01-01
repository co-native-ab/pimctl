package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TeamMemberSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If set to true, members can add and remove apps.
    allowAddRemoveApps *bool
    // If set to true, members can add and update private channels.
    allowCreatePrivateChannels *bool
    // If set to true, members can add and update channels.
    allowCreateUpdateChannels *bool
    // If set to true, members can add, update, and remove connectors.
    allowCreateUpdateRemoveConnectors *bool
    // If set to true, members can add, update, and remove tabs.
    allowCreateUpdateRemoveTabs *bool
    // If set to true, members can delete channels.
    allowDeleteChannels *bool
    // The OdataType property
    odataType *string
}
// NewTeamMemberSettings instantiates a new TeamMemberSettings and sets the default values.
func NewTeamMemberSettings()(*TeamMemberSettings) {
    m := &TeamMemberSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamMemberSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamMemberSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamMemberSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamMemberSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowAddRemoveApps gets the allowAddRemoveApps property value. If set to true, members can add and remove apps.
// returns a *bool when successful
func (m *TeamMemberSettings) GetAllowAddRemoveApps()(*bool) {
    return m.allowAddRemoveApps
}
// GetAllowCreatePrivateChannels gets the allowCreatePrivateChannels property value. If set to true, members can add and update private channels.
// returns a *bool when successful
func (m *TeamMemberSettings) GetAllowCreatePrivateChannels()(*bool) {
    return m.allowCreatePrivateChannels
}
// GetAllowCreateUpdateChannels gets the allowCreateUpdateChannels property value. If set to true, members can add and update channels.
// returns a *bool when successful
func (m *TeamMemberSettings) GetAllowCreateUpdateChannels()(*bool) {
    return m.allowCreateUpdateChannels
}
// GetAllowCreateUpdateRemoveConnectors gets the allowCreateUpdateRemoveConnectors property value. If set to true, members can add, update, and remove connectors.
// returns a *bool when successful
func (m *TeamMemberSettings) GetAllowCreateUpdateRemoveConnectors()(*bool) {
    return m.allowCreateUpdateRemoveConnectors
}
// GetAllowCreateUpdateRemoveTabs gets the allowCreateUpdateRemoveTabs property value. If set to true, members can add, update, and remove tabs.
// returns a *bool when successful
func (m *TeamMemberSettings) GetAllowCreateUpdateRemoveTabs()(*bool) {
    return m.allowCreateUpdateRemoveTabs
}
// GetAllowDeleteChannels gets the allowDeleteChannels property value. If set to true, members can delete channels.
// returns a *bool when successful
func (m *TeamMemberSettings) GetAllowDeleteChannels()(*bool) {
    return m.allowDeleteChannels
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamMemberSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowAddRemoveApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAddRemoveApps(val)
        }
        return nil
    }
    res["allowCreatePrivateChannels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCreatePrivateChannels(val)
        }
        return nil
    }
    res["allowCreateUpdateChannels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCreateUpdateChannels(val)
        }
        return nil
    }
    res["allowCreateUpdateRemoveConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCreateUpdateRemoveConnectors(val)
        }
        return nil
    }
    res["allowCreateUpdateRemoveTabs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCreateUpdateRemoveTabs(val)
        }
        return nil
    }
    res["allowDeleteChannels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeleteChannels(val)
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
func (m *TeamMemberSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamMemberSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowAddRemoveApps", m.GetAllowAddRemoveApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowCreatePrivateChannels", m.GetAllowCreatePrivateChannels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowCreateUpdateChannels", m.GetAllowCreateUpdateChannels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowCreateUpdateRemoveConnectors", m.GetAllowCreateUpdateRemoveConnectors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowCreateUpdateRemoveTabs", m.GetAllowCreateUpdateRemoveTabs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeleteChannels", m.GetAllowDeleteChannels())
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
func (m *TeamMemberSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowAddRemoveApps sets the allowAddRemoveApps property value. If set to true, members can add and remove apps.
func (m *TeamMemberSettings) SetAllowAddRemoveApps(value *bool)() {
    m.allowAddRemoveApps = value
}
// SetAllowCreatePrivateChannels sets the allowCreatePrivateChannels property value. If set to true, members can add and update private channels.
func (m *TeamMemberSettings) SetAllowCreatePrivateChannels(value *bool)() {
    m.allowCreatePrivateChannels = value
}
// SetAllowCreateUpdateChannels sets the allowCreateUpdateChannels property value. If set to true, members can add and update channels.
func (m *TeamMemberSettings) SetAllowCreateUpdateChannels(value *bool)() {
    m.allowCreateUpdateChannels = value
}
// SetAllowCreateUpdateRemoveConnectors sets the allowCreateUpdateRemoveConnectors property value. If set to true, members can add, update, and remove connectors.
func (m *TeamMemberSettings) SetAllowCreateUpdateRemoveConnectors(value *bool)() {
    m.allowCreateUpdateRemoveConnectors = value
}
// SetAllowCreateUpdateRemoveTabs sets the allowCreateUpdateRemoveTabs property value. If set to true, members can add, update, and remove tabs.
func (m *TeamMemberSettings) SetAllowCreateUpdateRemoveTabs(value *bool)() {
    m.allowCreateUpdateRemoveTabs = value
}
// SetAllowDeleteChannels sets the allowDeleteChannels property value. If set to true, members can delete channels.
func (m *TeamMemberSettings) SetAllowDeleteChannels(value *bool)() {
    m.allowDeleteChannels = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamMemberSettings) SetOdataType(value *string)() {
    m.odataType = value
}
type TeamMemberSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAddRemoveApps()(*bool)
    GetAllowCreatePrivateChannels()(*bool)
    GetAllowCreateUpdateChannels()(*bool)
    GetAllowCreateUpdateRemoveConnectors()(*bool)
    GetAllowCreateUpdateRemoveTabs()(*bool)
    GetAllowDeleteChannels()(*bool)
    GetOdataType()(*string)
    SetAllowAddRemoveApps(value *bool)()
    SetAllowCreatePrivateChannels(value *bool)()
    SetAllowCreateUpdateChannels(value *bool)()
    SetAllowCreateUpdateRemoveConnectors(value *bool)()
    SetAllowCreateUpdateRemoveTabs(value *bool)()
    SetAllowDeleteChannels(value *bool)()
    SetOdataType(value *string)()
}
