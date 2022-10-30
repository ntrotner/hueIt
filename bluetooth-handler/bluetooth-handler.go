package bluetooth_handler

import (
	bulb_handler "hue-it/bulb-handler"
	helpers "hue-it/helpers"
	"log"
	"os"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func InitializeBluetoothHandler(bulbs *map[string]bulb_handler.Bulb) {
	log.Println("Enable Bluetooth Adapter")
	err := adapter.Enable()

	if err != nil {
		os.Exit(helpers.BLUETOOTH_CLIENT_FAILED)
	}

	log.Println("Start Scanning...")
	err = adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		bulb, foundBulb := (*bulbs)[device.Address.String()]

		if foundBulb {
			log.Print("Found Supported Bulb")
			adapter.StopScan()
			log.Println(" - Stopped Scanning")
			bulb_handler.InitializeBulbStream(connectToBulb(&device, &bulb))
		}
	})

	if err != nil {
		os.Exit(helpers.BLUETOOTH_CLIENT_FAILED)
	}
}

func connectToBulb(device *bluetooth.ScanResult, bulb *bulb_handler.Bulb) *bulb_handler.Bulb {
	connectedDevice, err := adapter.Connect(device.Address, bluetooth.ConnectionParams{})

	if err != nil {
		os.Exit(helpers.BLUETOOTH_CONNECTION_FAILED)
	}

	log.Println("Established Connection to:", device.LocalName(), "-", device.Address)
	bulb.Connection = connectedDevice
	return bulb
}
