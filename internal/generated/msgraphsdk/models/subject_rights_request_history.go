package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubjectRightsRequestHistory struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identity of the user who changed the  subject rights request.
    changedBy IdentitySetable
    // Data and time when the entity was changed.
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The stage when the entity was changed. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue, approval. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: approval.
    stage *SubjectRightsRequestStage
    // The status of the stage when the entity was changed. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
    stageStatus *SubjectRightsRequestStageStatus
    // Type of history.
    typeEscaped *string
}
// NewSubjectRightsRequestHistory instantiates a new SubjectRightsRequestHistory and sets the default values.
func NewSubjectRightsRequestHistory()(*SubjectRightsRequestHistory) {
    m := &SubjectRightsRequestHistory{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSubjectRightsRequestHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubjectRightsRequestHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequestHistory(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SubjectRightsRequestHistory) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChangedBy gets the changedBy property value. Identity of the user who changed the  subject rights request.
// returns a IdentitySetable when successful
func (m *SubjectRightsRequestHistory) GetChangedBy()(IdentitySetable) {
    return m.changedBy
}
// GetEventDateTime gets the eventDateTime property value. Data and time when the entity was changed.
// returns a *Time when successful
func (m *SubjectRightsRequestHistory) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.eventDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubjectRightsRequestHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["changedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["eventDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
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
    res["stage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*SubjectRightsRequestStage))
        }
        return nil
    }
    res["stageStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStageStatus(val.(*SubjectRightsRequestStageStatus))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SubjectRightsRequestHistory) GetOdataType()(*string) {
    return m.odataType
}
// GetStage gets the stage property value. The stage when the entity was changed. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue, approval. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: approval.
// returns a *SubjectRightsRequestStage when successful
func (m *SubjectRightsRequestHistory) GetStage()(*SubjectRightsRequestStage) {
    return m.stage
}
// GetStageStatus gets the stageStatus property value. The status of the stage when the entity was changed. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
// returns a *SubjectRightsRequestStageStatus when successful
func (m *SubjectRightsRequestHistory) GetStageStatus()(*SubjectRightsRequestStageStatus) {
    return m.stageStatus
}
// GetTypeEscaped gets the type property value. Type of history.
// returns a *string when successful
func (m *SubjectRightsRequestHistory) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *SubjectRightsRequestHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("changedBy", m.GetChangedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
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
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
        err := writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStageStatus() != nil {
        cast := (*m.GetStageStatus()).String()
        err := writer.WriteStringValue("stageStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *SubjectRightsRequestHistory) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChangedBy sets the changedBy property value. Identity of the user who changed the  subject rights request.
func (m *SubjectRightsRequestHistory) SetChangedBy(value IdentitySetable)() {
    m.changedBy = value
}
// SetEventDateTime sets the eventDateTime property value. Data and time when the entity was changed.
func (m *SubjectRightsRequestHistory) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SubjectRightsRequestHistory) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStage sets the stage property value. The stage when the entity was changed. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue, approval. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: approval.
func (m *SubjectRightsRequestHistory) SetStage(value *SubjectRightsRequestStage)() {
    m.stage = value
}
// SetStageStatus sets the stageStatus property value. The status of the stage when the entity was changed. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
func (m *SubjectRightsRequestHistory) SetStageStatus(value *SubjectRightsRequestStageStatus)() {
    m.stageStatus = value
}
// SetTypeEscaped sets the type property value. Type of history.
func (m *SubjectRightsRequestHistory) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type SubjectRightsRequestHistoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangedBy()(IdentitySetable)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetStage()(*SubjectRightsRequestStage)
    GetStageStatus()(*SubjectRightsRequestStageStatus)
    GetTypeEscaped()(*string)
    SetChangedBy(value IdentitySetable)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetStage(value *SubjectRightsRequestStage)()
    SetStageStatus(value *SubjectRightsRequestStageStatus)()
    SetTypeEscaped(value *string)()
}
