package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingService represents a particular service offered by a booking business.
type BookingService struct {
    Entity
    // Additional information that is sent to the customer when an appointment is confirmed.
    additionalInformation *string
    // The date, time, and time zone when the service was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Contains the set of custom questions associated with a particular service.
    customQuestions []BookingQuestionAssignmentable
    // The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
    defaultDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The default physical location for the service.
    defaultLocation Locationable
    // The default monetary price for the service.
    defaultPrice *float64
    // Represents the type of pricing of a booking service.
    defaultPriceType *BookingPriceType
    // The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
    defaultReminders []BookingReminderable
    // A text description for the service.
    description *string
    // A service name.
    displayName *string
    // Indicates if an anonymousJoinWebUrl(webrtcUrl) is generated for the appointment booked for this service. The default value is false.
    isAnonymousJoinEnabled *bool
    // Indicates that the customer can manage bookings created by the staff. The default value is false.
    isCustomerAllowedToManageBooking *bool
    // True indicates that this service isn't available to customers for booking.
    isHiddenFromCustomers *bool
    // Indicates that the appointments for the service are held online. The default value is false.
    isLocationOnline *bool
    // The language of the self-service booking page.
    languageTag *string
    // The date, time, and time zone when the service was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The maximum number of customers allowed in a service. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
    maximumAttendeesCount *int32
    // Additional information about this service.
    notes *string
    // The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
    postBuffer *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The time to buffer before an appointment for this service can start.
    preBuffer *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The set of policies that determine how appointments for this type of service should be created and managed.
    schedulingPolicy BookingSchedulingPolicyable
    // True indicates SMS notifications can be sent to the customers for the appointment of the service. Default value is false.
    smsNotificationsEnabled *bool
    // Represents those staff members who provide this service.
    staffMemberIds []string
    // The URL a customer uses to access the service.
    webUrl *string
}
// NewBookingService instantiates a new BookingService and sets the default values.
func NewBookingService()(*BookingService) {
    m := &BookingService{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBookingServiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBookingServiceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingService(), nil
}
// GetAdditionalInformation gets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
// returns a *string when successful
func (m *BookingService) GetAdditionalInformation()(*string) {
    return m.additionalInformation
}
// GetCreatedDateTime gets the createdDateTime property value. The date, time, and time zone when the service was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *BookingService) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomQuestions gets the customQuestions property value. Contains the set of custom questions associated with a particular service.
// returns a []BookingQuestionAssignmentable when successful
func (m *BookingService) GetCustomQuestions()([]BookingQuestionAssignmentable) {
    return m.customQuestions
}
// GetDefaultDuration gets the defaultDuration property value. The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
// returns a *ISODuration when successful
func (m *BookingService) GetDefaultDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.defaultDuration
}
// GetDefaultLocation gets the defaultLocation property value. The default physical location for the service.
// returns a Locationable when successful
func (m *BookingService) GetDefaultLocation()(Locationable) {
    return m.defaultLocation
}
// GetDefaultPrice gets the defaultPrice property value. The default monetary price for the service.
// returns a *float64 when successful
func (m *BookingService) GetDefaultPrice()(*float64) {
    return m.defaultPrice
}
// GetDefaultPriceType gets the defaultPriceType property value. Represents the type of pricing of a booking service.
// returns a *BookingPriceType when successful
func (m *BookingService) GetDefaultPriceType()(*BookingPriceType) {
    return m.defaultPriceType
}
// GetDefaultReminders gets the defaultReminders property value. The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
// returns a []BookingReminderable when successful
func (m *BookingService) GetDefaultReminders()([]BookingReminderable) {
    return m.defaultReminders
}
// GetDescription gets the description property value. A text description for the service.
// returns a *string when successful
func (m *BookingService) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. A service name.
// returns a *string when successful
func (m *BookingService) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BookingService) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalInformation(val)
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
    res["customQuestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingQuestionAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingQuestionAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingQuestionAssignmentable)
                }
            }
            m.SetCustomQuestions(res)
        }
        return nil
    }
    res["defaultDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDuration(val)
        }
        return nil
    }
    res["defaultLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLocation(val.(Locationable))
        }
        return nil
    }
    res["defaultPrice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultPrice(val)
        }
        return nil
    }
    res["defaultPriceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingPriceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultPriceType(val.(*BookingPriceType))
        }
        return nil
    }
    res["defaultReminders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingReminderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingReminderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingReminderable)
                }
            }
            m.SetDefaultReminders(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["isAnonymousJoinEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAnonymousJoinEnabled(val)
        }
        return nil
    }
    res["isCustomerAllowedToManageBooking"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomerAllowedToManageBooking(val)
        }
        return nil
    }
    res["isHiddenFromCustomers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHiddenFromCustomers(val)
        }
        return nil
    }
    res["isLocationOnline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLocationOnline(val)
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
    res["maximumAttendeesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAttendeesCount(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["postBuffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostBuffer(val)
        }
        return nil
    }
    res["preBuffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreBuffer(val)
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
    res["smsNotificationsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsNotificationsEnabled(val)
        }
        return nil
    }
    res["staffMemberIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetStaffMemberIds(res)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetIsAnonymousJoinEnabled gets the isAnonymousJoinEnabled property value. Indicates if an anonymousJoinWebUrl(webrtcUrl) is generated for the appointment booked for this service. The default value is false.
// returns a *bool when successful
func (m *BookingService) GetIsAnonymousJoinEnabled()(*bool) {
    return m.isAnonymousJoinEnabled
}
// GetIsCustomerAllowedToManageBooking gets the isCustomerAllowedToManageBooking property value. Indicates that the customer can manage bookings created by the staff. The default value is false.
// returns a *bool when successful
func (m *BookingService) GetIsCustomerAllowedToManageBooking()(*bool) {
    return m.isCustomerAllowedToManageBooking
}
// GetIsHiddenFromCustomers gets the isHiddenFromCustomers property value. True indicates that this service isn't available to customers for booking.
// returns a *bool when successful
func (m *BookingService) GetIsHiddenFromCustomers()(*bool) {
    return m.isHiddenFromCustomers
}
// GetIsLocationOnline gets the isLocationOnline property value. Indicates that the appointments for the service are held online. The default value is false.
// returns a *bool when successful
func (m *BookingService) GetIsLocationOnline()(*bool) {
    return m.isLocationOnline
}
// GetLanguageTag gets the languageTag property value. The language of the self-service booking page.
// returns a *string when successful
func (m *BookingService) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The date, time, and time zone when the service was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *BookingService) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetMaximumAttendeesCount gets the maximumAttendeesCount property value. The maximum number of customers allowed in a service. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
// returns a *int32 when successful
func (m *BookingService) GetMaximumAttendeesCount()(*int32) {
    return m.maximumAttendeesCount
}
// GetNotes gets the notes property value. Additional information about this service.
// returns a *string when successful
func (m *BookingService) GetNotes()(*string) {
    return m.notes
}
// GetPostBuffer gets the postBuffer property value. The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
// returns a *ISODuration when successful
func (m *BookingService) GetPostBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.postBuffer
}
// GetPreBuffer gets the preBuffer property value. The time to buffer before an appointment for this service can start.
// returns a *ISODuration when successful
func (m *BookingService) GetPreBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.preBuffer
}
// GetSchedulingPolicy gets the schedulingPolicy property value. The set of policies that determine how appointments for this type of service should be created and managed.
// returns a BookingSchedulingPolicyable when successful
func (m *BookingService) GetSchedulingPolicy()(BookingSchedulingPolicyable) {
    return m.schedulingPolicy
}
// GetSmsNotificationsEnabled gets the smsNotificationsEnabled property value. True indicates SMS notifications can be sent to the customers for the appointment of the service. Default value is false.
// returns a *bool when successful
func (m *BookingService) GetSmsNotificationsEnabled()(*bool) {
    return m.smsNotificationsEnabled
}
// GetStaffMemberIds gets the staffMemberIds property value. Represents those staff members who provide this service.
// returns a []string when successful
func (m *BookingService) GetStaffMemberIds()([]string) {
    return m.staffMemberIds
}
// GetWebUrl gets the webUrl property value. The URL a customer uses to access the service.
// returns a *string when successful
func (m *BookingService) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *BookingService) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("additionalInformation", m.GetAdditionalInformation())
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
        err = writer.WriteISODurationValue("defaultDuration", m.GetDefaultDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultLocation", m.GetDefaultLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("defaultPrice", m.GetDefaultPrice())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultPriceType() != nil {
        cast := (*m.GetDefaultPriceType()).String()
        err = writer.WriteStringValue("defaultPriceType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefaultReminders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefaultReminders()))
        for i, v := range m.GetDefaultReminders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("defaultReminders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteBoolValue("isAnonymousJoinEnabled", m.GetIsAnonymousJoinEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCustomerAllowedToManageBooking", m.GetIsCustomerAllowedToManageBooking())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isHiddenFromCustomers", m.GetIsHiddenFromCustomers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLocationOnline", m.GetIsLocationOnline())
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
        err = writer.WriteInt32Value("maximumAttendeesCount", m.GetMaximumAttendeesCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("postBuffer", m.GetPostBuffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("preBuffer", m.GetPreBuffer())
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
    {
        err = writer.WriteBoolValue("smsNotificationsEnabled", m.GetSmsNotificationsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetStaffMemberIds() != nil {
        err = writer.WriteCollectionOfStringValues("staffMemberIds", m.GetStaffMemberIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalInformation sets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
func (m *BookingService) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date, time, and time zone when the service was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BookingService) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomQuestions sets the customQuestions property value. Contains the set of custom questions associated with a particular service.
func (m *BookingService) SetCustomQuestions(value []BookingQuestionAssignmentable)() {
    m.customQuestions = value
}
// SetDefaultDuration sets the defaultDuration property value. The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
func (m *BookingService) SetDefaultDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.defaultDuration = value
}
// SetDefaultLocation sets the defaultLocation property value. The default physical location for the service.
func (m *BookingService) SetDefaultLocation(value Locationable)() {
    m.defaultLocation = value
}
// SetDefaultPrice sets the defaultPrice property value. The default monetary price for the service.
func (m *BookingService) SetDefaultPrice(value *float64)() {
    m.defaultPrice = value
}
// SetDefaultPriceType sets the defaultPriceType property value. Represents the type of pricing of a booking service.
func (m *BookingService) SetDefaultPriceType(value *BookingPriceType)() {
    m.defaultPriceType = value
}
// SetDefaultReminders sets the defaultReminders property value. The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
func (m *BookingService) SetDefaultReminders(value []BookingReminderable)() {
    m.defaultReminders = value
}
// SetDescription sets the description property value. A text description for the service.
func (m *BookingService) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. A service name.
func (m *BookingService) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsAnonymousJoinEnabled sets the isAnonymousJoinEnabled property value. Indicates if an anonymousJoinWebUrl(webrtcUrl) is generated for the appointment booked for this service. The default value is false.
func (m *BookingService) SetIsAnonymousJoinEnabled(value *bool)() {
    m.isAnonymousJoinEnabled = value
}
// SetIsCustomerAllowedToManageBooking sets the isCustomerAllowedToManageBooking property value. Indicates that the customer can manage bookings created by the staff. The default value is false.
func (m *BookingService) SetIsCustomerAllowedToManageBooking(value *bool)() {
    m.isCustomerAllowedToManageBooking = value
}
// SetIsHiddenFromCustomers sets the isHiddenFromCustomers property value. True indicates that this service isn't available to customers for booking.
func (m *BookingService) SetIsHiddenFromCustomers(value *bool)() {
    m.isHiddenFromCustomers = value
}
// SetIsLocationOnline sets the isLocationOnline property value. Indicates that the appointments for the service are held online. The default value is false.
func (m *BookingService) SetIsLocationOnline(value *bool)() {
    m.isLocationOnline = value
}
// SetLanguageTag sets the languageTag property value. The language of the self-service booking page.
func (m *BookingService) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The date, time, and time zone when the service was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BookingService) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetMaximumAttendeesCount sets the maximumAttendeesCount property value. The maximum number of customers allowed in a service. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
func (m *BookingService) SetMaximumAttendeesCount(value *int32)() {
    m.maximumAttendeesCount = value
}
// SetNotes sets the notes property value. Additional information about this service.
func (m *BookingService) SetNotes(value *string)() {
    m.notes = value
}
// SetPostBuffer sets the postBuffer property value. The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
func (m *BookingService) SetPostBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.postBuffer = value
}
// SetPreBuffer sets the preBuffer property value. The time to buffer before an appointment for this service can start.
func (m *BookingService) SetPreBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.preBuffer = value
}
// SetSchedulingPolicy sets the schedulingPolicy property value. The set of policies that determine how appointments for this type of service should be created and managed.
func (m *BookingService) SetSchedulingPolicy(value BookingSchedulingPolicyable)() {
    m.schedulingPolicy = value
}
// SetSmsNotificationsEnabled sets the smsNotificationsEnabled property value. True indicates SMS notifications can be sent to the customers for the appointment of the service. Default value is false.
func (m *BookingService) SetSmsNotificationsEnabled(value *bool)() {
    m.smsNotificationsEnabled = value
}
// SetStaffMemberIds sets the staffMemberIds property value. Represents those staff members who provide this service.
func (m *BookingService) SetStaffMemberIds(value []string)() {
    m.staffMemberIds = value
}
// SetWebUrl sets the webUrl property value. The URL a customer uses to access the service.
func (m *BookingService) SetWebUrl(value *string)() {
    m.webUrl = value
}
type BookingServiceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalInformation()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomQuestions()([]BookingQuestionAssignmentable)
    GetDefaultDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetDefaultLocation()(Locationable)
    GetDefaultPrice()(*float64)
    GetDefaultPriceType()(*BookingPriceType)
    GetDefaultReminders()([]BookingReminderable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsAnonymousJoinEnabled()(*bool)
    GetIsCustomerAllowedToManageBooking()(*bool)
    GetIsHiddenFromCustomers()(*bool)
    GetIsLocationOnline()(*bool)
    GetLanguageTag()(*string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMaximumAttendeesCount()(*int32)
    GetNotes()(*string)
    GetPostBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPreBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetSchedulingPolicy()(BookingSchedulingPolicyable)
    GetSmsNotificationsEnabled()(*bool)
    GetStaffMemberIds()([]string)
    GetWebUrl()(*string)
    SetAdditionalInformation(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomQuestions(value []BookingQuestionAssignmentable)()
    SetDefaultDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetDefaultLocation(value Locationable)()
    SetDefaultPrice(value *float64)()
    SetDefaultPriceType(value *BookingPriceType)()
    SetDefaultReminders(value []BookingReminderable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsAnonymousJoinEnabled(value *bool)()
    SetIsCustomerAllowedToManageBooking(value *bool)()
    SetIsHiddenFromCustomers(value *bool)()
    SetIsLocationOnline(value *bool)()
    SetLanguageTag(value *string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMaximumAttendeesCount(value *int32)()
    SetNotes(value *string)()
    SetPostBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPreBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetSchedulingPolicy(value BookingSchedulingPolicyable)()
    SetSmsNotificationsEnabled(value *bool)()
    SetStaffMemberIds(value []string)()
    SetWebUrl(value *string)()
}