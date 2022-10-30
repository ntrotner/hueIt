package helpers

import "tinygo.org/x/bluetooth"

func MustParse(inputUUID string) bluetooth.UUID {
	uuid, _ := bluetooth.ParseUUID(inputUUID)
	return uuid
}
