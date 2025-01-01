package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompanySubscription struct {
    Entity
    // The ID of this subscription in the commerce system. Alternate key.
    commerceSubscriptionId *string
    // The date and time when this subscription was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Whether the subscription is a free trial or purchased.
    isTrial *bool
    // The date and time when the subscription will move to the next state (as defined by the status property) if not renewed by the tenant. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    nextLifecycleDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The object ID of the account admin.
    ownerId *string
    // The unique identifier for the Microsoft partner tenant that created the subscription on a customer tenant.
    ownerTenantId *string
    // Indicates the entity that ownerId belongs to, for example, 'User'.
    ownerType *string
    // The provisioning status of each service included in this subscription.
    serviceStatus []ServicePlanInfoable
    // The object ID of the SKU associated with this subscription.
    skuId *string
    // The SKU associated with this subscription.
    skuPartNumber *string
    // The status of this subscription. Possible values are: Enabled, Deleted, Suspended, Warning, LockedOut.
    status *string
    // The number of licenses included in this subscription.
    totalLicenses *int32
}
// NewCompanySubscription instantiates a new CompanySubscription and sets the default values.
func NewCompanySubscription()(*CompanySubscription) {
    m := &CompanySubscription{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCompanySubscriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompanySubscriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompanySubscription(), nil
}
// GetCommerceSubscriptionId gets the commerceSubscriptionId property value. The ID of this subscription in the commerce system. Alternate key.
// returns a *string when successful
func (m *CompanySubscription) GetCommerceSubscriptionId()(*string) {
    return m.commerceSubscriptionId
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when this subscription was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *CompanySubscription) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompanySubscription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["commerceSubscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommerceSubscriptionId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["isTrial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTrial(val)
        }
        return nil
    }
    res["nextLifecycleDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextLifecycleDateTime(val)
        }
        return nil
    }
    res["ownerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerId(val)
        }
        return nil
    }
    res["ownerTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerTenantId(val)
        }
        return nil
    }
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val)
        }
        return nil
    }
    res["serviceStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetServiceStatus(res)
        }
        return nil
    }
    res["skuId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["totalLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLicenses(val)
        }
        return nil
    }
    return res
}
// GetIsTrial gets the isTrial property value. Whether the subscription is a free trial or purchased.
// returns a *bool when successful
func (m *CompanySubscription) GetIsTrial()(*bool) {
    return m.isTrial
}
// GetNextLifecycleDateTime gets the nextLifecycleDateTime property value. The date and time when the subscription will move to the next state (as defined by the status property) if not renewed by the tenant. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *CompanySubscription) GetNextLifecycleDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.nextLifecycleDateTime
}
// GetOwnerId gets the ownerId property value. The object ID of the account admin.
// returns a *string when successful
func (m *CompanySubscription) GetOwnerId()(*string) {
    return m.ownerId
}
// GetOwnerTenantId gets the ownerTenantId property value. The unique identifier for the Microsoft partner tenant that created the subscription on a customer tenant.
// returns a *string when successful
func (m *CompanySubscription) GetOwnerTenantId()(*string) {
    return m.ownerTenantId
}
// GetOwnerType gets the ownerType property value. Indicates the entity that ownerId belongs to, for example, 'User'.
// returns a *string when successful
func (m *CompanySubscription) GetOwnerType()(*string) {
    return m.ownerType
}
// GetServiceStatus gets the serviceStatus property value. The provisioning status of each service included in this subscription.
// returns a []ServicePlanInfoable when successful
func (m *CompanySubscription) GetServiceStatus()([]ServicePlanInfoable) {
    return m.serviceStatus
}
// GetSkuId gets the skuId property value. The object ID of the SKU associated with this subscription.
// returns a *string when successful
func (m *CompanySubscription) GetSkuId()(*string) {
    return m.skuId
}
// GetSkuPartNumber gets the skuPartNumber property value. The SKU associated with this subscription.
// returns a *string when successful
func (m *CompanySubscription) GetSkuPartNumber()(*string) {
    return m.skuPartNumber
}
// GetStatus gets the status property value. The status of this subscription. Possible values are: Enabled, Deleted, Suspended, Warning, LockedOut.
// returns a *string when successful
func (m *CompanySubscription) GetStatus()(*string) {
    return m.status
}
// GetTotalLicenses gets the totalLicenses property value. The number of licenses included in this subscription.
// returns a *int32 when successful
func (m *CompanySubscription) GetTotalLicenses()(*int32) {
    return m.totalLicenses
}
// Serialize serializes information the current object
func (m *CompanySubscription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("commerceSubscriptionId", m.GetCommerceSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTrial", m.GetIsTrial())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("nextLifecycleDateTime", m.GetNextLifecycleDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerId", m.GetOwnerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerTenantId", m.GetOwnerTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerType", m.GetOwnerType())
        if err != nil {
            return err
        }
    }
    if m.GetServiceStatus() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServiceStatus()))
        for i, v := range m.GetServiceStatus() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("serviceStatus", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuId", m.GetSkuId())
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
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalLicenses", m.GetTotalLicenses())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCommerceSubscriptionId sets the commerceSubscriptionId property value. The ID of this subscription in the commerce system. Alternate key.
func (m *CompanySubscription) SetCommerceSubscriptionId(value *string)() {
    m.commerceSubscriptionId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when this subscription was created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CompanySubscription) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetIsTrial sets the isTrial property value. Whether the subscription is a free trial or purchased.
func (m *CompanySubscription) SetIsTrial(value *bool)() {
    m.isTrial = value
}
// SetNextLifecycleDateTime sets the nextLifecycleDateTime property value. The date and time when the subscription will move to the next state (as defined by the status property) if not renewed by the tenant. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CompanySubscription) SetNextLifecycleDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.nextLifecycleDateTime = value
}
// SetOwnerId sets the ownerId property value. The object ID of the account admin.
func (m *CompanySubscription) SetOwnerId(value *string)() {
    m.ownerId = value
}
// SetOwnerTenantId sets the ownerTenantId property value. The unique identifier for the Microsoft partner tenant that created the subscription on a customer tenant.
func (m *CompanySubscription) SetOwnerTenantId(value *string)() {
    m.ownerTenantId = value
}
// SetOwnerType sets the ownerType property value. Indicates the entity that ownerId belongs to, for example, 'User'.
func (m *CompanySubscription) SetOwnerType(value *string)() {
    m.ownerType = value
}
// SetServiceStatus sets the serviceStatus property value. The provisioning status of each service included in this subscription.
func (m *CompanySubscription) SetServiceStatus(value []ServicePlanInfoable)() {
    m.serviceStatus = value
}
// SetSkuId sets the skuId property value. The object ID of the SKU associated with this subscription.
func (m *CompanySubscription) SetSkuId(value *string)() {
    m.skuId = value
}
// SetSkuPartNumber sets the skuPartNumber property value. The SKU associated with this subscription.
func (m *CompanySubscription) SetSkuPartNumber(value *string)() {
    m.skuPartNumber = value
}
// SetStatus sets the status property value. The status of this subscription. Possible values are: Enabled, Deleted, Suspended, Warning, LockedOut.
func (m *CompanySubscription) SetStatus(value *string)() {
    m.status = value
}
// SetTotalLicenses sets the totalLicenses property value. The number of licenses included in this subscription.
func (m *CompanySubscription) SetTotalLicenses(value *int32)() {
    m.totalLicenses = value
}
type CompanySubscriptionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommerceSubscriptionId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsTrial()(*bool)
    GetNextLifecycleDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOwnerId()(*string)
    GetOwnerTenantId()(*string)
    GetOwnerType()(*string)
    GetServiceStatus()([]ServicePlanInfoable)
    GetSkuId()(*string)
    GetSkuPartNumber()(*string)
    GetStatus()(*string)
    GetTotalLicenses()(*int32)
    SetCommerceSubscriptionId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsTrial(value *bool)()
    SetNextLifecycleDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOwnerId(value *string)()
    SetOwnerTenantId(value *string)()
    SetOwnerType(value *string)()
    SetServiceStatus(value []ServicePlanInfoable)()
    SetSkuId(value *string)()
    SetSkuPartNumber(value *string)()
    SetStatus(value *string)()
    SetTotalLicenses(value *int32)()
}
