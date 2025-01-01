package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BaseCollectionPaginationCountResponse
    // The value property
    value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestable
}
// NewPrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse instantiates a new PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse()(*PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse) {
    m := &PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse{
        BaseCollectionPaginationCountResponse: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreatePrivilegedAccessGroupAssignmentScheduleRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []PrivilegedAccessGroupAssignmentScheduleRequestable when successful
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse) GetValue()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestable) {
    return m.value
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse) SetValue(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestable)() {
    m.value = value
}
type PrivilegedAccessGroupAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestable)
    SetValue(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupAssignmentScheduleRequestable)()
}
