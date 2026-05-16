package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ronexlemon/agent-kit/pkg/types"
)

const (
	maxAgentNameLength        = 100
	maxAgentDescriptionLength = 2000
)

func CreateAgentMetadata(input types.CreateAgentInput) (types.AgentMetadata, error) {
	name := strings.TrimSpace(input.Name)
	description := strings.TrimSpace(input.Description)
	image := strings.TrimSpace(input.Image)
	serviceName := strings.TrimSpace(input.ServiceName)
	serviceEndpoint := strings.TrimSpace(input.ServiceEndpoint)
	serviceVersion := strings.TrimSpace(input.ServiceVersion)

	if name == "" {
		return types.AgentMetadata{}, fmt.Errorf("agent name is required")
	}

	if len(name) > maxAgentNameLength {
		return types.AgentMetadata{}, fmt.Errorf("agent name must be at most %d characters", maxAgentNameLength)
	}

	if description == "" {
		return types.AgentMetadata{}, fmt.Errorf("agent description is required")
	}

	if len(description) > maxAgentDescriptionLength {
		return types.AgentMetadata{}, fmt.Errorf("agent description must be at most %d characters", maxAgentDescriptionLength)
	}

	if image != "" && !isValidURI(image) {
		return types.AgentMetadata{}, fmt.Errorf("agent image must be a valid URI")
	}

	if serviceName == "" && serviceEndpoint != "" {
		return types.AgentMetadata{}, fmt.Errorf("service name is required when service endpoint is provided")
	}

	if serviceName != "" && serviceEndpoint == "" {
		return types.AgentMetadata{}, fmt.Errorf("service endpoint is required when service name is provided")
	}

	metadata := types.AgentMetadata{
		Type:        types.AgentRegistrationType,
		Name:        name,
		Description: description,
		Image:       image,
	}

	if serviceName != "" && serviceEndpoint != "" {
		if !isValidURI(serviceEndpoint) {
			return types.AgentMetadata{}, fmt.Errorf("service endpoint must be a valid URI")
		}

		metadata.Services = append(metadata.Services, types.AgentService{
			Name:     serviceName,
			Endpoint: serviceEndpoint,
			Version:  serviceVersion,
		})
	}

	return metadata, nil
}

func isValidURI(raw string) bool {
	parsed, err := url.ParseRequestURI(raw)
	if err != nil {
		return false
	}

	return parsed.Scheme != "" && parsed.Host != ""
}
