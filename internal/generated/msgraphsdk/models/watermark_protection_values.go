package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WatermarkProtectionValues struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether to apply a watermark to any shared content.
    isEnabledForContentSharing *bool
    // Indicates whether to apply a watermark to everyone's video feed.
    isEnabledForVideo *bool
    // The OdataType property
    odataType *string
}
// NewWatermarkProtectionValues instantiates a new WatermarkProtectionValues and sets the default values.
func NewWatermarkProtectionValues()(*WatermarkProtectionValues) {
    m := &WatermarkProtectionValues{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWatermarkProtectionValuesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWatermarkProtectionValuesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWatermarkProtectionValues(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WatermarkProtectionValues) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WatermarkProtectionValues) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isEnabledForContentSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForContentSharing(val)
        }
        return nil
    }
    res["isEnabledForVideo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForVideo(val)
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
// GetIsEnabledForContentSharing gets the isEnabledForContentSharing property value. Indicates whether to apply a watermark to any shared content.
// returns a *bool when successful
func (m *WatermarkProtectionValues) GetIsEnabledForContentSharing()(*bool) {
    return m.isEnabledForContentSharing
}
// GetIsEnabledForVideo gets the isEnabledForVideo property value. Indicates whether to apply a watermark to everyone's video feed.
// returns a *bool when successful
func (m *WatermarkProtectionValues) GetIsEnabledForVideo()(*bool) {
    return m.isEnabledForVideo
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WatermarkProtectionValues) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *WatermarkProtectionValues) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isEnabledForContentSharing", m.GetIsEnabledForContentSharing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabledForVideo", m.GetIsEnabledForVideo())
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
func (m *WatermarkProtectionValues) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsEnabledForContentSharing sets the isEnabledForContentSharing property value. Indicates whether to apply a watermark to any shared content.
func (m *WatermarkProtectionValues) SetIsEnabledForContentSharing(value *bool)() {
    m.isEnabledForContentSharing = value
}
// SetIsEnabledForVideo sets the isEnabledForVideo property value. Indicates whether to apply a watermark to everyone's video feed.
func (m *WatermarkProtectionValues) SetIsEnabledForVideo(value *bool)() {
    m.isEnabledForVideo = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WatermarkProtectionValues) SetOdataType(value *string)() {
    m.odataType = value
}
type WatermarkProtectionValuesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsEnabledForContentSharing()(*bool)
    GetIsEnabledForVideo()(*bool)
    GetOdataType()(*string)
    SetIsEnabledForContentSharing(value *bool)()
    SetIsEnabledForVideo(value *bool)()
    SetOdataType(value *string)()
}
