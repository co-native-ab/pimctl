package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BaseCollectionPaginationCountResponse
    // The value property
    value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleRequestable
}
// NewDirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse instantiates a new DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse and sets the default values.
func NewDirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse()(*DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse) {
    m := &DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse{
        BaseCollectionPaginationCountResponse: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleRequestable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []UnifiedRoleAssignmentScheduleRequestable when successful
func (m *DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse) GetValue()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleRequestable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponse) SetValue(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleRequestable)() {
    m.value = value
}
type DirectoryRoleAssignmentScheduleRequestsFilterByCurrentUserWithOnGetResponseable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleRequestable)
    SetValue(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.UnifiedRoleAssignmentScheduleRequestable)()
}
