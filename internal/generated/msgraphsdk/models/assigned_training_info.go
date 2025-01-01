package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AssignedTrainingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Number of users who were assigned the training in an attack simulation and training campaign.
    assignedUserCount *int32
    // Number of users who completed the training in an attack simulation and training campaign.
    completedUserCount *int32
    // Display name of the training in an attack simulation and training campaign.
    displayName *string
    // The OdataType property
    odataType *string
}
// NewAssignedTrainingInfo instantiates a new AssignedTrainingInfo and sets the default values.
func NewAssignedTrainingInfo()(*AssignedTrainingInfo) {
    m := &AssignedTrainingInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAssignedTrainingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAssignedTrainingInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignedTrainingInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AssignedTrainingInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignedUserCount gets the assignedUserCount property value. Number of users who were assigned the training in an attack simulation and training campaign.
// returns a *int32 when successful
func (m *AssignedTrainingInfo) GetAssignedUserCount()(*int32) {
    return m.assignedUserCount
}
// GetCompletedUserCount gets the completedUserCount property value. Number of users who completed the training in an attack simulation and training campaign.
// returns a *int32 when successful
func (m *AssignedTrainingInfo) GetCompletedUserCount()(*int32) {
    return m.completedUserCount
}
// GetDisplayName gets the displayName property value. Display name of the training in an attack simulation and training campaign.
// returns a *string when successful
func (m *AssignedTrainingInfo) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AssignedTrainingInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedUserCount(val)
        }
        return nil
    }
    res["completedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedUserCount(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
func (m *AssignedTrainingInfo) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AssignedTrainingInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("assignedUserCount", m.GetAssignedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("completedUserCount", m.GetCompletedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
func (m *AssignedTrainingInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignedUserCount sets the assignedUserCount property value. Number of users who were assigned the training in an attack simulation and training campaign.
func (m *AssignedTrainingInfo) SetAssignedUserCount(value *int32)() {
    m.assignedUserCount = value
}
// SetCompletedUserCount sets the completedUserCount property value. Number of users who completed the training in an attack simulation and training campaign.
func (m *AssignedTrainingInfo) SetCompletedUserCount(value *int32)() {
    m.completedUserCount = value
}
// SetDisplayName sets the displayName property value. Display name of the training in an attack simulation and training campaign.
func (m *AssignedTrainingInfo) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignedTrainingInfo) SetOdataType(value *string)() {
    m.odataType = value
}
type AssignedTrainingInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedUserCount()(*int32)
    GetCompletedUserCount()(*int32)
    GetDisplayName()(*string)
    GetOdataType()(*string)
    SetAssignedUserCount(value *int32)()
    SetCompletedUserCount(value *int32)()
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
}
