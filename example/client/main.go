package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Print("client start...")
	time.Sleep(time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:2022")
	if err != nil {
		log.Fatalf("client net dial failed, error: %s", err.Error())
	}

	for {
		if _, err := conn.Write([]byte("Hello Zinx v0.1...")); err != nil {
			log.Fatalf("client conn write data failed, error: %s", err.Error())
		}

		buf := make([]byte, 512)
		if _, err := conn.Read(buf); err != nil {
			log.Fatalf("client conn read data failed, error: %s", err.Error())
		}

		log.Printf("client receive server response: %s", string(buf))

		time.Sleep(time.Second)
	}
}
