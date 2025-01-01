package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MultiTenantOrganizationMemberTransitionDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Role of the tenant in the multitenant organization. The possible values are: owner, member, unknownFutureValue.
    desiredRole *MultiTenantOrganizationMemberRole
    // State of the tenant in the multitenant organization currently being processed. The possible values are: pending, active, removed, unknownFutureValue. Read-only.
    desiredState *MultiTenantOrganizationMemberState
    // Details that explain the processing status if any. Read-only.
    details *string
    // The OdataType property
    odataType *string
    // Processing state of the asynchronous job. The possible values are: notStarted, running, succeeded, failed, unknownFutureValue. Read-only.
    status *MultiTenantOrganizationMemberProcessingStatus
}
// NewMultiTenantOrganizationMemberTransitionDetails instantiates a new MultiTenantOrganizationMemberTransitionDetails and sets the default values.
func NewMultiTenantOrganizationMemberTransitionDetails()(*MultiTenantOrganizationMemberTransitionDetails) {
    m := &MultiTenantOrganizationMemberTransitionDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMultiTenantOrganizationMemberTransitionDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMultiTenantOrganizationMemberTransitionDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMultiTenantOrganizationMemberTransitionDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MultiTenantOrganizationMemberTransitionDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDesiredRole gets the desiredRole property value. Role of the tenant in the multitenant organization. The possible values are: owner, member, unknownFutureValue.
// returns a *MultiTenantOrganizationMemberRole when successful
func (m *MultiTenantOrganizationMemberTransitionDetails) GetDesiredRole()(*MultiTenantOrganizationMemberRole) {
    return m.desiredRole
}
// GetDesiredState gets the desiredState property value. State of the tenant in the multitenant organization currently being processed. The possible values are: pending, active, removed, unknownFutureValue. Read-only.
// returns a *MultiTenantOrganizationMemberState when successful
func (m *MultiTenantOrganizationMemberTransitionDetails) GetDesiredState()(*MultiTenantOrganizationMemberState) {
    return m.desiredState
}
// GetDetails gets the details property value. Details that explain the processing status if any. Read-only.
// returns a *string when successful
func (m *MultiTenantOrganizationMemberTransitionDetails) GetDetails()(*string) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MultiTenantOrganizationMemberTransitionDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["desiredRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiTenantOrganizationMemberRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDesiredRole(val.(*MultiTenantOrganizationMemberRole))
        }
        return nil
    }
    res["desiredState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiTenantOrganizationMemberState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDesiredState(val.(*MultiTenantOrganizationMemberState))
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
func (m *MultiTenantOrganizationMemberTransitionDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. Processing state of the asynchronous job. The possible values are: notStarted, running, succeeded, failed, unknownFutureValue. Read-only.
// returns a *MultiTenantOrganizationMemberProcessingStatus when successful
func (m *MultiTenantOrganizationMemberTransitionDetails) GetStatus()(*MultiTenantOrganizationMemberProcessingStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *MultiTenantOrganizationMemberTransitionDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDesiredRole() != nil {
        cast := (*m.GetDesiredRole()).String()
        err := writer.WriteStringValue("desiredRole", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDesiredState() != nil {
        cast := (*m.GetDesiredState()).String()
        err := writer.WriteStringValue("desiredState", &cast)
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
func (m *MultiTenantOrganizationMemberTransitionDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDesiredRole sets the desiredRole property value. Role of the tenant in the multitenant organization. The possible values are: owner, member, unknownFutureValue.
func (m *MultiTenantOrganizationMemberTransitionDetails) SetDesiredRole(value *MultiTenantOrganizationMemberRole)() {
    m.desiredRole = value
}
// SetDesiredState sets the desiredState property value. State of the tenant in the multitenant organization currently being processed. The possible values are: pending, active, removed, unknownFutureValue. Read-only.
func (m *MultiTenantOrganizationMemberTransitionDetails) SetDesiredState(value *MultiTenantOrganizationMemberState)() {
    m.desiredState = value
}
// SetDetails sets the details property value. Details that explain the processing status if any. Read-only.
func (m *MultiTenantOrganizationMemberTransitionDetails) SetDetails(value *string)() {
    m.details = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MultiTenantOrganizationMemberTransitionDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. Processing state of the asynchronous job. The possible values are: notStarted, running, succeeded, failed, unknownFutureValue. Read-only.
func (m *MultiTenantOrganizationMemberTransitionDetails) SetStatus(value *MultiTenantOrganizationMemberProcessingStatus)() {
    m.status = value
}
type MultiTenantOrganizationMemberTransitionDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDesiredRole()(*MultiTenantOrganizationMemberRole)
    GetDesiredState()(*MultiTenantOrganizationMemberState)
    GetDetails()(*string)
    GetOdataType()(*string)
    GetStatus()(*MultiTenantOrganizationMemberProcessingStatus)
    SetDesiredRole(value *MultiTenantOrganizationMemberRole)()
    SetDesiredState(value *MultiTenantOrganizationMemberState)()
    SetDetails(value *string)()
    SetOdataType(value *string)()
    SetStatus(value *MultiTenantOrganizationMemberProcessingStatus)()
}
