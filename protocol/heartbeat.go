package protocol

type HeartbeatReq struct {
	id     int
	sn     string
	conn   int
	status ConnStatus
}
