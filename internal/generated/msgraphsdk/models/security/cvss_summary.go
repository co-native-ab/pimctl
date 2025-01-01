package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CvssSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The CVSS score about this vulnerability.
    score *float64
    // The CVSS severity rating for this vulnerability. The possible values are: none, low, medium, high, critical, unknownFutureValue.
    severity *VulnerabilitySeverity
    // The CVSS vector string for this vulnerability.
    vectorString *string
}
// NewCvssSummary instantiates a new CvssSummary and sets the default values.
func NewCvssSummary()(*CvssSummary) {
    m := &CvssSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCvssSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCvssSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCvssSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CvssSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CvssSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVulnerabilitySeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*VulnerabilitySeverity))
        }
        return nil
    }
    res["vectorString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVectorString(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CvssSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetScore gets the score property value. The CVSS score about this vulnerability.
// returns a *float64 when successful
func (m *CvssSummary) GetScore()(*float64) {
    return m.score
}
// GetSeverity gets the severity property value. The CVSS severity rating for this vulnerability. The possible values are: none, low, medium, high, critical, unknownFutureValue.
// returns a *VulnerabilitySeverity when successful
func (m *CvssSummary) GetSeverity()(*VulnerabilitySeverity) {
    return m.severity
}
// GetVectorString gets the vectorString property value. The CVSS vector string for this vulnerability.
// returns a *string when successful
func (m *CvssSummary) GetVectorString()(*string) {
    return m.vectorString
}
// Serialize serializes information the current object
func (m *CvssSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("score", m.GetScore())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err := writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vectorString", m.GetVectorString())
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
func (m *CvssSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CvssSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetScore sets the score property value. The CVSS score about this vulnerability.
func (m *CvssSummary) SetScore(value *float64)() {
    m.score = value
}
// SetSeverity sets the severity property value. The CVSS severity rating for this vulnerability. The possible values are: none, low, medium, high, critical, unknownFutureValue.
func (m *CvssSummary) SetSeverity(value *VulnerabilitySeverity)() {
    m.severity = value
}
// SetVectorString sets the vectorString property value. The CVSS vector string for this vulnerability.
func (m *CvssSummary) SetVectorString(value *string)() {
    m.vectorString = value
}
type CvssSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetScore()(*float64)
    GetSeverity()(*VulnerabilitySeverity)
    GetVectorString()(*string)
    SetOdataType(value *string)()
    SetScore(value *float64)()
    SetSeverity(value *VulnerabilitySeverity)()
    SetVectorString(value *string)()
}
