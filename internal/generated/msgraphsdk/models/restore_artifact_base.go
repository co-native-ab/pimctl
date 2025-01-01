package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RestoreArtifactBase struct {
    Entity
    // The time when restoration of restore artifact is completed.
    completionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates the restoration destination. The possible values are: new, inPlace, unknownFutureValue.
    destinationType *DestinationType
    // Contains error details if the restore session fails or completes with an error.
    error PublicErrorable
    // Represents the date and time when an artifact is protected by a protectionPolicy and can be restored.
    restorePoint RestorePointable
    // The time when restoration of restore artifact is started.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The individual restoration status of the restore artifact. The possible values are: added, scheduling, scheduled, inProgress, succeeded, failed, unknownFutureValue.
    status *ArtifactRestoreStatus
}
// NewRestoreArtifactBase instantiates a new RestoreArtifactBase and sets the default values.
func NewRestoreArtifactBase()(*RestoreArtifactBase) {
    m := &RestoreArtifactBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRestoreArtifactBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRestoreArtifactBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.driveRestoreArtifact":
                        return NewDriveRestoreArtifact(), nil
                    case "#microsoft.graph.granularMailboxRestoreArtifact":
                        return NewGranularMailboxRestoreArtifact(), nil
                    case "#microsoft.graph.mailboxRestoreArtifact":
                        return NewMailboxRestoreArtifact(), nil
                    case "#microsoft.graph.siteRestoreArtifact":
                        return NewSiteRestoreArtifact(), nil
                }
            }
        }
    }
    return NewRestoreArtifactBase(), nil
}
// GetCompletionDateTime gets the completionDateTime property value. The time when restoration of restore artifact is completed.
// returns a *Time when successful
func (m *RestoreArtifactBase) GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completionDateTime
}
// GetDestinationType gets the destinationType property value. Indicates the restoration destination. The possible values are: new, inPlace, unknownFutureValue.
// returns a *DestinationType when successful
func (m *RestoreArtifactBase) GetDestinationType()(*DestinationType) {
    return m.destinationType
}
// GetError gets the error property value. Contains error details if the restore session fails or completes with an error.
// returns a PublicErrorable when successful
func (m *RestoreArtifactBase) GetError()(PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RestoreArtifactBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionDateTime(val)
        }
        return nil
    }
    res["destinationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDestinationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationType(val.(*DestinationType))
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(PublicErrorable))
        }
        return nil
    }
    res["restorePoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRestorePointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestorePoint(val.(RestorePointable))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseArtifactRestoreStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ArtifactRestoreStatus))
        }
        return nil
    }
    return res
}
// GetRestorePoint gets the restorePoint property value. Represents the date and time when an artifact is protected by a protectionPolicy and can be restored.
// returns a RestorePointable when successful
func (m *RestoreArtifactBase) GetRestorePoint()(RestorePointable) {
    return m.restorePoint
}
// GetStartDateTime gets the startDateTime property value. The time when restoration of restore artifact is started.
// returns a *Time when successful
func (m *RestoreArtifactBase) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStatus gets the status property value. The individual restoration status of the restore artifact. The possible values are: added, scheduling, scheduled, inProgress, succeeded, failed, unknownFutureValue.
// returns a *ArtifactRestoreStatus when successful
func (m *RestoreArtifactBase) GetStatus()(*ArtifactRestoreStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *RestoreArtifactBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("completionDateTime", m.GetCompletionDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDestinationType() != nil {
        cast := (*m.GetDestinationType()).String()
        err = writer.WriteStringValue("destinationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("restorePoint", m.GetRestorePoint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
    return nil
}
// SetCompletionDateTime sets the completionDateTime property value. The time when restoration of restore artifact is completed.
func (m *RestoreArtifactBase) SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completionDateTime = value
}
// SetDestinationType sets the destinationType property value. Indicates the restoration destination. The possible values are: new, inPlace, unknownFutureValue.
func (m *RestoreArtifactBase) SetDestinationType(value *DestinationType)() {
    m.destinationType = value
}
// SetError sets the error property value. Contains error details if the restore session fails or completes with an error.
func (m *RestoreArtifactBase) SetError(value PublicErrorable)() {
    m.error = value
}
// SetRestorePoint sets the restorePoint property value. Represents the date and time when an artifact is protected by a protectionPolicy and can be restored.
func (m *RestoreArtifactBase) SetRestorePoint(value RestorePointable)() {
    m.restorePoint = value
}
// SetStartDateTime sets the startDateTime property value. The time when restoration of restore artifact is started.
func (m *RestoreArtifactBase) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. The individual restoration status of the restore artifact. The possible values are: added, scheduling, scheduled, inProgress, succeeded, failed, unknownFutureValue.
func (m *RestoreArtifactBase) SetStatus(value *ArtifactRestoreStatus)() {
    m.status = value
}
type RestoreArtifactBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDestinationType()(*DestinationType)
    GetError()(PublicErrorable)
    GetRestorePoint()(RestorePointable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*ArtifactRestoreStatus)
    SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDestinationType(value *DestinationType)()
    SetError(value PublicErrorable)()
    SetRestorePoint(value RestorePointable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *ArtifactRestoreStatus)()
}
