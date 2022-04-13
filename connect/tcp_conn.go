package main

import (
	"bufio"
	"fmt"
	"github.com/GeneralRelativityTheory/clientProject/util"
	"log"
	"net"
)

func main() {
	//util.String2Ascii("898601YY8SSXXXXXXXXP")

	//result := util.BootNotificationRequest()
	//fmt.Printf("%X", result)

	//Server2()

	util.GetStringLength("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
}

func Server2() {
	conn, err := net.Dial("tcp", "36.137.212.81:32887")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go func() {
		for {
			msg := make([]byte, 256)
			_, err := bufio.NewReader(conn).Read(msg)
			//time.Sleep(time.Second * 3)
			log.Println(msg[0:40])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}()

	b := util.BootNotificationRequest()
	_, err = conn.Write(b)
	if err != nil {
		return
	}
	select {}
}

//func Server() {
//	// create a tcp connection to localhost:8000 on the server side
//	conn, err := net.Dial("tcp", "36.137.212.81:32887")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	defer conn.Close()
//	go func() {
//		for {
//			// read the message from the client
//			message, err := bufio.NewReader(conn).ReadString('\n')
//			if err != nil {
//				log.Println(err)
//				return
//			}
//			log.Println("Message from client: " + message)
//		}
//	}()
//
//	go func() {
//		for {
//			// send a response
//			_, err = conn.Write(BootNotificationRequest())
//			if err != nil {
//				panic(err)
//			}
//		}
//	}()
//}

func hello() {
	fmt.Println("hello")
}
