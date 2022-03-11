package tcp

import "testing"

func TestNewConn(t *testing.T) {
	conn, _ := NewConn("tcp", "localhost", "9999")

	conn.Send([]byte("is connected"))
	select {}
}
