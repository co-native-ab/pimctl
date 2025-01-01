package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TeamsAppDefinition struct {
    Entity
    // Authorization requirements specified in the Teams app manifest.
    authorization TeamsAppAuthorizationable
    // The details of the bot specified in the Teams app manifest.
    bot TeamworkBotable
    // The createdBy property
    createdBy IdentitySetable
    // Verbose description of the application.
    description *string
    // The name of the app provided by the app developer.
    displayName *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The published status of a specific version of a Teams app. Possible values are:submitted—The specific version of the Teams app was submitted and is under review.published—The request to publish the specific version of the Teams app was approved by the admin and the app is published.rejected—The admin rejected the request to publish the specific version of the Teams app.
    publishingState *TeamsAppPublishingState
    // Short description of the application.
    shortDescription *string
    // The ID from the Teams app manifest.
    teamsAppId *string
    // The version number of the application.
    version *string
}
// NewTeamsAppDefinition instantiates a new TeamsAppDefinition and sets the default values.
func NewTeamsAppDefinition()(*TeamsAppDefinition) {
    m := &TeamsAppDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAppDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamsAppDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppDefinition(), nil
}
// GetAuthorization gets the authorization property value. Authorization requirements specified in the Teams app manifest.
// returns a TeamsAppAuthorizationable when successful
func (m *TeamsAppDefinition) GetAuthorization()(TeamsAppAuthorizationable) {
    return m.authorization
}
// GetBot gets the bot property value. The details of the bot specified in the Teams app manifest.
// returns a TeamworkBotable when successful
func (m *TeamsAppDefinition) GetBot()(TeamworkBotable) {
    return m.bot
}
// GetCreatedBy gets the createdBy property value. The createdBy property
// returns a IdentitySetable when successful
func (m *TeamsAppDefinition) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetDescription gets the description property value. Verbose description of the application.
// returns a *string when successful
func (m *TeamsAppDefinition) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of the app provided by the app developer.
// returns a *string when successful
func (m *TeamsAppDefinition) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamsAppDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authorization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppAuthorizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorization(val.(TeamsAppAuthorizationable))
        }
        return nil
    }
    res["bot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkBotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBot(val.(TeamworkBotable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
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
    res["publishingState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppPublishingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingState(val.(*TeamsAppPublishingState))
        }
        return nil
    }
    res["shortDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortDescription(val)
        }
        return nil
    }
    res["teamsAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppId(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
// returns a *Time when successful
func (m *TeamsAppDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPublishingState gets the publishingState property value. The published status of a specific version of a Teams app. Possible values are:submitted—The specific version of the Teams app was submitted and is under review.published—The request to publish the specific version of the Teams app was approved by the admin and the app is published.rejected—The admin rejected the request to publish the specific version of the Teams app.
// returns a *TeamsAppPublishingState when successful
func (m *TeamsAppDefinition) GetPublishingState()(*TeamsAppPublishingState) {
    return m.publishingState
}
// GetShortDescription gets the shortDescription property value. Short description of the application.
// returns a *string when successful
func (m *TeamsAppDefinition) GetShortDescription()(*string) {
    return m.shortDescription
}
// GetTeamsAppId gets the teamsAppId property value. The ID from the Teams app manifest.
// returns a *string when successful
func (m *TeamsAppDefinition) GetTeamsAppId()(*string) {
    return m.teamsAppId
}
// GetVersion gets the version property value. The version number of the application.
// returns a *string when successful
func (m *TeamsAppDefinition) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *TeamsAppDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("authorization", m.GetAuthorization())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bot", m.GetBot())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPublishingState() != nil {
        cast := (*m.GetPublishingState()).String()
        err = writer.WriteStringValue("publishingState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shortDescription", m.GetShortDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamsAppId", m.GetTeamsAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorization sets the authorization property value. Authorization requirements specified in the Teams app manifest.
func (m *TeamsAppDefinition) SetAuthorization(value TeamsAppAuthorizationable)() {
    m.authorization = value
}
// SetBot sets the bot property value. The details of the bot specified in the Teams app manifest.
func (m *TeamsAppDefinition) SetBot(value TeamworkBotable)() {
    m.bot = value
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *TeamsAppDefinition) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetDescription sets the description property value. Verbose description of the application.
func (m *TeamsAppDefinition) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of the app provided by the app developer.
func (m *TeamsAppDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *TeamsAppDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPublishingState sets the publishingState property value. The published status of a specific version of a Teams app. Possible values are:submitted—The specific version of the Teams app was submitted and is under review.published—The request to publish the specific version of the Teams app was approved by the admin and the app is published.rejected—The admin rejected the request to publish the specific version of the Teams app.
func (m *TeamsAppDefinition) SetPublishingState(value *TeamsAppPublishingState)() {
    m.publishingState = value
}
// SetShortDescription sets the shortDescription property value. Short description of the application.
func (m *TeamsAppDefinition) SetShortDescription(value *string)() {
    m.shortDescription = value
}
// SetTeamsAppId sets the teamsAppId property value. The ID from the Teams app manifest.
func (m *TeamsAppDefinition) SetTeamsAppId(value *string)() {
    m.teamsAppId = value
}
// SetVersion sets the version property value. The version number of the application.
func (m *TeamsAppDefinition) SetVersion(value *string)() {
    m.version = value
}
type TeamsAppDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorization()(TeamsAppAuthorizationable)
    GetBot()(TeamworkBotable)
    GetCreatedBy()(IdentitySetable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPublishingState()(*TeamsAppPublishingState)
    GetShortDescription()(*string)
    GetTeamsAppId()(*string)
    GetVersion()(*string)
    SetAuthorization(value TeamsAppAuthorizationable)()
    SetBot(value TeamworkBotable)()
    SetCreatedBy(value IdentitySetable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPublishingState(value *TeamsAppPublishingState)()
    SetShortDescription(value *string)()
    SetTeamsAppId(value *string)()
    SetVersion(value *string)()
}
