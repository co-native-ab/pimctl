package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageAssignmentApprovalSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If false, then approval isn't required for new requests in this policy.
    isApprovalRequiredForAdd *bool
    // If false, then approval isn't required for updates to requests in this policy.
    isApprovalRequiredForUpdate *bool
    // The OdataType property
    odataType *string
    // If approval is required, the one, two or three elements of this collection define each of the stages of approval. An empty array is present if no approval is required.
    stages []AccessPackageApprovalStageable
}
// NewAccessPackageAssignmentApprovalSettings instantiates a new AccessPackageAssignmentApprovalSettings and sets the default values.
func NewAccessPackageAssignmentApprovalSettings()(*AccessPackageAssignmentApprovalSettings) {
    m := &AccessPackageAssignmentApprovalSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessPackageAssignmentApprovalSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageAssignmentApprovalSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentApprovalSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccessPackageAssignmentApprovalSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageAssignmentApprovalSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isApprovalRequiredForAdd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequiredForAdd(val)
        }
        return nil
    }
    res["isApprovalRequiredForUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequiredForUpdate(val)
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
    res["stages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageApprovalStageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageApprovalStageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageApprovalStageable)
                }
            }
            m.SetStages(res)
        }
        return nil
    }
    return res
}
// GetIsApprovalRequiredForAdd gets the isApprovalRequiredForAdd property value. If false, then approval isn't required for new requests in this policy.
// returns a *bool when successful
func (m *AccessPackageAssignmentApprovalSettings) GetIsApprovalRequiredForAdd()(*bool) {
    return m.isApprovalRequiredForAdd
}
// GetIsApprovalRequiredForUpdate gets the isApprovalRequiredForUpdate property value. If false, then approval isn't required for updates to requests in this policy.
// returns a *bool when successful
func (m *AccessPackageAssignmentApprovalSettings) GetIsApprovalRequiredForUpdate()(*bool) {
    return m.isApprovalRequiredForUpdate
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AccessPackageAssignmentApprovalSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetStages gets the stages property value. If approval is required, the one, two or three elements of this collection define each of the stages of approval. An empty array is present if no approval is required.
// returns a []AccessPackageApprovalStageable when successful
func (m *AccessPackageAssignmentApprovalSettings) GetStages()([]AccessPackageApprovalStageable) {
    return m.stages
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentApprovalSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isApprovalRequiredForAdd", m.GetIsApprovalRequiredForAdd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequiredForUpdate", m.GetIsApprovalRequiredForUpdate())
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
    if m.GetStages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStages()))
        for i, v := range m.GetStages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("stages", cast)
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
func (m *AccessPackageAssignmentApprovalSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsApprovalRequiredForAdd sets the isApprovalRequiredForAdd property value. If false, then approval isn't required for new requests in this policy.
func (m *AccessPackageAssignmentApprovalSettings) SetIsApprovalRequiredForAdd(value *bool)() {
    m.isApprovalRequiredForAdd = value
}
// SetIsApprovalRequiredForUpdate sets the isApprovalRequiredForUpdate property value. If false, then approval isn't required for updates to requests in this policy.
func (m *AccessPackageAssignmentApprovalSettings) SetIsApprovalRequiredForUpdate(value *bool)() {
    m.isApprovalRequiredForUpdate = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageAssignmentApprovalSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStages sets the stages property value. If approval is required, the one, two or three elements of this collection define each of the stages of approval. An empty array is present if no approval is required.
func (m *AccessPackageAssignmentApprovalSettings) SetStages(value []AccessPackageApprovalStageable)() {
    m.stages = value
}
type AccessPackageAssignmentApprovalSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsApprovalRequiredForAdd()(*bool)
    GetIsApprovalRequiredForUpdate()(*bool)
    GetOdataType()(*string)
    GetStages()([]AccessPackageApprovalStageable)
    SetIsApprovalRequiredForAdd(value *bool)()
    SetIsApprovalRequiredForUpdate(value *bool)()
    SetOdataType(value *string)()
    SetStages(value []AccessPackageApprovalStageable)()
}
