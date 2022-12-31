package iface

type IRequest interface {
	GetConnection() IConnection

	GetMessageData() []byte

	GetMessageID() uint32
}
