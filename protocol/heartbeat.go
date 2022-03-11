package protocol

import "fmt"

type HeartbeatReq struct {
	id     int
	sn     string
	conn   int
	status ConnStatus
}

func NewHeartbeatReq(id int, sn string, conn int, status ConnStatus) *HeartbeatReq {
	return &HeartbeatReq{id: id, sn: sn, conn: conn, status: status}
}

func (h *HeartbeatReq) Len() int {
	return 0x0D
}

func (h *HeartbeatReq) MsgID() int {
	return h.id
}

func (h *HeartbeatReq) Action() byte {
	return HeartbeatReqType
}

func (h *HeartbeatReq) IsRequest() bool {
	return true
}

func (h *HeartbeatReq) Marshal() []byte {
	pkg := make([]byte, 0)
	pkg = append(pkg,
		0x68,
		h.getLen())
	pkg = append(pkg, h.getID()...)
	pkg = append(pkg, h.getType())
	pkg = append(pkg, h.getSN()...)
	pkg = append(pkg, h.getStatus(),
		0x68,
		0x90)
	return pkg
}

func (h *HeartbeatReq) UnMarshal(bytes []byte) error {
	return nil
}

func (h *HeartbeatReq) getLen() byte {
	return byte(h.Len())
}

func (h *HeartbeatReq) getID() []byte {
	step := fmt.Sprintf("%04X", h.MsgID())
	result, _ := convertStringToByte(step)
	return result
}

func (h *HeartbeatReq) getType() byte {
	return h.Action()
}

func (h *HeartbeatReq) getSN() []byte {
	return encodeSN(h.sn)
}

func (h *HeartbeatReq) getStatus() byte {
	return byte(h.status)
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
	return 0x0D
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
	return nil
}

func (h *HeartbeatResp) UnMarshal(pkg []byte) error {
	h.setID(pkg[2:4])
	mid := make([]byte, len(pkg[6:13]))
	copy(mid, pkg[6:13])
	removeZero(&mid)
	sn := fmt.Sprintf("%X", mid)

	h.setSN(sn)
	h.setGunNumber(pkg[13])
	h.setHeartbeatReply()
	return nil
}

func (h *HeartbeatResp) setID(pkg []byte) {
	h.id = decodeSequence(pkg)
}

func (h *HeartbeatResp) setSN(sn string) {
	h.sn = sn
}

func (h *HeartbeatResp) setGunNumber(nu byte) {
	h.gunNumber = int(nu)
}

func (h *HeartbeatResp) setHeartbeatReply() {
	h.heartbeatReply = COMMON
}
