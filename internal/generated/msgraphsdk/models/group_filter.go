package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupFilter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The includedGroups property
    includedGroups []string
    // The OdataType property
    odataType *string
}
// NewGroupFilter instantiates a new GroupFilter and sets the default values.
func NewGroupFilter()(*GroupFilter) {
    m := &GroupFilter{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGroupFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupFilter(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GroupFilter) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["includedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetIncludedGroups(res)
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
// GetIncludedGroups gets the includedGroups property value. The includedGroups property
// returns a []string when successful
func (m *GroupFilter) GetIncludedGroups()([]string) {
    return m.includedGroups
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *GroupFilter) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *GroupFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIncludedGroups() != nil {
        err := writer.WriteCollectionOfStringValues("includedGroups", m.GetIncludedGroups())
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
func (m *GroupFilter) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIncludedGroups sets the includedGroups property value. The includedGroups property
func (m *GroupFilter) SetIncludedGroups(value []string)() {
    m.includedGroups = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GroupFilter) SetOdataType(value *string)() {
    m.odataType = value
}
type GroupFilterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIncludedGroups()([]string)
    GetOdataType()(*string)
    SetIncludedGroups(value []string)()
    SetOdataType(value *string)()
}
