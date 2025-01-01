package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OAuthConsentAppDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // App scope. Possible values are: unknown, readCalendar, readContact, readMail, readAllChat, readAllFile, readAndWriteMail, sendMail, unknownFutureValue.
    appScope *OAuthAppScope
    // App display logo.
    displayLogo *string
    // App name.
    displayName *string
    // The OdataType property
    odataType *string
}
// NewOAuthConsentAppDetail instantiates a new OAuthConsentAppDetail and sets the default values.
func NewOAuthConsentAppDetail()(*OAuthConsentAppDetail) {
    m := &OAuthConsentAppDetail{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOAuthConsentAppDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOAuthConsentAppDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOAuthConsentAppDetail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OAuthConsentAppDetail) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppScope gets the appScope property value. App scope. Possible values are: unknown, readCalendar, readContact, readMail, readAllChat, readAllFile, readAndWriteMail, sendMail, unknownFutureValue.
// returns a *OAuthAppScope when successful
func (m *OAuthConsentAppDetail) GetAppScope()(*OAuthAppScope) {
    return m.appScope
}
// GetDisplayLogo gets the displayLogo property value. App display logo.
// returns a *string when successful
func (m *OAuthConsentAppDetail) GetDisplayLogo()(*string) {
    return m.displayLogo
}
// GetDisplayName gets the displayName property value. App name.
// returns a *string when successful
func (m *OAuthConsentAppDetail) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OAuthConsentAppDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOAuthAppScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppScope(val.(*OAuthAppScope))
        }
        return nil
    }
    res["displayLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayLogo(val)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *OAuthConsentAppDetail) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *OAuthConsentAppDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppScope() != nil {
        cast := (*m.GetAppScope()).String()
        err := writer.WriteStringValue("appScope", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayLogo", m.GetDisplayLogo())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OAuthConsentAppDetail) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppScope sets the appScope property value. App scope. Possible values are: unknown, readCalendar, readContact, readMail, readAllChat, readAllFile, readAndWriteMail, sendMail, unknownFutureValue.
func (m *OAuthConsentAppDetail) SetAppScope(value *OAuthAppScope)() {
    m.appScope = value
}
// SetDisplayLogo sets the displayLogo property value. App display logo.
func (m *OAuthConsentAppDetail) SetDisplayLogo(value *string)() {
    m.displayLogo = value
}
// SetDisplayName sets the displayName property value. App name.
func (m *OAuthConsentAppDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OAuthConsentAppDetail) SetOdataType(value *string)() {
    m.odataType = value
}
type OAuthConsentAppDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppScope()(*OAuthAppScope)
    GetDisplayLogo()(*string)
    GetDisplayName()(*string)
    GetOdataType()(*string)
    SetAppScope(value *OAuthAppScope)()
    SetDisplayLogo(value *string)()
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
}
