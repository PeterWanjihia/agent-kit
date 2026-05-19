package erc8004

import "testing"

func TestNormalizeNetworkDefaultsToTestnet(t *testing.T) {
	got := normalizeNetwork("")
	if got != NetworkTestnet {
		t.Fatalf("expected %q, got %q", NetworkTestnet, got)
	}
}

func TestNormalizeNetworkLowercasesInput(t *testing.T) {
	got := normalizeNetwork(" MainNet ")
	if got != NetworkMainnet {
		t.Fatalf("expected %q, got %q", NetworkMainnet, got)
	}
}

func TestNormalizeChainKeyCelo(t *testing.T) {
	got, err := normalizeChainKey(" celo ")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got != "Celo" {
		t.Fatalf("expected Celo, got %q", got)
	}
}

func TestNormalizeChainKeyRejectsUnsupportedChain(t *testing.T) {
	_, err := normalizeChainKey("solana")
	if err == nil {
		t.Fatal("expected error for unsupported chain")
	}
}

func TestResolveRPCURLCeloTestnet(t *testing.T) {
	rpcURL, err := resolveRPCURL("Celo", NetworkTestnet)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if rpcURL == "" {
		t.Fatal("expected rpc url")
	}
}

func TestResolveRPCURLCeloMainnet(t *testing.T) {
	rpcURL, err := resolveRPCURL("Celo", NetworkMainnet)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if rpcURL == "" {
		t.Fatal("expected rpc url")
	}
}

func TestResolveRPCURLRejectsUnsupportedNetwork(t *testing.T) {
	_, err := resolveRPCURL("Celo", "devnet")
	if err == nil {
		t.Fatal("expected error for unsupported network")
	}
}

func TestResolveContractAddressesCeloTestnet(t *testing.T) {
	identity, reputation, err := resolveContractAddresses("Celo", NetworkTestnet)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if identity == "" {
		t.Fatal("expected identity registry address")
	}

	if reputation == "" {
		t.Fatal("expected reputation registry address")
	}
}

func TestResolveContractAddressesCeloMainnet(t *testing.T) {
	identity, reputation, err := resolveContractAddresses("Celo", NetworkMainnet)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if identity == "" {
		t.Fatal("expected identity registry address")
	}

	if reputation == "" {
		t.Fatal("expected reputation registry address")
	}
}

func TestResolveContractAddressesRejectsUnsupportedNetwork(t *testing.T) {
	_, _, err := resolveContractAddresses("Celo", "devnet")
	if err == nil {
		t.Fatal("expected error for unsupported network")
	}
}
