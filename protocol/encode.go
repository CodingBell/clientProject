package protocol

func SignIn(req *LoginReq) []byte {
	pkg := make([]byte, 0)
	pkg = append(pkg, 0x68,
		0x22,
		0x00,
		0x00,
		0x00,
		0x01)
	pkg = append(pkg, req.getSN()...)
	pkg = append(pkg, req.getCSType(),
		req.getGunNumber(),
		req.getTenMultiVersion(),
	)
	pkg = append(pkg, req.getAsciiToByte()...)
	pkg = append(pkg, req.getNetType())
	pkg = append(pkg, req.getSim()...)
	pkg = append(pkg, req.getOperator())
	return pkg
}
