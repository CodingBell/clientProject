package main

import (
	"log"
	"net"
)

func recvMessage(client net.Conn) error {
	var message []byte
	message = make([]byte, 1025)
	for {
		read, _ := client.Read(message)
		if read > 0 {
			log.Printf("%s", message[0:read])
		}
	}
	return nil
}
func main() {
	server, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Fatal("start server failed!\n")
	}
	defer server.Close()

	log.Println("server is running...")
	for {
		client, err := server.Accept()
		if err != nil {
			log.Fatal("Accept error\n")
		}
		log.Println("the client is connected...")
		go recvMessage(client)
	}
}
