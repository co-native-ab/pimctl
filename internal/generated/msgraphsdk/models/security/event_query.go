package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EventQuery struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Represents unique identification for the  query. 'Asset ID' for SharePoint Online and OneDrive for Business, 'keywords' for Exchange Online.
    query *string
    // Represents the type of query associated with an event. 'files' for SPO and ODB and 'messages' for EXO.The possible values are: files, messages, unknownFutureValue.
    queryType *QueryType
}
// NewEventQuery instantiates a new EventQuery and sets the default values.
func NewEventQuery()(*EventQuery) {
    m := &EventQuery{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEventQueryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEventQueryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEventQuery(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EventQuery) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EventQuery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["query"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val)
        }
        return nil
    }
    res["queryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseQueryType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryType(val.(*QueryType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *EventQuery) GetOdataType()(*string) {
    return m.odataType
}
// GetQuery gets the query property value. Represents unique identification for the  query. 'Asset ID' for SharePoint Online and OneDrive for Business, 'keywords' for Exchange Online.
// returns a *string when successful
func (m *EventQuery) GetQuery()(*string) {
    return m.query
}
// GetQueryType gets the queryType property value. Represents the type of query associated with an event. 'files' for SPO and ODB and 'messages' for EXO.The possible values are: files, messages, unknownFutureValue.
// returns a *QueryType when successful
func (m *EventQuery) GetQueryType()(*QueryType) {
    return m.queryType
}
// Serialize serializes information the current object
func (m *EventQuery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    if m.GetQueryType() != nil {
        cast := (*m.GetQueryType()).String()
        err := writer.WriteStringValue("queryType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EventQuery) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EventQuery) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQuery sets the query property value. Represents unique identification for the  query. 'Asset ID' for SharePoint Online and OneDrive for Business, 'keywords' for Exchange Online.
func (m *EventQuery) SetQuery(value *string)() {
    m.query = value
}
// SetQueryType sets the queryType property value. Represents the type of query associated with an event. 'files' for SPO and ODB and 'messages' for EXO.The possible values are: files, messages, unknownFutureValue.
func (m *EventQuery) SetQueryType(value *QueryType)() {
    m.queryType = value
}
type EventQueryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetQuery()(*string)
    GetQueryType()(*QueryType)
    SetOdataType(value *string)()
    SetQuery(value *string)()
    SetQueryType(value *QueryType)()
}
