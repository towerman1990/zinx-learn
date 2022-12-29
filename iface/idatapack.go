package iface

type IDataPack interface {
	GetHeadLength() uint32

	Pack(message IMessage) []byte

	UnPack([]byte) (IMessage, error)
}
