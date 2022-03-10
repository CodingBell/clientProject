package tcp

import (
	"log"
	"net"
	"sync"
	"time"
)

var AUTOINCREMENTID int = 0

type MyConn struct {
	conn   net.Conn
	myChan chan []byte
	mu     sync.Mutex
	timer  *time.Ticker
}

func NewConn(protoType, address, port string) (*MyConn, error) {
	conn, err := createClient(protoType, address, port)
	if err != nil {
		log.Println("连接失败")
		return nil, err
	}
	m := &MyConn{
		conn:   conn,
		myChan: make(chan []byte, 1),
		timer:  time.NewTicker(5 * time.Second),
	}
	return m, nil
}

func createClient(protoType, address, port string) (client net.Conn, err error) {
	location := address + ":" + port
	client, err = net.Dial(protoType, location)
	return
}

//func (m *MyConn) ProcessMessage(protoType, address, port string, proto protocol.Proto) error {
//	err := m.SetConn(protoType, address, port)
//	if err != nil {
//		log.Println("无法取得连接")
//		return err
//	}
//	defer m.conn.Close()
//	if proto.IsRequest() {
//		go m.Write()
//	} else {
//		go m.Read()
//	}
//}

func (m *MyConn) Read() {
	pkg := make([]byte, 256)
	for {
		_, err := m.conn.Read(pkg)
		if err != nil {
			log.Println(err)
			break
		}
		// todo 处理协议数据
	}
}

func (m *MyConn) Write() {
	for {
		select {
		case pkg := <-m.myChan:
			_, err := m.conn.Write(pkg)
			if err != nil {
				log.Println(err)
				break
			}
		case <-m.timer.C:
			// todo 心跳
		}
	}
}

//func SelectRequestMethod(req *protocol.LoginReq) {
//	pkg := protocol.SignIn(req)
//	if pkg[0] != 0x68 {
//		log.Fatal("the message is not a request type")
//		return
//	}
//	conn, err := createClient("tcp", "localhost", "9999")
//	defer conn.Close()
//	if err != nil {
//		log.Fatal("there is no server")
//		return
//	}
//	_, err = conn.Write(pkg)
//	if err != nil {
//		log.Fatal("the server is failed")
//		return
//	}
//	conn.Close()
//}
