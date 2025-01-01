package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type X509CertificateUserBinding struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The priority of the binding. Microsoft Entra ID uses the binding with the highest priority. This value must be a non-negative integer and unique in the collection of objects in the certificateUserBindings property of an x509CertificateAuthenticationMethodConfiguration object. Required
    priority *int32
    // The trustAffinityLevel property
    trustAffinityLevel *X509CertificateAffinityLevel
    // Defines the Microsoft Entra user property of the user object to use for the binding. The possible values are: userPrincipalName, onPremisesUserPrincipalName, certificateUserIds. Required.
    userProperty *string
    // The field on the X.509 certificate to use for the binding. The possible values are: PrincipalName, RFC822Name, SubjectKeyIdentifier, SHA1PublicKey.
    x509CertificateField *string
}
// NewX509CertificateUserBinding instantiates a new X509CertificateUserBinding and sets the default values.
func NewX509CertificateUserBinding()(*X509CertificateUserBinding) {
    m := &X509CertificateUserBinding{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateX509CertificateUserBindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateX509CertificateUserBindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewX509CertificateUserBinding(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *X509CertificateUserBinding) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *X509CertificateUserBinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["trustAffinityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseX509CertificateAffinityLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustAffinityLevel(val.(*X509CertificateAffinityLevel))
        }
        return nil
    }
    res["userProperty"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserProperty(val)
        }
        return nil
    }
    res["x509CertificateField"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX509CertificateField(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *X509CertificateUserBinding) GetOdataType()(*string) {
    return m.odataType
}
// GetPriority gets the priority property value. The priority of the binding. Microsoft Entra ID uses the binding with the highest priority. This value must be a non-negative integer and unique in the collection of objects in the certificateUserBindings property of an x509CertificateAuthenticationMethodConfiguration object. Required
// returns a *int32 when successful
func (m *X509CertificateUserBinding) GetPriority()(*int32) {
    return m.priority
}
// GetTrustAffinityLevel gets the trustAffinityLevel property value. The trustAffinityLevel property
// returns a *X509CertificateAffinityLevel when successful
func (m *X509CertificateUserBinding) GetTrustAffinityLevel()(*X509CertificateAffinityLevel) {
    return m.trustAffinityLevel
}
// GetUserProperty gets the userProperty property value. Defines the Microsoft Entra user property of the user object to use for the binding. The possible values are: userPrincipalName, onPremisesUserPrincipalName, certificateUserIds. Required.
// returns a *string when successful
func (m *X509CertificateUserBinding) GetUserProperty()(*string) {
    return m.userProperty
}
// GetX509CertificateField gets the x509CertificateField property value. The field on the X.509 certificate to use for the binding. The possible values are: PrincipalName, RFC822Name, SubjectKeyIdentifier, SHA1PublicKey.
// returns a *string when successful
func (m *X509CertificateUserBinding) GetX509CertificateField()(*string) {
    return m.x509CertificateField
}
// Serialize serializes information the current object
func (m *X509CertificateUserBinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    if m.GetTrustAffinityLevel() != nil {
        cast := (*m.GetTrustAffinityLevel()).String()
        err := writer.WriteStringValue("trustAffinityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userProperty", m.GetUserProperty())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("x509CertificateField", m.GetX509CertificateField())
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
func (m *X509CertificateUserBinding) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *X509CertificateUserBinding) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPriority sets the priority property value. The priority of the binding. Microsoft Entra ID uses the binding with the highest priority. This value must be a non-negative integer and unique in the collection of objects in the certificateUserBindings property of an x509CertificateAuthenticationMethodConfiguration object. Required
func (m *X509CertificateUserBinding) SetPriority(value *int32)() {
    m.priority = value
}
// SetTrustAffinityLevel sets the trustAffinityLevel property value. The trustAffinityLevel property
func (m *X509CertificateUserBinding) SetTrustAffinityLevel(value *X509CertificateAffinityLevel)() {
    m.trustAffinityLevel = value
}
// SetUserProperty sets the userProperty property value. Defines the Microsoft Entra user property of the user object to use for the binding. The possible values are: userPrincipalName, onPremisesUserPrincipalName, certificateUserIds. Required.
func (m *X509CertificateUserBinding) SetUserProperty(value *string)() {
    m.userProperty = value
}
// SetX509CertificateField sets the x509CertificateField property value. The field on the X.509 certificate to use for the binding. The possible values are: PrincipalName, RFC822Name, SubjectKeyIdentifier, SHA1PublicKey.
func (m *X509CertificateUserBinding) SetX509CertificateField(value *string)() {
    m.x509CertificateField = value
}
type X509CertificateUserBindingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPriority()(*int32)
    GetTrustAffinityLevel()(*X509CertificateAffinityLevel)
    GetUserProperty()(*string)
    GetX509CertificateField()(*string)
    SetOdataType(value *string)()
    SetPriority(value *int32)()
    SetTrustAffinityLevel(value *X509CertificateAffinityLevel)()
    SetUserProperty(value *string)()
    SetX509CertificateField(value *string)()
}
