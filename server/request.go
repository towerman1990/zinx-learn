package server

import "towerman1990.cn/zinx-learn/iface"

type Request struct {
	conn iface.IConnecton

	data []byte
}

func (r *Request) GetConnection() iface.IConnecton {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
