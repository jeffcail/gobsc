package gobsc

import (
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

type rpcClient struct {
	c *rpc.Client
}

var Cli *rpcClient

func NewClient(netWork string) {
	client, err := rpc.Dial(netWork)
	if err != nil {
		log.Fatal(err)
	}
	Cli = &rpcClient{c: client}
}
