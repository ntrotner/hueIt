package bulb_handler

import (
	"hue-it/helpers"
	"log"
	"os"

	"tinygo.org/x/bluetooth"
)

type Bulb struct {
	UUID              bluetooth.UUID
	Connection        *bluetooth.Device
	Characteristics   AvailableCharacteristics
	SetBrightnessFunc func(*Bulb, uint16) bool
	SetColorFunc      func(*Bulb, []byte) bool
	SetOnFunc         func(*Bulb, bool) bool
}

type CharacteristicEndpoint struct {
	Value []byte
	UUID  bluetooth.UUID
}

type AvailableCharacteristics struct {
	On CharacteristicEndpoint

	Brightness CharacteristicEndpoint

	Color CharacteristicEndpoint
}

func (b *Bulb) SetBrightness(level uint16) bool {
	log.Println("SetBrightness to", level)
	return b.SetBrightnessFunc(b, level)
}

func (b *Bulb) GetWritableCharacteristic(characteristic CharacteristicEndpoint) *bluetooth.DeviceCharacteristic {
	writeableServices, err := b.Connection.DiscoverServices([]bluetooth.UUID{b.UUID})

	if err != nil || len(writeableServices) != 1 {
		os.Exit(helpers.BULB_MISSING_SERVICE_UUID)
	}

	writeableCharacteristics, err := writeableServices[0].DiscoverCharacteristics([]bluetooth.UUID{characteristic.UUID})
	if err != nil && len(writeableCharacteristics) != 1 {
		os.Exit(helpers.BULB_MISSING_CHARACTERISTIC_UUID)
	}
	log.Println("Found Characteristic:", writeableCharacteristics[0].String())
	return &(writeableCharacteristics[0])
}
