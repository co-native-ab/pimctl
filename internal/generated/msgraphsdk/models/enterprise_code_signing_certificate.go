package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseCodeSigningCertificate struct {
    Entity
    // The Windows Enterprise Code-Signing Certificate in the raw data format. Set to null once certificate has been uploaded and other properties have been populated.
    content []byte
    // The cert expiration date and time (using ISO 8601 format, in UTC time). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The issuer value for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
    issuer *string
    // The issuer name for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
    issuerName *string
    // The status property
    status *CertificateStatus
    // The subject value for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
    subject *string
    // The subject name for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
    subjectName *string
    // The date time of CodeSigning Cert when it is uploaded (using ISO 8601 format, in UTC time). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
    uploadDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewEnterpriseCodeSigningCertificate instantiates a new EnterpriseCodeSigningCertificate and sets the default values.
func NewEnterpriseCodeSigningCertificate()(*EnterpriseCodeSigningCertificate) {
    m := &EnterpriseCodeSigningCertificate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEnterpriseCodeSigningCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseCodeSigningCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseCodeSigningCertificate(), nil
}
// GetContent gets the content property value. The Windows Enterprise Code-Signing Certificate in the raw data format. Set to null once certificate has been uploaded and other properties have been populated.
// returns a []byte when successful
func (m *EnterpriseCodeSigningCertificate) GetContent()([]byte) {
    return m.content
}
// GetExpirationDateTime gets the expirationDateTime property value. The cert expiration date and time (using ISO 8601 format, in UTC time). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
// returns a *Time when successful
func (m *EnterpriseCodeSigningCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseCodeSigningCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["issuerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CertificateStatus))
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["subjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectName(val)
        }
        return nil
    }
    res["uploadDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadDateTime(val)
        }
        return nil
    }
    return res
}
// GetIssuer gets the issuer property value. The issuer value for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
// returns a *string when successful
func (m *EnterpriseCodeSigningCertificate) GetIssuer()(*string) {
    return m.issuer
}
// GetIssuerName gets the issuerName property value. The issuer name for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
// returns a *string when successful
func (m *EnterpriseCodeSigningCertificate) GetIssuerName()(*string) {
    return m.issuerName
}
// GetStatus gets the status property value. The status property
// returns a *CertificateStatus when successful
func (m *EnterpriseCodeSigningCertificate) GetStatus()(*CertificateStatus) {
    return m.status
}
// GetSubject gets the subject property value. The subject value for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
// returns a *string when successful
func (m *EnterpriseCodeSigningCertificate) GetSubject()(*string) {
    return m.subject
}
// GetSubjectName gets the subjectName property value. The subject name for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
// returns a *string when successful
func (m *EnterpriseCodeSigningCertificate) GetSubjectName()(*string) {
    return m.subjectName
}
// GetUploadDateTime gets the uploadDateTime property value. The date time of CodeSigning Cert when it is uploaded (using ISO 8601 format, in UTC time). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
// returns a *Time when successful
func (m *EnterpriseCodeSigningCertificate) GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.uploadDateTime
}
// Serialize serializes information the current object
func (m *EnterpriseCodeSigningCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuerName", m.GetIssuerName())
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
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectName", m.GetSubjectName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("uploadDateTime", m.GetUploadDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The Windows Enterprise Code-Signing Certificate in the raw data format. Set to null once certificate has been uploaded and other properties have been populated.
func (m *EnterpriseCodeSigningCertificate) SetContent(value []byte)() {
    m.content = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The cert expiration date and time (using ISO 8601 format, in UTC time). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
func (m *EnterpriseCodeSigningCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetIssuer sets the issuer property value. The issuer value for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
func (m *EnterpriseCodeSigningCertificate) SetIssuer(value *string)() {
    m.issuer = value
}
// SetIssuerName sets the issuerName property value. The issuer name for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
func (m *EnterpriseCodeSigningCertificate) SetIssuerName(value *string)() {
    m.issuerName = value
}
// SetStatus sets the status property value. The status property
func (m *EnterpriseCodeSigningCertificate) SetStatus(value *CertificateStatus)() {
    m.status = value
}
// SetSubject sets the subject property value. The subject value for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
func (m *EnterpriseCodeSigningCertificate) SetSubject(value *string)() {
    m.subject = value
}
// SetSubjectName sets the subjectName property value. The subject name for the cert. This might contain information such as country (C), state or province (S), locality (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
func (m *EnterpriseCodeSigningCertificate) SetSubjectName(value *string)() {
    m.subjectName = value
}
// SetUploadDateTime sets the uploadDateTime property value. The date time of CodeSigning Cert when it is uploaded (using ISO 8601 format, in UTC time). Uploading a valid cert file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top, $OrderBy, $skip. $Search is not supported.
func (m *EnterpriseCodeSigningCertificate) SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.uploadDateTime = value
}
type EnterpriseCodeSigningCertificateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIssuer()(*string)
    GetIssuerName()(*string)
    GetStatus()(*CertificateStatus)
    GetSubject()(*string)
    GetSubjectName()(*string)
    GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetContent(value []byte)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIssuer(value *string)()
    SetIssuerName(value *string)()
    SetStatus(value *CertificateStatus)()
    SetSubject(value *string)()
    SetSubjectName(value *string)()
    SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
