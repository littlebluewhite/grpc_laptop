package service

import (
	"context"
	"errors"
	"grpc_test/pb/message"
	"grpc_test/util"
	"log"
	"sync"
)

// ErrAlreadyExists is returned when a record with the same ID already exists in the Store
var ErrAlreadyExists = errors.New("record already exists")

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	// Save saves the laptop to th store
	Save(laptop *message.Laptop) error
	// Find finds a laptop by ID
	Find(id string) (*message.Laptop, error)
	// Search searches for laptops with filter, returns one by one via the found function
	Search(ctx context.Context, filter *message.Filter, found func(laptop *message.Laptop) error) error
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
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
	other, err := util.DeepCopy(laptop)
	if err != nil {
		return err
	}

	store.data[other.Id] = other
	return nil
}

func (store *InMemoryLaptopStore) Find(id string) (*message.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	// deep copy
	return util.DeepCopy(laptop)
}

func (store *InMemoryLaptopStore) Search(
	ctx context.Context,
	filter *message.Filter,
	found func(laptop *message.Laptop) error,
) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	for _, laptop := range store.data {
		// heavy processing
		//time.Sleep(time.Second)
		log.Print("checking laptop id: ", laptop.Id)

		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Print("context is canceled")
			return errors.New("context is canceled")
		}

		if isQualified(filter, laptop) {
			other, err := util.DeepCopy(laptop)
			if err != nil {
				return err
			}
			err = found(other)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isQualified(filter *message.Filter, laptop *message.Laptop) bool {
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd() {
		return false
	}

	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}

	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}

	if toBit(laptop.GetRam()) < toBit(filter.GetMinRam()) {
		return false
	}

	return true
}

func toBit(memory *message.Memory) uint64 {
	value := memory.GetValue()

	switch memory.GetUnit() {
	case message.Memory_BIT:
		return value
	case message.Memory_BYTE:
		return value << 3 // 8 = 2^3
	case message.Memory_KILOBYTE:
		return value << 13 // 1024 * 8 = 2^10 * 2^3 = 2^13
	case message.Memory_MEGABYTE:
		return value << 23 // 1024 * 1024 * 8 = 2^20 * 2^10 * 2^3 = 2^23
	case message.Memory_GIGABYTE:
		return value << 33 // 1024 * 1024 * 1024 * 8 = 2^30 * 2^20 * 2^10 * 2^3 = 2^33
	case message.Memory_TERABYTE:
		return value << 43 // 1024 * 1024 * 1024 * 1024 * 8 = 2^40 * 2^30 * 2^20 * 2^10 * 2^3 = 2^43
	default:
		return 0
	}
}
