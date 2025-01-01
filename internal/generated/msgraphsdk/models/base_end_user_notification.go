package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BaseEndUserNotification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The default language for the end user notification.
    defaultLanguage *string
    // The endUserNotification property
    endUserNotification EndUserNotificationable
    // The OdataType property
    odataType *string
}
// NewBaseEndUserNotification instantiates a new BaseEndUserNotification and sets the default values.
func NewBaseEndUserNotification()(*BaseEndUserNotification) {
    m := &BaseEndUserNotification{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBaseEndUserNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBaseEndUserNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.positiveReinforcementNotification":
                        return NewPositiveReinforcementNotification(), nil
                    case "#microsoft.graph.simulationNotification":
                        return NewSimulationNotification(), nil
                    case "#microsoft.graph.trainingReminderNotification":
                        return NewTrainingReminderNotification(), nil
                }
            }
        }
    }
    return NewBaseEndUserNotification(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BaseEndUserNotification) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultLanguage gets the defaultLanguage property value. The default language for the end user notification.
// returns a *string when successful
func (m *BaseEndUserNotification) GetDefaultLanguage()(*string) {
    return m.defaultLanguage
}
// GetEndUserNotification gets the endUserNotification property value. The endUserNotification property
// returns a EndUserNotificationable when successful
func (m *BaseEndUserNotification) GetEndUserNotification()(EndUserNotificationable) {
    return m.endUserNotification
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BaseEndUserNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLanguage(val)
        }
        return nil
    }
    res["endUserNotification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEndUserNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndUserNotification(val.(EndUserNotificationable))
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
func (m *BaseEndUserNotification) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *BaseEndUserNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultLanguage", m.GetDefaultLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("endUserNotification", m.GetEndUserNotification())
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
func (m *BaseEndUserNotification) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultLanguage sets the defaultLanguage property value. The default language for the end user notification.
func (m *BaseEndUserNotification) SetDefaultLanguage(value *string)() {
    m.defaultLanguage = value
}
// SetEndUserNotification sets the endUserNotification property value. The endUserNotification property
func (m *BaseEndUserNotification) SetEndUserNotification(value EndUserNotificationable)() {
    m.endUserNotification = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BaseEndUserNotification) SetOdataType(value *string)() {
    m.odataType = value
}
type BaseEndUserNotificationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultLanguage()(*string)
    GetEndUserNotification()(EndUserNotificationable)
    GetOdataType()(*string)
    SetDefaultLanguage(value *string)()
    SetEndUserNotification(value EndUserNotificationable)()
    SetOdataType(value *string)()
}
