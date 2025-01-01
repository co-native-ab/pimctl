package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type EdiscoveryCase struct {
    CaseEscaped
    // The user who closed the case.
    closedBy ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable
    // The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    closedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Returns a list of case ediscoveryCustodian objects for this case.
    custodians []EdiscoveryCustodianable
    // The external case number for customer reference.
    externalId *string
    // Returns a list of case ediscoveryNoncustodialDataSource objects for this case.
    noncustodialDataSources []EdiscoveryNoncustodialDataSourceable
    // Returns a list of case caseOperation objects for this case.
    operations []CaseOperationable
    // Returns a list of eDiscoveryReviewSet objects in the case.
    reviewSets []EdiscoveryReviewSetable
    // Returns a list of eDiscoverySearch objects associated with this case.
    searches []EdiscoverySearchable
    // Returns a list of eDIscoverySettings objects in the case.
    settings EdiscoveryCaseSettingsable
    // Returns a list of ediscoveryReviewTag objects associated to this case.
    tags []EdiscoveryReviewTagable
}
// NewEdiscoveryCase instantiates a new EdiscoveryCase and sets the default values.
func NewEdiscoveryCase()(*EdiscoveryCase) {
    m := &EdiscoveryCase{
        CaseEscaped: *NewCaseEscaped(),
    }
    odataTypeValue := "#microsoft.graph.security.ediscoveryCase"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEdiscoveryCaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdiscoveryCaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCase(), nil
}
// GetClosedBy gets the closedBy property value. The user who closed the case.
// returns a IdentitySetable when successful
func (m *EdiscoveryCase) GetClosedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable) {
    return m.closedBy
}
// GetClosedDateTime gets the closedDateTime property value. The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *EdiscoveryCase) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.closedDateTime
}
// GetCustodians gets the custodians property value. Returns a list of case ediscoveryCustodian objects for this case.
// returns a []EdiscoveryCustodianable when successful
func (m *EdiscoveryCase) GetCustodians()([]EdiscoveryCustodianable) {
    return m.custodians
}
// GetExternalId gets the externalId property value. The external case number for customer reference.
// returns a *string when successful
func (m *EdiscoveryCase) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EdiscoveryCase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseEscaped.GetFieldDeserializers()
    res["closedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedBy(val.(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable))
        }
        return nil
    }
    res["closedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedDateTime(val)
        }
        return nil
    }
    res["custodians"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryCustodianFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryCustodianable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EdiscoveryCustodianable)
                }
            }
            m.SetCustodians(res)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["noncustodialDataSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryNoncustodialDataSourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EdiscoveryNoncustodialDataSourceable)
                }
            }
            m.SetNoncustodialDataSources(res)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCaseOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CaseOperationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CaseOperationable)
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["reviewSets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryReviewSetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EdiscoveryReviewSetable)
                }
            }
            m.SetReviewSets(res)
        }
        return nil
    }
    res["searches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoverySearchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoverySearchable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EdiscoverySearchable)
                }
            }
            m.SetSearches(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryCaseSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(EdiscoveryCaseSettingsable))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryReviewTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryReviewTagable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EdiscoveryReviewTagable)
                }
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
// GetNoncustodialDataSources gets the noncustodialDataSources property value. Returns a list of case ediscoveryNoncustodialDataSource objects for this case.
// returns a []EdiscoveryNoncustodialDataSourceable when successful
func (m *EdiscoveryCase) GetNoncustodialDataSources()([]EdiscoveryNoncustodialDataSourceable) {
    return m.noncustodialDataSources
}
// GetOperations gets the operations property value. Returns a list of case caseOperation objects for this case.
// returns a []CaseOperationable when successful
func (m *EdiscoveryCase) GetOperations()([]CaseOperationable) {
    return m.operations
}
// GetReviewSets gets the reviewSets property value. Returns a list of eDiscoveryReviewSet objects in the case.
// returns a []EdiscoveryReviewSetable when successful
func (m *EdiscoveryCase) GetReviewSets()([]EdiscoveryReviewSetable) {
    return m.reviewSets
}
// GetSearches gets the searches property value. Returns a list of eDiscoverySearch objects associated with this case.
// returns a []EdiscoverySearchable when successful
func (m *EdiscoveryCase) GetSearches()([]EdiscoverySearchable) {
    return m.searches
}
// GetSettings gets the settings property value. Returns a list of eDIscoverySettings objects in the case.
// returns a EdiscoveryCaseSettingsable when successful
func (m *EdiscoveryCase) GetSettings()(EdiscoveryCaseSettingsable) {
    return m.settings
}
// GetTags gets the tags property value. Returns a list of ediscoveryReviewTag objects associated to this case.
// returns a []EdiscoveryReviewTagable when successful
func (m *EdiscoveryCase) GetTags()([]EdiscoveryReviewTagable) {
    return m.tags
}
// Serialize serializes information the current object
func (m *EdiscoveryCase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseEscaped.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("closedBy", m.GetClosedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("closedDateTime", m.GetClosedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCustodians() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustodians()))
        for i, v := range m.GetCustodians() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("custodians", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    if m.GetNoncustodialDataSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNoncustodialDataSources()))
        for i, v := range m.GetNoncustodialDataSources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("noncustodialDataSources", cast)
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
    if m.GetReviewSets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewSets()))
        for i, v := range m.GetReviewSets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("reviewSets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSearches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSearches()))
        for i, v := range m.GetSearches() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("searches", cast)
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
    if m.GetTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTags()))
        for i, v := range m.GetTags() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("tags", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClosedBy sets the closedBy property value. The user who closed the case.
func (m *EdiscoveryCase) SetClosedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)() {
    m.closedBy = value
}
// SetClosedDateTime sets the closedDateTime property value. The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EdiscoveryCase) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closedDateTime = value
}
// SetCustodians sets the custodians property value. Returns a list of case ediscoveryCustodian objects for this case.
func (m *EdiscoveryCase) SetCustodians(value []EdiscoveryCustodianable)() {
    m.custodians = value
}
// SetExternalId sets the externalId property value. The external case number for customer reference.
func (m *EdiscoveryCase) SetExternalId(value *string)() {
    m.externalId = value
}
// SetNoncustodialDataSources sets the noncustodialDataSources property value. Returns a list of case ediscoveryNoncustodialDataSource objects for this case.
func (m *EdiscoveryCase) SetNoncustodialDataSources(value []EdiscoveryNoncustodialDataSourceable)() {
    m.noncustodialDataSources = value
}
// SetOperations sets the operations property value. Returns a list of case caseOperation objects for this case.
func (m *EdiscoveryCase) SetOperations(value []CaseOperationable)() {
    m.operations = value
}
// SetReviewSets sets the reviewSets property value. Returns a list of eDiscoveryReviewSet objects in the case.
func (m *EdiscoveryCase) SetReviewSets(value []EdiscoveryReviewSetable)() {
    m.reviewSets = value
}
// SetSearches sets the searches property value. Returns a list of eDiscoverySearch objects associated with this case.
func (m *EdiscoveryCase) SetSearches(value []EdiscoverySearchable)() {
    m.searches = value
}
// SetSettings sets the settings property value. Returns a list of eDIscoverySettings objects in the case.
func (m *EdiscoveryCase) SetSettings(value EdiscoveryCaseSettingsable)() {
    m.settings = value
}
// SetTags sets the tags property value. Returns a list of ediscoveryReviewTag objects associated to this case.
func (m *EdiscoveryCase) SetTags(value []EdiscoveryReviewTagable)() {
    m.tags = value
}
type EdiscoveryCaseable interface {
    CaseEscapedable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClosedBy()(ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)
    GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodians()([]EdiscoveryCustodianable)
    GetExternalId()(*string)
    GetNoncustodialDataSources()([]EdiscoveryNoncustodialDataSourceable)
    GetOperations()([]CaseOperationable)
    GetReviewSets()([]EdiscoveryReviewSetable)
    GetSearches()([]EdiscoverySearchable)
    GetSettings()(EdiscoveryCaseSettingsable)
    GetTags()([]EdiscoveryReviewTagable)
    SetClosedBy(value ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.IdentitySetable)()
    SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodians(value []EdiscoveryCustodianable)()
    SetExternalId(value *string)()
    SetNoncustodialDataSources(value []EdiscoveryNoncustodialDataSourceable)()
    SetOperations(value []CaseOperationable)()
    SetReviewSets(value []EdiscoveryReviewSetable)()
    SetSearches(value []EdiscoverySearchable)()
    SetSettings(value EdiscoveryCaseSettingsable)()
    SetTags(value []EdiscoveryReviewTagable)()
}
