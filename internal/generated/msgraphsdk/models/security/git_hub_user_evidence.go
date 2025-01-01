package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GitHubUserEvidence struct {
    AlertEvidence
    // The email property
    email *string
    // The login property
    login *string
    // The name property
    name *string
    // The userId property
    userId *string
    // The webUrl property
    webUrl *string
}
// NewGitHubUserEvidence instantiates a new GitHubUserEvidence and sets the default values.
func NewGitHubUserEvidence()(*GitHubUserEvidence) {
    m := &GitHubUserEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.gitHubUserEvidence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGitHubUserEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGitHubUserEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGitHubUserEvidence(), nil
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *GitHubUserEvidence) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GitHubUserEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogin(val)
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
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *GitHubUserEvidence) GetLogin()(*string) {
    return m.login
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *GitHubUserEvidence) GetName()(*string) {
    return m.name
}
// GetUserId gets the userId property value. The userId property
// returns a *string when successful
func (m *GitHubUserEvidence) GetUserId()(*string) {
    return m.userId
}
// GetWebUrl gets the webUrl property value. The webUrl property
// returns a *string when successful
func (m *GitHubUserEvidence) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *GitHubUserEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("login", m.GetLogin())
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
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmail sets the email property value. The email property
func (m *GitHubUserEvidence) SetEmail(value *string)() {
    m.email = value
}
// SetLogin sets the login property value. The login property
func (m *GitHubUserEvidence) SetLogin(value *string)() {
    m.login = value
}
// SetName sets the name property value. The name property
func (m *GitHubUserEvidence) SetName(value *string)() {
    m.name = value
}
// SetUserId sets the userId property value. The userId property
func (m *GitHubUserEvidence) SetUserId(value *string)() {
    m.userId = value
}
// SetWebUrl sets the webUrl property value. The webUrl property
func (m *GitHubUserEvidence) SetWebUrl(value *string)() {
    m.webUrl = value
}
type GitHubUserEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetLogin()(*string)
    GetName()(*string)
    GetUserId()(*string)
    GetWebUrl()(*string)
    SetEmail(value *string)()
    SetLogin(value *string)()
    SetName(value *string)()
    SetUserId(value *string)()
    SetWebUrl(value *string)()
}
