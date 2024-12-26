package opensea

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	testnetAPI = "https://testnets-api.opensea.io"
	// rinkebyAPI is already declared
	contractEP = "/api/v1/asset_contract"
	assetEP    = "/api/v1/asset"
)

// Client represents an OpenSea API client
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

type OpenseaClient struct {
	API        string
	APIKey     string
	httpClient *http.Client
}

type errorResponse struct {
	Success bool `json:"success" bson:"success"`
}

func (e errorResponse) Error() string {
	return "Operation unsuccessful"
}

func NewOpenseaMainnet(apiKey string) *OpenseaClient {
	return &OpenseaClient{
		API:        mainnetAPI,
		APIKey:     apiKey,
		httpClient: newHttpClient(),
	}
}

// NewOpenseaRinkeby initializes an Opensea instance for the Rinkeby testnet.
func NewOpenseaRinkeby(apiKey string) *OpenseaClient {
	return &OpenseaClient{
		API:        rinkebyAPI,
		APIKey:     apiKey,
		httpClient: newHttpClient(),
	}
}

// newHttpClient creates a default HTTP client with a timeout.
func newHttpClient() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}

	if c.apiKey != "" {
		req.Header.Set("X-API-KEY", c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
