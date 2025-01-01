package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GoogleCloudResourceEvidence struct {
    AlertEvidence
    // The fullResourceName property
    fullResourceName *string
    // The zone or region where the resource is located.
    location *string
    // The type of location. Possible values are: unknown, regional, zonal, global, unknownFutureValue.
    locationType *GoogleCloudLocationType
    // The Google project ID as defined by the user.
    projectId *string
    // The project number assigned by Google.
    projectNumber *int64
    // The name of the resource.
    resourceName *string
    // The type of the resource.
    resourceType *string
}
// NewGoogleCloudResourceEvidence instantiates a new GoogleCloudResourceEvidence and sets the default values.
func NewGoogleCloudResourceEvidence()(*GoogleCloudResourceEvidence) {
    m := &GoogleCloudResourceEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.googleCloudResourceEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGoogleCloudResourceEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGoogleCloudResourceEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGoogleCloudResourceEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GoogleCloudResourceEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["fullResourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullResourceName(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    res["locationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGoogleCloudLocationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationType(val.(*GoogleCloudLocationType))
        }
        return nil
    }
    res["projectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProjectId(val)
        }
        return nil
    }
    res["projectNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProjectNumber(val)
        }
        return nil
    }
    res["resourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
        }
        return nil
    }
    res["resourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val)
        }
        return nil
    }
    return res
}
// GetFullResourceName gets the fullResourceName property value. The fullResourceName property
// returns a *string when successful
func (m *GoogleCloudResourceEvidence) GetFullResourceName()(*string) {
    return m.fullResourceName
}
// GetLocation gets the location property value. The zone or region where the resource is located.
// returns a *string when successful
func (m *GoogleCloudResourceEvidence) GetLocation()(*string) {
    return m.location
}
// GetLocationType gets the locationType property value. The type of location. Possible values are: unknown, regional, zonal, global, unknownFutureValue.
// returns a *GoogleCloudLocationType when successful
func (m *GoogleCloudResourceEvidence) GetLocationType()(*GoogleCloudLocationType) {
    return m.locationType
}
// GetProjectId gets the projectId property value. The Google project ID as defined by the user.
// returns a *string when successful
func (m *GoogleCloudResourceEvidence) GetProjectId()(*string) {
    return m.projectId
}
// GetProjectNumber gets the projectNumber property value. The project number assigned by Google.
// returns a *int64 when successful
func (m *GoogleCloudResourceEvidence) GetProjectNumber()(*int64) {
    return m.projectNumber
}
// GetResourceName gets the resourceName property value. The name of the resource.
// returns a *string when successful
func (m *GoogleCloudResourceEvidence) GetResourceName()(*string) {
    return m.resourceName
}
// GetResourceType gets the resourceType property value. The type of the resource.
// returns a *string when successful
func (m *GoogleCloudResourceEvidence) GetResourceType()(*string) {
    return m.resourceType
}
// Serialize serializes information the current object
func (m *GoogleCloudResourceEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fullResourceName", m.GetFullResourceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    if m.GetLocationType() != nil {
        cast := (*m.GetLocationType()).String()
        err = writer.WriteStringValue("locationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("projectId", m.GetProjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("projectNumber", m.GetProjectNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceType", m.GetResourceType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFullResourceName sets the fullResourceName property value. The fullResourceName property
func (m *GoogleCloudResourceEvidence) SetFullResourceName(value *string)() {
    m.fullResourceName = value
}
// SetLocation sets the location property value. The zone or region where the resource is located.
func (m *GoogleCloudResourceEvidence) SetLocation(value *string)() {
    m.location = value
}
// SetLocationType sets the locationType property value. The type of location. Possible values are: unknown, regional, zonal, global, unknownFutureValue.
func (m *GoogleCloudResourceEvidence) SetLocationType(value *GoogleCloudLocationType)() {
    m.locationType = value
}
// SetProjectId sets the projectId property value. The Google project ID as defined by the user.
func (m *GoogleCloudResourceEvidence) SetProjectId(value *string)() {
    m.projectId = value
}
// SetProjectNumber sets the projectNumber property value. The project number assigned by Google.
func (m *GoogleCloudResourceEvidence) SetProjectNumber(value *int64)() {
    m.projectNumber = value
}
// SetResourceName sets the resourceName property value. The name of the resource.
func (m *GoogleCloudResourceEvidence) SetResourceName(value *string)() {
    m.resourceName = value
}
// SetResourceType sets the resourceType property value. The type of the resource.
func (m *GoogleCloudResourceEvidence) SetResourceType(value *string)() {
    m.resourceType = value
}
type GoogleCloudResourceEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFullResourceName()(*string)
    GetLocation()(*string)
    GetLocationType()(*GoogleCloudLocationType)
    GetProjectId()(*string)
    GetProjectNumber()(*int64)
    GetResourceName()(*string)
    GetResourceType()(*string)
    SetFullResourceName(value *string)()
    SetLocation(value *string)()
    SetLocationType(value *GoogleCloudLocationType)()
    SetProjectId(value *string)()
    SetProjectNumber(value *int64)()
    SetResourceName(value *string)()
    SetResourceType(value *string)()
}
