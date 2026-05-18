package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const defaultPinataJSONEndpoint = "https://api.pinata.cloud/pinning/pinJSONToIPFS"

type PinataClient struct {
	JWT      string
	Endpoint string
	Client   *http.Client
}

type PublishJSONInput struct {
	Name string
	Data any
}

type PublishResult struct {
	CID         string `json:"cid"`
	URI         string `json:"uri"`
	Size        int64  `json:"size,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	IsDuplicate bool   `json:"isDuplicate,omitempty"`
}

type pinataJSONRequest struct {
	PinataMetadata pinataMetadata `json:"pinataMetadata,omitempty"`
	PinataContent  any            `json:"pinataContent"`
}

type pinataMetadata struct {
	Name string `json:"name,omitempty"`
}

type pinataJSONResponse struct {
	IpfsHash    string `json:"IpfsHash"`
	PinSize     int64  `json:"PinSize"`
	Timestamp   string `json:"Timestamp"`
	IsDuplicate bool   `json:"isDuplicate"`
}

func NewPinataClient(jwt string) (*PinataClient, error) {
	jwt = strings.TrimSpace(jwt)
	if jwt == "" {
		return nil, fmt.Errorf("pinata jwt is required")
	}

	return &PinataClient{
		JWT:      jwt,
		Endpoint: defaultPinataJSONEndpoint,
		Client:   http.DefaultClient,
	}, nil
}

func (c *PinataClient) PublishJSON(ctx context.Context, input PublishJSONInput) (PublishResult, error) {
	if c == nil {
		return PublishResult{}, fmt.Errorf("pinata client is nil")
	}

	if strings.TrimSpace(c.JWT) == "" {
		return PublishResult{}, fmt.Errorf("pinata jwt is required")
	}

	endpoint := strings.TrimSpace(c.Endpoint)
	if endpoint == "" {
		endpoint = defaultPinataJSONEndpoint
	}

	httpClient := c.Client
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	body := pinataJSONRequest{
		PinataMetadata: pinataMetadata{
			Name: strings.TrimSpace(input.Name),
		},
		PinataContent: input.Data,
	}

	encoded, err := json.Marshal(body)
	if err != nil {
		return PublishResult{}, fmt.Errorf("failed to encode pinata request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(encoded))
	if err != nil {
		return PublishResult{}, fmt.Errorf("failed to create pinata request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.JWT)

	resp, err := httpClient.Do(req)
	if err != nil {
		return PublishResult{}, fmt.Errorf("failed to upload json to pinata: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return PublishResult{}, fmt.Errorf("failed to read pinata response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return PublishResult{}, fmt.Errorf("pinata upload failed: status=%d body=%s", resp.StatusCode, string(respBody))
	}

	var decoded pinataJSONResponse
	if err := json.Unmarshal(respBody, &decoded); err != nil {
		return PublishResult{}, fmt.Errorf("failed to decode pinata response: %w", err)
	}

	if strings.TrimSpace(decoded.IpfsHash) == "" {
		return PublishResult{}, fmt.Errorf("pinata response missing IpfsHash")
	}

	return PublishResult{
		CID:         decoded.IpfsHash,
		URI:         "ipfs://" + decoded.IpfsHash,
		Size:        decoded.PinSize,
		Timestamp:   decoded.Timestamp,
		IsDuplicate: decoded.IsDuplicate,
	}, nil
}
