package types

const AgentRegistrationType = "https://eips.ethereum.org/EIPS/eip-8004#registration-v1"

type AgentMetadata struct {
	Type        string         `json:"type"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       string         `json:"image,omitempty"`
	Services    []AgentService `json:"services,omitempty"`
}

type AgentService struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Version  string `json:"version,omitempty"`
}

type CreateAgentInput struct {
	Name            string
	Description     string
	Image           string
	Services    []AgentService
}
