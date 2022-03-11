package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var t int = 0xA
	//k := byte(t)
	//fmt.Printf("%X", k)
	//result := fmt.Sprintf("%04X", 0x04)
	//resultj, j := convertStringToByte(result)
	//fmt.Printf("%X %d", resultj, j)

	var tr byte = 0xA
	fmt.Println(int(tr))

	//pkg := []byte{
	//	0x68,
	//	0x0C,
	//	0x00,
	//	0x00,
	//	0x00,
	//	0x02,
	//	0x55,
	//	0x03,
	//	0x14,
	//	0x12,
	//	0x78,
	//	0x23,
	//	0x05,
	//	0x00,
	//	0xDA,
	//	0x4C,
	//}
	//seq := protocol.NewLoginReq()
	//tcp.SelectRequestMethod(seq)
	//fmt.Printf("%X", protocol.SignIn(seq))
}
func convertStringToByte(str string) ([]byte, int) {
	result := make([]byte, 0)
	j := 0
	for i, convert := 0, ""; i < len(str); i += 2 {
		convert = str[i : i+2]
		bt, _ := strconv.ParseInt(convert, 16, 32)
		result = append(result, byte(bt))
		j++
	}
	return result, j
}
