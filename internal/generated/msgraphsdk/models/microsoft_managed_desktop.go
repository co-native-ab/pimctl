package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MicrosoftManagedDesktop struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates the provisioning policy associated with Microsoft Managed Desktop settings. Possible values are: notManaged, premiumManaged, standardManaged, starterManaged, unknownFutureValue. The default is notManaged.
    managedType *MicrosoftManagedDesktopType
    // The OdataType property
    odataType *string
    // The name of the Microsoft Managed Desktop profile that the Windows 365 Cloud PC is associated with.
    profile *string
}
// NewMicrosoftManagedDesktop instantiates a new MicrosoftManagedDesktop and sets the default values.
func NewMicrosoftManagedDesktop()(*MicrosoftManagedDesktop) {
    m := &MicrosoftManagedDesktop{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftManagedDesktopFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMicrosoftManagedDesktopFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftManagedDesktop(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MicrosoftManagedDesktop) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MicrosoftManagedDesktop) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftManagedDesktopType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedType(val.(*MicrosoftManagedDesktopType))
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
    res["profile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfile(val)
        }
        return nil
    }
    return res
}
// GetManagedType gets the managedType property value. Indicates the provisioning policy associated with Microsoft Managed Desktop settings. Possible values are: notManaged, premiumManaged, standardManaged, starterManaged, unknownFutureValue. The default is notManaged.
// returns a *MicrosoftManagedDesktopType when successful
func (m *MicrosoftManagedDesktop) GetManagedType()(*MicrosoftManagedDesktopType) {
    return m.managedType
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *MicrosoftManagedDesktop) GetOdataType()(*string) {
    return m.odataType
}
// GetProfile gets the profile property value. The name of the Microsoft Managed Desktop profile that the Windows 365 Cloud PC is associated with.
// returns a *string when successful
func (m *MicrosoftManagedDesktop) GetProfile()(*string) {
    return m.profile
}
// Serialize serializes information the current object
func (m *MicrosoftManagedDesktop) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManagedType() != nil {
        cast := (*m.GetManagedType()).String()
        err := writer.WriteStringValue("managedType", &cast)
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
    {
        err := writer.WriteStringValue("profile", m.GetProfile())
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
func (m *MicrosoftManagedDesktop) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetManagedType sets the managedType property value. Indicates the provisioning policy associated with Microsoft Managed Desktop settings. Possible values are: notManaged, premiumManaged, standardManaged, starterManaged, unknownFutureValue. The default is notManaged.
func (m *MicrosoftManagedDesktop) SetManagedType(value *MicrosoftManagedDesktopType)() {
    m.managedType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MicrosoftManagedDesktop) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProfile sets the profile property value. The name of the Microsoft Managed Desktop profile that the Windows 365 Cloud PC is associated with.
func (m *MicrosoftManagedDesktop) SetProfile(value *string)() {
    m.profile = value
}
type MicrosoftManagedDesktopable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManagedType()(*MicrosoftManagedDesktopType)
    GetOdataType()(*string)
    GetProfile()(*string)
    SetManagedType(value *MicrosoftManagedDesktopType)()
    SetOdataType(value *string)()
    SetProfile(value *string)()
}
