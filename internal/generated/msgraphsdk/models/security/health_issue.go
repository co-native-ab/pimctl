package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type HealthIssue struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // Contains additional information about the issue, such as a list of items to fix.
    additionalInformation []string
    // The date and time when the health issue was generated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Contains more detailed information about the health issue.
    description *string
    // The display name of the health issue.
    displayName *string
    // A list of the fully qualified domain names of the domains or the sensors the health issue is related to.
    domainNames []string
    // The type of the health issue. The possible values are: sensor, global, unknownFutureValue. For a list of all health issues and their identifiers, see Microsoft Defender for Identity health issues.
    healthIssueType *HealthIssueType
    // The type identifier of the health issue. For a list of all health issues and their identifiers, see Microsoft Defender for Identity health issues.
    issueTypeId *string
    // The date and time when the health issue was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A list of recommended actions that can be taken to resolve the issue effectively and efficiently. These actions might include instructions for further investigation and aren't limited to prewritten responses.
    recommendations []string
    // A list of commands from the PowerShell module for the product that can be used to resolve the issue, if available. If no commands can be used to solve the issue, this property is empty. The commands, if present, provide a quick and efficient way to address the issue. These commands run in sequence for the single recommended fix.
    recommendedActionCommands []string
    // A list of the DNS names of the sensors the health issue is related to.
    sensorDNSNames []string
    // The severity of the health issue. The possible values are: low, medium, high, unknownFutureValue.
    severity *HealthIssueSeverity
    // The status of the health issue. The possible values are: open, closed, suppressed, unknownFutureValue.
    status *HealthIssueStatus
}
// NewHealthIssue instantiates a new HealthIssue and sets the default values.
func NewHealthIssue()(*HealthIssue) {
    m := &HealthIssue{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateHealthIssueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHealthIssueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHealthIssue(), nil
}
// GetAdditionalInformation gets the additionalInformation property value. Contains additional information about the issue, such as a list of items to fix.
// returns a []string when successful
func (m *HealthIssue) GetAdditionalInformation()([]string) {
    return m.additionalInformation
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the health issue was generated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *HealthIssue) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Contains more detailed information about the health issue.
// returns a *string when successful
func (m *HealthIssue) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the health issue.
// returns a *string when successful
func (m *HealthIssue) GetDisplayName()(*string) {
    return m.displayName
}
// GetDomainNames gets the domainNames property value. A list of the fully qualified domain names of the domains or the sensors the health issue is related to.
// returns a []string when successful
func (m *HealthIssue) GetDomainNames()([]string) {
    return m.domainNames
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HealthIssue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAdditionalInformation(res)
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["domainNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDomainNames(res)
        }
        return nil
    }
    res["healthIssueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHealthIssueType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthIssueType(val.(*HealthIssueType))
        }
        return nil
    }
    res["issueTypeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueTypeId(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["recommendations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRecommendations(res)
        }
        return nil
    }
    res["recommendedActionCommands"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRecommendedActionCommands(res)
        }
        return nil
    }
    res["sensorDNSNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSensorDNSNames(res)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHealthIssueSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*HealthIssueSeverity))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHealthIssueStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*HealthIssueStatus))
        }
        return nil
    }
    return res
}
// GetHealthIssueType gets the healthIssueType property value. The type of the health issue. The possible values are: sensor, global, unknownFutureValue. For a list of all health issues and their identifiers, see Microsoft Defender for Identity health issues.
// returns a *HealthIssueType when successful
func (m *HealthIssue) GetHealthIssueType()(*HealthIssueType) {
    return m.healthIssueType
}
// GetIssueTypeId gets the issueTypeId property value. The type identifier of the health issue. For a list of all health issues and their identifiers, see Microsoft Defender for Identity health issues.
// returns a *string when successful
func (m *HealthIssue) GetIssueTypeId()(*string) {
    return m.issueTypeId
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the health issue was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *HealthIssue) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRecommendations gets the recommendations property value. A list of recommended actions that can be taken to resolve the issue effectively and efficiently. These actions might include instructions for further investigation and aren't limited to prewritten responses.
// returns a []string when successful
func (m *HealthIssue) GetRecommendations()([]string) {
    return m.recommendations
}
// GetRecommendedActionCommands gets the recommendedActionCommands property value. A list of commands from the PowerShell module for the product that can be used to resolve the issue, if available. If no commands can be used to solve the issue, this property is empty. The commands, if present, provide a quick and efficient way to address the issue. These commands run in sequence for the single recommended fix.
// returns a []string when successful
func (m *HealthIssue) GetRecommendedActionCommands()([]string) {
    return m.recommendedActionCommands
}
// GetSensorDNSNames gets the sensorDNSNames property value. A list of the DNS names of the sensors the health issue is related to.
// returns a []string when successful
func (m *HealthIssue) GetSensorDNSNames()([]string) {
    return m.sensorDNSNames
}
// GetSeverity gets the severity property value. The severity of the health issue. The possible values are: low, medium, high, unknownFutureValue.
// returns a *HealthIssueSeverity when successful
func (m *HealthIssue) GetSeverity()(*HealthIssueSeverity) {
    return m.severity
}
// GetStatus gets the status property value. The status of the health issue. The possible values are: open, closed, suppressed, unknownFutureValue.
// returns a *HealthIssueStatus when successful
func (m *HealthIssue) GetStatus()(*HealthIssueStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *HealthIssue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalInformation() != nil {
        err = writer.WriteCollectionOfStringValues("additionalInformation", m.GetAdditionalInformation())
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetDomainNames() != nil {
        err = writer.WriteCollectionOfStringValues("domainNames", m.GetDomainNames())
        if err != nil {
            return err
        }
    }
    if m.GetHealthIssueType() != nil {
        cast := (*m.GetHealthIssueType()).String()
        err = writer.WriteStringValue("healthIssueType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issueTypeId", m.GetIssueTypeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRecommendations() != nil {
        err = writer.WriteCollectionOfStringValues("recommendations", m.GetRecommendations())
        if err != nil {
            return err
        }
    }
    if m.GetRecommendedActionCommands() != nil {
        err = writer.WriteCollectionOfStringValues("recommendedActionCommands", m.GetRecommendedActionCommands())
        if err != nil {
            return err
        }
    }
    if m.GetSensorDNSNames() != nil {
        err = writer.WriteCollectionOfStringValues("sensorDNSNames", m.GetSensorDNSNames())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err = writer.WriteStringValue("severity", &cast)
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
// SetAdditionalInformation sets the additionalInformation property value. Contains additional information about the issue, such as a list of items to fix.
func (m *HealthIssue) SetAdditionalInformation(value []string)() {
    m.additionalInformation = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the health issue was generated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *HealthIssue) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Contains more detailed information about the health issue.
func (m *HealthIssue) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the health issue.
func (m *HealthIssue) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDomainNames sets the domainNames property value. A list of the fully qualified domain names of the domains or the sensors the health issue is related to.
func (m *HealthIssue) SetDomainNames(value []string)() {
    m.domainNames = value
}
// SetHealthIssueType sets the healthIssueType property value. The type of the health issue. The possible values are: sensor, global, unknownFutureValue. For a list of all health issues and their identifiers, see Microsoft Defender for Identity health issues.
func (m *HealthIssue) SetHealthIssueType(value *HealthIssueType)() {
    m.healthIssueType = value
}
// SetIssueTypeId sets the issueTypeId property value. The type identifier of the health issue. For a list of all health issues and their identifiers, see Microsoft Defender for Identity health issues.
func (m *HealthIssue) SetIssueTypeId(value *string)() {
    m.issueTypeId = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the health issue was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *HealthIssue) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRecommendations sets the recommendations property value. A list of recommended actions that can be taken to resolve the issue effectively and efficiently. These actions might include instructions for further investigation and aren't limited to prewritten responses.
func (m *HealthIssue) SetRecommendations(value []string)() {
    m.recommendations = value
}
// SetRecommendedActionCommands sets the recommendedActionCommands property value. A list of commands from the PowerShell module for the product that can be used to resolve the issue, if available. If no commands can be used to solve the issue, this property is empty. The commands, if present, provide a quick and efficient way to address the issue. These commands run in sequence for the single recommended fix.
func (m *HealthIssue) SetRecommendedActionCommands(value []string)() {
    m.recommendedActionCommands = value
}
// SetSensorDNSNames sets the sensorDNSNames property value. A list of the DNS names of the sensors the health issue is related to.
func (m *HealthIssue) SetSensorDNSNames(value []string)() {
    m.sensorDNSNames = value
}
// SetSeverity sets the severity property value. The severity of the health issue. The possible values are: low, medium, high, unknownFutureValue.
func (m *HealthIssue) SetSeverity(value *HealthIssueSeverity)() {
    m.severity = value
}
// SetStatus sets the status property value. The status of the health issue. The possible values are: open, closed, suppressed, unknownFutureValue.
func (m *HealthIssue) SetStatus(value *HealthIssueStatus)() {
    m.status = value
}
type HealthIssueable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalInformation()([]string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDomainNames()([]string)
    GetHealthIssueType()(*HealthIssueType)
    GetIssueTypeId()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecommendations()([]string)
    GetRecommendedActionCommands()([]string)
    GetSensorDNSNames()([]string)
    GetSeverity()(*HealthIssueSeverity)
    GetStatus()(*HealthIssueStatus)
    SetAdditionalInformation(value []string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDomainNames(value []string)()
    SetHealthIssueType(value *HealthIssueType)()
    SetIssueTypeId(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecommendations(value []string)()
    SetRecommendedActionCommands(value []string)()
    SetSensorDNSNames(value []string)()
    SetSeverity(value *HealthIssueSeverity)()
    SetStatus(value *HealthIssueStatus)()
}
