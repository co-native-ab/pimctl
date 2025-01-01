package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc "github.com/co-native-ab/pimctl/internal/generated/msgraphbetasdk/models"
)

type DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse struct {
    ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.BaseCollectionPaginationCountResponse
    // The value property
    value []ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable
}
// NewDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse instantiates a new DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse and sets the default values.
func NewDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse()(*DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse) {
    m := &DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse{
        BaseCollectionPaginationCountResponse: *ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.CreateApprovalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []Approvalable when successful
func (m *DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse) GetValue()([]ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponse) SetValue(value []ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable)() {
    m.value = value
}
type DirectoryRoleAssignmentApprovalsFilterByCurrentUserWithOnGetResponseable interface {
    ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable)
    SetValue(value []ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Approvalable)()
}
