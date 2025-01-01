package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TrainingSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Type of setting. Possible values are: microsoftCustom, microsoftManaged, noTraining, custom, unknownFutureValue.
    settingType *TrainingSettingType
}
// NewTrainingSetting instantiates a new TrainingSetting and sets the default values.
func NewTrainingSetting()(*TrainingSetting) {
    m := &TrainingSetting{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTrainingSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTrainingSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.customTrainingSetting":
                        return NewCustomTrainingSetting(), nil
                    case "#microsoft.graph.microsoftCustomTrainingSetting":
                        return NewMicrosoftCustomTrainingSetting(), nil
                    case "#microsoft.graph.microsoftManagedTrainingSetting":
                        return NewMicrosoftManagedTrainingSetting(), nil
                    case "#microsoft.graph.microsoftTrainingAssignmentMapping":
                        return NewMicrosoftTrainingAssignmentMapping(), nil
                    case "#microsoft.graph.noTrainingSetting":
                        return NewNoTrainingSetting(), nil
                }
            }
        }
    }
    return NewTrainingSetting(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TrainingSetting) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TrainingSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["settingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrainingSettingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingType(val.(*TrainingSettingType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TrainingSetting) GetOdataType()(*string) {
    return m.odataType
}
// GetSettingType gets the settingType property value. Type of setting. Possible values are: microsoftCustom, microsoftManaged, noTraining, custom, unknownFutureValue.
// returns a *TrainingSettingType when successful
func (m *TrainingSetting) GetSettingType()(*TrainingSettingType) {
    return m.settingType
}
// Serialize serializes information the current object
func (m *TrainingSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *TrainingSetting) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TrainingSetting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSettingType sets the settingType property value. Type of setting. Possible values are: microsoftCustom, microsoftManaged, noTraining, custom, unknownFutureValue.
func (m *TrainingSetting) SetSettingType(value *TrainingSettingType)() {
    m.settingType = value
}
type TrainingSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetSettingType()(*TrainingSettingType)
    SetOdataType(value *string)()
    SetSettingType(value *TrainingSettingType)()
}
