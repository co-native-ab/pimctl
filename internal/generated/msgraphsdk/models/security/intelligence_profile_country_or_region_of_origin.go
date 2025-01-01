package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IntelligenceProfileCountryOrRegionOfOrigin struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A codified representation for this country/region of origin.
    code *string
    // A display label for this ountry/region of origin.
    label *string
    // The OdataType property
    odataType *string
}
// NewIntelligenceProfileCountryOrRegionOfOrigin instantiates a new IntelligenceProfileCountryOrRegionOfOrigin and sets the default values.
func NewIntelligenceProfileCountryOrRegionOfOrigin()(*IntelligenceProfileCountryOrRegionOfOrigin) {
    m := &IntelligenceProfileCountryOrRegionOfOrigin{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIntelligenceProfileCountryOrRegionOfOriginFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntelligenceProfileCountryOrRegionOfOriginFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntelligenceProfileCountryOrRegionOfOrigin(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IntelligenceProfileCountryOrRegionOfOrigin) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. A codified representation for this country/region of origin.
// returns a *string when successful
func (m *IntelligenceProfileCountryOrRegionOfOrigin) GetCode()(*string) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IntelligenceProfileCountryOrRegionOfOrigin) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
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
// GetLabel gets the label property value. A display label for this ountry/region of origin.
// returns a *string when successful
func (m *IntelligenceProfileCountryOrRegionOfOrigin) GetLabel()(*string) {
    return m.label
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *IntelligenceProfileCountryOrRegionOfOrigin) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *IntelligenceProfileCountryOrRegionOfOrigin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("label", m.GetLabel())
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
func (m *IntelligenceProfileCountryOrRegionOfOrigin) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. A codified representation for this country/region of origin.
func (m *IntelligenceProfileCountryOrRegionOfOrigin) SetCode(value *string)() {
    m.code = value
}
// SetLabel sets the label property value. A display label for this ountry/region of origin.
func (m *IntelligenceProfileCountryOrRegionOfOrigin) SetLabel(value *string)() {
    m.label = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IntelligenceProfileCountryOrRegionOfOrigin) SetOdataType(value *string)() {
    m.odataType = value
}
type IntelligenceProfileCountryOrRegionOfOriginable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetLabel()(*string)
    GetOdataType()(*string)
    SetCode(value *string)()
    SetLabel(value *string)()
    SetOdataType(value *string)()
}
