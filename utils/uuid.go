package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateDeviceId() string {
	var b [16]byte
	_, _ = rand.Read(b[:])

	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

var DeviceId = GenerateDeviceId()
