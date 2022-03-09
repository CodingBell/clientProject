package protocol

type Protocol struct {
	FrameHeader   [4]byte
	MessageHeader [8]byte
	MessageBody   *[]byte
	FrameTail     [4]byte
}

func NewProtocol(frameHeader [4]byte,
	messageHeader [8]byte, messageBody *[]byte, frameTail [4]byte) *Protocol {
	return &Protocol{FrameHeader: frameHeader, MessageHeader: messageHeader,
		MessageBody: messageBody, FrameTail: frameTail}
}
