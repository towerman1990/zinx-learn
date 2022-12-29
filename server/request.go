package server

import "towerman1990.cn/zinx-learn/iface"

type Request struct {
	conn iface.IConnecton

	msg iface.IMessage
}

func (r *Request) GetConnection() iface.IConnecton {
	return r.conn
}

func (r *Request) GetMessageData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMessageID() uint32 {
	return r.msg.GetID()
}
