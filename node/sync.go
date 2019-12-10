package node

import (
	"time"
	"context"
	"fmt"
	"net/http"
)

func (n *Node) sync(ctx context.Context) error {
	ticker := time.NewTicker(45 * time.Second)

	for {
		select {
		case <- ticker.C:
			fmt.Println("Searching for new Peers and Blocks...")

			n.fetchNewBlocksAndPeers()

		case <- ctx.Done():
			ticker.Stop()
		}
	}
}

func (n *Node) fetchNewBlocksAndPeers() {
	for _, peer := range n.knownPeers {
		status, err := queryPeerStatus(peer)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			continue
		}

		localBlockNumber := n.state.LatestBlock().Header.Number
		if localBlockNumber < status.Number {
			newBlocksCount := status.Number - localBlockNumber

			fmt.Printf("Found %d new blocks from Peer %s\n", newBlocksCount, peer.IP)
		}

		for _, statusPeer := range status.KnownPeers {
			newPeer, isKnownPeer := n.knownPeers[statusPeer.TcpAddress()]
			if !isKnownPeer {
				fmt.Printf("Found new Peer %s\n", peer.TcpAddress())

				n.knownPeers[statusPeer.TcpAddress()] = newPeer
			}
		}
	}
}

func queryPeerStatus(peer PeerNode) (StatusRes, error) {
	url := fmt.Sprintf("http://%s/%s", peer.TcpAddress(), endpointStatus)
	res, err := http.Get(url)
	if err != nil {
		return StatusRes{}, err
	}

	statusRes := StatusRes{}
	err = readRes(res, &statusRes)
	if err != nil {
		return StatusRes{}, err
	}

	return statusRes, nil
}
