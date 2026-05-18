package chain

import (
	"context"
	"testing"
)


func TestInvalidRpcURL(t *testing.T){

	client,err:= NewCreateClient("https://test.ethereum.com")
if err == nil {
		t.Fatal("expected error but got nil")
	}

	if client != nil {
		t.Fatal("expected client to be nil")
	}
}

func TestValidRpcURL(t *testing.T) {
	client, err := NewCreateClient("https://celo-sepolia.drpc.org")
	if err != nil {
		t.Fatalf("failed to connect to rpc: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		t.Fatalf("failed to get chain id: %v", err)
	}

	if chainID == nil {
		t.Fatal("chain id is nil")
	}

	t.Logf("connected successfully to chain id %s", chainID.String())
}