package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EndUserNotificationSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Notification preference. Possible values are: unknown, microsoft, custom, unknownFutureValue.
    notificationPreference *EndUserNotificationPreference
    // The OdataType property
    odataType *string
    // Positive reinforcement detail.
    positiveReinforcement PositiveReinforcementNotificationable
    // End user notification type. Possible values are: unknown, noTraining, trainingSelected, noNotification, unknownFutureValue.
    settingType *EndUserNotificationSettingType
}
// NewEndUserNotificationSetting instantiates a new EndUserNotificationSetting and sets the default values.
func NewEndUserNotificationSetting()(*EndUserNotificationSetting) {
    m := &EndUserNotificationSetting{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEndUserNotificationSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEndUserNotificationSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.noTrainingNotificationSetting":
                        return NewNoTrainingNotificationSetting(), nil
                    case "#microsoft.graph.trainingNotificationSetting":
                        return NewTrainingNotificationSetting(), nil
                }
            }
        }
    }
    return NewEndUserNotificationSetting(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EndUserNotificationSetting) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EndUserNotificationSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notificationPreference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndUserNotificationPreference)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationPreference(val.(*EndUserNotificationPreference))
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
    res["positiveReinforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePositiveReinforcementNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPositiveReinforcement(val.(PositiveReinforcementNotificationable))
        }
        return nil
    }
    res["settingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndUserNotificationSettingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingType(val.(*EndUserNotificationSettingType))
        }
        return nil
    }
    return res
}
// GetNotificationPreference gets the notificationPreference property value. Notification preference. Possible values are: unknown, microsoft, custom, unknownFutureValue.
// returns a *EndUserNotificationPreference when successful
func (m *EndUserNotificationSetting) GetNotificationPreference()(*EndUserNotificationPreference) {
    return m.notificationPreference
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *EndUserNotificationSetting) GetOdataType()(*string) {
    return m.odataType
}
// GetPositiveReinforcement gets the positiveReinforcement property value. Positive reinforcement detail.
// returns a PositiveReinforcementNotificationable when successful
func (m *EndUserNotificationSetting) GetPositiveReinforcement()(PositiveReinforcementNotificationable) {
    return m.positiveReinforcement
}
// GetSettingType gets the settingType property value. End user notification type. Possible values are: unknown, noTraining, trainingSelected, noNotification, unknownFutureValue.
// returns a *EndUserNotificationSettingType when successful
func (m *EndUserNotificationSetting) GetSettingType()(*EndUserNotificationSettingType) {
    return m.settingType
}
// Serialize serializes information the current object
func (m *EndUserNotificationSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetNotificationPreference() != nil {
        cast := (*m.GetNotificationPreference()).String()
        err := writer.WriteStringValue("notificationPreference", &cast)
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
        err := writer.WriteObjectValue("positiveReinforcement", m.GetPositiveReinforcement())
        if err != nil {
            return err
        }
    }
    if m.GetSettingType() != nil {
        cast := (*m.GetSettingType()).String()
        err := writer.WriteStringValue("settingType", &cast)
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
func (m *EndUserNotificationSetting) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNotificationPreference sets the notificationPreference property value. Notification preference. Possible values are: unknown, microsoft, custom, unknownFutureValue.
func (m *EndUserNotificationSetting) SetNotificationPreference(value *EndUserNotificationPreference)() {
    m.notificationPreference = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EndUserNotificationSetting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPositiveReinforcement sets the positiveReinforcement property value. Positive reinforcement detail.
func (m *EndUserNotificationSetting) SetPositiveReinforcement(value PositiveReinforcementNotificationable)() {
    m.positiveReinforcement = value
}
// SetSettingType sets the settingType property value. End user notification type. Possible values are: unknown, noTraining, trainingSelected, noNotification, unknownFutureValue.
func (m *EndUserNotificationSetting) SetSettingType(value *EndUserNotificationSettingType)() {
    m.settingType = value
}
type EndUserNotificationSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNotificationPreference()(*EndUserNotificationPreference)
    GetOdataType()(*string)
    GetPositiveReinforcement()(PositiveReinforcementNotificationable)
    GetSettingType()(*EndUserNotificationSettingType)
    SetNotificationPreference(value *EndUserNotificationPreference)()
    SetOdataType(value *string)()
    SetPositiveReinforcement(value PositiveReinforcementNotificationable)()
    SetSettingType(value *EndUserNotificationSettingType)()
}
