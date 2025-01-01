package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type WhoisContact struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The physical address of the entity.
    address ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PhysicalAddressable
    // The email of this WHOIS contact.
    email *string
    // The fax of this WHOIS contact. No format is guaranteed.
    fax *string
    // The name of this WHOIS contact.
    name *string
    // The OdataType property
    odataType *string
    // The organization of this WHOIS contact.
    organization *string
    // The telephone of this WHOIS contact. No format is guaranteed.
    telephone *string
}
// NewWhoisContact instantiates a new WhoisContact and sets the default values.
func NewWhoisContact()(*WhoisContact) {
    m := &WhoisContact{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWhoisContactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWhoisContactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWhoisContact(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WhoisContact) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The physical address of the entity.
// returns a PhysicalAddressable when successful
func (m *WhoisContact) GetAddress()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PhysicalAddressable) {
    return m.address
}
// GetEmail gets the email property value. The email of this WHOIS contact.
// returns a *string when successful
func (m *WhoisContact) GetEmail()(*string) {
    return m.email
}
// GetFax gets the fax property value. The fax of this WHOIS contact. No format is guaranteed.
// returns a *string when successful
func (m *WhoisContact) GetFax()(*string) {
    return m.fax
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WhoisContact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PhysicalAddressable))
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["fax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFax(val)
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
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val)
        }
        return nil
    }
    res["telephone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTelephone(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of this WHOIS contact.
// returns a *string when successful
func (m *WhoisContact) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WhoisContact) GetOdataType()(*string) {
    return m.odataType
}
// GetOrganization gets the organization property value. The organization of this WHOIS contact.
// returns a *string when successful
func (m *WhoisContact) GetOrganization()(*string) {
    return m.organization
}
// GetTelephone gets the telephone property value. The telephone of this WHOIS contact. No format is guaranteed.
// returns a *string when successful
func (m *WhoisContact) GetTelephone()(*string) {
    return m.telephone
}
// Serialize serializes information the current object
func (m *WhoisContact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fax", m.GetFax())
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
        err := writer.WriteStringValue("organization", m.GetOrganization())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("telephone", m.GetTelephone())
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
func (m *WhoisContact) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The physical address of the entity.
func (m *WhoisContact) SetAddress(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PhysicalAddressable)() {
    m.address = value
}
// SetEmail sets the email property value. The email of this WHOIS contact.
func (m *WhoisContact) SetEmail(value *string)() {
    m.email = value
}
// SetFax sets the fax property value. The fax of this WHOIS contact. No format is guaranteed.
func (m *WhoisContact) SetFax(value *string)() {
    m.fax = value
}
// SetName sets the name property value. The name of this WHOIS contact.
func (m *WhoisContact) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WhoisContact) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrganization sets the organization property value. The organization of this WHOIS contact.
func (m *WhoisContact) SetOrganization(value *string)() {
    m.organization = value
}
// SetTelephone sets the telephone property value. The telephone of this WHOIS contact. No format is guaranteed.
func (m *WhoisContact) SetTelephone(value *string)() {
    m.telephone = value
}
type WhoisContactable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PhysicalAddressable)
    GetEmail()(*string)
    GetFax()(*string)
    GetName()(*string)
    GetOdataType()(*string)
    GetOrganization()(*string)
    GetTelephone()(*string)
    SetAddress(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PhysicalAddressable)()
    SetEmail(value *string)()
    SetFax(value *string)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetOrganization(value *string)()
    SetTelephone(value *string)()
}
