package protocol

const (
	LoginReqType      = 0x01
	LoginRespType     = LoginReqType + 1
	HeartbeatReqType  = LoginRespType + 1
	HeartbeatRespType = HeartbeatReqType + 1
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

type ConnStatus byte

const (
	// Available 正常
	Available ConnStatus = 0x00
	// Faulted 故障
	Faulted ConnStatus = 0x01
)

type Header struct {
	Start   byte   // 起始标志 1个字节 固定0x68
	Length  int    // 长度 1个字节 不超过200
	MsgID   []byte // 序列号域 2个字节
	Encrypt byte   // 加密标志 1个字节
	Type    byte   // 帧类型 1个字节
	CRC     []byte // 帧校验域 2个字节
}

type Protocol interface {
	// Len 数据长度 一个字节 数据长度不超过200字节
	Len() int
	// MsgID 序列号域 两个字节 转化成int号使用些 数据包发送数据
	// 从0开始往上加
	MsgID() int
	// Action 帧类型标志 一个字节 到时候注册map直接使用byte
	Action() byte
	// IsRequest 是否是请求
	IsRequest() bool
	// Marshal 序列化
	Marshal() []byte
	// UnMarshal 反序列化
	UnMarshal([]byte) error
}
