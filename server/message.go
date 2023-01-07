package server

import "github.com/towerman1990/zinx-learn/iface"

type Message struct {
	ID         uint32
	DataLength uint32
	Data       []byte
}

func (m *Message) GetID() uint32 {
	return m.ID
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) GetDataLength() uint32 {
	return m.DataLength
}

func (m *Message) SetID(messageID uint32) {
	m.ID = messageID
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) SetDataLength(dataLength uint32) {
	m.DataLength = dataLength
}

func NewMessage(ID uint32, data []byte) iface.IMessage {
	return &Message{
		ID:         ID,
		DataLength: uint32(len(data)),
		Data:       data,
	}
}
