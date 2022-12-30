package server

import (
	"fmt"
	"io"
	"log"
	"net"

	"towerman1990.cn/zinx-learn/iface"
)

type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	MessageHandler iface.IMessageHandler

	isClosed bool

	ExitChan chan bool
}

func (c *Connection) Read() {
	log.Printf("reading data from connection [%d] ", c.ConnID)
	defer log.Printf("connection [%d] stopt reading data", c.ConnID)
	defer c.Close()

	for {
		dataPack := NewDataPack()
		headData := make([]byte, dataPack.GetHeadLength())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			log.Printf("connection [%d] read data failed, error: %s", c.ConnID, err.Error())
			break
		}

		msg, err := dataPack.UnPack(headData)
		if err != nil {
			log.Printf("conn [%d] unpack data failed, error: %s", c.ConnID, err.Error())
			break
		}

		var data []byte
		if msg.GetDataLength() > 0 {
			data = make([]byte, msg.GetDataLength())
			if _, err = io.ReadFull(c.GetTCPConnection(), data); err != nil {
				log.Printf("conn [%d] read data failed, error: %s", c.ConnID, err.Error())
				break
			}
		}
		msg.SetData(data)

		req := &Request{
			conn: c,
			msg:  msg,
		}

		go c.MessageHandler.ExecHandler(req)
	}
}

func (c *Connection) SendMsg(msgID uint32, data []byte) (err error) {
	if c.isClosed {
		return fmt.Errorf("connection [%d] has closed", c.ConnID)
	}

	dataPack := NewDataPack()
	binaryMsg, err := dataPack.Pack(NewMessage(msgID, data))
	if err != nil {
		return fmt.Errorf("conn [%d] packed message [%d] failed, error: %s", c.ConnID, msgID, err.Error())
	}

	if _, err := c.GetTCPConnection().Write(binaryMsg); err != nil {
		return fmt.Errorf("conn [%d] wrote message [%d] failed, error: %s", c.ConnID, msgID, err.Error())
	}

	return
}

func (c *Connection) Open() {
	log.Printf("open connection [%d]", c.ConnID)
	c.Read()
}

func (c *Connection) Close() (err error) {
	log.Printf("close connection [%d]", c.ConnID)

	if c.isClosed {
		return
	}

	c.isClosed = true

	return c.Conn.Close()
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func NewConnection(conn *net.TCPConn, connID uint32, messageHandler iface.IMessageHandler) *Connection {
	c := &Connection{
		Conn:           conn,
		ConnID:         connID,
		MessageHandler: messageHandler,
		isClosed:       false,
		ExitChan:       make(chan bool),
	}

	return c
}
