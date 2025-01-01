package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthenticationMethodsRegistrationCampaign struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Users and groups of users that are excluded from being prompted to set up the authentication method.
    excludeTargets []ExcludeTargetable
    // Users and groups of users that are prompted to set up the authentication method.
    includeTargets []AuthenticationMethodsRegistrationCampaignIncludeTargetable
    // The OdataType property
    odataType *string
    // Specifies the number of days that the user sees a prompt again if they select 'Not now' and snoozes the prompt. Minimum: 0 days. Maximum: 14 days. If the value is '0', the user is prompted during every MFA attempt.
    snoozeDurationInDays *int32
    // The state property
    state *AdvancedConfigState
}
// NewAuthenticationMethodsRegistrationCampaign instantiates a new AuthenticationMethodsRegistrationCampaign and sets the default values.
func NewAuthenticationMethodsRegistrationCampaign()(*AuthenticationMethodsRegistrationCampaign) {
    m := &AuthenticationMethodsRegistrationCampaign{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationMethodsRegistrationCampaignFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationMethodsRegistrationCampaignFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationMethodsRegistrationCampaign(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthenticationMethodsRegistrationCampaign) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExcludeTargets gets the excludeTargets property value. Users and groups of users that are excluded from being prompted to set up the authentication method.
// returns a []ExcludeTargetable when successful
func (m *AuthenticationMethodsRegistrationCampaign) GetExcludeTargets()([]ExcludeTargetable) {
    return m.excludeTargets
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationMethodsRegistrationCampaign) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExcludeTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExcludeTargetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExcludeTargetable)
                }
            }
            m.SetExcludeTargets(res)
        }
        return nil
    }
    res["includeTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodsRegistrationCampaignIncludeTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodsRegistrationCampaignIncludeTargetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationMethodsRegistrationCampaignIncludeTargetable)
                }
            }
            m.SetIncludeTargets(res)
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
    res["snoozeDurationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnoozeDurationInDays(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdvancedConfigState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AdvancedConfigState))
        }
        return nil
    }
    return res
}
// GetIncludeTargets gets the includeTargets property value. Users and groups of users that are prompted to set up the authentication method.
// returns a []AuthenticationMethodsRegistrationCampaignIncludeTargetable when successful
func (m *AuthenticationMethodsRegistrationCampaign) GetIncludeTargets()([]AuthenticationMethodsRegistrationCampaignIncludeTargetable) {
    return m.includeTargets
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AuthenticationMethodsRegistrationCampaign) GetOdataType()(*string) {
    return m.odataType
}
// GetSnoozeDurationInDays gets the snoozeDurationInDays property value. Specifies the number of days that the user sees a prompt again if they select 'Not now' and snoozes the prompt. Minimum: 0 days. Maximum: 14 days. If the value is '0', the user is prompted during every MFA attempt.
// returns a *int32 when successful
func (m *AuthenticationMethodsRegistrationCampaign) GetSnoozeDurationInDays()(*int32) {
    return m.snoozeDurationInDays
}
// GetState gets the state property value. The state property
// returns a *AdvancedConfigState when successful
func (m *AuthenticationMethodsRegistrationCampaign) GetState()(*AdvancedConfigState) {
    return m.state
}
// Serialize serializes information the current object
func (m *AuthenticationMethodsRegistrationCampaign) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludeTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExcludeTargets()))
        for i, v := range m.GetExcludeTargets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("excludeTargets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncludeTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludeTargets()))
        for i, v := range m.GetIncludeTargets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("includeTargets", cast)
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
        err := writer.WriteInt32Value("snoozeDurationInDays", m.GetSnoozeDurationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *AuthenticationMethodsRegistrationCampaign) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExcludeTargets sets the excludeTargets property value. Users and groups of users that are excluded from being prompted to set up the authentication method.
func (m *AuthenticationMethodsRegistrationCampaign) SetExcludeTargets(value []ExcludeTargetable)() {
    m.excludeTargets = value
}
// SetIncludeTargets sets the includeTargets property value. Users and groups of users that are prompted to set up the authentication method.
func (m *AuthenticationMethodsRegistrationCampaign) SetIncludeTargets(value []AuthenticationMethodsRegistrationCampaignIncludeTargetable)() {
    m.includeTargets = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationMethodsRegistrationCampaign) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSnoozeDurationInDays sets the snoozeDurationInDays property value. Specifies the number of days that the user sees a prompt again if they select 'Not now' and snoozes the prompt. Minimum: 0 days. Maximum: 14 days. If the value is '0', the user is prompted during every MFA attempt.
func (m *AuthenticationMethodsRegistrationCampaign) SetSnoozeDurationInDays(value *int32)() {
    m.snoozeDurationInDays = value
}
// SetState sets the state property value. The state property
func (m *AuthenticationMethodsRegistrationCampaign) SetState(value *AdvancedConfigState)() {
    m.state = value
}
type AuthenticationMethodsRegistrationCampaignable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludeTargets()([]ExcludeTargetable)
    GetIncludeTargets()([]AuthenticationMethodsRegistrationCampaignIncludeTargetable)
    GetOdataType()(*string)
    GetSnoozeDurationInDays()(*int32)
    GetState()(*AdvancedConfigState)
    SetExcludeTargets(value []ExcludeTargetable)()
    SetIncludeTargets(value []AuthenticationMethodsRegistrationCampaignIncludeTargetable)()
    SetOdataType(value *string)()
    SetSnoozeDurationInDays(value *int32)()
    SetState(value *AdvancedConfigState)()
}
