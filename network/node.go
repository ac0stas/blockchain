package network

import (
	"context"
	"fmt"
	"strings"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
)

func createSourceNode() host.Host {
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}
	return node
}

func createTargetNode() host.Host {
	node, err := libp2p.New(
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/8007",
		),
	)
	if err != nil {
		panic(err)
	}
	return node
}

func connectToTargetNode(sourceNode host.Host, targetNode host.Host) {
	targetNodeAddressInfo := host.InfoFromHost(targetNode)
	err := sourceNode.Connect(context.Background(), *targetNodeAddressInfo)
	if err != nil {
		panic(err)
	}
}

func countSourceNodePeers(sourceNode host.Host) int {
	return len(sourceNode.Network().Peers())
}

func printNodeID(node host.Host) {
	println(fmt.Sprintf("Node ID: %s", node.ID().String()))
}

func printNodeAddress(node host.Host) {
	addressesString := make([]string, 0)
	for _, address := range node.Addrs() {
		addressesString = append(addressesString, address.String())
	}
	println(fmt.Sprintf("Multiaddresses: %s", strings.Join(addressesString, ",")))

}
func connectNode() {
	sourceNode := createSourceNode()
	println("-- SOURCE NODE INFO --")
	printNodeID(sourceNode)
	printNodeAddress(sourceNode)

	targetNode := createTargetNode()
	println("-- TARGET NODE INFO --")
	printNodeID(targetNode)
	printNodeAddress(targetNode)

	connectToTargetNode(sourceNode, targetNode)

	println(fmt.Sprintf("Source node peers: %d", countSourceNodePeers(sourceNode)))
}
