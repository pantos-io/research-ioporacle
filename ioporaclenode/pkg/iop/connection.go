package iop

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ConnectionManager struct {
	sync.RWMutex
	registryContract *RegistryContractWrapper
	account          common.Address
	connections      map[common.Address]*grpc.ClientConn
}

func NewConnectionManager(registryContract *RegistryContractWrapper, account common.Address) *ConnectionManager {
	return &ConnectionManager{
		registryContract: registryContract,
		account:          account,
		connections:      make(map[common.Address]*grpc.ClientConn),
	}
}

func (m *ConnectionManager) WatchAndHandleRegisterOracleNodeLog(ctx context.Context) error {
	sink := make(chan *RegistryContractRegisterOracleNode)
	defer close(sink)

	sub, err := m.registryContract.WatchRegisterOracleNode(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			log.Infof("Received register oracle node event %s", event.Sender.String())
			if err = m.HandleRegisterOracleNodeLog(event); err != nil {
				log.Errorf("Handle register oracle node log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (m *ConnectionManager) HandleRegisterOracleNodeLog(event *RegistryContractRegisterOracleNode) error {
	node, err := m.registryContract.FindOracleNodeByAddress(nil, event.Sender)
	if err != nil {
		return fmt.Errorf("find oracle node by address %s: %w", event.Sender, err)
	}
	if _, err = m.NewConnection(node); err != nil {
		return fmt.Errorf("new connection: %w", err)
	}
	return nil
}

func (m *ConnectionManager) InitConnections() error {
	log.Info("Initialize connections to other nodes")
	nodes, err := m.registryContract.FindOracleNodes()
	if err != nil {
		return fmt.Errorf("find nodes: %w", err)
	}
	for _, node := range nodes {
		if node.Addr == m.account {
			continue
		}
		if _, err := m.NewConnection(node); err != nil {
			log.Errorf("New connection %s: %v", node.IpAddr, err)
			continue
		}
	}
	return nil
}

func (m *ConnectionManager) NewConnection(node RegistryContractOracleNode) (*grpc.ClientConn, error) {
	m.Lock()
	defer m.Unlock()
	if conn, ok := m.connections[node.Addr]; ok {
		return conn, nil
	}
	conn, err := grpc.Dial(node.IpAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("dial %s: %v", node.IpAddr, err)
	}
	m.connections[node.Addr] = conn

	log.Infof("New connection to %s with index %d and IP address %s", node.Addr, node.Index, node.IpAddr)

	return conn, nil
}

func (m *ConnectionManager) FindByAddress(address common.Address) (*grpc.ClientConn, error) {
	m.RLock()
	defer m.RUnlock()
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
