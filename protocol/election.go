package protocol

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
		protocolVersion:    "V1.4",
		programmingVersion: "v1.0.1",
		netType:            SIM,
		sim:                "01010101010101010101",
		operator:           TELECOM}
}

//func SelectRequestMethod(p *Protocol) {
//	messageHear := p.MessageHeader
//	if messageHear[0] != 0x68 {
//		log.Fatal("the message is not a request type")
//		return
//	}
//	var result string
//	switch messageHear[1] {
//	// login to sign in
//	case 0x10:
//		result = p.SignIn(p.MessageBody)
//	// authority
//	case 0x11:
//		result = p.Authority(p.MessageBody)
//	}
//
//	// return message to server
//	conn, err := createClient("tcp", "localhost", "9999")
//	defer conn.Close()
//	if err != nil {
//		log.Fatal("there is no server")
//		return
//	}
//	_, err = conn.Write([]byte(result))
//	if err != nil {
//		log.Fatal("the server is failed")
//		return
//	}
//	conn.Close()
//
//}
//
//func createClient(protoType, address, port string) (client net.Conn, err error) {
//	location := address + ":" + port
//	client, err = net.Dial(protoType, location)
//	return
//}
