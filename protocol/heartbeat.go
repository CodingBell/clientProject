package protocol

type HeartbeatReq struct {
	id     int
	sn     string
	conn   int
	status ConnStatus
}

type HeartbeatResp struct {
	id             int
	sn             string             // 桩编码
	gunNumber      int                // 枪号
	heartbeatReply HeartbeatReplyType // 心跳应答
}

func NewHeartbeatResp(id int, sn string,
	gunNumber int,
	heartbeatReply HeartbeatReplyType) *HeartbeatResp {
	return &HeartbeatResp{id: id,
		sn:             sn,
		gunNumber:      gunNumber,
		heartbeatReply: heartbeatReply}
}

func (h *HeartbeatResp) Len() int {
	panic("implement me")
}

func (h *HeartbeatResp) MsgID() int {
	return h.id
}

func (h *HeartbeatResp) Action() byte {
	return HeartbeatRespType
}

func (h *HeartbeatResp) IsRequest() bool {
	return false
}

func (h *HeartbeatResp) Marshal() []byte {
	panic("implement me")
}

func (h *HeartbeatResp) UnMarshal(pkg []byte) error {
	panic("implement me")
}
