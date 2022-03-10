package tcp

import (
	"github.com/GeneralRelativityTheory/clientProject/protocol"
	"log"
	"net"
)

func SelectRequestMethod(req *protocol.LoginReq) {
	pkg := protocol.SignIn(req)
	if pkg[0] != 0x68 {
		log.Fatal("the message is not a request type")
		return
	}
	//var result string
	//switch messageHear[1] {
	//// login to sign in
	//case 0x10:
	//	result = p.SignIn(p.MessageBody)
	//// authority
	//case 0x11:
	//	result = p.Authority(p.MessageBody)
	//}

	// return message to server
	conn, err := createClient("tcp", "localhost", "9999")
	defer conn.Close()
	if err != nil {
		log.Fatal("there is no server")
		return
	}
	_, err = conn.Write(pkg)
	if err != nil {
		log.Fatal("the server is failed")
		return
	}
	conn.Close()

}

func createClient(protoType, address, port string) (client net.Conn, err error) {
	location := address + ":" + port
	client, err = net.Dial(protoType, location)
	return
}
