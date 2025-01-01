package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RetentionSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The frequency of the backup.
    interval *string
    // The OdataType property
    odataType *string
    // The period of time to retain the protected data for a single Microsoft 365 service.
    period *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewRetentionSetting instantiates a new RetentionSetting and sets the default values.
func NewRetentionSetting()(*RetentionSetting) {
    m := &RetentionSetting{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRetentionSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRetentionSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionSetting(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RetentionSetting) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RetentionSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["interval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInterval(val)
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
    res["period"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriod(val)
        }
        return nil
    }
    return res
}
// GetInterval gets the interval property value. The frequency of the backup.
// returns a *string when successful
func (m *RetentionSetting) GetInterval()(*string) {
    return m.interval
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *RetentionSetting) GetOdataType()(*string) {
    return m.odataType
}
// GetPeriod gets the period property value. The period of time to retain the protected data for a single Microsoft 365 service.
// returns a *ISODuration when successful
func (m *RetentionSetting) GetPeriod()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.period
}
// Serialize serializes information the current object
func (m *RetentionSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("interval", m.GetInterval())
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
        err := writer.WriteISODurationValue("period", m.GetPeriod())
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
func (m *RetentionSetting) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInterval sets the interval property value. The frequency of the backup.
func (m *RetentionSetting) SetInterval(value *string)() {
    m.interval = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RetentionSetting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPeriod sets the period property value. The period of time to retain the protected data for a single Microsoft 365 service.
func (m *RetentionSetting) SetPeriod(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.period = value
}
type RetentionSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInterval()(*string)
    GetOdataType()(*string)
    GetPeriod()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    SetInterval(value *string)()
    SetOdataType(value *string)()
    SetPeriod(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
}
