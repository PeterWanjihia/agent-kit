package client

import (
	"testing"

	"github.com/ronexlemon/agent-kit/pkg/types"
)

func TestCreateAgentMetadataRequiresName(t *testing.T) {
	_, err := CreateAgentMetadata(types.CreateAgentInput{
		Description: "Explains Celo concepts to developers.",
	})

	if err == nil {
		t.Fatal("expected error for missing agent name")
	}
}

func TestCreateAgentMetadataRequiresDescription(t *testing.T) {
	_, err := CreateAgentMetadata(types.CreateAgentInput{
		Name: "CeloTutorAgent",
	})

	if err == nil {
		t.Fatal("expected error for missing agent description")
	}
}

func TestCreateAgentMetadataTrimsInput(t *testing.T) {
	metadata, err := CreateAgentMetadata(types.CreateAgentInput{
		Name:        "  CeloTutorAgent  ",
		Description: "  Explains Celo concepts to developers.  ",
	})

	if err != nil {
		t.Fatalf("expected metadata creation to succeed: %v", err)
	}

	if metadata.Name != "CeloTutorAgent" {
		t.Fatalf("expected trimmed name, got %q", metadata.Name)
	}

	if metadata.Description != "Explains Celo concepts to developers." {
		t.Fatalf("expected trimmed description, got %q", metadata.Description)
	}
}

func TestCreateAgentMetadataRejectsServiceNameWithoutEndpoint(t *testing.T) {
	_, err := CreateAgentMetadata(types.CreateAgentInput{
		Name:        "CeloTutorAgent",
		Description: "Explains Celo concepts to developers.",
		Services: []types.AgentService{
			{Name: "A2A",Version: "1.0.0"},
			{Name:"MCP",Version: "1.0.1"},
		},
	})

	if err == nil {
		t.Fatal("expected error for service name without endpoint")
	}
}

func TestCreateAgentMetadataRejectsServiceEndpointWithoutName(t *testing.T) {
	_, err := CreateAgentMetadata(types.CreateAgentInput{
		Name:            "CeloTutorAgent",
		Description:     "Explains Celo concepts to developers.",
		Services: []types.AgentService{
			{Endpoint: "https://agent.example/.well-known/agent-card.json",Version: "1.0.0"},
			{Endpoint: "https://agent.example2/.well-known/agent-card.json",Version: "1.0.0"},
		},
		
	})

	if err == nil {
		t.Fatal("expected error for service endpoint without name")
	}
}

func TestCreateAgentMetadataIncludesValidService(t *testing.T) {
	metadata, err := CreateAgentMetadata(types.CreateAgentInput{
		Name:            "CeloTutorAgent",
		Description:     "Explains Celo concepts to developers.",
		Services: []types.AgentService{
			{Name:"A2A",Endpoint: "https://agent.example/.well-known/agent-card.json",Version: "1.0.0"},
		},
	})

	if err != nil {
		t.Fatalf("expected metadata creation to succeed: %v", err)
	}

	if len(metadata.Services) != 1 {
		t.Fatalf("expected one service, got %d", len(metadata.Services))
	}

	service := metadata.Services[0]

	if service.Name != "A2A" {
		t.Fatalf("expected service name A2A, got %q", service.Name)
	}

	if service.Endpoint != "https://agent.example/.well-known/agent-card.json" {
		t.Fatalf("unexpected service endpoint: %q", service.Endpoint)
	}

	if service.Version != "1.0.0" {
		t.Fatalf("expected service version 1.0.0, got %q", service.Version)
	}
}
