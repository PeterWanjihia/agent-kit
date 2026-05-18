package types

import (
	"fmt"
	"strconv"
	"strings"
)

// AgentRef is the canonical pointer to a registered ERC-8004 agent.

type AgentRef struct {
	Namespace       string `json:"namespace"`
	ChainID         int64  `json:"chainId"`
	RegistryAddress string `json:"registryAddress"`
	AgentID         string `json:"agentId"`
}

// AgentRegistry returns the ERC-8004 registry identifier.
//
//	{namespace}:{chainId}:{identityRegistry}
func (r AgentRef) AgentRegistry() string {
	return fmt.Sprintf("%s:%d:%s", r.Namespace, r.ChainID, r.RegistryAddress)
}

// String returns the full canonical agent reference.
//
//	{namespace}:{chainId}:{identityRegistry}:{agentId}
func (r AgentRef) String() string {
	return fmt.Sprintf("%s:%d:%s:%s", r.Namespace, r.ChainID, r.RegistryAddress, r.AgentID)
}

// ParseAgentRef parses a canonical agent reference string.
//
//	eip155:44787:0xRegistryAddress:1
func ParseAgentRef(value string) (AgentRef, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return AgentRef{}, fmt.Errorf("agent reference is required")
	}

	parts := strings.Split(value, ":")
	if len(parts) != 4 {
		return AgentRef{}, fmt.Errorf(
			"invalid agent reference %q: expected namespace:chainId:registryAddress:agentId",
			value,
		)
	}

	namespace := strings.TrimSpace(parts[0])
	chainIDValue := strings.TrimSpace(parts[1])
	registryAddress := strings.TrimSpace(parts[2])
	agentID := strings.TrimSpace(parts[3])

	if namespace == "" {
		return AgentRef{}, fmt.Errorf("namespace is required")
	}

	if chainIDValue == "" {
		return AgentRef{}, fmt.Errorf("chain id is required")
	}

	chainID, err := strconv.ParseInt(chainIDValue, 10, 64)
	if err != nil {
		return AgentRef{}, fmt.Errorf("invalid chain id %q: %w", chainIDValue, err)
	}

	if chainID <= 0 {
		return AgentRef{}, fmt.Errorf("chain id must be greater than zero")
	}

	if registryAddress == "" {
		return AgentRef{}, fmt.Errorf("registry address is required")
	}

	if agentID == "" {
		return AgentRef{}, fmt.Errorf("agent id is required")
	}

	return AgentRef{
		Namespace:       namespace,
		ChainID:         chainID,
		RegistryAddress: registryAddress,
		AgentID:         agentID,
	}, nil
}
