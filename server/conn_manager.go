package server

import (
	"fmt"
	"log"
	"sync"

	"github.com/towerman1990/zinx-learn/iface"
)

type ConnManager struct {
	connMap map[uint32]iface.IConnection

	connLock sync.RWMutex
}

func (cm *ConnManager) Add(conn iface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connMap[conn.GetConnID()] = conn
	log.Printf("connection add to connection manager successfully, connID = [%d]", conn.GetConnID())
}

func (cm *ConnManager) Remove(conn iface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	delete(cm.connMap, conn.GetConnID())
}

func (cm *ConnManager) Get(connID uint32) (conn iface.IConnection, err error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()
	if conn, ok := cm.connMap[conn.GetConnID()]; ok {
		return conn, err
	}

	return conn, fmt.Errorf("conn [%d] was not added", connID)
}

func (cm *ConnManager) Count() int {
	return len(cm.connMap)
}

func (cm *ConnManager) Clear() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for _, conn := range cm.connMap {
		conn.Close()
		cm.Remove(conn)
	}
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connMap:  make(map[uint32]iface.IConnection),
		connLock: sync.RWMutex{},
	}
}
