package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthenticationConditionsApplications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The includeApplications property
    includeApplications []AuthenticationConditionApplicationable
    // The OdataType property
    odataType *string
}
// NewAuthenticationConditionsApplications instantiates a new AuthenticationConditionsApplications and sets the default values.
func NewAuthenticationConditionsApplications()(*AuthenticationConditionsApplications) {
    m := &AuthenticationConditionsApplications{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationConditionsApplicationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationConditionsApplicationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationConditionsApplications(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthenticationConditionsApplications) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationConditionsApplications) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["includeApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationConditionApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationConditionApplicationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationConditionApplicationable)
                }
            }
            m.SetIncludeApplications(res)
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
// GetIncludeApplications gets the includeApplications property value. The includeApplications property
// returns a []AuthenticationConditionApplicationable when successful
func (m *AuthenticationConditionsApplications) GetIncludeApplications()([]AuthenticationConditionApplicationable) {
    return m.includeApplications
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AuthenticationConditionsApplications) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AuthenticationConditionsApplications) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIncludeApplications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludeApplications()))
        for i, v := range m.GetIncludeApplications() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("includeApplications", cast)
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
func (m *AuthenticationConditionsApplications) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIncludeApplications sets the includeApplications property value. The includeApplications property
func (m *AuthenticationConditionsApplications) SetIncludeApplications(value []AuthenticationConditionApplicationable)() {
    m.includeApplications = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationConditionsApplications) SetOdataType(value *string)() {
    m.odataType = value
}
type AuthenticationConditionsApplicationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIncludeApplications()([]AuthenticationConditionApplicationable)
    GetOdataType()(*string)
    SetIncludeApplications(value []AuthenticationConditionApplicationable)()
    SetOdataType(value *string)()
}
