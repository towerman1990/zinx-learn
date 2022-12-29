package main

import (
	"io"
	"log"
	"net"
	"time"

	"towerman1990.cn/zinx-learn/server"
)

func main() {
	log.Print("client start...")
	var msgID uint32 = 0
	conn, err := net.Dial("tcp", "127.0.0.1:2022")
	if err != nil {
		log.Fatalf("client net dial failed, error: %s", err.Error())
	}

	for {
		dataPack := server.NewDataPack()
		binaryMsg, err := dataPack.Pack(server.NewMessage(msgID, []byte("I'm from earth")))
		if err != nil {
			log.Fatalf("pack data failed, error: %s", err.Error())
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			log.Fatalf("client conn write data failed, error: %s", err.Error())
		}

		msgID++

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
