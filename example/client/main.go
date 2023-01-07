package main

import (
	"io"
	"log"
	"net"
	"time"

	"github.com/towerman1990/zinx-learn/server"
)

func main() {
	log.Print("client start...")
	conn, err := net.Dial("tcp", "127.0.0.1:2022")
	if err != nil {
		log.Fatalf("client net dial failed, error: %s", err.Error())
	}

	for {
		var msgID int64
		var data []byte
		if msgID = time.Now().Unix() % 2; msgID == 0 {
			data = []byte("ping message")
		} else {
			data = []byte("hello message")
		}
		message := server.NewMessage(uint32(msgID), data)
		dataPack := server.NewDataPack()
		binaryMsg, err := dataPack.Pack(message)
		if err != nil {
			log.Fatalf("pack data failed, error: %s", err.Error())
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			log.Fatalf("client conn write data failed, error: %s", err.Error())
		}

		binaryHead := make([]byte, dataPack.GetHeadLength())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			log.Printf("client read head data failed, error: %s", err.Error())
			break
		}

		msgHead, err := dataPack.UnPack(binaryHead)
		if err != nil {
			log.Printf("unpack server response message head failed, error: %s", err.Error())
			break
		}

		if msgHead.GetDataLength() > 0 {
			msg := msgHead.(*server.Message)
			msg.Data = make([]byte, msg.GetDataLength())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				log.Fatalf("read server response message data failed, error: %s", err.Error())
			}

			log.Printf("client receive server response:\n%s", string(msg.Data))
		}

		time.Sleep(time.Second)
	}
}
