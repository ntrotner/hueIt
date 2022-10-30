package bulb_handler

import (
	"tinygo.org/x/bluetooth"
)

type Bulb struct {
	UUID            bluetooth.UUID
	Connection      *bluetooth.Device
	Characteristics AvailableCharacteristics
}

type CharacteristicEndpoint struct {
	Value []byte
	UUID  bluetooth.UUID
}

type AvailableCharacteristics struct {
	On         CharacteristicEndpoint
	Brightness CharacteristicEndpoint
	Color      CharacteristicEndpoint
}
