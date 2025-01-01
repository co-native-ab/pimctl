package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagement singleton entity that acts as a container for all device management functionality.
type DeviceManagement struct {
    Entity
    // Apple push notification certificate.
    applePushNotificationCertificate ApplePushNotificationCertificateable
    // The Audit Events
    auditEvents []AuditEventable
    // The list of Compliance Management Partners configured by the tenant.
    complianceManagementPartners []ComplianceManagementPartnerable
    // The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
    conditionalAccessSettings OnPremisesConditionalAccessSettingsable
    // The list of detected apps associated with a device.
    detectedApps []DetectedAppable
    // The list of device categories with the tenant.
    deviceCategories []DeviceCategoryable
    // The device compliance policies.
    deviceCompliancePolicies []DeviceCompliancePolicyable
    // The device compliance state summary for this account.
    deviceCompliancePolicyDeviceStateSummary DeviceCompliancePolicyDeviceStateSummaryable
    // The summary states of compliance policy settings for this account.
    deviceCompliancePolicySettingStateSummaries []DeviceCompliancePolicySettingStateSummaryable
    // The device configuration device state summary for this account.
    deviceConfigurationDeviceStateSummaries DeviceConfigurationDeviceStateSummaryable
    // The device configurations.
    deviceConfigurations []DeviceConfigurationable
    // The list of device enrollment configurations
    deviceEnrollmentConfigurations []DeviceEnrollmentConfigurationable
    // The list of Device Management Partners configured by the tenant.
    deviceManagementPartners []DeviceManagementPartnerable
    // Device protection overview.
    deviceProtectionOverview DeviceProtectionOverviewable
    // The list of Exchange Connectors configured by the tenant.
    exchangeConnectors []DeviceManagementExchangeConnectorable
    // Collection of imported Windows autopilot devices.
    importedWindowsAutopilotDeviceIdentities []ImportedWindowsAutopilotDeviceIdentityable
    // Intune Account Id for given tenant
    intuneAccountId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
    intuneBrand IntuneBrandable
    // The IOS software update installation statuses for this account.
    iosUpdateStatuses []IosUpdateDeviceStatusable
    // Device overview
    managedDeviceOverview ManagedDeviceOverviewable
    // The list of managed devices.
    managedDevices []ManagedDeviceable
    // The collection property of MobileAppTroubleshootingEvent.
    mobileAppTroubleshootingEvents []MobileAppTroubleshootingEventable
    // The list of Mobile threat Defense connectors configured by the tenant.
    mobileThreatDefenseConnectors []MobileThreatDefenseConnectorable
    // The Notification Message Templates.
    notificationMessageTemplates []NotificationMessageTemplateable
    // The remote assist partners.
    remoteAssistancePartners []RemoteAssistancePartnerable
    // Reports singleton
    reports DeviceManagementReportsable
    // The Resource Operations.
    resourceOperations []ResourceOperationable
    // The Role Assignments.
    roleAssignments []DeviceAndAppManagementRoleAssignmentable
    // The Role Definitions.
    roleDefinitions []RoleDefinitionable
    // Account level settings.
    settings DeviceManagementSettingsable
    // The software update status summary.
    softwareUpdateStatusSummary SoftwareUpdateStatusSummaryable
    // Tenant mobile device management subscription state.
    subscriptionState *DeviceManagementSubscriptionState
    // The telecom expense management partners.
    telecomExpenseManagementPartners []TelecomExpenseManagementPartnerable
    // The terms and conditions associated with device management of the company.
    termsAndConditions []TermsAndConditionsable
    // The list of troubleshooting events for the tenant.
    troubleshootingEvents []DeviceManagementTroubleshootingEventable
    // User experience analytics appHealth Application Performance
    userExperienceAnalyticsAppHealthApplicationPerformance []UserExperienceAnalyticsAppHealthApplicationPerformanceable
    // User experience analytics appHealth Application Performance by App Version details
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable
    // User experience analytics appHealth Application Performance by App Version Device Id
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable
    // User experience analytics appHealth Application Performance by OS Version
    userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable
    // User experience analytics appHealth Model Performance
    userExperienceAnalyticsAppHealthDeviceModelPerformance []UserExperienceAnalyticsAppHealthDeviceModelPerformanceable
    // User experience analytics appHealth Device Performance
    userExperienceAnalyticsAppHealthDevicePerformance []UserExperienceAnalyticsAppHealthDevicePerformanceable
    // User experience analytics device performance details
    userExperienceAnalyticsAppHealthDevicePerformanceDetails []UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable
    // User experience analytics appHealth OS version Performance
    userExperienceAnalyticsAppHealthOSVersionPerformance []UserExperienceAnalyticsAppHealthOSVersionPerformanceable
    // User experience analytics appHealth overview
    userExperienceAnalyticsAppHealthOverview UserExperienceAnalyticsCategoryable
    // User experience analytics baselines
    userExperienceAnalyticsBaselines []UserExperienceAnalyticsBaselineable
    // User experience analytics categories
    userExperienceAnalyticsCategories []UserExperienceAnalyticsCategoryable
    // User experience analytics device performance
    userExperienceAnalyticsDevicePerformance []UserExperienceAnalyticsDevicePerformanceable
    // User experience analytics device scores
    userExperienceAnalyticsDeviceScores []UserExperienceAnalyticsDeviceScoresable
    // User experience analytics device Startup History
    userExperienceAnalyticsDeviceStartupHistory []UserExperienceAnalyticsDeviceStartupHistoryable
    // User experience analytics device Startup Processes
    userExperienceAnalyticsDeviceStartupProcesses []UserExperienceAnalyticsDeviceStartupProcessable
    // User experience analytics device Startup Process Performance
    userExperienceAnalyticsDeviceStartupProcessPerformance []UserExperienceAnalyticsDeviceStartupProcessPerformanceable
    // User experience analytics metric history
    userExperienceAnalyticsMetricHistory []UserExperienceAnalyticsMetricHistoryable
    // User experience analytics model scores
    userExperienceAnalyticsModelScores []UserExperienceAnalyticsModelScoresable
    // User experience analytics overview
    userExperienceAnalyticsOverview UserExperienceAnalyticsOverviewable
    // User experience analytics device Startup Score History
    userExperienceAnalyticsScoreHistory []UserExperienceAnalyticsScoreHistoryable
    // User experience analytics device settings
    userExperienceAnalyticsSettings UserExperienceAnalyticsSettingsable
    // User experience analytics work from anywhere hardware readiness metrics.
    userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable
    // User experience analytics work from anywhere metrics.
    userExperienceAnalyticsWorkFromAnywhereMetrics []UserExperienceAnalyticsWorkFromAnywhereMetricable
    // The user experience analytics work from anywhere model performance
    userExperienceAnalyticsWorkFromAnywhereModelPerformance []UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable
    // Virtual endpoint
    virtualEndpoint VirtualEndpointable
    // The Windows autopilot device identities contained collection.
    windowsAutopilotDeviceIdentities []WindowsAutopilotDeviceIdentityable
    // The windows information protection app learning summaries.
    windowsInformationProtectionAppLearningSummaries []WindowsInformationProtectionAppLearningSummaryable
    // The windows information protection network learning summaries.
    windowsInformationProtectionNetworkLearningSummaries []WindowsInformationProtectionNetworkLearningSummaryable
    // The list of affected malware in the tenant.
    windowsMalwareInformation []WindowsMalwareInformationable
    // Malware overview for windows devices.
    windowsMalwareOverview WindowsMalwareOverviewable
}
// NewDeviceManagement instantiates a new DeviceManagement and sets the default values.
func NewDeviceManagement()(*DeviceManagement) {
    m := &DeviceManagement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagement(), nil
}
// GetApplePushNotificationCertificate gets the applePushNotificationCertificate property value. Apple push notification certificate.
// returns a ApplePushNotificationCertificateable when successful
func (m *DeviceManagement) GetApplePushNotificationCertificate()(ApplePushNotificationCertificateable) {
    return m.applePushNotificationCertificate
}
// GetAuditEvents gets the auditEvents property value. The Audit Events
// returns a []AuditEventable when successful
func (m *DeviceManagement) GetAuditEvents()([]AuditEventable) {
    return m.auditEvents
}
// GetComplianceManagementPartners gets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
// returns a []ComplianceManagementPartnerable when successful
func (m *DeviceManagement) GetComplianceManagementPartners()([]ComplianceManagementPartnerable) {
    return m.complianceManagementPartners
}
// GetConditionalAccessSettings gets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
// returns a OnPremisesConditionalAccessSettingsable when successful
func (m *DeviceManagement) GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable) {
    return m.conditionalAccessSettings
}
// GetDetectedApps gets the detectedApps property value. The list of detected apps associated with a device.
// returns a []DetectedAppable when successful
func (m *DeviceManagement) GetDetectedApps()([]DetectedAppable) {
    return m.detectedApps
}
// GetDeviceCategories gets the deviceCategories property value. The list of device categories with the tenant.
// returns a []DeviceCategoryable when successful
func (m *DeviceManagement) GetDeviceCategories()([]DeviceCategoryable) {
    return m.deviceCategories
}
// GetDeviceCompliancePolicies gets the deviceCompliancePolicies property value. The device compliance policies.
// returns a []DeviceCompliancePolicyable when successful
func (m *DeviceManagement) GetDeviceCompliancePolicies()([]DeviceCompliancePolicyable) {
    return m.deviceCompliancePolicies
}
// GetDeviceCompliancePolicyDeviceStateSummary gets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
// returns a DeviceCompliancePolicyDeviceStateSummaryable when successful
func (m *DeviceManagement) GetDeviceCompliancePolicyDeviceStateSummary()(DeviceCompliancePolicyDeviceStateSummaryable) {
    return m.deviceCompliancePolicyDeviceStateSummary
}
// GetDeviceCompliancePolicySettingStateSummaries gets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
// returns a []DeviceCompliancePolicySettingStateSummaryable when successful
func (m *DeviceManagement) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable) {
    return m.deviceCompliancePolicySettingStateSummaries
}
// GetDeviceConfigurationDeviceStateSummaries gets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
// returns a DeviceConfigurationDeviceStateSummaryable when successful
func (m *DeviceManagement) GetDeviceConfigurationDeviceStateSummaries()(DeviceConfigurationDeviceStateSummaryable) {
    return m.deviceConfigurationDeviceStateSummaries
}
// GetDeviceConfigurations gets the deviceConfigurations property value. The device configurations.
// returns a []DeviceConfigurationable when successful
func (m *DeviceManagement) GetDeviceConfigurations()([]DeviceConfigurationable) {
    return m.deviceConfigurations
}
// GetDeviceEnrollmentConfigurations gets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
// returns a []DeviceEnrollmentConfigurationable when successful
func (m *DeviceManagement) GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfigurationable) {
    return m.deviceEnrollmentConfigurations
}
// GetDeviceManagementPartners gets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
// returns a []DeviceManagementPartnerable when successful
func (m *DeviceManagement) GetDeviceManagementPartners()([]DeviceManagementPartnerable) {
    return m.deviceManagementPartners
}
// GetDeviceProtectionOverview gets the deviceProtectionOverview property value. Device protection overview.
// returns a DeviceProtectionOverviewable when successful
func (m *DeviceManagement) GetDeviceProtectionOverview()(DeviceProtectionOverviewable) {
    return m.deviceProtectionOverview
}
// GetExchangeConnectors gets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
// returns a []DeviceManagementExchangeConnectorable when successful
func (m *DeviceManagement) GetExchangeConnectors()([]DeviceManagementExchangeConnectorable) {
    return m.exchangeConnectors
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applePushNotificationCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplePushNotificationCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplePushNotificationCertificate(val.(ApplePushNotificationCertificateable))
        }
        return nil
    }
    res["auditEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuditEventable)
                }
            }
            m.SetAuditEvents(res)
        }
        return nil
    }
    res["complianceManagementPartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceManagementPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceManagementPartnerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ComplianceManagementPartnerable)
                }
            }
            m.SetComplianceManagementPartners(res)
        }
        return nil
    }
    res["conditionalAccessSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesConditionalAccessSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessSettings(val.(OnPremisesConditionalAccessSettingsable))
        }
        return nil
    }
    res["detectedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDetectedAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DetectedAppable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DetectedAppable)
                }
            }
            m.SetDetectedApps(res)
        }
        return nil
    }
    res["deviceCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceCategoryable)
                }
            }
            m.SetDeviceCategories(res)
        }
        return nil
    }
    res["deviceCompliancePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceCompliancePolicyable)
                }
            }
            m.SetDeviceCompliancePolicies(res)
        }
        return nil
    }
    res["deviceCompliancePolicyDeviceStateSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCompliancePolicyDeviceStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePolicyDeviceStateSummary(val.(DeviceCompliancePolicyDeviceStateSummaryable))
        }
        return nil
    }
    res["deviceCompliancePolicySettingStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicySettingStateSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceCompliancePolicySettingStateSummaryable)
                }
            }
            m.SetDeviceCompliancePolicySettingStateSummaries(res)
        }
        return nil
    }
    res["deviceConfigurationDeviceStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationDeviceStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationDeviceStateSummaries(val.(DeviceConfigurationDeviceStateSummaryable))
        }
        return nil
    }
    res["deviceConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceConfigurationable)
                }
            }
            m.SetDeviceConfigurations(res)
        }
        return nil
    }
    res["deviceEnrollmentConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceEnrollmentConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceEnrollmentConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceEnrollmentConfigurationable)
                }
            }
            m.SetDeviceEnrollmentConfigurations(res)
        }
        return nil
    }
    res["deviceManagementPartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementPartnerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementPartnerable)
                }
            }
            m.SetDeviceManagementPartners(res)
        }
        return nil
    }
    res["deviceProtectionOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceProtectionOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceProtectionOverview(val.(DeviceProtectionOverviewable))
        }
        return nil
    }
    res["exchangeConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementExchangeConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementExchangeConnectorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementExchangeConnectorable)
                }
            }
            m.SetExchangeConnectors(res)
        }
        return nil
    }
    res["importedWindowsAutopilotDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImportedWindowsAutopilotDeviceIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ImportedWindowsAutopilotDeviceIdentityable)
                }
            }
            m.SetImportedWindowsAutopilotDeviceIdentities(res)
        }
        return nil
    }
    res["intuneAccountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneAccountId(val)
        }
        return nil
    }
    res["intuneBrand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntuneBrandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneBrand(val.(IntuneBrandable))
        }
        return nil
    }
    res["iosUpdateStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosUpdateDeviceStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosUpdateDeviceStatusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IosUpdateDeviceStatusable)
                }
            }
            m.SetIosUpdateStatuses(res)
        }
        return nil
    }
    res["managedDeviceOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedDeviceOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceOverview(val.(ManagedDeviceOverviewable))
        }
        return nil
    }
    res["managedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedDeviceable)
                }
            }
            m.SetManagedDevices(res)
        }
        return nil
    }
    res["mobileAppTroubleshootingEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppTroubleshootingEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppTroubleshootingEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MobileAppTroubleshootingEventable)
                }
            }
            m.SetMobileAppTroubleshootingEvents(res)
        }
        return nil
    }
    res["mobileThreatDefenseConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileThreatDefenseConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileThreatDefenseConnectorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MobileThreatDefenseConnectorable)
                }
            }
            m.SetMobileThreatDefenseConnectors(res)
        }
        return nil
    }
    res["notificationMessageTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNotificationMessageTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NotificationMessageTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NotificationMessageTemplateable)
                }
            }
            m.SetNotificationMessageTemplates(res)
        }
        return nil
    }
    res["remoteAssistancePartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRemoteAssistancePartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RemoteAssistancePartnerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RemoteAssistancePartnerable)
                }
            }
            m.SetRemoteAssistancePartners(res)
        }
        return nil
    }
    res["reports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementReportsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReports(val.(DeviceManagementReportsable))
        }
        return nil
    }
    res["resourceOperations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceOperationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ResourceOperationable)
                }
            }
            m.SetResourceOperations(res)
        }
        return nil
    }
    res["roleAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceAndAppManagementRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAndAppManagementRoleAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceAndAppManagementRoleAssignmentable)
                }
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RoleDefinitionable)
                }
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(DeviceManagementSettingsable))
        }
        return nil
    }
    res["softwareUpdateStatusSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSoftwareUpdateStatusSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdateStatusSummary(val.(SoftwareUpdateStatusSummaryable))
        }
        return nil
    }
    res["subscriptionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementSubscriptionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionState(val.(*DeviceManagementSubscriptionState))
        }
        return nil
    }
    res["telecomExpenseManagementPartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTelecomExpenseManagementPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TelecomExpenseManagementPartnerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TelecomExpenseManagementPartnerable)
                }
            }
            m.SetTelecomExpenseManagementPartners(res)
        }
        return nil
    }
    res["termsAndConditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTermsAndConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TermsAndConditionsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TermsAndConditionsable)
                }
            }
            m.SetTermsAndConditions(res)
        }
        return nil
    }
    res["troubleshootingEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTroubleshootingEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTroubleshootingEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementTroubleshootingEventable)
                }
            }
            m.SetTroubleshootingEvents(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthApplicationPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthApplicationPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthApplicationPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthDeviceModelPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthDevicePerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDevicePerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthDevicePerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthDevicePerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthDevicePerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthDevicePerformanceDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthOSVersionPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthOSVersionPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthOSVersionPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthOSVersionPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthOSVersionPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsAppHealthOverview(val.(UserExperienceAnalyticsCategoryable))
        }
        return nil
    }
    res["userExperienceAnalyticsBaselines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBaselineable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsBaselineable)
                }
            }
            m.SetUserExperienceAnalyticsBaselines(res)
        }
        return nil
    }
    res["userExperienceAnalyticsCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsCategoryable)
                }
            }
            m.SetUserExperienceAnalyticsCategories(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDevicePerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDevicePerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDevicePerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDevicePerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsDevicePerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceScores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceScoresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceScoresable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceScoresable)
                }
            }
            m.SetUserExperienceAnalyticsDeviceScores(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceStartupHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceStartupHistoryable)
                }
            }
            m.SetUserExperienceAnalyticsDeviceStartupHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupProcesses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupProcessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceStartupProcessable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceStartupProcessable)
                }
            }
            m.SetUserExperienceAnalyticsDeviceStartupProcesses(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupProcessPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceStartupProcessPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceStartupProcessPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsDeviceStartupProcessPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsMetricHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsMetricHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsMetricHistoryable)
                }
            }
            m.SetUserExperienceAnalyticsMetricHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsModelScores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsModelScoresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsModelScoresable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsModelScoresable)
                }
            }
            m.SetUserExperienceAnalyticsModelScores(res)
        }
        return nil
    }
    res["userExperienceAnalyticsOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsOverview(val.(UserExperienceAnalyticsOverviewable))
        }
        return nil
    }
    res["userExperienceAnalyticsScoreHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsScoreHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsScoreHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsScoreHistoryable)
                }
            }
            m.SetUserExperienceAnalyticsScoreHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsSettings(val.(UserExperienceAnalyticsSettingsable))
        }
        return nil
    }
    res["userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable))
        }
        return nil
    }
    res["userExperienceAnalyticsWorkFromAnywhereMetrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsWorkFromAnywhereMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsWorkFromAnywhereMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsWorkFromAnywhereMetricable)
                }
            }
            m.SetUserExperienceAnalyticsWorkFromAnywhereMetrics(res)
        }
        return nil
    }
    res["userExperienceAnalyticsWorkFromAnywhereModelPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance(res)
        }
        return nil
    }
    res["virtualEndpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualEndpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualEndpoint(val.(VirtualEndpointable))
        }
        return nil
    }
    res["windowsAutopilotDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsAutopilotDeviceIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsAutopilotDeviceIdentityable)
                }
            }
            m.SetWindowsAutopilotDeviceIdentities(res)
        }
        return nil
    }
    res["windowsInformationProtectionAppLearningSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionAppLearningSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionAppLearningSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsInformationProtectionAppLearningSummaryable)
                }
            }
            m.SetWindowsInformationProtectionAppLearningSummaries(res)
        }
        return nil
    }
    res["windowsInformationProtectionNetworkLearningSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionNetworkLearningSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionNetworkLearningSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsInformationProtectionNetworkLearningSummaryable)
                }
            }
            m.SetWindowsInformationProtectionNetworkLearningSummaries(res)
        }
        return nil
    }
    res["windowsMalwareInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsMalwareInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsMalwareInformationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsMalwareInformationable)
                }
            }
            m.SetWindowsMalwareInformation(res)
        }
        return nil
    }
    res["windowsMalwareOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsMalwareOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsMalwareOverview(val.(WindowsMalwareOverviewable))
        }
        return nil
    }
    return res
}
// GetImportedWindowsAutopilotDeviceIdentities gets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
// returns a []ImportedWindowsAutopilotDeviceIdentityable when successful
func (m *DeviceManagement) GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentityable) {
    return m.importedWindowsAutopilotDeviceIdentities
}
// GetIntuneAccountId gets the intuneAccountId property value. Intune Account Id for given tenant
// returns a *UUID when successful
func (m *DeviceManagement) GetIntuneAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.intuneAccountId
}
// GetIntuneBrand gets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
// returns a IntuneBrandable when successful
func (m *DeviceManagement) GetIntuneBrand()(IntuneBrandable) {
    return m.intuneBrand
}
// GetIosUpdateStatuses gets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
// returns a []IosUpdateDeviceStatusable when successful
func (m *DeviceManagement) GetIosUpdateStatuses()([]IosUpdateDeviceStatusable) {
    return m.iosUpdateStatuses
}
// GetManagedDeviceOverview gets the managedDeviceOverview property value. Device overview
// returns a ManagedDeviceOverviewable when successful
func (m *DeviceManagement) GetManagedDeviceOverview()(ManagedDeviceOverviewable) {
    return m.managedDeviceOverview
}
// GetManagedDevices gets the managedDevices property value. The list of managed devices.
// returns a []ManagedDeviceable when successful
func (m *DeviceManagement) GetManagedDevices()([]ManagedDeviceable) {
    return m.managedDevices
}
// GetMobileAppTroubleshootingEvents gets the mobileAppTroubleshootingEvents property value. The collection property of MobileAppTroubleshootingEvent.
// returns a []MobileAppTroubleshootingEventable when successful
func (m *DeviceManagement) GetMobileAppTroubleshootingEvents()([]MobileAppTroubleshootingEventable) {
    return m.mobileAppTroubleshootingEvents
}
// GetMobileThreatDefenseConnectors gets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
// returns a []MobileThreatDefenseConnectorable when successful
func (m *DeviceManagement) GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnectorable) {
    return m.mobileThreatDefenseConnectors
}
// GetNotificationMessageTemplates gets the notificationMessageTemplates property value. The Notification Message Templates.
// returns a []NotificationMessageTemplateable when successful
func (m *DeviceManagement) GetNotificationMessageTemplates()([]NotificationMessageTemplateable) {
    return m.notificationMessageTemplates
}
// GetRemoteAssistancePartners gets the remoteAssistancePartners property value. The remote assist partners.
// returns a []RemoteAssistancePartnerable when successful
func (m *DeviceManagement) GetRemoteAssistancePartners()([]RemoteAssistancePartnerable) {
    return m.remoteAssistancePartners
}
// GetReports gets the reports property value. Reports singleton
// returns a DeviceManagementReportsable when successful
func (m *DeviceManagement) GetReports()(DeviceManagementReportsable) {
    return m.reports
}
// GetResourceOperations gets the resourceOperations property value. The Resource Operations.
// returns a []ResourceOperationable when successful
func (m *DeviceManagement) GetResourceOperations()([]ResourceOperationable) {
    return m.resourceOperations
}
// GetRoleAssignments gets the roleAssignments property value. The Role Assignments.
// returns a []DeviceAndAppManagementRoleAssignmentable when successful
func (m *DeviceManagement) GetRoleAssignments()([]DeviceAndAppManagementRoleAssignmentable) {
    return m.roleAssignments
}
// GetRoleDefinitions gets the roleDefinitions property value. The Role Definitions.
// returns a []RoleDefinitionable when successful
func (m *DeviceManagement) GetRoleDefinitions()([]RoleDefinitionable) {
    return m.roleDefinitions
}
// GetSettings gets the settings property value. Account level settings.
// returns a DeviceManagementSettingsable when successful
func (m *DeviceManagement) GetSettings()(DeviceManagementSettingsable) {
    return m.settings
}
// GetSoftwareUpdateStatusSummary gets the softwareUpdateStatusSummary property value. The software update status summary.
// returns a SoftwareUpdateStatusSummaryable when successful
func (m *DeviceManagement) GetSoftwareUpdateStatusSummary()(SoftwareUpdateStatusSummaryable) {
    return m.softwareUpdateStatusSummary
}
// GetSubscriptionState gets the subscriptionState property value. Tenant mobile device management subscription state.
// returns a *DeviceManagementSubscriptionState when successful
func (m *DeviceManagement) GetSubscriptionState()(*DeviceManagementSubscriptionState) {
    return m.subscriptionState
}
// GetTelecomExpenseManagementPartners gets the telecomExpenseManagementPartners property value. The telecom expense management partners.
// returns a []TelecomExpenseManagementPartnerable when successful
func (m *DeviceManagement) GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartnerable) {
    return m.telecomExpenseManagementPartners
}
// GetTermsAndConditions gets the termsAndConditions property value. The terms and conditions associated with device management of the company.
// returns a []TermsAndConditionsable when successful
func (m *DeviceManagement) GetTermsAndConditions()([]TermsAndConditionsable) {
    return m.termsAndConditions
}
// GetTroubleshootingEvents gets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
// returns a []DeviceManagementTroubleshootingEventable when successful
func (m *DeviceManagement) GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable) {
    return m.troubleshootingEvents
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformance gets the userExperienceAnalyticsAppHealthApplicationPerformance property value. User experience analytics appHealth Application Performance
// returns a []UserExperienceAnalyticsAppHealthApplicationPerformanceable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformance()([]UserExperienceAnalyticsAppHealthApplicationPerformanceable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformance
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property value. User experience analytics appHealth Application Performance by App Version details
// returns a []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property value. User experience analytics appHealth Application Performance by App Version Device Id
// returns a []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion gets the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property value. User experience analytics appHealth Application Performance by OS Version
// returns a []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion
}
// GetUserExperienceAnalyticsAppHealthDeviceModelPerformance gets the userExperienceAnalyticsAppHealthDeviceModelPerformance property value. User experience analytics appHealth Model Performance
// returns a []UserExperienceAnalyticsAppHealthDeviceModelPerformanceable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()([]UserExperienceAnalyticsAppHealthDeviceModelPerformanceable) {
    return m.userExperienceAnalyticsAppHealthDeviceModelPerformance
}
// GetUserExperienceAnalyticsAppHealthDevicePerformance gets the userExperienceAnalyticsAppHealthDevicePerformance property value. User experience analytics appHealth Device Performance
// returns a []UserExperienceAnalyticsAppHealthDevicePerformanceable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformance()([]UserExperienceAnalyticsAppHealthDevicePerformanceable) {
    return m.userExperienceAnalyticsAppHealthDevicePerformance
}
// GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails gets the userExperienceAnalyticsAppHealthDevicePerformanceDetails property value. User experience analytics device performance details
// returns a []UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()([]UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable) {
    return m.userExperienceAnalyticsAppHealthDevicePerformanceDetails
}
// GetUserExperienceAnalyticsAppHealthOSVersionPerformance gets the userExperienceAnalyticsAppHealthOSVersionPerformance property value. User experience analytics appHealth OS version Performance
// returns a []UserExperienceAnalyticsAppHealthOSVersionPerformanceable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOSVersionPerformance()([]UserExperienceAnalyticsAppHealthOSVersionPerformanceable) {
    return m.userExperienceAnalyticsAppHealthOSVersionPerformance
}
// GetUserExperienceAnalyticsAppHealthOverview gets the userExperienceAnalyticsAppHealthOverview property value. User experience analytics appHealth overview
// returns a UserExperienceAnalyticsCategoryable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOverview()(UserExperienceAnalyticsCategoryable) {
    return m.userExperienceAnalyticsAppHealthOverview
}
// GetUserExperienceAnalyticsBaselines gets the userExperienceAnalyticsBaselines property value. User experience analytics baselines
// returns a []UserExperienceAnalyticsBaselineable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsBaselines()([]UserExperienceAnalyticsBaselineable) {
    return m.userExperienceAnalyticsBaselines
}
// GetUserExperienceAnalyticsCategories gets the userExperienceAnalyticsCategories property value. User experience analytics categories
// returns a []UserExperienceAnalyticsCategoryable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsCategories()([]UserExperienceAnalyticsCategoryable) {
    return m.userExperienceAnalyticsCategories
}
// GetUserExperienceAnalyticsDevicePerformance gets the userExperienceAnalyticsDevicePerformance property value. User experience analytics device performance
// returns a []UserExperienceAnalyticsDevicePerformanceable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsDevicePerformance()([]UserExperienceAnalyticsDevicePerformanceable) {
    return m.userExperienceAnalyticsDevicePerformance
}
// GetUserExperienceAnalyticsDeviceScores gets the userExperienceAnalyticsDeviceScores property value. User experience analytics device scores
// returns a []UserExperienceAnalyticsDeviceScoresable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScores()([]UserExperienceAnalyticsDeviceScoresable) {
    return m.userExperienceAnalyticsDeviceScores
}
// GetUserExperienceAnalyticsDeviceStartupHistory gets the userExperienceAnalyticsDeviceStartupHistory property value. User experience analytics device Startup History
// returns a []UserExperienceAnalyticsDeviceStartupHistoryable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupHistory()([]UserExperienceAnalyticsDeviceStartupHistoryable) {
    return m.userExperienceAnalyticsDeviceStartupHistory
}
// GetUserExperienceAnalyticsDeviceStartupProcesses gets the userExperienceAnalyticsDeviceStartupProcesses property value. User experience analytics device Startup Processes
// returns a []UserExperienceAnalyticsDeviceStartupProcessable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcesses()([]UserExperienceAnalyticsDeviceStartupProcessable) {
    return m.userExperienceAnalyticsDeviceStartupProcesses
}
// GetUserExperienceAnalyticsDeviceStartupProcessPerformance gets the userExperienceAnalyticsDeviceStartupProcessPerformance property value. User experience analytics device Startup Process Performance
// returns a []UserExperienceAnalyticsDeviceStartupProcessPerformanceable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcessPerformance()([]UserExperienceAnalyticsDeviceStartupProcessPerformanceable) {
    return m.userExperienceAnalyticsDeviceStartupProcessPerformance
}
// GetUserExperienceAnalyticsMetricHistory gets the userExperienceAnalyticsMetricHistory property value. User experience analytics metric history
// returns a []UserExperienceAnalyticsMetricHistoryable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsMetricHistory()([]UserExperienceAnalyticsMetricHistoryable) {
    return m.userExperienceAnalyticsMetricHistory
}
// GetUserExperienceAnalyticsModelScores gets the userExperienceAnalyticsModelScores property value. User experience analytics model scores
// returns a []UserExperienceAnalyticsModelScoresable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsModelScores()([]UserExperienceAnalyticsModelScoresable) {
    return m.userExperienceAnalyticsModelScores
}
// GetUserExperienceAnalyticsOverview gets the userExperienceAnalyticsOverview property value. User experience analytics overview
// returns a UserExperienceAnalyticsOverviewable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsOverview()(UserExperienceAnalyticsOverviewable) {
    return m.userExperienceAnalyticsOverview
}
// GetUserExperienceAnalyticsScoreHistory gets the userExperienceAnalyticsScoreHistory property value. User experience analytics device Startup Score History
// returns a []UserExperienceAnalyticsScoreHistoryable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsScoreHistory()([]UserExperienceAnalyticsScoreHistoryable) {
    return m.userExperienceAnalyticsScoreHistory
}
// GetUserExperienceAnalyticsSettings gets the userExperienceAnalyticsSettings property value. User experience analytics device settings
// returns a UserExperienceAnalyticsSettingsable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsSettings()(UserExperienceAnalyticsSettingsable) {
    return m.userExperienceAnalyticsSettings
}
// GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric gets the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property value. User experience analytics work from anywhere hardware readiness metrics.
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable) {
    return m.userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric
}
// GetUserExperienceAnalyticsWorkFromAnywhereMetrics gets the userExperienceAnalyticsWorkFromAnywhereMetrics property value. User experience analytics work from anywhere metrics.
// returns a []UserExperienceAnalyticsWorkFromAnywhereMetricable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereMetrics()([]UserExperienceAnalyticsWorkFromAnywhereMetricable) {
    return m.userExperienceAnalyticsWorkFromAnywhereMetrics
}
// GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance gets the userExperienceAnalyticsWorkFromAnywhereModelPerformance property value. The user experience analytics work from anywhere model performance
// returns a []UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable when successful
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance()([]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable) {
    return m.userExperienceAnalyticsWorkFromAnywhereModelPerformance
}
// GetVirtualEndpoint gets the virtualEndpoint property value. Virtual endpoint
// returns a VirtualEndpointable when successful
func (m *DeviceManagement) GetVirtualEndpoint()(VirtualEndpointable) {
    return m.virtualEndpoint
}
// GetWindowsAutopilotDeviceIdentities gets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
// returns a []WindowsAutopilotDeviceIdentityable when successful
func (m *DeviceManagement) GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentityable) {
    return m.windowsAutopilotDeviceIdentities
}
// GetWindowsInformationProtectionAppLearningSummaries gets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
// returns a []WindowsInformationProtectionAppLearningSummaryable when successful
func (m *DeviceManagement) GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummaryable) {
    return m.windowsInformationProtectionAppLearningSummaries
}
// GetWindowsInformationProtectionNetworkLearningSummaries gets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
// returns a []WindowsInformationProtectionNetworkLearningSummaryable when successful
func (m *DeviceManagement) GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummaryable) {
    return m.windowsInformationProtectionNetworkLearningSummaries
}
// GetWindowsMalwareInformation gets the windowsMalwareInformation property value. The list of affected malware in the tenant.
// returns a []WindowsMalwareInformationable when successful
func (m *DeviceManagement) GetWindowsMalwareInformation()([]WindowsMalwareInformationable) {
    return m.windowsMalwareInformation
}
// GetWindowsMalwareOverview gets the windowsMalwareOverview property value. Malware overview for windows devices.
// returns a WindowsMalwareOverviewable when successful
func (m *DeviceManagement) GetWindowsMalwareOverview()(WindowsMalwareOverviewable) {
    return m.windowsMalwareOverview
}
// Serialize serializes information the current object
func (m *DeviceManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("applePushNotificationCertificate", m.GetApplePushNotificationCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetAuditEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuditEvents()))
        for i, v := range m.GetAuditEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceManagementPartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceManagementPartners()))
        for i, v := range m.GetComplianceManagementPartners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("complianceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("conditionalAccessSettings", m.GetConditionalAccessSettings())
        if err != nil {
            return err
        }
    }
    if m.GetDetectedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetectedApps()))
        for i, v := range m.GetDetectedApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("detectedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCategories()))
        for i, v := range m.GetDeviceCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCompliancePolicies()))
        for i, v := range m.GetDeviceCompliancePolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceCompliancePolicyDeviceStateSummary", m.GetDeviceCompliancePolicyDeviceStateSummary())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicySettingStateSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCompliancePolicySettingStateSummaries()))
        for i, v := range m.GetDeviceCompliancePolicySettingStateSummaries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicySettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceConfigurationDeviceStateSummaries", m.GetDeviceConfigurationDeviceStateSummaries())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurations()))
        for i, v := range m.GetDeviceConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceEnrollmentConfigurations()))
        for i, v := range m.GetDeviceEnrollmentConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceEnrollmentConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementPartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceManagementPartners()))
        for i, v := range m.GetDeviceManagementPartners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceProtectionOverview", m.GetDeviceProtectionOverview())
        if err != nil {
            return err
        }
    }
    if m.GetExchangeConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExchangeConnectors()))
        for i, v := range m.GetExchangeConnectors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("exchangeConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportedWindowsAutopilotDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedWindowsAutopilotDeviceIdentities()))
        for i, v := range m.GetImportedWindowsAutopilotDeviceIdentities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("importedWindowsAutopilotDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("intuneAccountId", m.GetIntuneAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("intuneBrand", m.GetIntuneBrand())
        if err != nil {
            return err
        }
    }
    if m.GetIosUpdateStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIosUpdateStatuses()))
        for i, v := range m.GetIosUpdateStatuses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("iosUpdateStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managedDeviceOverview", m.GetManagedDeviceOverview())
        if err != nil {
            return err
        }
    }
    if m.GetManagedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDevices()))
        for i, v := range m.GetManagedDevices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("managedDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppTroubleshootingEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppTroubleshootingEvents()))
        for i, v := range m.GetMobileAppTroubleshootingEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppTroubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileThreatDefenseConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileThreatDefenseConnectors()))
        for i, v := range m.GetMobileThreatDefenseConnectors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mobileThreatDefenseConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotificationMessageTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotificationMessageTemplates()))
        for i, v := range m.GetNotificationMessageTemplates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("notificationMessageTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoteAssistancePartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemoteAssistancePartners()))
        for i, v := range m.GetRemoteAssistancePartners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("remoteAssistancePartners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reports", m.GetReports())
        if err != nil {
            return err
        }
    }
    if m.GetResourceOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceOperations()))
        for i, v := range m.GetResourceOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceOperations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("softwareUpdateStatusSummary", m.GetSoftwareUpdateStatusSummary())
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptionState() != nil {
        cast := (*m.GetSubscriptionState()).String()
        err = writer.WriteStringValue("subscriptionState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTelecomExpenseManagementPartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTelecomExpenseManagementPartners()))
        for i, v := range m.GetTelecomExpenseManagementPartners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("telecomExpenseManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTermsAndConditions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTermsAndConditions()))
        for i, v := range m.GetTermsAndConditions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("termsAndConditions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTroubleshootingEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTroubleshootingEvents()))
        for i, v := range m.GetTroubleshootingEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("troubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDeviceModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDevicePerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDevicePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDevicePerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDevicePerformanceDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthOSVersionPerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsAppHealthOverview", m.GetUserExperienceAnalyticsAppHealthOverview())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBaselines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBaselines()))
        for i, v := range m.GetUserExperienceAnalyticsBaselines() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBaselines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsCategories()))
        for i, v := range m.GetUserExperienceAnalyticsCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDevicePerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDevicePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsDevicePerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceScores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceScores()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceScores() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceScores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupHistory()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupProcesses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupProcesses()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupProcesses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupProcesses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupProcessPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsMetricHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsMetricHistory()))
        for i, v := range m.GetUserExperienceAnalyticsMetricHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsMetricHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsModelScores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsModelScores()))
        for i, v := range m.GetUserExperienceAnalyticsModelScores() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsModelScores", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsOverview", m.GetUserExperienceAnalyticsOverview())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsScoreHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsScoreHistory()))
        for i, v := range m.GetUserExperienceAnalyticsScoreHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsScoreHistory", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsSettings", m.GetUserExperienceAnalyticsSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric", m.GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics()))
        for i, v := range m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsWorkFromAnywhereMetrics", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsWorkFromAnywhereModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("virtualEndpoint", m.GetVirtualEndpoint())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsAutopilotDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsAutopilotDeviceIdentities()))
        for i, v := range m.GetWindowsAutopilotDeviceIdentities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsAutopilotDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionAppLearningSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionAppLearningSummaries()))
        for i, v := range m.GetWindowsInformationProtectionAppLearningSummaries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionAppLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionNetworkLearningSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionNetworkLearningSummaries()))
        for i, v := range m.GetWindowsInformationProtectionNetworkLearningSummaries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionNetworkLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsMalwareInformation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsMalwareInformation()))
        for i, v := range m.GetWindowsMalwareInformation() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsMalwareInformation", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsMalwareOverview", m.GetWindowsMalwareOverview())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplePushNotificationCertificate sets the applePushNotificationCertificate property value. Apple push notification certificate.
