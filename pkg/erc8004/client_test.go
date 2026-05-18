package erc8004

import (
	"context"
	"testing"
	"github.com/ronexlemon/agent-kit/pkg/chain"
)

func TestSuccesfulAgentRegistration(t *testing.T){
	client, err := chain.NewCreateClient("https://celo-sepolia.drpc.org")
	if err != nil {
		t.Fatalf("failed to connect to rpc: %v", err)
	}
	Identity:= "0x8004A818BFB912233c491871b3d84c89A494BD9e"
	Reputation:="0x8004B663056A597Dffe9eCcC1965A193B7388713"
	agent_instance,err := NewERC8004Client(client,Identity,Reputation)


	chainID, err := client.ChainID(context.Background())
	if err != nil {
		t.Fatalf("failed to get chain id: %v", err)
	}

	// Replace with a real valid test private key
	privateKey := "0xYOUR_PRIVATE_KEY"

	auth, err := chain.MakeAuth(privateKey, chainID.Int64())
if err != nil {
	t.Fatal(err)
}

	
	if err != nil {
		t.Fatalf("failed to create auth: %v", err)
	}

	tx,err := agent_instance.RegisterAgent(auth,"ipfs://bafkreiajlac5sd7dzo7li3c6jzy4dv4rqpc7d3cpxbiqn557addkrngley")

	if err !=nil{
		t.Fatal("Failed to Register agent",err)
	}
	t.Log("Agent registered Successful",tx)	
}