package types

import "testing"

func TestAgentRefString(t *testing.T) {
	ref := AgentRef{
		Namespace:       "eip155",
		ChainID:         44787,
		RegistryAddress: "0xRegistryAddress",
		AgentID:         "1",
	}

	got := ref.String()
	want := "eip155:44787:0xRegistryAddress:1"

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestAgentRefAgentRegistry(t *testing.T) {
	ref := AgentRef{
		Namespace:       "eip155",
		ChainID:         44787,
		RegistryAddress: "0xRegistryAddress",
		AgentID:         "1",
	}

	got := ref.AgentRegistry()
	want := "eip155:44787:0xRegistryAddress"

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestParseAgentRef(t *testing.T) {
	ref, err := ParseAgentRef("eip155:44787:0xRegistryAddress:1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if ref.Namespace != "eip155" {
		t.Fatalf("expected namespace eip155, got %q", ref.Namespace)
	}

	if ref.ChainID != 44787 {
		t.Fatalf("expected chain id 44787, got %d", ref.ChainID)
	}

	if ref.RegistryAddress != "0xRegistryAddress" {
		t.Fatalf("expected registry address 0xRegistryAddress, got %q", ref.RegistryAddress)
	}

	if ref.AgentID != "1" {
		t.Fatalf("expected agent id 1, got %q", ref.AgentID)
	}
}

func TestParseAgentRefRejectsInvalidFormat(t *testing.T) {
	_, err := ParseAgentRef("eip155:44787:0xRegistryAddress")
	if err == nil {
		t.Fatal("expected error for invalid agent ref")
	}
}

func TestParseAgentRefRejectsInvalidChainID(t *testing.T) {
	_, err := ParseAgentRef("eip155:not-a-chain:0xRegistryAddress:1")
	if err == nil {
		t.Fatal("expected error for invalid chain id")
	}
}

func TestParseAgentRefRejectsZeroChainID(t *testing.T) {
	_, err := ParseAgentRef("eip155:0:0xRegistryAddress:1")
	if err == nil {
		t.Fatal("expected error for zero chain id")
	}
}
