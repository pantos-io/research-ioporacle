package iop

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/iotaledger/hive.go/serializer"
	iota "github.com/iotaledger/iota.go/v2"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	dkg "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	"go.dedis.ch/kyber/v3/suites"
)

type DistKeyGenerator struct {
	sync.Mutex
	dkg               *dkg.DistKeyGenerator
	suite             suites.Suite
	connectionManager *ConnectionManager
	aggregator        *Aggregator
	mqttClient        mqtt.Client
	mqttTopic         []byte
	iotaClient        *iota.NodeHTTPAPIClient
	registryContract  *RegistryContractWrapper
	distKeyContract   *DistKeyContract
	ecdsaPrivateKey   *ecdsa.PrivateKey
	blsPrivateKey     kyber.Scalar
	account           common.Address
	deals             map[uint32]*dkg.Deal
	pendingResp       map[uint32][]*dkg.Response
	index             uint32
}

func NewDistKeyGenerator(
	suite suites.Suite,
	connectionManager *ConnectionManager,
	aggregator *Aggregator,
	mqttClient mqtt.Client,
	mqttTopic []byte,
	iotaClient *iota.NodeHTTPAPIClient,
	registryContract *RegistryContractWrapper,
	distKeyContract *DistKeyContract,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	blsPrivateKey kyber.Scalar,
	account common.Address,
) *DistKeyGenerator {
	return &DistKeyGenerator{
		suite:             suite,
		connectionManager: connectionManager,
		aggregator:        aggregator,
		mqttClient:        mqttClient,
		mqttTopic:         mqttTopic,
		iotaClient:        iotaClient,
		registryContract:  registryContract,
		distKeyContract:   distKeyContract,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		blsPrivateKey:     blsPrivateKey,
		account:           account,
		deals:             make(map[uint32]*dkg.Deal),
		pendingResp:       make(map[uint32][]*dkg.Response),
	}
}

func (g *DistKeyGenerator) ListenAndProcess(ctx context.Context) error {
	go func() {
		if err := g.WatchAndHandleDistKeyGenerationLog(ctx); err != nil {
			log.Errorf("watch and handle dkg log: %v", err)
		}
	}()
	go func() {
		if err := g.ListenAndProcessResponse(); err != nil {
			log.Errorf("listen and process response: %v", err)
		}
	}()
	return nil
}

