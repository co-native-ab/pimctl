package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcAuditProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The display name for this property.
    displayName *string
    // The new value for this property.
    newValue *string
    // The OdataType property
    odataType *string
    // The old value for this property.
    oldValue *string
}
// NewCloudPcAuditProperty instantiates a new CloudPcAuditProperty and sets the default values.
func NewCloudPcAuditProperty()(*CloudPcAuditProperty) {
    m := &CloudPcAuditProperty{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcAuditPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcAuditPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcAuditProperty(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudPcAuditProperty) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The display name for this property.
// returns a *string when successful
func (m *CloudPcAuditProperty) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcAuditProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["newValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewValue(val)
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
    res["oldValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldValue(val)
        }
        return nil
    }
    return res
}
// GetNewValue gets the newValue property value. The new value for this property.
// returns a *string when successful
func (m *CloudPcAuditProperty) GetNewValue()(*string) {
    return m.newValue
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CloudPcAuditProperty) GetOdataType()(*string) {
    return m.odataType
}
// GetOldValue gets the oldValue property value. The old value for this property.
// returns a *string when successful
func (m *CloudPcAuditProperty) GetOldValue()(*string) {
    return m.oldValue
}
// Serialize serializes information the current object
func (m *CloudPcAuditProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("newValue", m.GetNewValue())
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
        err := writer.WriteStringValue("oldValue", m.GetOldValue())
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
func (m *CloudPcAuditProperty) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The display name for this property.
func (m *CloudPcAuditProperty) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetNewValue sets the newValue property value. The new value for this property.
func (m *CloudPcAuditProperty) SetNewValue(value *string)() {
    m.newValue = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcAuditProperty) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOldValue sets the oldValue property value. The old value for this property.
func (m *CloudPcAuditProperty) SetOldValue(value *string)() {
    m.oldValue = value
}
type CloudPcAuditPropertyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetNewValue()(*string)
    GetOdataType()(*string)
    GetOldValue()(*string)
    SetDisplayName(value *string)()
    SetNewValue(value *string)()
    SetOdataType(value *string)()
    SetOldValue(value *string)()
}
