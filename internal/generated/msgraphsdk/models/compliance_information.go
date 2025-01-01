package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ComplianceInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Collection of the certification controls associated with the certification.
    certificationControls []CertificationControlable
    // The name of the compliance certification, for example, ISO 27018:2014, GDPR, FedRAMP, and NIST 800-171.
    certificationName *string
    // The OdataType property
    odataType *string
}
// NewComplianceInformation instantiates a new ComplianceInformation and sets the default values.
func NewComplianceInformation()(*ComplianceInformation) {
    m := &ComplianceInformation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComplianceInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateComplianceInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComplianceInformation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ComplianceInformation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCertificationControls gets the certificationControls property value. Collection of the certification controls associated with the certification.
// returns a []CertificationControlable when successful
func (m *ComplianceInformation) GetCertificationControls()([]CertificationControlable) {
    return m.certificationControls
}
// GetCertificationName gets the certificationName property value. The name of the compliance certification, for example, ISO 27018:2014, GDPR, FedRAMP, and NIST 800-171.
// returns a *string when successful
func (m *ComplianceInformation) GetCertificationName()(*string) {
    return m.certificationName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ComplianceInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificationControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCertificationControlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificationControlable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CertificationControlable)
                }
            }
            m.SetCertificationControls(res)
        }
        return nil
    }
    res["certificationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationName(val)
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
func (m *ComplianceInformation) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ComplianceInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCertificationControls() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificationControls()))
        for i, v := range m.GetCertificationControls() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("certificationControls", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("certificationName", m.GetCertificationName())
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
func (m *ComplianceInformation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCertificationControls sets the certificationControls property value. Collection of the certification controls associated with the certification.
func (m *ComplianceInformation) SetCertificationControls(value []CertificationControlable)() {
    m.certificationControls = value
}
// SetCertificationName sets the certificationName property value. The name of the compliance certification, for example, ISO 27018:2014, GDPR, FedRAMP, and NIST 800-171.
func (m *ComplianceInformation) SetCertificationName(value *string)() {
    m.certificationName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ComplianceInformation) SetOdataType(value *string)() {
    m.odataType = value
}
type ComplianceInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificationControls()([]CertificationControlable)
    GetCertificationName()(*string)
    GetOdataType()(*string)
    SetCertificationControls(value []CertificationControlable)()
    SetCertificationName(value *string)()
    SetOdataType(value *string)()
}
