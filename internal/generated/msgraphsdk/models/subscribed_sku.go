package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubscribedSku struct {
    Entity
    // The unique ID of the account this SKU belongs to.
    accountId *string
    // The name of the account this SKU belongs to.
    accountName *string
    // The target class for this SKU. Only SKUs with target class User are assignable. Possible values are: User, Company.
    appliesTo *string
    // Enabled indicates that the prepaidUnits property has at least one unit that is enabled. LockedOut indicates that the customer canceled their subscription. Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut.
    capabilityStatus *string
    // The number of licenses that have been assigned.
    consumedUnits *int32
    // Information about the number and status of prepaid licenses.
    prepaidUnits LicenseUnitsDetailable
    // Information about the service plans that are available with the SKU. Not nullable.
    servicePlans []ServicePlanInfoable
    // The unique identifier (GUID) for the service SKU.
    skuId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The SKU part number; for example: AAD_PREMIUM or RMSBASIC. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus.
    skuPartNumber *string
    // A list of all subscription IDs associated with this SKU.
    subscriptionIds []string
}
// NewSubscribedSku instantiates a new SubscribedSku and sets the default values.
func NewSubscribedSku()(*SubscribedSku) {
    m := &SubscribedSku{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSubscribedSkuFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubscribedSkuFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubscribedSku(), nil
}
// GetAccountId gets the accountId property value. The unique ID of the account this SKU belongs to.
// returns a *string when successful
func (m *SubscribedSku) GetAccountId()(*string) {
    return m.accountId
}
// GetAccountName gets the accountName property value. The name of the account this SKU belongs to.
// returns a *string when successful
func (m *SubscribedSku) GetAccountName()(*string) {
    return m.accountName
}
// GetAppliesTo gets the appliesTo property value. The target class for this SKU. Only SKUs with target class User are assignable. Possible values are: User, Company.
// returns a *string when successful
func (m *SubscribedSku) GetAppliesTo()(*string) {
    return m.appliesTo
}
// GetCapabilityStatus gets the capabilityStatus property value. Enabled indicates that the prepaidUnits property has at least one unit that is enabled. LockedOut indicates that the customer canceled their subscription. Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut.
// returns a *string when successful
func (m *SubscribedSku) GetCapabilityStatus()(*string) {
    return m.capabilityStatus
}
// GetConsumedUnits gets the consumedUnits property value. The number of licenses that have been assigned.
// returns a *int32 when successful
func (m *SubscribedSku) GetConsumedUnits()(*int32) {
    return m.consumedUnits
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubscribedSku) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
        }
        return nil
    }
    res["accountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountName(val)
        }
        return nil
    }
    res["appliesTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesTo(val)
        }
        return nil
    }
    res["capabilityStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilityStatus(val)
        }
        return nil
    }
    res["consumedUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsumedUnits(val)
        }
        return nil
    }
    res["prepaidUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLicenseUnitsDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrepaidUnits(val.(LicenseUnitsDetailable))
        }
        return nil
    }
    res["servicePlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePlanInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePlanInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServicePlanInfoable)
                }
            }
            m.SetServicePlans(res)
        }
        return nil
    }
    res["skuId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuId(val)
        }
        return nil
    }
    res["skuPartNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuPartNumber(val)
        }
        return nil
    }
    res["subscriptionIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSubscriptionIds(res)
        }
        return nil
    }
    return res
}
// GetPrepaidUnits gets the prepaidUnits property value. Information about the number and status of prepaid licenses.
// returns a LicenseUnitsDetailable when successful
func (m *SubscribedSku) GetPrepaidUnits()(LicenseUnitsDetailable) {
    return m.prepaidUnits
}
// GetServicePlans gets the servicePlans property value. Information about the service plans that are available with the SKU. Not nullable.
// returns a []ServicePlanInfoable when successful
func (m *SubscribedSku) GetServicePlans()([]ServicePlanInfoable) {
    return m.servicePlans
}
// GetSkuId gets the skuId property value. The unique identifier (GUID) for the service SKU.
// returns a *UUID when successful
func (m *SubscribedSku) GetSkuId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.skuId
}
// GetSkuPartNumber gets the skuPartNumber property value. The SKU part number; for example: AAD_PREMIUM or RMSBASIC. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus.
// returns a *string when successful
func (m *SubscribedSku) GetSkuPartNumber()(*string) {
    return m.skuPartNumber
}
// GetSubscriptionIds gets the subscriptionIds property value. A list of all subscription IDs associated with this SKU.
// returns a []string when successful
func (m *SubscribedSku) GetSubscriptionIds()([]string) {
    return m.subscriptionIds
}
// Serialize serializes information the current object
func (m *SubscribedSku) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accountId", m.GetAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("accountName", m.GetAccountName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appliesTo", m.GetAppliesTo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("capabilityStatus", m.GetCapabilityStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("consumedUnits", m.GetConsumedUnits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("prepaidUnits", m.GetPrepaidUnits())
        if err != nil {
            return err
        }
    }
    if m.GetServicePlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServicePlans()))
        for i, v := range m.GetServicePlans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("servicePlans", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("skuId", m.GetSkuId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuPartNumber", m.GetSkuPartNumber())
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptionIds() != nil {
        err = writer.WriteCollectionOfStringValues("subscriptionIds", m.GetSubscriptionIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountId sets the accountId property value. The unique ID of the account this SKU belongs to.
func (m *SubscribedSku) SetAccountId(value *string)() {
    m.accountId = value
}
// SetAccountName sets the accountName property value. The name of the account this SKU belongs to.
func (m *SubscribedSku) SetAccountName(value *string)() {
    m.accountName = value
}
// SetAppliesTo sets the appliesTo property value. The target class for this SKU. Only SKUs with target class User are assignable. Possible values are: User, Company.
func (m *SubscribedSku) SetAppliesTo(value *string)() {
    m.appliesTo = value
}
// SetCapabilityStatus sets the capabilityStatus property value. Enabled indicates that the prepaidUnits property has at least one unit that is enabled. LockedOut indicates that the customer canceled their subscription. Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut.
func (m *SubscribedSku) SetCapabilityStatus(value *string)() {
    m.capabilityStatus = value
}
// SetConsumedUnits sets the consumedUnits property value. The number of licenses that have been assigned.
func (m *SubscribedSku) SetConsumedUnits(value *int32)() {
    m.consumedUnits = value
}
// SetPrepaidUnits sets the prepaidUnits property value. Information about the number and status of prepaid licenses.
func (m *SubscribedSku) SetPrepaidUnits(value LicenseUnitsDetailable)() {
    m.prepaidUnits = value
}
// SetServicePlans sets the servicePlans property value. Information about the service plans that are available with the SKU. Not nullable.
func (m *SubscribedSku) SetServicePlans(value []ServicePlanInfoable)() {
    m.servicePlans = value
}
// SetSkuId sets the skuId property value. The unique identifier (GUID) for the service SKU.
func (m *SubscribedSku) SetSkuId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.skuId = value
}
// SetSkuPartNumber sets the skuPartNumber property value. The SKU part number; for example: AAD_PREMIUM or RMSBASIC. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus.
func (m *SubscribedSku) SetSkuPartNumber(value *string)() {
    m.skuPartNumber = value
}
// SetSubscriptionIds sets the subscriptionIds property value. A list of all subscription IDs associated with this SKU.
func (m *SubscribedSku) SetSubscriptionIds(value []string)() {
    m.subscriptionIds = value
}
type SubscribedSkuable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountId()(*string)
    GetAccountName()(*string)
    GetAppliesTo()(*string)
    GetCapabilityStatus()(*string)
    GetConsumedUnits()(*int32)
    GetPrepaidUnits()(LicenseUnitsDetailable)
    GetServicePlans()([]ServicePlanInfoable)
    GetSkuId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetSkuPartNumber()(*string)
    GetSubscriptionIds()([]string)
    SetAccountId(value *string)()
    SetAccountName(value *string)()
    SetAppliesTo(value *string)()
    SetCapabilityStatus(value *string)()
    SetConsumedUnits(value *int32)()
    SetPrepaidUnits(value LicenseUnitsDetailable)()
    SetServicePlans(value []ServicePlanInfoable)()
    SetSkuId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetSkuPartNumber(value *string)()
    SetSubscriptionIds(value []string)()
}
