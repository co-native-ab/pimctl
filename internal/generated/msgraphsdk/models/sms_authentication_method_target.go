package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SmsAuthenticationMethodTarget struct {
    AuthenticationMethodTarget
    // Determines if users can use this authentication method to sign in to Microsoft Entra ID. true if users can use this method for primary authentication, otherwise false.
    isUsableForSignIn *bool
}
// NewSmsAuthenticationMethodTarget instantiates a new SmsAuthenticationMethodTarget and sets the default values.
func NewSmsAuthenticationMethodTarget()(*SmsAuthenticationMethodTarget) {
    m := &SmsAuthenticationMethodTarget{
        AuthenticationMethodTarget: *NewAuthenticationMethodTarget(),
    }
    return m
}
// CreateSmsAuthenticationMethodTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSmsAuthenticationMethodTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSmsAuthenticationMethodTarget(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SmsAuthenticationMethodTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodTarget.GetFieldDeserializers()
    res["isUsableForSignIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUsableForSignIn(val)
        }
        return nil
    }
    return res
}
// GetIsUsableForSignIn gets the isUsableForSignIn property value. Determines if users can use this authentication method to sign in to Microsoft Entra ID. true if users can use this method for primary authentication, otherwise false.
// returns a *bool when successful
func (m *SmsAuthenticationMethodTarget) GetIsUsableForSignIn()(*bool) {
    return m.isUsableForSignIn
}
// Serialize serializes information the current object
func (m *SmsAuthenticationMethodTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodTarget.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isUsableForSignIn", m.GetIsUsableForSignIn())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsUsableForSignIn sets the isUsableForSignIn property value. Determines if users can use this authentication method to sign in to Microsoft Entra ID. true if users can use this method for primary authentication, otherwise false.
func (m *SmsAuthenticationMethodTarget) SetIsUsableForSignIn(value *bool)() {
    m.isUsableForSignIn = value
}
type SmsAuthenticationMethodTargetable interface {
    AuthenticationMethodTargetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsUsableForSignIn()(*bool)
    SetIsUsableForSignIn(value *bool)()
}
