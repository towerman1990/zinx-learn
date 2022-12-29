package iface

type IRequest interface {
	GetConnection() IConnecton

	GetMessageData() []byte

	GetMessageID() uint32
}
