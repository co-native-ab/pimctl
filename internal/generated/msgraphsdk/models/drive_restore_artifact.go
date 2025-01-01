package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DriveRestoreArtifact struct {
    RestoreArtifactBase
    // The new site identifier if destinationType is new, and the input site ID if the destinationType is inPlace.
    restoredSiteId *string
    // The name of the restored site.
    restoredSiteName *string
    // The web URL of the restored site.
    restoredSiteWebUrl *string
}
// NewDriveRestoreArtifact instantiates a new DriveRestoreArtifact and sets the default values.
func NewDriveRestoreArtifact()(*DriveRestoreArtifact) {
    m := &DriveRestoreArtifact{
        RestoreArtifactBase: *NewRestoreArtifactBase(),
    }
    return m
}
// CreateDriveRestoreArtifactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDriveRestoreArtifactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriveRestoreArtifact(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DriveRestoreArtifact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RestoreArtifactBase.GetFieldDeserializers()
    res["restoredSiteId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoredSiteId(val)
        }
        return nil
    }
    res["restoredSiteName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoredSiteName(val)
        }
        return nil
    }
    res["restoredSiteWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoredSiteWebUrl(val)
        }
        return nil
    }
    return res
}
// GetRestoredSiteId gets the restoredSiteId property value. The new site identifier if destinationType is new, and the input site ID if the destinationType is inPlace.
// returns a *string when successful
func (m *DriveRestoreArtifact) GetRestoredSiteId()(*string) {
    return m.restoredSiteId
}
// GetRestoredSiteName gets the restoredSiteName property value. The name of the restored site.
// returns a *string when successful
func (m *DriveRestoreArtifact) GetRestoredSiteName()(*string) {
    return m.restoredSiteName
}
// GetRestoredSiteWebUrl gets the restoredSiteWebUrl property value. The web URL of the restored site.
// returns a *string when successful
func (m *DriveRestoreArtifact) GetRestoredSiteWebUrl()(*string) {
    return m.restoredSiteWebUrl
}
// Serialize serializes information the current object
func (m *DriveRestoreArtifact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RestoreArtifactBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("restoredSiteId", m.GetRestoredSiteId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRestoredSiteId sets the restoredSiteId property value. The new site identifier if destinationType is new, and the input site ID if the destinationType is inPlace.
func (m *DriveRestoreArtifact) SetRestoredSiteId(value *string)() {
    m.restoredSiteId = value
}
// SetRestoredSiteName sets the restoredSiteName property value. The name of the restored site.
func (m *DriveRestoreArtifact) SetRestoredSiteName(value *string)() {
    m.restoredSiteName = value
}
// SetRestoredSiteWebUrl sets the restoredSiteWebUrl property value. The web URL of the restored site.
func (m *DriveRestoreArtifact) SetRestoredSiteWebUrl(value *string)() {
    m.restoredSiteWebUrl = value
}
type DriveRestoreArtifactable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RestoreArtifactBaseable
    GetRestoredSiteId()(*string)
    GetRestoredSiteName()(*string)
    GetRestoredSiteWebUrl()(*string)
    SetRestoredSiteId(value *string)()
    SetRestoredSiteName(value *string)()
    SetRestoredSiteWebUrl(value *string)()
}
