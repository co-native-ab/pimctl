package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserSettings struct {
    Entity
    // Reflects the organization level setting controlling delegate access to the trending API. When set to true, the organization doesn't have access to Office Delve. The relevancy of the content displayed in Microsoft 365, for example in Suggested sites in SharePoint Home and the Discover view in OneDrive for work or school is affected for the whole organization. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
    contributionToContentDiscoveryAsOrganizationDisabled *bool
    // When set to true, the delegate access to the user's trending API is disabled. When set to true, documents in the user's Office Delve are disabled. When set to true, the relevancy of the content displayed in Microsoft 365, for example in Suggested sites in SharePoint Home and the Discover view in OneDrive for work or school is affected. Users can control this setting in Office Delve.
    contributionToContentDiscoveryDisabled *bool
    // The user's settings for the visibility of meeting hour insights, and insights derived between a user and other items in Microsoft 365, such as documents or sites. Get userInsightsSettings through this navigation property.
    itemInsights UserInsightsSettingsable
    // The shiftPreferences property
    shiftPreferences ShiftPreferencesable
    // The storage property
    storage UserStorageable
    // The windows property
    windows []WindowsSettingable
}
// NewUserSettings instantiates a new UserSettings and sets the default values.
func NewUserSettings()(*UserSettings) {
    m := &UserSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSettings(), nil
}
// GetContributionToContentDiscoveryAsOrganizationDisabled gets the contributionToContentDiscoveryAsOrganizationDisabled property value. Reflects the organization level setting controlling delegate access to the trending API. When set to true, the organization doesn't have access to Office Delve. The relevancy of the content displayed in Microsoft 365, for example in Suggested sites in SharePoint Home and the Discover view in OneDrive for work or school is affected for the whole organization. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
// returns a *bool when successful
func (m *UserSettings) GetContributionToContentDiscoveryAsOrganizationDisabled()(*bool) {
    return m.contributionToContentDiscoveryAsOrganizationDisabled
}
// GetContributionToContentDiscoveryDisabled gets the contributionToContentDiscoveryDisabled property value. When set to true, the delegate access to the user's trending API is disabled. When set to true, documents in the user's Office Delve are disabled. When set to true, the relevancy of the content displayed in Microsoft 365, for example in Suggested sites in SharePoint Home and the Discover view in OneDrive for work or school is affected. Users can control this setting in Office Delve.
// returns a *bool when successful
func (m *UserSettings) GetContributionToContentDiscoveryDisabled()(*bool) {
    return m.contributionToContentDiscoveryDisabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contributionToContentDiscoveryAsOrganizationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributionToContentDiscoveryAsOrganizationDisabled(val)
        }
        return nil
    }
    res["contributionToContentDiscoveryDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributionToContentDiscoveryDisabled(val)
        }
        return nil
    }
    res["itemInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserInsightsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemInsights(val.(UserInsightsSettingsable))
        }
        return nil
    }
    res["shiftPreferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShiftPreferencesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShiftPreferences(val.(ShiftPreferencesable))
        }
        return nil
    }
    res["storage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserStorageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorage(val.(UserStorageable))
        }
        return nil
    }
    res["windows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsSettingable)
                }
            }
            m.SetWindows(res)
        }
        return nil
    }
    return res
}
// GetItemInsights gets the itemInsights property value. The user's settings for the visibility of meeting hour insights, and insights derived between a user and other items in Microsoft 365, such as documents or sites. Get userInsightsSettings through this navigation property.
// returns a UserInsightsSettingsable when successful
func (m *UserSettings) GetItemInsights()(UserInsightsSettingsable) {
    return m.itemInsights
}
// GetShiftPreferences gets the shiftPreferences property value. The shiftPreferences property
// returns a ShiftPreferencesable when successful
func (m *UserSettings) GetShiftPreferences()(ShiftPreferencesable) {
    return m.shiftPreferences
}
// GetStorage gets the storage property value. The storage property
// returns a UserStorageable when successful
func (m *UserSettings) GetStorage()(UserStorageable) {
    return m.storage
}
// GetWindows gets the windows property value. The windows property
// returns a []WindowsSettingable when successful
func (m *UserSettings) GetWindows()([]WindowsSettingable) {
    return m.windows
}
// Serialize serializes information the current object
func (m *UserSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("contributionToContentDiscoveryAsOrganizationDisabled", m.GetContributionToContentDiscoveryAsOrganizationDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contributionToContentDiscoveryDisabled", m.GetContributionToContentDiscoveryDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("itemInsights", m.GetItemInsights())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shiftPreferences", m.GetShiftPreferences())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("storage", m.GetStorage())
        if err != nil {
            return err
        }
    }
    if m.GetWindows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindows()))
        for i, v := range m.GetWindows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windows", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContributionToContentDiscoveryAsOrganizationDisabled sets the contributionToContentDiscoveryAsOrganizationDisabled property value. Reflects the organization level setting controlling delegate access to the trending API. When set to true, the organization doesn't have access to Office Delve. The relevancy of the content displayed in Microsoft 365, for example in Suggested sites in SharePoint Home and the Discover view in OneDrive for work or school is affected for the whole organization. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
func (m *UserSettings) SetContributionToContentDiscoveryAsOrganizationDisabled(value *bool)() {
    m.contributionToContentDiscoveryAsOrganizationDisabled = value
}
// SetContributionToContentDiscoveryDisabled sets the contributionToContentDiscoveryDisabled property value. When set to true, the delegate access to the user's trending API is disabled. When set to true, documents in the user's Office Delve are disabled. When set to true, the relevancy of the content displayed in Microsoft 365, for example in Suggested sites in SharePoint Home and the Discover view in OneDrive for work or school is affected. Users can control this setting in Office Delve.
func (m *UserSettings) SetContributionToContentDiscoveryDisabled(value *bool)() {
    m.contributionToContentDiscoveryDisabled = value
}
// SetItemInsights sets the itemInsights property value. The user's settings for the visibility of meeting hour insights, and insights derived between a user and other items in Microsoft 365, such as documents or sites. Get userInsightsSettings through this navigation property.
func (m *UserSettings) SetItemInsights(value UserInsightsSettingsable)() {
    m.itemInsights = value
}
// SetShiftPreferences sets the shiftPreferences property value. The shiftPreferences property
func (m *UserSettings) SetShiftPreferences(value ShiftPreferencesable)() {
    m.shiftPreferences = value
}
// SetStorage sets the storage property value. The storage property
func (m *UserSettings) SetStorage(value UserStorageable)() {
    m.storage = value
}
// SetWindows sets the windows property value. The windows property
func (m *UserSettings) SetWindows(value []WindowsSettingable)() {
    m.windows = value
}
type UserSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContributionToContentDiscoveryAsOrganizationDisabled()(*bool)
    GetContributionToContentDiscoveryDisabled()(*bool)
    GetItemInsights()(UserInsightsSettingsable)
    GetShiftPreferences()(ShiftPreferencesable)
    GetStorage()(UserStorageable)
    GetWindows()([]WindowsSettingable)
    SetContributionToContentDiscoveryAsOrganizationDisabled(value *bool)()
    SetContributionToContentDiscoveryDisabled(value *bool)()
    SetItemInsights(value UserInsightsSettingsable)()
    SetShiftPreferences(value ShiftPreferencesable)()
    SetStorage(value UserStorageable)()
    SetWindows(value []WindowsSettingable)()
}
