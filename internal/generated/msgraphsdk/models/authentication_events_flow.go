package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthenticationEventsFlow struct {
    Entity
    // The conditions representing the context of the authentication request that's used to decide whether the events policy is invoked.  Supports $filter (eq). See support for filtering on user flows for syntax information.
    conditions AuthenticationConditionsable
    // The description of the events policy.
    description *string
    // Required. The display name for the events policy.
    displayName *string
}
// NewAuthenticationEventsFlow instantiates a new AuthenticationEventsFlow and sets the default values.
func NewAuthenticationEventsFlow()(*AuthenticationEventsFlow) {
    m := &AuthenticationEventsFlow{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationEventsFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationEventsFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.externalUsersSelfServiceSignUpEventsFlow":
                        return NewExternalUsersSelfServiceSignUpEventsFlow(), nil
                }
            }
        }
    }
    return NewAuthenticationEventsFlow(), nil
}
// GetConditions gets the conditions property value. The conditions representing the context of the authentication request that's used to decide whether the events policy is invoked.  Supports $filter (eq). See support for filtering on user flows for syntax information.
// returns a AuthenticationConditionsable when successful
func (m *AuthenticationEventsFlow) GetConditions()(AuthenticationConditionsable) {
    return m.conditions
}
// GetDescription gets the description property value. The description of the events policy.
// returns a *string when successful
func (m *AuthenticationEventsFlow) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Required. The display name for the events policy.
// returns a *string when successful
func (m *AuthenticationEventsFlow) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationEventsFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditions(val.(AuthenticationConditionsable))
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
    return res
}
// Serialize serializes information the current object
func (m *AuthenticationEventsFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("conditions", m.GetConditions())
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
    return nil
}
// SetConditions sets the conditions property value. The conditions representing the context of the authentication request that's used to decide whether the events policy is invoked.  Supports $filter (eq). See support for filtering on user flows for syntax information.
func (m *AuthenticationEventsFlow) SetConditions(value AuthenticationConditionsable)() {
    m.conditions = value
}
// SetDescription sets the description property value. The description of the events policy.
func (m *AuthenticationEventsFlow) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Required. The display name for the events policy.
func (m *AuthenticationEventsFlow) SetDisplayName(value *string)() {
    m.displayName = value
}
type AuthenticationEventsFlowable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConditions()(AuthenticationConditionsable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    SetConditions(value AuthenticationConditionsable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
}
