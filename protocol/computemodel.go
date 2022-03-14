package protocol

type ComputeModelReq struct {
	id int
	// 桩编号
	sn string
	// 计费模型编号
	computeModelNo ComputeModelType
}

func NewComputeModelReq(id int,
	sn string,
	computeModelNo ComputeModelType) *ComputeModelReq {
	return &ComputeModelReq{id: id,
		sn:             sn,
		computeModelNo: computeModelNo}
}

func (c *ComputeModelReq) Len() int {
	return 0x0D
}

func (c *ComputeModelReq) MsgID() int {
	return c.id
}

func (c *ComputeModelReq) Action() byte {
	return ComputeModelReqType
}

func (c *ComputeModelReq) IsRequest() bool {
	return true
}

func (c *ComputeModelReq) Marshal() []byte {
	pkg := make([]byte, 0)
	pkg = append(pkg, 0x68,
		c.getLen())
	pkg =
}

func (c *ComputeModelReq) UnMarshal(bytes []byte) error {
	return nil
}

func (c *ComputeModelReq) getLen() byte {
	return byte(c.Len())
}

func (c *ComputeModelReq) getID() []byte {

}