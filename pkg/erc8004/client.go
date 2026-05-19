package erc8004

import (
	"fmt"
	"math/big"

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
type FeedbackInput struct{
	AgentId *big.Int
	Value *big.Int
	ValueDecimals uint8
    Tag1 string
    Tag2 string
    Endpoint string
    FeedbackURI string
    FeedbackHash [32]byte
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

func (erc *ERC8004Client) OwnerOfAgent(	opts *bind.CallOpts,agentID *big.Int) (common.Address, error) {

	if opts == nil {
		return common.Address{}, fmt.Errorf("call options are nil")
	}

	owner, err := erc.Registry.OwnerOf(opts, agentID)
	if err != nil {
		return common.Address{}, err
	}

	return owner, nil
}

func (erc *ERC8004Client) AgentWallet(opts *bind.CallOpts,agent_id *big.Int)(common.Address,error){
	address,err:=erc.Registry.GetAgentWallet(opts,agent_id)
	if err !=nil{
		return common.Address{},fmt.Errorf("Failed to get Agentic wallet")
	}
	return address,nil
}


func (erc *ERC8004Client) GetAgentMetadata(opts *bind.CallOpts,agent_id *big.Int,metadataKey string)([]byte,error){
	data,err:=erc.Registry.GetMetadata(opts,agent_id,metadataKey)
	if err !=nil{
		return nil,fmt.Errorf("Failed to get metadata, reconfirm metadata key")
	}
	return  data,nil
}


func (erc *ERC8004Client) GetAgentTokenURI(
	opts *bind.CallOpts,
	agentID *big.Int,
) (string, error) {

	if opts == nil {
		return "", fmt.Errorf("call opts is nil")
	}

	if agentID == nil {
		return "", fmt.Errorf("agentID is required")
	}

	return erc.Registry.TokenURI(opts, agentID)
}


func (erc *ERC8004Client) AgentGiveFeedBack(opts *bind.TransactOpts,input FeedbackInput)(*types.Transaction,error){
	return erc.Reputation.GiveFeedback(opts,input.AgentId,
		input.Value,
		input.ValueDecimals,
		input.Tag1,
		input.Tag2,
		input.Endpoint,
		input.FeedbackURI,
		input.FeedbackHash)
}
