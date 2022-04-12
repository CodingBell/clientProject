package protocol

type BootNotificationReq struct {
	id    int
	sn    string
	model string
	iccid string
}

type StatusType byte

const (
	Accepted StatusType = 0x01
	Rejected StatusType = 0x02
)

type BootNotificationResp struct {
	id            int
	status        StatusType
	keepalive     int
	orderInterval int
	key           string
	baseUrl       string
	//currentTime
}
