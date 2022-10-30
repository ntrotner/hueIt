package bulb_handler

import (
	"hue-it/helpers"
	"os"

	"tinygo.org/x/bluetooth"
)

func InitializeBulbStream(bulb *Bulb) {
	authenticate(bulb.Connection)
	setBrightness(bulb)
}

func authenticate(connection *bluetooth.Device) {
	writeableServices, err := connection.DiscoverServices([]bluetooth.UUID{helpers.MustParse("b8843add-0000-4aa1-8794-c3f462030bda")})

	if err != nil && len(writeableServices) != 1 {
		os.Exit(helpers.BULB_MISSING_SERVICE_UUID)
	}

	writeableCharacteristics, err := writeableServices[0].DiscoverCharacteristics([]bluetooth.UUID{helpers.MustParse("b8843add-0002-4aa1-8794-c3f462030bda")})
	if err != nil && len(writeableCharacteristics) != 1 {
		os.Exit(helpers.BULB_MISSING_CHARACTERISTIC_UUID)
	}
	output, err := writeableCharacteristics[0].WriteWithoutResponse([]byte{128, 0, 6, 21, 0, 1, 20, 1, 11, 16, 0, 0})

	println(output)
	println(err.Error())
}

func getWritableCharacteristic(connection *bluetooth.Device, uuid *bluetooth.UUID, characteristic CharacteristicEndpoint) bluetooth.DeviceCharacteristic {
	writeableServices, err := connection.DiscoverServices([]bluetooth.UUID{*uuid})

	if err != nil && len(writeableServices) != 1 {
		os.Exit(helpers.BULB_MISSING_SERVICE_UUID)
	}

	writeableCharacteristics, err := writeableServices[0].DiscoverCharacteristics([]bluetooth.UUID{characteristic.UUID})
	if err != nil && len(writeableCharacteristics) != 1 {
		os.Exit(helpers.BULB_MISSING_CHARACTERISTIC_UUID)
	}

	return writeableCharacteristics[0]
}

func setBrightness(bulb *Bulb) {
	writeableCharacteristic := getWritableCharacteristic(bulb.Connection, &bulb.UUID, bulb.Characteristics.Brightness)

	writeableCharacteristic.EnableNotifications(func(value []byte) {
		println(value)
	})

	output, err := writeableCharacteristic.WriteWithoutResponse(bulb.Characteristics.Brightness.Value)

	println(output)
	println(err.Error())
}
