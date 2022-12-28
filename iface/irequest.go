package iface

type IRequest interface {
	GetConnection() IConnecton

	GetData() []byte
}
