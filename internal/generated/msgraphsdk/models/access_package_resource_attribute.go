package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageResourceAttribute struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Information about how to set the attribute, currently a accessPackageUserDirectoryAttributeStore type.
    destination AccessPackageResourceAttributeDestinationable
    // The isEditable property
    isEditable *bool
    // The isPersistedOnAssignmentRemoval property
    isPersistedOnAssignmentRemoval *bool
    // The name of the attribute in the end system. If the destination is accessPackageUserDirectoryAttributeStore, then a user property such as jobTitle or a directory schema extension for the user object type, such as extension2b676109c7c74ae2b41549205f1947edpersonalTitle.
    name *string
    // The OdataType property
    odataType *string
    // Information about how to populate the attribute value when an accessPackageAssignmentRequest is being fulfilled, currently a accessPackageResourceAttributeQuestion type.
    source AccessPackageResourceAttributeSourceable
}
// NewAccessPackageResourceAttribute instantiates a new AccessPackageResourceAttribute and sets the default values.
func NewAccessPackageResourceAttribute()(*AccessPackageResourceAttribute) {
    m := &AccessPackageResourceAttribute{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessPackageResourceAttributeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageResourceAttributeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceAttribute(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccessPackageResourceAttribute) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDestination gets the destination property value. Information about how to set the attribute, currently a accessPackageUserDirectoryAttributeStore type.
// returns a AccessPackageResourceAttributeDestinationable when successful
func (m *AccessPackageResourceAttribute) GetDestination()(AccessPackageResourceAttributeDestinationable) {
    return m.destination
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageResourceAttribute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceAttributeDestinationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestination(val.(AccessPackageResourceAttributeDestinationable))
        }
        return nil
    }
    res["isEditable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEditable(val)
        }
        return nil
    }
    res["isPersistedOnAssignmentRemoval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPersistedOnAssignmentRemoval(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceAttributeSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(AccessPackageResourceAttributeSourceable))
        }
        return nil
    }
    return res
}
// GetIsEditable gets the isEditable property value. The isEditable property
// returns a *bool when successful
func (m *AccessPackageResourceAttribute) GetIsEditable()(*bool) {
    return m.isEditable
}
// GetIsPersistedOnAssignmentRemoval gets the isPersistedOnAssignmentRemoval property value. The isPersistedOnAssignmentRemoval property
// returns a *bool when successful
func (m *AccessPackageResourceAttribute) GetIsPersistedOnAssignmentRemoval()(*bool) {
    return m.isPersistedOnAssignmentRemoval
}
// GetName gets the name property value. The name of the attribute in the end system. If the destination is accessPackageUserDirectoryAttributeStore, then a user property such as jobTitle or a directory schema extension for the user object type, such as extension2b676109c7c74ae2b41549205f1947edpersonalTitle.
// returns a *string when successful
func (m *AccessPackageResourceAttribute) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AccessPackageResourceAttribute) GetOdataType()(*string) {
    return m.odataType
}
// GetSource gets the source property value. Information about how to populate the attribute value when an accessPackageAssignmentRequest is being fulfilled, currently a accessPackageResourceAttributeQuestion type.
// returns a AccessPackageResourceAttributeSourceable when successful
func (m *AccessPackageResourceAttribute) GetSource()(AccessPackageResourceAttributeSourceable) {
    return m.source
}
// Serialize serializes information the current object
func (m *AccessPackageResourceAttribute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("destination", m.GetDestination())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEditable", m.GetIsEditable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPersistedOnAssignmentRemoval", m.GetIsPersistedOnAssignmentRemoval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteObjectValue("source", m.GetSource())
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
func (m *AccessPackageResourceAttribute) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDestination sets the destination property value. Information about how to set the attribute, currently a accessPackageUserDirectoryAttributeStore type.
func (m *AccessPackageResourceAttribute) SetDestination(value AccessPackageResourceAttributeDestinationable)() {
    m.destination = value
}
// SetIsEditable sets the isEditable property value. The isEditable property
func (m *AccessPackageResourceAttribute) SetIsEditable(value *bool)() {
    m.isEditable = value
}
// SetIsPersistedOnAssignmentRemoval sets the isPersistedOnAssignmentRemoval property value. The isPersistedOnAssignmentRemoval property
func (m *AccessPackageResourceAttribute) SetIsPersistedOnAssignmentRemoval(value *bool)() {
    m.isPersistedOnAssignmentRemoval = value
}
// SetName sets the name property value. The name of the attribute in the end system. If the destination is accessPackageUserDirectoryAttributeStore, then a user property such as jobTitle or a directory schema extension for the user object type, such as extension2b676109c7c74ae2b41549205f1947edpersonalTitle.
func (m *AccessPackageResourceAttribute) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageResourceAttribute) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSource sets the source property value. Information about how to populate the attribute value when an accessPackageAssignmentRequest is being fulfilled, currently a accessPackageResourceAttributeQuestion type.
func (m *AccessPackageResourceAttribute) SetSource(value AccessPackageResourceAttributeSourceable)() {
    m.source = value
}
type AccessPackageResourceAttributeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDestination()(AccessPackageResourceAttributeDestinationable)
    GetIsEditable()(*bool)
    GetIsPersistedOnAssignmentRemoval()(*bool)
    GetName()(*string)
    GetOdataType()(*string)
    GetSource()(AccessPackageResourceAttributeSourceable)
    SetDestination(value AccessPackageResourceAttributeDestinationable)()
    SetIsEditable(value *bool)()
    SetIsPersistedOnAssignmentRemoval(value *bool)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetSource(value AccessPackageResourceAttributeSourceable)()
}
