package zinxlearn

import (
	"log"
	"net"
)

type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	HandleFunc HandleFunc

	ExitChan chan bool
}

func (c *Connection) Read() {
	log.Printf("connection [%d] is reading data", c.ConnID)
	defer log.Printf("connection [%d] stopt reading data", c.ConnID)
	defer c.Close()

	for {
		buf := make([]byte, 512)
		dataLen, err := c.Conn.Read(buf)
		if err != nil {
			log.Printf("connection [%d] read data failed, error: %s", c.ConnID, err.Error())
			continue
		}

		if err := c.HandleFunc(c.Conn, buf, dataLen); err != nil {
			log.Printf("connection [%d] handle data failed, error: %s", c.ConnID, err.Error())
		}
	}
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

func (c *Connection) Send(data []byte) error {
	return nil
}

func NewConnection(conn *net.TCPConn, connID uint32, HandleFunc HandleFunc) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		HandleFunc: HandleFunc,
		isClosed:   false,
		ExitChan:   make(chan bool),
	}

	return c
}
