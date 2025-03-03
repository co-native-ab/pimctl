package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FileStorage struct {
    Entity
    // The containers property
    containers []FileStorageContainerable
    // The deletedContainers property
    deletedContainers []FileStorageContainerable
}
// NewFileStorage instantiates a new FileStorage and sets the default values.
func NewFileStorage()(*FileStorage) {
    m := &FileStorage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateFileStorageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorage(), nil
}
// GetContainers gets the containers property value. The containers property
// returns a []FileStorageContainerable when successful
func (m *FileStorage) GetContainers()([]FileStorageContainerable) {
    return m.containers
}
// GetDeletedContainers gets the deletedContainers property value. The deletedContainers property
// returns a []FileStorageContainerable when successful
func (m *FileStorage) GetDeletedContainers()([]FileStorageContainerable) {
    return m.deletedContainers
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FileStorage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["containers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFileStorageContainerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileStorageContainerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FileStorageContainerable)
                }
            }
            m.SetContainers(res)
        }
        return nil
    }
    res["deletedContainers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFileStorageContainerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileStorageContainerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FileStorageContainerable)
                }
            }
            m.SetDeletedContainers(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *FileStorage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContainers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContainers()))
        for i, v := range m.GetContainers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("containers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedContainers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedContainers()))
        for i, v := range m.GetDeletedContainers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deletedContainers", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContainers sets the containers property value. The containers property
func (m *FileStorage) SetContainers(value []FileStorageContainerable)() {
    m.containers = value
}
// SetDeletedContainers sets the deletedContainers property value. The deletedContainers property
func (m *FileStorage) SetDeletedContainers(value []FileStorageContainerable)() {
    m.deletedContainers = value
}
type FileStorageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContainers()([]FileStorageContainerable)
    GetDeletedContainers()([]FileStorageContainerable)
    SetContainers(value []FileStorageContainerable)()
    SetDeletedContainers(value []FileStorageContainerable)()
}
