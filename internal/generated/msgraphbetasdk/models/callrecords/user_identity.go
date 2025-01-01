package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc "github.com/co-native-ab/pimctl/internal/generated/msgraphbetasdk/models"
)

type UserIdentity struct {
    ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Identity
    // The userPrincipalName property
    userPrincipalName *string
}
// NewUserIdentity instantiates a new UserIdentity and sets the default values.
func NewUserIdentity()(*UserIdentity) {
    m := &UserIdentity{
        Identity: *ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.callRecords.userIdentity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUserIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
// returns a *string when successful
func (m *UserIdentity) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *UserIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *UserIdentity) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
type UserIdentityable interface {
    ic0b244548f4d248264a1f750ce59f6e064a76c1a39ca69f620ae700132501ecc.Identityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUserPrincipalName()(*string)
    SetUserPrincipalName(value *string)()
}
