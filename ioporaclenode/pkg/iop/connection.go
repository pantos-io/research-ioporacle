package iop

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
)

type ConnectionManager struct {
	connections map[common.Address]*grpc.ClientConn
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[common.Address]*grpc.ClientConn),
	}
}

func (m *ConnectionManager) NewConnection(node RegistryContractOracleNode) (*grpc.ClientConn, error) {
	if conn, ok := m.connections[node.Addr]; ok {
		return conn, nil
	}
	conn, err := grpc.Dial(node.IpAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("dial %s: %v", node.IpAddr, err)
	}
	m.connections[node.Addr] = conn
	return conn, nil
}

func (m *ConnectionManager) FindByAddress(address common.Address) (*grpc.ClientConn, error) {
	if conn, ok := m.connections[address]; ok {
		return conn, nil
	}
	return nil, fmt.Errorf("connection to node with address %s not found", address)
}

func (m *ConnectionManager) Close() {
	for _, conn := range m.connections {
		_ = conn.Close()
	}
}
