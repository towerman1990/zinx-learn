package server

import (
	"io"
	"log"
	"net"
	"testing"
	"time"

	"github.com/towerman1990/zinx-learn/utils"
)

func TestDataPack(t *testing.T) {
	addr := "127.0.0.1:2022"
	// server
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("server [%s] listen TCP failed, error: %s", utils.GlobalObject.Name, err.Error())
	}
	log.Printf("server listenning: %s", addr)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatalf("listener accept TCP failed, error: %s", err.Error())
				continue
			}

			go func(conn net.Conn) {
				dataPack := NewDataPack()
				for {
					headData := make([]byte, dataPack.GetHeadLength())
					if _, err := io.ReadFull(conn, headData); err != nil {
						log.Printf("read head data failed, error: %s", err.Error())
						continue
					}

					messageHead, err := dataPack.UnPack(headData)
					if err != nil {
						log.Printf("unpack data head failed, error: %s", err.Error())
						return
					}

					log.Printf("message ID: [%d], length: [%d]", messageHead.GetID(), messageHead.GetDataLength())
					if messageHead.GetDataLength() > 0 {
						message := messageHead.(*Message)
						message.Data = make([]byte, message.DataLength)
						if _, err := io.ReadFull(conn, message.Data); err != nil {
							log.Printf("read data failed, error: %s", err.Error())
							return
						}

						log.Printf("receive message [%d] successfully, content: [%s]", message.ID, string(message.Data))
					}

				}
			}(conn)
		}
	}()

	// client
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("client dial TCP failed, error: %s", err.Error())
		return
	}
	log.Printf("client dial: %s", addr)

	dataPack := NewDataPack()

	message1 := &Message{
		ID:         1,
		DataLength: 5,
		Data:       []byte{'h', 'e', 'l', 'l', 'o'},
	}

	sendData1, err := dataPack.Pack(message1)
	if err != nil {
		log.Printf("pack message failed, error: %s", err.Error())
		return
	}

	message2 := &Message{
		ID:         2,
		DataLength: 6,
		Data:       []byte{'w', 'o', 'r', 'l', 'd', '!'},
	}

	sendData2, err := dataPack.Pack(message2)
	if err != nil {
		log.Printf("pack message failed, error: %s", err.Error())
		return
	}

	message3 := &Message{
		ID:         3,
		DataLength: 3,
		Data:       []byte{'1', '2', '3'},
	}

	sendData3, err := dataPack.Pack(message3)
	if err != nil {
		log.Printf("pack message failed, error: %s", err.Error())
		return
	}

	sendData1 = append(sendData1, sendData2...)

	conn.Write(sendData1)
	time.Sleep(time.Second)
	conn.Write(sendData3)

	select {}
}
