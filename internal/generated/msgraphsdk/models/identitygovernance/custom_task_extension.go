package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type CustomTaskExtension struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CustomCalloutExtension
    // The callback configuration for a custom task extension.
    callbackConfiguration ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CustomExtensionCallbackConfigurationable
    // The unique identifier of the Microsoft Entra user that created the custom task extension.Supports $filter(eq, ne) and $expand.
    createdBy ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable
    // When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the Microsoft Entra user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
    lastModifiedBy ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable
    // When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewCustomTaskExtension instantiates a new CustomTaskExtension and sets the default values.
func NewCustomTaskExtension()(*CustomTaskExtension) {
    m := &CustomTaskExtension{
        CustomCalloutExtension: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewCustomCalloutExtension(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.customTaskExtension"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomTaskExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomTaskExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomTaskExtension(), nil
}
// GetCallbackConfiguration gets the callbackConfiguration property value. The callback configuration for a custom task extension.
// returns a CustomExtensionCallbackConfigurationable when successful
func (m *CustomTaskExtension) GetCallbackConfiguration()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CustomExtensionCallbackConfigurationable) {
    return m.callbackConfiguration
}
// GetCreatedBy gets the createdBy property value. The unique identifier of the Microsoft Entra user that created the custom task extension.Supports $filter(eq, ne) and $expand.
// returns a Userable when successful
func (m *CustomTaskExtension) GetCreatedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
// returns a *Time when successful
func (m *CustomTaskExtension) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomTaskExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomCalloutExtension.GetFieldDeserializers()
    res["callbackConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateCustomExtensionCallbackConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackConfiguration(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CustomExtensionCallbackConfigurationable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable))
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
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The unique identifier of the Microsoft Entra user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
// returns a Userable when successful
func (m *CustomTaskExtension) GetLastModifiedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
// returns a *Time when successful
func (m *CustomTaskExtension) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// Serialize serializes information the current object
func (m *CustomTaskExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomCalloutExtension.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("callbackConfiguration", m.GetCallbackConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCallbackConfiguration sets the callbackConfiguration property value. The callback configuration for a custom task extension.
func (m *CustomTaskExtension) SetCallbackConfiguration(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CustomExtensionCallbackConfigurationable)() {
    m.callbackConfiguration = value
}
// SetCreatedBy sets the createdBy property value. The unique identifier of the Microsoft Entra user that created the custom task extension.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) SetCreatedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The unique identifier of the Microsoft Entra user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) SetLastModifiedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
type CustomTaskExtensionable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CustomCalloutExtensionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallbackConfiguration()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CustomExtensionCallbackConfigurationable)
    GetCreatedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCallbackConfiguration(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CustomExtensionCallbackConfigurationable)()
    SetCreatedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Userable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
