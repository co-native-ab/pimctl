package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PrinterShare struct {
    PrinterBase
    // If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties.
    allowAllUsers *bool
    // The groups whose users have access to print using the printer.
    allowedGroups []Groupable
    // The users who have access to print using the printer.
    allowedUsers []Userable
    // The DateTimeOffset when the printer share was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The printer that this printer share is related to.
    printer Printerable
    // Additional data for a printer share as viewed by the signed-in user.
    viewPoint PrinterShareViewpointable
}
// NewPrinterShare instantiates a new PrinterShare and sets the default values.
func NewPrinterShare()(*PrinterShare) {
    m := &PrinterShare{
        PrinterBase: *NewPrinterBase(),
    }
    odataTypeValue := "#microsoft.graph.printerShare"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePrinterShareFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrinterShareFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinterShare(), nil
}
// GetAllowAllUsers gets the allowAllUsers property value. If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties.
// returns a *bool when successful
func (m *PrinterShare) GetAllowAllUsers()(*bool) {
    return m.allowAllUsers
}
// GetAllowedGroups gets the allowedGroups property value. The groups whose users have access to print using the printer.
// returns a []Groupable when successful
func (m *PrinterShare) GetAllowedGroups()([]Groupable) {
    return m.allowedGroups
}
// GetAllowedUsers gets the allowedUsers property value. The users who have access to print using the printer.
// returns a []Userable when successful
func (m *PrinterShare) GetAllowedUsers()([]Userable) {
    return m.allowedUsers
}
// GetCreatedDateTime gets the createdDateTime property value. The DateTimeOffset when the printer share was created. Read-only.
// returns a *Time when successful
func (m *PrinterShare) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrinterShare) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrinterBase.GetFieldDeserializers()
    res["allowAllUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAllUsers(val)
        }
        return nil
    }
    res["allowedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Groupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Groupable)
                }
            }
            m.SetAllowedGroups(res)
        }
        return nil
    }
    res["allowedUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetAllowedUsers(res)
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
    res["printer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinter(val.(Printerable))
        }
        return nil
    }
    res["viewPoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterShareViewpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewPoint(val.(PrinterShareViewpointable))
        }
        return nil
    }
    return res
}
// GetPrinter gets the printer property value. The printer that this printer share is related to.
// returns a Printerable when successful
func (m *PrinterShare) GetPrinter()(Printerable) {
    return m.printer
}
// GetViewPoint gets the viewPoint property value. Additional data for a printer share as viewed by the signed-in user.
// returns a PrinterShareViewpointable when successful
func (m *PrinterShare) GetViewPoint()(PrinterShareViewpointable) {
    return m.viewPoint
}
// Serialize serializes information the current object
func (m *PrinterShare) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrinterBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAllUsers", m.GetAllowAllUsers())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAllowedGroups()))
        for i, v := range m.GetAllowedGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("allowedGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAllowedUsers()))
        for i, v := range m.GetAllowedUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("allowedUsers", cast)
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
        err = writer.WriteObjectValue("printer", m.GetPrinter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("viewPoint", m.GetViewPoint())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowAllUsers sets the allowAllUsers property value. If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties.
func (m *PrinterShare) SetAllowAllUsers(value *bool)() {
    m.allowAllUsers = value
}
// SetAllowedGroups sets the allowedGroups property value. The groups whose users have access to print using the printer.
func (m *PrinterShare) SetAllowedGroups(value []Groupable)() {
    m.allowedGroups = value
}
// SetAllowedUsers sets the allowedUsers property value. The users who have access to print using the printer.
func (m *PrinterShare) SetAllowedUsers(value []Userable)() {
    m.allowedUsers = value
}
// SetCreatedDateTime sets the createdDateTime property value. The DateTimeOffset when the printer share was created. Read-only.
func (m *PrinterShare) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetPrinter sets the printer property value. The printer that this printer share is related to.
func (m *PrinterShare) SetPrinter(value Printerable)() {
    m.printer = value
}
// SetViewPoint sets the viewPoint property value. Additional data for a printer share as viewed by the signed-in user.
func (m *PrinterShare) SetViewPoint(value PrinterShareViewpointable)() {
    m.viewPoint = value
}
type PrinterShareable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrinterBaseable
    GetAllowAllUsers()(*bool)
    GetAllowedGroups()([]Groupable)
    GetAllowedUsers()([]Userable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPrinter()(Printerable)
    GetViewPoint()(PrinterShareViewpointable)
    SetAllowAllUsers(value *bool)()
    SetAllowedGroups(value []Groupable)()
    SetAllowedUsers(value []Userable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPrinter(value Printerable)()
    SetViewPoint(value PrinterShareViewpointable)()
}
