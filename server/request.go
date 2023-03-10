package server

import "github.com/towerman1990/zinx-learn/iface"

type Request struct {
	conn iface.IConnection

	msg iface.IMessage
}

func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

func (r *Request) GetMessageData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMessageID() uint32 {
	return r.msg.GetID()
}
