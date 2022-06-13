package state

import (
	"context"
	
	"go-blockchain-api/config"
	
	"github.com/ethereum/go-ethereum/ethclient"
)

func connectRpc(){
	client, err := ethclient.DialContext(context.Background(), config.Global.Rpc)
	if err!= nil {

	}
	defer client.Close()
}