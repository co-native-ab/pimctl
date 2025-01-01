package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcRestorePointSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The time interval in hours to take snapshots (restore points) of a Cloud PC automatically. Possible values are: default, fourHours, sixHours, twelveHours, sixteenHours, twentyFourHours, unknownFutureValue. The default value is default that indicates that the time interval for automatic capturing of restore point snapshots is set to 12 hours.
    frequencyType *CloudPcRestorePointFrequencyType
    // The OdataType property
    odataType *string
    // If true, the user has the ability to use snapshots to restore Cloud PCs. If false, non-admin users can't use snapshots to restore the Cloud PC.
    userRestoreEnabled *bool
}
// NewCloudPcRestorePointSetting instantiates a new CloudPcRestorePointSetting and sets the default values.
func NewCloudPcRestorePointSetting()(*CloudPcRestorePointSetting) {
    m := &CloudPcRestorePointSetting{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcRestorePointSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcRestorePointSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcRestorePointSetting(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudPcRestorePointSetting) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcRestorePointSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["frequencyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcRestorePointFrequencyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrequencyType(val.(*CloudPcRestorePointFrequencyType))
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
    res["userRestoreEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRestoreEnabled(val)
        }
        return nil
    }
    return res
}
// GetFrequencyType gets the frequencyType property value. The time interval in hours to take snapshots (restore points) of a Cloud PC automatically. Possible values are: default, fourHours, sixHours, twelveHours, sixteenHours, twentyFourHours, unknownFutureValue. The default value is default that indicates that the time interval for automatic capturing of restore point snapshots is set to 12 hours.
// returns a *CloudPcRestorePointFrequencyType when successful
func (m *CloudPcRestorePointSetting) GetFrequencyType()(*CloudPcRestorePointFrequencyType) {
    return m.frequencyType
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CloudPcRestorePointSetting) GetOdataType()(*string) {
    return m.odataType
}
// GetUserRestoreEnabled gets the userRestoreEnabled property value. If true, the user has the ability to use snapshots to restore Cloud PCs. If false, non-admin users can't use snapshots to restore the Cloud PC.
// returns a *bool when successful
func (m *CloudPcRestorePointSetting) GetUserRestoreEnabled()(*bool) {
    return m.userRestoreEnabled
}
// Serialize serializes information the current object
func (m *CloudPcRestorePointSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFrequencyType() != nil {
        cast := (*m.GetFrequencyType()).String()
        err := writer.WriteStringValue("frequencyType", &cast)
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
        err := writer.WriteBoolValue("userRestoreEnabled", m.GetUserRestoreEnabled())
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
func (m *CloudPcRestorePointSetting) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFrequencyType sets the frequencyType property value. The time interval in hours to take snapshots (restore points) of a Cloud PC automatically. Possible values are: default, fourHours, sixHours, twelveHours, sixteenHours, twentyFourHours, unknownFutureValue. The default value is default that indicates that the time interval for automatic capturing of restore point snapshots is set to 12 hours.
func (m *CloudPcRestorePointSetting) SetFrequencyType(value *CloudPcRestorePointFrequencyType)() {
    m.frequencyType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcRestorePointSetting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserRestoreEnabled sets the userRestoreEnabled property value. If true, the user has the ability to use snapshots to restore Cloud PCs. If false, non-admin users can't use snapshots to restore the Cloud PC.
func (m *CloudPcRestorePointSetting) SetUserRestoreEnabled(value *bool)() {
    m.userRestoreEnabled = value
}
type CloudPcRestorePointSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFrequencyType()(*CloudPcRestorePointFrequencyType)
    GetOdataType()(*string)
    GetUserRestoreEnabled()(*bool)
    SetFrequencyType(value *CloudPcRestorePointFrequencyType)()
    SetOdataType(value *string)()
    SetUserRestoreEnabled(value *bool)()
}
