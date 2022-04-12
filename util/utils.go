package util

import (
	"fmt"
	"github.com/Kotodian/gokit/ac/lib"
	"strconv"
)

// String2Ascii change string to ascii
func String2Ascii(s string) {
	for i := 0; i < len(s); i++ {
		result := int(s[i])
		fmt.Printf("%X", result)
	}
}

// String2Bytes change string to bytes
func String2Bytes(s string) []byte {
	if len(s)%2 != 0 {
		fmt.Println(len(s))
		return nil
	}
	result := make([]byte, 0)
	for i := 0; i < len(s); i += 2 {
		temp := s[i : i+2]
		num, _ := strconv.ParseInt(temp, 16, 32)
		result = append(result, byte(num))
	}
	return result
}

// GetBytes generate bytes from string
func GetBytes(str []string) []byte {
	result := make([]byte, 0)
	for _, s := range str {
		result = append(result, String2Bytes(s)...)
	}
	return result
}

func addHeadAndTail(ty byte, b []byte) {
	by := []byte{0x68, 0x00, 0x01, 1 << 6, 89, 0x00, 0x00, ty, 0x00, 0x01}
	by = append(by, b...)
	crc := lib.CheckSum(by)
	by = append(by, crc...)
}

// BootNotificationRequest return the necessary bytes
func BootNotificationRequest() []byte {
	str := []string{
		"5431363431373335323131",
		"4A4F59534F4EFFFFFFFF",
		"4A4F59534F4EFFFFFFFF",
		"76312E302E31FFFF",
		"3839383630335959584D48484858585858585850",
		"303132333435363738393031323334",
		"00",
		"7F000001",
		"31323334353637383930",
	}

	result := GetBytes(str)

	addHeadAndTail(0x01, result)
	return result
}
