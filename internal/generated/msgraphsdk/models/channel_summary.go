package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ChannelSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Count of guests in a channel.
    guestsCount *int32
    // Indicates whether external members are included on the channel.
    hasMembersFromOtherTenants *bool
    // Count of members in a channel.
    membersCount *int32
    // The OdataType property
    odataType *string
    // Count of owners in a channel.
    ownersCount *int32
}
// NewChannelSummary instantiates a new ChannelSummary and sets the default values.
func NewChannelSummary()(*ChannelSummary) {
    m := &ChannelSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChannelSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateChannelSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ChannelSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ChannelSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["guestsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuestsCount(val)
        }
        return nil
    }
    res["hasMembersFromOtherTenants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMembersFromOtherTenants(val)
        }
        return nil
    }
    res["membersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCount(val)
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
    res["ownersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnersCount(val)
        }
        return nil
    }
    return res
}
// GetGuestsCount gets the guestsCount property value. Count of guests in a channel.
// returns a *int32 when successful
func (m *ChannelSummary) GetGuestsCount()(*int32) {
    return m.guestsCount
}
// GetHasMembersFromOtherTenants gets the hasMembersFromOtherTenants property value. Indicates whether external members are included on the channel.
// returns a *bool when successful
func (m *ChannelSummary) GetHasMembersFromOtherTenants()(*bool) {
    return m.hasMembersFromOtherTenants
}
// GetMembersCount gets the membersCount property value. Count of members in a channel.
// returns a *int32 when successful
func (m *ChannelSummary) GetMembersCount()(*int32) {
    return m.membersCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ChannelSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetOwnersCount gets the ownersCount property value. Count of owners in a channel.
// returns a *int32 when successful
func (m *ChannelSummary) GetOwnersCount()(*int32) {
    return m.ownersCount
}
// Serialize serializes information the current object
func (m *ChannelSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("guestsCount", m.GetGuestsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasMembersFromOtherTenants", m.GetHasMembersFromOtherTenants())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("membersCount", m.GetMembersCount())
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
        err := writer.WriteInt32Value("ownersCount", m.GetOwnersCount())
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
func (m *ChannelSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGuestsCount sets the guestsCount property value. Count of guests in a channel.
func (m *ChannelSummary) SetGuestsCount(value *int32)() {
    m.guestsCount = value
}
// SetHasMembersFromOtherTenants sets the hasMembersFromOtherTenants property value. Indicates whether external members are included on the channel.
func (m *ChannelSummary) SetHasMembersFromOtherTenants(value *bool)() {
    m.hasMembersFromOtherTenants = value
}
// SetMembersCount sets the membersCount property value. Count of members in a channel.
func (m *ChannelSummary) SetMembersCount(value *int32)() {
    m.membersCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChannelSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOwnersCount sets the ownersCount property value. Count of owners in a channel.
func (m *ChannelSummary) SetOwnersCount(value *int32)() {
    m.ownersCount = value
}
type ChannelSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGuestsCount()(*int32)
    GetHasMembersFromOtherTenants()(*bool)
    GetMembersCount()(*int32)
    GetOdataType()(*string)
    GetOwnersCount()(*int32)
    SetGuestsCount(value *int32)()
    SetHasMembersFromOtherTenants(value *bool)()
    SetMembersCount(value *int32)()
    SetOdataType(value *string)()
    SetOwnersCount(value *int32)()
}
