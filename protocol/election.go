package protocol

import (
	"fmt"
	"strconv"
	"strings"
)

type CSType int

const (
	// 直流
	DC CSType = 0
	// 交流
	AC CSType = 1
)

type NetType byte

const (
	SIM       NetType = 0x00
	LAN       NetType = 0x01
	WAN       NetType = 0x02
	NET_OTHER NetType = 0x03
)

type OperatorType byte

const (
	MOBILE         OperatorType = 0x00
	TELECOM        OperatorType = 0x02
	UNICOM         OperatorType = 0x03
	OPERATOR_OTHER OperatorType = 0x04
)

type LoginReq struct {
	id int
	// 桩编码
	sn string
	// 桩类型
	csType CSType
	// 充电枪数量
	gunNumber int
	// 通信协议版本
	protocolVersion string
	// 程序版本
	programmingVersion string
	// 网络链接类型
	netType NetType
	// sim卡
	sim string
	// 运营商类型
	operator OperatorType
}

func NewLoginReq() *LoginReq {
	return &LoginReq{id: 1,
		sn:                 "55031412782305",
		csType:             DC,
		gunNumber:          2,
		protocolVersion:    "V1.0",
		programmingVersion: "v1.0.1",
		netType:            SIM,
		sim:                "01010101010101010101",
		operator:           TELECOM}
}

func (l *LoginReq) getSN() []byte {
	step, i := convertStringToByte(l.sn)
	if i = 7 - i; i > 0 {
		addZero(&step, i)
	}
	return step
}

func (l *LoginReq) getCSType() byte {
	return getHex(int(l.csType))
}

func (l *LoginReq) getGunNumber() byte {
	return getHex(l.gunNumber)
}

func (l *LoginReq) getTenMultiVersion() byte {
	str := strings.Split(l.protocolVersion, "V")[1]
	fl, _ := strconv.ParseFloat(str, 64)
	return getHex(int(fl * 10))
}

func (l *LoginReq) getAsciiToByte() []byte {
	str := fmt.Sprintf("%X", []byte(l.programmingVersion))
	step, i := convertStringToByte(str)
	if i = 8 - i; i > 0 {
		addZero(&step, i)
	}
	return step
}

func (l *LoginReq) getNetType() byte {
	return byte(l.netType)
}

func (l *LoginReq) getSim() []byte {
	step, i := convertStringToByte(l.sim)
	if i = 10 - i; i > 0 {
		addZero(&step, i)
	}
	return step
}

func (l *LoginReq) getOperator() byte {
	return byte(l.operator)
}
