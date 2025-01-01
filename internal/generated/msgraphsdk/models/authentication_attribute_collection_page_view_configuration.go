package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthenticationAttributeCollectionPageViewConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description of the page.
    description *string
    // The display configuration of attributes being collected on the attribute collection page. You must specify all attributes that you want to retain, otherwise they're removed from the user flow.
    inputs []AuthenticationAttributeCollectionInputConfigurationable
    // The OdataType property
    odataType *string
    // The title of the attribute collection page.
    title *string
}
// NewAuthenticationAttributeCollectionPageViewConfiguration instantiates a new AuthenticationAttributeCollectionPageViewConfiguration and sets the default values.
func NewAuthenticationAttributeCollectionPageViewConfiguration()(*AuthenticationAttributeCollectionPageViewConfiguration) {
    m := &AuthenticationAttributeCollectionPageViewConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationAttributeCollectionPageViewConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationAttributeCollectionPageViewConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationAttributeCollectionPageViewConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthenticationAttributeCollectionPageViewConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description of the page.
// returns a *string when successful
func (m *AuthenticationAttributeCollectionPageViewConfiguration) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationAttributeCollectionPageViewConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["inputs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationAttributeCollectionInputConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationAttributeCollectionInputConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationAttributeCollectionInputConfigurationable)
                }
            }
            m.SetInputs(res)
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
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetInputs gets the inputs property value. The display configuration of attributes being collected on the attribute collection page. You must specify all attributes that you want to retain, otherwise they're removed from the user flow.
// returns a []AuthenticationAttributeCollectionInputConfigurationable when successful
func (m *AuthenticationAttributeCollectionPageViewConfiguration) GetInputs()([]AuthenticationAttributeCollectionInputConfigurationable) {
    return m.inputs
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AuthenticationAttributeCollectionPageViewConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetTitle gets the title property value. The title of the attribute collection page.
// returns a *string when successful
func (m *AuthenticationAttributeCollectionPageViewConfiguration) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *AuthenticationAttributeCollectionPageViewConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetInputs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInputs()))
        for i, v := range m.GetInputs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("inputs", cast)
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
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *AuthenticationAttributeCollectionPageViewConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description of the page.
func (m *AuthenticationAttributeCollectionPageViewConfiguration) SetDescription(value *string)() {
    m.description = value
}
// SetInputs sets the inputs property value. The display configuration of attributes being collected on the attribute collection page. You must specify all attributes that you want to retain, otherwise they're removed from the user flow.
func (m *AuthenticationAttributeCollectionPageViewConfiguration) SetInputs(value []AuthenticationAttributeCollectionInputConfigurationable)() {
    m.inputs = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationAttributeCollectionPageViewConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTitle sets the title property value. The title of the attribute collection page.
func (m *AuthenticationAttributeCollectionPageViewConfiguration) SetTitle(value *string)() {
    m.title = value
}
type AuthenticationAttributeCollectionPageViewConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetInputs()([]AuthenticationAttributeCollectionInputConfigurationable)
    GetOdataType()(*string)
    GetTitle()(*string)
    SetDescription(value *string)()
    SetInputs(value []AuthenticationAttributeCollectionInputConfigurationable)()
    SetOdataType(value *string)()
    SetTitle(value *string)()
}
