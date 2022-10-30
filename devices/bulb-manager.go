package devices

import (
	bulb_handler "hue-it/bulb-handler"
	hue_filament_bulb "hue-it/devices/hue-filament-bulb"
	"hue-it/helpers"

	"tinygo.org/x/bluetooth"
)

func GetAllBulbs() *map[string]bulb_handler.Bulb {
	var supportedBulbs = map[string]bulb_handler.Bulb{
		"E2:A0:E1:C4:62:00": buildBulb(
			"932c32bd-0000-47a2-835a-a8d455b859dd",
			&bluetooth.Device{},
			hue_filament_bulb.GetServices(),
		),
	}
	return &supportedBulbs
}

func buildBulb(inputUUID string, connection *bluetooth.Device, characteristics bulb_handler.AvailableCharacteristics) bulb_handler.Bulb {
	return bulb_handler.Bulb{
		UUID:            helpers.MustParse(inputUUID),
		Connection:      connection,
		Characteristics: hue_filament_bulb.GetServices(),
	}
}
