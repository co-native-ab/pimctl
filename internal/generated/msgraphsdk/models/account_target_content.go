package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccountTargetContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The type of account target content. Possible values are: unknown, includeAll, addressBook, unknownFutureValue.
    typeEscaped *AccountTargetContentType
}
// NewAccountTargetContent instantiates a new AccountTargetContent and sets the default values.
func NewAccountTargetContent()(*AccountTargetContent) {
    m := &AccountTargetContent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccountTargetContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountTargetContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.addressBookAccountTargetContent":
                        return NewAddressBookAccountTargetContent(), nil
                    case "#microsoft.graph.includeAllAccountTargetContent":
                        return NewIncludeAllAccountTargetContent(), nil
                }
            }
        }
    }
    return NewAccountTargetContent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccountTargetContent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountTargetContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccountTargetContentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*AccountTargetContentType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AccountTargetContent) GetOdataType()(*string) {
    return m.odataType
}
// GetTypeEscaped gets the type property value. The type of account target content. Possible values are: unknown, includeAll, addressBook, unknownFutureValue.
// returns a *AccountTargetContentType when successful
func (m *AccountTargetContent) GetTypeEscaped()(*AccountTargetContentType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *AccountTargetContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *AccountTargetContent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccountTargetContent) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTypeEscaped sets the type property value. The type of account target content. Possible values are: unknown, includeAll, addressBook, unknownFutureValue.
func (m *AccountTargetContent) SetTypeEscaped(value *AccountTargetContentType)() {
    m.typeEscaped = value
}
type AccountTargetContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetTypeEscaped()(*AccountTargetContentType)
    SetOdataType(value *string)()
    SetTypeEscaped(value *AccountTargetContentType)()
}
