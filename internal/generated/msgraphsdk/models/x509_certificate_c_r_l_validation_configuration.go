package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type X509CertificateCRLValidationConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The exemptedCertificateAuthoritiesSubjectKeyIdentifiers property
    exemptedCertificateAuthoritiesSubjectKeyIdentifiers []string
    // The OdataType property
    odataType *string
    // The state property
    state *X509CertificateCRLValidationConfigurationState
}
// NewX509CertificateCRLValidationConfiguration instantiates a new X509CertificateCRLValidationConfiguration and sets the default values.
func NewX509CertificateCRLValidationConfiguration()(*X509CertificateCRLValidationConfiguration) {
    m := &X509CertificateCRLValidationConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateX509CertificateCRLValidationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateX509CertificateCRLValidationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewX509CertificateCRLValidationConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *X509CertificateCRLValidationConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExemptedCertificateAuthoritiesSubjectKeyIdentifiers gets the exemptedCertificateAuthoritiesSubjectKeyIdentifiers property value. The exemptedCertificateAuthoritiesSubjectKeyIdentifiers property
// returns a []string when successful
func (m *X509CertificateCRLValidationConfiguration) GetExemptedCertificateAuthoritiesSubjectKeyIdentifiers()([]string) {
    return m.exemptedCertificateAuthoritiesSubjectKeyIdentifiers
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *X509CertificateCRLValidationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["exemptedCertificateAuthoritiesSubjectKeyIdentifiers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetExemptedCertificateAuthoritiesSubjectKeyIdentifiers(res)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseX509CertificateCRLValidationConfigurationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*X509CertificateCRLValidationConfigurationState))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *X509CertificateCRLValidationConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetState gets the state property value. The state property
// returns a *X509CertificateCRLValidationConfigurationState when successful
func (m *X509CertificateCRLValidationConfiguration) GetState()(*X509CertificateCRLValidationConfigurationState) {
    return m.state
}
// Serialize serializes information the current object
func (m *X509CertificateCRLValidationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExemptedCertificateAuthoritiesSubjectKeyIdentifiers() != nil {
        err := writer.WriteCollectionOfStringValues("exemptedCertificateAuthoritiesSubjectKeyIdentifiers", m.GetExemptedCertificateAuthoritiesSubjectKeyIdentifiers())
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
func (m *X509CertificateCRLValidationConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExemptedCertificateAuthoritiesSubjectKeyIdentifiers sets the exemptedCertificateAuthoritiesSubjectKeyIdentifiers property value. The exemptedCertificateAuthoritiesSubjectKeyIdentifiers property
func (m *X509CertificateCRLValidationConfiguration) SetExemptedCertificateAuthoritiesSubjectKeyIdentifiers(value []string)() {
    m.exemptedCertificateAuthoritiesSubjectKeyIdentifiers = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *X509CertificateCRLValidationConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetState sets the state property value. The state property
func (m *X509CertificateCRLValidationConfiguration) SetState(value *X509CertificateCRLValidationConfigurationState)() {
    m.state = value
}
type X509CertificateCRLValidationConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExemptedCertificateAuthoritiesSubjectKeyIdentifiers()([]string)
    GetOdataType()(*string)
    GetState()(*X509CertificateCRLValidationConfigurationState)
    SetExemptedCertificateAuthoritiesSubjectKeyIdentifiers(value []string)()
    SetOdataType(value *string)()
    SetState(value *X509CertificateCRLValidationConfigurationState)()
}
