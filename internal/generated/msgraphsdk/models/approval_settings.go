package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApprovalSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // One of SingleStage, Serial, Parallel, NoApproval (default). NoApproval is used when isApprovalRequired is false.
    approvalMode *string
    // If approval is required, the one or two elements of this collection define each of the stages of approval. An empty array if no approval is required.
    approvalStages []UnifiedApprovalStageable
    // Indicates whether approval is required for requests in this policy.
    isApprovalRequired *bool
    // Indicates whether approval is required for a user to extend their assignment.
    isApprovalRequiredForExtension *bool
    // Indicates whether the requestor is required to supply a justification in their request.
    isRequestorJustificationRequired *bool
    // The OdataType property
    odataType *string
}
// NewApprovalSettings instantiates a new ApprovalSettings and sets the default values.
func NewApprovalSettings()(*ApprovalSettings) {
    m := &ApprovalSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApprovalSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApprovalSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApprovalSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApprovalMode gets the approvalMode property value. One of SingleStage, Serial, Parallel, NoApproval (default). NoApproval is used when isApprovalRequired is false.
// returns a *string when successful
func (m *ApprovalSettings) GetApprovalMode()(*string) {
    return m.approvalMode
}
// GetApprovalStages gets the approvalStages property value. If approval is required, the one or two elements of this collection define each of the stages of approval. An empty array if no approval is required.
// returns a []UnifiedApprovalStageable when successful
func (m *ApprovalSettings) GetApprovalStages()([]UnifiedApprovalStageable) {
    return m.approvalStages
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApprovalSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["approvalMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalMode(val)
        }
        return nil
    }
    res["approvalStages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedApprovalStageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedApprovalStageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedApprovalStageable)
                }
            }
            m.SetApprovalStages(res)
        }
        return nil
    }
    res["isApprovalRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequired(val)
        }
        return nil
    }
    res["isApprovalRequiredForExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequiredForExtension(val)
        }
        return nil
    }
    res["isRequestorJustificationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequestorJustificationRequired(val)
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
// GetIsApprovalRequired gets the isApprovalRequired property value. Indicates whether approval is required for requests in this policy.
// returns a *bool when successful
func (m *ApprovalSettings) GetIsApprovalRequired()(*bool) {
    return m.isApprovalRequired
}
// GetIsApprovalRequiredForExtension gets the isApprovalRequiredForExtension property value. Indicates whether approval is required for a user to extend their assignment.
// returns a *bool when successful
func (m *ApprovalSettings) GetIsApprovalRequiredForExtension()(*bool) {
    return m.isApprovalRequiredForExtension
}
// GetIsRequestorJustificationRequired gets the isRequestorJustificationRequired property value. Indicates whether the requestor is required to supply a justification in their request.
// returns a *bool when successful
func (m *ApprovalSettings) GetIsRequestorJustificationRequired()(*bool) {
    return m.isRequestorJustificationRequired
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ApprovalSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ApprovalSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("approvalMode", m.GetApprovalMode())
        if err != nil {
            return err
        }
    }
    if m.GetApprovalStages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApprovalStages()))
        for i, v := range m.GetApprovalStages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("approvalStages", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequired", m.GetIsApprovalRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequiredForExtension", m.GetIsApprovalRequiredForExtension())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRequestorJustificationRequired", m.GetIsRequestorJustificationRequired())
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
func (m *ApprovalSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApprovalMode sets the approvalMode property value. One of SingleStage, Serial, Parallel, NoApproval (default). NoApproval is used when isApprovalRequired is false.
func (m *ApprovalSettings) SetApprovalMode(value *string)() {
    m.approvalMode = value
}
// SetApprovalStages sets the approvalStages property value. If approval is required, the one or two elements of this collection define each of the stages of approval. An empty array if no approval is required.
func (m *ApprovalSettings) SetApprovalStages(value []UnifiedApprovalStageable)() {
    m.approvalStages = value
}
// SetIsApprovalRequired sets the isApprovalRequired property value. Indicates whether approval is required for requests in this policy.
func (m *ApprovalSettings) SetIsApprovalRequired(value *bool)() {
    m.isApprovalRequired = value
}
// SetIsApprovalRequiredForExtension sets the isApprovalRequiredForExtension property value. Indicates whether approval is required for a user to extend their assignment.
func (m *ApprovalSettings) SetIsApprovalRequiredForExtension(value *bool)() {
    m.isApprovalRequiredForExtension = value
}
// SetIsRequestorJustificationRequired sets the isRequestorJustificationRequired property value. Indicates whether the requestor is required to supply a justification in their request.
func (m *ApprovalSettings) SetIsRequestorJustificationRequired(value *bool)() {
    m.isRequestorJustificationRequired = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ApprovalSettings) SetOdataType(value *string)() {
    m.odataType = value
}
type ApprovalSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovalMode()(*string)
    GetApprovalStages()([]UnifiedApprovalStageable)
    GetIsApprovalRequired()(*bool)
    GetIsApprovalRequiredForExtension()(*bool)
    GetIsRequestorJustificationRequired()(*bool)
    GetOdataType()(*string)
    SetApprovalMode(value *string)()
    SetApprovalStages(value []UnifiedApprovalStageable)()
    SetIsApprovalRequired(value *bool)()
    SetIsApprovalRequiredForExtension(value *bool)()
    SetIsRequestorJustificationRequired(value *bool)()
    SetOdataType(value *string)()
}
