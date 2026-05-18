package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewPinataClientRequiresJWT(t *testing.T) {
	client, err := NewPinataClient("")
	if err == nil {
		t.Fatal("expected error for missing jwt")
	}

	if client != nil {
		t.Fatal("expected client to be nil")
	}
}

func TestPublishJSONUploadsContentAndReturnsURI(t *testing.T) {
	var sawAuthHeader bool
	var sawContent bool

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}

		if r.Header.Get("Authorization") == "Bearer test-jwt" {
			sawAuthHeader = true
		}

		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}

		if _, ok := body["pinataContent"]; ok {
			sawContent = true
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"IpfsHash":"bafytest","PinSize":123,"Timestamp":"2026-01-01T00:00:00Z","isDuplicate":false}`)
	}))
	defer server.Close()

	client := &PinataClient{
		JWT:      "test-jwt",
		Endpoint: server.URL,
		Client:   server.Client(),
	}

	result, err := client.PublishJSON(context.Background(), PublishJSONInput{
		Name: "agent.json",
		Data: map[string]any{
			"name":        "Test Agent",
			"description": "Testing",
		},
	})
	if err != nil {
		t.Fatalf("expected upload to succeed: %v", err)
	}

	if !sawAuthHeader {
		t.Fatal("expected Authorization header")
	}

	if !sawContent {
		t.Fatal("expected pinataContent in request")
	}

	if result.CID != "bafytest" {
		t.Fatalf("expected CID bafytest, got %q", result.CID)
	}

	if result.URI != "ipfs://bafytest" {
		t.Fatalf("expected ipfs URI, got %q", result.URI)
	}

	if result.Size != 123 {
		t.Fatalf("expected size 123, got %d", result.Size)
	}
}

func TestPublishJSONHandlesPinataError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}))
	defer server.Close()

	client := &PinataClient{
		JWT:      "test-jwt",
		Endpoint: server.URL,
		Client:   server.Client(),
	}

	_, err := client.PublishJSON(context.Background(), PublishJSONInput{
		Name: "agent.json",
		Data: map[string]any{"name": "Test Agent"},
	})
	if err == nil {
		t.Fatal("expected error for non-2xx response")
	}
}

func TestPublishJSONRejectsMissingIpfsHash(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{}`)
	}))
	defer server.Close()

	client := &PinataClient{
		JWT:      "test-jwt",
		Endpoint: server.URL,
		Client:   server.Client(),
	}

	_, err := client.PublishJSON(context.Background(), PublishJSONInput{
		Name: "agent.json",
		Data: map[string]any{"name": "Test Agent"},
	})
	if err == nil {
		t.Fatal("expected error for missing IpfsHash")
	}
}
