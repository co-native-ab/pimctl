package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ChatMessagePolicyViolationPolicyTip struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URL a user can visit to read about the data loss prevention policies for the organization. (ie, policies about what users shouldn't say in chats)
    complianceUrl *string
    // Explanatory text shown to the sender of the message.
    generalText *string
    // The list of improper data in the message that was detected by the data loss prevention app. Each DLP app defines its own conditions, examples include 'Credit Card Number' and 'Social Security Number'.
    matchedConditionDescriptions []string
    // The OdataType property
    odataType *string
}
// NewChatMessagePolicyViolationPolicyTip instantiates a new ChatMessagePolicyViolationPolicyTip and sets the default values.
func NewChatMessagePolicyViolationPolicyTip()(*ChatMessagePolicyViolationPolicyTip) {
    m := &ChatMessagePolicyViolationPolicyTip{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChatMessagePolicyViolationPolicyTipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateChatMessagePolicyViolationPolicyTipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessagePolicyViolationPolicyTip(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ChatMessagePolicyViolationPolicyTip) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComplianceUrl gets the complianceUrl property value. The URL a user can visit to read about the data loss prevention policies for the organization. (ie, policies about what users shouldn't say in chats)
// returns a *string when successful
func (m *ChatMessagePolicyViolationPolicyTip) GetComplianceUrl()(*string) {
    return m.complianceUrl
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ChatMessagePolicyViolationPolicyTip) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["complianceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceUrl(val)
        }
        return nil
    }
    res["generalText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneralText(val)
        }
        return nil
    }
    res["matchedConditionDescriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetMatchedConditionDescriptions(res)
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
// GetGeneralText gets the generalText property value. Explanatory text shown to the sender of the message.
// returns a *string when successful
func (m *ChatMessagePolicyViolationPolicyTip) GetGeneralText()(*string) {
    return m.generalText
}
// GetMatchedConditionDescriptions gets the matchedConditionDescriptions property value. The list of improper data in the message that was detected by the data loss prevention app. Each DLP app defines its own conditions, examples include 'Credit Card Number' and 'Social Security Number'.
// returns a []string when successful
func (m *ChatMessagePolicyViolationPolicyTip) GetMatchedConditionDescriptions()([]string) {
    return m.matchedConditionDescriptions
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ChatMessagePolicyViolationPolicyTip) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ChatMessagePolicyViolationPolicyTip) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("complianceUrl", m.GetComplianceUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("generalText", m.GetGeneralText())
        if err != nil {
            return err
        }
    }
    if m.GetMatchedConditionDescriptions() != nil {
        err := writer.WriteCollectionOfStringValues("matchedConditionDescriptions", m.GetMatchedConditionDescriptions())
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
func (m *ChatMessagePolicyViolationPolicyTip) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComplianceUrl sets the complianceUrl property value. The URL a user can visit to read about the data loss prevention policies for the organization. (ie, policies about what users shouldn't say in chats)
func (m *ChatMessagePolicyViolationPolicyTip) SetComplianceUrl(value *string)() {
    m.complianceUrl = value
}
// SetGeneralText sets the generalText property value. Explanatory text shown to the sender of the message.
func (m *ChatMessagePolicyViolationPolicyTip) SetGeneralText(value *string)() {
    m.generalText = value
}
// SetMatchedConditionDescriptions sets the matchedConditionDescriptions property value. The list of improper data in the message that was detected by the data loss prevention app. Each DLP app defines its own conditions, examples include 'Credit Card Number' and 'Social Security Number'.
func (m *ChatMessagePolicyViolationPolicyTip) SetMatchedConditionDescriptions(value []string)() {
    m.matchedConditionDescriptions = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChatMessagePolicyViolationPolicyTip) SetOdataType(value *string)() {
    m.odataType = value
}
type ChatMessagePolicyViolationPolicyTipable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComplianceUrl()(*string)
    GetGeneralText()(*string)
    GetMatchedConditionDescriptions()([]string)
    GetOdataType()(*string)
    SetComplianceUrl(value *string)()
    SetGeneralText(value *string)()
    SetMatchedConditionDescriptions(value []string)()
    SetOdataType(value *string)()
}
