package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthenticationMethodsRegistrationCampaignIncludeTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The object identifier of a Microsoft Entra user or group.
    id *string
    // The OdataType property
    odataType *string
    // The authentication method that the user is prompted to register. The value must be microsoftAuthenticator.
    targetedAuthenticationMethod *string
    // The targetType property
    targetType *AuthenticationMethodTargetType
}
// NewAuthenticationMethodsRegistrationCampaignIncludeTarget instantiates a new AuthenticationMethodsRegistrationCampaignIncludeTarget and sets the default values.
func NewAuthenticationMethodsRegistrationCampaignIncludeTarget()(*AuthenticationMethodsRegistrationCampaignIncludeTarget) {
    m := &AuthenticationMethodsRegistrationCampaignIncludeTarget{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationMethodsRegistrationCampaignIncludeTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationMethodsRegistrationCampaignIncludeTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationMethodsRegistrationCampaignIncludeTarget(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["targetedAuthenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedAuthenticationMethod(val)
        }
        return nil
    }
    res["targetType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodTargetType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetType(val.(*AuthenticationMethodTargetType))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The object identifier of a Microsoft Entra user or group.
// returns a *string when successful
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetOdataType()(*string) {
    return m.odataType
}
// GetTargetedAuthenticationMethod gets the targetedAuthenticationMethod property value. The authentication method that the user is prompted to register. The value must be microsoftAuthenticator.
// returns a *string when successful
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetTargetedAuthenticationMethod()(*string) {
    return m.targetedAuthenticationMethod
}
// GetTargetType gets the targetType property value. The targetType property
// returns a *AuthenticationMethodTargetType when successful
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) GetTargetType()(*AuthenticationMethodTargetType) {
    return m.targetType
}
// Serialize serializes information the current object
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
        err := writer.WriteStringValue("targetedAuthenticationMethod", m.GetTargetedAuthenticationMethod())
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := (*m.GetTargetType()).String()
        err := writer.WriteStringValue("targetType", &cast)
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
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The object identifier of a Microsoft Entra user or group.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTargetedAuthenticationMethod sets the targetedAuthenticationMethod property value. The authentication method that the user is prompted to register. The value must be microsoftAuthenticator.
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetTargetedAuthenticationMethod(value *string)() {
    m.targetedAuthenticationMethod = value
}
// SetTargetType sets the targetType property value. The targetType property
func (m *AuthenticationMethodsRegistrationCampaignIncludeTarget) SetTargetType(value *AuthenticationMethodTargetType)() {
    m.targetType = value
}
type AuthenticationMethodsRegistrationCampaignIncludeTargetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetOdataType()(*string)
    GetTargetedAuthenticationMethod()(*string)
    GetTargetType()(*AuthenticationMethodTargetType)
    SetId(value *string)()
    SetOdataType(value *string)()
    SetTargetedAuthenticationMethod(value *string)()
    SetTargetType(value *AuthenticationMethodTargetType)()
}
