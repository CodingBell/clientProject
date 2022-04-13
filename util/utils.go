package util

import (
	"fmt"
	"github.com/Kotodian/gokit/ac/lib"
	"strconv"
	"strings"
)

// String2Ascii change string to ascii
func String2Ascii(s string) {
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		result := int(s[i])
		fmt.Printf("%X", result)
	}
}

// Ascii2String change ascii to string
func Ascii2String(by []byte) string {
	var s string
	for _, v := range by {
		s += string(v)
	}
	return s
}

// StringAscii2String input a ascii string and return the mapping string
func StringAscii2String(s string) string {
	s = TrimRight(s)
	by := String2Bytes(s)
	return Ascii2String(by)
}

// String2Bytes change string to bytes
func String2Bytes(s string) []byte {
	if len(s)%2 != 0 {
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

// addHeadAndTail is a common function to add head and tail to the string
func addHeadAndTail(ty byte, b []byte) []byte {
	by := []byte{0x68, 0x00, 0x01, 1 << 6, 89, 0x00, 0x00, ty, 0x00, 0x01}
	by = append(by, b...)
	crc := lib.CheckSum(by)
	by = append(by, crc...)
	return by
}

// GetStringLength return the length of the input string
func GetStringLength(s string) {
	fmt.Println(len(s))
}

// GenerateBytes generate all FF string
func GenerateBytes(length int) string {
	var str string
	for i := 0; i < length; i++ {
		str += "FF"
	}
	return str
}

// TrimRight return a string without FF
func TrimRight(s string) string {
	s = strings.Trim(s, "FF")
	return s
}

func Int2Bytes(i int) []byte {
	return []byte{byte(i >> 8), byte(i)}
}
