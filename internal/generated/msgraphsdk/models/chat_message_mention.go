package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ChatMessageMention struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Index of an entity being mentioned in the specified chatMessage. Matches the {index} value in the corresponding <at id='{index}'> tag in the message body.
    id *int32
    // The entity (user, application, team, or channel) that was @mentioned.
    mentioned ChatMessageMentionedIdentitySetable
    // String used to represent the mention. For example, a user's display name, a team name.
    mentionText *string
    // The OdataType property
    odataType *string
}
// NewChatMessageMention instantiates a new ChatMessageMention and sets the default values.
func NewChatMessageMention()(*ChatMessageMention) {
    m := &ChatMessageMention{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChatMessageMentionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateChatMessageMentionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageMention(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ChatMessageMention) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ChatMessageMention) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["mentioned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatMessageMentionedIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentioned(val.(ChatMessageMentionedIdentitySetable))
        }
        return nil
    }
    res["mentionText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentionText(val)
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
// GetId gets the id property value. Index of an entity being mentioned in the specified chatMessage. Matches the {index} value in the corresponding <at id='{index}'> tag in the message body.
// returns a *int32 when successful
func (m *ChatMessageMention) GetId()(*int32) {
    return m.id
}
// GetMentioned gets the mentioned property value. The entity (user, application, team, or channel) that was @mentioned.
// returns a ChatMessageMentionedIdentitySetable when successful
func (m *ChatMessageMention) GetMentioned()(ChatMessageMentionedIdentitySetable) {
    return m.mentioned
}
// GetMentionText gets the mentionText property value. String used to represent the mention. For example, a user's display name, a team name.
// returns a *string when successful
func (m *ChatMessageMention) GetMentionText()(*string) {
    return m.mentionText
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ChatMessageMention) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ChatMessageMention) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mentioned", m.GetMentioned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mentionText", m.GetMentionText())
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
func (m *ChatMessageMention) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Index of an entity being mentioned in the specified chatMessage. Matches the {index} value in the corresponding <at id='{index}'> tag in the message body.
func (m *ChatMessageMention) SetId(value *int32)() {
    m.id = value
}
// SetMentioned sets the mentioned property value. The entity (user, application, team, or channel) that was @mentioned.
func (m *ChatMessageMention) SetMentioned(value ChatMessageMentionedIdentitySetable)() {
    m.mentioned = value
}
// SetMentionText sets the mentionText property value. String used to represent the mention. For example, a user's display name, a team name.
func (m *ChatMessageMention) SetMentionText(value *string)() {
    m.mentionText = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChatMessageMention) SetOdataType(value *string)() {
    m.odataType = value
}
type ChatMessageMentionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetMentioned()(ChatMessageMentionedIdentitySetable)
    GetMentionText()(*string)
    GetOdataType()(*string)
    SetId(value *int32)()
    SetMentioned(value ChatMessageMentionedIdentitySetable)()
    SetMentionText(value *string)()
    SetOdataType(value *string)()
}
