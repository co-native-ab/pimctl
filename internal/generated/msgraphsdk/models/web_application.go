package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WebApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Home page or landing page of the application.
    homePageUrl *string
    // Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
    implicitGrantSettings ImplicitGrantSettingsable
    // Specifies the URL that is used by Microsoft's authorization service to log out a user using front-channel, back-channel or SAML logout protocols.
    logoutUrl *string
    // The OdataType property
    odataType *string
    // Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
    redirectUris []string
    // The redirectUriSettings property
    redirectUriSettings []RedirectUriSettingsable
}
// NewWebApplication instantiates a new WebApplication and sets the default values.
func NewWebApplication()(*WebApplication) {
    m := &WebApplication{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWebApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWebApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebApplication(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WebApplication) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WebApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["homePageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomePageUrl(val)
        }
        return nil
    }
    res["implicitGrantSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateImplicitGrantSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImplicitGrantSettings(val.(ImplicitGrantSettingsable))
        }
        return nil
    }
    res["logoutUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogoutUrl(val)
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
    res["redirectUris"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRedirectUris(res)
        }
        return nil
    }
    res["redirectUriSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRedirectUriSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RedirectUriSettingsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RedirectUriSettingsable)
                }
            }
            m.SetRedirectUriSettings(res)
        }
        return nil
    }
    return res
}
// GetHomePageUrl gets the homePageUrl property value. Home page or landing page of the application.
// returns a *string when successful
func (m *WebApplication) GetHomePageUrl()(*string) {
    return m.homePageUrl
}
// GetImplicitGrantSettings gets the implicitGrantSettings property value. Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
// returns a ImplicitGrantSettingsable when successful
func (m *WebApplication) GetImplicitGrantSettings()(ImplicitGrantSettingsable) {
    return m.implicitGrantSettings
}
// GetLogoutUrl gets the logoutUrl property value. Specifies the URL that is used by Microsoft's authorization service to log out a user using front-channel, back-channel or SAML logout protocols.
// returns a *string when successful
func (m *WebApplication) GetLogoutUrl()(*string) {
    return m.logoutUrl
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WebApplication) GetOdataType()(*string) {
    return m.odataType
}
// GetRedirectUris gets the redirectUris property value. Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
// returns a []string when successful
func (m *WebApplication) GetRedirectUris()([]string) {
    return m.redirectUris
}
// GetRedirectUriSettings gets the redirectUriSettings property value. The redirectUriSettings property
// returns a []RedirectUriSettingsable when successful
func (m *WebApplication) GetRedirectUriSettings()([]RedirectUriSettingsable) {
    return m.redirectUriSettings
}
// Serialize serializes information the current object
func (m *WebApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("homePageUrl", m.GetHomePageUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("implicitGrantSettings", m.GetImplicitGrantSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logoutUrl", m.GetLogoutUrl())
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
    if m.GetRedirectUris() != nil {
        err := writer.WriteCollectionOfStringValues("redirectUris", m.GetRedirectUris())
        if err != nil {
            return err
        }
    }
    if m.GetRedirectUriSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRedirectUriSettings()))
        for i, v := range m.GetRedirectUriSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("redirectUriSettings", cast)
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
func (m *WebApplication) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHomePageUrl sets the homePageUrl property value. Home page or landing page of the application.
func (m *WebApplication) SetHomePageUrl(value *string)() {
    m.homePageUrl = value
}
// SetImplicitGrantSettings sets the implicitGrantSettings property value. Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
func (m *WebApplication) SetImplicitGrantSettings(value ImplicitGrantSettingsable)() {
    m.implicitGrantSettings = value
}
// SetLogoutUrl sets the logoutUrl property value. Specifies the URL that is used by Microsoft's authorization service to log out a user using front-channel, back-channel or SAML logout protocols.
func (m *WebApplication) SetLogoutUrl(value *string)() {
    m.logoutUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WebApplication) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRedirectUris sets the redirectUris property value. Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
func (m *WebApplication) SetRedirectUris(value []string)() {
    m.redirectUris = value
}
// SetRedirectUriSettings sets the redirectUriSettings property value. The redirectUriSettings property
func (m *WebApplication) SetRedirectUriSettings(value []RedirectUriSettingsable)() {
    m.redirectUriSettings = value
}
type WebApplicationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHomePageUrl()(*string)
    GetImplicitGrantSettings()(ImplicitGrantSettingsable)
    GetLogoutUrl()(*string)
    GetOdataType()(*string)
    GetRedirectUris()([]string)
    GetRedirectUriSettings()([]RedirectUriSettingsable)
    SetHomePageUrl(value *string)()
    SetImplicitGrantSettings(value ImplicitGrantSettingsable)()
    SetLogoutUrl(value *string)()
    SetOdataType(value *string)()
    SetRedirectUris(value []string)()
    SetRedirectUriSettings(value []RedirectUriSettingsable)()
}
