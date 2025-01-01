package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type X509CertificateRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The identifier of the X.509 certificate. Required.
    identifier *string
    // The issuerSubjectIdentifier property
    issuerSubjectIdentifier *string
    // The OdataType property
    odataType *string
    // The policyOidIdentifier property
    policyOidIdentifier *string
    // The type of strong authentication mode. The possible values are: x509CertificateSingleFactor, x509CertificateMultiFactor, unknownFutureValue. Required.
    x509CertificateAuthenticationMode *X509CertificateAuthenticationMode
    // The x509CertificateRequiredAffinityLevel property
    x509CertificateRequiredAffinityLevel *X509CertificateAffinityLevel
    // The type of the X.509 certificate mode configuration rule. The possible values are: issuerSubject, policyOID, unknownFutureValue. Required.
    x509CertificateRuleType *X509CertificateRuleType
}
// NewX509CertificateRule instantiates a new X509CertificateRule and sets the default values.
func NewX509CertificateRule()(*X509CertificateRule) {
    m := &X509CertificateRule{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateX509CertificateRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateX509CertificateRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewX509CertificateRule(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *X509CertificateRule) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *X509CertificateRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    res["issuerSubjectIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerSubjectIdentifier(val)
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
    res["policyOidIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyOidIdentifier(val)
        }
        return nil
    }
    res["x509CertificateAuthenticationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseX509CertificateAuthenticationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX509CertificateAuthenticationMode(val.(*X509CertificateAuthenticationMode))
        }
        return nil
    }
    res["x509CertificateRequiredAffinityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseX509CertificateAffinityLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX509CertificateRequiredAffinityLevel(val.(*X509CertificateAffinityLevel))
        }
        return nil
    }
    res["x509CertificateRuleType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseX509CertificateRuleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX509CertificateRuleType(val.(*X509CertificateRuleType))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. The identifier of the X.509 certificate. Required.
// returns a *string when successful
func (m *X509CertificateRule) GetIdentifier()(*string) {
    return m.identifier
}
// GetIssuerSubjectIdentifier gets the issuerSubjectIdentifier property value. The issuerSubjectIdentifier property
// returns a *string when successful
func (m *X509CertificateRule) GetIssuerSubjectIdentifier()(*string) {
    return m.issuerSubjectIdentifier
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *X509CertificateRule) GetOdataType()(*string) {
    return m.odataType
}
// GetPolicyOidIdentifier gets the policyOidIdentifier property value. The policyOidIdentifier property
// returns a *string when successful
func (m *X509CertificateRule) GetPolicyOidIdentifier()(*string) {
    return m.policyOidIdentifier
}
// GetX509CertificateAuthenticationMode gets the x509CertificateAuthenticationMode property value. The type of strong authentication mode. The possible values are: x509CertificateSingleFactor, x509CertificateMultiFactor, unknownFutureValue. Required.
// returns a *X509CertificateAuthenticationMode when successful
func (m *X509CertificateRule) GetX509CertificateAuthenticationMode()(*X509CertificateAuthenticationMode) {
    return m.x509CertificateAuthenticationMode
}
// GetX509CertificateRequiredAffinityLevel gets the x509CertificateRequiredAffinityLevel property value. The x509CertificateRequiredAffinityLevel property
// returns a *X509CertificateAffinityLevel when successful
func (m *X509CertificateRule) GetX509CertificateRequiredAffinityLevel()(*X509CertificateAffinityLevel) {
    return m.x509CertificateRequiredAffinityLevel
}
// GetX509CertificateRuleType gets the x509CertificateRuleType property value. The type of the X.509 certificate mode configuration rule. The possible values are: issuerSubject, policyOID, unknownFutureValue. Required.
// returns a *X509CertificateRuleType when successful
func (m *X509CertificateRule) GetX509CertificateRuleType()(*X509CertificateRuleType) {
    return m.x509CertificateRuleType
}
// Serialize serializes information the current object
func (m *X509CertificateRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("identifier", m.GetIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issuerSubjectIdentifier", m.GetIssuerSubjectIdentifier())
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
        err := writer.WriteStringValue("policyOidIdentifier", m.GetPolicyOidIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetX509CertificateAuthenticationMode() != nil {
        cast := (*m.GetX509CertificateAuthenticationMode()).String()
        err := writer.WriteStringValue("x509CertificateAuthenticationMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetX509CertificateRequiredAffinityLevel() != nil {
        cast := (*m.GetX509CertificateRequiredAffinityLevel()).String()
        err := writer.WriteStringValue("x509CertificateRequiredAffinityLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetX509CertificateRuleType() != nil {
        cast := (*m.GetX509CertificateRuleType()).String()
        err := writer.WriteStringValue("x509CertificateRuleType", &cast)
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
func (m *X509CertificateRule) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIdentifier sets the identifier property value. The identifier of the X.509 certificate. Required.
func (m *X509CertificateRule) SetIdentifier(value *string)() {
    m.identifier = value
}
// SetIssuerSubjectIdentifier sets the issuerSubjectIdentifier property value. The issuerSubjectIdentifier property
func (m *X509CertificateRule) SetIssuerSubjectIdentifier(value *string)() {
    m.issuerSubjectIdentifier = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *X509CertificateRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPolicyOidIdentifier sets the policyOidIdentifier property value. The policyOidIdentifier property
func (m *X509CertificateRule) SetPolicyOidIdentifier(value *string)() {
    m.policyOidIdentifier = value
}
// SetX509CertificateAuthenticationMode sets the x509CertificateAuthenticationMode property value. The type of strong authentication mode. The possible values are: x509CertificateSingleFactor, x509CertificateMultiFactor, unknownFutureValue. Required.
func (m *X509CertificateRule) SetX509CertificateAuthenticationMode(value *X509CertificateAuthenticationMode)() {
    m.x509CertificateAuthenticationMode = value
}
// SetX509CertificateRequiredAffinityLevel sets the x509CertificateRequiredAffinityLevel property value. The x509CertificateRequiredAffinityLevel property
func (m *X509CertificateRule) SetX509CertificateRequiredAffinityLevel(value *X509CertificateAffinityLevel)() {
    m.x509CertificateRequiredAffinityLevel = value
}
// SetX509CertificateRuleType sets the x509CertificateRuleType property value. The type of the X.509 certificate mode configuration rule. The possible values are: issuerSubject, policyOID, unknownFutureValue. Required.
func (m *X509CertificateRule) SetX509CertificateRuleType(value *X509CertificateRuleType)() {
    m.x509CertificateRuleType = value
}
type X509CertificateRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentifier()(*string)
    GetIssuerSubjectIdentifier()(*string)
    GetOdataType()(*string)
    GetPolicyOidIdentifier()(*string)
    GetX509CertificateAuthenticationMode()(*X509CertificateAuthenticationMode)
    GetX509CertificateRequiredAffinityLevel()(*X509CertificateAffinityLevel)
    GetX509CertificateRuleType()(*X509CertificateRuleType)
    SetIdentifier(value *string)()
    SetIssuerSubjectIdentifier(value *string)()
    SetOdataType(value *string)()
    SetPolicyOidIdentifier(value *string)()
    SetX509CertificateAuthenticationMode(value *X509CertificateAuthenticationMode)()
    SetX509CertificateRequiredAffinityLevel(value *X509CertificateAffinityLevel)()
    SetX509CertificateRuleType(value *X509CertificateRuleType)()
}
