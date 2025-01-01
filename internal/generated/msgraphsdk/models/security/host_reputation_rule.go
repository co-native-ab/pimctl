package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HostReputationRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description of the rule that gives more context.
    description *string
    // The name of the rule.
    name *string
    // The OdataType property
    odataType *string
    // Link to a web page with details related to this rule.
    relatedDetailsUrl *string
    // The severity property
    severity *HostReputationRuleSeverity
}
// NewHostReputationRule instantiates a new HostReputationRule and sets the default values.
func NewHostReputationRule()(*HostReputationRule) {
    m := &HostReputationRule{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHostReputationRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHostReputationRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHostReputationRule(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HostReputationRule) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description of the rule that gives more context.
// returns a *string when successful
func (m *HostReputationRule) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HostReputationRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["relatedDetailsUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelatedDetailsUrl(val)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHostReputationRuleSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*HostReputationRuleSeverity))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the rule.
// returns a *string when successful
func (m *HostReputationRule) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *HostReputationRule) GetOdataType()(*string) {
    return m.odataType
}
// GetRelatedDetailsUrl gets the relatedDetailsUrl property value. Link to a web page with details related to this rule.
// returns a *string when successful
func (m *HostReputationRule) GetRelatedDetailsUrl()(*string) {
    return m.relatedDetailsUrl
}
// GetSeverity gets the severity property value. The severity property
// returns a *HostReputationRuleSeverity when successful
func (m *HostReputationRule) GetSeverity()(*HostReputationRuleSeverity) {
    return m.severity
}
// Serialize serializes information the current object
func (m *HostReputationRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteStringValue("relatedDetailsUrl", m.GetRelatedDetailsUrl())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err := writer.WriteStringValue("severity", &cast)
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
func (m *HostReputationRule) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description of the rule that gives more context.
func (m *HostReputationRule) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. The name of the rule.
func (m *HostReputationRule) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *HostReputationRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRelatedDetailsUrl sets the relatedDetailsUrl property value. Link to a web page with details related to this rule.
func (m *HostReputationRule) SetRelatedDetailsUrl(value *string)() {
    m.relatedDetailsUrl = value
}
// SetSeverity sets the severity property value. The severity property
func (m *HostReputationRule) SetSeverity(value *HostReputationRuleSeverity)() {
    m.severity = value
}
type HostReputationRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetName()(*string)
    GetOdataType()(*string)
    GetRelatedDetailsUrl()(*string)
    GetSeverity()(*HostReputationRuleSeverity)
    SetDescription(value *string)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetRelatedDetailsUrl(value *string)()
    SetSeverity(value *HostReputationRuleSeverity)()
}
