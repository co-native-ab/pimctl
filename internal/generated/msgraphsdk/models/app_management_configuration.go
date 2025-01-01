package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AppManagementConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Collection of keyCredential restrictions settings to be applied to an application or service principal.
    keyCredentials []KeyCredentialConfigurationable
    // The OdataType property
    odataType *string
    // Collection of password restrictions settings to be applied to an application or service principal.
    passwordCredentials []PasswordCredentialConfigurationable
}
// NewAppManagementConfiguration instantiates a new AppManagementConfiguration and sets the default values.
func NewAppManagementConfiguration()(*AppManagementConfiguration) {
    m := &AppManagementConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppManagementConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppManagementConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.appManagementApplicationConfiguration":
                        return NewAppManagementApplicationConfiguration(), nil
                    case "#microsoft.graph.appManagementServicePrincipalConfiguration":
                        return NewAppManagementServicePrincipalConfiguration(), nil
                    case "#microsoft.graph.customAppManagementConfiguration":
                        return NewCustomAppManagementConfiguration(), nil
                }
            }
        }
    }
    return NewAppManagementConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AppManagementConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppManagementConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyCredentialConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyCredentialConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyCredentialConfigurationable)
                }
            }
            m.SetKeyCredentials(res)
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
    res["passwordCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePasswordCredentialConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordCredentialConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PasswordCredentialConfigurationable)
                }
            }
            m.SetPasswordCredentials(res)
        }
        return nil
    }
    return res
}
// GetKeyCredentials gets the keyCredentials property value. Collection of keyCredential restrictions settings to be applied to an application or service principal.
// returns a []KeyCredentialConfigurationable when successful
func (m *AppManagementConfiguration) GetKeyCredentials()([]KeyCredentialConfigurationable) {
    return m.keyCredentials
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AppManagementConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetPasswordCredentials gets the passwordCredentials property value. Collection of password restrictions settings to be applied to an application or service principal.
// returns a []PasswordCredentialConfigurationable when successful
func (m *AppManagementConfiguration) GetPasswordCredentials()([]PasswordCredentialConfigurationable) {
    return m.passwordCredentials
}
// Serialize serializes information the current object
func (m *AppManagementConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetKeyCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKeyCredentials()))
        for i, v := range m.GetKeyCredentials() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("keyCredentials", cast)
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
    if m.GetPasswordCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPasswordCredentials()))
        for i, v := range m.GetPasswordCredentials() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("passwordCredentials", cast)
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
func (m *AppManagementConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKeyCredentials sets the keyCredentials property value. Collection of keyCredential restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) SetKeyCredentials(value []KeyCredentialConfigurationable)() {
    m.keyCredentials = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppManagementConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPasswordCredentials sets the passwordCredentials property value. Collection of password restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) SetPasswordCredentials(value []PasswordCredentialConfigurationable)() {
    m.passwordCredentials = value
}
type AppManagementConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeyCredentials()([]KeyCredentialConfigurationable)
    GetOdataType()(*string)
    GetPasswordCredentials()([]PasswordCredentialConfigurationable)
    SetKeyCredentials(value []KeyCredentialConfigurationable)()
    SetOdataType(value *string)()
    SetPasswordCredentials(value []PasswordCredentialConfigurationable)()
}
