package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RecycleBinSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Recycle bin retention period override in days for deleted content. The default value is 93; the value range is 7 to 180. The setting applies to newly deleted content only. Setting this property to null reverts to its default value. Read-write.
    retentionPeriodOverrideDays *int32
}
// NewRecycleBinSettings instantiates a new RecycleBinSettings and sets the default values.
func NewRecycleBinSettings()(*RecycleBinSettings) {
    m := &RecycleBinSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRecycleBinSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRecycleBinSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecycleBinSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RecycleBinSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RecycleBinSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["retentionPeriodOverrideDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetentionPeriodOverrideDays(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *RecycleBinSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetRetentionPeriodOverrideDays gets the retentionPeriodOverrideDays property value. Recycle bin retention period override in days for deleted content. The default value is 93; the value range is 7 to 180. The setting applies to newly deleted content only. Setting this property to null reverts to its default value. Read-write.
// returns a *int32 when successful
func (m *RecycleBinSettings) GetRetentionPeriodOverrideDays()(*int32) {
    return m.retentionPeriodOverrideDays
}
// Serialize serializes information the current object
func (m *RecycleBinSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("retentionPeriodOverrideDays", m.GetRetentionPeriodOverrideDays())
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
func (m *RecycleBinSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RecycleBinSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRetentionPeriodOverrideDays sets the retentionPeriodOverrideDays property value. Recycle bin retention period override in days for deleted content. The default value is 93; the value range is 7 to 180. The setting applies to newly deleted content only. Setting this property to null reverts to its default value. Read-write.
func (m *RecycleBinSettings) SetRetentionPeriodOverrideDays(value *int32)() {
    m.retentionPeriodOverrideDays = value
}
type RecycleBinSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetRetentionPeriodOverrideDays()(*int32)
    SetOdataType(value *string)()
    SetRetentionPeriodOverrideDays(value *int32)()
}
