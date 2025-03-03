package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Simulation struct {
    Entity
    // The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment, linkToMalwareFile, unknownFutureValue, oAuthConsentGrant. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: oAuthConsentGrant. For more information on the types of social engineering attack techniques, see simulations.
    attackTechnique *SimulationAttackTechnique
    // Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, social, cloud, endpoint, unknownFutureValue.
    attackType *SimulationAttackType
    // Unique identifier for the attack simulation automation.
    automationId *string
    // Date and time of completion of the attack simulation and training campaign. Supports $filter and $orderby.
    completionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Identity of the user who created the attack simulation and training campaign.
    createdBy EmailIdentityable
    // Date and time of creation of the attack simulation and training campaign.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description of the attack simulation and training campaign.
    description *string
    // Display name of the attack simulation and training campaign. Supports $filter and $orderby.
    displayName *string
    // Simulation duration in days.
    durationInDays *int32
    // Details about the end user notification setting.
    endUserNotificationSetting EndUserNotificationSettingable
    // Users excluded from the simulation.
    excludedAccountTarget AccountTargetContentable
    // Users targeted in the simulation.
    includedAccountTarget AccountTargetContentable
    // Flag that represents if the attack simulation and training campaign was created from a simulation automation flow. Supports $filter and $orderby.
    isAutomated *bool
    // The landing page associated with a simulation during its creation.
    landingPage LandingPageable
    // Identity of the user who most recently modified the attack simulation and training campaign.
    lastModifiedBy EmailIdentityable
    // Date and time of the most recent modification of the attack simulation and training campaign.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date and time of the launch/start of the attack simulation and training campaign. Supports $filter and $orderby.
    launchDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The login page associated with a simulation during its creation.
    loginPage LoginPageable
    // OAuth app details for the OAuth technique.
    oAuthConsentAppDetail OAuthConsentAppDetailable
    // The payload associated with a simulation during its creation.
    payload Payloadable
    // Method of delivery of the phishing payload used in the attack simulation and training campaign. Possible values are: unknown, sms, email, teams, unknownFutureValue.
    payloadDeliveryPlatform *PayloadDeliveryPlatform
    // Report of the attack simulation and training campaign.
    report SimulationReportable
    // Status of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, draft, running, scheduled, succeeded, failed, cancelled, excluded, unknownFutureValue.
    status *SimulationStatus
    // Details about the training settings for a simulation.
    trainingSetting TrainingSettingable
}
// NewSimulation instantiates a new Simulation and sets the default values.
func NewSimulation()(*Simulation) {
    m := &Simulation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSimulationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSimulationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimulation(), nil
}
// GetAttackTechnique gets the attackTechnique property value. The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment, linkToMalwareFile, unknownFutureValue, oAuthConsentGrant. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: oAuthConsentGrant. For more information on the types of social engineering attack techniques, see simulations.
// returns a *SimulationAttackTechnique when successful
func (m *Simulation) GetAttackTechnique()(*SimulationAttackTechnique) {
    return m.attackTechnique
}
// GetAttackType gets the attackType property value. Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, social, cloud, endpoint, unknownFutureValue.
// returns a *SimulationAttackType when successful
func (m *Simulation) GetAttackType()(*SimulationAttackType) {
    return m.attackType
}
// GetAutomationId gets the automationId property value. Unique identifier for the attack simulation automation.
// returns a *string when successful
func (m *Simulation) GetAutomationId()(*string) {
    return m.automationId
}
// GetCompletionDateTime gets the completionDateTime property value. Date and time of completion of the attack simulation and training campaign. Supports $filter and $orderby.
// returns a *Time when successful
func (m *Simulation) GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completionDateTime
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the attack simulation and training campaign.
// returns a EmailIdentityable when successful
func (m *Simulation) GetCreatedBy()(EmailIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of creation of the attack simulation and training campaign.
// returns a *Time when successful
func (m *Simulation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description of the attack simulation and training campaign.
// returns a *string when successful
func (m *Simulation) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name of the attack simulation and training campaign. Supports $filter and $orderby.
// returns a *string when successful
func (m *Simulation) GetDisplayName()(*string) {
    return m.displayName
}
// GetDurationInDays gets the durationInDays property value. Simulation duration in days.
// returns a *int32 when successful
func (m *Simulation) GetDurationInDays()(*int32) {
    return m.durationInDays
}
// GetEndUserNotificationSetting gets the endUserNotificationSetting property value. Details about the end user notification setting.
// returns a EndUserNotificationSettingable when successful
func (m *Simulation) GetEndUserNotificationSetting()(EndUserNotificationSettingable) {
    return m.endUserNotificationSetting
}
// GetExcludedAccountTarget gets the excludedAccountTarget property value. Users excluded from the simulation.
// returns a AccountTargetContentable when successful
func (m *Simulation) GetExcludedAccountTarget()(AccountTargetContentable) {
    return m.excludedAccountTarget
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Simulation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attackTechnique"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAttackTechnique)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackTechnique(val.(*SimulationAttackTechnique))
        }
        return nil
    }
    res["attackType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAttackType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackType(val.(*SimulationAttackType))
        }
        return nil
    }
    res["automationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomationId(val)
        }
        return nil
    }
    res["completionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionDateTime(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(EmailIdentityable))
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
    res["durationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInDays(val)
        }
        return nil
    }
    res["endUserNotificationSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEndUserNotificationSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndUserNotificationSetting(val.(EndUserNotificationSettingable))
        }
        return nil
    }
    res["excludedAccountTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountTargetContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludedAccountTarget(val.(AccountTargetContentable))
        }
        return nil
    }
    res["includedAccountTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountTargetContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludedAccountTarget(val.(AccountTargetContentable))
        }
        return nil
    }
    res["isAutomated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAutomated(val)
        }
        return nil
    }
    res["landingPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLandingPageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLandingPage(val.(LandingPageable))
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(EmailIdentityable))
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
    res["launchDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLaunchDateTime(val)
        }
        return nil
    }
    res["loginPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLoginPageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginPage(val.(LoginPageable))
        }
        return nil
    }
    res["oAuthConsentAppDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOAuthConsentAppDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOAuthConsentAppDetail(val.(OAuthConsentAppDetailable))
        }
        return nil
    }
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePayloadFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val.(Payloadable))
        }
        return nil
    }
    res["payloadDeliveryPlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadDeliveryPlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadDeliveryPlatform(val.(*PayloadDeliveryPlatform))
        }
        return nil
    }
    res["report"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimulationReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReport(val.(SimulationReportable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SimulationStatus))
        }
        return nil
    }
    res["trainingSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTrainingSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingSetting(val.(TrainingSettingable))
        }
        return nil
    }
    return res
}
// GetIncludedAccountTarget gets the includedAccountTarget property value. Users targeted in the simulation.
// returns a AccountTargetContentable when successful
func (m *Simulation) GetIncludedAccountTarget()(AccountTargetContentable) {
    return m.includedAccountTarget
}
// GetIsAutomated gets the isAutomated property value. Flag that represents if the attack simulation and training campaign was created from a simulation automation flow. Supports $filter and $orderby.
// returns a *bool when successful
func (m *Simulation) GetIsAutomated()(*bool) {
    return m.isAutomated
}
// GetLandingPage gets the landingPage property value. The landing page associated with a simulation during its creation.
// returns a LandingPageable when successful
func (m *Simulation) GetLandingPage()(LandingPageable) {
    return m.landingPage
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who most recently modified the attack simulation and training campaign.
// returns a EmailIdentityable when successful
func (m *Simulation) GetLastModifiedBy()(EmailIdentityable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time of the most recent modification of the attack simulation and training campaign.
// returns a *Time when successful
func (m *Simulation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLaunchDateTime gets the launchDateTime property value. Date and time of the launch/start of the attack simulation and training campaign. Supports $filter and $orderby.
// returns a *Time when successful
func (m *Simulation) GetLaunchDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.launchDateTime
}
// GetLoginPage gets the loginPage property value. The login page associated with a simulation during its creation.
// returns a LoginPageable when successful
func (m *Simulation) GetLoginPage()(LoginPageable) {
    return m.loginPage
}
// GetOAuthConsentAppDetail gets the oAuthConsentAppDetail property value. OAuth app details for the OAuth technique.
// returns a OAuthConsentAppDetailable when successful
func (m *Simulation) GetOAuthConsentAppDetail()(OAuthConsentAppDetailable) {
    return m.oAuthConsentAppDetail
}
// GetPayload gets the payload property value. The payload associated with a simulation during its creation.
// returns a Payloadable when successful
func (m *Simulation) GetPayload()(Payloadable) {
    return m.payload
}
// GetPayloadDeliveryPlatform gets the payloadDeliveryPlatform property value. Method of delivery of the phishing payload used in the attack simulation and training campaign. Possible values are: unknown, sms, email, teams, unknownFutureValue.
// returns a *PayloadDeliveryPlatform when successful
func (m *Simulation) GetPayloadDeliveryPlatform()(*PayloadDeliveryPlatform) {
    return m.payloadDeliveryPlatform
}
// GetReport gets the report property value. Report of the attack simulation and training campaign.
// returns a SimulationReportable when successful
func (m *Simulation) GetReport()(SimulationReportable) {
    return m.report
}
// GetStatus gets the status property value. Status of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, draft, running, scheduled, succeeded, failed, cancelled, excluded, unknownFutureValue.
// returns a *SimulationStatus when successful
func (m *Simulation) GetStatus()(*SimulationStatus) {
    return m.status
}
// GetTrainingSetting gets the trainingSetting property value. Details about the training settings for a simulation.
// returns a TrainingSettingable when successful
func (m *Simulation) GetTrainingSetting()(TrainingSettingable) {
    return m.trainingSetting
}
// Serialize serializes information the current object
func (m *Simulation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttackTechnique() != nil {
        cast := (*m.GetAttackTechnique()).String()
        err = writer.WriteStringValue("attackTechnique", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAttackType() != nil {
        cast := (*m.GetAttackType()).String()
        err = writer.WriteStringValue("attackType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("automationId", m.GetAutomationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completionDateTime", m.GetCompletionDateTime())
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
    {
        err = writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("endUserNotificationSetting", m.GetEndUserNotificationSetting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("excludedAccountTarget", m.GetExcludedAccountTarget())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("includedAccountTarget", m.GetIncludedAccountTarget())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAutomated", m.GetIsAutomated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("landingPage", m.GetLandingPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
    {
        err = writer.WriteTimeValue("launchDateTime", m.GetLaunchDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("loginPage", m.GetLoginPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("oAuthConsentAppDetail", m.GetOAuthConsentAppDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    if m.GetPayloadDeliveryPlatform() != nil {
        cast := (*m.GetPayloadDeliveryPlatform()).String()
        err = writer.WriteStringValue("payloadDeliveryPlatform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("report", m.GetReport())
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
    {
        err = writer.WriteObjectValue("trainingSetting", m.GetTrainingSetting())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttackTechnique sets the attackTechnique property value. The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment, linkToMalwareFile, unknownFutureValue, oAuthConsentGrant. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: oAuthConsentGrant. For more information on the types of social engineering attack techniques, see simulations.
func (m *Simulation) SetAttackTechnique(value *SimulationAttackTechnique)() {
    m.attackTechnique = value
}
// SetAttackType sets the attackType property value. Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, social, cloud, endpoint, unknownFutureValue.
func (m *Simulation) SetAttackType(value *SimulationAttackType)() {
    m.attackType = value
}
// SetAutomationId sets the automationId property value. Unique identifier for the attack simulation automation.
func (m *Simulation) SetAutomationId(value *string)() {
    m.automationId = value
}
// SetCompletionDateTime sets the completionDateTime property value. Date and time of completion of the attack simulation and training campaign. Supports $filter and $orderby.
func (m *Simulation) SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completionDateTime = value
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the attack simulation and training campaign.
func (m *Simulation) SetCreatedBy(value EmailIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of creation of the attack simulation and training campaign.
func (m *Simulation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description of the attack simulation and training campaign.
func (m *Simulation) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name of the attack simulation and training campaign. Supports $filter and $orderby.
func (m *Simulation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDurationInDays sets the durationInDays property value. Simulation duration in days.
func (m *Simulation) SetDurationInDays(value *int32)() {
    m.durationInDays = value
}
// SetEndUserNotificationSetting sets the endUserNotificationSetting property value. Details about the end user notification setting.
func (m *Simulation) SetEndUserNotificationSetting(value EndUserNotificationSettingable)() {
    m.endUserNotificationSetting = value
}
// SetExcludedAccountTarget sets the excludedAccountTarget property value. Users excluded from the simulation.
func (m *Simulation) SetExcludedAccountTarget(value AccountTargetContentable)() {
    m.excludedAccountTarget = value
}
// SetIncludedAccountTarget sets the includedAccountTarget property value. Users targeted in the simulation.
func (m *Simulation) SetIncludedAccountTarget(value AccountTargetContentable)() {
    m.includedAccountTarget = value
}
// SetIsAutomated sets the isAutomated property value. Flag that represents if the attack simulation and training campaign was created from a simulation automation flow. Supports $filter and $orderby.
func (m *Simulation) SetIsAutomated(value *bool)() {
    m.isAutomated = value
}
// SetLandingPage sets the landingPage property value. The landing page associated with a simulation during its creation.
func (m *Simulation) SetLandingPage(value LandingPageable)() {
    m.landingPage = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who most recently modified the attack simulation and training campaign.
func (m *Simulation) SetLastModifiedBy(value EmailIdentityable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time of the most recent modification of the attack simulation and training campaign.
func (m *Simulation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLaunchDateTime sets the launchDateTime property value. Date and time of the launch/start of the attack simulation and training campaign. Supports $filter and $orderby.
func (m *Simulation) SetLaunchDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.launchDateTime = value
}
// SetLoginPage sets the loginPage property value. The login page associated with a simulation during its creation.
func (m *Simulation) SetLoginPage(value LoginPageable)() {
    m.loginPage = value
}
// SetOAuthConsentAppDetail sets the oAuthConsentAppDetail property value. OAuth app details for the OAuth technique.
func (m *Simulation) SetOAuthConsentAppDetail(value OAuthConsentAppDetailable)() {
    m.oAuthConsentAppDetail = value
}
// SetPayload sets the payload property value. The payload associated with a simulation during its creation.
func (m *Simulation) SetPayload(value Payloadable)() {
    m.payload = value
}
// SetPayloadDeliveryPlatform sets the payloadDeliveryPlatform property value. Method of delivery of the phishing payload used in the attack simulation and training campaign. Possible values are: unknown, sms, email, teams, unknownFutureValue.
func (m *Simulation) SetPayloadDeliveryPlatform(value *PayloadDeliveryPlatform)() {
    m.payloadDeliveryPlatform = value
}
// SetReport sets the report property value. Report of the attack simulation and training campaign.
func (m *Simulation) SetReport(value SimulationReportable)() {
    m.report = value
}
// SetStatus sets the status property value. Status of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown, draft, running, scheduled, succeeded, failed, cancelled, excluded, unknownFutureValue.
func (m *Simulation) SetStatus(value *SimulationStatus)() {
    m.status = value
}
// SetTrainingSetting sets the trainingSetting property value. Details about the training settings for a simulation.
func (m *Simulation) SetTrainingSetting(value TrainingSettingable)() {
    m.trainingSetting = value
}
type Simulationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttackTechnique()(*SimulationAttackTechnique)
    GetAttackType()(*SimulationAttackType)
    GetAutomationId()(*string)
    GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedBy()(EmailIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDurationInDays()(*int32)
    GetEndUserNotificationSetting()(EndUserNotificationSettingable)
    GetExcludedAccountTarget()(AccountTargetContentable)
    GetIncludedAccountTarget()(AccountTargetContentable)
    GetIsAutomated()(*bool)
    GetLandingPage()(LandingPageable)
    GetLastModifiedBy()(EmailIdentityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLaunchDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLoginPage()(LoginPageable)
    GetOAuthConsentAppDetail()(OAuthConsentAppDetailable)
    GetPayload()(Payloadable)
    GetPayloadDeliveryPlatform()(*PayloadDeliveryPlatform)
    GetReport()(SimulationReportable)
    GetStatus()(*SimulationStatus)
    GetTrainingSetting()(TrainingSettingable)
    SetAttackTechnique(value *SimulationAttackTechnique)()
    SetAttackType(value *SimulationAttackType)()
    SetAutomationId(value *string)()
    SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedBy(value EmailIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDurationInDays(value *int32)()
    SetEndUserNotificationSetting(value EndUserNotificationSettingable)()
    SetExcludedAccountTarget(value AccountTargetContentable)()
    SetIncludedAccountTarget(value AccountTargetContentable)()
    SetIsAutomated(value *bool)()
    SetLandingPage(value LandingPageable)()
    SetLastModifiedBy(value EmailIdentityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLaunchDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLoginPage(value LoginPageable)()
    SetOAuthConsentAppDetail(value OAuthConsentAppDetailable)()
    SetPayload(value Payloadable)()
    SetPayloadDeliveryPlatform(value *PayloadDeliveryPlatform)()
    SetReport(value SimulationReportable)()
    SetStatus(value *SimulationStatus)()
    SetTrainingSetting(value TrainingSettingable)()
}
