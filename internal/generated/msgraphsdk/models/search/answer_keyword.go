package search

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AnswerKeyword struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A collection of keywords used to trigger the search answer.
    keywords []string
    // If true, indicates that the search term contains similar words to the keywords that should trigger the search answer.
    matchSimilarKeywords *bool
    // The OdataType property
    odataType *string
    // Unique keywords that guarantee the search answer is triggered.
    reservedKeywords []string
}
// NewAnswerKeyword instantiates a new AnswerKeyword and sets the default values.
func NewAnswerKeyword()(*AnswerKeyword) {
    m := &AnswerKeyword{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAnswerKeywordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAnswerKeywordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnswerKeyword(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AnswerKeyword) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AnswerKeyword) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keywords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetKeywords(res)
        }
        return nil
    }
    res["matchSimilarKeywords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchSimilarKeywords(val)
        }
        return nil
    }
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
    res["reservedKeywords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetReservedKeywords(res)
        }
        return nil
    }
    return res
}
// GetKeywords gets the keywords property value. A collection of keywords used to trigger the search answer.
// returns a []string when successful
func (m *AnswerKeyword) GetKeywords()([]string) {
    return m.keywords
}
// GetMatchSimilarKeywords gets the matchSimilarKeywords property value. If true, indicates that the search term contains similar words to the keywords that should trigger the search answer.
// returns a *bool when successful
func (m *AnswerKeyword) GetMatchSimilarKeywords()(*bool) {
    return m.matchSimilarKeywords
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AnswerKeyword) GetOdataType()(*string) {
    return m.odataType
}
// GetReservedKeywords gets the reservedKeywords property value. Unique keywords that guarantee the search answer is triggered.
// returns a []string when successful
func (m *AnswerKeyword) GetReservedKeywords()([]string) {
    return m.reservedKeywords
}
// Serialize serializes information the current object
func (m *AnswerKeyword) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetKeywords() != nil {
        err := writer.WriteCollectionOfStringValues("keywords", m.GetKeywords())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("matchSimilarKeywords", m.GetMatchSimilarKeywords())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetReservedKeywords() != nil {
        err := writer.WriteCollectionOfStringValues("reservedKeywords", m.GetReservedKeywords())
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
func (m *AnswerKeyword) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKeywords sets the keywords property value. A collection of keywords used to trigger the search answer.
func (m *AnswerKeyword) SetKeywords(value []string)() {
    m.keywords = value
}
// SetMatchSimilarKeywords sets the matchSimilarKeywords property value. If true, indicates that the search term contains similar words to the keywords that should trigger the search answer.
func (m *AnswerKeyword) SetMatchSimilarKeywords(value *bool)() {
    m.matchSimilarKeywords = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AnswerKeyword) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReservedKeywords sets the reservedKeywords property value. Unique keywords that guarantee the search answer is triggered.
func (m *AnswerKeyword) SetReservedKeywords(value []string)() {
    m.reservedKeywords = value
}
type AnswerKeywordable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeywords()([]string)
    GetMatchSimilarKeywords()(*bool)
    GetOdataType()(*string)
    GetReservedKeywords()([]string)
    SetKeywords(value []string)()
    SetMatchSimilarKeywords(value *bool)()
    SetOdataType(value *string)()
    SetReservedKeywords(value []string)()
}
