package iface

type IMessage interface {
	GetID() uint32

	GetData() []byte

	GetDataLength() uint32

	SetID(uint32)

	SetData([]byte)

	SetDataLength(uint32)
}
