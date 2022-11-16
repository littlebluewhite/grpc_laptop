package sample

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc_test/pb/message"
)

func NewKeyboard() *message.Keyboard {
	keyboard := &message.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCPU() *message.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGhz := randomFloat(2.0, 3.5)
	maxGhz := randomFloat(minGhz, 5.0)
	cpu := &message.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func NewGPU() *message.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat(1.0, 1.5)
	maxGhz := randomFloat(minGhz, 2.0)

	memory := &message.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  message.Memory_GIGABYTE,
	}
	gpu := &message.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}

func NewRAM() *message.Memory {
	ram := &message.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  message.Memory_GIGABYTE,
	}
	return ram
}

func NewSSD() *message.Storage {
	ssd := &message.Storage{
		Driver: message.Storage_SDD,
		Memory: &message.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  message.Memory_GIGABYTE,
		},
	}
	return ssd
}

func NewHDD() *message.Storage {
	hdd := &message.Storage{
		Driver: message.Storage_HDD,
		Memory: &message.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  message.Memory_TERABYTE,
		},
	}
	return hdd
}

func NewScreen() *message.Screen {
	screen := &message.Screen{
		SizeInch:   randomFloat[float32](13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
	return screen
}

func NewLaptop() *message.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &message.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*message.GPU{NewGPU()},
		Storages: []*message.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &message.Laptop_WeightKg{
			WeightKg: randomFloat(1.0, 3.0),
		},
		PriceUsd:    randomFloat[float64](1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdateAt:    timestamppb.Now(),
	}
	return laptop
}
