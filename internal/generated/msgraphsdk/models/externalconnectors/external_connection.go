package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type ExternalConnection struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // Collects configurable settings related to activities involving connector content.
    activitySettings ActivitySettingsable
    // Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
    configuration Configurationable
    // The Teams app ID. Optional.
    connectorId *string
    // Description of the connection displayed in the Microsoft 365 admin center. Optional.
    description *string
    // The groups property
    groups []ExternalGroupable
    // The items property
    items []ExternalItemable
    // The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
    name *string
    // The operations property
    operations []ConnectionOperationable
    // The schema property
    schema Schemaable
    // The settings configuring the search experience for content in this connection, such as the display templates for search results.
    searchSettings SearchSettingsable
    // Indicates the current state of the connection. Possible values are: draft, ready, obsolete, limitExceeded, unknownFutureValue.
    state *ConnectionState
}
// NewExternalConnection instantiates a new ExternalConnection and sets the default values.
func NewExternalConnection()(*ExternalConnection) {
    m := &ExternalConnection{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateExternalConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalConnection(), nil
}
// GetActivitySettings gets the activitySettings property value. Collects configurable settings related to activities involving connector content.
// returns a ActivitySettingsable when successful
func (m *ExternalConnection) GetActivitySettings()(ActivitySettingsable) {
    return m.activitySettings
}
// GetConfiguration gets the configuration property value. Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
// returns a Configurationable when successful
func (m *ExternalConnection) GetConfiguration()(Configurationable) {
    return m.configuration
}
// GetConnectorId gets the connectorId property value. The Teams app ID. Optional.
// returns a *string when successful
func (m *ExternalConnection) GetConnectorId()(*string) {
    return m.connectorId
}
// GetDescription gets the description property value. Description of the connection displayed in the Microsoft 365 admin center. Optional.
// returns a *string when successful
func (m *ExternalConnection) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExternalConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activitySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateActivitySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivitySettings(val.(ActivitySettingsable))
        }
        return nil
    }
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(Configurationable))
        }
        return nil
    }
    res["connectorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorId(val)
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
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalGroupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExternalGroupable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExternalItemable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectionOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectionOperationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConnectionOperationable)
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["schema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchema(val.(Schemaable))
        }
        return nil
    }
    res["searchSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchSettings(val.(SearchSettingsable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ConnectionState))
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. The groups property
// returns a []ExternalGroupable when successful
func (m *ExternalConnection) GetGroups()([]ExternalGroupable) {
    return m.groups
}
// GetItems gets the items property value. The items property
// returns a []ExternalItemable when successful
func (m *ExternalConnection) GetItems()([]ExternalItemable) {
    return m.items
}
// GetName gets the name property value. The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
// returns a *string when successful
func (m *ExternalConnection) GetName()(*string) {
    return m.name
}
// GetOperations gets the operations property value. The operations property
// returns a []ConnectionOperationable when successful
func (m *ExternalConnection) GetOperations()([]ConnectionOperationable) {
    return m.operations
}
// GetSchema gets the schema property value. The schema property
// returns a Schemaable when successful
func (m *ExternalConnection) GetSchema()(Schemaable) {
    return m.schema
}
// GetSearchSettings gets the searchSettings property value. The settings configuring the search experience for content in this connection, such as the display templates for search results.
// returns a SearchSettingsable when successful
func (m *ExternalConnection) GetSearchSettings()(SearchSettingsable) {
    return m.searchSettings
}
// GetState gets the state property value. Indicates the current state of the connection. Possible values are: draft, ready, obsolete, limitExceeded, unknownFutureValue.
// returns a *ConnectionState when successful
func (m *ExternalConnection) GetState()(*ConnectionState) {
    return m.state
}
// Serialize serializes information the current object
func (m *ExternalConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activitySettings", m.GetActivitySettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("connectorId", m.GetConnectorId())
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
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schema", m.GetSchema())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("searchSettings", m.GetSearchSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivitySettings sets the activitySettings property value. Collects configurable settings related to activities involving connector content.
func (m *ExternalConnection) SetActivitySettings(value ActivitySettingsable)() {
    m.activitySettings = value
}
// SetConfiguration sets the configuration property value. Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
func (m *ExternalConnection) SetConfiguration(value Configurationable)() {
    m.configuration = value
}
// SetConnectorId sets the connectorId property value. The Teams app ID. Optional.
func (m *ExternalConnection) SetConnectorId(value *string)() {
    m.connectorId = value
}
// SetDescription sets the description property value. Description of the connection displayed in the Microsoft 365 admin center. Optional.
func (m *ExternalConnection) SetDescription(value *string)() {
    m.description = value
}
// SetGroups sets the groups property value. The groups property
func (m *ExternalConnection) SetGroups(value []ExternalGroupable)() {
    m.groups = value
}
// SetItems sets the items property value. The items property
func (m *ExternalConnection) SetItems(value []ExternalItemable)() {
    m.items = value
}
// SetName sets the name property value. The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
func (m *ExternalConnection) SetName(value *string)() {
    m.name = value
}
// SetOperations sets the operations property value. The operations property
func (m *ExternalConnection) SetOperations(value []ConnectionOperationable)() {
    m.operations = value
}
// SetSchema sets the schema property value. The schema property
func (m *ExternalConnection) SetSchema(value Schemaable)() {
    m.schema = value
}
// SetSearchSettings sets the searchSettings property value. The settings configuring the search experience for content in this connection, such as the display templates for search results.
func (m *ExternalConnection) SetSearchSettings(value SearchSettingsable)() {
    m.searchSettings = value
}
// SetState sets the state property value. Indicates the current state of the connection. Possible values are: draft, ready, obsolete, limitExceeded, unknownFutureValue.
func (m *ExternalConnection) SetState(value *ConnectionState)() {
    m.state = value
}
type ExternalConnectionable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivitySettings()(ActivitySettingsable)
    GetConfiguration()(Configurationable)
    GetConnectorId()(*string)
    GetDescription()(*string)
    GetGroups()([]ExternalGroupable)
    GetItems()([]ExternalItemable)
    GetName()(*string)
    GetOperations()([]ConnectionOperationable)
    GetSchema()(Schemaable)
    GetSearchSettings()(SearchSettingsable)
    GetState()(*ConnectionState)
    SetActivitySettings(value ActivitySettingsable)()
    SetConfiguration(value Configurationable)()
    SetConnectorId(value *string)()
    SetDescription(value *string)()
    SetGroups(value []ExternalGroupable)()
    SetItems(value []ExternalItemable)()
    SetName(value *string)()
    SetOperations(value []ConnectionOperationable)()
    SetSchema(value Schemaable)()
    SetSearchSettings(value SearchSettingsable)()
    SetState(value *ConnectionState)()
}
