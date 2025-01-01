package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FileStorageContainer struct {
    Entity
    // Container type ID of the fileStorageContainer. For details about container types, see Container Types. Each container must have only one container type. Read-only.
    containerTypeId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Date and time of the fileStorageContainer creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Custom property collection for the fileStorageContainer. Read-write.
    customProperties FileStorageContainerCustomPropertyDictionaryable
    // Provides a user-visible description of the fileStorageContainer. Read-write.
    description *string
    // The display name of the fileStorageContainer. Read-write.
    displayName *string
    // The drive of the resource fileStorageContainer. Read-only.
    drive Driveable
    // Indicates the lock state of the fileStorageContainer. The possible values are unlocked and lockedReadOnly. Read-only.
    lockState *SiteLockState
    // The set of permissions for users in the fileStorageContainer. Permission for each user is set by the roles property. The possible values are: reader, writer, manager, and owner. Read-write.
    permissions []Permissionable
    // Recycle bin of the fileStorageContainer. Read-only.
    recycleBin RecycleBinable
    // The settings property
    settings FileStorageContainerSettingsable
    // Status of the fileStorageContainer. Containers are created as inactive and require activation. Inactive containers are subjected to automatic deletion in 24 hours. The possible values are: inactive, active. Read-only.
    status *FileStorageContainerStatus
    // Data specific to the current user. Read-only.
    viewpoint FileStorageContainerViewpointable
}
// NewFileStorageContainer instantiates a new FileStorageContainer and sets the default values.
func NewFileStorageContainer()(*FileStorageContainer) {
    m := &FileStorageContainer{
        Entity: *NewEntity(),
    }
    return m
}
// CreateFileStorageContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainer(), nil
}
// GetContainerTypeId gets the containerTypeId property value. Container type ID of the fileStorageContainer. For details about container types, see Container Types. Each container must have only one container type. Read-only.
// returns a *UUID when successful
func (m *FileStorageContainer) GetContainerTypeId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.containerTypeId
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of the fileStorageContainer creation. Read-only.
// returns a *Time when successful
func (m *FileStorageContainer) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomProperties gets the customProperties property value. Custom property collection for the fileStorageContainer. Read-write.
// returns a FileStorageContainerCustomPropertyDictionaryable when successful
func (m *FileStorageContainer) GetCustomProperties()(FileStorageContainerCustomPropertyDictionaryable) {
    return m.customProperties
}
// GetDescription gets the description property value. Provides a user-visible description of the fileStorageContainer. Read-write.
// returns a *string when successful
func (m *FileStorageContainer) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the fileStorageContainer. Read-write.
// returns a *string when successful
func (m *FileStorageContainer) GetDisplayName()(*string) {
    return m.displayName
}
// GetDrive gets the drive property value. The drive of the resource fileStorageContainer. Read-only.
// returns a Driveable when successful
func (m *FileStorageContainer) GetDrive()(Driveable) {
    return m.drive
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FileStorageContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["containerTypeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerTypeId(val)
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
    res["customProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileStorageContainerCustomPropertyDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomProperties(val.(FileStorageContainerCustomPropertyDictionaryable))
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
    res["drive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrive(val.(Driveable))
        }
        return nil
    }
    res["lockState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSiteLockState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockState(val.(*SiteLockState))
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Permissionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Permissionable)
                }
            }
            m.SetPermissions(res)
        }
        return nil
    }
    res["recycleBin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecycleBinFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecycleBin(val.(RecycleBinable))
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileStorageContainerSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(FileStorageContainerSettingsable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileStorageContainerStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*FileStorageContainerStatus))
        }
        return nil
    }
    res["viewpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileStorageContainerViewpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewpoint(val.(FileStorageContainerViewpointable))
        }
        return nil
    }
    return res
}
// GetLockState gets the lockState property value. Indicates the lock state of the fileStorageContainer. The possible values are unlocked and lockedReadOnly. Read-only.
// returns a *SiteLockState when successful
func (m *FileStorageContainer) GetLockState()(*SiteLockState) {
    return m.lockState
}
// GetPermissions gets the permissions property value. The set of permissions for users in the fileStorageContainer. Permission for each user is set by the roles property. The possible values are: reader, writer, manager, and owner. Read-write.
// returns a []Permissionable when successful
func (m *FileStorageContainer) GetPermissions()([]Permissionable) {
    return m.permissions
}
// GetRecycleBin gets the recycleBin property value. Recycle bin of the fileStorageContainer. Read-only.
// returns a RecycleBinable when successful
func (m *FileStorageContainer) GetRecycleBin()(RecycleBinable) {
    return m.recycleBin
}
// GetSettings gets the settings property value. The settings property
// returns a FileStorageContainerSettingsable when successful
func (m *FileStorageContainer) GetSettings()(FileStorageContainerSettingsable) {
    return m.settings
}
// GetStatus gets the status property value. Status of the fileStorageContainer. Containers are created as inactive and require activation. Inactive containers are subjected to automatic deletion in 24 hours. The possible values are: inactive, active. Read-only.
// returns a *FileStorageContainerStatus when successful
func (m *FileStorageContainer) GetStatus()(*FileStorageContainerStatus) {
    return m.status
}
// GetViewpoint gets the viewpoint property value. Data specific to the current user. Read-only.
// returns a FileStorageContainerViewpointable when successful
func (m *FileStorageContainer) GetViewpoint()(FileStorageContainerViewpointable) {
    return m.viewpoint
}
// Serialize serializes information the current object
func (m *FileStorageContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteUUIDValue("containerTypeId", m.GetContainerTypeId())
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
        err = writer.WriteObjectValue("customProperties", m.GetCustomProperties())
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
        err = writer.WriteObjectValue("drive", m.GetDrive())
        if err != nil {
            return err
        }
    }
    if m.GetLockState() != nil {
        cast := (*m.GetLockState()).String()
        err = writer.WriteStringValue("lockState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissions()))
        for i, v := range m.GetPermissions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("permissions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recycleBin", m.GetRecycleBin())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("viewpoint", m.GetViewpoint())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContainerTypeId sets the containerTypeId property value. Container type ID of the fileStorageContainer. For details about container types, see Container Types. Each container must have only one container type. Read-only.
func (m *FileStorageContainer) SetContainerTypeId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.containerTypeId = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of the fileStorageContainer creation. Read-only.
func (m *FileStorageContainer) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomProperties sets the customProperties property value. Custom property collection for the fileStorageContainer. Read-write.
func (m *FileStorageContainer) SetCustomProperties(value FileStorageContainerCustomPropertyDictionaryable)() {
    m.customProperties = value
}
// SetDescription sets the description property value. Provides a user-visible description of the fileStorageContainer. Read-write.
func (m *FileStorageContainer) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the fileStorageContainer. Read-write.
func (m *FileStorageContainer) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDrive sets the drive property value. The drive of the resource fileStorageContainer. Read-only.
func (m *FileStorageContainer) SetDrive(value Driveable)() {
    m.drive = value
}
// SetLockState sets the lockState property value. Indicates the lock state of the fileStorageContainer. The possible values are unlocked and lockedReadOnly. Read-only.
func (m *FileStorageContainer) SetLockState(value *SiteLockState)() {
    m.lockState = value
}
// SetPermissions sets the permissions property value. The set of permissions for users in the fileStorageContainer. Permission for each user is set by the roles property. The possible values are: reader, writer, manager, and owner. Read-write.
func (m *FileStorageContainer) SetPermissions(value []Permissionable)() {
    m.permissions = value
}
// SetRecycleBin sets the recycleBin property value. Recycle bin of the fileStorageContainer. Read-only.
func (m *FileStorageContainer) SetRecycleBin(value RecycleBinable)() {
    m.recycleBin = value
}
// SetSettings sets the settings property value. The settings property
func (m *FileStorageContainer) SetSettings(value FileStorageContainerSettingsable)() {
    m.settings = value
}
// SetStatus sets the status property value. Status of the fileStorageContainer. Containers are created as inactive and require activation. Inactive containers are subjected to automatic deletion in 24 hours. The possible values are: inactive, active. Read-only.
func (m *FileStorageContainer) SetStatus(value *FileStorageContainerStatus)() {
    m.status = value
}
// SetViewpoint sets the viewpoint property value. Data specific to the current user. Read-only.
func (m *FileStorageContainer) SetViewpoint(value FileStorageContainerViewpointable)() {
    m.viewpoint = value
}
type FileStorageContainerable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContainerTypeId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomProperties()(FileStorageContainerCustomPropertyDictionaryable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDrive()(Driveable)
    GetLockState()(*SiteLockState)
    GetPermissions()([]Permissionable)
    GetRecycleBin()(RecycleBinable)
    GetSettings()(FileStorageContainerSettingsable)
    GetStatus()(*FileStorageContainerStatus)
    GetViewpoint()(FileStorageContainerViewpointable)
    SetContainerTypeId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomProperties(value FileStorageContainerCustomPropertyDictionaryable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDrive(value Driveable)()
    SetLockState(value *SiteLockState)()
    SetPermissions(value []Permissionable)()
    SetRecycleBin(value RecycleBinable)()
    SetSettings(value FileStorageContainerSettingsable)()
    SetStatus(value *FileStorageContainerStatus)()
    SetViewpoint(value FileStorageContainerViewpointable)()
}