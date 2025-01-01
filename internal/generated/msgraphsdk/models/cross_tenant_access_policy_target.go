package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CrossTenantAccessPolicyTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Defines the target for cross-tenant access policy settings and can have one of the following values:  The unique identifier of the user, group, or application  AllUsers  AllApplications - Refers to any Microsoft cloud application.  Office365 - Includes the applications mentioned as part of the Office 365 suite.
    target *string
    // The type of resource that you want to target. The possible values are: user, group, application, unknownFutureValue.
    targetType *CrossTenantAccessPolicyTargetType
}
// NewCrossTenantAccessPolicyTarget instantiates a new CrossTenantAccessPolicyTarget and sets the default values.
func NewCrossTenantAccessPolicyTarget()(*CrossTenantAccessPolicyTarget) {
    m := &CrossTenantAccessPolicyTarget{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCrossTenantAccessPolicyTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCrossTenantAccessPolicyTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicyTarget(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CrossTenantAccessPolicyTarget) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CrossTenantAccessPolicyTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    res["targetType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCrossTenantAccessPolicyTargetType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetType(val.(*CrossTenantAccessPolicyTargetType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CrossTenantAccessPolicyTarget) GetOdataType()(*string) {
    return m.odataType
}
// GetTarget gets the target property value. Defines the target for cross-tenant access policy settings and can have one of the following values:  The unique identifier of the user, group, or application  AllUsers  AllApplications - Refers to any Microsoft cloud application.  Office365 - Includes the applications mentioned as part of the Office 365 suite.
// returns a *string when successful
func (m *CrossTenantAccessPolicyTarget) GetTarget()(*string) {
    return m.target
}
// GetTargetType gets the targetType property value. The type of resource that you want to target. The possible values are: user, group, application, unknownFutureValue.
// returns a *CrossTenantAccessPolicyTargetType when successful
func (m *CrossTenantAccessPolicyTarget) GetTargetType()(*CrossTenantAccessPolicyTargetType) {
    return m.targetType
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := (*m.GetTargetType()).String()
        err := writer.WriteStringValue("targetType", &cast)
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
func (m *CrossTenantAccessPolicyTarget) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CrossTenantAccessPolicyTarget) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTarget sets the target property value. Defines the target for cross-tenant access policy settings and can have one of the following values:  The unique identifier of the user, group, or application  AllUsers  AllApplications - Refers to any Microsoft cloud application.  Office365 - Includes the applications mentioned as part of the Office 365 suite.
func (m *CrossTenantAccessPolicyTarget) SetTarget(value *string)() {
    m.target = value
}
// SetTargetType sets the targetType property value. The type of resource that you want to target. The possible values are: user, group, application, unknownFutureValue.
func (m *CrossTenantAccessPolicyTarget) SetTargetType(value *CrossTenantAccessPolicyTargetType)() {
    m.targetType = value
}
type CrossTenantAccessPolicyTargetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetTarget()(*string)
    GetTargetType()(*CrossTenantAccessPolicyTargetType)
    SetOdataType(value *string)()
    SetTarget(value *string)()
    SetTargetType(value *CrossTenantAccessPolicyTargetType)()
}
