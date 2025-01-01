package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FilePlanDescriptorBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique string that defines the name for the file plan descriptor associated with a particular retention label.
    displayName *string
    // The OdataType property
    odataType *string
}
// NewFilePlanDescriptorBase instantiates a new FilePlanDescriptorBase and sets the default values.
func NewFilePlanDescriptorBase()(*FilePlanDescriptorBase) {
    m := &FilePlanDescriptorBase{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFilePlanDescriptorBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilePlanDescriptorBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.security.filePlanAppliedCategory":
                        return NewFilePlanAppliedCategory(), nil
                    case "#microsoft.graph.security.filePlanAuthority":
                        return NewFilePlanAuthority(), nil
                    case "#microsoft.graph.security.filePlanCitation":
                        return NewFilePlanCitation(), nil
                    case "#microsoft.graph.security.filePlanDepartment":
                        return NewFilePlanDepartment(), nil
                    case "#microsoft.graph.security.filePlanReference":
                        return NewFilePlanReference(), nil
                    case "#microsoft.graph.security.filePlanSubcategory":
                        return NewFilePlanSubcategory(), nil
                }
            }
        }
    }
    return NewFilePlanDescriptorBase(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FilePlanDescriptorBase) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Unique string that defines the name for the file plan descriptor associated with a particular retention label.
// returns a *string when successful
func (m *FilePlanDescriptorBase) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FilePlanDescriptorBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
func (m *FilePlanDescriptorBase) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *FilePlanDescriptorBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
func (m *FilePlanDescriptorBase) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Unique string that defines the name for the file plan descriptor associated with a particular retention label.
func (m *FilePlanDescriptorBase) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FilePlanDescriptorBase) SetOdataType(value *string)() {
    m.odataType = value
}
type FilePlanDescriptorBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetOdataType()(*string)
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
}
