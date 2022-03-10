package protocol

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func SignIn(req *LoginReq) []byte {
	pkg := make([]byte, 0)
	pkg = append(pkg, 0x68,
		0x22,
		0x00,
		0x00,
		0x00,
		0x01)
	pkg = append(pkg, getSN(req.sn)...)
	pkg = append(pkg, getHex(int(req.csType)),
		getHex(req.gunNumber),
		getTenMultiVersion(req.protocolVersion),
	)
	pkg = append(pkg, getAsciiToByte(req.programmingVersion)...)
	pkg = append(pkg, byte(req.netType))
	pkg = append(pkg, getSim(req.sim)...)
	pkg = append(pkg, byte(req.operator))
	return pkg
}

func getSN(sn string) []byte {
	step, i := convertStringToByte(sn)
	if i = 7 - i; i > 0 {
		addZero(&step, i)
	}
	return step
}

func getSim(sim string) []byte {
	step, i := convertStringToByte(sim)
	if i = 10 - i; i > 0 {
		addZero(&step, i)
	}
	return step
}

func getHex(i int) (hex byte) {
	t := fmt.Sprintf("%X", i)
	mid, _ := strconv.ParseInt(t, 16, 32)
	hex = byte(mid)
	return
}

func getTenMultiVersion(version string) (result byte) {
	str := strings.Split(version, "V")[1]
	fl, _ := strconv.ParseFloat(str, 64)
	result = getHex(int(fl * 10))
	return
}

func getAsciiToByte(value string) []byte {
	str := fmt.Sprintf("%X", []byte(value))
	step, i := convertStringToByte(str)
	if i = 8 - i; i > 0 {
		addZero(&step, i)
	}
	return step
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

func removeZero(b []byte) {
	bytes.TrimRight(b, "\x00")
}
