package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConditionalAccessGuestsOrExternalUsers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The tenant IDs of the selected types of external users. Either all B2B tenant or a collection of tenant IDs. External tenants can be specified only when the property guestOrExternalUserTypes isn't null or an empty String.
    externalTenants ConditionalAccessExternalTenantsable
    // The guestOrExternalUserTypes property
    guestOrExternalUserTypes *ConditionalAccessGuestOrExternalUserTypes
    // The OdataType property
    odataType *string
}
// NewConditionalAccessGuestsOrExternalUsers instantiates a new ConditionalAccessGuestsOrExternalUsers and sets the default values.
func NewConditionalAccessGuestsOrExternalUsers()(*ConditionalAccessGuestsOrExternalUsers) {
    m := &ConditionalAccessGuestsOrExternalUsers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessGuestsOrExternalUsersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionalAccessGuestsOrExternalUsersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessGuestsOrExternalUsers(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConditionalAccessGuestsOrExternalUsers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExternalTenants gets the externalTenants property value. The tenant IDs of the selected types of external users. Either all B2B tenant or a collection of tenant IDs. External tenants can be specified only when the property guestOrExternalUserTypes isn't null or an empty String.
// returns a ConditionalAccessExternalTenantsable when successful
func (m *ConditionalAccessGuestsOrExternalUsers) GetExternalTenants()(ConditionalAccessExternalTenantsable) {
    return m.externalTenants
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConditionalAccessGuestsOrExternalUsers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalTenants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessExternalTenantsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalTenants(val.(ConditionalAccessExternalTenantsable))
        }
        return nil
    }
    res["guestOrExternalUserTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessGuestOrExternalUserTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuestOrExternalUserTypes(val.(*ConditionalAccessGuestOrExternalUserTypes))
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
// GetGuestOrExternalUserTypes gets the guestOrExternalUserTypes property value. The guestOrExternalUserTypes property
// returns a *ConditionalAccessGuestOrExternalUserTypes when successful
func (m *ConditionalAccessGuestsOrExternalUsers) GetGuestOrExternalUserTypes()(*ConditionalAccessGuestOrExternalUserTypes) {
    return m.guestOrExternalUserTypes
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ConditionalAccessGuestsOrExternalUsers) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ConditionalAccessGuestsOrExternalUsers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("externalTenants", m.GetExternalTenants())
        if err != nil {
            return err
        }
    }
    if m.GetGuestOrExternalUserTypes() != nil {
        cast := (*m.GetGuestOrExternalUserTypes()).String()
        err := writer.WriteStringValue("guestOrExternalUserTypes", &cast)
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
func (m *ConditionalAccessGuestsOrExternalUsers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExternalTenants sets the externalTenants property value. The tenant IDs of the selected types of external users. Either all B2B tenant or a collection of tenant IDs. External tenants can be specified only when the property guestOrExternalUserTypes isn't null or an empty String.
func (m *ConditionalAccessGuestsOrExternalUsers) SetExternalTenants(value ConditionalAccessExternalTenantsable)() {
    m.externalTenants = value
}
// SetGuestOrExternalUserTypes sets the guestOrExternalUserTypes property value. The guestOrExternalUserTypes property
func (m *ConditionalAccessGuestsOrExternalUsers) SetGuestOrExternalUserTypes(value *ConditionalAccessGuestOrExternalUserTypes)() {
    m.guestOrExternalUserTypes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessGuestsOrExternalUsers) SetOdataType(value *string)() {
    m.odataType = value
}
type ConditionalAccessGuestsOrExternalUsersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExternalTenants()(ConditionalAccessExternalTenantsable)
    GetGuestOrExternalUserTypes()(*ConditionalAccessGuestOrExternalUserTypes)
    GetOdataType()(*string)
    SetExternalTenants(value ConditionalAccessExternalTenantsable)()
    SetGuestOrExternalUserTypes(value *ConditionalAccessGuestOrExternalUserTypes)()
    SetOdataType(value *string)()
}
