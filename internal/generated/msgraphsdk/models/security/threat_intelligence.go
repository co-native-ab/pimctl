package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba "github.com/co-native-ab/pimctl/internal/generated/msgraphsdk/models"
)

type ThreatIntelligence struct {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entity
    // Refers to indicators of threat or compromise highlighted in an article.Note: List retrieval is not yet supported.
    articleIndicators []ArticleIndicatorable
    // A list of article objects.
    articles []Articleable
    // Retrieve details about hostComponent objects.Note: List retrieval is not yet supported.
    hostComponents []HostComponentable
    // Retrieve details about hostCookie objects.Note: List retrieval is not yet supported.
    hostCookies []HostCookieable
    // Retrieve details about hostTracker objects.Note: List retrieval is not yet supported.
    hostPairs []HostPairable
    // Retrieve details about hostPort objects.Note: List retrieval is not yet supported.
    hostPorts []HostPortable
    // Refers to host objects that Microsoft Threat Intelligence has observed.Note: List retrieval is not yet supported.
    hosts []Hostable
    // Retrieve details about hostSslCertificate objects.Note: List retrieval is not yet supported.
    hostSslCertificates []HostSslCertificateable
    // Retrieve details about hostTracker objects.Note: List retrieval is not yet supported.
    hostTrackers []HostTrackerable
    // The intelligenceProfileIndicators property
    intelligenceProfileIndicators []IntelligenceProfileIndicatorable
    // A list of intelligenceProfile objects.
    intelProfiles []IntelligenceProfileable
    // Retrieve details about passiveDnsRecord objects.Note: List retrieval is not yet supported.
    passiveDnsRecords []PassiveDnsRecordable
    // Retrieve details about sslCertificate objects.Note: List retrieval is not yet supported.
    sslCertificates []SslCertificateable
    // Retrieve details about the subdomain.Note: List retrieval is not yet supported.
    subdomains []Subdomainable
    // Retrieve details about vulnerabilities.Note: List retrieval is not yet supported.
    vulnerabilities []Vulnerabilityable
    // Retrieve details about whoisHistoryRecord objects.Note: List retrieval is not yet supported.
    whoisHistoryRecords []WhoisHistoryRecordable
    // A list of whoisRecord objects.
    whoisRecords []WhoisRecordable
}
// NewThreatIntelligence instantiates a new ThreatIntelligence and sets the default values.
func NewThreatIntelligence()(*ThreatIntelligence) {
    m := &ThreatIntelligence{
        Entity: *ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.NewEntity(),
    }
    return m
}
// CreateThreatIntelligenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateThreatIntelligenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThreatIntelligence(), nil
}
// GetArticleIndicators gets the articleIndicators property value. Refers to indicators of threat or compromise highlighted in an article.Note: List retrieval is not yet supported.
// returns a []ArticleIndicatorable when successful
func (m *ThreatIntelligence) GetArticleIndicators()([]ArticleIndicatorable) {
    return m.articleIndicators
}
// GetArticles gets the articles property value. A list of article objects.
// returns a []Articleable when successful
func (m *ThreatIntelligence) GetArticles()([]Articleable) {
    return m.articles
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ThreatIntelligence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["articleIndicators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateArticleIndicatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ArticleIndicatorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ArticleIndicatorable)
                }
            }
            m.SetArticleIndicators(res)
        }
        return nil
    }
    res["articles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateArticleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Articleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Articleable)
                }
            }
            m.SetArticles(res)
        }
        return nil
    }
    res["hostComponents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostComponentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostComponentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostComponentable)
                }
            }
            m.SetHostComponents(res)
        }
        return nil
    }
    res["hostCookies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostCookieFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostCookieable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostCookieable)
                }
            }
            m.SetHostCookies(res)
        }
        return nil
    }
    res["hostPairs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostPairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostPairable)
                }
            }
            m.SetHostPairs(res)
        }
        return nil
    }
    res["hostPorts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostPortFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostPortable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostPortable)
                }
            }
            m.SetHostPorts(res)
        }
        return nil
    }
    res["hosts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Hostable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Hostable)
                }
            }
            m.SetHosts(res)
        }
        return nil
    }
    res["hostSslCertificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostSslCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostSslCertificateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostSslCertificateable)
                }
            }
            m.SetHostSslCertificates(res)
        }
        return nil
    }
    res["hostTrackers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostTrackerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostTrackerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HostTrackerable)
                }
            }
            m.SetHostTrackers(res)
        }
        return nil
    }
    res["intelligenceProfileIndicators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntelligenceProfileIndicatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntelligenceProfileIndicatorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IntelligenceProfileIndicatorable)
                }
            }
            m.SetIntelligenceProfileIndicators(res)
        }
        return nil
    }
    res["intelProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntelligenceProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntelligenceProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IntelligenceProfileable)
                }
            }
            m.SetIntelProfiles(res)
        }
        return nil
    }
    res["passiveDnsRecords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePassiveDnsRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PassiveDnsRecordable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PassiveDnsRecordable)
                }
            }
            m.SetPassiveDnsRecords(res)
        }
        return nil
    }
    res["sslCertificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSslCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SslCertificateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SslCertificateable)
                }
            }
            m.SetSslCertificates(res)
        }
        return nil
    }
    res["subdomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubdomainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Subdomainable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Subdomainable)
                }
            }
            m.SetSubdomains(res)
        }
        return nil
    }
    res["vulnerabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVulnerabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Vulnerabilityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Vulnerabilityable)
                }
            }
            m.SetVulnerabilities(res)
        }
        return nil
    }
    res["whoisHistoryRecords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWhoisHistoryRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WhoisHistoryRecordable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WhoisHistoryRecordable)
                }
            }
            m.SetWhoisHistoryRecords(res)
        }
        return nil
    }
    res["whoisRecords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWhoisRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WhoisRecordable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WhoisRecordable)
                }
            }
            m.SetWhoisRecords(res)
        }
        return nil
    }
    return res
}
// GetHostComponents gets the hostComponents property value. Retrieve details about hostComponent objects.Note: List retrieval is not yet supported.
// returns a []HostComponentable when successful
func (m *ThreatIntelligence) GetHostComponents()([]HostComponentable) {
    return m.hostComponents
}
// GetHostCookies gets the hostCookies property value. Retrieve details about hostCookie objects.Note: List retrieval is not yet supported.
// returns a []HostCookieable when successful
func (m *ThreatIntelligence) GetHostCookies()([]HostCookieable) {
    return m.hostCookies
}
// GetHostPairs gets the hostPairs property value. Retrieve details about hostTracker objects.Note: List retrieval is not yet supported.
// returns a []HostPairable when successful
func (m *ThreatIntelligence) GetHostPairs()([]HostPairable) {
    return m.hostPairs
}
// GetHostPorts gets the hostPorts property value. Retrieve details about hostPort objects.Note: List retrieval is not yet supported.
// returns a []HostPortable when successful
func (m *ThreatIntelligence) GetHostPorts()([]HostPortable) {
    return m.hostPorts
}
// GetHosts gets the hosts property value. Refers to host objects that Microsoft Threat Intelligence has observed.Note: List retrieval is not yet supported.
// returns a []Hostable when successful
func (m *ThreatIntelligence) GetHosts()([]Hostable) {
    return m.hosts
}
// GetHostSslCertificates gets the hostSslCertificates property value. Retrieve details about hostSslCertificate objects.Note: List retrieval is not yet supported.
// returns a []HostSslCertificateable when successful
func (m *ThreatIntelligence) GetHostSslCertificates()([]HostSslCertificateable) {
    return m.hostSslCertificates
}
// GetHostTrackers gets the hostTrackers property value. Retrieve details about hostTracker objects.Note: List retrieval is not yet supported.
// returns a []HostTrackerable when successful
func (m *ThreatIntelligence) GetHostTrackers()([]HostTrackerable) {
    return m.hostTrackers
}
// GetIntelligenceProfileIndicators gets the intelligenceProfileIndicators property value. The intelligenceProfileIndicators property
// returns a []IntelligenceProfileIndicatorable when successful
func (m *ThreatIntelligence) GetIntelligenceProfileIndicators()([]IntelligenceProfileIndicatorable) {
    return m.intelligenceProfileIndicators
}
// GetIntelProfiles gets the intelProfiles property value. A list of intelligenceProfile objects.
// returns a []IntelligenceProfileable when successful
func (m *ThreatIntelligence) GetIntelProfiles()([]IntelligenceProfileable) {
    return m.intelProfiles
}
// GetPassiveDnsRecords gets the passiveDnsRecords property value. Retrieve details about passiveDnsRecord objects.Note: List retrieval is not yet supported.
// returns a []PassiveDnsRecordable when successful
func (m *ThreatIntelligence) GetPassiveDnsRecords()([]PassiveDnsRecordable) {
    return m.passiveDnsRecords
}
// GetSslCertificates gets the sslCertificates property value. Retrieve details about sslCertificate objects.Note: List retrieval is not yet supported.
// returns a []SslCertificateable when successful
func (m *ThreatIntelligence) GetSslCertificates()([]SslCertificateable) {
    return m.sslCertificates
}
// GetSubdomains gets the subdomains property value. Retrieve details about the subdomain.Note: List retrieval is not yet supported.
// returns a []Subdomainable when successful
func (m *ThreatIntelligence) GetSubdomains()([]Subdomainable) {
    return m.subdomains
}
// GetVulnerabilities gets the vulnerabilities property value. Retrieve details about vulnerabilities.Note: List retrieval is not yet supported.
// returns a []Vulnerabilityable when successful
func (m *ThreatIntelligence) GetVulnerabilities()([]Vulnerabilityable) {
    return m.vulnerabilities
}
// GetWhoisHistoryRecords gets the whoisHistoryRecords property value. Retrieve details about whoisHistoryRecord objects.Note: List retrieval is not yet supported.
// returns a []WhoisHistoryRecordable when successful
func (m *ThreatIntelligence) GetWhoisHistoryRecords()([]WhoisHistoryRecordable) {
    return m.whoisHistoryRecords
}
// GetWhoisRecords gets the whoisRecords property value. A list of whoisRecord objects.
// returns a []WhoisRecordable when successful
func (m *ThreatIntelligence) GetWhoisRecords()([]WhoisRecordable) {
    return m.whoisRecords
}
// Serialize serializes information the current object
func (m *ThreatIntelligence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetArticleIndicators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetArticleIndicators()))
        for i, v := range m.GetArticleIndicators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("articleIndicators", cast)
        if err != nil {
            return err
        }
    }
    if m.GetArticles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetArticles()))
        for i, v := range m.GetArticles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("articles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostComponents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostComponents()))
        for i, v := range m.GetHostComponents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hostComponents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostCookies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostCookies()))
        for i, v := range m.GetHostCookies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hostCookies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostPairs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostPairs()))
        for i, v := range m.GetHostPairs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hostPairs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostPorts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostPorts()))
        for i, v := range m.GetHostPorts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hostPorts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHosts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHosts()))
        for i, v := range m.GetHosts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hosts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostSslCertificates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostSslCertificates()))
        for i, v := range m.GetHostSslCertificates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hostSslCertificates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostTrackers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostTrackers()))
        for i, v := range m.GetHostTrackers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("hostTrackers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIntelligenceProfileIndicators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIntelligenceProfileIndicators()))
        for i, v := range m.GetIntelligenceProfileIndicators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("intelligenceProfileIndicators", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIntelProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIntelProfiles()))
        for i, v := range m.GetIntelProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("intelProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPassiveDnsRecords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPassiveDnsRecords()))
        for i, v := range m.GetPassiveDnsRecords() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("passiveDnsRecords", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSslCertificates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSslCertificates()))
        for i, v := range m.GetSslCertificates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sslCertificates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubdomains() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubdomains()))
        for i, v := range m.GetSubdomains() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("subdomains", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVulnerabilities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVulnerabilities()))
        for i, v := range m.GetVulnerabilities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("vulnerabilities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWhoisHistoryRecords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWhoisHistoryRecords()))
        for i, v := range m.GetWhoisHistoryRecords() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("whoisHistoryRecords", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWhoisRecords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWhoisRecords()))
        for i, v := range m.GetWhoisRecords() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("whoisRecords", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArticleIndicators sets the articleIndicators property value. Refers to indicators of threat or compromise highlighted in an article.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetArticleIndicators(value []ArticleIndicatorable)() {
    m.articleIndicators = value
}
// SetArticles sets the articles property value. A list of article objects.
func (m *ThreatIntelligence) SetArticles(value []Articleable)() {
    m.articles = value
}
// SetHostComponents sets the hostComponents property value. Retrieve details about hostComponent objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetHostComponents(value []HostComponentable)() {
    m.hostComponents = value
}
// SetHostCookies sets the hostCookies property value. Retrieve details about hostCookie objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetHostCookies(value []HostCookieable)() {
    m.hostCookies = value
}
// SetHostPairs sets the hostPairs property value. Retrieve details about hostTracker objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetHostPairs(value []HostPairable)() {
    m.hostPairs = value
}
// SetHostPorts sets the hostPorts property value. Retrieve details about hostPort objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetHostPorts(value []HostPortable)() {
    m.hostPorts = value
}
// SetHosts sets the hosts property value. Refers to host objects that Microsoft Threat Intelligence has observed.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetHosts(value []Hostable)() {
    m.hosts = value
}
// SetHostSslCertificates sets the hostSslCertificates property value. Retrieve details about hostSslCertificate objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetHostSslCertificates(value []HostSslCertificateable)() {
    m.hostSslCertificates = value
}
// SetHostTrackers sets the hostTrackers property value. Retrieve details about hostTracker objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetHostTrackers(value []HostTrackerable)() {
    m.hostTrackers = value
}
// SetIntelligenceProfileIndicators sets the intelligenceProfileIndicators property value. The intelligenceProfileIndicators property
func (m *ThreatIntelligence) SetIntelligenceProfileIndicators(value []IntelligenceProfileIndicatorable)() {
    m.intelligenceProfileIndicators = value
}
// SetIntelProfiles sets the intelProfiles property value. A list of intelligenceProfile objects.
func (m *ThreatIntelligence) SetIntelProfiles(value []IntelligenceProfileable)() {
    m.intelProfiles = value
}
// SetPassiveDnsRecords sets the passiveDnsRecords property value. Retrieve details about passiveDnsRecord objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetPassiveDnsRecords(value []PassiveDnsRecordable)() {
    m.passiveDnsRecords = value
}
// SetSslCertificates sets the sslCertificates property value. Retrieve details about sslCertificate objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetSslCertificates(value []SslCertificateable)() {
    m.sslCertificates = value
}
// SetSubdomains sets the subdomains property value. Retrieve details about the subdomain.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetSubdomains(value []Subdomainable)() {
    m.subdomains = value
}
// SetVulnerabilities sets the vulnerabilities property value. Retrieve details about vulnerabilities.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetVulnerabilities(value []Vulnerabilityable)() {
    m.vulnerabilities = value
}
// SetWhoisHistoryRecords sets the whoisHistoryRecords property value. Retrieve details about whoisHistoryRecord objects.Note: List retrieval is not yet supported.
func (m *ThreatIntelligence) SetWhoisHistoryRecords(value []WhoisHistoryRecordable)() {
    m.whoisHistoryRecords = value
}
// SetWhoisRecords sets the whoisRecords property value. A list of whoisRecord objects.
func (m *ThreatIntelligence) SetWhoisRecords(value []WhoisRecordable)() {
    m.whoisRecords = value
}
type ThreatIntelligenceable interface {
    ib77c81ae8501035869703744ba13b6f711366c9348e33eae916d2aea3d8b34ba.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArticleIndicators()([]ArticleIndicatorable)
    GetArticles()([]Articleable)
    GetHostComponents()([]HostComponentable)
    GetHostCookies()([]HostCookieable)
    GetHostPairs()([]HostPairable)
    GetHostPorts()([]HostPortable)
    GetHosts()([]Hostable)
    GetHostSslCertificates()([]HostSslCertificateable)
    GetHostTrackers()([]HostTrackerable)
    GetIntelligenceProfileIndicators()([]IntelligenceProfileIndicatorable)
    GetIntelProfiles()([]IntelligenceProfileable)
    GetPassiveDnsRecords()([]PassiveDnsRecordable)
    GetSslCertificates()([]SslCertificateable)
    GetSubdomains()([]Subdomainable)
    GetVulnerabilities()([]Vulnerabilityable)
    GetWhoisHistoryRecords()([]WhoisHistoryRecordable)
    GetWhoisRecords()([]WhoisRecordable)
    SetArticleIndicators(value []ArticleIndicatorable)()
    SetArticles(value []Articleable)()
    SetHostComponents(value []HostComponentable)()
    SetHostCookies(value []HostCookieable)()
    SetHostPairs(value []HostPairable)()
    SetHostPorts(value []HostPortable)()
    SetHosts(value []Hostable)()
    SetHostSslCertificates(value []HostSslCertificateable)()
    SetHostTrackers(value []HostTrackerable)()
    SetIntelligenceProfileIndicators(value []IntelligenceProfileIndicatorable)()
    SetIntelProfiles(value []IntelligenceProfileable)()
    SetPassiveDnsRecords(value []PassiveDnsRecordable)()
    SetSslCertificates(value []SslCertificateable)()
    SetSubdomains(value []Subdomainable)()
    SetVulnerabilities(value []Vulnerabilityable)()
    SetWhoisHistoryRecords(value []WhoisHistoryRecordable)()
    SetWhoisRecords(value []WhoisRecordable)()
}
