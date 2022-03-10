package protocol

import (
	"fmt"
	"log"
)

func SignResp(pkg []byte) *loginResp {
	start := pkg[0]
	if start != 0x68 {
		log.Println("数据格式错误")
		return nil
	}

	postByte := make([]byte, len(pkg[6:13]))
	copy(postByte, pkg[6:13])
	removeZero(&postByte)
	postSn := fmt.Sprintf("%X", postByte)

	result := pkg[13]
	var success bool
	if result != 0x00 && result != 0x01 {
		log.Println("登录结果格式错误")
		return nil
	}
	if result == 0x00 {
		success = true
	} else {
		success = false
	}

	resp := newLoginResp(3, postSn, success)
	return resp
}