func (g *DistKeyGenerator) WatchAndHandleDistKeyGenerationLog(ctx context.Context) error {
	sink := make(chan *DistKeyContractDistKeyGenerationLog)
	defer close(sink)

	sub, err := g.distKeyContract.WatchDistKeyGenerationLog(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			log.Infof("Received distributed key generation event with t %s", event.Threshold)
			if err := g.HandleDistributedKeyGenerationLog(event); err != nil {
				log.Errorf("handle dkg log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (g *DistKeyGenerator) HandleDistributedKeyGenerationLog(event *DistKeyContractDistKeyGenerationLog) error {

	g.deals = make(map[uint32]*dkg.Deal)
	g.pendingResp = make(map[uint32][]*dkg.Response)

	nodes, err := g.registryContract.FindOracleNodes()
	if err != nil {
		return fmt.Errorf("find nodes: %w", err)
	}

	pubKeys := make([]kyber.Point, len(nodes))
	for i, node := range nodes {
		if node.Addr == g.account {
			g.index = uint32(node.Index.Uint64())
		}
		pubKey := g.suite.Point()
		if err := pubKey.UnmarshalBinary(node.PubKey); err != nil {
			return fmt.Errorf("unmarshal public key: %w", err)
		}
		pubKeys[i] = pubKey
	}

	threshold := int(event.Threshold.Int64())
	g.aggregator.SetThreshold(threshold)
	g.dkg, err = dkg.NewDistKeyGenerator(
		g.suite,
		g.blsPrivateKey,
		pubKeys,
		threshold,
	)
	if err != nil {
		return fmt.Errorf("new dkg: %w", err)
	}

	//Wait until every participant is prepared. TODO: Wait for head n
	time.Sleep(5 * time.Second)

	deals, err := g.dkg.Deals()
	if err != nil {
		return fmt.Errorf("deals: %w", err)
	}

	g.SendDeals(nodes, deals)

	timeout := time.After(DkgTimeout)
loop:
	for {
		select {
		case <-timeout:
			g.dkg.SetTimeout()
			break loop
		default:
			if g.dkg.Certified() {
				break loop
			}
			time.Sleep(500 * time.Millisecond)
		}
	}

	log.Infof("Qualified shares: %v", g.dkg.QualifiedShares())
	log.Infof("QUAL: %v", g.dkg.QUAL())

	distKey, err := g.dkg.DistKeyShare()
	if err != nil {
		return fmt.Errorf("distributed key share: %w", err)
	}

	log.Infof("Public Key: %v", distKey.Public())
	pubKey, err := PubKeyToBig(distKey.Public())
	if err != nil {
		return fmt.Errorf("public key to big int: %w", err)
	}
	numberOfValidators := len(g.dkg.QUAL())
	auth := bind.NewKeyedTransactor(g.ecdsaPrivateKey)
	if _, err = g.distKeyContract.SetPublicKey(auth, pubKey, big.NewInt(int64(numberOfValidators))); err != nil {
		return fmt.Errorf("set public key: %w", err)
	}

	return nil
}

func (g *DistKeyGenerator) SendDeals(nodes []RegistryContractOracleNode, deals map[int]*dkg.Deal) {
	for i, deal := range deals {
		conn, err := g.connectionManager.FindByAddress(nodes[i].Addr)
		if err != nil {
			log.Errorf("find connection by address: %v", err)
			continue
		}
		client := NewOracleNodeClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		request := &SendDealRequest{
			Deal: DealToPb(deal),
		}
		log.Infof("Sending deal to node %d", i)
		if _, err := client.SendDeal(ctx, request); err != nil {
			log.Errorf("send deal: %v", err)
		}
		cancel()
	}
}

func (g *DistKeyGenerator) ListenAndProcessResponse() error {
	if token := g.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("connect to broker: %w", token.Error())
	}

	topic := fmt.Sprintf("messages/indexation/%s", hex.EncodeToString(g.mqttTopic))
	if token := g.mqttClient.Subscribe(topic, 1, func(c mqtt.Client, m mqtt.Message) {
		g.publishHandler(c, m)
	}); token.Wait() && token.Error() != nil {
		return fmt.Errorf("subscribe to topic: %w", token.Error())
	}

	return nil
}

func (g *DistKeyGenerator) publishHandler(c mqtt.Client, msg mqtt.Message) {
	iotaMsg := &iota.Message{}
	if _, err := iotaMsg.Deserialize(msg.Payload(), serializer.DeSeriModeNoValidation); err != nil {
		log.Errorf("malformed mqtt message: %w", err)
		return
	}

	var response dkg.Response
	if err := json.Unmarshal(iotaMsg.Payload.(*iota.Indexation).Data, &response); err != nil {
		log.Errorf("unmarshal response: %w", err)
	}

	go func() {
		if err := g.ProcessResponse(&response); err != nil {
			log.Errorf("handle response: %w", err)
		}
	}()
}

func (g *DistKeyGenerator) ProcessDeal(deal *dkg.Deal) (*dkg.Response, error) {
	g.Lock()
	defer g.Unlock()

	response, err := g.dkg.ProcessDeal(deal)
	if err != nil {
		return nil, fmt.Errorf("process deal: %w", err)
	}
	g.deals[deal.Index] = deal

	if _, ok := g.pendingResp[deal.Index]; ok {
		for _, val := range g.pendingResp[deal.Index] {
			_, err = g.dkg.ProcessResponse(val)
			if err != nil {
				return nil, fmt.Errorf("process deal: %w", err)
			}
		}
	}

	if err = g.BroadcastResponse(response); err != nil {
		return nil, fmt.Errorf("broadcast response: %w", err)
	}
	return response, nil
}

func (g *DistKeyGenerator) BroadcastResponse(response *dkg.Response) error {
	log.Infof("Broadcasting response for deal from %d", response.Index)
	b, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("marshal response: %w", err)
	}

	payload := &iota.Indexation{
		Index: 	g.mqttTopic,
		Data:	b,
	}

	msg, err := iota.NewMessageBuilder().
		Payload(payload).
		Build()
	if err != nil {
		return fmt.Errorf("build iota message: %w", err)
	}

	if _, err := g.iotaClient.SubmitMessage(context.Background(), msg); err != nil {
		return fmt.Errorf("submit message: %w", err)
	}
	return nil
}

func (g *DistKeyGenerator) ProcessResponse(response *dkg.Response) error {
	log.Infof("Handle response for dealer %d and verifier %d", response.Index, response.Response.Index)
	g.Lock()
	defer g.Unlock()

	if response.Response.Index == g.index {
		return nil
	}

	if _, ok := g.deals[response.Index]; !ok && response.Index != g.index {
		g.pendingResp[response.Index] = append(g.pendingResp[response.Index], response)
		return nil
	}

	if _, err := g.dkg.ProcessResponse(response); err != nil {
		return fmt.Errorf("process response: %w", err)
	}
	return nil
}

func (g *DistKeyGenerator) DistKeyShare() (*dkg.DistKeyShare, error) {
	return g.dkg.DistKeyShare()
}
