package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"grpc_test/pb/message"
	"sync"
)

// ErrAlreadyExists is returned when a record with the same ID already exists in the Store
var ErrAlreadyExists = errors.New("record already exists")

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	Save(laptop *message.Laptop) error
}

type InMemoryLaptopStore struct {
	mutex sync.Mutex
	data  map[string]*message.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*message.Laptop),
	}
}

// Save saves the laptop to the store
func (store *InMemoryLaptopStore) Save(laptop *message.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// deep copy
	other := &message.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	store.data[other.Id] = other
	return nil
}
