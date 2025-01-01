package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EdiscoverySearchExportOperation struct {
    CaseOperation
    // The additional items to include in the export. The possible values are: none, teamsAndYammerConversations, cloudAttachments, allDocumentVersions, subfolderContents, listAttachments, unknownFutureValue.
    additionalOptions *AdditionalOptions
    // The description of the export by the user.
    description *string
    // The name of export provided by the user.
    displayName *string
    // Items to be included in the export. The possible values are: searchHits, partiallyIndexed, unknownFutureValue.
    exportCriteria *ExportCriteria
    // Contains the properties for an export file metadata, including downloadUrl, fileName, and size.
    exportFileMetadata []ExportFileMetadataable
    // Format of the emails of the export. The possible values are: pst, msg, eml, unknownFutureValue.
    exportFormat *ExportFormat
    // Location scope for partially indexed items. You can choose to include partially indexed items only in responsive locations with search hits or in all targeted locations. The possible values are: responsiveLocations, nonresponsiveLocations, unknownFutureValue.
    exportLocation *ExportLocation
    // Indicates whether to export single items.
    exportSingleItems *bool
    // The eDiscovery searches under each case.
    search EdiscoverySearchable
}
// NewEdiscoverySearchExportOperation instantiates a new EdiscoverySearchExportOperation and sets the default values.
func NewEdiscoverySearchExportOperation()(*EdiscoverySearchExportOperation) {
    m := &EdiscoverySearchExportOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEdiscoverySearchExportOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdiscoverySearchExportOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoverySearchExportOperation(), nil
}
// GetAdditionalOptions gets the additionalOptions property value. The additional items to include in the export. The possible values are: none, teamsAndYammerConversations, cloudAttachments, allDocumentVersions, subfolderContents, listAttachments, unknownFutureValue.
// returns a *AdditionalOptions when successful
func (m *EdiscoverySearchExportOperation) GetAdditionalOptions()(*AdditionalOptions) {
    return m.additionalOptions
}
// GetDescription gets the description property value. The description of the export by the user.
// returns a *string when successful
func (m *EdiscoverySearchExportOperation) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of export provided by the user.
// returns a *string when successful
func (m *EdiscoverySearchExportOperation) GetDisplayName()(*string) {
    return m.displayName
}
// GetExportCriteria gets the exportCriteria property value. Items to be included in the export. The possible values are: searchHits, partiallyIndexed, unknownFutureValue.
// returns a *ExportCriteria when successful
func (m *EdiscoverySearchExportOperation) GetExportCriteria()(*ExportCriteria) {
    return m.exportCriteria
}
// GetExportFileMetadata gets the exportFileMetadata property value. Contains the properties for an export file metadata, including downloadUrl, fileName, and size.
// returns a []ExportFileMetadataable when successful
func (m *EdiscoverySearchExportOperation) GetExportFileMetadata()([]ExportFileMetadataable) {
    return m.exportFileMetadata
}
// GetExportFormat gets the exportFormat property value. Format of the emails of the export. The possible values are: pst, msg, eml, unknownFutureValue.
// returns a *ExportFormat when successful
func (m *EdiscoverySearchExportOperation) GetExportFormat()(*ExportFormat) {
    return m.exportFormat
}
// GetExportLocation gets the exportLocation property value. Location scope for partially indexed items. You can choose to include partially indexed items only in responsive locations with search hits or in all targeted locations. The possible values are: responsiveLocations, nonresponsiveLocations, unknownFutureValue.
// returns a *ExportLocation when successful
func (m *EdiscoverySearchExportOperation) GetExportLocation()(*ExportLocation) {
    return m.exportLocation
}
// GetExportSingleItems gets the exportSingleItems property value. Indicates whether to export single items.
// returns a *bool when successful
func (m *EdiscoverySearchExportOperation) GetExportSingleItems()(*bool) {
    return m.exportSingleItems
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EdiscoverySearchExportOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["additionalOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdditionalOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalOptions(val.(*AdditionalOptions))
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
    res["exportCriteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExportCriteria)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportCriteria(val.(*ExportCriteria))
        }
        return nil
    }
    res["exportFileMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExportFileMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExportFileMetadataable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExportFileMetadataable)
                }
            }
            m.SetExportFileMetadata(res)
        }
        return nil
    }
    res["exportFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExportFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportFormat(val.(*ExportFormat))
        }
        return nil
    }
    res["exportLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExportLocation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportLocation(val.(*ExportLocation))
        }
        return nil
    }
    res["exportSingleItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportSingleItems(val)
        }
        return nil
    }
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoverySearchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val.(EdiscoverySearchable))
        }
        return nil
    }
    return res
}
// GetSearch gets the search property value. The eDiscovery searches under each case.
// returns a EdiscoverySearchable when successful
func (m *EdiscoverySearchExportOperation) GetSearch()(EdiscoverySearchable) {
    return m.search
}
// Serialize serializes information the current object
func (m *EdiscoverySearchExportOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalOptions() != nil {
        cast := (*m.GetAdditionalOptions()).String()
        err = writer.WriteStringValue("additionalOptions", &cast)
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
    if m.GetExportCriteria() != nil {
        cast := (*m.GetExportCriteria()).String()
        err = writer.WriteStringValue("exportCriteria", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportFileMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExportFileMetadata()))
        for i, v := range m.GetExportFileMetadata() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("exportFileMetadata", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportFormat() != nil {
        cast := (*m.GetExportFormat()).String()
        err = writer.WriteStringValue("exportFormat", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportLocation() != nil {
        cast := (*m.GetExportLocation()).String()
        err = writer.WriteStringValue("exportLocation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("exportSingleItems", m.GetExportSingleItems())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalOptions sets the additionalOptions property value. The additional items to include in the export. The possible values are: none, teamsAndYammerConversations, cloudAttachments, allDocumentVersions, subfolderContents, listAttachments, unknownFutureValue.
func (m *EdiscoverySearchExportOperation) SetAdditionalOptions(value *AdditionalOptions)() {
    m.additionalOptions = value
}
// SetDescription sets the description property value. The description of the export by the user.
func (m *EdiscoverySearchExportOperation) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of export provided by the user.
func (m *EdiscoverySearchExportOperation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExportCriteria sets the exportCriteria property value. Items to be included in the export. The possible values are: searchHits, partiallyIndexed, unknownFutureValue.
func (m *EdiscoverySearchExportOperation) SetExportCriteria(value *ExportCriteria)() {
    m.exportCriteria = value
}
// SetExportFileMetadata sets the exportFileMetadata property value. Contains the properties for an export file metadata, including downloadUrl, fileName, and size.
func (m *EdiscoverySearchExportOperation) SetExportFileMetadata(value []ExportFileMetadataable)() {
    m.exportFileMetadata = value
}
// SetExportFormat sets the exportFormat property value. Format of the emails of the export. The possible values are: pst, msg, eml, unknownFutureValue.
func (m *EdiscoverySearchExportOperation) SetExportFormat(value *ExportFormat)() {
    m.exportFormat = value
}
// SetExportLocation sets the exportLocation property value. Location scope for partially indexed items. You can choose to include partially indexed items only in responsive locations with search hits or in all targeted locations. The possible values are: responsiveLocations, nonresponsiveLocations, unknownFutureValue.
func (m *EdiscoverySearchExportOperation) SetExportLocation(value *ExportLocation)() {
    m.exportLocation = value
}
// SetExportSingleItems sets the exportSingleItems property value. Indicates whether to export single items.
func (m *EdiscoverySearchExportOperation) SetExportSingleItems(value *bool)() {
    m.exportSingleItems = value
}
// SetSearch sets the search property value. The eDiscovery searches under each case.
func (m *EdiscoverySearchExportOperation) SetSearch(value EdiscoverySearchable)() {
    m.search = value
}
type EdiscoverySearchExportOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalOptions()(*AdditionalOptions)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExportCriteria()(*ExportCriteria)
    GetExportFileMetadata()([]ExportFileMetadataable)
    GetExportFormat()(*ExportFormat)
    GetExportLocation()(*ExportLocation)
    GetExportSingleItems()(*bool)
    GetSearch()(EdiscoverySearchable)
    SetAdditionalOptions(value *AdditionalOptions)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExportCriteria(value *ExportCriteria)()
    SetExportFileMetadata(value []ExportFileMetadataable)()
    SetExportFormat(value *ExportFormat)()
    SetExportLocation(value *ExportLocation)()
    SetExportSingleItems(value *bool)()
    SetSearch(value EdiscoverySearchable)()
}
