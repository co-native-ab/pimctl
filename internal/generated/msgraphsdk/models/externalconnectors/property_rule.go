package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type PropertyRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The operation property
    operation *RuleOperation
    // The property from the externalItem schema. Required.
    property *string
    // A collection with one or many strings. One or more specified strings are matched with the specified property using the specified operation. Required.
    values []string
    // The valuesJoinedBy property
    valuesJoinedBy *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BinaryOperator
}
// NewPropertyRule instantiates a new PropertyRule and sets the default values.
func NewPropertyRule()(*PropertyRule) {
    m := &PropertyRule{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePropertyRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePropertyRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPropertyRule(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PropertyRule) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PropertyRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["operation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleOperation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperation(val.(*RuleOperation))
        }
        return nil
    }
    res["property"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperty(val)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetValues(res)
        }
        return nil
    }
    res["valuesJoinedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.ParseBinaryOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValuesJoinedBy(val.(*ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BinaryOperator))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *PropertyRule) GetOdataType()(*string) {
    return m.odataType
}
// GetOperation gets the operation property value. The operation property
// returns a *RuleOperation when successful
func (m *PropertyRule) GetOperation()(*RuleOperation) {
    return m.operation
}
// GetProperty gets the property property value. The property from the externalItem schema. Required.
// returns a *string when successful
func (m *PropertyRule) GetProperty()(*string) {
    return m.property
}
// GetValues gets the values property value. A collection with one or many strings. One or more specified strings are matched with the specified property using the specified operation. Required.
// returns a []string when successful
func (m *PropertyRule) GetValues()([]string) {
    return m.values
}
// GetValuesJoinedBy gets the valuesJoinedBy property value. The valuesJoinedBy property
// returns a *BinaryOperator when successful
func (m *PropertyRule) GetValuesJoinedBy()(*ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BinaryOperator) {
    return m.valuesJoinedBy
}
// Serialize serializes information the current object
func (m *PropertyRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOperation() != nil {
        cast := (*m.GetOperation()).String()
        err := writer.WriteStringValue("operation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("property", m.GetProperty())
        if err != nil {
            return err
        }
    }
    if m.GetValues() != nil {
        err := writer.WriteCollectionOfStringValues("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    if m.GetValuesJoinedBy() != nil {
        cast := (*m.GetValuesJoinedBy()).String()
        err := writer.WriteStringValue("valuesJoinedBy", &cast)
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
func (m *PropertyRule) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PropertyRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperation sets the operation property value. The operation property
func (m *PropertyRule) SetOperation(value *RuleOperation)() {
    m.operation = value
}
// SetProperty sets the property property value. The property from the externalItem schema. Required.
func (m *PropertyRule) SetProperty(value *string)() {
    m.property = value
}
// SetValues sets the values property value. A collection with one or many strings. One or more specified strings are matched with the specified property using the specified operation. Required.
func (m *PropertyRule) SetValues(value []string)() {
    m.values = value
}
// SetValuesJoinedBy sets the valuesJoinedBy property value. The valuesJoinedBy property
func (m *PropertyRule) SetValuesJoinedBy(value *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BinaryOperator)() {
    m.valuesJoinedBy = value
}
type PropertyRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetOperation()(*RuleOperation)
    GetProperty()(*string)
    GetValues()([]string)
    GetValuesJoinedBy()(*ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BinaryOperator)
    SetOdataType(value *string)()
    SetOperation(value *RuleOperation)()
    SetProperty(value *string)()
    SetValues(value []string)()
    SetValuesJoinedBy(value *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BinaryOperator)()
}