func (m *DeviceManagement) SetApplePushNotificationCertificate(value ApplePushNotificationCertificateable)() {
    m.applePushNotificationCertificate = value
}
// SetAuditEvents sets the auditEvents property value. The Audit Events
func (m *DeviceManagement) SetAuditEvents(value []AuditEventable)() {
    m.auditEvents = value
}
// SetComplianceManagementPartners sets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
func (m *DeviceManagement) SetComplianceManagementPartners(value []ComplianceManagementPartnerable)() {
    m.complianceManagementPartners = value
}
// SetConditionalAccessSettings sets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagement) SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)() {
    m.conditionalAccessSettings = value
}
// SetDetectedApps sets the detectedApps property value. The list of detected apps associated with a device.
func (m *DeviceManagement) SetDetectedApps(value []DetectedAppable)() {
    m.detectedApps = value
}
// SetDeviceCategories sets the deviceCategories property value. The list of device categories with the tenant.
func (m *DeviceManagement) SetDeviceCategories(value []DeviceCategoryable)() {
    m.deviceCategories = value
}
// SetDeviceCompliancePolicies sets the deviceCompliancePolicies property value. The device compliance policies.
func (m *DeviceManagement) SetDeviceCompliancePolicies(value []DeviceCompliancePolicyable)() {
    m.deviceCompliancePolicies = value
}
// SetDeviceCompliancePolicyDeviceStateSummary sets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
func (m *DeviceManagement) SetDeviceCompliancePolicyDeviceStateSummary(value DeviceCompliancePolicyDeviceStateSummaryable)() {
    m.deviceCompliancePolicyDeviceStateSummary = value
}
// SetDeviceCompliancePolicySettingStateSummaries sets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
func (m *DeviceManagement) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)() {
    m.deviceCompliancePolicySettingStateSummaries = value
}
// SetDeviceConfigurationDeviceStateSummaries sets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
func (m *DeviceManagement) SetDeviceConfigurationDeviceStateSummaries(value DeviceConfigurationDeviceStateSummaryable)() {
    m.deviceConfigurationDeviceStateSummaries = value
}
// SetDeviceConfigurations sets the deviceConfigurations property value. The device configurations.
func (m *DeviceManagement) SetDeviceConfigurations(value []DeviceConfigurationable)() {
    m.deviceConfigurations = value
}
// SetDeviceEnrollmentConfigurations sets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
func (m *DeviceManagement) SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfigurationable)() {
    m.deviceEnrollmentConfigurations = value
}
// SetDeviceManagementPartners sets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
func (m *DeviceManagement) SetDeviceManagementPartners(value []DeviceManagementPartnerable)() {
    m.deviceManagementPartners = value
}
// SetDeviceProtectionOverview sets the deviceProtectionOverview property value. Device protection overview.
func (m *DeviceManagement) SetDeviceProtectionOverview(value DeviceProtectionOverviewable)() {
    m.deviceProtectionOverview = value
}
// SetExchangeConnectors sets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
func (m *DeviceManagement) SetExchangeConnectors(value []DeviceManagementExchangeConnectorable)() {
    m.exchangeConnectors = value
}
// SetImportedWindowsAutopilotDeviceIdentities sets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
func (m *DeviceManagement) SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentityable)() {
    m.importedWindowsAutopilotDeviceIdentities = value
}
// SetIntuneAccountId sets the intuneAccountId property value. Intune Account Id for given tenant
func (m *DeviceManagement) SetIntuneAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.intuneAccountId = value
}
// SetIntuneBrand sets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
func (m *DeviceManagement) SetIntuneBrand(value IntuneBrandable)() {
    m.intuneBrand = value
}
// SetIosUpdateStatuses sets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
func (m *DeviceManagement) SetIosUpdateStatuses(value []IosUpdateDeviceStatusable)() {
    m.iosUpdateStatuses = value
}
// SetManagedDeviceOverview sets the managedDeviceOverview property value. Device overview
func (m *DeviceManagement) SetManagedDeviceOverview(value ManagedDeviceOverviewable)() {
    m.managedDeviceOverview = value
}
// SetManagedDevices sets the managedDevices property value. The list of managed devices.
func (m *DeviceManagement) SetManagedDevices(value []ManagedDeviceable)() {
    m.managedDevices = value
}
// SetMobileAppTroubleshootingEvents sets the mobileAppTroubleshootingEvents property value. The collection property of MobileAppTroubleshootingEvent.
func (m *DeviceManagement) SetMobileAppTroubleshootingEvents(value []MobileAppTroubleshootingEventable)() {
    m.mobileAppTroubleshootingEvents = value
}
// SetMobileThreatDefenseConnectors sets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
func (m *DeviceManagement) SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnectorable)() {
    m.mobileThreatDefenseConnectors = value
}
// SetNotificationMessageTemplates sets the notificationMessageTemplates property value. The Notification Message Templates.
func (m *DeviceManagement) SetNotificationMessageTemplates(value []NotificationMessageTemplateable)() {
    m.notificationMessageTemplates = value
}
// SetRemoteAssistancePartners sets the remoteAssistancePartners property value. The remote assist partners.
func (m *DeviceManagement) SetRemoteAssistancePartners(value []RemoteAssistancePartnerable)() {
    m.remoteAssistancePartners = value
}
// SetReports sets the reports property value. Reports singleton
func (m *DeviceManagement) SetReports(value DeviceManagementReportsable)() {
    m.reports = value
}
// SetResourceOperations sets the resourceOperations property value. The Resource Operations.
func (m *DeviceManagement) SetResourceOperations(value []ResourceOperationable)() {
    m.resourceOperations = value
}
// SetRoleAssignments sets the roleAssignments property value. The Role Assignments.
func (m *DeviceManagement) SetRoleAssignments(value []DeviceAndAppManagementRoleAssignmentable)() {
    m.roleAssignments = value
}
// SetRoleDefinitions sets the roleDefinitions property value. The Role Definitions.
func (m *DeviceManagement) SetRoleDefinitions(value []RoleDefinitionable)() {
    m.roleDefinitions = value
}
// SetSettings sets the settings property value. Account level settings.
func (m *DeviceManagement) SetSettings(value DeviceManagementSettingsable)() {
    m.settings = value
}
// SetSoftwareUpdateStatusSummary sets the softwareUpdateStatusSummary property value. The software update status summary.
func (m *DeviceManagement) SetSoftwareUpdateStatusSummary(value SoftwareUpdateStatusSummaryable)() {
    m.softwareUpdateStatusSummary = value
}
// SetSubscriptionState sets the subscriptionState property value. Tenant mobile device management subscription state.
func (m *DeviceManagement) SetSubscriptionState(value *DeviceManagementSubscriptionState)() {
    m.subscriptionState = value
}
// SetTelecomExpenseManagementPartners sets the telecomExpenseManagementPartners property value. The telecom expense management partners.
func (m *DeviceManagement) SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartnerable)() {
    m.telecomExpenseManagementPartners = value
}
// SetTermsAndConditions sets the termsAndConditions property value. The terms and conditions associated with device management of the company.
func (m *DeviceManagement) SetTermsAndConditions(value []TermsAndConditionsable)() {
    m.termsAndConditions = value
}
// SetTroubleshootingEvents sets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
func (m *DeviceManagement) SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)() {
    m.troubleshootingEvents = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformance sets the userExperienceAnalyticsAppHealthApplicationPerformance property value. User experience analytics appHealth Application Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformance(value []UserExperienceAnalyticsAppHealthApplicationPerformanceable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformance = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property value. User experience analytics appHealth Application Performance by App Version details
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property value. User experience analytics appHealth Application Performance by App Version Device Id
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion sets the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property value. User experience analytics appHealth Application Performance by OS Version
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion = value
}
// SetUserExperienceAnalyticsAppHealthDeviceModelPerformance sets the userExperienceAnalyticsAppHealthDeviceModelPerformance property value. User experience analytics appHealth Model Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(value []UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)() {
    m.userExperienceAnalyticsAppHealthDeviceModelPerformance = value
}
// SetUserExperienceAnalyticsAppHealthDevicePerformance sets the userExperienceAnalyticsAppHealthDevicePerformance property value. User experience analytics appHealth Device Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformance(value []UserExperienceAnalyticsAppHealthDevicePerformanceable)() {
    m.userExperienceAnalyticsAppHealthDevicePerformance = value
}
// SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails sets the userExperienceAnalyticsAppHealthDevicePerformanceDetails property value. User experience analytics device performance details
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(value []UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)() {
    m.userExperienceAnalyticsAppHealthDevicePerformanceDetails = value
}
// SetUserExperienceAnalyticsAppHealthOSVersionPerformance sets the userExperienceAnalyticsAppHealthOSVersionPerformance property value. User experience analytics appHealth OS version Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOSVersionPerformance(value []UserExperienceAnalyticsAppHealthOSVersionPerformanceable)() {
    m.userExperienceAnalyticsAppHealthOSVersionPerformance = value
}
// SetUserExperienceAnalyticsAppHealthOverview sets the userExperienceAnalyticsAppHealthOverview property value. User experience analytics appHealth overview
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOverview(value UserExperienceAnalyticsCategoryable)() {
    m.userExperienceAnalyticsAppHealthOverview = value
}
// SetUserExperienceAnalyticsBaselines sets the userExperienceAnalyticsBaselines property value. User experience analytics baselines
func (m *DeviceManagement) SetUserExperienceAnalyticsBaselines(value []UserExperienceAnalyticsBaselineable)() {
    m.userExperienceAnalyticsBaselines = value
}
// SetUserExperienceAnalyticsCategories sets the userExperienceAnalyticsCategories property value. User experience analytics categories
func (m *DeviceManagement) SetUserExperienceAnalyticsCategories(value []UserExperienceAnalyticsCategoryable)() {
    m.userExperienceAnalyticsCategories = value
}
// SetUserExperienceAnalyticsDevicePerformance sets the userExperienceAnalyticsDevicePerformance property value. User experience analytics device performance
func (m *DeviceManagement) SetUserExperienceAnalyticsDevicePerformance(value []UserExperienceAnalyticsDevicePerformanceable)() {
    m.userExperienceAnalyticsDevicePerformance = value
}
// SetUserExperienceAnalyticsDeviceScores sets the userExperienceAnalyticsDeviceScores property value. User experience analytics device scores
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScores(value []UserExperienceAnalyticsDeviceScoresable)() {
    m.userExperienceAnalyticsDeviceScores = value
}
// SetUserExperienceAnalyticsDeviceStartupHistory sets the userExperienceAnalyticsDeviceStartupHistory property value. User experience analytics device Startup History
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupHistory(value []UserExperienceAnalyticsDeviceStartupHistoryable)() {
    m.userExperienceAnalyticsDeviceStartupHistory = value
}
// SetUserExperienceAnalyticsDeviceStartupProcesses sets the userExperienceAnalyticsDeviceStartupProcesses property value. User experience analytics device Startup Processes
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcesses(value []UserExperienceAnalyticsDeviceStartupProcessable)() {
    m.userExperienceAnalyticsDeviceStartupProcesses = value
}
// SetUserExperienceAnalyticsDeviceStartupProcessPerformance sets the userExperienceAnalyticsDeviceStartupProcessPerformance property value. User experience analytics device Startup Process Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcessPerformance(value []UserExperienceAnalyticsDeviceStartupProcessPerformanceable)() {
    m.userExperienceAnalyticsDeviceStartupProcessPerformance = value
}
// SetUserExperienceAnalyticsMetricHistory sets the userExperienceAnalyticsMetricHistory property value. User experience analytics metric history
func (m *DeviceManagement) SetUserExperienceAnalyticsMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)() {
    m.userExperienceAnalyticsMetricHistory = value
}
// SetUserExperienceAnalyticsModelScores sets the userExperienceAnalyticsModelScores property value. User experience analytics model scores
func (m *DeviceManagement) SetUserExperienceAnalyticsModelScores(value []UserExperienceAnalyticsModelScoresable)() {
    m.userExperienceAnalyticsModelScores = value
}
// SetUserExperienceAnalyticsOverview sets the userExperienceAnalyticsOverview property value. User experience analytics overview
func (m *DeviceManagement) SetUserExperienceAnalyticsOverview(value UserExperienceAnalyticsOverviewable)() {
    m.userExperienceAnalyticsOverview = value
}
// SetUserExperienceAnalyticsScoreHistory sets the userExperienceAnalyticsScoreHistory property value. User experience analytics device Startup Score History
func (m *DeviceManagement) SetUserExperienceAnalyticsScoreHistory(value []UserExperienceAnalyticsScoreHistoryable)() {
    m.userExperienceAnalyticsScoreHistory = value
}
// SetUserExperienceAnalyticsSettings sets the userExperienceAnalyticsSettings property value. User experience analytics device settings
func (m *DeviceManagement) SetUserExperienceAnalyticsSettings(value UserExperienceAnalyticsSettingsable)() {
    m.userExperienceAnalyticsSettings = value
}
// SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric sets the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property value. User experience analytics work from anywhere hardware readiness metrics.
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable)() {
    m.userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric = value
}
// SetUserExperienceAnalyticsWorkFromAnywhereMetrics sets the userExperienceAnalyticsWorkFromAnywhereMetrics property value. User experience analytics work from anywhere metrics.
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereMetrics(value []UserExperienceAnalyticsWorkFromAnywhereMetricable)() {
    m.userExperienceAnalyticsWorkFromAnywhereMetrics = value
}
// SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance sets the userExperienceAnalyticsWorkFromAnywhereModelPerformance property value. The user experience analytics work from anywhere model performance
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance(value []UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)() {
    m.userExperienceAnalyticsWorkFromAnywhereModelPerformance = value
}
// SetVirtualEndpoint sets the virtualEndpoint property value. Virtual endpoint
func (m *DeviceManagement) SetVirtualEndpoint(value VirtualEndpointable)() {
    m.virtualEndpoint = value
}
// SetWindowsAutopilotDeviceIdentities sets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
func (m *DeviceManagement) SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentityable)() {
    m.windowsAutopilotDeviceIdentities = value
}
// SetWindowsInformationProtectionAppLearningSummaries sets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
func (m *DeviceManagement) SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummaryable)() {
    m.windowsInformationProtectionAppLearningSummaries = value
}
// SetWindowsInformationProtectionNetworkLearningSummaries sets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
func (m *DeviceManagement) SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummaryable)() {
    m.windowsInformationProtectionNetworkLearningSummaries = value
}
// SetWindowsMalwareInformation sets the windowsMalwareInformation property value. The list of affected malware in the tenant.
func (m *DeviceManagement) SetWindowsMalwareInformation(value []WindowsMalwareInformationable)() {
    m.windowsMalwareInformation = value
}
// SetWindowsMalwareOverview sets the windowsMalwareOverview property value. Malware overview for windows devices.
func (m *DeviceManagement) SetWindowsMalwareOverview(value WindowsMalwareOverviewable)() {
    m.windowsMalwareOverview = value
}
type DeviceManagementable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplePushNotificationCertificate()(ApplePushNotificationCertificateable)
    GetAuditEvents()([]AuditEventable)
    GetComplianceManagementPartners()([]ComplianceManagementPartnerable)
    GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable)
    GetDetectedApps()([]DetectedAppable)
    GetDeviceCategories()([]DeviceCategoryable)
    GetDeviceCompliancePolicies()([]DeviceCompliancePolicyable)
    GetDeviceCompliancePolicyDeviceStateSummary()(DeviceCompliancePolicyDeviceStateSummaryable)
    GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable)
    GetDeviceConfigurationDeviceStateSummaries()(DeviceConfigurationDeviceStateSummaryable)
    GetDeviceConfigurations()([]DeviceConfigurationable)
    GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfigurationable)
    GetDeviceManagementPartners()([]DeviceManagementPartnerable)
    GetDeviceProtectionOverview()(DeviceProtectionOverviewable)
    GetExchangeConnectors()([]DeviceManagementExchangeConnectorable)
    GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentityable)
    GetIntuneAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetIntuneBrand()(IntuneBrandable)
    GetIosUpdateStatuses()([]IosUpdateDeviceStatusable)
    GetManagedDeviceOverview()(ManagedDeviceOverviewable)
    GetManagedDevices()([]ManagedDeviceable)
    GetMobileAppTroubleshootingEvents()([]MobileAppTroubleshootingEventable)
    GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnectorable)
    GetNotificationMessageTemplates()([]NotificationMessageTemplateable)
    GetRemoteAssistancePartners()([]RemoteAssistancePartnerable)
    GetReports()(DeviceManagementReportsable)
    GetResourceOperations()([]ResourceOperationable)
    GetRoleAssignments()([]DeviceAndAppManagementRoleAssignmentable)
    GetRoleDefinitions()([]RoleDefinitionable)
    GetSettings()(DeviceManagementSettingsable)
    GetSoftwareUpdateStatusSummary()(SoftwareUpdateStatusSummaryable)
    GetSubscriptionState()(*DeviceManagementSubscriptionState)
    GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartnerable)
    GetTermsAndConditions()([]TermsAndConditionsable)
    GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformance()([]UserExperienceAnalyticsAppHealthApplicationPerformanceable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)
    GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()([]UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)
    GetUserExperienceAnalyticsAppHealthDevicePerformance()([]UserExperienceAnalyticsAppHealthDevicePerformanceable)
    GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()([]UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)
    GetUserExperienceAnalyticsAppHealthOSVersionPerformance()([]UserExperienceAnalyticsAppHealthOSVersionPerformanceable)
    GetUserExperienceAnalyticsAppHealthOverview()(UserExperienceAnalyticsCategoryable)
    GetUserExperienceAnalyticsBaselines()([]UserExperienceAnalyticsBaselineable)
    GetUserExperienceAnalyticsCategories()([]UserExperienceAnalyticsCategoryable)
    GetUserExperienceAnalyticsDevicePerformance()([]UserExperienceAnalyticsDevicePerformanceable)
    GetUserExperienceAnalyticsDeviceScores()([]UserExperienceAnalyticsDeviceScoresable)
    GetUserExperienceAnalyticsDeviceStartupHistory()([]UserExperienceAnalyticsDeviceStartupHistoryable)
    GetUserExperienceAnalyticsDeviceStartupProcesses()([]UserExperienceAnalyticsDeviceStartupProcessable)
    GetUserExperienceAnalyticsDeviceStartupProcessPerformance()([]UserExperienceAnalyticsDeviceStartupProcessPerformanceable)
    GetUserExperienceAnalyticsMetricHistory()([]UserExperienceAnalyticsMetricHistoryable)
    GetUserExperienceAnalyticsModelScores()([]UserExperienceAnalyticsModelScoresable)
    GetUserExperienceAnalyticsOverview()(UserExperienceAnalyticsOverviewable)
    GetUserExperienceAnalyticsScoreHistory()([]UserExperienceAnalyticsScoreHistoryable)
    GetUserExperienceAnalyticsSettings()(UserExperienceAnalyticsSettingsable)
    GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable)
    GetUserExperienceAnalyticsWorkFromAnywhereMetrics()([]UserExperienceAnalyticsWorkFromAnywhereMetricable)
    GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance()([]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)
    GetVirtualEndpoint()(VirtualEndpointable)
    GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentityable)
    GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummaryable)
    GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummaryable)
    GetWindowsMalwareInformation()([]WindowsMalwareInformationable)
    GetWindowsMalwareOverview()(WindowsMalwareOverviewable)
    SetApplePushNotificationCertificate(value ApplePushNotificationCertificateable)()
    SetAuditEvents(value []AuditEventable)()
    SetComplianceManagementPartners(value []ComplianceManagementPartnerable)()
    SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)()
    SetDetectedApps(value []DetectedAppable)()
    SetDeviceCategories(value []DeviceCategoryable)()
    SetDeviceCompliancePolicies(value []DeviceCompliancePolicyable)()
    SetDeviceCompliancePolicyDeviceStateSummary(value DeviceCompliancePolicyDeviceStateSummaryable)()
    SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)()
    SetDeviceConfigurationDeviceStateSummaries(value DeviceConfigurationDeviceStateSummaryable)()
    SetDeviceConfigurations(value []DeviceConfigurationable)()
    SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfigurationable)()
    SetDeviceManagementPartners(value []DeviceManagementPartnerable)()
    SetDeviceProtectionOverview(value DeviceProtectionOverviewable)()
    SetExchangeConnectors(value []DeviceManagementExchangeConnectorable)()
    SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentityable)()
    SetIntuneAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetIntuneBrand(value IntuneBrandable)()
    SetIosUpdateStatuses(value []IosUpdateDeviceStatusable)()
    SetManagedDeviceOverview(value ManagedDeviceOverviewable)()
    SetManagedDevices(value []ManagedDeviceable)()
    SetMobileAppTroubleshootingEvents(value []MobileAppTroubleshootingEventable)()
    SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnectorable)()
    SetNotificationMessageTemplates(value []NotificationMessageTemplateable)()
    SetRemoteAssistancePartners(value []RemoteAssistancePartnerable)()
    SetReports(value DeviceManagementReportsable)()
    SetResourceOperations(value []ResourceOperationable)()
    SetRoleAssignments(value []DeviceAndAppManagementRoleAssignmentable)()
    SetRoleDefinitions(value []RoleDefinitionable)()
    SetSettings(value DeviceManagementSettingsable)()
    SetSoftwareUpdateStatusSummary(value SoftwareUpdateStatusSummaryable)()
    SetSubscriptionState(value *DeviceManagementSubscriptionState)()
    SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartnerable)()
    SetTermsAndConditions(value []TermsAndConditionsable)()
    SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformance(value []UserExperienceAnalyticsAppHealthApplicationPerformanceable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)()
    SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(value []UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)()
    SetUserExperienceAnalyticsAppHealthDevicePerformance(value []UserExperienceAnalyticsAppHealthDevicePerformanceable)()
    SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(value []UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)()
    SetUserExperienceAnalyticsAppHealthOSVersionPerformance(value []UserExperienceAnalyticsAppHealthOSVersionPerformanceable)()
    SetUserExperienceAnalyticsAppHealthOverview(value UserExperienceAnalyticsCategoryable)()
    SetUserExperienceAnalyticsBaselines(value []UserExperienceAnalyticsBaselineable)()
    SetUserExperienceAnalyticsCategories(value []UserExperienceAnalyticsCategoryable)()
    SetUserExperienceAnalyticsDevicePerformance(value []UserExperienceAnalyticsDevicePerformanceable)()
    SetUserExperienceAnalyticsDeviceScores(value []UserExperienceAnalyticsDeviceScoresable)()
    SetUserExperienceAnalyticsDeviceStartupHistory(value []UserExperienceAnalyticsDeviceStartupHistoryable)()
    SetUserExperienceAnalyticsDeviceStartupProcesses(value []UserExperienceAnalyticsDeviceStartupProcessable)()
    SetUserExperienceAnalyticsDeviceStartupProcessPerformance(value []UserExperienceAnalyticsDeviceStartupProcessPerformanceable)()
    SetUserExperienceAnalyticsMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)()
    SetUserExperienceAnalyticsModelScores(value []UserExperienceAnalyticsModelScoresable)()
    SetUserExperienceAnalyticsOverview(value UserExperienceAnalyticsOverviewable)()
    SetUserExperienceAnalyticsScoreHistory(value []UserExperienceAnalyticsScoreHistoryable)()
    SetUserExperienceAnalyticsSettings(value UserExperienceAnalyticsSettingsable)()
    SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable)()
    SetUserExperienceAnalyticsWorkFromAnywhereMetrics(value []UserExperienceAnalyticsWorkFromAnywhereMetricable)()
    SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance(value []UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)()
    SetVirtualEndpoint(value VirtualEndpointable)()
    SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentityable)()
    SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummaryable)()
    SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummaryable)()
    SetWindowsMalwareInformation(value []WindowsMalwareInformationable)()
    SetWindowsMalwareOverview(value WindowsMalwareOverviewable)()
}
