package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type CaseOperation struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // The type of action the operation represents. Possible values are: contentExport,  applyTags, convertToPdf, index, estimateStatistics, addToReviewSet, holdUpdate, unknownFutureValue, purgeData, exportReport, exportResult. You must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: purgeData, exportReport, exportResult.
    action *CaseAction
    // The date and time the operation was completed.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user that created the operation.
    createdBy ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable
    // The date and time the operation was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The progress of the operation.
    percentProgress *int32
    // Contains success and failure-specific result information.
    resultInfo ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ResultInfoable
    // The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
    status *CaseOperationStatus
}
// NewCaseOperation instantiates a new CaseOperation and sets the default values.
func NewCaseOperation()(*CaseOperation) {
    m := &CaseOperation{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateCaseOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCaseOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.security.ediscoveryAddToReviewSetOperation":
                        return NewEdiscoveryAddToReviewSetOperation(), nil
                    case "#microsoft.graph.security.ediscoveryEstimateOperation":
                        return NewEdiscoveryEstimateOperation(), nil
                    case "#microsoft.graph.security.ediscoveryExportOperation":
                        return NewEdiscoveryExportOperation(), nil
                    case "#microsoft.graph.security.ediscoveryHoldOperation":
                        return NewEdiscoveryHoldOperation(), nil
                    case "#microsoft.graph.security.ediscoveryIndexOperation":
                        return NewEdiscoveryIndexOperation(), nil
                    case "#microsoft.graph.security.ediscoveryPurgeDataOperation":
                        return NewEdiscoveryPurgeDataOperation(), nil
                    case "#microsoft.graph.security.ediscoverySearchExportOperation":
                        return NewEdiscoverySearchExportOperation(), nil
                    case "#microsoft.graph.security.ediscoveryTagOperation":
                        return NewEdiscoveryTagOperation(), nil
                }
            }
        }
    }
    return NewCaseOperation(), nil
}
// GetAction gets the action property value. The type of action the operation represents. Possible values are: contentExport,  applyTags, convertToPdf, index, estimateStatistics, addToReviewSet, holdUpdate, unknownFutureValue, purgeData, exportReport, exportResult. You must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: purgeData, exportReport, exportResult.
// returns a *CaseAction when successful
func (m *CaseOperation) GetAction()(*CaseAction) {
    return m.action
}
// GetCompletedDateTime gets the completedDateTime property value. The date and time the operation was completed.
// returns a *Time when successful
func (m *CaseOperation) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetCreatedBy gets the createdBy property value. The user that created the operation.
// returns a IdentitySetable when successful
func (m *CaseOperation) GetCreatedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the operation was created.
// returns a *Time when successful
func (m *CaseOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CaseOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCaseAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*CaseAction))
        }
        return nil
    }
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["percentProgress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentProgress(val)
        }
        return nil
    }
    res["resultInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateResultInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultInfo(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ResultInfoable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCaseOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CaseOperationStatus))
        }
        return nil
    }
    return res
}
// GetPercentProgress gets the percentProgress property value. The progress of the operation.
// returns a *int32 when successful
func (m *CaseOperation) GetPercentProgress()(*int32) {
    return m.percentProgress
}
// GetResultInfo gets the resultInfo property value. Contains success and failure-specific result information.
// returns a ResultInfoable when successful
func (m *CaseOperation) GetResultInfo()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ResultInfoable) {
    return m.resultInfo
}
// GetStatus gets the status property value. The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
// returns a *CaseOperationStatus when successful
func (m *CaseOperation) GetStatus()(*CaseOperationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CaseOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("percentProgress", m.GetPercentProgress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resultInfo", m.GetResultInfo())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The type of action the operation represents. Possible values are: contentExport,  applyTags, convertToPdf, index, estimateStatistics, addToReviewSet, holdUpdate, unknownFutureValue, purgeData, exportReport, exportResult. You must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: purgeData, exportReport, exportResult.
func (m *CaseOperation) SetAction(value *CaseAction)() {
    m.action = value
}
// SetCompletedDateTime sets the completedDateTime property value. The date and time the operation was completed.
func (m *CaseOperation) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetCreatedBy sets the createdBy property value. The user that created the operation.
func (m *CaseOperation) SetCreatedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the operation was created.
func (m *CaseOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetPercentProgress sets the percentProgress property value. The progress of the operation.
func (m *CaseOperation) SetPercentProgress(value *int32)() {
    m.percentProgress = value
}
// SetResultInfo sets the resultInfo property value. Contains success and failure-specific result information.
func (m *CaseOperation) SetResultInfo(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ResultInfoable)() {
    m.resultInfo = value
}
// SetStatus sets the status property value. The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
func (m *CaseOperation) SetStatus(value *CaseOperationStatus)() {
    m.status = value
}
type CaseOperationable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*CaseAction)
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPercentProgress()(*int32)
    GetResultInfo()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ResultInfoable)
    GetStatus()(*CaseOperationStatus)
    SetAction(value *CaseAction)()
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPercentProgress(value *int32)()
    SetResultInfo(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ResultInfoable)()
    SetStatus(value *CaseOperationStatus)()
}
