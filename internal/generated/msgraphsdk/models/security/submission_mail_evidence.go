package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubmissionMailEvidence struct {
    AlertEvidence
    // The networkMessageId property
    networkMessageId *string
    // The recipient property
    recipient *string
    // The reportType property
    reportType *string
    // The sender property
    sender *string
    // The senderIp property
    senderIp *string
    // The subject property
    subject *string
    // The submissionDateTime property
    submissionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The submissionId property
    submissionId *string
    // The submitter property
    submitter *string
}
// NewSubmissionMailEvidence instantiates a new SubmissionMailEvidence and sets the default values.
func NewSubmissionMailEvidence()(*SubmissionMailEvidence) {
    m := &SubmissionMailEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.submissionMailEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSubmissionMailEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubmissionMailEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubmissionMailEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubmissionMailEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["networkMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkMessageId(val)
        }
        return nil
    }
    res["recipient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipient(val)
        }
        return nil
    }
    res["reportType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportType(val)
        }
        return nil
    }
    res["sender"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSender(val)
        }
        return nil
    }
    res["senderIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderIp(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["submissionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmissionDateTime(val)
        }
        return nil
    }
    res["submissionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmissionId(val)
        }
        return nil
    }
    res["submitter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmitter(val)
        }
        return nil
    }
    return res
}
// GetNetworkMessageId gets the networkMessageId property value. The networkMessageId property
// returns a *string when successful
func (m *SubmissionMailEvidence) GetNetworkMessageId()(*string) {
    return m.networkMessageId
}
// GetRecipient gets the recipient property value. The recipient property
// returns a *string when successful
func (m *SubmissionMailEvidence) GetRecipient()(*string) {
    return m.recipient
}
// GetReportType gets the reportType property value. The reportType property
// returns a *string when successful
func (m *SubmissionMailEvidence) GetReportType()(*string) {
    return m.reportType
}
// GetSender gets the sender property value. The sender property
// returns a *string when successful
func (m *SubmissionMailEvidence) GetSender()(*string) {
    return m.sender
}
// GetSenderIp gets the senderIp property value. The senderIp property
// returns a *string when successful
func (m *SubmissionMailEvidence) GetSenderIp()(*string) {
    return m.senderIp
}
// GetSubject gets the subject property value. The subject property
// returns a *string when successful
func (m *SubmissionMailEvidence) GetSubject()(*string) {
    return m.subject
}
// GetSubmissionDateTime gets the submissionDateTime property value. The submissionDateTime property
// returns a *Time when successful
func (m *SubmissionMailEvidence) GetSubmissionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.submissionDateTime
}
// GetSubmissionId gets the submissionId property value. The submissionId property
// returns a *string when successful
func (m *SubmissionMailEvidence) GetSubmissionId()(*string) {
    return m.submissionId
}
// GetSubmitter gets the submitter property value. The submitter property
// returns a *string when successful
func (m *SubmissionMailEvidence) GetSubmitter()(*string) {
    return m.submitter
}
// Serialize serializes information the current object
func (m *SubmissionMailEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("networkMessageId", m.GetNetworkMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipient", m.GetRecipient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportType", m.GetReportType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sender", m.GetSender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderIp", m.GetSenderIp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("submissionDateTime", m.GetSubmissionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("submissionId", m.GetSubmissionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("submitter", m.GetSubmitter())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNetworkMessageId sets the networkMessageId property value. The networkMessageId property
func (m *SubmissionMailEvidence) SetNetworkMessageId(value *string)() {
    m.networkMessageId = value
}
// SetRecipient sets the recipient property value. The recipient property
func (m *SubmissionMailEvidence) SetRecipient(value *string)() {
    m.recipient = value
}
// SetReportType sets the reportType property value. The reportType property
func (m *SubmissionMailEvidence) SetReportType(value *string)() {
    m.reportType = value
}
// SetSender sets the sender property value. The sender property
func (m *SubmissionMailEvidence) SetSender(value *string)() {
    m.sender = value
}
// SetSenderIp sets the senderIp property value. The senderIp property
func (m *SubmissionMailEvidence) SetSenderIp(value *string)() {
    m.senderIp = value
}
// SetSubject sets the subject property value. The subject property
func (m *SubmissionMailEvidence) SetSubject(value *string)() {
    m.subject = value
}
// SetSubmissionDateTime sets the submissionDateTime property value. The submissionDateTime property
func (m *SubmissionMailEvidence) SetSubmissionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.submissionDateTime = value
}
// SetSubmissionId sets the submissionId property value. The submissionId property
func (m *SubmissionMailEvidence) SetSubmissionId(value *string)() {
    m.submissionId = value
}
// SetSubmitter sets the submitter property value. The submitter property
func (m *SubmissionMailEvidence) SetSubmitter(value *string)() {
    m.submitter = value
}
type SubmissionMailEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNetworkMessageId()(*string)
    GetRecipient()(*string)
    GetReportType()(*string)
    GetSender()(*string)
    GetSenderIp()(*string)
    GetSubject()(*string)
    GetSubmissionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubmissionId()(*string)
    GetSubmitter()(*string)
    SetNetworkMessageId(value *string)()
    SetRecipient(value *string)()
    SetReportType(value *string)()
    SetSender(value *string)()
    SetSenderIp(value *string)()
    SetSubject(value *string)()
    SetSubmissionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubmissionId(value *string)()
    SetSubmitter(value *string)()
}
