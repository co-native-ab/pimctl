package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingBusiness represents a Microsoft Bookings Business.
type BookingBusiness struct {
    Entity
    // The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
    address PhysicalAddressable
    // All the appointments of this business. Read-only. Nullable.
    appointments []BookingAppointmentable
    // Settings for the published booking page.
    bookingPageSettings BookingPageSettingsable
    // The hours of operation for the business.
    businessHours []BookingWorkHoursable
    // The type of business.
    businessType *string
    // The set of appointments of this business in a specified date range. Read-only. Nullable.
    calendarView []BookingAppointmentable
    // The date, time, and time zone when the booking business was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // All the customers of this business. Read-only. Nullable.
    customers []BookingCustomerBaseable
    // All the custom questions of this business. Read-only. Nullable.
    customQuestions []BookingCustomQuestionable
    // The code for the currency that the business operates in on Microsoft Bookings.
    defaultCurrencyIso *string
    // The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
    displayName *string
    // The email address for the business.
    email *string
    // The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
    isPublished *bool
    // The language of the self-service booking page.
    languageTag *string
    // The date, time, and time zone when the booking business was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
    phone *string
    // The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
    publicUrl *string
    // Specifies how bookings can be created for this business.
    schedulingPolicy BookingSchedulingPolicyable
    // All the services offered by this business. Read-only. Nullable.
    services []BookingServiceable
    // All the staff members that provide services in this business. Read-only. Nullable.
    staffMembers []BookingStaffMemberBaseable
    // The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
    webSiteUrl *string
}
// NewBookingBusiness instantiates a new BookingBusiness and sets the default values.
func NewBookingBusiness()(*BookingBusiness) {
    m := &BookingBusiness{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBookingBusinessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBookingBusinessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingBusiness(), nil
}
// GetAddress gets the address property value. The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
// returns a PhysicalAddressable when successful
func (m *BookingBusiness) GetAddress()(PhysicalAddressable) {
    return m.address
}
// GetAppointments gets the appointments property value. All the appointments of this business. Read-only. Nullable.
// returns a []BookingAppointmentable when successful
func (m *BookingBusiness) GetAppointments()([]BookingAppointmentable) {
    return m.appointments
}
// GetBookingPageSettings gets the bookingPageSettings property value. Settings for the published booking page.
// returns a BookingPageSettingsable when successful
func (m *BookingBusiness) GetBookingPageSettings()(BookingPageSettingsable) {
    return m.bookingPageSettings
}
// GetBusinessHours gets the businessHours property value. The hours of operation for the business.
// returns a []BookingWorkHoursable when successful
func (m *BookingBusiness) GetBusinessHours()([]BookingWorkHoursable) {
    return m.businessHours
}
// GetBusinessType gets the businessType property value. The type of business.
// returns a *string when successful
func (m *BookingBusiness) GetBusinessType()(*string) {
    return m.businessType
}
// GetCalendarView gets the calendarView property value. The set of appointments of this business in a specified date range. Read-only. Nullable.
// returns a []BookingAppointmentable when successful
func (m *BookingBusiness) GetCalendarView()([]BookingAppointmentable) {
    return m.calendarView
}
// GetCreatedDateTime gets the createdDateTime property value. The date, time, and time zone when the booking business was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *BookingBusiness) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomers gets the customers property value. All the customers of this business. Read-only. Nullable.
// returns a []BookingCustomerBaseable when successful
func (m *BookingBusiness) GetCustomers()([]BookingCustomerBaseable) {
    return m.customers
}
// GetCustomQuestions gets the customQuestions property value. All the custom questions of this business. Read-only. Nullable.
// returns a []BookingCustomQuestionable when successful
func (m *BookingBusiness) GetCustomQuestions()([]BookingCustomQuestionable) {
    return m.customQuestions
}
// GetDefaultCurrencyIso gets the defaultCurrencyIso property value. The code for the currency that the business operates in on Microsoft Bookings.
// returns a *string when successful
func (m *BookingBusiness) GetDefaultCurrencyIso()(*string) {
    return m.defaultCurrencyIso
}
// GetDisplayName gets the displayName property value. The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
// returns a *string when successful
func (m *BookingBusiness) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The email address for the business.
// returns a *string when successful
func (m *BookingBusiness) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BookingBusiness) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["appointments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingAppointmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingAppointmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingAppointmentable)
                }
            }
            m.SetAppointments(res)
        }
        return nil
    }
    res["bookingPageSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBookingPageSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBookingPageSettings(val.(BookingPageSettingsable))
        }
        return nil
    }
    res["businessHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingWorkHoursFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingWorkHoursable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingWorkHoursable)
                }
            }
            m.SetBusinessHours(res)
        }
        return nil
    }
    res["businessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessType(val)
        }
        return nil
    }
    res["calendarView"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingAppointmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingAppointmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingAppointmentable)
                }
            }
            m.SetCalendarView(res)
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
    res["customers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingCustomerBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCustomerBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingCustomerBaseable)
                }
            }
            m.SetCustomers(res)
        }
        return nil
    }
    res["customQuestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingCustomQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCustomQuestionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingCustomQuestionable)
                }
            }
            m.SetCustomQuestions(res)
        }
        return nil
    }
    res["defaultCurrencyIso"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultCurrencyIso(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["isPublished"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPublished(val)
        }
        return nil
    }
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    res["publicUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicUrl(val)
        }
        return nil
    }
    res["schedulingPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBookingSchedulingPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingPolicy(val.(BookingSchedulingPolicyable))
        }
        return nil
    }
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingServiceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingServiceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingServiceable)
                }
            }
            m.SetServices(res)
        }
        return nil
    }
    res["staffMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingStaffMemberBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingStaffMemberBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingStaffMemberBaseable)
                }
            }
            m.SetStaffMembers(res)
        }
        return nil
    }
    res["webSiteUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebSiteUrl(val)
        }
        return nil
    }
    return res
}
// GetIsPublished gets the isPublished property value. The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
// returns a *bool when successful
func (m *BookingBusiness) GetIsPublished()(*bool) {
    return m.isPublished
}
// GetLanguageTag gets the languageTag property value. The language of the self-service booking page.
// returns a *string when successful
func (m *BookingBusiness) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The date, time, and time zone when the booking business was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *BookingBusiness) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetPhone gets the phone property value. The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
// returns a *string when successful
func (m *BookingBusiness) GetPhone()(*string) {
    return m.phone
}
// GetPublicUrl gets the publicUrl property value. The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
// returns a *string when successful
func (m *BookingBusiness) GetPublicUrl()(*string) {
    return m.publicUrl
}
// GetSchedulingPolicy gets the schedulingPolicy property value. Specifies how bookings can be created for this business.
// returns a BookingSchedulingPolicyable when successful
func (m *BookingBusiness) GetSchedulingPolicy()(BookingSchedulingPolicyable) {
    return m.schedulingPolicy
}
// GetServices gets the services property value. All the services offered by this business. Read-only. Nullable.
// returns a []BookingServiceable when successful
func (m *BookingBusiness) GetServices()([]BookingServiceable) {
    return m.services
}
// GetStaffMembers gets the staffMembers property value. All the staff members that provide services in this business. Read-only. Nullable.
// returns a []BookingStaffMemberBaseable when successful
func (m *BookingBusiness) GetStaffMembers()([]BookingStaffMemberBaseable) {
    return m.staffMembers
}
// GetWebSiteUrl gets the webSiteUrl property value. The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
// returns a *string when successful
func (m *BookingBusiness) GetWebSiteUrl()(*string) {
    return m.webSiteUrl
}
// Serialize serializes information the current object
func (m *BookingBusiness) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    if m.GetAppointments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppointments()))
        for i, v := range m.GetAppointments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appointments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bookingPageSettings", m.GetBookingPageSettings())
        if err != nil {
            return err
        }
    }
    if m.GetBusinessHours() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBusinessHours()))
        for i, v := range m.GetBusinessHours() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("businessHours", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("businessType", m.GetBusinessType())
        if err != nil {
            return err
        }
    }
    if m.GetCalendarView() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCalendarView()))
        for i, v := range m.GetCalendarView() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
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
    if m.GetCustomers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomers()))
        for i, v := range m.GetCustomers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomQuestions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomQuestions()))
        for i, v := range m.GetCustomQuestions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customQuestions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultCurrencyIso", m.GetDefaultCurrencyIso())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedulingPolicy", m.GetSchedulingPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServices()))
        for i, v := range m.GetServices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("services", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStaffMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStaffMembers()))
        for i, v := range m.GetStaffMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("staffMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webSiteUrl", m.GetWebSiteUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
func (m *BookingBusiness) SetAddress(value PhysicalAddressable)() {
    m.address = value
}
// SetAppointments sets the appointments property value. All the appointments of this business. Read-only. Nullable.
func (m *BookingBusiness) SetAppointments(value []BookingAppointmentable)() {
    m.appointments = value
}
// SetBookingPageSettings sets the bookingPageSettings property value. Settings for the published booking page.
func (m *BookingBusiness) SetBookingPageSettings(value BookingPageSettingsable)() {
    m.bookingPageSettings = value
}
// SetBusinessHours sets the businessHours property value. The hours of operation for the business.
func (m *BookingBusiness) SetBusinessHours(value []BookingWorkHoursable)() {
    m.businessHours = value
}
// SetBusinessType sets the businessType property value. The type of business.
func (m *BookingBusiness) SetBusinessType(value *string)() {
    m.businessType = value
}
// SetCalendarView sets the calendarView property value. The set of appointments of this business in a specified date range. Read-only. Nullable.
func (m *BookingBusiness) SetCalendarView(value []BookingAppointmentable)() {
    m.calendarView = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date, time, and time zone when the booking business was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BookingBusiness) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomers sets the customers property value. All the customers of this business. Read-only. Nullable.
func (m *BookingBusiness) SetCustomers(value []BookingCustomerBaseable)() {
    m.customers = value
}
// SetCustomQuestions sets the customQuestions property value. All the custom questions of this business. Read-only. Nullable.
func (m *BookingBusiness) SetCustomQuestions(value []BookingCustomQuestionable)() {
    m.customQuestions = value
}
// SetDefaultCurrencyIso sets the defaultCurrencyIso property value. The code for the currency that the business operates in on Microsoft Bookings.
func (m *BookingBusiness) SetDefaultCurrencyIso(value *string)() {
    m.defaultCurrencyIso = value
}
// SetDisplayName sets the displayName property value. The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
func (m *BookingBusiness) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The email address for the business.
func (m *BookingBusiness) SetEmail(value *string)() {
    m.email = value
}
// SetIsPublished sets the isPublished property value. The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
func (m *BookingBusiness) SetIsPublished(value *bool)() {
    m.isPublished = value
}
// SetLanguageTag sets the languageTag property value. The language of the self-service booking page.
func (m *BookingBusiness) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The date, time, and time zone when the booking business was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BookingBusiness) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetPhone sets the phone property value. The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
func (m *BookingBusiness) SetPhone(value *string)() {
    m.phone = value
}
// SetPublicUrl sets the publicUrl property value. The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
func (m *BookingBusiness) SetPublicUrl(value *string)() {
    m.publicUrl = value
}
// SetSchedulingPolicy sets the schedulingPolicy property value. Specifies how bookings can be created for this business.
func (m *BookingBusiness) SetSchedulingPolicy(value BookingSchedulingPolicyable)() {
    m.schedulingPolicy = value
}
// SetServices sets the services property value. All the services offered by this business. Read-only. Nullable.
func (m *BookingBusiness) SetServices(value []BookingServiceable)() {
    m.services = value
}
// SetStaffMembers sets the staffMembers property value. All the staff members that provide services in this business. Read-only. Nullable.
func (m *BookingBusiness) SetStaffMembers(value []BookingStaffMemberBaseable)() {
    m.staffMembers = value
}
// SetWebSiteUrl sets the webSiteUrl property value. The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
func (m *BookingBusiness) SetWebSiteUrl(value *string)() {
    m.webSiteUrl = value
}
type BookingBusinessable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(PhysicalAddressable)
    GetAppointments()([]BookingAppointmentable)
    GetBookingPageSettings()(BookingPageSettingsable)
    GetBusinessHours()([]BookingWorkHoursable)
    GetBusinessType()(*string)
    GetCalendarView()([]BookingAppointmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomers()([]BookingCustomerBaseable)
    GetCustomQuestions()([]BookingCustomQuestionable)
    GetDefaultCurrencyIso()(*string)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetIsPublished()(*bool)
    GetLanguageTag()(*string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPhone()(*string)
    GetPublicUrl()(*string)
    GetSchedulingPolicy()(BookingSchedulingPolicyable)
    GetServices()([]BookingServiceable)
    GetStaffMembers()([]BookingStaffMemberBaseable)
    GetWebSiteUrl()(*string)
    SetAddress(value PhysicalAddressable)()
    SetAppointments(value []BookingAppointmentable)()
    SetBookingPageSettings(value BookingPageSettingsable)()
    SetBusinessHours(value []BookingWorkHoursable)()
    SetBusinessType(value *string)()
    SetCalendarView(value []BookingAppointmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomers(value []BookingCustomerBaseable)()
    SetCustomQuestions(value []BookingCustomQuestionable)()
    SetDefaultCurrencyIso(value *string)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetIsPublished(value *bool)()
    SetLanguageTag(value *string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPhone(value *string)()
    SetPublicUrl(value *string)()
    SetSchedulingPolicy(value BookingSchedulingPolicyable)()
    SetServices(value []BookingServiceable)()
    SetStaffMembers(value []BookingStaffMemberBaseable)()
    SetWebSiteUrl(value *string)()
}
