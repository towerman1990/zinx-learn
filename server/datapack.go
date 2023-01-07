package server

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/towerman1990/zinx-learn/iface"
	"github.com/towerman1990/zinx-learn/utils"
)

type DataPack struct {
}

func (dp *DataPack) GetHeadLength() uint32 {
	return 8
}

func (dp *DataPack) Pack(message iface.IMessage) (packageData []byte, err error) {
	dataBuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuff, binary.LittleEndian, message.GetDataLength()); err != nil {
		return packageData, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, message.GetID()); err != nil {
		return packageData, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, message.GetData()); err != nil {
		return packageData, err
	}

	packageData = dataBuff.Bytes()

	return packageData, err
}

func (dp *DataPack) UnPack(binaryData []byte) (iface.IMessage, error) {
	message := &Message{}
	dataBuff := bytes.NewBuffer(binaryData)

	if err := binary.Read(dataBuff, binary.LittleEndian, &message.DataLength); err != nil {
		return message, err
	}

	if utils.GlobalObject.MaxPackageSize > 0 && message.DataLength > utils.GlobalObject.MaxPackageSize {
		return message, fmt.Errorf("data length [%d] beyond max package size", message.DataLength)
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &message.ID); err != nil {
		return message, err
	}

	return message, nil
}

func NewDataPack() *DataPack {
	return &DataPack{}
}
