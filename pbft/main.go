package main

import (
	"os"
	"github.com/Snooowgh/consensusPBFT/pbft/network"
)

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID)

	server.Start()
}
