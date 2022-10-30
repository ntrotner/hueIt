package devices

import (
	"encoding/binary"
	bulb_handler "hue-it/bulb-handler"
	"log"
)

func SetBrightnessDefault(bulb *bulb_handler.Bulb, level uint16) bool {
	if level <= 0 || level >= 100 {
		log.Println("Invalid Brightness: Expected 1..99 got", level)
		return false
	}

	characteristic := bulb.GetWritableCharacteristic(bulb.Characteristics.Brightness)

	binary.LittleEndian.AppendUint16([]byte{}, level&0xff)
	_, err := characteristic.WriteWithoutResponse([]byte{})

	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
