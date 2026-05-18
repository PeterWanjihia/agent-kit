package chain

import (
	"context"
	"testing"
)

func TestIncorrectPrivateKey(t *testing.T) {
	client, err := NewCreateClient("https://celo-sepolia.drpc.org")
	if err != nil {
		t.Fatalf("failed to connect to rpc: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		t.Fatalf("failed to get chain id: %v", err)
	}

	auth, err := MakeAuth("0x6425782757258275872", chainID.Int64())

	if err == nil {
		t.Fatal("expected invalid private key error but got nil")
	}

	if auth != nil {
		t.Fatal("expected auth to be nil")
	}
}

func TestValidPrivateKey(t *testing.T) {
	client, err := NewCreateClient("https://celo-sepolia.drpc.org")
	if err != nil {
		t.Fatalf("failed to connect to rpc: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		t.Fatalf("failed to get chain id: %v", err)
	}

	// Replace with a real valid test private key
	privateKey := "0xYOUR_PRIVATE_KEY"

	auth, err := MakeAuth(privateKey, chainID.Int64())

	
	if err != nil {
		t.Fatalf("failed to create auth: %v", err)
	}

	
	if auth == nil {
		t.Fatal("expected auth but got nil")
	}


	t.Log("valid private key auth created successfully")
}