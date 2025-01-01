package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageResourceRequest struct {
    Entity
    // The catalog property
    catalog AccessPackageCatalogable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of the request. Use adminAdd to add a resource, if the caller is an administrator or resource owner, adminUpdate to update a resource, or adminRemove to remove a resource.
    requestType *AccessPackageRequestType
    // The resource property
    resource AccessPackageResourceable
    // The outcome of whether the service was able to add the resource to the catalog. The value is delivered if the resource was added or removed, and deliveryFailed if it couldn't be added or removed. Read-only.
    state *AccessPackageRequestState
}
// NewAccessPackageResourceRequest instantiates a new AccessPackageResourceRequest and sets the default values.
func NewAccessPackageResourceRequest()(*AccessPackageResourceRequest) {
    m := &AccessPackageResourceRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageResourceRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageResourceRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceRequest(), nil
}
// GetCatalog gets the catalog property value. The catalog property
// returns a AccessPackageCatalogable when successful
func (m *AccessPackageResourceRequest) GetCatalog()(AccessPackageCatalogable) {
    return m.catalog
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *AccessPackageResourceRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccessPackageResourceRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["catalog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalog(val.(AccessPackageCatalogable))
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
    res["requestType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageRequestType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestType(val.(*AccessPackageRequestType))
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(AccessPackageResourceable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageRequestState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AccessPackageRequestState))
        }
        return nil
    }
    return res
}
// GetRequestType gets the requestType property value. The type of the request. Use adminAdd to add a resource, if the caller is an administrator or resource owner, adminUpdate to update a resource, or adminRemove to remove a resource.
// returns a *AccessPackageRequestType when successful
func (m *AccessPackageResourceRequest) GetRequestType()(*AccessPackageRequestType) {
    return m.requestType
}
// GetResource gets the resource property value. The resource property
// returns a AccessPackageResourceable when successful
func (m *AccessPackageResourceRequest) GetResource()(AccessPackageResourceable) {
    return m.resource
}
// GetState gets the state property value. The outcome of whether the service was able to add the resource to the catalog. The value is delivered if the resource was added or removed, and deliveryFailed if it couldn't be added or removed. Read-only.
// returns a *AccessPackageRequestState when successful
func (m *AccessPackageResourceRequest) GetState()(*AccessPackageRequestState) {
    return m.state
}
// Serialize serializes information the current object
func (m *AccessPackageResourceRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("catalog", m.GetCatalog())
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
    if m.GetRequestType() != nil {
        cast := (*m.GetRequestType()).String()
        err = writer.WriteStringValue("requestType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCatalog sets the catalog property value. The catalog property
func (m *AccessPackageResourceRequest) SetCatalog(value AccessPackageCatalogable)() {
    m.catalog = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageResourceRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetRequestType sets the requestType property value. The type of the request. Use adminAdd to add a resource, if the caller is an administrator or resource owner, adminUpdate to update a resource, or adminRemove to remove a resource.
func (m *AccessPackageResourceRequest) SetRequestType(value *AccessPackageRequestType)() {
    m.requestType = value
}
// SetResource sets the resource property value. The resource property
func (m *AccessPackageResourceRequest) SetResource(value AccessPackageResourceable)() {
    m.resource = value
}
// SetState sets the state property value. The outcome of whether the service was able to add the resource to the catalog. The value is delivered if the resource was added or removed, and deliveryFailed if it couldn't be added or removed. Read-only.
func (m *AccessPackageResourceRequest) SetState(value *AccessPackageRequestState)() {
    m.state = value
}
type AccessPackageResourceRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCatalog()(AccessPackageCatalogable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestType()(*AccessPackageRequestType)
    GetResource()(AccessPackageResourceable)
    GetState()(*AccessPackageRequestState)
    SetCatalog(value AccessPackageCatalogable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestType(value *AccessPackageRequestType)()
    SetResource(value AccessPackageResourceable)()
    SetState(value *AccessPackageRequestState)()
}
