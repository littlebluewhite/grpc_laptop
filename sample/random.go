package sample

import (
	"github.com/google/uuid"
	"grpc_test/pb/message"
	"math/rand"
)

func randomKeyboardLayout() message.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return message.Keyboard_QWERTY
	case 2:
		return message.Keyboard_QWERTZ
	default:
		return message.Keyboard_AZERTY
	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	} else {
		return randomStringFromSet(
			"Ryzen 7 Pro 2700U",
			"Ryzen 5 PRO 3500U",
			"Ryzen 3 PRO 3200GE",
		)
	}
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	} else {
		return randomStringFromSet(
			"RX 590",
			"RX580",
			"RX 5700-XT",
			"RX Vega-56",
		)
	}
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad x1", "Thinkpad P1", "Thinkpad P53")
	}
}

func randomScreenResolution() *message.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &message.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return resolution
}

func randomScreenPanel() message.Screen_Panel {
	if rand.Intn(2) == 1 {
		return message.Screen_IPS
	} else {
		return message.Screen_OLED
	}
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat[T float64 | float32](min, max T) T {
	return min + T(rand.Float64())*(max-min)
}

func randomID() string {
	return uuid.New().String()
}
