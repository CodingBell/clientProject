package protocol

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// 登录认证请求
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

func NewLoginReq(id int,
	sn string,
	csType CSType,
	gunNumber int,
	protocolVersion string,
	programmingVersion string,
	netType NetType,
	sim string,
	operator OperatorType) *LoginReq {
	return &LoginReq{id: id,
		sn:                 sn,
		csType:             csType,
		gunNumber:          gunNumber,
		protocolVersion:    protocolVersion,
		programmingVersion: programmingVersion,
		netType:            netType,
		sim:                sim,
		operator:           operator}
}

func NewLoginReqTest() *LoginReq {
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

func (l *LoginReq) Len() int {
	return 0x22
}

func (l *LoginReq) MsgID() int {
	return l.id
}

func (l *LoginReq) Action() byte {
	return LoginReqType
}

func (l *LoginReq) IsRequest() bool {
	return true
}

func (l *LoginReq) Marshal() []byte {
	pkg := make([]byte, 0)

	pkg = append(pkg, 0x68,
		l.getLen())
	pkg = append(pkg, l.getID()...)
	pkg = append(pkg, 0x00, 0x01)
	pkg = append(pkg, l.getSN()...)
	pkg = append(pkg, l.getCSType(),
		l.getGunNumber(),
		l.getProtocolVersion())
	pkg = append(pkg, l.getProgrammingVersion()...)
	pkg = append(pkg, l.getNetType())
	pkg = append(pkg, l.getSim()...)
	pkg = append(pkg, l.getOperator())
	return pkg
}

func (l *LoginReq) UnMarshal(bytes []byte) error {
	return nil
}

func (l *LoginReq) getSN() []byte {
	return encodeSN(l.sn)
}

func (l *LoginReq) getCSType() byte {
	return getHex(int(l.csType))
}

func (l *LoginReq) getGunNumber() byte {
	return getHex(l.gunNumber)
}

func (l *LoginReq) getProtocolVersion() byte {
	str := strings.Split(l.protocolVersion, "V")[1]
	fl, _ := strconv.ParseFloat(str, 64)
	return getHex(int(fl * 10))
}

func (l *LoginReq) getProgrammingVersion() []byte {
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

func (l *LoginReq) getLen() byte {
	return byte(l.Len())
}

func (l *LoginReq) getID() []byte {
	strID := fmt.Sprintf("%04X", l.id)
	id, _ := convertStringToByte(strID)
	return id
}

// 登录认证应答
type loginResp struct {
	id int
	// 桩编号
	sn string
	// 登录结果
	success bool
}

func newLoginResp(id int, sn string, success bool) *loginResp {
	return &loginResp{id: id, sn: sn, success: success}
}

func (l *loginResp) Len() int {
	return 0x0C
}

func (l *loginResp) MsgID() int {
	return l.id
}

func (l *loginResp) Action() byte {
	return LoginRespType
}

func (l *loginResp) IsRequest() bool {
	return false
}

func (l *loginResp) Marshal() []byte {
	return nil
}

func (l *loginResp) UnMarshal(pkg []byte) error {
	start := pkg[0]
	if start != 0x68 {
		log.Println("数据格式错误")
		return errors.New("数据格式错误")
	}
	l.setID(pkg[2:4])
	postByte := make([]byte, len(pkg[6:13]))
	copy(postByte, pkg[6:13])
	removeZero(&postByte)
	postSn := fmt.Sprintf("%X", postByte)
	l.setSN(postSn)
	result := pkg[13]
	var success bool
	if result != 0x00 && result != 0x01 {
		log.Println("登录结果格式错误")
		return errors.New("登录结果格式错误")
	}
	if result == 0x00 {
		success = true
	} else {
		success = false
	}
	l.setSuccess(success)
	return nil
}

func (l *loginResp) setID(id []byte) {
	l.id = decodeSequence(id)
}

func (l *loginResp) setSN(sn string) {
	l.sn = sn
}

func (l *loginResp) setSuccess(su bool) {
	l.success = su
}
