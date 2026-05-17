package erc8004

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ronexlemon/agent-kit/pkg/contract/registry"
	"github.com/ronexlemon/agent-kit/pkg/contract/reputation"
)


type ERC8004Client struct{
	Registry *registry.Registry
	Reputation *reputation.Reputation
}

func NewERC8004Client(eth *ethclient.Client,registryAddr string,reputationAddr string)(*ERC8004Client,error){
	reg, err := registry.NewRegistry(
		common.HexToAddress(registryAddr),
		eth,
	)
	if err != nil {
		return nil, err
	}

	rep, err := reputation.NewReputation(
		common.HexToAddress(reputationAddr),
		eth,
	)
	if err != nil {
		return nil, err
	}

	return &ERC8004Client{
		Registry:   reg,
		Reputation: rep,
	}, nil

}

func (erc *ERC8004Client) RegisterAgent(auth *bind.TransactOpts,ipfsMetadata string)(*types.Transaction,error){

	tx, err := erc.Registry.Register1(
		auth,
		ipfsMetadata,
	)
	if err !=nil{
		return nil,fmt.Errorf("Failed to register Agent")
	}

	return tx,nil

}
