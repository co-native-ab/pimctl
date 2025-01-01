package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MultiTenantOrganizationJoinRequestTransitionDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // State of the tenant in the multitenant organization currently being processed. The possible values are: pending, active, removed, unknownFutureValue. Read-only.
    desiredMemberState *MultiTenantOrganizationMemberState
    // Details that explain the processing status if any. Read-only.
    details *string
    // The OdataType property
    odataType *string
    // Processing state of the asynchronous job. The possible values are: notStarted, running, succeeded, failed, unknownFutureValue. Read-only.
    status *MultiTenantOrganizationMemberProcessingStatus
}
// NewMultiTenantOrganizationJoinRequestTransitionDetails instantiates a new MultiTenantOrganizationJoinRequestTransitionDetails and sets the default values.
func NewMultiTenantOrganizationJoinRequestTransitionDetails()(*MultiTenantOrganizationJoinRequestTransitionDetails) {
    m := &MultiTenantOrganizationJoinRequestTransitionDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMultiTenantOrganizationJoinRequestTransitionDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMultiTenantOrganizationJoinRequestTransitionDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMultiTenantOrganizationJoinRequestTransitionDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDesiredMemberState gets the desiredMemberState property value. State of the tenant in the multitenant organization currently being processed. The possible values are: pending, active, removed, unknownFutureValue. Read-only.
// returns a *MultiTenantOrganizationMemberState when successful
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetDesiredMemberState()(*MultiTenantOrganizationMemberState) {
    return m.desiredMemberState
}
// GetDetails gets the details property value. Details that explain the processing status if any. Read-only.
// returns a *string when successful
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetDetails()(*string) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["desiredMemberState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiTenantOrganizationMemberState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDesiredMemberState(val.(*MultiTenantOrganizationMemberState))
        }
        return nil
    }
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiTenantOrganizationMemberProcessingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*MultiTenantOrganizationMemberProcessingStatus))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. Processing state of the asynchronous job. The possible values are: notStarted, running, succeeded, failed, unknownFutureValue. Read-only.
// returns a *MultiTenantOrganizationMemberProcessingStatus when successful
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetStatus()(*MultiTenantOrganizationMemberProcessingStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDesiredMemberState() != nil {
        cast := (*m.GetDesiredMemberState()).String()
        err := writer.WriteStringValue("desiredMemberState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("details", m.GetDetails())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDesiredMemberState sets the desiredMemberState property value. State of the tenant in the multitenant organization currently being processed. The possible values are: pending, active, removed, unknownFutureValue. Read-only.
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) SetDesiredMemberState(value *MultiTenantOrganizationMemberState)() {
    m.desiredMemberState = value
}
// SetDetails sets the details property value. Details that explain the processing status if any. Read-only.
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) SetDetails(value *string)() {
    m.details = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. Processing state of the asynchronous job. The possible values are: notStarted, running, succeeded, failed, unknownFutureValue. Read-only.
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) SetStatus(value *MultiTenantOrganizationMemberProcessingStatus)() {
    m.status = value
}
type MultiTenantOrganizationJoinRequestTransitionDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDesiredMemberState()(*MultiTenantOrganizationMemberState)
    GetDetails()(*string)
    GetOdataType()(*string)
    GetStatus()(*MultiTenantOrganizationMemberProcessingStatus)
    SetDesiredMemberState(value *MultiTenantOrganizationMemberState)()
    SetDetails(value *string)()
    SetOdataType(value *string)()
    SetStatus(value *MultiTenantOrganizationMemberProcessingStatus)()
}
