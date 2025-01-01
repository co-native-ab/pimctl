package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type SslCertificateEntity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A physical address of the entity.
    address ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PhysicalAddressable
    // Alternate names for this entity that are part of the certificate.
    alternateNames []string
    // A common name for this entity.
    commonName *string
    // An email for this entity.
    email *string
    // If the entity is a person, this is the person's given name (first name).
    givenName *string
    // The OdataType property
    odataType *string
    // If the entity is an organization, this is the name of the organization.
    organizationName *string
    // If the entity is an organization, this communicates if a unit in the organization is named on the entity.
    organizationUnitName *string
    // A serial number assigned to the entity; usually only available if the entity is the issuer.
    serialNumber *string
    // If the entity is a person, this is the person's surname (last name).
    surname *string
}
// NewSslCertificateEntity instantiates a new SslCertificateEntity and sets the default values.
func NewSslCertificateEntity()(*SslCertificateEntity) {
    m := &SslCertificateEntity{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSslCertificateEntityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSslCertificateEntityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSslCertificateEntity(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SslCertificateEntity) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. A physical address of the entity.
// returns a PhysicalAddressable when successful
func (m *SslCertificateEntity) GetAddress()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PhysicalAddressable) {
    return m.address
}
// GetAlternateNames gets the alternateNames property value. Alternate names for this entity that are part of the certificate.
// returns a []string when successful
func (m *SslCertificateEntity) GetAlternateNames()([]string) {
    return m.alternateNames
}
// GetCommonName gets the commonName property value. A common name for this entity.
// returns a *string when successful
func (m *SslCertificateEntity) GetCommonName()(*string) {
    return m.commonName
}
// GetEmail gets the email property value. An email for this entity.
// returns a *string when successful
func (m *SslCertificateEntity) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SslCertificateEntity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["alternateNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAlternateNames(res)
        }
        return nil
    }
    res["commonName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommonName(val)
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
    res["givenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
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
    res["organizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationName(val)
        }
        return nil
    }
    res["organizationUnitName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationUnitName(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["surname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    return res
}
// GetGivenName gets the givenName property value. If the entity is a person, this is the person's given name (first name).
// returns a *string when successful
func (m *SslCertificateEntity) GetGivenName()(*string) {
    return m.givenName
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SslCertificateEntity) GetOdataType()(*string) {
    return m.odataType
}
// GetOrganizationName gets the organizationName property value. If the entity is an organization, this is the name of the organization.
// returns a *string when successful
func (m *SslCertificateEntity) GetOrganizationName()(*string) {
    return m.organizationName
}
// GetOrganizationUnitName gets the organizationUnitName property value. If the entity is an organization, this communicates if a unit in the organization is named on the entity.
// returns a *string when successful
func (m *SslCertificateEntity) GetOrganizationUnitName()(*string) {
    return m.organizationUnitName
}
// GetSerialNumber gets the serialNumber property value. A serial number assigned to the entity; usually only available if the entity is the issuer.
// returns a *string when successful
func (m *SslCertificateEntity) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetSurname gets the surname property value. If the entity is a person, this is the person's surname (last name).
// returns a *string when successful
func (m *SslCertificateEntity) GetSurname()(*string) {
    return m.surname
}
// Serialize serializes information the current object
func (m *SslCertificateEntity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    if m.GetAlternateNames() != nil {
        err := writer.WriteCollectionOfStringValues("alternateNames", m.GetAlternateNames())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commonName", m.GetCommonName())
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
        err := writer.WriteStringValue("givenName", m.GetGivenName())
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
        err := writer.WriteStringValue("organizationName", m.GetOrganizationName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organizationUnitName", m.GetOrganizationUnitName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("surname", m.GetSurname())
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
func (m *SslCertificateEntity) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. A physical address of the entity.
func (m *SslCertificateEntity) SetAddress(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PhysicalAddressable)() {
    m.address = value
}
// SetAlternateNames sets the alternateNames property value. Alternate names for this entity that are part of the certificate.
func (m *SslCertificateEntity) SetAlternateNames(value []string)() {
    m.alternateNames = value
}
// SetCommonName sets the commonName property value. A common name for this entity.
func (m *SslCertificateEntity) SetCommonName(value *string)() {
    m.commonName = value
}
// SetEmail sets the email property value. An email for this entity.
func (m *SslCertificateEntity) SetEmail(value *string)() {
    m.email = value
}
// SetGivenName sets the givenName property value. If the entity is a person, this is the person's given name (first name).
func (m *SslCertificateEntity) SetGivenName(value *string)() {
    m.givenName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SslCertificateEntity) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrganizationName sets the organizationName property value. If the entity is an organization, this is the name of the organization.
func (m *SslCertificateEntity) SetOrganizationName(value *string)() {
    m.organizationName = value
}
// SetOrganizationUnitName sets the organizationUnitName property value. If the entity is an organization, this communicates if a unit in the organization is named on the entity.
func (m *SslCertificateEntity) SetOrganizationUnitName(value *string)() {
    m.organizationUnitName = value
}
// SetSerialNumber sets the serialNumber property value. A serial number assigned to the entity; usually only available if the entity is the issuer.
func (m *SslCertificateEntity) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetSurname sets the surname property value. If the entity is a person, this is the person's surname (last name).
func (m *SslCertificateEntity) SetSurname(value *string)() {
    m.surname = value
}
type SslCertificateEntityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PhysicalAddressable)
    GetAlternateNames()([]string)
    GetCommonName()(*string)
    GetEmail()(*string)
    GetGivenName()(*string)
    GetOdataType()(*string)
    GetOrganizationName()(*string)
    GetOrganizationUnitName()(*string)
    GetSerialNumber()(*string)
    GetSurname()(*string)
    SetAddress(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PhysicalAddressable)()
    SetAlternateNames(value []string)()
    SetCommonName(value *string)()
    SetEmail(value *string)()
    SetGivenName(value *string)()
    SetOdataType(value *string)()
    SetOrganizationName(value *string)()
    SetOrganizationUnitName(value *string)()
    SetSerialNumber(value *string)()
    SetSurname(value *string)()
}
