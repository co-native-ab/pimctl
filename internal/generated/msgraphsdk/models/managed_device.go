package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevice devices that are managed or pre-enrolled through Intune
type ManagedDevice struct {
    Entity
    // The code that allows the Activation Lock on managed device to be bypassed. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity in LIST call. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
    activationLockBypassCode *string
    // Android security patch level. This property is read-only.
    androidSecurityPatchLevel *string
    // The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
    azureADDeviceId *string
    // Whether the device is Azure Active Directory registered. This property is read-only.
    azureADRegistered *bool
    // The DateTime when device compliance grace period expires. This property is read-only.
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Compliance state.
    complianceState *ComplianceState
    // ConfigrMgr client enabled features. This property is read-only.
    configurationManagerClientEnabledFeatures ConfigurationManagerClientEnabledFeaturesable
    // List of ComplexType deviceActionResult objects. This property is read-only.
    deviceActionResults []DeviceActionResultable
    // Device category
    deviceCategory DeviceCategoryable
    // Device category display name. Default is an empty string. Supports $filter operator 'eq' and 'or'. This property is read-only.
    deviceCategoryDisplayName *string
    // Device compliance policy states for this device.
    deviceCompliancePolicyStates []DeviceCompliancePolicyStateable
    // Device configuration states for this device.
    deviceConfigurationStates []DeviceConfigurationStateable
    // Possible ways of adding a mobile device to management.
    deviceEnrollmentType *DeviceEnrollmentType
    // The device health attestation state. This property is read-only.
    deviceHealthAttestationState DeviceHealthAttestationStateable
    // Name of the device. This property is read-only.
    deviceName *string
    // Device registration status.
    deviceRegistrationState *DeviceRegistrationState
    // Whether the device is Exchange ActiveSync activated. This property is read-only.
    easActivated *bool
    // Exchange ActivationSync activation time of the device. This property is read-only.
    easActivationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Exchange ActiveSync Id of the device. This property is read-only.
    easDeviceId *string
    // Email(s) for the user associated with the device. This property is read-only.
    emailAddress *string
    // Enrollment time of the device. Supports $filter operator 'lt' and 'gt'. This property is read-only.
    enrolledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
    enrollmentProfileName *string
    // Indicates Ethernet MAC Address of the device. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity. Individual get call with select query options is needed to retrieve actual values. Example: deviceManagement/managedDevices({managedDeviceId})?$select=ethernetMacAddress Supports: $select. $Search is not supported. Read-only. This property is read-only.
    ethernetMacAddress *string
    // Device Exchange Access State.
    exchangeAccessState *DeviceManagementExchangeAccessState
    // Device Exchange Access State Reason.
    exchangeAccessStateReason *DeviceManagementExchangeAccessStateReason
    // Last time the device contacted Exchange. This property is read-only.
    exchangeLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Free Storage in Bytes. Default value is 0. Read-only. This property is read-only.
    freeStorageSpaceInBytes *int64
    // Integrated Circuit Card Identifier, it is A SIM card's unique identification number. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
    iccid *string
    // IMEI. This property is read-only.
    imei *string
    // Device encryption status. This property is read-only.
    isEncrypted *bool
    // Device supervised status. This property is read-only.
    isSupervised *bool
    // Whether the device is jail broken or rooted. Default is an empty string. Supports $filter operator 'eq' and 'or'. This property is read-only.
    jailBroken *string
    // The date and time that the device last completed a successful sync with Intune. Supports $filter operator 'lt' and 'gt'. This property is read-only.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of log collection requests
    logCollectionRequests []DeviceLogCollectionResponseable
    // Automatically generated name to identify a device. Can be overwritten to a user friendly name.
    managedDeviceName *string
    // Owner type of device.
    managedDeviceOwnerType *ManagedDeviceOwnerType
    // The managementAgent property
    managementAgent *ManagementAgentType
    // Reports device management certificate expiration date. This property is read-only.
    managementCertificateExpirationDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Manufacturer of the device. This property is read-only.
    manufacturer *string
    // MEID. This property is read-only.
    meid *string
    // Model of the device. This property is read-only.
    model *string
    // Notes on the device created by IT Admin. Default is null. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported.
    notes *string
    // Operating system of the device. Windows, iOS, etc. This property is read-only.
    operatingSystem *string
    // Operating system version of the device. This property is read-only.
    osVersion *string
    // Available health states for the Device Health API
    partnerReportedThreatState *ManagedDevicePartnerReportedHealthState
    // Phone number of the device. This property is read-only.
    phoneNumber *string
    // Total Memory in Bytes. Default is 0. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. Read-only. This property is read-only.
    physicalMemoryInBytes *int64
    // An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
    remoteAssistanceSessionErrorDetails *string
    // Url that allows a Remote Assistance session to be established with the device. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. This property is read-only.
    remoteAssistanceSessionUrl *string
    // Reports if the managed iOS device is user approval enrollment. This property is read-only.
    requireUserEnrollmentApproval *bool
    // SerialNumber. This property is read-only.
    serialNumber *string
    // Subscriber Carrier. This property is read-only.
    subscriberCarrier *string
    // Total Storage in Bytes. This property is read-only.
    totalStorageSpaceInBytes *int64
    // Unique Device Identifier for iOS and macOS devices. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
    udid *string
    // User display name. This property is read-only.
    userDisplayName *string
    // Unique Identifier for the user associated with the device. This property is read-only.
    userId *string
    // Device user principal name. This property is read-only.
    userPrincipalName *string
    // The primary users associated with the managed device.
    users []Userable
    // Wi-Fi MAC. This property is read-only.
    wiFiMacAddress *string
    // The device protection status. This property is read-only.
    windowsProtectionState WindowsProtectionStateable
}
// NewManagedDevice instantiates a new ManagedDevice and sets the default values.
func NewManagedDevice()(*ManagedDevice) {
    m := &ManagedDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagedDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevice(), nil
}
// GetActivationLockBypassCode gets the activationLockBypassCode property value. The code that allows the Activation Lock on managed device to be bypassed. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity in LIST call. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetActivationLockBypassCode()(*string) {
    return m.activationLockBypassCode
}
// GetAndroidSecurityPatchLevel gets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetAndroidSecurityPatchLevel()(*string) {
    return m.androidSecurityPatchLevel
}
// GetAzureADDeviceId gets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetAzureADDeviceId()(*string) {
    return m.azureADDeviceId
}
// GetAzureADRegistered gets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
// returns a *bool when successful
func (m *ManagedDevice) GetAzureADRegistered()(*bool) {
    return m.azureADRegistered
}
// GetComplianceGracePeriodExpirationDateTime gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
// returns a *Time when successful
func (m *ManagedDevice) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.complianceGracePeriodExpirationDateTime
}
// GetComplianceState gets the complianceState property value. Compliance state.
// returns a *ComplianceState when successful
func (m *ManagedDevice) GetComplianceState()(*ComplianceState) {
    return m.complianceState
}
// GetConfigurationManagerClientEnabledFeatures gets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
// returns a ConfigurationManagerClientEnabledFeaturesable when successful
func (m *ManagedDevice) GetConfigurationManagerClientEnabledFeatures()(ConfigurationManagerClientEnabledFeaturesable) {
    return m.configurationManagerClientEnabledFeatures
}
// GetDeviceActionResults gets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
// returns a []DeviceActionResultable when successful
func (m *ManagedDevice) GetDeviceActionResults()([]DeviceActionResultable) {
    return m.deviceActionResults
}
// GetDeviceCategory gets the deviceCategory property value. Device category
// returns a DeviceCategoryable when successful
func (m *ManagedDevice) GetDeviceCategory()(DeviceCategoryable) {
    return m.deviceCategory
}
// GetDeviceCategoryDisplayName gets the deviceCategoryDisplayName property value. Device category display name. Default is an empty string. Supports $filter operator 'eq' and 'or'. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetDeviceCategoryDisplayName()(*string) {
    return m.deviceCategoryDisplayName
}
// GetDeviceCompliancePolicyStates gets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
// returns a []DeviceCompliancePolicyStateable when successful
func (m *ManagedDevice) GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyStateable) {
    return m.deviceCompliancePolicyStates
}
// GetDeviceConfigurationStates gets the deviceConfigurationStates property value. Device configuration states for this device.
// returns a []DeviceConfigurationStateable when successful
func (m *ManagedDevice) GetDeviceConfigurationStates()([]DeviceConfigurationStateable) {
    return m.deviceConfigurationStates
}
// GetDeviceEnrollmentType gets the deviceEnrollmentType property value. Possible ways of adding a mobile device to management.
// returns a *DeviceEnrollmentType when successful
func (m *ManagedDevice) GetDeviceEnrollmentType()(*DeviceEnrollmentType) {
    return m.deviceEnrollmentType
}
// GetDeviceHealthAttestationState gets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
// returns a DeviceHealthAttestationStateable when successful
func (m *ManagedDevice) GetDeviceHealthAttestationState()(DeviceHealthAttestationStateable) {
    return m.deviceHealthAttestationState
}
// GetDeviceName gets the deviceName property value. Name of the device. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDeviceRegistrationState gets the deviceRegistrationState property value. Device registration status.
// returns a *DeviceRegistrationState when successful
func (m *ManagedDevice) GetDeviceRegistrationState()(*DeviceRegistrationState) {
    return m.deviceRegistrationState
}
// GetEasActivated gets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
// returns a *bool when successful
func (m *ManagedDevice) GetEasActivated()(*bool) {
    return m.easActivated
}
// GetEasActivationDateTime gets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
// returns a *Time when successful
func (m *ManagedDevice) GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.easActivationDateTime
}
// GetEasDeviceId gets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetEasDeviceId()(*string) {
    return m.easDeviceId
}
// GetEmailAddress gets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetEnrolledDateTime gets the enrolledDateTime property value. Enrollment time of the device. Supports $filter operator 'lt' and 'gt'. This property is read-only.
// returns a *Time when successful
func (m *ManagedDevice) GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.enrolledDateTime
}
// GetEnrollmentProfileName gets the enrollmentProfileName property value. Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetEnrollmentProfileName()(*string) {
    return m.enrollmentProfileName
}
// GetEthernetMacAddress gets the ethernetMacAddress property value. Indicates Ethernet MAC Address of the device. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity. Individual get call with select query options is needed to retrieve actual values. Example: deviceManagement/managedDevices({managedDeviceId})?$select=ethernetMacAddress Supports: $select. $Search is not supported. Read-only. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetEthernetMacAddress()(*string) {
    return m.ethernetMacAddress
}
// GetExchangeAccessState gets the exchangeAccessState property value. Device Exchange Access State.
// returns a *DeviceManagementExchangeAccessState when successful
func (m *ManagedDevice) GetExchangeAccessState()(*DeviceManagementExchangeAccessState) {
    return m.exchangeAccessState
}
// GetExchangeAccessStateReason gets the exchangeAccessStateReason property value. Device Exchange Access State Reason.
// returns a *DeviceManagementExchangeAccessStateReason when successful
func (m *ManagedDevice) GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason) {
    return m.exchangeAccessStateReason
}
// GetExchangeLastSuccessfulSyncDateTime gets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
// returns a *Time when successful
func (m *ManagedDevice) GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.exchangeLastSuccessfulSyncDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ManagedDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationLockBypassCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationLockBypassCode(val)
        }
        return nil
    }
    res["androidSecurityPatchLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidSecurityPatchLevel(val)
        }
        return nil
    }
    res["azureADDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADDeviceId(val)
        }
        return nil
    }
    res["azureADRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADRegistered(val)
        }
        return nil
    }
    res["complianceGracePeriodExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceGracePeriodExpirationDateTime(val)
        }
        return nil
    }
    res["complianceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceState(val.(*ComplianceState))
        }
        return nil
    }
    res["configurationManagerClientEnabledFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationManagerClientEnabledFeaturesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientEnabledFeatures(val.(ConfigurationManagerClientEnabledFeaturesable))
        }
        return nil
    }
    res["deviceActionResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceActionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceActionResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceActionResultable)
                }
            }
            m.SetDeviceActionResults(res)
        }
        return nil
    }
    res["deviceCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategory(val.(DeviceCategoryable))
        }
        return nil
    }
    res["deviceCategoryDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategoryDisplayName(val)
        }
        return nil
    }
    res["deviceCompliancePolicyStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicyStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicyStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceCompliancePolicyStateable)
                }
            }
            m.SetDeviceCompliancePolicyStates(res)
        }
        return nil
    }
    res["deviceConfigurationStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceConfigurationStateable)
                }
            }
            m.SetDeviceConfigurationStates(res)
        }
        return nil
    }
    res["deviceEnrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceEnrollmentType(val.(*DeviceEnrollmentType))
        }
        return nil
    }
    res["deviceHealthAttestationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceHealthAttestationStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceHealthAttestationState(val.(DeviceHealthAttestationStateable))
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceRegistrationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRegistrationState(val.(*DeviceRegistrationState))
        }
        return nil
    }
    res["easActivated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivated(val)
        }
        return nil
    }
    res["easActivationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivationDateTime(val)
        }
        return nil
    }
    res["easDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasDeviceId(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["enrolledDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDateTime(val)
        }
        return nil
    }
    res["enrollmentProfileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentProfileName(val)
        }
        return nil
    }
    res["ethernetMacAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEthernetMacAddress(val)
        }
        return nil
    }
    res["exchangeAccessState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAccessState(val.(*DeviceManagementExchangeAccessState))
        }
        return nil
    }
    res["exchangeAccessStateReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessStateReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAccessStateReason(val.(*DeviceManagementExchangeAccessStateReason))
        }
        return nil
    }
    res["exchangeLastSuccessfulSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["freeStorageSpaceInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeStorageSpaceInBytes(val)
        }
        return nil
    }
    res["iccid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIccid(val)
        }
        return nil
    }
    res["imei"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImei(val)
        }
        return nil
    }
    res["isEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["isSupervised"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
        }
        return nil
    }
    res["jailBroken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJailBroken(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["logCollectionRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceLogCollectionResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceLogCollectionResponseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceLogCollectionResponseable)
                }
            }
            m.SetLogCollectionRequests(res)
        }
        return nil
    }
    res["managedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["managedDeviceOwnerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceOwnerType(val.(*ManagedDeviceOwnerType))
        }
        return nil
    }
    res["managementAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementAgent(val.(*ManagementAgentType))
        }
        return nil
    }
    res["managementCertificateExpirationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementCertificateExpirationDate(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["meid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeid(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
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
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["partnerReportedThreatState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDevicePartnerReportedHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerReportedThreatState(val.(*ManagedDevicePartnerReportedHealthState))
        }
        return nil
    }
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["physicalMemoryInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhysicalMemoryInBytes(val)
        }
        return nil
    }
    res["remoteAssistanceSessionErrorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionErrorDetails(val)
        }
        return nil
    }
    res["remoteAssistanceSessionUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionUrl(val)
        }
        return nil
    }
    res["requireUserEnrollmentApproval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireUserEnrollmentApproval(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["subscriberCarrier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriberCarrier(val)
        }
        return nil
    }
    res["totalStorageSpaceInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageSpaceInBytes(val)
        }
        return nil
    }
    res["udid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUdid(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
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
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetUsers(res)
        }
        return nil
    }
    res["wiFiMacAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiMacAddress(val)
        }
        return nil
    }
    res["windowsProtectionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsProtectionStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsProtectionState(val.(WindowsProtectionStateable))
        }
        return nil
    }
    return res
}
// GetFreeStorageSpaceInBytes gets the freeStorageSpaceInBytes property value. Free Storage in Bytes. Default value is 0. Read-only. This property is read-only.
// returns a *int64 when successful
func (m *ManagedDevice) GetFreeStorageSpaceInBytes()(*int64) {
    return m.freeStorageSpaceInBytes
}
// GetIccid gets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetIccid()(*string) {
    return m.iccid
}
// GetImei gets the imei property value. IMEI. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetImei()(*string) {
    return m.imei
}
// GetIsEncrypted gets the isEncrypted property value. Device encryption status. This property is read-only.
// returns a *bool when successful
func (m *ManagedDevice) GetIsEncrypted()(*bool) {
    return m.isEncrypted
}
// GetIsSupervised gets the isSupervised property value. Device supervised status. This property is read-only.
// returns a *bool when successful
func (m *ManagedDevice) GetIsSupervised()(*bool) {
    return m.isSupervised
}
// GetJailBroken gets the jailBroken property value. Whether the device is jail broken or rooted. Default is an empty string. Supports $filter operator 'eq' and 'or'. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetJailBroken()(*string) {
    return m.jailBroken
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. Supports $filter operator 'lt' and 'gt'. This property is read-only.
// returns a *Time when successful
func (m *ManagedDevice) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetLogCollectionRequests gets the logCollectionRequests property value. List of log collection requests
// returns a []DeviceLogCollectionResponseable when successful
func (m *ManagedDevice) GetLogCollectionRequests()([]DeviceLogCollectionResponseable) {
    return m.logCollectionRequests
}
// GetManagedDeviceName gets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
// returns a *string when successful
func (m *ManagedDevice) GetManagedDeviceName()(*string) {
    return m.managedDeviceName
}
// GetManagedDeviceOwnerType gets the managedDeviceOwnerType property value. Owner type of device.
// returns a *ManagedDeviceOwnerType when successful
func (m *ManagedDevice) GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType) {
    return m.managedDeviceOwnerType
}
// GetManagementAgent gets the managementAgent property value. The managementAgent property
// returns a *ManagementAgentType when successful
func (m *ManagedDevice) GetManagementAgent()(*ManagementAgentType) {
    return m.managementAgent
}
// GetManagementCertificateExpirationDate gets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
// returns a *Time when successful
func (m *ManagedDevice) GetManagementCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.managementCertificateExpirationDate
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetMeid gets the meid property value. MEID. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetMeid()(*string) {
    return m.meid
}
// GetModel gets the model property value. Model of the device. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetModel()(*string) {
    return m.model
}
// GetNotes gets the notes property value. Notes on the device created by IT Admin. Default is null. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported.
// returns a *string when successful
func (m *ManagedDevice) GetNotes()(*string) {
    return m.notes
}
// GetOperatingSystem gets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetOsVersion gets the osVersion property value. Operating system version of the device. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetOsVersion()(*string) {
    return m.osVersion
}
// GetPartnerReportedThreatState gets the partnerReportedThreatState property value. Available health states for the Device Health API
// returns a *ManagedDevicePartnerReportedHealthState when successful
func (m *ManagedDevice) GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState) {
    return m.partnerReportedThreatState
}
// GetPhoneNumber gets the phoneNumber property value. Phone number of the device. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPhysicalMemoryInBytes gets the physicalMemoryInBytes property value. Total Memory in Bytes. Default is 0. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. Read-only. This property is read-only.
// returns a *int64 when successful
func (m *ManagedDevice) GetPhysicalMemoryInBytes()(*int64) {
    return m.physicalMemoryInBytes
}
// GetRemoteAssistanceSessionErrorDetails gets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetRemoteAssistanceSessionErrorDetails()(*string) {
    return m.remoteAssistanceSessionErrorDetails
}
// GetRemoteAssistanceSessionUrl gets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetRemoteAssistanceSessionUrl()(*string) {
    return m.remoteAssistanceSessionUrl
}
// GetRequireUserEnrollmentApproval gets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
// returns a *bool when successful
func (m *ManagedDevice) GetRequireUserEnrollmentApproval()(*bool) {
    return m.requireUserEnrollmentApproval
}
// GetSerialNumber gets the serialNumber property value. SerialNumber. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetSubscriberCarrier gets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetSubscriberCarrier()(*string) {
    return m.subscriberCarrier
}
// GetTotalStorageSpaceInBytes gets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
// returns a *int64 when successful
func (m *ManagedDevice) GetTotalStorageSpaceInBytes()(*int64) {
    return m.totalStorageSpaceInBytes
}
// GetUdid gets the udid property value. Unique Device Identifier for iOS and macOS devices. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetUdid()(*string) {
    return m.udid
}
// GetUserDisplayName gets the userDisplayName property value. User display name. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserId gets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. Device user principal name. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUsers gets the users property value. The primary users associated with the managed device.
// returns a []Userable when successful
func (m *ManagedDevice) GetUsers()([]Userable) {
    return m.users
}
// GetWiFiMacAddress gets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
// returns a *string when successful
func (m *ManagedDevice) GetWiFiMacAddress()(*string) {
    return m.wiFiMacAddress
}
// GetWindowsProtectionState gets the windowsProtectionState property value. The device protection status. This property is read-only.
// returns a WindowsProtectionStateable when successful
func (m *ManagedDevice) GetWindowsProtectionState()(WindowsProtectionStateable) {
    return m.windowsProtectionState
}
// Serialize serializes information the current object
func (m *ManagedDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetComplianceState() != nil {
        cast := (*m.GetComplianceState()).String()
        err = writer.WriteStringValue("complianceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceCategory", m.GetDeviceCategory())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicyStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCompliancePolicyStates()))
        for i, v := range m.GetDeviceCompliancePolicyStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicyStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurationStates()))
        for i, v := range m.GetDeviceConfigurationStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentType() != nil {
        cast := (*m.GetDeviceEnrollmentType()).String()
        err = writer.WriteStringValue("deviceEnrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceRegistrationState() != nil {
        cast := (*m.GetDeviceRegistrationState()).String()
        err = writer.WriteStringValue("deviceRegistrationState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeAccessState() != nil {
        cast := (*m.GetExchangeAccessState()).String()
        err = writer.WriteStringValue("exchangeAccessState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeAccessStateReason() != nil {
        cast := (*m.GetExchangeAccessStateReason()).String()
        err = writer.WriteStringValue("exchangeAccessStateReason", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLogCollectionRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLogCollectionRequests()))
        for i, v := range m.GetLogCollectionRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("logCollectionRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceOwnerType() != nil {
        cast := (*m.GetManagedDeviceOwnerType()).String()
        err = writer.WriteStringValue("managedDeviceOwnerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementAgent() != nil {
        cast := (*m.GetManagementAgent()).String()
        err = writer.WriteStringValue("managementAgent", &cast)
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
    if m.GetPartnerReportedThreatState() != nil {
        cast := (*m.GetPartnerReportedThreatState()).String()
        err = writer.WriteStringValue("partnerReportedThreatState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("users", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsProtectionState", m.GetWindowsProtectionState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationLockBypassCode sets the activationLockBypassCode property value. The code that allows the Activation Lock on managed device to be bypassed. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity in LIST call. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetActivationLockBypassCode(value *string)() {
    m.activationLockBypassCode = value
}
// SetAndroidSecurityPatchLevel sets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) SetAndroidSecurityPatchLevel(value *string)() {
    m.androidSecurityPatchLevel = value
}
// SetAzureADDeviceId sets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) SetAzureADDeviceId(value *string)() {
    m.azureADDeviceId = value
}
// SetAzureADRegistered sets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) SetAzureADRegistered(value *bool)() {
    m.azureADRegistered = value
}
// SetComplianceGracePeriodExpirationDateTime sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceGracePeriodExpirationDateTime = value
}
// SetComplianceState sets the complianceState property value. Compliance state.
func (m *ManagedDevice) SetComplianceState(value *ComplianceState)() {
    m.complianceState = value
}
// SetConfigurationManagerClientEnabledFeatures sets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) SetConfigurationManagerClientEnabledFeatures(value ConfigurationManagerClientEnabledFeaturesable)() {
    m.configurationManagerClientEnabledFeatures = value
}
// SetDeviceActionResults sets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) SetDeviceActionResults(value []DeviceActionResultable)() {
    m.deviceActionResults = value
}
// SetDeviceCategory sets the deviceCategory property value. Device category
func (m *ManagedDevice) SetDeviceCategory(value DeviceCategoryable)() {
    m.deviceCategory = value
}
// SetDeviceCategoryDisplayName sets the deviceCategoryDisplayName property value. Device category display name. Default is an empty string. Supports $filter operator 'eq' and 'or'. This property is read-only.
func (m *ManagedDevice) SetDeviceCategoryDisplayName(value *string)() {
    m.deviceCategoryDisplayName = value
}
// SetDeviceCompliancePolicyStates sets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyStateable)() {
    m.deviceCompliancePolicyStates = value
}
// SetDeviceConfigurationStates sets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) SetDeviceConfigurationStates(value []DeviceConfigurationStateable)() {
    m.deviceConfigurationStates = value
}
// SetDeviceEnrollmentType sets the deviceEnrollmentType property value. Possible ways of adding a mobile device to management.
func (m *ManagedDevice) SetDeviceEnrollmentType(value *DeviceEnrollmentType)() {
    m.deviceEnrollmentType = value
}
// SetDeviceHealthAttestationState sets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) SetDeviceHealthAttestationState(value DeviceHealthAttestationStateable)() {
    m.deviceHealthAttestationState = value
}
// SetDeviceName sets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDeviceRegistrationState sets the deviceRegistrationState property value. Device registration status.
func (m *ManagedDevice) SetDeviceRegistrationState(value *DeviceRegistrationState)() {
    m.deviceRegistrationState = value
}
// SetEasActivated sets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) SetEasActivated(value *bool)() {
    m.easActivated = value
}
// SetEasActivationDateTime sets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.easActivationDateTime = value
}
// SetEasDeviceId sets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) SetEasDeviceId(value *string)() {
    m.easDeviceId = value
}
// SetEmailAddress sets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// SetEnrolledDateTime sets the enrolledDateTime property value. Enrollment time of the device. Supports $filter operator 'lt' and 'gt'. This property is read-only.
func (m *ManagedDevice) SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrolledDateTime = value
}
// SetEnrollmentProfileName sets the enrollmentProfileName property value. Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
func (m *ManagedDevice) SetEnrollmentProfileName(value *string)() {
    m.enrollmentProfileName = value
}
// SetEthernetMacAddress sets the ethernetMacAddress property value. Indicates Ethernet MAC Address of the device. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity. Individual get call with select query options is needed to retrieve actual values. Example: deviceManagement/managedDevices({managedDeviceId})?$select=ethernetMacAddress Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetEthernetMacAddress(value *string)() {
    m.ethernetMacAddress = value
}
// SetExchangeAccessState sets the exchangeAccessState property value. Device Exchange Access State.
func (m *ManagedDevice) SetExchangeAccessState(value *DeviceManagementExchangeAccessState)() {
    m.exchangeAccessState = value
}
// SetExchangeAccessStateReason sets the exchangeAccessStateReason property value. Device Exchange Access State Reason.
func (m *ManagedDevice) SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)() {
    m.exchangeAccessStateReason = value
}
// SetExchangeLastSuccessfulSyncDateTime sets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.exchangeLastSuccessfulSyncDateTime = value
}
// SetFreeStorageSpaceInBytes sets the freeStorageSpaceInBytes property value. Free Storage in Bytes. Default value is 0. Read-only. This property is read-only.
func (m *ManagedDevice) SetFreeStorageSpaceInBytes(value *int64)() {
    m.freeStorageSpaceInBytes = value
}
// SetIccid sets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetIccid(value *string)() {
    m.iccid = value
}
// SetImei sets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) SetImei(value *string)() {
    m.imei = value
}
// SetIsEncrypted sets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
// SetIsSupervised sets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) SetIsSupervised(value *bool)() {
    m.isSupervised = value
}
// SetJailBroken sets the jailBroken property value. Whether the device is jail broken or rooted. Default is an empty string. Supports $filter operator 'eq' and 'or'. This property is read-only.
func (m *ManagedDevice) SetJailBroken(value *string)() {
    m.jailBroken = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. Supports $filter operator 'lt' and 'gt'. This property is read-only.
func (m *ManagedDevice) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetLogCollectionRequests sets the logCollectionRequests property value. List of log collection requests
func (m *ManagedDevice) SetLogCollectionRequests(value []DeviceLogCollectionResponseable)() {
    m.logCollectionRequests = value
}
// SetManagedDeviceName sets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetManagedDeviceOwnerType sets the managedDeviceOwnerType property value. Owner type of device.
func (m *ManagedDevice) SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)() {
    m.managedDeviceOwnerType = value
}
// SetManagementAgent sets the managementAgent property value. The managementAgent property
func (m *ManagedDevice) SetManagementAgent(value *ManagementAgentType)() {
    m.managementAgent = value
}
// SetManagementCertificateExpirationDate sets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
func (m *ManagedDevice) SetManagementCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.managementCertificateExpirationDate = value
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetMeid sets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) SetMeid(value *string)() {
    m.meid = value
}
// SetModel sets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) SetModel(value *string)() {
    m.model = value
}
// SetNotes sets the notes property value. Notes on the device created by IT Admin. Default is null. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported.
func (m *ManagedDevice) SetNotes(value *string)() {
    m.notes = value
}
// SetOperatingSystem sets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetOsVersion sets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetPartnerReportedThreatState sets the partnerReportedThreatState property value. Available health states for the Device Health API
func (m *ManagedDevice) SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)() {
    m.partnerReportedThreatState = value
}
// SetPhoneNumber sets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPhysicalMemoryInBytes sets the physicalMemoryInBytes property value. Total Memory in Bytes. Default is 0. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. Read-only. This property is read-only.
func (m *ManagedDevice) SetPhysicalMemoryInBytes(value *int64)() {
    m.physicalMemoryInBytes = value
}
// SetRemoteAssistanceSessionErrorDetails sets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionErrorDetails(value *string)() {
    m.remoteAssistanceSessionErrorDetails = value
}
// SetRemoteAssistanceSessionUrl sets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionUrl(value *string)() {
    m.remoteAssistanceSessionUrl = value
}
// SetRequireUserEnrollmentApproval sets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
func (m *ManagedDevice) SetRequireUserEnrollmentApproval(value *bool)() {
    m.requireUserEnrollmentApproval = value
}
// SetSerialNumber sets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetSubscriberCarrier sets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) SetSubscriberCarrier(value *string)() {
    m.subscriberCarrier = value
}
// SetTotalStorageSpaceInBytes sets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) SetTotalStorageSpaceInBytes(value *int64)() {
    m.totalStorageSpaceInBytes = value
}
// SetUdid sets the udid property value. Unique Device Identifier for iOS and macOS devices. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetUdid(value *string)() {
    m.udid = value
}
// SetUserDisplayName sets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserId sets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUsers sets the users property value. The primary users associated with the managed device.
func (m *ManagedDevice) SetUsers(value []Userable)() {
    m.users = value
}
// SetWiFiMacAddress sets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) SetWiFiMacAddress(value *string)() {
    m.wiFiMacAddress = value
}
// SetWindowsProtectionState sets the windowsProtectionState property value. The device protection status. This property is read-only.
func (m *ManagedDevice) SetWindowsProtectionState(value WindowsProtectionStateable)() {
    m.windowsProtectionState = value
}
type ManagedDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivationLockBypassCode()(*string)
    GetAndroidSecurityPatchLevel()(*string)
    GetAzureADDeviceId()(*string)
    GetAzureADRegistered()(*bool)
    GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetComplianceState()(*ComplianceState)
    GetConfigurationManagerClientEnabledFeatures()(ConfigurationManagerClientEnabledFeaturesable)
    GetDeviceActionResults()([]DeviceActionResultable)
    GetDeviceCategory()(DeviceCategoryable)
    GetDeviceCategoryDisplayName()(*string)
    GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyStateable)
    GetDeviceConfigurationStates()([]DeviceConfigurationStateable)
    GetDeviceEnrollmentType()(*DeviceEnrollmentType)
    GetDeviceHealthAttestationState()(DeviceHealthAttestationStateable)
    GetDeviceName()(*string)
    GetDeviceRegistrationState()(*DeviceRegistrationState)
    GetEasActivated()(*bool)
    GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEasDeviceId()(*string)
    GetEmailAddress()(*string)
    GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEnrollmentProfileName()(*string)
    GetEthernetMacAddress()(*string)
    GetExchangeAccessState()(*DeviceManagementExchangeAccessState)
    GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason)
    GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFreeStorageSpaceInBytes()(*int64)
    GetIccid()(*string)
    GetImei()(*string)
    GetIsEncrypted()(*bool)
    GetIsSupervised()(*bool)
    GetJailBroken()(*string)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLogCollectionRequests()([]DeviceLogCollectionResponseable)
    GetManagedDeviceName()(*string)
    GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType)
    GetManagementAgent()(*ManagementAgentType)
    GetManagementCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManufacturer()(*string)
    GetMeid()(*string)
    GetModel()(*string)
    GetNotes()(*string)
    GetOperatingSystem()(*string)
    GetOsVersion()(*string)
    GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState)
    GetPhoneNumber()(*string)
    GetPhysicalMemoryInBytes()(*int64)
    GetRemoteAssistanceSessionErrorDetails()(*string)
    GetRemoteAssistanceSessionUrl()(*string)
    GetRequireUserEnrollmentApproval()(*bool)
    GetSerialNumber()(*string)
    GetSubscriberCarrier()(*string)
    GetTotalStorageSpaceInBytes()(*int64)
    GetUdid()(*string)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetUsers()([]Userable)
    GetWiFiMacAddress()(*string)
    GetWindowsProtectionState()(WindowsProtectionStateable)
    SetActivationLockBypassCode(value *string)()
    SetAndroidSecurityPatchLevel(value *string)()
    SetAzureADDeviceId(value *string)()
    SetAzureADRegistered(value *bool)()
    SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetComplianceState(value *ComplianceState)()
    SetConfigurationManagerClientEnabledFeatures(value ConfigurationManagerClientEnabledFeaturesable)()
    SetDeviceActionResults(value []DeviceActionResultable)()
    SetDeviceCategory(value DeviceCategoryable)()
    SetDeviceCategoryDisplayName(value *string)()
    SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyStateable)()
    SetDeviceConfigurationStates(value []DeviceConfigurationStateable)()
    SetDeviceEnrollmentType(value *DeviceEnrollmentType)()
    SetDeviceHealthAttestationState(value DeviceHealthAttestationStateable)()
    SetDeviceName(value *string)()
    SetDeviceRegistrationState(value *DeviceRegistrationState)()
    SetEasActivated(value *bool)()
    SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEasDeviceId(value *string)()
    SetEmailAddress(value *string)()
    SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEnrollmentProfileName(value *string)()
    SetEthernetMacAddress(value *string)()
    SetExchangeAccessState(value *DeviceManagementExchangeAccessState)()
    SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)()
    SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFreeStorageSpaceInBytes(value *int64)()
    SetIccid(value *string)()
    SetImei(value *string)()
    SetIsEncrypted(value *bool)()
    SetIsSupervised(value *bool)()
    SetJailBroken(value *string)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLogCollectionRequests(value []DeviceLogCollectionResponseable)()
    SetManagedDeviceName(value *string)()
    SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)()
    SetManagementAgent(value *ManagementAgentType)()
    SetManagementCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManufacturer(value *string)()
    SetMeid(value *string)()
    SetModel(value *string)()
    SetNotes(value *string)()
    SetOperatingSystem(value *string)()
    SetOsVersion(value *string)()
    SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)()
    SetPhoneNumber(value *string)()
    SetPhysicalMemoryInBytes(value *int64)()
    SetRemoteAssistanceSessionErrorDetails(value *string)()
    SetRemoteAssistanceSessionUrl(value *string)()
    SetRequireUserEnrollmentApproval(value *bool)()
    SetSerialNumber(value *string)()
    SetSubscriberCarrier(value *string)()
    SetTotalStorageSpaceInBytes(value *int64)()
    SetUdid(value *string)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetUsers(value []Userable)()
    SetWiFiMacAddress(value *string)()
    SetWindowsProtectionState(value WindowsProtectionStateable)()
}
