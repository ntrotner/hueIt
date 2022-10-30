package hue_filament_bulb

import (
	bulb_handler "hue-it/bulb-handler"
	"hue-it/helpers"
)

func GetServices() bulb_handler.AvailableCharacteristics {
	return bulb_handler.AvailableCharacteristics{
		On: bulb_handler.CharacteristicEndpoint{
			UUID:  helpers.MustParse("932c32bd-0002-47a2-835a-a8d455b859dd"),
			Value: []byte{1},
		},
		Brightness: bulb_handler.CharacteristicEndpoint{
			UUID:  helpers.MustParse("932c32bd-0003-47a2-835a-a8d455b859dd"),
			Value: []byte{254},
		},
		Color: bulb_handler.CharacteristicEndpoint{
			UUID:  helpers.MustParse("932c32bd-0004-47a2-835a-a8d455b859dd"),
			Value: []byte{200, 1},
		},
	}
}
