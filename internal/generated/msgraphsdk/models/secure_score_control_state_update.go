package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecureScoreControlStateUpdate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Assigns the control to the user who will take the action.
    assignedTo *string
    // Provides optional comment about the control.
    comment *string
    // The OdataType property
    odataType *string
    // State of the control, which can be modified via a PATCH command (for example, ignored, thirdParty).
    state *string
    // ID of the user who updated tenant state.
    updatedBy *string
    // Time at which the control state was updated.
    updatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewSecureScoreControlStateUpdate instantiates a new SecureScoreControlStateUpdate and sets the default values.
func NewSecureScoreControlStateUpdate()(*SecureScoreControlStateUpdate) {
    m := &SecureScoreControlStateUpdate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecureScoreControlStateUpdateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecureScoreControlStateUpdateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecureScoreControlStateUpdate(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecureScoreControlStateUpdate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignedTo gets the assignedTo property value. Assigns the control to the user who will take the action.
// returns a *string when successful
func (m *SecureScoreControlStateUpdate) GetAssignedTo()(*string) {
    return m.assignedTo
}
// GetComment gets the comment property value. Provides optional comment about the control.
// returns a *string when successful
func (m *SecureScoreControlStateUpdate) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecureScoreControlStateUpdate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val)
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["updatedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedBy(val)
        }
        return nil
    }
    res["updatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedDateTime(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SecureScoreControlStateUpdate) GetOdataType()(*string) {
    return m.odataType
}
// GetState gets the state property value. State of the control, which can be modified via a PATCH command (for example, ignored, thirdParty).
// returns a *string when successful
func (m *SecureScoreControlStateUpdate) GetState()(*string) {
    return m.state
}
// GetUpdatedBy gets the updatedBy property value. ID of the user who updated tenant state.
// returns a *string when successful
func (m *SecureScoreControlStateUpdate) GetUpdatedBy()(*string) {
    return m.updatedBy
}
// GetUpdatedDateTime gets the updatedDateTime property value. Time at which the control state was updated.
// returns a *Time when successful
func (m *SecureScoreControlStateUpdate) GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedDateTime
}
// Serialize serializes information the current object
func (m *SecureScoreControlStateUpdate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("comment", m.GetComment())
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
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updatedBy", m.GetUpdatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedDateTime", m.GetUpdatedDateTime())
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
func (m *SecureScoreControlStateUpdate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignedTo sets the assignedTo property value. Assigns the control to the user who will take the action.
func (m *SecureScoreControlStateUpdate) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
// SetComment sets the comment property value. Provides optional comment about the control.
func (m *SecureScoreControlStateUpdate) SetComment(value *string)() {
    m.comment = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SecureScoreControlStateUpdate) SetOdataType(value *string)() {
    m.odataType = value
}
// SetState sets the state property value. State of the control, which can be modified via a PATCH command (for example, ignored, thirdParty).
func (m *SecureScoreControlStateUpdate) SetState(value *string)() {
    m.state = value
}
// SetUpdatedBy sets the updatedBy property value. ID of the user who updated tenant state.
func (m *SecureScoreControlStateUpdate) SetUpdatedBy(value *string)() {
    m.updatedBy = value
}
// SetUpdatedDateTime sets the updatedDateTime property value. Time at which the control state was updated.
func (m *SecureScoreControlStateUpdate) SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedDateTime = value
}
type SecureScoreControlStateUpdateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedTo()(*string)
    GetComment()(*string)
    GetOdataType()(*string)
    GetState()(*string)
    GetUpdatedBy()(*string)
    GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAssignedTo(value *string)()
    SetComment(value *string)()
    SetOdataType(value *string)()
    SetState(value *string)()
    SetUpdatedBy(value *string)()
    SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
