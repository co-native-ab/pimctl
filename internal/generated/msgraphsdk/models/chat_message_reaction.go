package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ChatMessageReaction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the reaction.
    displayName *string
    // The OdataType property
    odataType *string
    // The hosted content URL for the custom reaction type.
    reactionContentUrl *string
    // The reaction type. Supported values include Unicode characters, custom, and some backward-compatible reaction types, such as like, angry, sad, laugh, heart, and surprised.
    reactionType *string
    // The user property
    user ChatMessageReactionIdentitySetable
}
// NewChatMessageReaction instantiates a new ChatMessageReaction and sets the default values.
func NewChatMessageReaction()(*ChatMessageReaction) {
    m := &ChatMessageReaction{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChatMessageReactionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateChatMessageReactionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageReaction(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ChatMessageReaction) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedDateTime gets the createdDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *ChatMessageReaction) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The name of the reaction.
// returns a *string when successful
func (m *ChatMessageReaction) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ChatMessageReaction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["reactionContentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReactionContentUrl(val)
        }
        return nil
    }
    res["reactionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReactionType(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatMessageReactionIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(ChatMessageReactionIdentitySetable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ChatMessageReaction) GetOdataType()(*string) {
    return m.odataType
}
// GetReactionContentUrl gets the reactionContentUrl property value. The hosted content URL for the custom reaction type.
// returns a *string when successful
func (m *ChatMessageReaction) GetReactionContentUrl()(*string) {
    return m.reactionContentUrl
}
// GetReactionType gets the reactionType property value. The reaction type. Supported values include Unicode characters, custom, and some backward-compatible reaction types, such as like, angry, sad, laugh, heart, and surprised.
// returns a *string when successful
func (m *ChatMessageReaction) GetReactionType()(*string) {
    return m.reactionType
}
// GetUser gets the user property value. The user property
// returns a ChatMessageReactionIdentitySetable when successful
func (m *ChatMessageReaction) GetUser()(ChatMessageReactionIdentitySetable) {
    return m.user
}
// Serialize serializes information the current object
func (m *ChatMessageReaction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err := writer.WriteStringValue("reactionContentUrl", m.GetReactionContentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reactionType", m.GetReactionType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *ChatMessageReaction) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedDateTime sets the createdDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ChatMessageReaction) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The name of the reaction.
func (m *ChatMessageReaction) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChatMessageReaction) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReactionContentUrl sets the reactionContentUrl property value. The hosted content URL for the custom reaction type.
func (m *ChatMessageReaction) SetReactionContentUrl(value *string)() {
    m.reactionContentUrl = value
}
// SetReactionType sets the reactionType property value. The reaction type. Supported values include Unicode characters, custom, and some backward-compatible reaction types, such as like, angry, sad, laugh, heart, and surprised.
func (m *ChatMessageReaction) SetReactionType(value *string)() {
    m.reactionType = value
}
// SetUser sets the user property value. The user property
func (m *ChatMessageReaction) SetUser(value ChatMessageReactionIdentitySetable)() {
    m.user = value
}
type ChatMessageReactionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetOdataType()(*string)
    GetReactionContentUrl()(*string)
    GetReactionType()(*string)
    GetUser()(ChatMessageReactionIdentitySetable)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
    SetReactionContentUrl(value *string)()
    SetReactionType(value *string)()
    SetUser(value ChatMessageReactionIdentitySetable)()
}
