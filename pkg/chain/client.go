package chain

import (
	
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewPublicClient(rpcUrl string)(*ethclient.Client,error){
	if rpcUrl == ""{
		return nil,fmt.Errorf("rpcul missing")
	}

	client,err :=  ethclient.Dial(rpcUrl)
	if err !=nil{
		return nil, fmt.Errorf("failed to connect to rpc: %w", err)
	}
	return client,nil
}


func  NewCreateClient(rpcUrl string)(*ethclient.Client,error){
	return  NewPublicClient(rpcUrl)
}
