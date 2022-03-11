package protocol

import (
	"bytes"
	"fmt"
	"strconv"
)

func getHex(i int) (hex byte) {
	t := fmt.Sprintf("%X", i)
	mid, _ := strconv.ParseInt(t, 16, 32)
	hex = byte(mid)
	return
}

func convertStringToByte(str string) ([]byte, int) {
	result := make([]byte, 0)
	j := 0
	for i, convert := 0, ""; i < len(str)-1; i += 2 {
		convert = str[i : i+2]
		bt, _ := strconv.ParseInt(convert, 16, 32)
		result = append(result, byte(bt))
		j++
	}
	return result, j
}

func addZero(b *[]byte, n int) {
	for i := 0; i < n; i++ {
		*b = append(*b, byte(0x00))
	}
}

func removeZero(b *[]byte) {
	bytes.TrimRight(*b, "\x00")
}

func encodeSN(sn string) []byte {
	step, i := convertStringToByte(sn)
	if i = 7 - i; i > 0 {
		addZero(&step, i)
	}
	return step
}
