package util

type Protocol struct {
	FrameHeader   *[]byte
	MessageHeader *[]byte
	MessageBody   *[]byte
	FrameTail     *[]byte
}

func NewProtocol(frameHeader *[]byte,
	messageHeader *[]byte,
	messageBody *[]byte,
	frameTail *[]byte) *Protocol {
	return &Protocol{FrameHeader: frameHeader,
		MessageHeader: messageHeader,
		MessageBody:   messageBody,
		FrameTail:     frameTail}
}

func (p *Protocol) getFrame() []byte {
	result := make([]byte, 0)
	result = append(result, *p.FrameHeader...)
	result = append(result, *p.MessageHeader...)
	result = append(result, *p.MessageBody...)
	result = append(result, *p.FrameTail...)
	return result
}
