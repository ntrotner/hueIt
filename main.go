package main

import (
	bluetooth_handler "hue-it/bluetooth-handler"
	bulb_manager "hue-it/devices"
)

func main() {
	bulbs := bulb_manager.GetAllBulbs()
	bluetooth_handler.InitializeBluetoothHandler(bulbs)
}
