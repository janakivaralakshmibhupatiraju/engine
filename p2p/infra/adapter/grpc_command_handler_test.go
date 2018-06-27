package adapter_test

import (
	"encoding/json"
	"testing"

	"github.com/it-chain/it-chain-Engine/p2p"
	"github.com/it-chain/it-chain-Engine/p2p/infra/adapter"
	"github.com/it-chain/midgard"
	"github.com/magiconair/properties/assert"
)

type MockLeaderApi struct{}

func (mla MockLeaderApi) UpdateLeader(leader p2p.Leader) error { return nil }
func (mla MockLeaderApi) DeliverLeaderInfo(nodeId p2p.PeerId)  {}

type MockPeerApi struct{}

func (mna MockPeerApi) UpdatePeerList(nodeList []p2p.Peer) error { return nil }
func (mna MockPeerApi) DeliverPeerList(nodeId p2p.PeerId) error  { return nil }
func (mna MockPeerApi) AddPeer(node p2p.Peer)                    {}

func TestGrpcMessageHandler_HandleMessageReceive(t *testing.T) {

	leader := p2p.Leader{}
	leaderByte, _ := json.Marshal(leader)

	nodeList := make([]p2p.Peer, 0)
	newPeerList := append(nodeList, p2p.Peer{PeerId: p2p.PeerId{Id: "123"}})
	nodeListByte, _ := json.Marshal(newPeerList)

	tests := map[string]struct {
		input struct {
			command p2p.GrpcRequestCommand
		}
		err error
	}{
		"leader info deliver test success": {
			input: struct {
				command p2p.GrpcRequestCommand
			}{
				command: p2p.GrpcRequestCommand{
					CommandModel: midgard.CommandModel{ID: "123"},
					Data:         leaderByte,
					Protocol:     "LeaderInfoDeliverProtocol",
				},
			},
			err: nil,
		},
		"node list deliver test success": {
			input: struct {
				command p2p.GrpcRequestCommand
			}{
				command: p2p.GrpcRequestCommand{
					CommandModel: midgard.CommandModel{ID: "123"},
					Data:         nodeListByte,
					Protocol:     "PeerListDeliverProtocol",
				},
			},
			err: nil,
		},
	}
	leaderApi := MockLeaderApi{}
	nodeApi := MockPeerApi{}
	messageHandler := adapter.NewGrpcCommandHandler(leaderApi, nodeApi)

	for testName, test := range tests {
		t.Logf("running test case %s", testName)
		err := messageHandler.HandleMessageReceive(test.input.command)
		assert.Equal(t, err, test.err)
	}

}
