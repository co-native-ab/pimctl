package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AppliedConditionalAccessPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Refers to the name of the conditional access policy (example: 'Require MFA for Salesforce').
    displayName *string
    // Refers to the grant controls enforced by the conditional access policy (example: 'Require multifactor authentication').
    enforcedGrantControls []string
    // Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
    enforcedSessionControls []string
    // An identifier of the conditional access policy. Supports $filter (eq).
    id *string
    // The OdataType property
    odataType *string
    // Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (policy isn't applied because policy conditions weren't met), notEnabled (This is due to the policy in a disabled state), unknown, unknownFutureValue, reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted.
    result *AppliedConditionalAccessPolicyResult
}
// NewAppliedConditionalAccessPolicy instantiates a new AppliedConditionalAccessPolicy and sets the default values.
func NewAppliedConditionalAccessPolicy()(*AppliedConditionalAccessPolicy) {
    m := &AppliedConditionalAccessPolicy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppliedConditionalAccessPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppliedConditionalAccessPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppliedConditionalAccessPolicy(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AppliedConditionalAccessPolicy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Refers to the name of the conditional access policy (example: 'Require MFA for Salesforce').
// returns a *string when successful
func (m *AppliedConditionalAccessPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnforcedGrantControls gets the enforcedGrantControls property value. Refers to the grant controls enforced by the conditional access policy (example: 'Require multifactor authentication').
// returns a []string when successful
func (m *AppliedConditionalAccessPolicy) GetEnforcedGrantControls()([]string) {
    return m.enforcedGrantControls
}
// GetEnforcedSessionControls gets the enforcedSessionControls property value. Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
// returns a []string when successful
func (m *AppliedConditionalAccessPolicy) GetEnforcedSessionControls()([]string) {
    return m.enforcedSessionControls
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppliedConditionalAccessPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["enforcedGrantControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetEnforcedGrantControls(res)
        }
        return nil
    }
    res["enforcedSessionControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetEnforcedSessionControls(res)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppliedConditionalAccessPolicyResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(*AppliedConditionalAccessPolicyResult))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. An identifier of the conditional access policy. Supports $filter (eq).
// returns a *string when successful
func (m *AppliedConditionalAccessPolicy) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AppliedConditionalAccessPolicy) GetOdataType()(*string) {
    return m.odataType
}
// GetResult gets the result property value. Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (policy isn't applied because policy conditions weren't met), notEnabled (This is due to the policy in a disabled state), unknown, unknownFutureValue, reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted.
// returns a *AppliedConditionalAccessPolicyResult when successful
func (m *AppliedConditionalAccessPolicy) GetResult()(*AppliedConditionalAccessPolicyResult) {
    return m.result
}
// Serialize serializes information the current object
func (m *AppliedConditionalAccessPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcedGrantControls() != nil {
        err := writer.WriteCollectionOfStringValues("enforcedGrantControls", m.GetEnforcedGrantControls())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcedSessionControls() != nil {
        err := writer.WriteCollectionOfStringValues("enforcedSessionControls", m.GetEnforcedSessionControls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
    if m.GetResult() != nil {
        cast := (*m.GetResult()).String()
        err := writer.WriteStringValue("result", &cast)
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
func (m *AppliedConditionalAccessPolicy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Refers to the name of the conditional access policy (example: 'Require MFA for Salesforce').
func (m *AppliedConditionalAccessPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnforcedGrantControls sets the enforcedGrantControls property value. Refers to the grant controls enforced by the conditional access policy (example: 'Require multifactor authentication').
func (m *AppliedConditionalAccessPolicy) SetEnforcedGrantControls(value []string)() {
    m.enforcedGrantControls = value
}
// SetEnforcedSessionControls sets the enforcedSessionControls property value. Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
func (m *AppliedConditionalAccessPolicy) SetEnforcedSessionControls(value []string)() {
    m.enforcedSessionControls = value
}
// SetId sets the id property value. An identifier of the conditional access policy. Supports $filter (eq).
func (m *AppliedConditionalAccessPolicy) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppliedConditionalAccessPolicy) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResult sets the result property value. Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (policy isn't applied because policy conditions weren't met), notEnabled (This is due to the policy in a disabled state), unknown, unknownFutureValue, reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted.
func (m *AppliedConditionalAccessPolicy) SetResult(value *AppliedConditionalAccessPolicyResult)() {
    m.result = value
}
type AppliedConditionalAccessPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetEnforcedGrantControls()([]string)
    GetEnforcedSessionControls()([]string)
    GetId()(*string)
    GetOdataType()(*string)
    GetResult()(*AppliedConditionalAccessPolicyResult)
    SetDisplayName(value *string)()
    SetEnforcedGrantControls(value []string)()
    SetEnforcedSessionControls(value []string)()
    SetId(value *string)()
    SetOdataType(value *string)()
    SetResult(value *AppliedConditionalAccessPolicyResult)()
}
