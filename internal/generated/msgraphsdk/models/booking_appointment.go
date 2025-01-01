package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingAppointment represents a booked appointment of a service by a customer in a business.
type BookingAppointment struct {
    Entity
    // Additional information that is sent to the customer when an appointment is confirmed.
    additionalInformation *string
    // The URL of the meeting to join anonymously.
    anonymousJoinWebUrl *string
    // The custom label that can be stamped on this appointment by users.
    appointmentLabel *string
    // The date, time, and time zone when the appointment was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The SMTP address of the bookingCustomer who books the appointment.
    customerEmailAddress *string
    // The customer's name.
    customerName *string
    // Notes from the customer associated with this appointment. You can get the value only when you read this bookingAppointment by its ID. You can set this property only when you initially create an appointment with a new customer.
    customerNotes *string
    // The customer's phone number.
    customerPhone *string
    // A collection of customer properties for an appointment. An appointment contains a list of customer information and each unit will indicate the properties of a customer who is part of that appointment. Optional.
    customers []BookingCustomerInformationBaseable
    // The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
    customerTimeZone *string
    // The length of the appointment, denoted in ISO8601 format.
    duration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The endDateTime property
    endDateTime DateTimeTimeZoneable
    // The current number of customers in the appointment.
    filledAttendeesCount *int32
    // Indicates that the customer can manage bookings created by the staff. The default value is false.
    isCustomerAllowedToManageBooking *bool
    // Indicates that the appointment is held online. The default value is false.
    isLocationOnline *bool
    // The URL of the online meeting for the appointment.
    joinWebUrl *string
    // The date, time, and time zone when the booking business was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The maximum number of customers allowed in an appointment. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
    maximumAttendeesCount *int32
    // If true indicates that the bookingCustomer for this appointment doesn't wish to receive a confirmation for this appointment.
    optOutOfCustomerEmail *bool
    // The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
    postBuffer *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
    preBuffer *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The regular price for an appointment for the specified bookingService.
    price *float64
    // Represents the type of pricing of a booking service.
    priceType *BookingPriceType
    // The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
    reminders []BookingReminderable
    // Another tracking ID for the appointment, if the appointment was created directly by the customer on the scheduling page, as opposed to by a staff member on behalf of the customer.
    selfServiceAppointmentId *string
    // The ID of the bookingService associated with this appointment.
    serviceId *string
    // The location where the service is delivered.
    serviceLocation Locationable
    // The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it's computed from the service associated with the appointment by the serviceId property.
    serviceName *string
    // Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
    serviceNotes *string
    // If true, indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
    smsNotificationsEnabled *bool
    // The ID of each bookingStaffMember who is scheduled in this appointment.
    staffMemberIds []string
    // The startDateTime property
    startDateTime DateTimeTimeZoneable
}
// NewBookingAppointment instantiates a new BookingAppointment and sets the default values.
func NewBookingAppointment()(*BookingAppointment) {
    m := &BookingAppointment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBookingAppointmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBookingAppointmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingAppointment(), nil
}
// GetAdditionalInformation gets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
// returns a *string when successful
func (m *BookingAppointment) GetAdditionalInformation()(*string) {
    return m.additionalInformation
}
// GetAnonymousJoinWebUrl gets the anonymousJoinWebUrl property value. The URL of the meeting to join anonymously.
// returns a *string when successful
func (m *BookingAppointment) GetAnonymousJoinWebUrl()(*string) {
    return m.anonymousJoinWebUrl
}
// GetAppointmentLabel gets the appointmentLabel property value. The custom label that can be stamped on this appointment by users.
// returns a *string when successful
func (m *BookingAppointment) GetAppointmentLabel()(*string) {
    return m.appointmentLabel
}
// GetCreatedDateTime gets the createdDateTime property value. The date, time, and time zone when the appointment was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *BookingAppointment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomerEmailAddress gets the customerEmailAddress property value. The SMTP address of the bookingCustomer who books the appointment.
// returns a *string when successful
func (m *BookingAppointment) GetCustomerEmailAddress()(*string) {
    return m.customerEmailAddress
}
// GetCustomerName gets the customerName property value. The customer's name.
// returns a *string when successful
func (m *BookingAppointment) GetCustomerName()(*string) {
    return m.customerName
}
// GetCustomerNotes gets the customerNotes property value. Notes from the customer associated with this appointment. You can get the value only when you read this bookingAppointment by its ID. You can set this property only when you initially create an appointment with a new customer.
// returns a *string when successful
func (m *BookingAppointment) GetCustomerNotes()(*string) {
    return m.customerNotes
}
// GetCustomerPhone gets the customerPhone property value. The customer's phone number.
// returns a *string when successful
func (m *BookingAppointment) GetCustomerPhone()(*string) {
    return m.customerPhone
}
// GetCustomers gets the customers property value. A collection of customer properties for an appointment. An appointment contains a list of customer information and each unit will indicate the properties of a customer who is part of that appointment. Optional.
// returns a []BookingCustomerInformationBaseable when successful
func (m *BookingAppointment) GetCustomers()([]BookingCustomerInformationBaseable) {
    return m.customers
}
// GetCustomerTimeZone gets the customerTimeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
// returns a *string when successful
func (m *BookingAppointment) GetCustomerTimeZone()(*string) {
    return m.customerTimeZone
}
// GetDuration gets the duration property value. The length of the appointment, denoted in ISO8601 format.
// returns a *ISODuration when successful
func (m *BookingAppointment) GetDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.duration
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
// returns a DateTimeTimeZoneable when successful
func (m *BookingAppointment) GetEndDateTime()(DateTimeTimeZoneable) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BookingAppointment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["anonymousJoinWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnonymousJoinWebUrl(val)
        }
        return nil
    }
    res["appointmentLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppointmentLabel(val)
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
    res["customerEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerEmailAddress(val)
        }
        return nil
    }
    res["customerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerName(val)
        }
        return nil
    }
    res["customerNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerNotes(val)
        }
        return nil
    }
    res["customerPhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerPhone(val)
        }
        return nil
    }
    res["customers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingCustomerInformationBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCustomerInformationBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BookingCustomerInformationBaseable)
                }
            }
            m.SetCustomers(res)
        }
        return nil
    }
    res["customerTimeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerTimeZone(val)
        }
        return nil
    }
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["filledAttendeesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilledAttendeesCount(val)
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
    res["joinWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinWebUrl(val)
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
    res["optOutOfCustomerEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptOutOfCustomerEmail(val)
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
    res["price"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrice(val)
        }
        return nil
    }
    res["priceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingPriceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriceType(val.(*BookingPriceType))
        }
        return nil
    }
    res["reminders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetReminders(res)
        }
        return nil
    }
    res["selfServiceAppointmentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelfServiceAppointmentId(val)
        }
        return nil
    }
    res["serviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceId(val)
        }
        return nil
    }
    res["serviceLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceLocation(val.(Locationable))
        }
        return nil
    }
    res["serviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceName(val)
        }
        return nil
    }
    res["serviceNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceNotes(val)
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
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetFilledAttendeesCount gets the filledAttendeesCount property value. The current number of customers in the appointment.
// returns a *int32 when successful
func (m *BookingAppointment) GetFilledAttendeesCount()(*int32) {
    return m.filledAttendeesCount
}
// GetIsCustomerAllowedToManageBooking gets the isCustomerAllowedToManageBooking property value. Indicates that the customer can manage bookings created by the staff. The default value is false.
// returns a *bool when successful
func (m *BookingAppointment) GetIsCustomerAllowedToManageBooking()(*bool) {
    return m.isCustomerAllowedToManageBooking
}
// GetIsLocationOnline gets the isLocationOnline property value. Indicates that the appointment is held online. The default value is false.
// returns a *bool when successful
func (m *BookingAppointment) GetIsLocationOnline()(*bool) {
    return m.isLocationOnline
}
// GetJoinWebUrl gets the joinWebUrl property value. The URL of the online meeting for the appointment.
// returns a *string when successful
func (m *BookingAppointment) GetJoinWebUrl()(*string) {
    return m.joinWebUrl
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The date, time, and time zone when the booking business was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *BookingAppointment) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetMaximumAttendeesCount gets the maximumAttendeesCount property value. The maximum number of customers allowed in an appointment. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
// returns a *int32 when successful
func (m *BookingAppointment) GetMaximumAttendeesCount()(*int32) {
    return m.maximumAttendeesCount
}
// GetOptOutOfCustomerEmail gets the optOutOfCustomerEmail property value. If true indicates that the bookingCustomer for this appointment doesn't wish to receive a confirmation for this appointment.
// returns a *bool when successful
func (m *BookingAppointment) GetOptOutOfCustomerEmail()(*bool) {
    return m.optOutOfCustomerEmail
}
// GetPostBuffer gets the postBuffer property value. The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
// returns a *ISODuration when successful
func (m *BookingAppointment) GetPostBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.postBuffer
}
// GetPreBuffer gets the preBuffer property value. The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
// returns a *ISODuration when successful
func (m *BookingAppointment) GetPreBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.preBuffer
}
// GetPrice gets the price property value. The regular price for an appointment for the specified bookingService.
// returns a *float64 when successful
func (m *BookingAppointment) GetPrice()(*float64) {
    return m.price
}
// GetPriceType gets the priceType property value. Represents the type of pricing of a booking service.
// returns a *BookingPriceType when successful
func (m *BookingAppointment) GetPriceType()(*BookingPriceType) {
    return m.priceType
}
// GetReminders gets the reminders property value. The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
// returns a []BookingReminderable when successful
func (m *BookingAppointment) GetReminders()([]BookingReminderable) {
    return m.reminders
}
// GetSelfServiceAppointmentId gets the selfServiceAppointmentId property value. Another tracking ID for the appointment, if the appointment was created directly by the customer on the scheduling page, as opposed to by a staff member on behalf of the customer.
// returns a *string when successful
func (m *BookingAppointment) GetSelfServiceAppointmentId()(*string) {
    return m.selfServiceAppointmentId
}
// GetServiceId gets the serviceId property value. The ID of the bookingService associated with this appointment.
// returns a *string when successful
func (m *BookingAppointment) GetServiceId()(*string) {
    return m.serviceId
}
// GetServiceLocation gets the serviceLocation property value. The location where the service is delivered.
// returns a Locationable when successful
func (m *BookingAppointment) GetServiceLocation()(Locationable) {
    return m.serviceLocation
}
// GetServiceName gets the serviceName property value. The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it's computed from the service associated with the appointment by the serviceId property.
// returns a *string when successful
func (m *BookingAppointment) GetServiceName()(*string) {
    return m.serviceName
}
// GetServiceNotes gets the serviceNotes property value. Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
// returns a *string when successful
func (m *BookingAppointment) GetServiceNotes()(*string) {
    return m.serviceNotes
}
// GetSmsNotificationsEnabled gets the smsNotificationsEnabled property value. If true, indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
// returns a *bool when successful
func (m *BookingAppointment) GetSmsNotificationsEnabled()(*bool) {
    return m.smsNotificationsEnabled
}
// GetStaffMemberIds gets the staffMemberIds property value. The ID of each bookingStaffMember who is scheduled in this appointment.
// returns a []string when successful
func (m *BookingAppointment) GetStaffMemberIds()([]string) {
    return m.staffMemberIds
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
// returns a DateTimeTimeZoneable when successful
func (m *BookingAppointment) GetStartDateTime()(DateTimeTimeZoneable) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *BookingAppointment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("anonymousJoinWebUrl", m.GetAnonymousJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appointmentLabel", m.GetAppointmentLabel())
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
        err = writer.WriteStringValue("customerEmailAddress", m.GetCustomerEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerName", m.GetCustomerName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerNotes", m.GetCustomerNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerPhone", m.GetCustomerPhone())
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
    {
        err = writer.WriteStringValue("customerTimeZone", m.GetCustomerTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("endDateTime", m.GetEndDateTime())
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
        err = writer.WriteBoolValue("isLocationOnline", m.GetIsLocationOnline())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
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
        err = writer.WriteBoolValue("optOutOfCustomerEmail", m.GetOptOutOfCustomerEmail())
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
        err = writer.WriteFloat64Value("price", m.GetPrice())
        if err != nil {
            return err
        }
    }
    if m.GetPriceType() != nil {
        cast := (*m.GetPriceType()).String()
        err = writer.WriteStringValue("priceType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetReminders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReminders()))
        for i, v := range m.GetReminders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("reminders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("selfServiceAppointmentId", m.GetSelfServiceAppointmentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceId", m.GetServiceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("serviceLocation", m.GetServiceLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceName", m.GetServiceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceNotes", m.GetServiceNotes())
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
    {
        err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalInformation sets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
func (m *BookingAppointment) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
// SetAnonymousJoinWebUrl sets the anonymousJoinWebUrl property value. The URL of the meeting to join anonymously.
func (m *BookingAppointment) SetAnonymousJoinWebUrl(value *string)() {
    m.anonymousJoinWebUrl = value
}
// SetAppointmentLabel sets the appointmentLabel property value. The custom label that can be stamped on this appointment by users.
func (m *BookingAppointment) SetAppointmentLabel(value *string)() {
    m.appointmentLabel = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date, time, and time zone when the appointment was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BookingAppointment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomerEmailAddress sets the customerEmailAddress property value. The SMTP address of the bookingCustomer who books the appointment.
func (m *BookingAppointment) SetCustomerEmailAddress(value *string)() {
    m.customerEmailAddress = value
}
// SetCustomerName sets the customerName property value. The customer's name.
func (m *BookingAppointment) SetCustomerName(value *string)() {
    m.customerName = value
}
// SetCustomerNotes sets the customerNotes property value. Notes from the customer associated with this appointment. You can get the value only when you read this bookingAppointment by its ID. You can set this property only when you initially create an appointment with a new customer.
func (m *BookingAppointment) SetCustomerNotes(value *string)() {
    m.customerNotes = value
}
// SetCustomerPhone sets the customerPhone property value. The customer's phone number.
func (m *BookingAppointment) SetCustomerPhone(value *string)() {
    m.customerPhone = value
}
// SetCustomers sets the customers property value. A collection of customer properties for an appointment. An appointment contains a list of customer information and each unit will indicate the properties of a customer who is part of that appointment. Optional.
func (m *BookingAppointment) SetCustomers(value []BookingCustomerInformationBaseable)() {
    m.customers = value
}
// SetCustomerTimeZone sets the customerTimeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
func (m *BookingAppointment) SetCustomerTimeZone(value *string)() {
    m.customerTimeZone = value
}
// SetDuration sets the duration property value. The length of the appointment, denoted in ISO8601 format.
func (m *BookingAppointment) SetDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.duration = value
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *BookingAppointment) SetEndDateTime(value DateTimeTimeZoneable)() {
    m.endDateTime = value
}
// SetFilledAttendeesCount sets the filledAttendeesCount property value. The current number of customers in the appointment.
func (m *BookingAppointment) SetFilledAttendeesCount(value *int32)() {
    m.filledAttendeesCount = value
}
// SetIsCustomerAllowedToManageBooking sets the isCustomerAllowedToManageBooking property value. Indicates that the customer can manage bookings created by the staff. The default value is false.
func (m *BookingAppointment) SetIsCustomerAllowedToManageBooking(value *bool)() {
    m.isCustomerAllowedToManageBooking = value
}
// SetIsLocationOnline sets the isLocationOnline property value. Indicates that the appointment is held online. The default value is false.
func (m *BookingAppointment) SetIsLocationOnline(value *bool)() {
    m.isLocationOnline = value
}
// SetJoinWebUrl sets the joinWebUrl property value. The URL of the online meeting for the appointment.
func (m *BookingAppointment) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The date, time, and time zone when the booking business was last updated. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BookingAppointment) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetMaximumAttendeesCount sets the maximumAttendeesCount property value. The maximum number of customers allowed in an appointment. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create bookingCustomer operation.
func (m *BookingAppointment) SetMaximumAttendeesCount(value *int32)() {
    m.maximumAttendeesCount = value
}
// SetOptOutOfCustomerEmail sets the optOutOfCustomerEmail property value. If true indicates that the bookingCustomer for this appointment doesn't wish to receive a confirmation for this appointment.
func (m *BookingAppointment) SetOptOutOfCustomerEmail(value *bool)() {
    m.optOutOfCustomerEmail = value
}
// SetPostBuffer sets the postBuffer property value. The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) SetPostBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.postBuffer = value
}
// SetPreBuffer sets the preBuffer property value. The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed in ISO8601 format.
func (m *BookingAppointment) SetPreBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.preBuffer = value
}
// SetPrice sets the price property value. The regular price for an appointment for the specified bookingService.
func (m *BookingAppointment) SetPrice(value *float64)() {
    m.price = value
}
// SetPriceType sets the priceType property value. Represents the type of pricing of a booking service.
func (m *BookingAppointment) SetPriceType(value *BookingPriceType)() {
    m.priceType = value
}
// SetReminders sets the reminders property value. The collection of customer reminders sent for this appointment. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) SetReminders(value []BookingReminderable)() {
    m.reminders = value
}
// SetSelfServiceAppointmentId sets the selfServiceAppointmentId property value. Another tracking ID for the appointment, if the appointment was created directly by the customer on the scheduling page, as opposed to by a staff member on behalf of the customer.
func (m *BookingAppointment) SetSelfServiceAppointmentId(value *string)() {
    m.selfServiceAppointmentId = value
}
// SetServiceId sets the serviceId property value. The ID of the bookingService associated with this appointment.
func (m *BookingAppointment) SetServiceId(value *string)() {
    m.serviceId = value
}
// SetServiceLocation sets the serviceLocation property value. The location where the service is delivered.
func (m *BookingAppointment) SetServiceLocation(value Locationable)() {
    m.serviceLocation = value
}
// SetServiceName sets the serviceName property value. The name of the bookingService associated with this appointment.This property is optional when creating a new appointment. If not specified, it's computed from the service associated with the appointment by the serviceId property.
func (m *BookingAppointment) SetServiceName(value *string)() {
    m.serviceName = value
}
// SetServiceNotes sets the serviceNotes property value. Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by its ID.
func (m *BookingAppointment) SetServiceNotes(value *string)() {
    m.serviceNotes = value
}
// SetSmsNotificationsEnabled sets the smsNotificationsEnabled property value. If true, indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
func (m *BookingAppointment) SetSmsNotificationsEnabled(value *bool)() {
    m.smsNotificationsEnabled = value
}
// SetStaffMemberIds sets the staffMemberIds property value. The ID of each bookingStaffMember who is scheduled in this appointment.
func (m *BookingAppointment) SetStaffMemberIds(value []string)() {
    m.staffMemberIds = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *BookingAppointment) SetStartDateTime(value DateTimeTimeZoneable)() {
    m.startDateTime = value
}
type BookingAppointmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalInformation()(*string)
    GetAnonymousJoinWebUrl()(*string)
    GetAppointmentLabel()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerEmailAddress()(*string)
    GetCustomerName()(*string)
    GetCustomerNotes()(*string)
    GetCustomerPhone()(*string)
    GetCustomers()([]BookingCustomerInformationBaseable)
    GetCustomerTimeZone()(*string)
    GetDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetEndDateTime()(DateTimeTimeZoneable)
    GetFilledAttendeesCount()(*int32)
    GetIsCustomerAllowedToManageBooking()(*bool)
    GetIsLocationOnline()(*bool)
    GetJoinWebUrl()(*string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMaximumAttendeesCount()(*int32)
    GetOptOutOfCustomerEmail()(*bool)
    GetPostBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPreBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPrice()(*float64)
    GetPriceType()(*BookingPriceType)
    GetReminders()([]BookingReminderable)
    GetSelfServiceAppointmentId()(*string)
    GetServiceId()(*string)
    GetServiceLocation()(Locationable)
    GetServiceName()(*string)
    GetServiceNotes()(*string)
    GetSmsNotificationsEnabled()(*bool)
    GetStaffMemberIds()([]string)
    GetStartDateTime()(DateTimeTimeZoneable)
    SetAdditionalInformation(value *string)()
    SetAnonymousJoinWebUrl(value *string)()
    SetAppointmentLabel(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerEmailAddress(value *string)()
    SetCustomerName(value *string)()
    SetCustomerNotes(value *string)()
    SetCustomerPhone(value *string)()
    SetCustomers(value []BookingCustomerInformationBaseable)()
    SetCustomerTimeZone(value *string)()
    SetDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetEndDateTime(value DateTimeTimeZoneable)()
    SetFilledAttendeesCount(value *int32)()
    SetIsCustomerAllowedToManageBooking(value *bool)()
    SetIsLocationOnline(value *bool)()
    SetJoinWebUrl(value *string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMaximumAttendeesCount(value *int32)()
    SetOptOutOfCustomerEmail(value *bool)()
    SetPostBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPreBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPrice(value *float64)()
    SetPriceType(value *BookingPriceType)()
    SetReminders(value []BookingReminderable)()
    SetSelfServiceAppointmentId(value *string)()
    SetServiceId(value *string)()
    SetServiceLocation(value Locationable)()
    SetServiceName(value *string)()
    SetServiceNotes(value *string)()
    SetSmsNotificationsEnabled(value *bool)()
    SetStaffMemberIds(value []string)()
    SetStartDateTime(value DateTimeTimeZoneable)()
}