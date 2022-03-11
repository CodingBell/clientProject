package protocol

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHeartbeatReq(t *testing.T) {
	//test := NewHeartbeatReqTest()
	//marshal := test.Marshal()
	//assert.NotNil(t, marshal)
	//fmt.Printf("%X", marshal)

	str := "680D3600000455031412782305010065B2"
	pkg, _ := convertStringToByte(str)
	test := HeartbeatResp{}
	err := test.UnMarshal(pkg)
	assert.Nil(t, err)
	assert.NotNil(t, test)
	fmt.Printf("%v", test)
}
