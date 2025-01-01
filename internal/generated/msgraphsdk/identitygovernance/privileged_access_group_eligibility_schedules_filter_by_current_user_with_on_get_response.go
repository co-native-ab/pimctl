package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BaseCollectionPaginationCountResponse
    // The value property
    value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable
}
// NewPrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse instantiates a new PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse and sets the default values.
func NewPrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse()(*PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse) {
    m := &PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse{
        BaseCollectionPaginationCountResponse: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []PrivilegedAccessGroupEligibilityScheduleable when successful
func (m *PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse) GetValue()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable) {
    return m.value
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponse) SetValue(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable)() {
    m.value = value
}
type PrivilegedAccessGroupEligibilitySchedulesFilterByCurrentUserWithOnGetResponseable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable)
    SetValue(value []ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.PrivilegedAccessGroupEligibilityScheduleable)()
}
