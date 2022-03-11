package protocol

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLoginReq(t *testing.T) {
	//test := NewHeartbeatReqTest()
	//pkg := test.Marshal()
	//assert.NotNil(t, pkg)
	//fmt.Printf("%X", pkg)

	str := "680C000000025503141278230500DA4C"
	pkg, _ := convertStringToByte(str)
	test := loginResp{}
	err := test.UnMarshal(pkg)
	assert.Nil(t, err)
	assert.NotNil(t, test)
	fmt.Printf("%v", test)
}
