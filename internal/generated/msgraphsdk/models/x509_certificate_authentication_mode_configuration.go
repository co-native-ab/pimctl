package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type X509CertificateAuthenticationModeConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Rules are configured in addition to the authentication mode to bind a specific x509CertificateRuleType to an x509CertificateAuthenticationMode. For example, bind the policyOID with identifier 1.32.132.343 to x509CertificateMultiFactor authentication mode.
    rules []X509CertificateRuleable
    // The type of strong authentication mode. The possible values are: x509CertificateSingleFactor, x509CertificateMultiFactor, unknownFutureValue.
    x509CertificateAuthenticationDefaultMode *X509CertificateAuthenticationMode
    // The x509CertificateDefaultRequiredAffinityLevel property
    x509CertificateDefaultRequiredAffinityLevel *X509CertificateAffinityLevel
}
// NewX509CertificateAuthenticationModeConfiguration instantiates a new X509CertificateAuthenticationModeConfiguration and sets the default values.
func NewX509CertificateAuthenticationModeConfiguration()(*X509CertificateAuthenticationModeConfiguration) {
    m := &X509CertificateAuthenticationModeConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateX509CertificateAuthenticationModeConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateX509CertificateAuthenticationModeConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewX509CertificateAuthenticationModeConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *X509CertificateAuthenticationModeConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *X509CertificateAuthenticationModeConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateX509CertificateRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]X509CertificateRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(X509CertificateRuleable)
                }
            }
            m.SetRules(res)
        }
        return nil
    }
    res["x509CertificateAuthenticationDefaultMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseX509CertificateAuthenticationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX509CertificateAuthenticationDefaultMode(val.(*X509CertificateAuthenticationMode))
        }
        return nil
    }
    res["x509CertificateDefaultRequiredAffinityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseX509CertificateAffinityLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX509CertificateDefaultRequiredAffinityLevel(val.(*X509CertificateAffinityLevel))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *X509CertificateAuthenticationModeConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetRules gets the rules property value. Rules are configured in addition to the authentication mode to bind a specific x509CertificateRuleType to an x509CertificateAuthenticationMode. For example, bind the policyOID with identifier 1.32.132.343 to x509CertificateMultiFactor authentication mode.
// returns a []X509CertificateRuleable when successful
func (m *X509CertificateAuthenticationModeConfiguration) GetRules()([]X509CertificateRuleable) {
    return m.rules
}
// GetX509CertificateAuthenticationDefaultMode gets the x509CertificateAuthenticationDefaultMode property value. The type of strong authentication mode. The possible values are: x509CertificateSingleFactor, x509CertificateMultiFactor, unknownFutureValue.
// returns a *X509CertificateAuthenticationMode when successful
func (m *X509CertificateAuthenticationModeConfiguration) GetX509CertificateAuthenticationDefaultMode()(*X509CertificateAuthenticationMode) {
    return m.x509CertificateAuthenticationDefaultMode
}
// GetX509CertificateDefaultRequiredAffinityLevel gets the x509CertificateDefaultRequiredAffinityLevel property value. The x509CertificateDefaultRequiredAffinityLevel property
// returns a *X509CertificateAffinityLevel when successful
func (m *X509CertificateAuthenticationModeConfiguration) GetX509CertificateDefaultRequiredAffinityLevel()(*X509CertificateAffinityLevel) {
    return m.x509CertificateDefaultRequiredAffinityLevel
}
// Serialize serializes information the current object
func (m *X509CertificateAuthenticationModeConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetX509CertificateAuthenticationDefaultMode() != nil {
        cast := (*m.GetX509CertificateAuthenticationDefaultMode()).String()
        err := writer.WriteStringValue("x509CertificateAuthenticationDefaultMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetX509CertificateDefaultRequiredAffinityLevel() != nil {
        cast := (*m.GetX509CertificateDefaultRequiredAffinityLevel()).String()
        err := writer.WriteStringValue("x509CertificateDefaultRequiredAffinityLevel", &cast)
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
func (m *X509CertificateAuthenticationModeConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *X509CertificateAuthenticationModeConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRules sets the rules property value. Rules are configured in addition to the authentication mode to bind a specific x509CertificateRuleType to an x509CertificateAuthenticationMode. For example, bind the policyOID with identifier 1.32.132.343 to x509CertificateMultiFactor authentication mode.
func (m *X509CertificateAuthenticationModeConfiguration) SetRules(value []X509CertificateRuleable)() {
    m.rules = value
}
// SetX509CertificateAuthenticationDefaultMode sets the x509CertificateAuthenticationDefaultMode property value. The type of strong authentication mode. The possible values are: x509CertificateSingleFactor, x509CertificateMultiFactor, unknownFutureValue.
func (m *X509CertificateAuthenticationModeConfiguration) SetX509CertificateAuthenticationDefaultMode(value *X509CertificateAuthenticationMode)() {
    m.x509CertificateAuthenticationDefaultMode = value
}
// SetX509CertificateDefaultRequiredAffinityLevel sets the x509CertificateDefaultRequiredAffinityLevel property value. The x509CertificateDefaultRequiredAffinityLevel property
func (m *X509CertificateAuthenticationModeConfiguration) SetX509CertificateDefaultRequiredAffinityLevel(value *X509CertificateAffinityLevel)() {
    m.x509CertificateDefaultRequiredAffinityLevel = value
}
type X509CertificateAuthenticationModeConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetRules()([]X509CertificateRuleable)
    GetX509CertificateAuthenticationDefaultMode()(*X509CertificateAuthenticationMode)
    GetX509CertificateDefaultRequiredAffinityLevel()(*X509CertificateAffinityLevel)
    SetOdataType(value *string)()
    SetRules(value []X509CertificateRuleable)()
    SetX509CertificateAuthenticationDefaultMode(value *X509CertificateAuthenticationMode)()
    SetX509CertificateDefaultRequiredAffinityLevel(value *X509CertificateAffinityLevel)()
}
